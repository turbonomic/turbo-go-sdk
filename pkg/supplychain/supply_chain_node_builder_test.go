package supplychain

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

var (
	testEntityType       = proto.EntityDTO_WORKLOAD_CONTROLLER
	testProviderType     = proto.Provider_HOSTING
	maxCardinalityMaxInt = int32(math.MaxInt32)
	maxCardinalityOne    = int32(1)
	minCardinalityOne    = int32(1)
	minCardinalityZero   = int32(0)
	isProviderOptional   = true

	testProvider1 = &provider{
		provider: &proto.Provider{
			TemplateClass:  &testEntityType,
			ProviderType:   &testProviderType,
			CardinalityMax: &maxCardinalityOne,
			CardinalityMin: &minCardinalityOne,
		},
	}
	testProvider2 = &provider{
		provider: &proto.Provider{
			TemplateClass:  &testEntityType,
			ProviderType:   &testProviderType,
			CardinalityMax: &maxCardinalityOne,
			CardinalityMin: &minCardinalityZero,
		},
		optional: &isProviderOptional,
	}

	cType1  = proto.CommodityDTO_VCPU
	cType2  = proto.CommodityDTO_VMEM
	fakeKey = "fakeKey"

	testTemplateCommodity1 = &proto.TemplateCommodity{
		CommodityType: &cType1,
		Key:           &fakeKey,
	}
	testTemplateCommodity2 = &proto.TemplateCommodity{
		CommodityType: &cType2,
		Key:           &fakeKey,
	}
)

