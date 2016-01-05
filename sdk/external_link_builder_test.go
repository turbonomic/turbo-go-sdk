package sdk

import (
	"github.com/stretchr/testify/assert"
	//	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	//	mathrand "math/rand"
	"testing"
)

func TestNewExternalEntityBuilder(t *testing.T) {
	assert := assert.New(t)
	newEELB := NewExternalEntityLinkBuilder()
	assert.NotNil(newEELB)
	assert.NotNil(newEELB.entityLink)
}
