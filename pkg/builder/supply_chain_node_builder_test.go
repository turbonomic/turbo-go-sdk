package builder

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"testing"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	"math"
)

func TestNewSupplyChainNodeBuilder(t *testing.T) {
	entityType := randomEntityType()
	defaultTemplateType := proto.TemplateDTO_BASE
	defaultPriority := int32(0)
	expectedBuilder := &SupplyChainNodeBuilder{
		templateClass: &entityType,
		templateType:  &defaultTemplateType,
		priority:      &defaultPriority,
	}
	builder := NewSupplyChainNodeBuilder(entityType)
	if !reflect.DeepEqual(expectedBuilder, builder) {
		t.Errorf("Expected %++v, got %++v", expectedBuilder, builder)
	}
}

func TestSupplyChainNodeBuilder_Create(t *testing.T) {
	table := []struct {
		templateClass proto.EntityDTO_EntityType
		templateType  proto.TemplateDTO_TemplateType
		priority      int32

		err error

		expectsError bool
	}{
		{
			templateClass: randomEntityType(),
			templateType:  proto.TemplateDTO_BASE,
			priority:      0,

			expectsError: false,
		},
		{
			err:          fmt.Errorf("Fake Error"),
			expectsError: true,
		},
	}

	for _, item := range table {
		builder := &SupplyChainNodeBuilder{
			templateClass: &item.templateClass,
			templateType:  &item.templateType,
			priority:      &item.priority,
			err:           item.err,
		}
		supplyChainNode, err := builder.Create()
		if item.expectsError {
			if err == nil {
				t.Errorf("Expected error, but got no error")
			}
		} else {
			expectedTemplateDTO := &proto.TemplateDTO{
				TemplateClass:    &item.templateClass,
				TemplatePriority: &item.priority,
				TemplateType:     &item.templateType,
			}
			if !reflect.DeepEqual(expectedTemplateDTO, supplyChainNode) {
				t.Errorf("\nExpected %v, \ngot %v", expectedTemplateDTO, supplyChainNode)
			}
		}
	}
}

func TestSupplyChainNodeBuilder_Sells(t *testing.T) {
	table := []struct {
		templateCommoditiesSold []*proto.TemplateCommodity
		err                     error
	}{
		{
			templateCommoditiesSold: []*proto.TemplateCommodity{
				randomTemplateCommodity(),
				randomTemplateCommodity(),
			},
			err: nil,
		},
		{
			err: fmt.Errorf("Fake"),
		},
	}
	for _, item := range table {
		base := randomBaseSupplyChainNodeBuilder()
		if item.err != nil {
			base.err = item.err
		}
		expectedBuilder := &SupplyChainNodeBuilder{
			templateClass: base.templateClass,
			templateType:  base.templateType,
			priority:      base.priority,

			commoditiesSold: item.templateCommoditiesSold,
			err:             item.err,
		}

		builder := base
		for _, comm := range item.templateCommoditiesSold {
			builder = builder.Sells(comm)
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected %++v, \ngot %++v", expectedBuilder, builder)
		}
	}
}

func TestSupplyChainNodeBuilder_Buys(t *testing.T) {
	table := []struct {
		templateCommoditiesBought []*proto.TemplateCommodity
		provider                  *proto.Provider
		existingErr               error
		newErr                    error
	}{
		{
			existingErr: fmt.Errorf("Fake"),
		},
		{
			templateCommoditiesBought: []*proto.TemplateCommodity{
				randomTemplateCommodity(),
				randomTemplateCommodity(),
			},
			provider: randomProvider(),
		},
		{
			templateCommoditiesBought: []*proto.TemplateCommodity{
				randomTemplateCommodity(),
				randomTemplateCommodity(),
			},
			newErr: fmt.Errorf("Provider must be set before calling Buys()."),
		},
	}
	for _, item := range table {
		base := randomBaseSupplyChainNodeBuilder()
		if item.existingErr != nil {
			base.err = item.existingErr
		}

		var expectedMap map[*proto.Provider][]*proto.TemplateCommodity
		if item.provider != nil {
			base.currentProvider = item.provider
			expectedMap = make(map[*proto.Provider][]*proto.TemplateCommodity)
			expectedMap[item.provider] = item.templateCommoditiesBought
		}
		expectedErr := item.existingErr
		if item.newErr != nil {
			expectedErr = item.newErr
		}
		expectedBuilder := &SupplyChainNodeBuilder{
			templateClass: base.templateClass,
			templateType:  base.templateType,
			priority:      base.priority,

			providerCommodityBoughtMap: expectedMap,
			currentProvider:            item.provider,
			err:                        expectedErr,
		}
		builder := base
		for _, comm := range item.templateCommoditiesBought {
			builder = builder.Buys(comm)
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected %++v, \ngot      %++v", expectedBuilder, builder)
		}

	}
}

