package communication

import  (
	"fmt"
	"time"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

)

// Abstraction to establish session using the specified protocol with the server
// and handle server messages for the different probes in the Mediation Container
type RemoteMediationClient struct {
	// All the probes
	allProbes         map[string]*ProbeProperties
	// The container info for all the registered probes
	containerConfig   *ContainerConfig
	// Associated Transport
	transport         ITransport
	// Map of Message Handlers to receive server messages
	MessageHandlers   map[string]RequestHandler
	// Channel for receiving probe responses
	probeResponseChan chan *proto.MediationClientMessage
}

// TODO: make singleton instance
func CreateRemoteMediationClient(allProbes map[string]*ProbeProperties,
				 containerConfig *ContainerConfig) (*RemoteMediationClient) {
	remoteMediationClient := &RemoteMediationClient{
		MessageHandlers: make(map[string]RequestHandler),
		allProbes: allProbes,
		containerConfig:  containerConfig,
		probeResponseChan: make(chan *proto.MediationClientMessage),
	}

	fmt.Printf("[RemoteMediationClient] Created probe Response channel %s\n", remoteMediationClient.probeResponseChan)

	// Create message handlers
	remoteMediationClient.createMessageHandlers(remoteMediationClient.probeResponseChan)

	fmt.Println("[RemoteMediationClient] ****** Created RemoteMediationClient")

	return remoteMediationClient
}


// Establish connection with the Turbo server
// - First register probes and targets
// - Then wait for server messages
//
func (remoteMediationClient *RemoteMediationClient) Init(probeRegisteredMsg chan bool) {
	// Assert that the probes are registered before starting the handshake ??

	//// --------- Create Websocket Transport
	connConfig := &WebsocketConnectionConfig{
		VmtServerAddress: remoteMediationClient.containerConfig.VmtServerAddress,
		ServerUsername: remoteMediationClient.containerConfig.VmtUserName,
		ServerPassword: remoteMediationClient.containerConfig.VmtPassword,
		IsSecure: remoteMediationClient.containerConfig.IsSecure,
		RequestURI: remoteMediationClient.containerConfig.ApplicationBase,
	}

	remoteMediationClient.transport = CreateClientWebsocketTransport(connConfig)
	// TODO: handle websocket creation errors

	// -------- Start protocol handler separate thread
	// Initiate protocol to connect to server
	fmt.Println("[RemoteMediationClient] START CLIENT PROTOCOL ........")
	done := make(chan bool, 1)
	sdkProtocolHandler := CreateSdkClientProtocolHandler(remoteMediationClient.allProbes, done)

	go sdkProtocolHandler.handleClientProtocol(remoteMediationClient.transport)

	// TODO: block till message received from the Protocol handler or timeout
	status := <-done

	fmt.Println("[RemoteMediationClient] END CLIENT PROTOCOL, Received DONE = " , status)
	// Send registration status to the upper layer
	probeRegisteredMsg <- status

	fmt.Printf("[RemoteMediationClient] Sent Registration Status on channel %s\n", probeRegisteredMsg)

	// --------- Listen for server messages
	remoteMediationClient.handleServerMessages(remoteMediationClient.transport)
}

func (remoteMediationClient *RemoteMediationClient) Stop() {
	remoteMediationClient.transport.CloseTransportPoint()
	// TODO: stop the go routines for message handling and response callback
}

// ======================== Listen for server messages ===================
// Performs registration, version negotiation , then notifies that the client protobuf endpoint is ready to be created

