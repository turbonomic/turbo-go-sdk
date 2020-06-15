package supplychain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

func TestNewSupplyChainBuilder(t *testing.T) {
	expectedBuilder := &SupplyChainBuilder{}
	builder := NewSupplyChainBuilder()
	if !reflect.DeepEqual(expectedBuilder, builder) {
		t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
	}
}

func TestSupplyChainBuilder_Create(t *testing.T) {
	table := []struct {
		supplyChainNodes []*proto.TemplateDTO
		err              error

		expectsError bool
	}{
		{
			supplyChainNodes: []*proto.TemplateDTO{
				randomSupplyChainNode(),
				randomSupplyChainNode(),
			},
			expectsError: false,
		},
		{
			err:          fmt.Errorf("Fake"),
			expectsError: true,
		},
	}

	for _, item := range table {
		builder := &SupplyChainBuilder{
			supplyChainNodes: item.supplyChainNodes,
			err:              item.err,
		}
		supplyChainNodes, err := builder.Create()
		if item.expectsError {
			if err == nil {
				t.Errorf("Expected error, but got no error")
			}
		} else {
			expectedNodes := item.supplyChainNodes
			if !reflect.DeepEqual(expectedNodes, supplyChainNodes) {
				t.Errorf("\nExpected %+v, \ngot %+v", expectedNodes, supplyChainNodes)
			}
		}

	}
}

func TestSupplyChainBuilder_Top(t *testing.T) {
	table := []struct {
		topNode  *proto.TemplateDTO
		newError error
	}{
		{
			topNode: randomSupplyChainNode(),
		},
		{
			newError: fmt.Errorf("top node cannot be nil"),
		},
	}

	for _, item := range table {
		base := baseSupplyChainBuilder()
		builder := base.Top(item.topNode)
		var expectedNodes []*proto.TemplateDTO
		if item.topNode != nil {
			expectedNodes = append(expectedNodes, item.topNode)
		}
		expectedBuilder := &SupplyChainBuilder{
			supplyChainNodes: expectedNodes,
			err:              item.newError,
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestSupplyChainBuilder_Entity(t *testing.T) {
	table := []struct {
		newNodes                 []*proto.TemplateDTO
		existingSupplyChainNodes []*proto.TemplateDTO

		existingErr error
		newErr      error
	}{
		{
			newNodes: []*proto.TemplateDTO{
				randomSupplyChainNode(),
				randomSupplyChainNode(),
			},
			existingSupplyChainNodes: []*proto.TemplateDTO{
				randomSupplyChainNode(),
			},
			existingErr: fmt.Errorf("Fake"),
		},
		{
			newNodes: []*proto.TemplateDTO{randomSupplyChainNode()},
			newErr:   fmt.Errorf("must set top supply chain node first"),
		},
		{
			newNodes:                 []*proto.TemplateDTO{randomSupplyChainNode()},
			existingSupplyChainNodes: []*proto.TemplateDTO{},
			newErr:                   fmt.Errorf("must set top supply chain node first"),
		},
		{
			newNodes: []*proto.TemplateDTO{
				randomSupplyChainNode(),
				randomSupplyChainNode(),
			},
			existingSupplyChainNodes: []*proto.TemplateDTO{
				randomSupplyChainNode(),
			},
		},
	}
	for _, item := range table {
		base := baseSupplyChainBuilder()
		expectedBuilder := &SupplyChainBuilder{}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			if item.existingSupplyChainNodes != nil {
				base.supplyChainNodes = item.existingSupplyChainNodes
			}
			var expectedNodes []*proto.TemplateDTO
			if item.existingErr == nil && item.existingSupplyChainNodes != nil {
				if len(item.existingSupplyChainNodes) > 0 {
					if item.newErr == nil {
						expectedNodes = append(item.existingSupplyChainNodes, item.newNodes...)
					}
				} else {
					expectedNodes = item.existingSupplyChainNodes
				}
			}
			expectedErr := item.existingErr
			if item.newErr != nil {
				expectedErr = item.newErr
			}
			expectedBuilder.supplyChainNodes = expectedNodes
			expectedBuilder.err = expectedErr
		}

		builder := base
		for _, newNode := range item.newNodes {
			builder = builder.Entity(newNode)
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestSupplyChainBuilder_ValidateChargedBySold(t *testing.T) {
	builder := &SupplyChainBuilder{
		supplyChainNodes: []*proto.TemplateDTO{
			randomSupplyChainNodeWithChargedBySold(),
		},
	}
	supplyChainNodes, err := builder.Create()
	errMsg := "commodity HEAP of entity template APPLICATION_COMPONENT is charged by commodity " +
		"REMAINING_GC_CAPACITY which is not declared to be sold by the entity"
	assert.Nil(t, supplyChainNodes)
	assert.EqualError(t, err, errMsg)
}

func TestSupplyChainBuilder_ValidateChargedBy(t *testing.T) {
	builder := &SupplyChainBuilder{
		supplyChainNodes: []*proto.TemplateDTO{
			randomSupplyChainNodeWithChargedBy(),
			randomSupplyChainNode(),
		},
	}
	supplyChainNodes, err := builder.Create()
	assert.NotNil(t, supplyChainNodes)
	assert.NoError(t, err)
}

func baseSupplyChainBuilder() *SupplyChainBuilder {
	return &SupplyChainBuilder{}
}

func randomSupplyChainNode() *proto.TemplateDTO {
	node, _ := testBaseSupplyChainNodeBuilder().Create()
	return node
}

func randomSupplyChainNodeWithChargedBySold() *proto.TemplateDTO {
	heap := proto.CommodityDTO_HEAP
	collection := proto.CommodityDTO_REMAINING_GC_CAPACITY
	node, _ := NewSupplyChainNodeBuilder(proto.EntityDTO_APPLICATION_COMPONENT).
		Sells(&proto.TemplateCommodity{
			CommodityType: &heap,
			ChargedBySold: []proto.CommodityDTO_CommodityType{collection},
		}).Create()
	return node
}

func randomSupplyChainNodeWithChargedBy() *proto.TemplateDTO {
	threads := proto.CommodityDTO_THREADS
	vcpu := proto.CommodityDTO_VCPU
	builder := NewSupplyChainNodeBuilder(proto.EntityDTO_APPLICATION_COMPONENT).
		Sells(&proto.TemplateCommodity{
			CommodityType: &threads,
			ChargedBy:     []proto.CommodityDTO_CommodityType{vcpu},
		}).
		Provider(proto.EntityDTO_CONTAINER, proto.Provider_HOSTING).
		Buys(&proto.TemplateCommodity{CommodityType: &vcpu})
	node, _ := builder.Create()
	return node
}