func TestBuildCommodityBought(t *testing.T) {
	table := []struct {
		providerList             []*proto.Provider
		allCommoditiesBoughtList [][]*proto.TemplateCommodity
	}{
		{
			providerList:             []*proto.Provider{},
			allCommoditiesBoughtList: [][]*proto.TemplateCommodity{},
		},
		{
			providerList: []*proto.Provider{
				randomProvider(),
				randomProvider(),
			},
			allCommoditiesBoughtList: [][]*proto.TemplateCommodity{
				{
					randomTemplateCommodity(),
				},
				{
					randomTemplateCommodity(),
					randomTemplateCommodity(),
				},
			},
		},
	}
	for _, item := range table {
		curMap := make(map[*proto.Provider][]*proto.TemplateCommodity)
		expectedPropsSet := make(map[*proto.Provider]map[*proto.TemplateCommodity]struct{})
		for _, provider := range item.providerList {
			for _, commList := range item.allCommoditiesBoughtList {
				curMap[provider] = commList

				if _, exist := expectedPropsSet[provider]; !exist {
					expectedPropsSet[provider] = make(map[*proto.TemplateCommodity]struct{})
				}
				for _, comm := range commList {
					expectedPropsSet[provider][comm] = struct{}{}
				}
			}
		}
		props := buildCommodityBought(curMap)
		for _, prop := range props {
			provider := prop.Key
			if _, exist := expectedPropsSet[provider]; !exist {
				t.Errorf("Unexpected provider %+v", prop.Key)
			}
			for _, comm := range prop.Value {
				if _, find := expectedPropsSet[provider][comm]; !find {
					t.Errorf("Unexpected commodity bought %+v", comm)
				}
				delete(expectedPropsSet[provider], comm)
				if len(expectedPropsSet[provider]) == 0 {
					delete(expectedPropsSet, provider)
				}
			}
		}

	}
}

func TestSupplyChainNodeBuilder_ConnectsTo(t *testing.T) {

}

func TestBuildExternalEntityLinkProperty(t *testing.T) {
	table := []struct {
		buyerRef   proto.EntityDTO_EntityType
		sellerRef  proto.EntityDTO_EntityType
		entityType proto.EntityDTO_EntityType
		key        proto.EntityDTO_EntityType

		expectsError bool
	}{
		{
			buyerRef:     proto.EntityDTO_VIRTUAL_MACHINE,
			sellerRef:    proto.EntityDTO_PHYSICAL_MACHINE,
			entityType:   proto.EntityDTO_VIRTUAL_MACHINE,
			key:          proto.EntityDTO_PHYSICAL_MACHINE,
			expectsError: false,
		},
		{
			sellerRef:    proto.EntityDTO_PHYSICAL_MACHINE,
			buyerRef:     proto.EntityDTO_VIRTUAL_MACHINE,
			entityType:   proto.EntityDTO_PHYSICAL_MACHINE,
			key:          proto.EntityDTO_VIRTUAL_MACHINE,
			expectsError: false,
		},
		{
			sellerRef:  proto.EntityDTO_PHYSICAL_MACHINE,
			buyerRef:   proto.EntityDTO_VIRTUAL_MACHINE,
			entityType: proto.EntityDTO_APPLICATION,

			expectsError: true,
		},
	}
	for _, item := range table {
		extLink := &proto.ExternalEntityLink{
			BuyerRef:  &item.buyerRef,
			SellerRef: &item.sellerRef,
		}
		builder := &SupplyChainNodeBuilder{
			templateClass: &item.entityType,
		}
		linkProp, err := builder.buildExternalEntityLinkProperty(extLink)
		if item.expectsError {
			if err == nil {
				t.Error("Expect error, got no error")
			}
		} else {
			expectedLinkProp := &proto.TemplateDTO_ExternalEntityLinkProp{
				Key:   &item.key,
				Value: extLink,
			}
			if !reflect.DeepEqual(linkProp, expectedLinkProp) {
				t.Errorf("\nExpected %++v, \ngot      %++v", expectedLinkProp, linkProp)
			}
		}
	}
}

func randomTemplateCommodity() *proto.TemplateCommodity {
	// a random commodity type.
	cType := randomCommodityType()
	// a random key
	key := rand.String(5)
	return &proto.TemplateCommodity{
		CommodityType: &cType,
		Key:           &key,
	}
}

func randomProvider() *proto.Provider {
	providerEntityType := randomEntityType()
	relationShip := randomProviderConsumerRelationship()
	maxCardinality := int32(math.MaxInt32)
	minCardinality := int32(0)
	return &proto.Provider{
		TemplateClass:  &providerEntityType,
		ProviderType:   &relationShip,
		CardinalityMax: &maxCardinality,
		CardinalityMin: &minCardinality,
	}
}

func randomProviderConsumerRelationship() proto.Provider_ProviderType {
	return proto.Provider_ProviderType(mathrand.Int31n(2))
}

// Create a random EntityDTOBuilder.
func randomBaseSupplyChainNodeBuilder() *SupplyChainNodeBuilder {
	return NewSupplyChainNodeBuilder(randomEntityType())
}
