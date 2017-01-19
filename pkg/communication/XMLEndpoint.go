package communication

// Endpoint to handle communication of a particular protobuf message type with the server
type XMLEndpoint interface {
	GetName() string
	CloseEndpoint()
	Send(messageToSend *XMLEndpointMessage)
	GetMessageHandler() XMLMessage
	MessageReceiver() chan string
	//AddEventHandler(handler EndpointEventHandler)
}


type XMLEndpointMessage struct {
	XMLMessage string
}

// =====================================================================================
// Parser interface for different server messages
type XMLMessage interface {
	parse(rawMsg []byte)
	GetMessage() string
}

