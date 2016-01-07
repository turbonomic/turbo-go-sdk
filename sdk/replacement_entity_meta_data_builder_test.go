package sdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	"testing"
)

// Tests to see if the method NewReplacementEntityMetaDataBuilder() returns a
// ReplacementEntityMetaDataBuilder with non-nil member variable metaData
// and tests that the member variables of the struct of type EntityDTO_ReplacementEntityMetaData
// named metaData, are non-nil
func TestNewReplacementEntityMetaDataBuilder(t *testing.T) {
	assert := assert.New(t)
	newReplacementEMB := NewReplacementEntityMetaDataBuilder()
	if assert.NotNil(newReplacementEMB.metaData) {
		//if assert.NotEqual(nil, newReplacementEMB.metaData.IdentifyingProp) {
		if assert.NotNil(&newReplacementEMB.metaData.IdentifyingProp) {
			assert.Equal(0, len(newReplacementEMB.metaData.IdentifyingProp))
		}
	}
	if assert.NotNil(newReplacementEMB.metaData) {
		if assert.NotNil(&newReplacementEMB.metaData.BuyingCommTypes) {
			assert.Equal(0, len(newReplacementEMB.metaData.BuyingCommTypes))
		}
	}
	if assert.NotNil(newReplacementEMB.metaData) {
		if assert.NotNil(&newReplacementEMB.metaData.SellingCommTypes) {
			assert.Equal(0, len(newReplacementEMB.metaData.SellingCommTypes))
		}
	}
}

func TestReplacementEntityMetaDataBuilder_Build(t *testing.T) {
	assert := assert.New(t)
	replacementEMBnil := new(ReplacementEntityMetaDataBuilder)
	assert.Equal((*EntityDTO_ReplacementEntityMetaData)(nil), replacementEMBnil.metaData)
	replacementEntityMD := new(EntityDTO_ReplacementEntityMetaData)
	replacementEMB := &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMD,
	}
	metaData := replacementEMB.Build()
	assert.Equal(replacementEntityMD, metaData)
	assert.Equal(*replacementEntityMD, *metaData)
}

func TestMatching(t *testing.T) {
	assert := assert.New(t)
	replacementEntityMD := new(EntityDTO_ReplacementEntityMetaData)
	replacementEMB := &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMD,
	}
	propStr := rand.String(6)
	rEMB := replacementEMB.Matching(propStr)
	assert.Equal(propStr, rEMB.metaData.IdentifyingProp[0])
}
