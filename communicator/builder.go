package communicator

import (
	"github.com/vmturbo/vmturbo-go-sdk/sdk"
)

// A ClientMessageBuilder builds a ClientMessage instance.
type ClientMessageBuilder struct {
	clientMessage *MediationClientMessage
}

// Get an instance of ClientMessageBuilder
func NewClientMessageBuilder(messageID int32) *ClientMessageBuilder {
	clientMessage := &MediationClientMessage{
		MessageID: &messageID,
	}
	return &ClientMessageBuilder{
		clientMessage: clientMessage,
	}
}

// Build an instance of ClientMessage.
func (cmb *ClientMessageBuilder) Create() *MediationClientMessage {
	return cmb.clientMessage
}

// // Set the ContainerInfo of the ClientMessage if necessary.
// func (cmb *ClientMessageBuilder) SetContainerInfo(containerInfo *ContainerInfo) *ClientMessageBuilder {
// 	cmb.clientMessage.ContainerInfo = containerInfo
// 	return cmb
// }

// set the validation response
func (cmb *ClientMessageBuilder) SetValidationResponse(validationResponse *ValidationResponse) *ClientMessageBuilder {
	cmb.clientMessage.ValidationResponse = validationResponse
	return cmb
}

// set discovery response
func (cmb *ClientMessageBuilder) SetDiscoveryResponse(discoveryResponse *DiscoveryResponse) *ClientMessageBuilder {
	cmb.clientMessage.DiscoveryResponse = discoveryResponse
	return cmb
}

// set discovery keep alive
func (cmb *ClientMessageBuilder) SetKeepAlive(keepAlive *KeepAlive) *ClientMessageBuilder {
	cmb.clientMessage.KeepAlive = keepAlive
	return cmb
}

// set action progress
func (cmb *ClientMessageBuilder) SetActionProgress(actionProgress *ActionProgress) *ClientMessageBuilder {
	cmb.clientMessage.ActionProgress = actionProgress
	return cmb
}

// set action response
func (cmb *ClientMessageBuilder) SetActionResponse(actionResponse *ActionResult) *ClientMessageBuilder {
	cmb.clientMessage.ActionResponse = actionResponse
	return cmb
}

// An AccountDefEntryBuilder builds an AccountDefEntry instance.
type AccountDefEntryBuilder struct {
	accountDefEntry *AccountDefEntry
}

func NewAccountDefEntryBuilder(name, displayName, description, verificationRegex string,
	entryType AccountDefEntry_AccountDefEntryType, isSecret bool) *AccountDefEntryBuilder {
	accountDefEntry := &AccountDefEntry{
		Name:              &name,
		DisplayName:       &displayName,
		Description:       &description,
		VerificationRegex: &verificationRegex,
		Type:              &entryType,
		IsSecret:          &isSecret,
	}
	return &AccountDefEntryBuilder{
		accountDefEntry: accountDefEntry,
	}
}

func (builder *AccountDefEntryBuilder) Create() *AccountDefEntry {
	return builder.accountDefEntry
}

// A ProbeInfoBuilder builds a ProbeInfo instance.
type ProbeInfoBuilder struct {
	probeInfo *ProbeInfo
}

func NewProbeInfoBuilder(probeType, probeCat string, supplyChainSet []*sdk.TemplateDTO, acctDef []*AccountDefEntry) *ProbeInfoBuilder {
	probeInfo := &ProbeInfo{
		ProbeType:                &probeType,
		ProbeCategory:            &probeCat,
		SupplyChainDefinitionSet: supplyChainSet,
		AccountDefinition:        acctDef,
	}
	return &ProbeInfoBuilder{
		probeInfo: probeInfo,
	}
}

func (builder *ProbeInfoBuilder) Create() *ProbeInfo {
	return builder.probeInfo
}
