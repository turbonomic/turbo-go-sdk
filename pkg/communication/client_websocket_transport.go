package communication

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

// Connection parameters for the websocket
type WebsocketConnectionConfig struct {
	IsSecure         bool
	VmtServerAddress string
	LocalAddress     string
	ServerUsername   string
	ServerPassword   string
	RequestURI       string
}

func CreateWebSocketTransportPointUrl(connConfig *WebsocketConnectionConfig) string {
	protocol := "ws"
	if connConfig.IsSecure {
		protocol = "wss"
	}

	vmtServerUrl := protocol + "://" + connConfig.VmtServerAddress + connConfig.RequestURI
	glog.Infof("######### [CreateTransportPointUrl] Created Webscoket URL : ", vmtServerUrl)
	return vmtServerUrl
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
func CreateClientWebsocketTransport(connConfig *WebsocketConnectionConfig) (*ClientWebSocketTransport, error) {
	websocket, err := newWebSocketConnection(connConfig) // will be nil if the server is not connected

	if err != nil {
		return nil, err
	}

	transport := &ClientWebSocketTransport{
		ws:          websocket, //newWebsocketConnection(connConfig),
		inputStream: make(chan []byte),
	}

	// Message handler for received messages
	go transport.waitForServerMessage()
	return transport, nil
}

// Create the websocket connection and establish session with the server
func newWebSocketConnection(connConfig *WebsocketConnectionConfig) (*websocket.Conn, error) {
	// Websocket URL
	vmtServerUrl := CreateWebSocketTransportPointUrl(connConfig)
	//
	localAddr := "http://127.0.0.1" //Note - required in url format, but ip is don't care for us
	glog.V(3).Infof("Dial Server: %s", vmtServerUrl)

	config, err := websocket.NewConfig(vmtServerUrl, localAddr)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	usrpasswd := []byte(connConfig.ServerUsername + ":" + connConfig.ServerPassword)

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

	glog.Infof("Created WebSocket : %s+v", vmtServerUrl)
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
