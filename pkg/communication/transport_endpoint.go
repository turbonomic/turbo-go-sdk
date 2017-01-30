package communication

import "fmt"

// Transport endpoint that sends and receives raw message bytes
type ITransport interface {
	CloseTransportPoint()
	Send(messageToSend *TransportMessage)
	RawMessageReceiver() chan []byte
}

type TransportEventHandler interface {
	EventReceiver() chan []byte
	OnClose()
	OnMessage(msgReceived []byte)
}

type TransportMessage struct {
	RawMsg []byte
}

// =====================================================================================

// Connection parameters for the websocket
type WebsocketConnectionConfig struct {
	IsSecure         bool
	VmtServerAddress string
	LocalAddress     string
	ServerUsername   string
	ServerPassword   string
	RequestURI       string
}

func CreateWebsocketTransportPointUrl(connConfig *WebsocketConnectionConfig) string {
	protocol := "ws"
	if connConfig.IsSecure {
		protocol = "wss"
	}

	vmtServerUrl := protocol + "://" + connConfig.VmtServerAddress + connConfig.RequestURI
	fmt.Println("######### [CreateTransportPointUrl] Created Webscoket URL : ", vmtServerUrl)
	return vmtServerUrl
}


