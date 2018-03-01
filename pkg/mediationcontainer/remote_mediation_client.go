package mediationcontainer

import (
	"time"

	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"github.com/golang/glog"
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
	MessageHandlers  map[RequestType]RequestHandler
	stopMsgHandlerCh chan struct{}

	// Channel for receiving responses from the registered probes to be sent to the server
	probeResponseChan chan *proto.MediationClientMessage
	// Channel to stop the mediation client and the underlying transport and message handling
	stopMediationClientCh  chan struct{}
	stoppedMediationClient bool // a flag indicating whether the stopMediationClientCh is stopped or not
}

func CreateRemoteMediationClient(allProbes map[string]*ProbeProperties,
	containerConfig *MediationContainerConfig) *remoteMediationClient {
	remoteMediationClient := &remoteMediationClient{
		MessageHandlers:       make(map[RequestType]RequestHandler),
		allProbes:             allProbes,
		containerConfig:       containerConfig,
		probeResponseChan:     make(chan *proto.MediationClientMessage),
		stopMediationClientCh: make(chan struct{}),
	}

	glog.V(4).Infof("Created channels : probeResponseChan %s, stopMediationClientCh %s\n",
		remoteMediationClient.probeResponseChan, remoteMediationClient.stopMediationClientCh)

	// Create message handlers
	remoteMediationClient.createMessageHandlers(remoteMediationClient.probeResponseChan)

	glog.V(2).Infof("Created remote mediation client")

	return remoteMediationClient
}

func (rclient *remoteMediationClient) Init(probeRegisteredMsg chan bool) {
	connConfig, err := CreateWebSocketConnectionConfig(rclient.containerConfig)
	if err != nil {
		glog.Errorf("Failed to start RemoteMediationClient: failed to create websocket config: %v", err)
		probeRegisteredMsg <- false
		return
	}

	firstTime := true
	for {
		rclient.initTransport(connConfig)
		flag := rclient.protocolHandShake()
		if firstTime {
			probeRegisteredMsg <- flag
			firstTime = false
		}

		if !flag {
			glog.Errorf("Protocol hand shake failed, exiting ...")
			// MediationContainer will call rclient.Stop() later
			//rclient.Stop()
			return
		}

		rclient.stoppedMediationClient = false
		rclient.HandleServerRequests()
		if rclient.stoppedMediationClient {
			glog.V(2).Infof("RemoteMediationClient is stopped, exiting ...")
			return
		}
	}
}

func (rclient *remoteMediationClient) HandleServerRequests() {
	go rclient.RunServerMessageHandler(rclient.Transport)

	select {
	case <-rclient.stopMediationClientCh:
		glog.V(1).Info("RemoteMediationClient is stopped. quitting ...")
		return
	case <-rclient.Transport.NotifyClosed():
		glog.V(1).Info("RemoteMediationClient transport is closed, starting reconnection ...")
		close(rclient.stopMsgHandlerCh)
		return
	}
}

// protocolHandShake will return only when connection has been built, or proposed protocol version is not accepted.
// TODO: add an channel indicating to stop this process
func (rclient *remoteMediationClient) protocolHandShake() bool {
	du := time.Second * 30
	for {
		//1. build websocket connection
		transport := rclient.Transport
		if err := transport.Connect(); err != nil {
			glog.Errorf("Not able to build websocket connection: %v", err)
			glog.Errorf("Will re-try in %v seconds.", du)
			time.Sleep(du)
			continue
		}

		//2. negotiate version
		protocolHandler := CreateSdkClientProtocolHandler(rclient.allProbes, rclient.containerConfig.Version)
		flag, err := protocolHandler.handleClientProtocol(rclient.Transport)
		if err == nil {
			if !flag {
				return false
			}
			return true
		}

		//3. err != nil
		glog.Errorf("Protocol handshake error: %v", err)
		glog.Errorf("Will retry protocol handshake in %v seconds.", du)
		transport.CloseTransportPoint()
		time.Sleep(du)
	}
}

// Stop the remote mediation client by closing the underlying transport and message handler routines
func (rclient *remoteMediationClient) Stop() {
	if rclient.stoppedMediationClient {
		glog.Errorf("stopping a stopped remoteMediationClient.")
		return
	}
	close(rclient.stopMediationClientCh)
	rclient.stoppedMediationClient = true

	rclient.stopTransport()
}

// ======================== Listen for server messages ===================
func (rclient *remoteMediationClient) initTransport(conf *WebSocketConnectionConfig) {
	if rclient.Transport != nil && !rclient.Transport.IsClosed() {
		glog.Errorf("websocket transport is not closed.")
	}

	rclient.Transport = CreateClientWebSocketTransport(conf)
	rclient.stopMsgHandlerCh = make(chan struct{})
}

