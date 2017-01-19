package probe

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

//// Helper methods to create AccountDefinition map for sub classes of the probe
// An AccountDefEntryBuilder builds an AccountDefEntry instance.
type AccountDefEntryBuilder struct {
	accountDefEntry *proto.AccountDefEntry
}

//func NewAccountDefEntryBuilder(name, displayName, description, verificationRegex string,
//	entryType proto.AccountDefEntry_AccountDefEntryType, isSecret bool) *AccountDefEntryBuilder {
//	accountDefEntry := &proto.AccountDefEntry{
//		Name:              &name,
//		DisplayName:       &displayName,
//		Description:       &description,
//		VerificationRegex: &verificationRegex,
//		Type:              &entryType,
//		IsSecret:          &isSecret,
//	}
//	return &AccountDefEntryBuilder{
//		accountDefEntry: accountDefEntry,
//	}
//}
//
func NewAccountDefEntryBuilder(name, displayName, description, verificationRegex string,
				mandatory  bool,	//proto.AccountDefEntry_AccountDefEntryType,
				isSecret bool) *AccountDefEntryBuilder {
	fieldType := &proto.CustomAccountDefEntry_PrimitiveValue_{
		PrimitiveValue: proto.CustomAccountDefEntry_STRING,
	}
	entry := &proto.CustomAccountDefEntry {

		Name: &name,
		DisplayName: &displayName,
		Description: &description,
		VerificationRegex: &verificationRegex,
		IsSecret: &isSecret,
		FieldType: fieldType,
	}

	customDef := &proto.AccountDefEntry_CustomDefinition {
		CustomDefinition: entry,
	}

	accountDefEntry := &proto.AccountDefEntry{
		Mandatory: &mandatory,
		Definition: customDef,
	}

	return &AccountDefEntryBuilder{
		accountDefEntry: accountDefEntry,
	}
}

func (builder *AccountDefEntryBuilder) Create() *proto.AccountDefEntry {
	return builder.accountDefEntry
}

// A ProbeInfoBuilder builds a ProbeInfo instance.
type ProbeInfoBuilder struct {
	probeInfo *proto.ProbeInfo
}

// NewProbeInfoBuilder builds the ProbeInfo DTO for the given probe
func NewProbeInfoBuilder(probeType, probeCat string,
				supplyChainSet []*proto.TemplateDTO,
				acctDef []*proto.AccountDefEntry) *ProbeInfoBuilder {
	// New ProbeInfo protobuf with this input
	probeInfo := &proto.ProbeInfo{
		ProbeType:                &probeType,
		ProbeCategory:            &probeCat,
		SupplyChainDefinitionSet: supplyChainSet,
		AccountDefinition:        acctDef,
	}
	return &ProbeInfoBuilder{
		probeInfo: probeInfo,
	}
}

func (builder *ProbeInfoBuilder) Create() *proto.ProbeInfo {
	return builder.probeInfo
}


//
//// A ProbeInfoBuilder builds a ProbeInfo instance.
//type ProbeInfoBuilder2 struct {
//	probeInfo *proto.ProbeInfo
//}
//
////func NewProbeInfoBuilder2(probeType, probeCat string,
////				acctDef []*proto.AccountDefEntry) *ProbeInfoBuilder2 {
////	// New ProbeInfo protobuf with this input
////	probeInfo := &proto.ProbeInfo{
////		ProbeType:                &probeType,
////		ProbeCategory:            &probeCat,
////		AccountDefinition:        acctDef,
////
////	}
////	return &ProbeInfoBuilder2{
////		probeInfo: probeInfo,
////	}
////}
//func (builder *ProbeInfoBuilder2) Create() *proto.ProbeInfo {
//	return builder.probeInfo
//}
