package mediationcontainer

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

const (
	Closed           TransportStatus = "closed"
	Ready            TransportStatus = "ready"
	handshakeTimeout                 = 60 * time.Second
	wsReadLimit                      = 33554432 // 32 MB
	writeWaitSeconds                 = 120
)

type TransportStatus string

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

// ============================ ClientWebSocketTransport - Start, Connect, Close ======================================
// Implementation of the ITransport for WebSocket communication to send and receive serialized protobuf message bytes
type ClientWebSocketTransport struct {
	// current status of the transport layer.
	status                   TransportStatus
	ws                       *websocket.Conn // created during Connect()
	connConfig               *WebSocketConnectionConfig
	inputStreamCh            chan []byte // unbuffered channel
	closeRequested           bool
	stopListenerCh           chan bool // buffered channel
	connClosedNotificationCh chan bool // channel where the transport connection error will be notified
}

// Instantiate a new ClientWebSocketTransport endpoint for the client
func CreateClientWebSocketTransport(connConfig *WebSocketConnectionConfig) *ClientWebSocketTransport {
	transport := &ClientWebSocketTransport{
		connConfig:               connConfig,
		connClosedNotificationCh: make(chan bool),
	}
	return transport
}

// WebSocket connection is established with the server
func (clientTransport *ClientWebSocketTransport) Connect() error {
	// Close any previous connected WebSocket connection and set current connection to nil.
	clientTransport.closeAndResetWebSocket()

	// loop till server is up or close received
	// TODO: give an optional timeout to wait for server in performWebSocketConnection()
	err := clientTransport.performWebSocketConnection() // Blocks or till transport is closed
	if err != nil {
		return fmt.Errorf("Cannot connect with server websocket %s", err)
	}

	glog.V(4).Infof("[Connect] Connected to server " + clientTransport.GetConnectionId())

	clientTransport.stopListenerCh = make(chan bool, 1) // Channel to stop the routine that listens for messages
	clientTransport.inputStreamCh = make(chan []byte)   // Message Queue
	// Message handler for received messages
	clientTransport.ListenForMessages() // spawns a new routine
	return nil
}

func (clientTransport *ClientWebSocketTransport) NotifyClosed() chan bool {
	return clientTransport.connClosedNotificationCh
}

func (clientTransport *ClientWebSocketTransport) GetConnectionId() string {
	if clientTransport.status == Closed {
		return ""
	}
	return clientTransport.ws.RemoteAddr().String() + "::" + clientTransport.ws.LocalAddr().String()
}

// Close the WebSocket Transport point
func (clientTransport *ClientWebSocketTransport) CloseTransportPoint() {
	glog.V(4).Infof("[CloseTransportPoint] closing transport endpoint and listener routine")
	clientTransport.closeRequested = true
	// close listener
	clientTransport.stopListenForMessages()
	clientTransport.closeAndResetWebSocket()
}

// Close current WebSocket connection and set it to nil.
func (clientTransport *ClientWebSocketTransport) closeAndResetWebSocket() {
	if clientTransport.status == Closed {
		return
	}
	// close WebSocket
	if clientTransport.ws != nil {
		clientTransport.ws.WriteMessage(websocket.CloseMessage, []byte{})
		clientTransport.ws.Close()
		clientTransport.ws = nil
	}
	clientTransport.status = Closed
}

// ================================================= Message Listener =============================================
func (clientTransport *ClientWebSocketTransport) stopListenForMessages() {
	if clientTransport.stopListenerCh != nil {
		glog.V(4).Infof("[StopListenForMessages] closing stopListenerCh %+v", clientTransport.stopListenerCh)
		clientTransport.stopListenerCh <- true
		close(clientTransport.stopListenerCh)
		glog.V(4).Infof("[StopListenForMessages] closed stopListenerCh %+v", clientTransport.stopListenerCh)
	}
}

// Routine to listen for messages on the websocket.
// The websocket is continuously checked for messages and queued on the clientTransport.inputStream channel
// Routine exits when a message is sent on clientTransport.stopListenerCh.
//
func (clientTransport *ClientWebSocketTransport) ListenForMessages() {
	glog.V(3).Infof("[ListenForMessages] %s : ENTER  ", time.Now())

	go func() {
		for {
			glog.V(4).Info("[ListenForMessages] waiting for messages on websocket transport")
			glog.V(4).Infof("[ListenForMessages] waiting for messages on websocket transport : %++v", clientTransport)
			select {
			case <-clientTransport.stopListenerCh:
				close(clientTransport.inputStreamCh) // This listener routine is the writer for this channel
				glog.V(4).Infof("[ListenForMessages] closed inputStreamCh %+v", clientTransport.inputStreamCh)
				return
			default:
				if clientTransport.status != Ready {
					glog.Errorf("WebSocket transport layer status is %s", clientTransport.status)
					glog.Errorf("WebSocket is not ready.")
					continue
				}
				glog.V(2).Infof("[ListenForMessages]: connected, waiting for server response ...")

				msgType, data, err := clientTransport.ws.ReadMessage()
				glog.V(3).Infof("Received websocket message of type %d and size %d", msgType, len(data))

				// Notify errors and break
				if err != nil {
					//TODO handle err according to possible type. Now reconnect when never there is an error, which may not necessary.
					if err == io.EOF {
						glog.Errorf("[ListenForMessages] received EOF on websocket %s", err)
					}

					glog.Errorf("[ListenForMessages] error during receive %v", err)
					//notify error with the connection
					clientTransport.connClosedNotificationCh <- true // Note: this will block till the message is received
					// close current WebSocket connection.
					clientTransport.closeAndResetWebSocket()
					clientTransport.stopListenForMessages()

					glog.V(2).Infof("[ListenForMessages] error notified, will re-establish websocket connection")
					break
				}
				// write the message on the channel
				glog.V(3).Infof("[ListenForMessages] received message on websocket of size %d", len(data))

				clientTransport.queueRawMessage(data) // Note: this will block till the message is read
				glog.V(4).Infof("[ListenForMessages] delivered websocket message, continue listening for server messages...")
			} //end select
		} //end for
	}()
	glog.V(4).Infof("[ListenForMessages] : END")
}

