package mediationcontainer

import (
	"time"

	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// Abstraction to establish session using the specified protocol with the server
// and handle server messages for the different probes in the Mediation Container
type remoteMediationClient struct {
	// All the probes
	allProbes map[string]*ProbeProperties
	// The container info containing the communication config for all the registered probes
	containerConfig *MediationContainerConfig
	// Associated Transport
	Transport ITransport
	// Map of Message Handlers to receive server messages
	MessageHandlers map[RequestType]RequestHandler
	// Channel for receiving responses from the registered probes to be sent to the server
	probeResponseChan chan *proto.MediationClientMessage
}

func CreateRemoteMediationClient(allProbes map[string]*ProbeProperties,
	containerConfig *MediationContainerConfig) *remoteMediationClient {
	remoteMediationClient := &remoteMediationClient{
		MessageHandlers:   make(map[RequestType]RequestHandler),
		allProbes:         allProbes,
		containerConfig:   containerConfig,
		probeResponseChan: make(chan *proto.MediationClientMessage),
	}

	glog.Infof("Created probe Response channel %s\n", remoteMediationClient.probeResponseChan)

	// Create message handlers
	remoteMediationClient.createMessageHandlers(remoteMediationClient.probeResponseChan)

	glog.Infof(" ****** Created RemoteMediationClient")

	return remoteMediationClient
}

// Establish connection with the Turbo server
// - First register probes and targets
// - Then wait for server messages
//
func (remoteMediationClient *remoteMediationClient) Init(probeRegisteredMsg chan bool) {
	// Assert that the probes are registered before starting the handshake ??
	glog.Infof("Probe Registration status channel %s\n", probeRegisteredMsg)

	//// --------- Create WebSocket Transport
	connConfig, err := CreateWebSocketConnectionConfig(remoteMediationClient.containerConfig)
	if err != nil {
		// TODO, extract error handling method.
		glog.Errorf("Initialization of remote mediation client failed, null transport : " + err.Error())
		remoteMediationClient.Stop()
		probeRegisteredMsg <- false
		return
	}

	transport, err := CreateClientWebSocketTransport(connConfig)

	// handle WebSocket creation errors
	if transport == nil || err != nil {
		// TODO, extract error handling method.
		glog.Errorf("Initialization of remote mediation client failed, null transport : " + err.Error())
		remoteMediationClient.Stop()
		probeRegisteredMsg <- false
		return
	}

	remoteMediationClient.Transport = transport

	// -------- Start protocol handler separate thread
	// Initiate protocol to connect to server
	glog.V(2).Infof("START CLIENT PROTOCOL ........")
	done := make(chan bool, 1)
	sdkProtocolHandler := CreateSdkClientProtocolHandler(remoteMediationClient.allProbes, done)
	go sdkProtocolHandler.handleClientProtocol(remoteMediationClient.Transport)

	// TODO: block till message received from the Protocol handler or timeout
	status := <-done

	glog.V(2).Infof("END CLIENT PROTOCOL, Received DONE = ", status)
	// Send registration status to the upper layer
	defer close(probeRegisteredMsg)
	defer close(done)
	probeRegisteredMsg <- status

	if !status {
		glog.Errorf("******* Registration with server failed")
		remoteMediationClient.Stop()
		return
	}

	glog.V(2).Infof("Sent Registration Status on channel %s\n", probeRegisteredMsg)

	// --------- Listen for server messages
	remoteMediationClient.HandleServerMessages(remoteMediationClient.Transport)
}

func (remoteMediationClient *remoteMediationClient) Stop() {
	if remoteMediationClient.Transport != nil {
		remoteMediationClient.Transport.CloseTransportPoint()
	}
	// TODO: stop the go routines for message handling and response callback
}

// ======================== Listen for server messages ===================
// Performs registration, version negotiation , then notifies that the client protobuf endpoint is ready to be created

func (remoteMediationClient *remoteMediationClient) HandleServerMessages(transport ITransport) {
	// Create Protobuf Endpoint to handle server messages
	protoMsg := &MediationRequest{} // parser for the server requests
	endpoint := CreateClientProtobufEndpoint("ServerRequestEndpoint", transport, protoMsg, false)

	// Spawn a new go routine that serves as a Callback for Probes when their response is ready
	glog.V(2).Infof("[" + endpoint.GetName() + "] : Start callback to receive probe responses")
	go remoteMediationClient.probeCallback(endpoint)

	// main loop for listening to server message.
	for {
		// Wait for the server request to be received and parsed by the protobuf endpoint
		glog.V(2).Infof("[" + endpoint.GetName() + "] : Waiting for parsed server message .....")
		parsedMsg, ok := <-endpoint.MessageReceiver() // block till a message appears on the endpoint's message channel
		if !ok {
			glog.Errorf("[" + endpoint.GetName() + "] : Endpoint Receiver channel is closed")
			endpoint.CloseEndpoint()
			return
		}
		glog.Infof("["+endpoint.GetName()+"] : Received: %s\n", parsedMsg)

		// Handler response - find the handler to handle the message
		serverRequest := parsedMsg.ServerMsg
		requestType := getRequestType(serverRequest)
		glog.Infof("["+endpoint.GetName()+"] : serverRequest: %s\n", serverRequest)

		requestHandler := remoteMediationClient.MessageHandlers[requestType]
		if requestHandler == nil {
			glog.Errorf("Cannot find message handler for " + string(requestType) + " Request")
		} else {
			// Dispatch on a new thread
			glog.V(2).Infof("Dispatching received %s %s", serverRequest, parsedMsg)
			// TODO: create MessageOperationRunner to handle this request for a specific message id
			go requestHandler.HandleMessage(serverRequest, remoteMediationClient.probeResponseChan)
			glog.Infof("Message dispatched, waiting for next one")
		}
	}
}

// Send the probe response to the server.
// Probe responses are put on the probeMsgChan by the different message handlers
func (remoteMediationClient *remoteMediationClient) probeCallback(endpoint ProtobufEndpoint) {
	glog.Infof("[probeCallback] Waiting for Probe messages ..... on  %s\n", remoteMediationClient.probeResponseChan)
	for {
		msg, ok := <-remoteMediationClient.probeResponseChan
		if !ok {
			glog.Errorf("[probeCallback] [" + endpoint.GetName() + "] : Probe Callback is closed")
			return
		}
		glog.V(2).Infof("[probeCallback] Received message on probe channel %s\n ", remoteMediationClient.probeResponseChan)
		endMsg := &EndpointMessage{
			ProtobufMessage: msg,
		}
		endpoint.Send(endMsg)

		//select {
		//case msg := <-remoteMediationClient.probeMsgChan :
		//	fmt.Printf("[RemoteMediationClient] [probeCallback] Received message on probe channle %s\n ", msg)
		//	endMsg := &EndpointMessage{
		//		ProtobufMessage: msg,
		//	}
		//	endpoint.Send(endMsg)
		//case <-remoteMediationClient.stop:
		//	return
		//}
	}
}

// ======================== Message Handlers ============================
type RequestType string

const (
	DISCOVERY_REQUEST  RequestType = "Discovery"
	VALIDATION_REQUEST RequestType = "Validation"
	INTERRUPT_REQUEST  RequestType = "Interrupt"
	ACTION_REQUEST     RequestType = "Action"
	UNKNOWN_REQUEST    RequestType = "Unknown"
)

func getRequestType(serverRequest proto.MediationServerMessage) RequestType {
	if serverRequest.GetValidationRequest() != nil {
		return VALIDATION_REQUEST
	} else if serverRequest.GetDiscoveryRequest() != nil {
		return DISCOVERY_REQUEST
	} else if serverRequest.GetActionRequest() != nil {
		return ACTION_REQUEST
	} else if serverRequest.GetInterruptOperation() > 0 {
		return INTERRUPT_REQUEST
	} else {
		return UNKNOWN_REQUEST
	}
}

type RequestHandler interface {
	HandleMessage(serverRequest proto.MediationServerMessage, probeMsgChan chan *proto.MediationClientMessage)
}

func (remoteMediationClient *remoteMediationClient) createMessageHandlers(probeMsgChan chan *proto.MediationClientMessage) {
	allProbes := remoteMediationClient.allProbes
	remoteMediationClient.MessageHandlers[DISCOVERY_REQUEST] = &DiscoveryRequestHandler{
		probes: allProbes,
	}
	remoteMediationClient.MessageHandlers[VALIDATION_REQUEST] = &ValidationRequestHandler{
		probes: allProbes,
	}
	remoteMediationClient.MessageHandlers[INTERRUPT_REQUEST] = &InterruptMessageHandler{
		probes: allProbes,
	}
	remoteMediationClient.MessageHandlers[ACTION_REQUEST] = &ActionMessageHandler{
		probes: allProbes,
	}

	var keys []RequestType
	for k := range remoteMediationClient.MessageHandlers {
		keys = append(keys, k)
	}
	glog.Infof("Created message handlers for server message types : [%s]", keys)
}

// -------------------------------- Discovery Request Handler -----------------------------------
type DiscoveryRequestHandler struct {
	probes map[string]*ProbeProperties
}

func (discReqHandler *DiscoveryRequestHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {
	request := serverRequest.GetDiscoveryRequest()
	probeType := request.ProbeType
	if discReqHandler.probes[*probeType] == nil {
		glog.Errorf("Received: Discovery request for unknown probe type : " + *probeType)
		return
	}

	glog.Infof("Received: discovery for probe type :, %s\n "+*probeType, serverRequest)
	probeProps := discReqHandler.probes[*probeType]
	turboProbe := probeProps.Probe

	msgID := serverRequest.GetMessageID()

	stopCh := make(chan struct{})
	defer close(stopCh)
	go func() {
		for {
			discReqHandler.keepDiscoveryAlive(msgID, probeMsgChan)

			t := time.NewTimer(time.Second * 10)
			select {
			case <-stopCh:
				glog.Infof("******** Cancel Keep alive for msgID ", msgID)
				return
			case <-t.C:
			}
		}

	}()

	var discoveryResponse *proto.DiscoveryResponse
	discoveryResponse = turboProbe.DiscoverTarget(request.GetAccountValue())
	clientMsg := NewClientMessageBuilder(msgID).SetDiscoveryResponse(discoveryResponse).Create()

	// Send the response on the callback channel to send to the server
	//fmt.Printf("[DiscoveryRequestHandler] send discovery response %s on %s\n", clientMsg,  probeMsgChan)
	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.Infof("Sent discovery response for ", clientMsg.GetMessageID())

	// Send empty response to signal completion of discovery
	discoveryResponse = &proto.DiscoveryResponse{}
	clientMsg = NewClientMessageBuilder(msgID).SetDiscoveryResponse(discoveryResponse).Create()

	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.Infof("Sent empty discovery response for ", clientMsg.GetMessageID())

	// Cancel keep alive
	// Note  : Keep alive routine is cancelled when the stopCh is closed at the end of this method
}

// Send the KeepAlive message to server in order to inform server the discovery is stil ongoing. Prevent timeout.
func (discReqHandler *DiscoveryRequestHandler) keepDiscoveryAlive(msgID int32, probeMsgChan chan *proto.MediationClientMessage) {
	keepAliveMsg := new(proto.KeepAlive)
	clientMsg := NewClientMessageBuilder(msgID).SetKeepAlive(keepAliveMsg).Create()

	// Send the response on the callback channel to send to the server
	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.Infof("Sent keepDiscoveryAlive response ", clientMsg.GetMessageID())
}

// -------------------------------- Validation Request Handler -----------------------------------
type ValidationRequestHandler struct {
	probes map[string]*ProbeProperties //TODO: synchronize access to the probes map
}

func (valReqHandler *ValidationRequestHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {
	request := serverRequest.GetValidationRequest()
	probeType := request.ProbeType
	if valReqHandler.probes[*probeType] == nil {
		glog.Errorf("Received: Validation request for unknown probe type : " + *probeType)
		return
	}
	glog.Infof("Received: validation for probe type :, %s\n "+*probeType, serverRequest)
	probeProps := valReqHandler.probes[*probeType]
	turboProbe := probeProps.Probe

	var validationResponse *proto.ValidationResponse
	validationResponse = turboProbe.ValidateTarget(request.GetAccountValue())

	msgID := serverRequest.GetMessageID()
	clientMsg := NewClientMessageBuilder(msgID).SetValidationResponse(validationResponse).Create()

	// Send the response on the callback channel to send to the server
	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.Infof("Sent validation response ", clientMsg.GetMessageID())
}

// -------------------------------- Action Request Handler -----------------------------------
// Message handler that will receive the Action Request for entities in the TurboProbe.
// Action request will be delegated to the right TurboProbe. Multiple ActionProgress and final ActionResult
// responses are sent back to the server.
type ActionMessageHandler struct {
	probes map[string]*ProbeProperties
}

func (actionReqHandler *ActionMessageHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {
	glog.Infof("[ActionMessageHandler] Received: action %s request", serverRequest)
	request := serverRequest.GetActionRequest()
	probeType := request.ProbeType
	if actionReqHandler.probes[*probeType] == nil {
		glog.Errorf("Received: Action request for unknown probe type : ", *probeType)
		return
	}

	glog.Infof("Received: action %s request for probe type :, %s\n ", request.ActionExecutionDTO.ActionType, probeType)
	probeProps := actionReqHandler.probes[*probeType]
	turboProbe := probeProps.Probe

	msgID := serverRequest.GetMessageID()

	worker := NewActionResponseWorker(msgID, turboProbe,
		request.ActionExecutionDTO, request.GetAccountValue(), probeMsgChan)
	worker.start()
}

// Worker Object that will receive multiple action progress responses from the TurboProbe
// before the final result. Action progress and result are sent to the server as responses for the action request.
// It implements the ActionProgressTracker interface.
type ActionResponseWorker struct {
	msgId              int32
	turboProbe         *probe.TurboProbe
	actionExecutionDto *proto.ActionExecutionDTO
	accountValues      []*proto.AccountValue
	probeMsgChan       chan *proto.MediationClientMessage
}

func NewActionResponseWorker(msgId int32, turboProbe *probe.TurboProbe,
	actionExecutionDto *proto.ActionExecutionDTO, accountValues []*proto.AccountValue,
	probeMsgChan chan *proto.MediationClientMessage) *ActionResponseWorker {
	worker := &ActionResponseWorker{
		msgId:              msgId,
		turboProbe:         turboProbe,
		actionExecutionDto: actionExecutionDto,
		accountValues:      accountValues,
		probeMsgChan:       probeMsgChan,
	}
	glog.Infof("**** New ActionResponseProtocolWorker for %s %s %s", msgId, turboProbe, actionExecutionDto.ActionType)
	return worker
}

func (actionWorker *ActionResponseWorker) start() {
	var actionResult *proto.ActionResult
	// Execute the action
	actionResult = actionWorker.turboProbe.ExecuteAction(actionWorker.actionExecutionDto, actionWorker.accountValues, actionWorker)
	clientMsg := NewClientMessageBuilder(actionWorker.msgId).SetActionResponse(actionResult).Create()

	// Send the response on the callback channel to send to the server
	//glog.Infof("[ActionResponseProtocolWorker] send action response %s on %s\n", clientMsg, actionWorker.probeMsgChan)
	actionWorker.probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.Infof("Sent action response for ", clientMsg.GetMessageID())
}

func (actionWorker *ActionResponseWorker) UpdateProgress(actionState proto.ActionResponseState,
	description string, progress int32) {
	// Build ActionProgress
	actionResponse := &proto.ActionResponse{
		ActionResponseState: &actionState,
		ResponseDescription: &description,
		Progress:            &progress,
	}

	actionProgress := &proto.ActionProgress{
		Response: actionResponse,
	}

	clientMsg := NewClientMessageBuilder(actionWorker.msgId).SetActionProgress(actionProgress).Create()
	// Send the response on the callback channel to send to the server
	//glog.Infof("[ActionResponseProtocolWorker] send action progress %s on %s\n", clientMsg, actionWorker.probeMsgChan)
	actionWorker.probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.Infof("Sent action progress for ", clientMsg.GetMessageID())

}

// -------------------------------- Interrupt Request Handler -----------------------------------
type InterruptMessageHandler struct {
	probes map[string]*ProbeProperties
}

func (intMsgHandler *InterruptMessageHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {

	msgID := serverRequest.GetMessageID()
	glog.Infof("******** Received: Interrupt Message for message Id: %d, %s\n ", msgID, serverRequest)
}
