package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that getIpHandler returns a pointer to a ExternalEntityLink_PropertyHandler struct
// Tests that the MethodName, DirectlyApply and EntityType member variables of the struct are not nil
// DirectlyApply is tested to be false
func TestgetIpHandler(t *testing.T) {
	assert := assert.New(t)
	propertyHandler := getIpHandler()
	assert.NotNil(propertyHandler.EntityType)
	assert.Equal(false, *propertyHandler.DirectlyApply)
	assert.Equal("getAddress", *propertyHandler.MethodName)
}

// Test that the Entity and PropertyHandler member variables of
// the struct returned bu getVirtualMachineIpProperty are not nil
// tests that the attribute string type member variable is "UsesEndPoints"
func TestgetVirtualMachineIpProperty(t *testing.T) {
	assert := assert.New(t)
	serverEntityProp := getVirtualMachineIpProperty()
	assert.Equal("UsesEndPoints", *serverEntityProp.Attribute)
	assert.NotNil(serverEntityProp.Entity)
	assert.NotNil(serverEntityProp.PropertyHandler)
}