func (clientTransport *ClientWebSocketTransport) queueRawMessage(data []byte) {
	//TODO: this read should be accumulative - see onMessageReceived in AbstractWebsocketTransport
	clientTransport.inputStreamCh <- data
	// TODO: how to ensure that the channel is open to write
}

func (clientTransport *ClientWebSocketTransport) RawMessageReceiver() chan []byte {
	return clientTransport.inputStreamCh
}

// ==================================================== Message Sender ===============================================
// Send serialized protobuf message bytes
func (clientTransport *ClientWebSocketTransport) Send(messageToSend *TransportMessage) error {
	if clientTransport.closeRequested {
		glog.Errorf("Cannot send message : transport endpoint is closed")
		return errors.New("Cannot send message: transport endpoint is closed")
	}

	if clientTransport.status != Ready {
		glog.Errorf("WebSocket transport layer status is %s", clientTransport.status)
		return errors.New("Cannot send message: web socket is not ready")
	}
	if messageToSend == nil { //.RawMsg == nil {
		glog.Errorf("Cannot send message : marshalled msg is nil")
		return errors.New("Cannot send message: marshalled msg is nil")
	}
	du := time.Second * time.Duration(writeWaitSeconds)
	clientTransport.ws.SetWriteDeadline(time.Now().Add(du))
	err := clientTransport.ws.WriteMessage(websocket.BinaryMessage, messageToSend.RawMsg)
	if err != nil {
		glog.Errorf("Error sending message on client transport: %s", err)
		return fmt.Errorf("Error sending message on client transport: %s", err)
	}
	glog.V(4).Infof("Successfully sent message on client transport")
	return nil
}

// ====================================== Websocket Connection =========================================================
// Establish connection to server websocket until connected or until the transport endpoint is closed
func (clientTransport *ClientWebSocketTransport) performWebSocketConnection() error {
	connRetryIntervalSeconds := time.Second * 30 // TODO: use ConnectionRetry parameter from the connConfig or default
	connConfig := clientTransport.connConfig
	// WebSocket URL
	vmtServerUrl := connConfig.TurboServer + connConfig.WebSocketPath
	glog.Infof("[performWebSocketConnection]: %s", vmtServerUrl)

	for !clientTransport.closeRequested { // only set when CloseTransportPoint() is called
		ws, err := openWebSocketConn(connConfig, vmtServerUrl)

		if err != nil {
			// print at debug level after some time
			glog.V(3).Infof("[performWebSocketConnection] %v : unable to connect to %s. Retrying in %v\n", time.Now(), vmtServerUrl, connRetryIntervalSeconds)

			time.Sleep(connRetryIntervalSeconds)
		} else {
			clientTransport.ws = ws
			clientTransport.status = Ready
			glog.V(2).Infof("[performWebSocketConnection]*********** Connected to server " + clientTransport.GetConnectionId())
			glog.V(2).Infof("WebSocket transport layer is ready.")

			return nil
		}
	}
	glog.V(4).Infof("[performWebSocketConnection] exit connect routine, close = %s ", clientTransport.closeRequested)
	return errors.New("Abort client socket connect, transport is closed")
}

func openWebSocketConn(connConfig *WebSocketConnectionConfig, vmtServerUrl string) (*websocket.Conn, error) {
	//1. set up dialer
	d := &websocket.Dialer{
		HandshakeTimeout: handshakeTimeout,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}

	//2. auth header
	header := getAuthHeader(connConfig.WebSocketUsername, connConfig.WebSocketPassword)

	//3. connect it
	c, _, err := d.Dial(vmtServerUrl, header)
	if err != nil {
		glog.Errorf("Failed to connect to server(%s): %v", vmtServerUrl, err)
		return nil, err
	}

	c.SetReadLimit(wsReadLimit)

	return c, nil
}

func getAuthHeader(user, password string) http.Header {
	dat := []byte(fmt.Sprintf("%s:%s", user, password))
	header := http.Header{
		"Authorization": {"Basic " + base64.StdEncoding.EncodeToString(dat)},
	}

	return header
}