//This can be called in two cases:
// (1) From upper layer: mediationClient tries to stop everything;
func (rclient *remoteMediationClient) stopTransport() {
	if rclient.Transport != nil {
		rclient.Transport.CloseTransportPoint()
		rclient.Transport = nil
	}
}

// Checks for incoming server messages received by the ProtoBuf endpoint created to handle server requests
func (rclient *remoteMediationClient) RunServerMessageHandler(transport ITransport) {
	glog.V(2).Infof("[handleServerMessages] %s : ENTER  ", time.Now())

	// Create Protobuf Endpoint to handle server messages
	protoMsg := &MediationRequest{} // parser for the server requests
	endpoint := CreateClientProtoBufEndpoint("ServerRequestEndpoint", transport, protoMsg, false)
	defer endpoint.CloseEndpoint()
	logPrefix := "[handleServerMessages][" + endpoint.GetName() + "] : "

	// Spawn a new go routine that serves as a Callback for Probes when their response is ready
	go rclient.runProbeCallback(endpoint) // this also exits using the stopMsgHandlerCh

	// main loop for listening to server message.
	for {
		glog.V(2).Infof(logPrefix + "waiting for parsed server message .....") // make debug
		// Wait for the server request to be received and parsed by the protobuf endpoint
		select {
		case <-rclient.stopMediationClientCh:
			glog.V(1).Infof(logPrefix + "Exit routine because of stopMediationClient signal. ")
			return
		case <-rclient.stopMsgHandlerCh:
			glog.V(1).Infof(logPrefix + "Exit routine ***************")
			return
		case parsedMsg, ok := <-endpoint.MessageReceiver(): // block till a message appears on the endpoint's message channel
			if !ok {
				glog.Errorf(logPrefix + "endpoint message channel is closed")
				break // return or continue ?
			}
			glog.V(3).Infof(logPrefix+"received: %++v\n", parsedMsg)

			// Handler response - find the handler to handle the message
			serverRequest := parsedMsg.ServerMsg
			requestType := getRequestType(serverRequest)

			requestHandler := rclient.MessageHandlers[requestType]
			if requestHandler == nil {
				glog.Errorf(logPrefix + "cannot find message handler for request type " + string(requestType))
			} else {
				// Dispatch on a new thread
				// TODO: create MessageOperationRunner to handle this request for a specific message id
				go requestHandler.HandleMessage(serverRequest, rclient.probeResponseChan)
				glog.V(2).Infof(logPrefix + "message dispatched, waiting for next one")
			}
		} //end select
	} //end for

}

// Run probe callback to the probe response to the server.
// Probe responses put on the probeResponseChan by the different message handlers are sent to the server
func (remoteMediationClient *remoteMediationClient) runProbeCallback(endpoint ProtobufEndpoint) {
	glog.V(4).Infof("[runProbeCallback] %s : ENTER  ", time.Now())
	for {
		glog.V(4).Infof("[probeCallback] waiting for probe responses")
		select {
		case <-remoteMediationClient.stopMsgHandlerCh:
			glog.V(2).Infof("[probeCallback] Exit routine *************")
			return
		case msg := <-remoteMediationClient.probeResponseChan:
			glog.V(3).Infof("[probeCallback] received response on probe channel %v\n ", remoteMediationClient.probeResponseChan)
			endMsg := &EndpointMessage{
				ProtobufMessage: msg,
			}
			endpoint.Send(endMsg)
		} // end select
	}
	glog.V(4).Infof("[probeCallback] DONE")
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
	glog.V(4).Infof("Created message handlers for server message types : [%s]", keys)
}

// -------------------------------- Discovery Request Handler -----------------------------------
type DiscoveryRequestHandler struct {
	probes map[string]*ProbeProperties
}

