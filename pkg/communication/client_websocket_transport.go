package communication

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"encoding/base64"

	"golang.org/x/net/websocket"

	"github.com/golang/glog"
)


// Implementation of the ITransport for websocket communication to send and receive serialized protobuf message bytes
type ClientWebsocketTransport struct {
	//VmtServerAddress string
	//LocalAddress     string
	//ServerUsername   string
	//ServerPassword   string
	ws               *websocket.Conn	// created during registration
	inputStream 	chan []byte
	closeReceived bool
}

// Instantiate a new ClientWebsocketTransport endpoint for the client
// Websocket connection is established with the server
func CreateClientWebsocketTransport (connConfig *WebsocketConnectionConfig) (*ClientWebsocketTransport) {
	transport := &ClientWebsocketTransport{
		ws: newWebsocketConnection(connConfig),
		inputStream: make(chan []byte),
	}

	// Message handler for received messages
	go transport.waitForServerMessage()
	return transport
}

// Create the websocket connection and establish session with the server
func  newWebsocketConnection(connConfig *WebsocketConnectionConfig) (*websocket.Conn) {
	// Websocket URL
	vmtServerUrl := CreateWebsocketTransportPointUrl(connConfig)
	//
	localAddr := "http://127.0.0.1"			//TODO: required in url format, but ip is don't care for us
	glog.V(3).Infof("Dial Server: %s", vmtServerUrl)

	config, err := websocket.NewConfig(vmtServerUrl, localAddr)
	if err != nil {
		glog.Fatal(err)
	}
	usrpasswd := []byte(connConfig.ServerUsername + ":" + connConfig.ServerPassword)

	config.Header = http.Header{
		"Authorization": {"Basic " + base64.StdEncoding.EncodeToString(usrpasswd)},
	}
	config.TlsConfig = &tls.Config{InsecureSkipVerify: true}
	webs, err := websocket.DialConfig(config)

	if err != nil {
		glog.Error(err)
		if webs == nil {
			glog.Error("The websocket is null, reset")
		}
		closeWebsocket(webs)
	}

	glog.V(3).Infof("CREATED WEBSOCKET : %s", vmtServerUrl)
	fmt.Println("[ClientWebsocketTransport] CREATED WEBSOCKET : %s+v", webs)
	return webs
}

func (clientTransport *ClientWebsocketTransport) RawMessageReceiver() chan []byte {
	return clientTransport.inputStream
}

// Close the Websocket Transport point
func (clientTransport *ClientWebsocketTransport) CloseTransportPoint() {
	fmt.Println("[ClientWebsocketTransport] Closing Transport endpoint and channel ....")
	clientTransport.closeReceived = true
	close(clientTransport.RawMessageReceiver())
	closeWebsocket(clientTransport.ws)
}
// Close the websocket connection
func  closeWebsocket(wsConn *websocket.Conn) {
	//TODO:
	wsConn.Close()
	wsConn = nil
}

// Send serialized protobuf message bytes
func (clientTransport *ClientWebsocketTransport) Send(messageToSend *TransportMessage) {
	//fmt.Printf("[ClientWebsocketTransport] SEND %s\n", messageToSend.RawMsg)
	if clientTransport.closeReceived {
		fmt.Println("[ClientWebsocketTransport] Transport endpoint is closed")
		return
	}

	if clientTransport.ws == nil {
		fmt.Println("[ClientWebsocketTransport] web socket is nil")
		return
	}
	if messageToSend.RawMsg == nil {
		fmt.Println("[ClientWebsocketTransport] marshalled msg is nil")
		return
	}
	if clientTransport.ws.IsClientConn() {
		fmt.Println("[ClientWebsocketTransport] Sending message on client transport %+v", clientTransport.ws)
	}
	err := websocket.Message.Send(clientTransport.ws, messageToSend.RawMsg)
	// _, err  := clientTransport.ws.Write(messageToSend.RawMsg)	// issues using this call for protobuf messages
	if err != nil {
		fmt.Println("[WebSocketCommunicator] Error sending message on client transport ", err)
		//TODO: re-establish connection when error ?
		// TODO: throw exception to the upper layer
		return
	}
	fmt.Println("[ClientWebsocketTransport] Successfully sent message on client transport %s",messageToSend )
}

// Receives serialized protobuf message bytes
func (clientTransport *ClientWebsocketTransport) waitForServerMessage() {
	glog.Info("ClientWebsocketTransport] : Waiting for server response")
	fmt.Println("[ClientWebsocketTransport] : ######### Waiting for server response #######")

	// main loop for listening server message.
	for {

		fmt.Println("[ClientWebsocketTransport] Waiting for server message ...")
		var data []byte = make([]byte, 1024)
		error := websocket.Message.Receive(clientTransport.ws, &data)
		if error != nil {
			fmt.Println("[ClientWebsocketTransport] Error during receive ", error)
		} else {
			fmt.Printf("[ClientWebsocketTransport] Received message on websocket %s\n")
		}
		msgChannel := clientTransport.RawMessageReceiver()
		msgChannel <- data
		//TODO: this read should be accumulative - see onMessageReceived in AbstractWebsocketTransport
		fmt.Println("[ClientWebsocketTransport] Delivered Raw Message on the channel, continue listening for server messages...")
	}
}

//// Should attach itself as event handler for the data received on the websocket
//func (clientTransport *WebSocketCommunicator) onMessageReceived(messageRcvd []byte) {	//*TransportMessage) {
//	for _, eh := range clientTransport.eventHandlers {
//		fmt.Printf("[WebSocketClientTransport] onMessageReceived : %s %s\n", eh, messageRcvd)
//		eh.OnMessage(messageRcvd)
//	}
//}

//==============================================

