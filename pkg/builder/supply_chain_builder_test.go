package builder

import (
	"fmt"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
	"reflect"
	"testing"
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
			newError: fmt.Errorf("topNode cannot be nil."),
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
			existingErr: fmt.Errorf("Fake"),
		},
		{
			newNodes: []*proto.TemplateDTO{randomSupplyChainNode()},
			newErr:   fmt.Errorf("Must set top supply chain node first."),
		},
		{
			newNodes:                 []*proto.TemplateDTO{randomSupplyChainNode()},
			existingSupplyChainNodes: []*proto.TemplateDTO{},
			newErr: fmt.Errorf("Must set top supply chain node first."),
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
		if item.existingErr != nil {
			base.err = item.existingErr
		}
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
		expectedBuilder := &SupplyChainBuilder{
			supplyChainNodes: expectedNodes,
			err:              expectedErr,
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

func baseSupplyChainBuilder() *SupplyChainBuilder {
	return &SupplyChainBuilder{}
}

func randomSupplyChainNode() *proto.TemplateDTO {
	node, _ := randomBaseSupplyChainNodeBuilder().Create()
	return node
}
