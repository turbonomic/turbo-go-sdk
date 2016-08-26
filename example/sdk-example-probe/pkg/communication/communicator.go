package communication

import (
	"fmt"
	"net/url"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/comm"
	//"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
	//	"github.com/vmturbo/vmturbo-go-sdk/pkg/sdk"

	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/metadata"
	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/registration"
	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/turboapi/client"
)

//type Callback struct {
//	c chan *proto.MediationClientMessage
//}
//
//func (this *Callback) GetClientMessage(clientMsg *proto.MediationClientMessage) {
//	c <- clientMsg
//}

type Communicator struct {
	wsComm *comm.WebSocketCommunicator

	handler comm.ServerMessageHandler

	stop chan struct{}
}

// use meta data and customized handler to create a communicator.
func NewCommunicator(meta *metadata.Meta, handler comm.ServerMessageHandler) *Communicator {
	//turboAPIClient, err := createTurboAPIClientFromMeta(meta)
	//if err != nil {
	//	return nil, err
	//}

	//clientMsgChan := make(chan *proto.MediationClientMessage)

	//handler := &ExampleServerMessageHandler{
	//	turboAPIClient: turboAPIClient,
	//	turboMetadata:  meta,
	//}

	wsCommunicator := &comm.WebSocketCommunicator{
		VmtServerAddress: meta.ServerAddress,
		LocalAddress:     meta.LocalAddress,
		ServerUsername:   meta.WebSocketUsername,
		ServerPassword:   meta.WebSocketPassword,
		ServerMsgHandler: handler,
	}

	return &Communicator{
		wsComm:  wsCommunicator,
		handler: handler,

		stop: make(chan struct{}),
	}

}

func (this *Communicator) Start() {
	go this.listenCallback()
}

func (this *Communicator) Stop() {
	this.stop <- struct{}{}
}

func createTurboAPIClientFromMeta(meta *metadata.Meta) (*client.Client, error) {
	serverAddress, err := url.Parse(meta.ServerAddress)
	if err != nil {
		return nil, fmt.Errorf("Invalid server address: %v", err)
	}
	apiClientConfig := client.NewConfigBuilder(serverAddress).BasicAuthentication(meta.Username, meta.Password).Create()
	return client.NewAPIClient(apiClientConfig), nil
}

// Listen to callback.c and send back any client message to server.
// Should be run in a separate goroutine.
func (this *Communicator) listenCallback() {
	for {
		select {
		case msg := <-this.handler.Callback():
			this.wsComm.SendClientMessage(msg)
		case <-this.stop:
			return
		}
	}
}

// Register target to server.
func (this *Communicator) RegisterExampleProbe(targetType string) {
	containerInfo := registration.NewMediationContainerInfoBuilder(targetType).Build()
	this.wsComm.RegisterAndListen(containerInfo)
}