func (discReqHandler *DiscoveryRequestHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {
	request := serverRequest.GetDiscoveryRequest()
	probeType := request.ProbeType

	probeProps, exist := discReqHandler.probes[*probeType]
	if !exist {
		glog.Errorf("Received: discovery request for unknown probe type: %s", *probeType)
		return
	}
	glog.V(3).Infof("Received: discovery for probe type: %s", *probeType)

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
				glog.V(4).Infof("Cancel keep alive for msgID ", msgID)
				return
			case <-t.C:
			}
		}

	}()

	accountValues := request.GetAccountValue()
	var discoveryResponse *proto.DiscoveryResponse
	switch requestType := request.GetDiscoveryType(); requestType {
	case proto.DiscoveryType_FULL:
		discoveryResponse = turboProbe.DiscoverTarget(accountValues)
	case proto.DiscoveryType_INCREMENTAL:
		discoveryResponse = turboProbe.DiscoverTargetIncremental(accountValues)
	case proto.DiscoveryType_PERFORMANCE:
		discoveryResponse = turboProbe.DiscoverTargetPerformance(accountValues)
	default:
		discoveryResponse = turboProbe.DiscoverTarget(accountValues)
	}

	clientMsg := NewClientMessageBuilder(msgID).SetDiscoveryResponse(discoveryResponse).Create()

	// Send the response on the callback channel to send to the server
	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.V(3).Infof("Sent discovery response for %d:%s", clientMsg.GetMessageID(), request.GetDiscoveryType())

	// Send empty response to signal completion of discovery
	discoveryResponse = &proto.DiscoveryResponse{}
	clientMsg = NewClientMessageBuilder(msgID).SetDiscoveryResponse(discoveryResponse).Create()

	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.V(2).Infof("Discovery has finished for %d:%s", clientMsg.GetMessageID(), request.GetDiscoveryType())

	// Cancel keep alive
	// Note  : Keep alive routine is cancelled when the stopCh is closed at the end of this method
	// when the discovery response is out on the probeMsgCha
}

// Send the KeepAlive message to server in order to inform server the discovery is stil ongoing. Prevent timeout.
func (discReqHandler *DiscoveryRequestHandler) keepDiscoveryAlive(msgID int32, probeMsgChan chan *proto.MediationClientMessage) {
	keepAliveMsg := new(proto.KeepAlive)
	clientMsg := NewClientMessageBuilder(msgID).SetKeepAlive(keepAliveMsg).Create()

	// Send the response on the callback channel to send to the server
	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.V(3).Infof("Sent keep alive response %d", clientMsg.GetMessageID())
}

// -------------------------------- Validation Request Handler -----------------------------------
type ValidationRequestHandler struct {
	probes map[string]*ProbeProperties //TODO: synchronize access to the probes map
}

func (valReqHandler *ValidationRequestHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {
	request := serverRequest.GetValidationRequest()
	probeType := request.ProbeType
	probeProps, exist := valReqHandler.probes[*probeType]
	if !exist {
		glog.Errorf("Received: validation request for unknown probe type : %s", *probeType)
		return
	}
	glog.V(3).Infof("Received: validation for probe type: %s\n ", *probeType)
	turboProbe := probeProps.Probe

	var validationResponse *proto.ValidationResponse
	validationResponse = turboProbe.ValidateTarget(request.GetAccountValue())

	msgID := serverRequest.GetMessageID()
	clientMsg := NewClientMessageBuilder(msgID).SetValidationResponse(validationResponse).Create()

	// Send the response on the callback channel to send to the server
	probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.V(3).Infof("Sent validation response %d", clientMsg.GetMessageID())
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
	glog.V(4).Infof("[ActionMessageHandler] Received: action %s request", serverRequest)
	request := serverRequest.GetActionRequest()
	probeType := request.ProbeType
	if actionReqHandler.probes[*probeType] == nil {
		glog.Errorf("Received: Action request for unknown probe type : %s", *probeType)
		return
	}

	glog.V(3).Infof("Received: action %s request for probe type: %s\n ",
		request.ActionExecutionDTO.ActionType, *probeType)
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
	glog.V(4).Infof("New ActionResponseProtocolWorker for %s %s %s", msgId, turboProbe,
		actionExecutionDto.ActionType)
	return worker
}

func (actionWorker *ActionResponseWorker) start() {
	var actionResult *proto.ActionResult
	// Execute the action
	actionResult = actionWorker.turboProbe.ExecuteAction(actionWorker.actionExecutionDto, actionWorker.accountValues, actionWorker)
	clientMsg := NewClientMessageBuilder(actionWorker.msgId).SetActionResponse(actionResult).Create()

	// Send the response on the callback channel to send to the server
	actionWorker.probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.V(3).Infof("Sent action response for %d.", clientMsg.GetMessageID())
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
	actionWorker.probeMsgChan <- clientMsg // This will block till the channel is ready to receive
	glog.V(3).Infof("Sent action progress for %d.", clientMsg.GetMessageID())

}

// -------------------------------- Interrupt Request Handler -----------------------------------
type InterruptMessageHandler struct {
	probes map[string]*ProbeProperties
}

func (intMsgHandler *InterruptMessageHandler) HandleMessage(serverRequest proto.MediationServerMessage,
	probeMsgChan chan *proto.MediationClientMessage) {

	msgID := serverRequest.GetMessageID()
	glog.V(3).Infof("Received: Interrupt Message for message ID: %d, %s\n ", msgID, serverRequest)
}
