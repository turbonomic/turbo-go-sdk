package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that getIpHandler returns a pointer to a ExternalEntityLink_PropertyHandler struct
// Tests that the MethodName, DirectlyApply and EntityType member variables of the struct are not nil
func TestgetIpHandler(t *testing.T) {
	assert := assert.New(t)
	propertyHandler := getIpHandler()
	assert.NotNil(propertyHandler.EntityType)
	assert.Equal(false, *propertyHandler.DirectlyApply)
	assert.Equal("getAddress", *propertyHandler.MethodName)
}