func (remoteMediationClient *RemoteMediationClient) handleServerMessages(transport ITransport) {
	// Create Protobuf Endpoint to handle server messages
	protoMsg := &ServerRequest {}	// parser for the server requests
	endpoint := CreateClientProtobufEndpoint3("ServerRequestEndpoint", transport, protoMsg)
	// Spawn a new go routine that serves as a Callback for Probes when their response is ready
	fmt.Println("[RemoteMediationClient] [" + endpoint.GetName() + "] : Start Client Callback")
	go remoteMediationClient.probeCallback(endpoint)

	// main loop for listening to server message.
	for {
		// Wait for the server request to be received and parsed by the protobuf endpoint
		fmt.Println("[RemoteMediationClient] [" + endpoint.GetName() + "] : Waiting for parsed server message .....")
		serverMsg, ok := <-endpoint.MessageReceiver()	// block till a message appears on the endpoint's message channel
		if !ok {
			fmt.Println("[" + endpoint.GetName() + "] : Endpoint Receiver channel is closed")
			endpoint.CloseEndpoint()
			return
		}
		fmt.Printf("[RemoteMediationClient] [" + endpoint.GetName() + "] : Received: %s\n", serverMsg)

		// Handler response - find the handler to handle the message
		serverRequest := protoMsg.ServerMsg
		//var requestHandler RequestHandler
		if serverRequest.GetValidationRequest() != nil {
			//requestHandler := getMessageHandler("Validation")
			requestHandler := remoteMediationClient.MessageHandlers["Validation"]
			if requestHandler == nil {
				fmt.Println("[RemoteMediationClient] : Cannot find message handler for Validation Request")
			} else {
				// Dispatch on a new thread
				// TODO: create MessageOperationRunner to handle this request for a specific message id
				go requestHandler.HandleMessage(serverRequest, remoteMediationClient.probeResponseChan)
			}
		} else if serverRequest.GetDiscoveryRequest() != nil {
			requestHandler := remoteMediationClient.MessageHandlers["Discovery"]
			if requestHandler == nil {
				fmt.Println("[RemoteMediationClient] : Cannot find message handler for Discovery Request")
			} else {
				// Dispatch on a new thread
				go requestHandler.HandleMessage(serverRequest, remoteMediationClient.probeResponseChan)
			}
		} else if serverRequest.GetActionRequest() != nil {
			fmt.Println("[" + endpoint.GetName() + "] : Received: Action Request")
		} else if serverRequest.GetInterruptOperation() > 0 {
			requestHandler := remoteMediationClient.MessageHandlers["Interrupt"]
			if requestHandler == nil {
				fmt.Println("[RemoteMediationClient] : Cannot find message handler for Interrupt")
			} else {
				// Dispatch on a new thread
				go requestHandler.HandleMessage(serverRequest, remoteMediationClient.probeResponseChan)
			}
		}

		//TODO: // Dispatch on a new thread
		//if (requestHandler != nil) {
		//	go requestHandler.HandleMessage(serverRequest, remoteMediationClient.probeMsgChan)
		//}
		fmt.Println("[RemoteMediationClient] : Message dispatched, waiting for next one")
	}
}

func (remoteMediationClient *RemoteMediationClient) getMessageHandler(msgType string) RequestHandler {
	requestHandler := remoteMediationClient.MessageHandlers[msgType]
	if requestHandler == nil {
		fmt.Println("[RemoteMediationClient] : Cannot find message handler for " + msgType)
		return nil
	} else {
		// Dispatch on a new thread
		return requestHandler
		//go requestHandler.HandleMessage(serverRequest, remoteMediationClient.probeMsgChan)
	}
}

