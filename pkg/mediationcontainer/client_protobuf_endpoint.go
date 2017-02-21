package mediationcontainer

import (
	"github.com/golang/glog"
	goproto "github.com/golang/protobuf/proto"
)

// =====================================================================================================
// Implementation of ProtobufEndpoint to handle all the server protobuf messages sent to the client
type ClientProtobufEndpoint struct {
	Name string
	// Transport used to send and receive messages
	transport ITransport
	// Parser for the message - this will vary with the type of message communication the endpoint is being used for
	messageHandler ProtobufMessage
	// Channel where the endpoint will send the parsed messages
	MessageChannel chan goproto.Message
	closeReceived  bool
	// TODO: add message waiting policy
}


// Create a new instance of the ClientProtobufEndpoint that handles communication
// for a specific message type using the given transport point
func CreateClientProtobufEndpoint(name string, transport ITransport, messageHandler ProtobufMessage, singleMessage bool) ProtobufEndpoint {
	endpoint := &ClientProtobufEndpoint{
		Name: name,

		transport:      transport, // the transport
		MessageChannel: make(chan goproto.Message),
		messageHandler: messageHandler, // the message parser
	}

	glog.Infof("Created Protobuf Endpoint " + endpoint.GetName())
	// Start a Message Handling routine to wait for messages arriving on the transport point
	if singleMessage {
		go endpoint.waitForSingleServerMessage() // TODO: redo using MessageWaiting policy
	} else {
		go endpoint.waitForServerMessage()
	}

	return endpoint
}

func (endpoint *ClientProtobufEndpoint) GetName() string {
	return endpoint.Name
}

func (endpoint *ClientProtobufEndpoint) GetTransport() ITransport {
	return endpoint.transport
}

func (endpoint *ClientProtobufEndpoint) MessageReceiver() chan goproto.Message {
	return endpoint.MessageChannel
}

func (endpoint *ClientProtobufEndpoint) GetMessageHandler() ProtobufMessage {
	return endpoint.messageHandler
}

func (endpoint *ClientProtobufEndpoint) CloseEndpoint() {
	endpoint.closeReceived = true
	glog.Infof("["+endpoint.Name+"] : CLOSING, close=", endpoint.closeReceived)
	//TODO: close the channel
	//close(endpoint.MessageReceiver())
}

func (endpoint *ClientProtobufEndpoint) Send(messageToSend *EndpointMessage) {
	glog.Infof("[" + endpoint.Name + "] : SENDING Protobuf message") // %s", messageToSend.ProtobufMessage)
	// Marshal protobuf message to raw bytes
	msgMarshalled, err := goproto.Marshal(messageToSend.ProtobufMessage) // marshal to byte array
	if err != nil {
		glog.Errorf("[ClientProtobufEndpoint] During Send - marshaling error: ", err)
		return
	}
	// Send using the underlying transport
	tmsg := &TransportMessage{
		RawMsg: msgMarshalled,
	}
	endpoint.transport.Send(tmsg) // TODO: catch any exceptions during send
}

func (endpoint *ClientProtobufEndpoint) waitForServerMessage() {
	glog.Infof("[" + endpoint.Name + "][waitForServerMessage] : Waiting for server request")

	// main loop for listening server message until its message receiver channel is closed.
	for {
		glog.V(2).Infof("["+endpoint.Name+"] : Waiting for server request ...", endpoint.closeReceived)
		if endpoint.closeReceived {
			glog.Errorf("[" + endpoint.Name + "] : Endpoint is closed")
			break
		}
		// Get the message bytes from the transport channel,
		// - this will block till the message appears on the channel
		rawBytes, ok := <-endpoint.transport.RawMessageReceiver()
		if !ok {
			glog.Errorf("[" + endpoint.Name + "] : Transport Message Receiver channel is closed")
			break
			// TODO: initiate reconnection ?
		} else {
			glog.Infof("[" + endpoint.Name + "] : Received message on Transport Receiver channel")
		}

		// Parse the input stream using the registered message handler
		messageHandler := endpoint.GetMessageHandler()
		messageHandler.parse(rawBytes)
		serverMsg := messageHandler.GetMessage()

		if serverMsg == nil {
			glog.Errorf("["+endpoint.Name+"][waitForServerMessage] : Received null message, dropping it")
			continue
		}

		glog.Infof("["+endpoint.Name+"][waitForServerMessage] : Received: %s\n", serverMsg)

		// Put the parsed message on the endpoint's channel
		// - this will block till the upper layer receives this message
		msgChannel := endpoint.MessageReceiver()
		if msgChannel != nil { // checking if the channel was closed before putting the message
			msgChannel <- serverMsg
		}

		glog.V(2).Infof("[" + endpoint.Name + "] : Parsed server message delivered on the message channel, continue to listen from transport ...")
	}
	glog.V(2).Infof("[" + endpoint.Name + "][waitForServerMessage] : DONE, Waiting for server request")
}

func (endpoint *ClientProtobufEndpoint) waitForSingleServerMessage() {
	glog.V(2).Infof("[" + endpoint.Name + "][waitForSingleServerMessage] : Waiting for server response")

	// listen for server message
	// - this will block till the message appears on the channel
	rawBytes := <-endpoint.transport.RawMessageReceiver()

	// Parse the input stream using the registered message handler
	messageHandler := endpoint.GetMessageHandler()
	messageHandler.parse(rawBytes)

	serverMsg := messageHandler.GetMessage()

	glog.Infof("["+endpoint.Name+"][waitForSingleServerMessage] : Received: %s\n", serverMsg)

	// - this will block till the upper layer receives this message
	msgChannel := endpoint.MessageReceiver()
	if msgChannel != nil { // checking if the channel was closed before putting the message
		msgChannel <- serverMsg
	}

	glog.V(2).Infof("[" + endpoint.Name + "][waitForSingleServerMessage] : DONE Waiting for server response")
}

// =====================================================================================

type MessageWaiter interface {
	getMessage(endpoint ProtobufEndpoint) goproto.Message
}

type SingleMessageWaiter struct {
}

func (messageWaiter *SingleMessageWaiter) getMessage(endpoint ProtobufEndpoint) {
	glog.Infof("[" + endpoint.GetName() + "] : ########## Waiting for server request #######")
	// listen for server message
	// - this will block till the message appears on the channel
	transport := endpoint.GetTransport()
	rawBytes := <-transport.RawMessageReceiver()
	//fmt.Printf("[" + endpoint.Name + "][waitForSingleServerMessage] : Received: message from transport channel %s\n", rawBytes)

	// Parse the input stream using the registered message handler
	messageHandler := endpoint.GetMessageHandler()
	err := messageHandler.parse(rawBytes)
	if err != nil {

	}
	serverMsg := messageHandler.GetMessage()

	glog.Infof("["+endpoint.GetName()+"][waitForSingleServerMessage] : Received: %s\n", serverMsg)

	// - this will block till the upper layer receives this message
	msgChannel := endpoint.MessageReceiver()
	if msgChannel != nil { // checking if the channel was closed before putting the message
		msgChannel <- serverMsg
	}
}


type ContinousMessageWaiter struct {
}

func (messageWaiter *ContinousMessageWaiter) getMessage(endpoint ProtobufEndpoint) {
	go func() {

	} ()
}