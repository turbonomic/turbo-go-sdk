package communicator

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"util/rand"
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
		assert(clientMsg, cm)
		assert(*clientMsg, *cm)

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
		assert.Equal(*cInfo, *cmb.clienMessage.ContainerInfo)
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
	if assert.NotEqual((*DiscoveryResponse)(nil), cmb.clientMessage.DiscoverResponse) {
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

func TestNewAccountDefEntryBuilder(t *testing.T) {
	assert := assert.New(t)
	name := rand.String(6)
	displayName := rand.String(7)
	description := rand.String(8)
	verificationRegex := rand.String(9)
	acctDefEntryBuilder := NewAccountDefEntryBuilder(name, displayName, description, verificationRegex)
	acctDef := acctDefEntryBuilder.accountDefEntry
	if assert.NotEqual((*AccountDefEntry)(nil), acctDef) {
		assert.Equal(name, *acctDef.name)
	}
}
