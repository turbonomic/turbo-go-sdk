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

// Tests that the Build() method returns the correct metaData member variable for the
// ReplacementEntityMetaDataBuilder which calls it
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

// Tests that the string passed to Matching() method is appended to the string array
// variable called IdentifyProp in this.metaData
func TestMatching(t *testing.T) {
	assert := assert.New(t)
	replacementEntityMD := new(EntityDTO_ReplacementEntityMetaData)
	replacementEMB := &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMD,
	}
	propStr := rand.String(6)
	rEMB := replacementEMB.Matching(propStr)
	assert.Equal(propStr, rEMB.metaData.IdentifyingProp[0])
	assert.Equal(&propStr, &rEMB.metaData.IdentifyingProp[0])
}

// Tests that the CommodityDTO_CommodityType passed to PatchBuying is appended to this.metaData.BuyingCo
//
func TestPatchBuying(t *testing.T) {
	assert := assert.New(t)
	replacementEntityMD := new(EntityDTO_ReplacementEntityMetaData)
	replacementEMB := &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMD,
	}
	commType := new(CommodityDTO_CommodityType)
	rEMB := replacementEMB.PatchBuying(*commType)
	assert.Equal(commType, &rEMB.metaData.BuyingCommTypes[0])
	assert.Equal(*commType, rEMB.metaData.BuyingCommTypes[0])
}

// Tests that the CommodityDTO_CommodityType passed to PatchSelling is appended
// to this.metaData.SellingCommTypes
func TestPatchSelling(t *testing.T) {
	assert := assert.New(t)
	replacementEntityMD := new(EntityDTO_ReplacementEntityMetaData)
	replacementEMB := &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMD,
	}
	commType := new(CommodityDTO_CommodityType)
	rEMB := replacementEMB.PatchSelling(*commType)
	assert.Equal(commType, &rEMB.metaData.SellingCommTypes[0])

}
