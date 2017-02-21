package mediationcontainer

import (
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
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
