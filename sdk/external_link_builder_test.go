package sdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	//	"reflect"
	"testing"
)

// Tests that the NewExternalEntityBuilder() function creates a new ExternalEntityLinkBuilder,
// instantiates an ExternalEntityLink and uses it to set the entityLink member variable of the
// ExternalEntityLinkBuilder returned
func TestNewExternalEntityBuilder(t *testing.T) {
	assert := assert.New(t)
	newEELB := NewExternalEntityLinkBuilder()
	assert.NotNil(newEELB)
	//	if assert.NotNil(newEELB.entityLink) {
	//		assert.Equal(*ExternalEntityLink, reflect.TypeOf(newEELB.entityLink))
	//	}
}

// Tests that the buyer, seller and relationship passed to the Link method are used to set
// this.entityLink.BuyerRef, this.entityLink.SellerRef and this.entityLink.Relationship
func TestLink(t *testing.T) {
	assert := assert.New(t)
	externalEntityLink := new(ExternalEntityLink)
	externalEntityLinkBuilder := &ExternalEntityLinkBuilder{
		entityLink: externalEntityLink,
	}
	buyer := new(EntityDTO_EntityType)
	seller := new(EntityDTO_EntityType)
	relationship := new(Provider_ProviderType)
	eelb := externalEntityLinkBuilder.Link(*buyer, *seller, *relationship)
	assert.Equal(buyer, eelb.entityLink.BuyerRef)
	assert.Equal(seller, eelb.entityLink.SellerRef)
	assert.Equal(relationship, eelb.entityLink.Relationship)
}

// Tests that the CommodityDTO_CommodityType passed to the Commodity() method
// is appended to the array this.entityLink.Commodities
func TestCommodity(t *testing.T) {
	assert := assert.New(t)
	externalEntityLink := new(ExternalEntityLink)
	externalEntityLinkBuilder := &ExternalEntityLinkBuilder{
		entityLink: externalEntityLink,
	}
	assert.Equal(0, len(externalEntityLink.Commodities))
	comm := new(CommodityDTO_CommodityType)
	eelb := externalEntityLinkBuilder.Commodity(*comm)
	assert.Equal(1, len(externalEntityLink.Commodities))
	assert.Equal(comm, &eelb.entityLink.Commodities[0])
	assert.Equal(*comm, eelb.entityLink.Commodities[0])
}

// Tests that the name and description string arguments passed to ProbeEntityPropertyDef() are
// used to set the Name and Description member variables of a newly created ExternalEntityLink_EntityPropertyDef
// struct .
// Tests that the created struct is appended to this.entityLink.ProbeEntityPropertyDef array
func TestProbeEntityPropertyDef(t *testing.T) {
	assert := assert.New(t)
	externalEntityLink := new(ExternalEntityLink)
	externalEntityLinkBuilder := &ExternalEntityLinkBuilder{
		entityLink: externalEntityLink,
	}
	name := rand.String(6)
	description := rand.String(6)
	assert.Equal(0, len(externalEntityLink.ProbeEntityPropertyDef))
	eelb := externalEntityLinkBuilder.ProbeEntityPropertyDef(name, description)
	if assert.Equal(1, len(externalEntityLink.ProbeEntityPropertyDef)) {
		assert.Equal(&name, eelb.entityLink.ProbeEntityPropertyDef[0].Name)
		assert.Equal(name, *eelb.entityLink.ProbeEntityPropertyDef[0].Name)
		assert.Equal(&description, eelb.entityLink.ProbeEntityPropertyDef[0].Description)
		assert.Equal(description, *eelb.entityLink.ProbeEntityPropertyDef[0].Description)
	}
}

// Tests that the ExternalEntityLink_ServerEntityPropDef passed to the method ExternalEntityPropertyDef
// is appended to this.entityLink.ExternalEntityPropertyDefs
func TestExternalEntityPropertyDef(t *testing.T) {
	assert := assert.New(t)
	externalEntityLink := new(ExternalEntityLink)
	externalEntityLinkBuilder := &ExternalEntityLinkBuilder{
		entityLink: externalEntityLink,
	}
	propertyDef := new(ExternalEntityLink_ServerEntityPropDef)
	eelb := externalEntityLinkBuilder.ExternalEntityPropertyDef(propertyDef)
	assert.Equal(1, len(eelb.entityLink.ExternalEntityPropertyDefs))
	assert.Equal(propertyDef, eelb.entityLink.ExternalEntityPropertyDefs[0])
}

// Tests that the Build method returns the entityLink member of this
func TestBuild(t *testing.T) {
	assert := assert.New(t)
	externalEntityLink := new(ExternalEntityLink)
	externalEntityLinkBuilder := &ExternalEntityLinkBuilder{
		entityLink: externalEntityLink,
	}
	el := externalEntityLinkBuilder.Build()
	assert.Equal(externalEntityLink, el)
	assert.Equal(*externalEntityLink, *el)
}