// Send the probe response to the server.
// Probe responses are put on the probeMsgChan by the different message handlers
func (remoteMediationClient *RemoteMediationClient) probeCallback(endpoint ProtobufEndpoint) {
	fmt.Printf("[RemoteMediationClient][probeCallback] Waiting for Probe messages ..... on  %s\n", remoteMediationClient.probeResponseChan)
	for {
		msg, ok := <- remoteMediationClient.probeResponseChan
		if !ok {
			fmt.Println("[RemoteMediationClient][probeCallback] [" + endpoint.GetName() + "] : Probe Callback is closed")
			return
		}
		fmt.Printf("[RemoteMediationClient] [probeCallback] Received message on probe channel %s\n ", remoteMediationClient.probeResponseChan)
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
type RequestHandler interface {
	HandleMessage(serverRequest *proto.MediationServerMessage, probeMsgChan chan *proto.MediationClientMessage)
}

func (remoteMediationClient *RemoteMediationClient) createMessageHandlers(probeMsgChan chan *proto.MediationClientMessage) {
	allProbes := remoteMediationClient.allProbes
	fmt.Println("[RemoteMediationClient] Creating message handlers ...")
	remoteMediationClient.MessageHandlers["Discovery"] = &DiscoveryRequestHandler{
									probes: allProbes,
								}
	remoteMediationClient.MessageHandlers["Validation"] = &ValidationRequestHandler{
									probes: allProbes,
								}
	remoteMediationClient.MessageHandlers["Interrupt"] = &InterruptMessageHandler{
									probes: allProbes,
								}
	remoteMediationClient.MessageHandlers["Action"] = &ActionMessageHandler{
									probes: allProbes,
								}

	var keys []string
	for k := range remoteMediationClient.MessageHandlers {
		keys = append(keys, k)
	}
	fmt.Println("[RemoteMediationClient] Created message handlers for : " , keys)
}


type ActionMessageHandler struct {
	probes map[string]*ProbeProperties
}

func (actionReqHandler *ActionMessageHandler) HandleMessage(serverRequest *proto.MediationServerMessage,
							probeMsgChan chan *proto.MediationClientMessage) {

	msgID := serverRequest.GetMessageID()
	fmt.Printf("************** [ActionMessageHandler] Received: Action Message for message Id: %d, %s\n ", msgID, serverRequest)
}

type InterruptMessageHandler struct {
	probes map[string]*ProbeProperties
}

func (intMsgHandler *InterruptMessageHandler) HandleMessage(serverRequest *proto.MediationServerMessage,
								probeMsgChan chan *proto.MediationClientMessage) {

	msgID := serverRequest.GetMessageID()
	fmt.Printf("************** [InterruptMessageHandler] Received: Interrupt Message for message Id: %d, %s\n ", msgID, serverRequest)
}

type DiscoveryRequestHandler struct {
	probes map[string]*ProbeProperties
}

func (discReqHandler *DiscoveryRequestHandler) HandleMessage(serverRequest *proto.MediationServerMessage,
								probeMsgChan chan *proto.MediationClientMessage) {
	request := serverRequest.GetDiscoveryRequest()
	probeType := request.ProbeType
	if discReqHandler.probes[*probeType] == nil {
		fmt.Println("Received: Discovery request for unknown probe type : " + *probeType)
		return
	}

	fmt.Println("DiscoveryRequestHandler : Found handler for probe " , probeType)
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
				fmt.Println("********** [DiscoveryRequestHandler] Cancel Keep alive for msgID  *********", msgID)
				return
			case <-t.C:
			}
		}

	}()

	var discoveryResponse *proto.DiscoveryResponse
	discoveryResponse = turboProbe.DiscoverTarget(request.GetAccountValue())

	//clientMsg := NewClientMessageBuilder(msgID).SetValidationResponse(validationResponse).Create()
	cb := NewClientMessageBuilder(msgID)
	clientMsg := cb.SetDiscoveryResponse(discoveryResponse).Create()

	//response := &proto.MediationClientMessage_DiscoveryResponse {
	//	DiscoveryResponse: discoveryResponse,
	//}
	//clientMsg := &proto.MediationClientMessage{
	//	MessageID: &msgID,
	//	MediationClientMessage: response,
	//}

	// Send the response on the callback channel to send to the server
	//fmt.Printf("[DiscoveryRequestHandler] send discovery response %s on %s\n", clientMsg,  probeMsgChan)
	probeMsgChan <- clientMsg	// This will block till the channel is ready to receive
	fmt.Println("[DiscoveryRequestHandler] sent discovery response for ", clientMsg.GetMessageID())

	// Send empty response to signal completion of discovery
	discoveryResponse = &proto.DiscoveryResponse{}
	//response = &proto.MediationClientMessage_DiscoveryResponse {
	//	DiscoveryResponse: discoveryResponse,
	//}
	//clientMsg = &proto.MediationClientMessage{
	//	MessageID: &msgID,
	//	MediationClientMessage: response,
	//}

	cb = NewClientMessageBuilder(msgID)
	clientMsg = cb.SetDiscoveryResponse(discoveryResponse).Create()
	probeMsgChan <- clientMsg	// This will block till the channel is ready to receive
	fmt.Println("[DiscoveryRequestHandler] sent empty discovery response for ", clientMsg.GetMessageID())

	// Cancel keep alive
	// Note  : Keep alive routine is cancelled when the stopCh is closed at the end of this method
}

