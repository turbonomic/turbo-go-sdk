package supplychain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that getIpHandler returns a pointer to a ExternalEntityLink_PropertyHandler struct
// Tests that the MethodName, DirectlyApply and EntityType member variables of the struct are not nil
// DirectlyApply is tested to be false
func TestGetIpHandler(t *testing.T) {
	assert := assert.New(t)
	propertyHandler := getIPHandler()
	assert.NotNil(propertyHandler.EntityType)
	if assert.NotNil(propertyHandler.DirectlyApply) {
		assert.Equal(false, *propertyHandler.DirectlyApply)
	}
	if assert.NotNil(propertyHandler.MethodName) {
		assert.Equal(getIPAddressMethodName, *propertyHandler.MethodName)
	}
}

// Test that the Entity and PropertyHandler member variables of
// the struct returned bu getVirtualMachineIpProperty are not nil
// tests that the attribute string type member variable is "UsesEndPoints"
func TestGetVirtualMachineIpProperty(t *testing.T) {
	assert := assert.New(t)
	serverEntityProp := getVirtualMachineIPProperty()
	if assert.NotNil(serverEntityProp.Attribute) {
		assert.Equal(serverIPProperty, *serverEntityProp.Attribute)

	}
	assert.NotNil(serverEntityProp.Entity)
	assert.NotNil(serverEntityProp.PropertyHandler)
}

func TestGetVirtualMachineUUIDProperty(t *testing.T) {
	assert := assert.New(t)
	serverEntityProp := getVirtualMachineUUIDProperty()
	if assert.NotNil(serverEntityProp.Attribute) {
		assert.Equal(serverUUIDProperty, *serverEntityProp.Attribute)
	}
	assert.NotNil(serverEntityProp.Entity)
}