func TestNewSupplyChainNodeBuilder(t *testing.T) {
	entityType := rand.RandomEntityType()
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
			templateClass: rand.RandomEntityType(),
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
				testTemplateCommodity1,
				testTemplateCommodity2,
			},
			err: nil,
		},
		{
			templateCommoditiesSold: []*proto.TemplateCommodity{
				testTemplateCommodity1,
				testTemplateCommodity2,
			},
			err: fmt.Errorf("Fake"),
		},
	}
	for _, item := range table {
		base := testBaseSupplyChainNodeBuilder()
		expectedBuilder := &SupplyChainNodeBuilder{
			templateClass: base.templateClass,
			templateType:  base.templateType,
			priority:      base.priority,
		}
		if item.err != nil {
			base.err = item.err
			expectedBuilder.err = item.err
		} else {
			expectedBuilder.commoditiesSold = item.templateCommoditiesSold
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
		provider                  *provider
		existingErr               error
		newErr                    error
	}{
		{
			templateCommoditiesBought: []*proto.TemplateCommodity{
				testTemplateCommodity1,
				testTemplateCommodity2,
			},
			provider:    testProvider1,
			existingErr: fmt.Errorf("Fake"),
		},
		{
			templateCommoditiesBought: []*proto.TemplateCommodity{
				testTemplateCommodity1,
				testTemplateCommodity2,
			},
			provider: testProvider1,
		},
		{
			templateCommoditiesBought: []*proto.TemplateCommodity{
				testTemplateCommodity1,
				testTemplateCommodity2,
			},
			newErr: fmt.Errorf("Provider must be set before calling Buys()."),
		},
	}
	for _, item := range table {
		base := testBaseSupplyChainNodeBuilder()
		expectedBuilder := &SupplyChainNodeBuilder{
			templateClass: base.templateClass,
			templateType:  base.templateType,
			priority:      base.priority,
		}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {

			var expectedMap map[*provider][]*proto.TemplateCommodity
			if item.provider != nil {
				base.currentProvider = item.provider
				expectedMap = make(map[*provider][]*proto.TemplateCommodity)
				expectedMap[item.provider] = item.templateCommoditiesBought
			}
			expectedErr := item.existingErr
			if item.newErr != nil {
				expectedErr = item.newErr
			}

			expectedBuilder.providerCommodityBoughtMap = expectedMap
			expectedBuilder.currentProvider = item.provider
			expectedBuilder.err = expectedErr
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
		providerList             []*provider
		allCommoditiesBoughtList [][]*proto.TemplateCommodity
	}{
		{
			providerList:             []*provider{},
			allCommoditiesBoughtList: [][]*proto.TemplateCommodity{},
		},
		{
			providerList: []*provider{
				testProvider1,
				testProvider2,
			},
			allCommoditiesBoughtList: [][]*proto.TemplateCommodity{
				{
					testTemplateCommodity1,
				},
				{
					testTemplateCommodity1,
					testTemplateCommodity2,
				},
			},
		},
	}
	for _, item := range table {
		curMap := make(map[*provider][]*proto.TemplateCommodity)
		expectedPropsSet := make(map[*proto.Provider]map[*proto.TemplateCommodity]struct{})
		for _, provider := range item.providerList {
			for _, commList := range item.allCommoditiesBoughtList {
				curMap[provider] = commList

				if _, exist := expectedPropsSet[provider.provider]; !exist {
					expectedPropsSet[provider.provider] = make(map[*proto.TemplateCommodity]struct{})
				}
				for _, comm := range commList {
					expectedPropsSet[provider.provider][comm] = struct{}{}
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

func TestProviderOptOptionalHost(t *testing.T) {
	builder := NewSupplyChainNodeBuilder(proto.EntityDTO_CONTAINER_POD)
	isProviderOptional := true
	entityType := proto.EntityDTO_WORKLOAD_CONTROLLER
	providerType := proto.Provider_HOSTING
	builder.ProviderOpt(entityType, providerType, &isProviderOptional)
	expectedCurrentProvider := &provider{
		provider: &proto.Provider{
			TemplateClass:  &entityType,
			ProviderType:   &providerType,
			CardinalityMax: &maxCardinalityOne,
			CardinalityMin: &minCardinalityZero,
		},
		optional: &isProviderOptional,
	}
	if !reflect.DeepEqual(expectedCurrentProvider, builder.currentProvider) {
		t.Errorf("\nExpected %++v, \ngot      %++v", expectedCurrentProvider, builder.currentProvider)
	}
}

func TestProviderOptNonOptionalLayeredOver(t *testing.T) {
	builder := NewSupplyChainNodeBuilder(proto.EntityDTO_VIRTUAL_MACHINE)
	isProviderOptional := false
	entityType := proto.EntityDTO_STORAGE
	providerType := proto.Provider_LAYERED_OVER
	builder.ProviderOpt(entityType, providerType, &isProviderOptional)
	expectedCurrentProvider := &provider{
		provider: &proto.Provider{
			TemplateClass:  &entityType,
			ProviderType:   &providerType,
			CardinalityMax: &maxCardinalityMaxInt,
			CardinalityMin: &minCardinalityOne,
		},
		optional: &isProviderOptional,
	}
	if !reflect.DeepEqual(expectedCurrentProvider, builder.currentProvider) {
		t.Errorf("\nExpected %++v, \ngot      %++v", expectedCurrentProvider, builder.currentProvider)
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

// Create a test EntityDTOBuilder.
func testBaseSupplyChainNodeBuilder() *SupplyChainNodeBuilder {
	return NewSupplyChainNodeBuilder(proto.EntityDTO_CONTAINER_POD)
}

func doTestPriority(t proto.EntityDTO_EntityType, p int32) error {
	builder := NewSupplyChainNodeBuilder(t)
	builder.SetPriority(p)
	pp := p
	p = p + 1

	dto, err := builder.Create()
	if err != nil {
		return err
	}
	if dto.GetTemplatePriority() != pp {
		err := fmt.Errorf("Wrong priority %d Vs. %d", dto.GetTemplatePriority(), p)
		return err
	}

	return nil
}

func TestSupplyChainNodeBuilder_SetPriority_2(t *testing.T) {
	etypes := []proto.EntityDTO_EntityType{
		proto.EntityDTO_VIRTUAL_MACHINE,
		proto.EntityDTO_CONTAINER_POD,
		proto.EntityDTO_CONTAINER,
		proto.EntityDTO_APPLICATION,
		proto.EntityDTO_VIRTUAL_APPLICATION,
	}

	ps := []int32{0, 1, -1, 2, -2, 3, -3, 100, -100, 1000, -1000}
	for _, p := range ps {
		for _, et := range etypes {
			err := doTestPriority(et, p)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
		}
	}
}

func TestSupplyChainNodeBuilder_SetPriority(t *testing.T) {
	builder := NewSupplyChainNodeBuilder(proto.EntityDTO_VIRTUAL_MACHINE)
	p := int32(-8)
	builder.SetPriority(p)

	dto, err := builder.Create()
	if err != nil {
		t.Errorf("Create NodeTemplate failed: %v", err)
		return
	}
	if dto.GetTemplatePriority() != p {
		t.Errorf("Wrong priority %d Vs. %d", dto.GetTemplatePriority(), p)
	}
}

func TestSupplyChainNodeBuilder_SetTemplateType(t *testing.T) {
	builder := NewSupplyChainNodeBuilder(proto.EntityDTO_VIRTUAL_MACHINE)
	tt := proto.TemplateDTO_EXTENSION
	builder.SetTemplateType(tt)

	dto, err := builder.Create()
	if err != nil {
		t.Errorf("Create node Template failed: %v", err)
		return
	}

	if dto.GetTemplateType() != tt {
		t.Errorf("Failed to set template type: %v Vs. %v", dto.GetTemplateType(), tt)
	}
}