// Send the KeepAlive message to server in order to inform server the discovery is stil ongoing. Prevent timeout.
func (discReqHandler *DiscoveryRequestHandler) keepDiscoveryAlive(msgID int32, probeMsgChan chan *proto.MediationClientMessage) {
	fmt.Println("Keep Alive is called for message with ID: %d", msgID)

	keepAliveMsg := new(proto.KeepAlive)
	cb := NewClientMessageBuilder(msgID)
	clientMsg := cb.SetKeepAlive(keepAliveMsg).Create()

	//response := &proto.MediationClientMessage_KeepAlive {
	//	KeepAlive: 	keepAliveMsg,
	//}
	//clientMsg := &proto.MediationClientMessage{
	//	MessageID: &msgID,
	//	MediationClientMessage: response,
	//}

	// Send the response on the callback channel to send to the server
	//fmt.Printf("[keepDiscoveryAlive] send keepDiscoveryAlive response %s on %s\n", clientMsg,  probeMsgChan)
	probeMsgChan <- clientMsg	// This will block till the channel is ready to receive
	fmt.Println("[keepDiscoveryAlive] sent keepDiscoveryAlive response ", clientMsg.GetMessageID())
}

type ValidationRequestHandler struct {
	probes map[string]*ProbeProperties	//TODO: synchronize access to the probes map
}

func (valReqHandler *ValidationRequestHandler) HandleMessage(serverRequest *proto.MediationServerMessage,
								probeMsgChan chan *proto.MediationClientMessage) {
	request := serverRequest.GetValidationRequest()
	probeType := request.ProbeType
	if valReqHandler.probes[*probeType] == nil {
		fmt.Println("Received: Validation request for unknown probe type : " + *probeType)
		return
	}
	fmt.Printf("[ValidationRequestHandler] Received: Validation for probe type :, %s\n " + *probeType, serverRequest)
	probeProps := valReqHandler.probes[*probeType]
	turboProbe := probeProps.Probe
	var validationResponse *proto.ValidationResponse
	validationResponse = turboProbe.ValidateTarget(request.GetAccountValue())

	// validationResponse := probeInterface.Validate(request.GetAccountValue())

	msgID := serverRequest.GetMessageID()
	//clientMsg := NewClientMessageBuilder(msgID).SetValidationResponse(validationResponse).Create()
	cb := NewClientMessageBuilder(msgID)
	clientMsg := cb.SetValidationResponse(validationResponse).Create()

	//response := &proto.MediationClientMessage_ValidationResponse {
	//	ValidationResponse: 	validationResponse,
	//}
	//clientMsg := &proto.MediationClientMessage{
	//	MessageID: &msgID,
	//	MediationClientMessage: response,
	//}

	// Send the response on the callback channel to send to the server
	//fmt.Printf("[ValidationRequestHandler] send validation response %s on %s\n", clientMsg,  probeMsgChan)

	probeMsgChan <- clientMsg	// This will block till the channel is ready to receive
	fmt.Println("[ValidationRequestHandler] sent validation response ", clientMsg.GetMessageID())
}

