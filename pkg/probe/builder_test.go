package probe

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
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
	clientMsg := new(proto.MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	cm := cmbuilder.Create()
	if assert.NotEqual((*proto.MediationClientMessage)(nil), cm) {
		assert.Equal(clientMsg, cm)
		assert.Equal(*clientMsg, *cm)

	}
}
//
//// Tests that the argument passed to SetContainerInfo is assigned to this.clientMessage.ContainerInfo
//func TestSetContainerInfo(t *testing.T) {
//	assert := assert.New(t)
//	clientMsg := new(proto.MediationClientMessage)
//	cmbuilder := &ClientMessageBuilder{
//		clientMessage: clientMsg,
//	}
//	cInfo := new(proto.ContainerInfo)
//	cmb := cmbuilder.SetContainerInfo(cInfo)
//	if assert.NotEqual((*ContainerInfo)(nil), cmb.clientMessage.ContainerInfo) {
//		assert.Equal(cInfo, cmb.clientMessage.ContainerInfo)
//		assert.Equal(*cInfo, *cmb.clientMessage.ContainerInfo)
//	}
//}

// Tests that the argument passed to SetValidationResponse is assigned to this.clientMessage.ValidationResponse
func TestSetValidationResponse(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(proto.MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	val := new(proto.ValidationResponse)
	cmb := cmbuilder.SetValidationResponse(val)
	if assert.NotEqual((*proto.ValidationResponse)(nil), cmb.clientMessage.GetValidationResponse()) {
		assert.Equal(val, cmb.clientMessage.GetValidationResponse())
		assert.Equal(*val, *cmb.clientMessage.GetValidationResponse())
	}
}

// Tests that the argument passed to SetDiscoveryResponse method is assigned
// to this.clientMessage.DiscoveryResponse
func TestSetDiscoveryResponse(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(proto.MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	discresponse := new(proto.DiscoveryResponse)
	cmb := cmbuilder.SetDiscoveryResponse(discresponse)
	if assert.NotEqual((*proto.DiscoveryResponse)(nil), cmb.clientMessage.GetDiscoveryResponse()) {
		assert.Equal(discresponse, cmb.clientMessage.GetDiscoveryResponse())
		assert.Equal(*discresponse, *cmb.clientMessage.GetDiscoveryResponse())
	}
}

// Tests that the argument passed to SetKeepAlive method is assigned to this.clientMessage.KeepAlive
func TestSetKeepAlive(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(proto.MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	keepAlive := new(proto.KeepAlive)
	cmb := cmbuilder.SetKeepAlive(keepAlive)
	if assert.NotEqual((*proto.KeepAlive)(nil), cmb.clientMessage.GetKeepAlive()) {
		assert.Equal(keepAlive, cmb.clientMessage.GetKeepAlive())
		assert.Equal(*keepAlive, *cmb.clientMessage.GetKeepAlive())
	}
}

// Test that the argument passed to SetActionProgress method is assigned to this.clientMessage.ActionProgress
func TestSetActionProgress(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(proto.MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	actionP := new(proto.ActionProgress)
	cmb := cmbuilder.SetActionProgress(actionP)
	if assert.NotEqual((*proto.ActionProgress)(nil), cmb.clientMessage.GetActionProgress()) {
		assert.Equal(actionP, cmb.clientMessage.GetActionProgress())
		assert.Equal(*actionP, *cmb.clientMessage.GetActionProgress())
	}
}

// Tests that the argument passed to SetActionResponse is assigned to this.clientMessage.ActionResponse
func TestSetActionResponset(t *testing.T) {
	assert := assert.New(t)
	clientMsg := new(proto.MediationClientMessage)
	cmbuilder := &ClientMessageBuilder{
		clientMessage: clientMsg,
	}
	actionResult := new(proto.ActionResult)
	cmb := cmbuilder.SetActionResponse(actionResult)
	if assert.NotEqual((*proto.ActionResult)(nil), cmb.clientMessage.GetActionResponse()) {
		assert.Equal(actionResult, cmb.clientMessage.GetActionResponse())
		assert.Equal(*actionResult, *cmb.clientMessage.GetActionResponse())
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
	entryType := false	//AccountDefEntry_OPTIONAL
	isSecret := true
	acctDefEntryBuilder := NewAccountDefEntryBuilder(name, displayName, description, verificationRegex, entryType, isSecret)
	acctDef := acctDefEntryBuilder.accountDefEntry
	customAcctDef := acctDef.GetCustomDefinition()
	if assert.NotEqual((*proto.AccountDefEntry)(nil), acctDef) {
		assert.Equal(name, *customAcctDef.Name)
		assert.Equal(displayName, *customAcctDef.DisplayName)
		assert.Equal(description, *customAcctDef.Description)
		assert.Equal(verificationRegex, *customAcctDef.VerificationRegex)
		assert.Equal(entryType, *acctDef.Mandatory)
		assert.Equal(isSecret, *customAcctDef.IsSecret)
	}
}

// Tests that the method Create returns the correct this.accountDefentry
func TestCreate(t *testing.T) {
	assert := assert.New(t)
	acctDefEntry := new(proto.AccountDefEntry)
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
	probeType := "ProbeType1"	//rand.String(6)
	probeCat := "ProbeCategory1"	//rand.String(7)
	var supplyCS []*proto.TemplateDTO
	var acctDef []*proto.AccountDefEntry
	probeInfoBldr := NewProbeInfoBuilder(probeType, probeCat, supplyCS, acctDef)
	assert.Equal(probeType, *probeInfoBldr.probeInfo.ProbeType)
}

// Tests that the method Create() returns a pointer to the ProbeInfo struct in the object that Create
// called on
func TestProbeInfo_Create(t *testing.T) {
	assert := assert.New(t)
	probeinfo := new(proto.ProbeInfo)
	builder := &ProbeInfoBuilder{
		probeInfo: probeinfo,
	}
	pi := builder.Create()
	assert.Equal(probeinfo, pi)
	assert.Equal(*probeinfo, *pi)

}
