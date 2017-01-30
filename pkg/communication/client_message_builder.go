package communication


import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// ========== Builder for proto messages created from the probe =============
// A ClientMessageBuilder builds a ClientMessage instance.
type ClientMessageBuilder struct {
	clientMessage *proto.MediationClientMessage
}

// Get an instance of ClientMessageBuilder
func NewClientMessageBuilder(messageID int32) *ClientMessageBuilder {
	clientMessage := &proto.MediationClientMessage{
		MessageID: &messageID,
	}
	return &ClientMessageBuilder{
		clientMessage: clientMessage,
	}
}

// Build an instance of ClientMessage.
func (cmb *ClientMessageBuilder) Create() *proto.MediationClientMessage {
	return cmb.clientMessage
}
//
//// Set the ContainerInfo of the ClientMessage if necessary.
//func (cmb *ClientMessageBuilder) SetContainerInfo(containerInfo *proto.ContainerInfo) *ClientMessageBuilder {
//	cmb.clientMessage.ContainerInfo = containerInfo
//	return cmb
//}

// set the validation response
func (cmb *ClientMessageBuilder) SetValidationResponse(validationResponse *proto.ValidationResponse) *ClientMessageBuilder {
	response := &proto.MediationClientMessage_ValidationResponse {
		ValidationResponse: 	validationResponse,
	}

	cmb.clientMessage = &proto.MediationClientMessage{
		MediationClientMessage: response,
	}
	//cmb.clientMessage.ValidationResponse = validationResponse
	return cmb
}

// set discovery response
func (cmb *ClientMessageBuilder) SetDiscoveryResponse(discoveryResponse *proto.DiscoveryResponse) *ClientMessageBuilder {
	response := &proto.MediationClientMessage_DiscoveryResponse {
		DiscoveryResponse: 	discoveryResponse,
	}
	cmb.clientMessage = &proto.MediationClientMessage{
		MediationClientMessage: response,
	}

	//cmb.clientMessage.DiscoveryResponse = discoveryResponse
	return cmb
}

// set discovery keep alive
func (cmb *ClientMessageBuilder) SetKeepAlive(keepAlive *proto.KeepAlive) *ClientMessageBuilder {
	response := &proto.MediationClientMessage_KeepAlive {
		KeepAlive: 	keepAlive,
	}
	cmb.clientMessage = &proto.MediationClientMessage{
		MediationClientMessage: response,
	}
	//cmb.clientMessage.KeepAlive = keepAlive
	return cmb
}

// set action progress
func (cmb *ClientMessageBuilder) SetActionProgress(actionProgress *proto.ActionProgress) *ClientMessageBuilder {
	response := &proto.MediationClientMessage_ActionProgress {
		ActionProgress: 	actionProgress,
	}
	cmb.clientMessage = &proto.MediationClientMessage{
		MediationClientMessage: response,
	}
	// cmb.clientMessage.ActionProgress = actionProgress
	return cmb
}

// set action response
func (cmb *ClientMessageBuilder) SetActionResponse(actionResponse *proto.ActionResult) *ClientMessageBuilder {
	response := &proto.MediationClientMessage_ActionResponse{
		ActionResponse: 	actionResponse,
	}
	cmb.clientMessage = &proto.MediationClientMessage{
		MediationClientMessage: response,
	}
	// cmb.clientMessage.ActionResponse = actionResponse
	return cmb
}
