package mediationcontainer

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"
	"net/url"
)

type WebSocketConnectionConfig MediationContainerConfig

func CreateWebSocketConnectionConfig(connConfig *MediationContainerConfig) (*WebSocketConnectionConfig, error) {
	_, err := url.ParseRequestURI(connConfig.LocalAddress)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse local URL from %s when create WebSocketConnectionConfig.", connConfig.LocalAddress)
	}
	// Change URL scheme from ws to http or wss to https.
	serverURL, err := url.ParseRequestURI(connConfig.TurboServer)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse turboServer URL from %s when create WebSocketConnectionConfig.", connConfig.TurboServer)
	}
	switch serverURL.Scheme {
	case "http":
		serverURL.Scheme = "ws"
	case "https":
		serverURL.Scheme = "wss"
	}
	wsConfig := WebSocketConnectionConfig(*connConfig)
	wsConfig.TurboServer = serverURL.String()
	return &wsConfig, nil
}

// ===================================================================================================================
// Implementation of the ITransport for websocket communication to send and receive serialized protobuf message bytes
type ClientWebSocketTransport struct {
	ws            *websocket.Conn // created during registration
	inputStream   chan []byte
	closeReceived bool
}

// Instantiate a new ClientWebsocketTransport endpoint for the client
// Websocket connection is established with the server
func CreateClientWebSocketTransport(connConfig *WebSocketConnectionConfig) (*ClientWebSocketTransport, error) {
	websocket, err := newWebSocketConnection(connConfig) // will be nil if the server is not connected

	if err != nil {
		return nil, err
	}

	transport := &ClientWebSocketTransport{
		ws:          websocket, //newWebSocketConnection(connConfig),
		inputStream: make(chan []byte),
	}

	// Message handler for received messages
	go transport.waitForServerMessage()
	return transport, nil
}

// Create the WebSocket connection and establish session with the server
func newWebSocketConnection(connConfig *WebSocketConnectionConfig) (*websocket.Conn, error) {
	glog.V(3).Infof("Dial Server: %s", connConfig.TurboServer)

	config, err := websocket.NewConfig(connConfig.TurboServer+connConfig.WebSocketPath, connConfig.LocalAddress)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	usrpasswd := []byte(connConfig.WebSocketUsername + ":" + connConfig.WebSocketPassword)

	config.Header = http.Header{
		"Authorization": {"Basic " + base64.StdEncoding.EncodeToString(usrpasswd)},
	}
	config.TlsConfig = &tls.Config{InsecureSkipVerify: true}
	webs, err := websocket.DialConfig(config)

	if err != nil {
		fmt.Println("Error : ", err)
		glog.Error(err)
		if webs == nil {
			glog.Error("The websocket is null, reset")
		}
		closeWebsocket(webs)
		return nil, err
	}

	glog.Infof("Created WebSocket connection to %s", connConfig.TurboServer)
	return webs, nil
}

func (clientTransport *ClientWebSocketTransport) RawMessageReceiver() chan []byte {
	return clientTransport.inputStream
}

// Close the Websocket Transport point
func (clientTransport *ClientWebSocketTransport) CloseTransportPoint() {
	glog.Infof("Closing Transport endpoint and channel ....")
	clientTransport.closeReceived = true
	close(clientTransport.RawMessageReceiver())
	closeWebsocket(clientTransport.ws)
}

// Close the websocket connection
func closeWebsocket(wsConn *websocket.Conn) {
	//TODO:
	if wsConn != nil {
		wsConn.Close()
		wsConn = nil
	}
}

// Send serialized protobuf message bytes
func (clientTransport *ClientWebSocketTransport) Send(messageToSend *TransportMessage) {
	if clientTransport.closeReceived {
		glog.Errorf("Cannot send message : Transport endpoint is closed")
		return
	}

	if clientTransport.ws == nil {
		glog.Errorf("Cannot send message : web socket is nil")
		return
	}
	if messageToSend == nil { //.RawMsg == nil {
		glog.Errorf("Cannot send message : marshalled msg is nil")
		return
	}
	if clientTransport.ws.IsClientConn() {
		glog.V(4).Infof("Sending message on client transport %+v", clientTransport.ws)
	}
	err := websocket.Message.Send(clientTransport.ws, messageToSend.RawMsg)
	if err != nil {
		glog.Errorf("Error sending message on client transport ", err)
		//TODO: re-establish connection when error ?
		// TODO: throw exception to the upper layer
		return
	}
	glog.V(3).Infof("Successfully sent message on client transport")
}

// Receives serialized protobuf message bytes
func (clientTransport *ClientWebSocketTransport) waitForServerMessage() {
	glog.Infof("Waiting for server response")

	if clientTransport.ws == nil {
		glog.Errorf("Websocket is nil")
		return
	}
	// main loop for listening server message.
	for {
		glog.V(3).Infof("Waiting for server message ...")
		var data []byte = make([]byte, 1024)
		error := websocket.Message.Receive(clientTransport.ws, &data)
		if error != nil {
			if error == io.EOF {
				glog.Errorf("Received EOF on websocket ", clientTransport.ws)
				clientTransport.closeReceived = true
			}
			glog.Errorf("Error during receive ", error)
			closeWebsocket(clientTransport.ws)
			//// TODO: Invoke Reconnect
			//for {
			//	if err := c.Dial(c.url, c.subprotocol); err == nil {
			//		break
			//	}
			//	time.Sleep(time.Second * 1)
			//}
			return
		} else {
			glog.Infof("Received message on websocket")
			msgChannel := clientTransport.RawMessageReceiver()
			msgChannel <- data
			//TODO: this read should be accumulative - see onMessageReceived in AbstractWebsocketTransport
			glog.Infof("Delivered Raw Message on the channel, continue listening for server messages...")
		}
	}
}
