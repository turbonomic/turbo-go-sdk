package sdk

import (
	"github.com/stretchr/testify/assert"
	//	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	//	mathrand "math/rand"
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
