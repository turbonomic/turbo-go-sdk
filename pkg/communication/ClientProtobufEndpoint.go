package communication

import (
	"fmt"

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
	closeReceived bool
	// TODO: add message waiting policy
}

// Create a new instance of the ClientProtobufEndpoint that handles communication
// for a specific message type using the given transport point
func CreateClientProtobufEndpoint2(name string, transport ITransport, messageHandler ProtobufMessage) (ProtobufEndpoint) {
	endpoint := &ClientProtobufEndpoint {
		Name: name,

		transport: transport,				// the transport
		MessageChannel: make(chan goproto.Message),
		messageHandler: messageHandler,			// the message parser
	}

	fmt.Println("[ProtobufEndpoint] Created endpoint " + endpoint.GetName())
	// Start a Message Handling routine to wait for messages arriving on the transport point
	go endpoint.waitForSingleServerMessage()	// TODO: redo using MessageWaiting policy
	return endpoint
}

// TODO:
func CreateClientProtobufEndpoint3(name string, transport ITransport, messageHandler ProtobufMessage) (ProtobufEndpoint) {
	endpoint := &ClientProtobufEndpoint {
		Name: name,
		transport: transport,				// the transport
		MessageChannel: make(chan goproto.Message),	// the unbuffered channel where parsed messages are sent
								// this channel will block until the message is received
								// by the top layer
		messageHandler: messageHandler,			// the message parser
	}

	fmt.Println("[ProtobufEndpoint] Created endpoint " + endpoint.GetName())
	// Start a Message Handling routine to wait for messages arriving on the transport point
	go endpoint.waitForServerMessage() 	 // TODO: redo using MessageWaiting policy
	return endpoint
}

func (endpoint *ClientProtobufEndpoint) GetName() string {
	return endpoint.Name
}

func (endpoint *ClientProtobufEndpoint) MessageReceiver() chan goproto.Message {
	return endpoint.MessageChannel
}

func (endpoint *ClientProtobufEndpoint) GetMessageHandler() ProtobufMessage {
	return endpoint.messageHandler
}

func (endpoint *ClientProtobufEndpoint) CloseEndpoint() {
	endpoint.closeReceived = true
	fmt.Println("[" + endpoint.Name + "] : CLOSING, close=", endpoint.closeReceived)
	//TODO: close the channel
	//close(endpoint.MessageReceiver())
}

func (endpoint *ClientProtobufEndpoint) Send(messageToSend *EndpointMessage) {
	fmt.Println("[" + endpoint.Name + "] : SENDING Protobuf message")	// %s", messageToSend.ProtobufMessage)
	// Marshal protobuf message to raw bytes
	msgMarshalled, err := goproto.Marshal(messageToSend.ProtobufMessage)	// marshal to byte array
	if err != nil {
		glog.Fatal("[ClientProtobufEndpoint] During Send - marshaling error: ", err)
		return
	}
	// Send using the underlying transport
	tmsg := &TransportMessage {
		RawMsg: msgMarshalled,
	}
	endpoint.transport.Send(tmsg)		// TODO: catch any exceptions during send
}

func (endpoint *ClientProtobufEndpoint) waitForServerMessage() {
	glog.Info("[" + endpoint.Name + "] : ########## Waiting for server response #######")
	fmt.Println("[" + endpoint.Name + "][waitForServerMessage] : ########## Waiting for server response #######")

	// main loop for listening server message until its message receiver channel is closed.
	for {
		fmt.Println("[" + endpoint.Name + "] : Waiting for server message ...", endpoint.closeReceived)
		if endpoint.closeReceived {
			fmt.Println("[" + endpoint.Name + "] : Endpoint is closed")
			break
		}
		// Get the message bytes from the transport channel,
		// - this will block till the message appears on the channel
		rawBytes, ok := <-endpoint.transport.RawMessageReceiver()
		if !ok {
			fmt.Println("[" + endpoint.Name + "] : Transport Message Receiver channel is closed")
			break
			// TODO: initiate reconnection ?
		} else {
			fmt.Println("[" + endpoint.Name + "] : Received message on Transport Receiver channel")
		}
		//fmt.Printf("[" + endpoint.Name + "][waitForServerMessage] : Received: message from transport channel %s \n", rawBytes)

		// Parse the input stream using the registered message handler
		messageHandler := endpoint.GetMessageHandler()
		messageHandler.parse(rawBytes)
		serverMsg := messageHandler.GetMessage()

		fmt.Printf("[" + endpoint.Name + "][waitForServerMessage] : Received: %s\n", serverMsg)

		// Put the parsed message on the endpoint's channel
		// - this will block till the upper layer receives this message
		msgChannel := endpoint.MessageReceiver()
		if msgChannel != nil {	// checking if the channel was closed before putting the message
			msgChannel <- serverMsg
		}

		fmt.Println("[" + endpoint.Name + "] : Parsed server message delivered on the message channel, continue to listen from transport ...")
	}
	fmt.Println("[" + endpoint.Name + "][waitForServerMessage] : DONE, Waiting for server response")
}

func (endpoint *ClientProtobufEndpoint) waitForSingleServerMessage() {
	glog.Info("[" + endpoint.Name + "][waitForSingleServerMessage] : ########## Waiting for server response #######")
	fmt.Println("[" + endpoint.Name + "][waitForSingleServerMessage] : ########## Waiting for server response #######")

	// listen for server message
	// - this will block till the message appears on the channel
	rawBytes := <-endpoint.transport.RawMessageReceiver()
	//fmt.Printf("[" + endpoint.Name + "][waitForSingleServerMessage] : Received: message from transport channel %s\n", rawBytes)

	// Parse the input stream using the registered message handler
	messageHandler := endpoint.GetMessageHandler()
	messageHandler.parse(rawBytes)
	serverMsg := messageHandler.GetMessage()	//endpoint.ParseFromData(rawBytes)

	fmt.Printf("[" + endpoint.Name + "][waitForSingleServerMessage] : Received: %s\n", serverMsg)

	// - this will block till the upper layer receives this message
	msgChannel := endpoint.MessageReceiver()
	if msgChannel != nil {	// checking if the channel was closed before putting the message
		msgChannel <- serverMsg
	}

	fmt.Println("[" + endpoint.Name + "][waitForSingleServerMessage] : DONE Waiting for server response")
}
// =====================================================================================

//type EndpointEventHandler interface {
//	EventReceiver() chan goproto.Message
//	OnClose()
//	OnMessage(protoMsg goproto.Message)
//}

//func (endpoint *ClientProtobufEndpoint) AddEventHandler(eventHandler EndpointEventHandler) {
//	fmt.Println("[ClientProtobufEndpoint] Adding Event Handler ...")
//	endpoint.eventHandlers = append(endpoint.eventHandlers, eventHandler)
//	fmt.Println("[ClientProtobufEndpoint] Number of Event Handlers ", len(endpoint.eventHandlers))
//}


// =====================================================================================
type MessageWaiter interface {
	getMessage(endpoint ProtobufEndpoint) goproto.Message
}

type SingleMessageWaiter struct {
}

func (messageWaiter *SingleMessageWaiter) getMessage(endpoint ProtobufEndpoint) {
	fmt.Println("[" + endpoint.GetName() + "] : ########## Waiting for server response #######")

	serverMsg := <-endpoint.MessageReceiver()
	fmt.Printf("[" + endpoint.GetName() + "] : Received: %s\n", serverMsg)
}
