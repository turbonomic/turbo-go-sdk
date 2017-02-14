package communication


// Transport endpoint that sends and receives raw message bytes
type ITransport interface {
	CloseTransportPoint()
	Send(messageToSend *TransportMessage)
	RawMessageReceiver() chan []byte
}

type TransportMessage struct {
	RawMsg []byte
}

//type TransportMessage []byte


