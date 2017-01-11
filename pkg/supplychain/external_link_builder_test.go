package supplychain

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

func TestExternalEntityLinkBuilder_Build(t *testing.T) {
	table := []struct {
		existingErr error
		expectError bool
	}{
		{
			existingErr: fmt.Errorf("Fake"),
			expectError: true,
		},
		{
			expectError: false,
		},
	}

	for _, item := range table {
		base := randomExternalEntityLinkBuilder()
		if item.existingErr != nil {
			base.err = item.existingErr
		}
		externalLink, err := base.Build()
		if item.expectError {
			if err == nil {
				t.Errorf("Expected erro, got no error.")
			}
		} else {
			expectedExternalLink := &proto.ExternalEntityLink{
				BuyerRef:     base.buyerRef,
				SellerRef:    base.sellerRef,
				Relationship: base.relationship,
			}
			if !reflect.DeepEqual(expectedExternalLink, externalLink) {
				t.Errorf("Expected %+v, got %+v", expectedExternalLink, externalLink)
			}
		}
	}
}

func TestExternalEntityLinkBuilder_Link(t *testing.T) {
	table := []struct {
		buyer        proto.EntityDTO_EntityType
		seller       proto.EntityDTO_EntityType
		relationship proto.Provider_ProviderType

		existingErr error
	}{
		{
			buyer:        rand.RandomEntityType(),
			seller:       rand.RandomEntityType(),
			relationship: rand.RandomProviderConsumerRelationship(),
		},
		{
			buyer:        rand.RandomEntityType(),
			seller:       rand.RandomEntityType(),
			relationship: rand.RandomProviderConsumerRelationship(),
			existingErr:  fmt.Errorf("Error!"),
		},
	}

	for _, item := range table {
		base := &ExternalEntityLinkBuilder{}
		expectedLinkBuilder := &ExternalEntityLinkBuilder{}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedLinkBuilder.err = item.existingErr
		} else {
			expectedLinkBuilder.buyerRef = &item.buyer
			expectedLinkBuilder.sellerRef = &item.seller
			expectedLinkBuilder.relationship = &item.relationship
		}
		builder := base.Link(item.buyer, item.seller, item.relationship)
		if !reflect.DeepEqual(expectedLinkBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedLinkBuilder, builder)
		}

	}
}

func TestExternalEntityLinkBuilder_Commodity(t *testing.T) {
	table := []struct {
		comm   proto.CommodityDTO_CommodityType
		hasKey bool

		existingErr error
	}{
		{
			comm:   rand.RandomCommodityType(),
			hasKey: true,
		},
		{
			comm:   rand.RandomCommodityType(),
			hasKey: false,
		},
		{
			comm:        rand.RandomCommodityType(),
			hasKey:      false,
			existingErr: fmt.Errorf("Error!"),
		},
		{
			comm:        rand.RandomCommodityType(),
			hasKey:      true,
			existingErr: fmt.Errorf("Error!"),
		},
	}

	for _, item := range table {
		base := &ExternalEntityLinkBuilder{}
		expectedLinkBuilder := &ExternalEntityLinkBuilder{}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedLinkBuilder.err = item.existingErr
		} else {
			expectedLinkBuilder.commodityDefs = []*proto.ExternalEntityLink_CommodityDef{
				{
					Type:   &item.comm,
					HasKey: &item.hasKey,
				},
			}
		}
		builder := base.Commodity(item.comm, item.hasKey)
		if !reflect.DeepEqual(expectedLinkBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedLinkBuilder, builder)
		}

	}
}

func TestExternalEntityLinkBuilder_ProbeEntityPropertyDef(t *testing.T) {
	table := []struct {
		name string
		desc string

		existingErr error
	}{
		{
			name: rand.String(5),
			desc: rand.String(5),
		},
		{
			name:        rand.String(5),
			desc:        rand.String(5),
			existingErr: fmt.Errorf("Error!"),
		},
	}

	for _, item := range table {
		base := &ExternalEntityLinkBuilder{}
		expectedLinkBuilder := &ExternalEntityLinkBuilder{}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedLinkBuilder.err = item.existingErr
		} else {
			expectedLinkBuilder.probeEntityPropertyDef = []*proto.ExternalEntityLink_EntityPropertyDef{
				{
					Name:        &item.name,
					Description: &item.desc,
				},
			}
		}
		builder := base.ProbeEntityPropertyDef(item.name, item.desc)
		if !reflect.DeepEqual(expectedLinkBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedLinkBuilder, builder)
		}

	}
}

func TestExternalEntityLinkBuilder_ExternalEntityPropertyDef(t *testing.T) {
	table := []struct {
		propertyDef *proto.ExternalEntityLink_ServerEntityPropDef
	}{
		{
			propertyDef: nil,
		},
		{
			propertyDef:rand.RandomExternalEntityLink_ServerEntityPropDef(),
		},
	}

	for _, item := range table {
		base := &ExternalEntityLinkBuilder{}
		expectedLinkBuilder := &ExternalEntityLinkBuilder{}
		if item.propertyDef != nil {
			expectedLinkBuilder.externalEntityPropertyDefs = []*proto.ExternalEntityLink_ServerEntityPropDef{
				item.propertyDef,
			}
		} else {
			expectedLinkBuilder.err = fmt.Errorf("Nil service entity property definition.")
		}
		builder := base.ExternalEntityPropertyDef(item.propertyDef)
		if !reflect.DeepEqual(expectedLinkBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedLinkBuilder, builder)
		}

	}
}

func randomExternalEntityLinkBuilder() *ExternalEntityLinkBuilder {
	buyerRef := rand.RandomEntityType()
	sellerRef := rand.RandomEntityType()
	relationship := rand.RandomProviderConsumerRelationship()
	return &ExternalEntityLinkBuilder{
		buyerRef:     &buyerRef,
		sellerRef:    &sellerRef,
		relationship: &relationship,
	}
}