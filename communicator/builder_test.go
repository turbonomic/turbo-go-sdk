package communicator

import (
	"github.com/stretchr/testify/assert"
	"github.com/vmturbo/vmturbo-go-sdk/sdk"
	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	"testing"
)

// Tests that the function NewSupplyChainBuilder() returns an initialized SupplyChainNodeBuilder struct
func TestNewClientMessageBuilder(t *testing.T) {
	assert := assert.New(t)
	var msgid int32 = 98
	clientMessageBuilder := NewClientMessageBuilder(msgid)
	assert.Equal(msgid, *clientMessageBuilder.clientMessage.MessageID)
}

// Tests that the Create method, when called on a ClientMessageBuilder, returns
// the clientMessage value inside the object it was called on
func TestCreate_ClientMessageBuilder(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	cm := cmbuilder.Create()
	if assert.NotEqual((*MediationClientMessage)(nil), cm) {
		assert.Equal(clientMsg, cm)
		assert.Equal(*clientMsg, *cm)

	}
}

// Tests that the argument passed to SetContainerInfo is assigned to this.clientMessage.ContainerInfo
func TestSetContainerInfo(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	cInfo := new(ContainerInfo)
	cmb := cmbuilder.SetContainerInfo(cInfo)
	if assert.NotEqual((*ContainerInfo)(nil), cmb.clientMessage.ContainerInfo) {
		assert.Equal(cInfo, cmb.clientMessage.ContainerInfo)
		assert.Equal(*cInfo, *cmb.clientMessage.ContainerInfo)
	}
}

// Tests that the argument passed to SetValidationResponse is assigned to this.clientMessage.ValidationResponse
func TestSetValidationResponse(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	val := new(ValidationResponse)
	cmb := cmbuilder.SetValidationResponse(val)
	if assert.NotEqual((*ValidationResponse)(nil), cmb.clientMessage.ValidationResponse) {
		assert.Equal(val, cmb.clientMessage.ValidationResponse)
		assert.Equal(*val, *cmb.clientMessage.ValidationResponse)
	}
}

// Tests that the argument passed to SetDiscoveryResponse method is assigned
// to this.clientMessage.DiscoveryResponse
func TestSetDiscoveryResponse(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	discresponse := new(DiscoveryResponse)
	cmb := cmbuilder.SetDiscoveryResponse(discresponse)
	if assert.NotEqual((*DiscoveryResponse)(nil), cmb.clientMessage.DiscoveryResponse) {
		assert.Equal(discresponse, cmb.clientMessage.DiscoveryResponse)
		assert.Equal(*discresponse, *cmb.clientMessage.DiscoveryResponse)
	}
}

// Tests that the argument passed to SetKeepAlive method is assigned to this.clientMessage.KeepAlive
func TestSetKeepAlive(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	keepAlive := new(KeepAlive)
	cmb := cmbuilder.SetKeepAlive(keepAlive)
	if assert.NotEqual((*KeepAlive)(nil), cmb.clientMessage.KeepAlive) {
		assert.Equal(keepAlive, cmb.clientMessage.KeepAlive)
		assert.Equal(*keepAlive, *cmb.clientMessage.KeepAlive)
	}
}

// Test that the argument passed to SetActionProgress method is assigned to this.clientMessage.ActionProgress
func TestSetActionProgress(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	actionP := new(ActionProgress)
	cmb := cmbuilder.SetActionProgress(actionP)
	if assert.NotEqual((*ActionProgress)(nil), cmb.clientMessage.ActionProgress) {
		assert.Equal(actionP, cmb.clientMessage.ActionProgress)
		assert.Equal(*actionP, *cmb.clientMessage.ActionProgress)
	}
}

// Tests that the argument passed to SetActionResponse is assigned to this.clientMessage.ActionResponse
func TestSetActionResponset(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	actionResult := new(ActionResult)
	cmb := cmbuilder.SetActionResponse(actionResult)
	if assert.NotEqual((*ActionResult)(nil), cmb.clientMessage.ActionResponse) {
		assert.Equal(actionResult, cmb.clientMessage.ActionResponse)
		assert.Equal(*actionResult, *cmb.clientMessage.ActionResponse)
	}
}

// Tests that the method NewAccountDefEntryBuilder creates an AccountDefEntry struct with member
// variables set equal to the pointers to the arguments passed to it
func TestNewAccountDefEntryBuilder(t *testing.T) {
	assert := assert.New(t)
	name := rand.String(6)
	displayName := rand.String(7)
	description := rand.String(8)
	verificationRegex := rand.String(9)
	entryType := AccountDefEntry_OPTIONAL
	isSecret := true
	acctDefEntryBuilder := NewAccountDefEntryBuilder(name, displayName, description, verificationRegex, entryType, isSecret)
	acctDef := acctDefEntryBuilder.accountDefEntry
	if assert.NotEqual((*AccountDefEntry)(nil), acctDef) {
		assert.Equal(name, *acctDef.Name)
		assert.Equal(displayName, *acctDef.DisplayName)
		assert.Equal(description, *acctDef.Description)
		assert.Equal(verificationRegex, *acctDef.VerificationRegex)
		assert.Equal(entryType, *acctDef.Type)
		assert.Equal(isSecret, *acctDef.IsSecret)
	}
}

// Tests that the method Create returns the correct this.accountDefentry
func TestCreate(t *testing.T) {
	assert := assert.New(t)
	acctDefEntry := new(AccountDefEntry)
	acctDefEBuilder := &AccountDefEntryBuilder{
		accountDefEntry: acctDefEntry,
	}
	acctDE := acctDefEBuilder.Create()
	assert.Equal(acctDefEntry, acctDE)
}

// Tests that NewProbeInfoBuilder creates a ProbeInfo struct with member variables set to the
// arguments passed to it and creates a ProbeInfoBuilder struct with its probeInfo variable set
// to the new ProbeInfo struct containing the passed arguments
func TestNewProbeInfoBuilder(t *testing.T) {
	assert := assert.New(t)
	probeType := rand.String(6)
	probeCat := rand.String(7)
	var supplyCS []*sdk.TemplateDTO
	var acctDef []*AccountDefEntry
	probeInfoBldr := NewProbeInfoBuilder(probeType, probeCat, supplyCS, acctDef)
	assert.Equal(probeType, *probeInfoBldr.probeInfo.ProbeType)
}

// Tests that the method Create() returns a pointer to the ProbeInfo struct in the object that Create
// called on
func TestProbeInfo_Create(t *testing.T) {
	assert := assert.New(t)
	probeinfo := new(ProbeInfo)
	builder := &ProbeInfoBuilder{
		probeInfo: probeinfo,
	}
	pi := builder.Create()
	assert.Equal(probeinfo, pi)
	assert.Equal(*probeinfo, *pi)

}
