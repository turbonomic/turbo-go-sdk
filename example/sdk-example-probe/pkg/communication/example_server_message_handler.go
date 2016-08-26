package communication

import (
	"fmt"
	"time"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/comm"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
	//	"github.com/vmturbo/vmturbo-go-sdk/pkg/sdk"

	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/metadata"
	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/probe"
	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/turboapi/client"

	"github.com/golang/glog"
)

// ExampleServerMessageHandler implement comm.ServerMessageHandler interface, which defines required actions of handlers that wants to work with Go SDK.
type ExampleServerMessageHandler struct {
	turboAPIClient *client.Client
	turboMetadata  *metadata.Meta

	clientMsgChan chan *proto.MediationClientMessage
}

func NewExampleServerMessageHandler(c *client.Client, d *metadata.Meta) *ExampleServerMessageHandler {
	return &ExampleServerMessageHandler{
		turboAPIClient: c,
		turboMetadata:  d,

		// Make sure only send one clien message a time.
		clientMsgChan: make(chan *proto.MediationClientMessage, 1),
	}
}

// Implement comm.ServerMessageHandler Interface -> AddTarget()
func (this *ExampleServerMessageHandler) AddTarget() {
	glog.V(2).Infof("Now adding a %s target to server.", this.turboMetadata.TargetType)

	this.turboAPIClient.AddTarget(
		this.turboMetadata.TargetType,
		this.turboMetadata.NameOrAddress,
		this.turboMetadata.TargetIdentifier,
		this.turboMetadata.Username,
		this.turboMetadata.Password)
}

// Implement comm.ServerMessageHandler Interface -> Validate()
func (this *ExampleServerMessageHandler) Validate(serverMsg *proto.MediationServerMessage) {
	glog.V(2).Infof("Now validating target %s", this.turboMetadata.NameOrAddress)
	// NOTE: Here we validate regardless of what info sent along with Validation request from server.
	// 1. Retreive message ID from server message.
	msgID := serverMsg.GetMessageID()
	// 2. Build validation response.
	validationResponse := new(proto.ValidationResponse)
	// 3. Create client message.
	clientMsg := comm.NewClientMessageBuilder(msgID).SetValidationResponse(validationResponse).Create()
	//4. Send it to client msg channel.
	this.clientMsgChan <- clientMsg
}

// Implement comm.ServerMessageHandler Interface -> DiscoverTarget()
func (this *ExampleServerMessageHandler) DiscoverTarget(serverMsg *proto.MediationServerMessage) {
	glog.V(2).Infof("Now discovering toplogy of target %s.", this.turboMetadata.NameOrAddress)

	// 1. Get message ID.
	msgID := serverMsg.GetMessageID()
	// 2. Stop being timed out. Send back KeepAlive message every 10s.
	stopCh := make(chan struct{})
	defer close(stopCh)
	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
			}

			this.keepDiscoveryAlive(msgID)

			t := time.NewTimer(time.Second * 10)
			select {
			case <-stopCh:
				return
			case <-t.C:
			}
		}

	}()
	// 3. Create probe and use probe to discover target topology.
	exampleProbe := probe.NewExampleProbe()
	discoveryResults, err := exampleProbe.Discover()
	// 4. Build discovery response.
	// If there is error during discovery, return an ErrorDTO.
	var discoveryResponse *proto.DiscoveryResponse
	if err != nil {
		// If there is error during discovery, return an ErrorDTO.
		serverity := proto.ErrorDTO_CRITICAL
		description := fmt.Sprintf("%v", err)
		errorDTO := &proto.ErrorDTO{
			Severity:    &serverity,
			Description: &description,
		}
		discoveryResponse = &proto.DiscoveryResponse{
			ErrorDTO: []*proto.ErrorDTO{errorDTO},
		}
	} else {
		// No error. Return the result entityDTOs.
		discoveryResponse = &proto.DiscoveryResponse{
			EntityDTO: discoveryResults,
		}
	}
	// 5. Build client message.
	clientMsg := comm.NewClientMessageBuilder(msgID).SetDiscoveryResponse(discoveryResponse).Create()
	// 6. Send to client msg channel.
	this.clientMsgChan <- clientMsg

}

// Implement comm.ServerMessageHandler Interface -> HandleAction()
func (this *ExampleServerMessageHandler) HandleAction(serverMsg *proto.MediationServerMessage) {
	// TODO
}

func (this *ExampleServerMessageHandler) Callback() <-chan *proto.MediationClientMessage {
	return this.clientMsgChan
}

// Send the KeepAlive message to server in order to inform server the discovery is stil ongoing. Prevent timeout.
func (this *ExampleServerMessageHandler) keepDiscoveryAlive(msgID int32) {
	glog.V(3).Infof("Keep Alive is called for message with ID: %d", msgID)

	keepAliveMsg := new(proto.KeepAlive)
	clientMsg := comm.NewClientMessageBuilder(msgID).SetKeepAlive(keepAliveMsg).Create()

	this.clientMsgChan <- clientMsg
}
