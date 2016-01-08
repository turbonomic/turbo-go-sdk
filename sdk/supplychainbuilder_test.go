package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests that the NewSupplyChainBuilder() method returns a SupplyChainBuilder struct
// with a nil pointer variable named currentNode and a nil map variable called SupplyChainNodes
func TestNewSupplyChainBuilder(t *testing.T) {
	assert := assert.New(t)
	newSCB := NewSupplyChainBuilder()
	assert.Equal((*SupplyChainNodeBuilder)(nil), newSCB.currentNode)
	assert.Equal((map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)(nil), newSCB.SupplyChainNodes)
}

func TestSupplyChainBuilder_Create(t *testing.T) {
	assert := assert.New(t)
	scb := &SupplyChainBuilder{
	//      SupplyChainNodes is nil for now
	}
	allNodes := scb.Create()
	assert.Equal(0, len(allNodes))
	scn := make(map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)
	scb_map := &SupplyChainBuilder{
		SupplyChainNodes: scn,
	}
	entityDTO := new(EntityDTO_EntityType)
	supplyTest := new(SupplyChainNodeBuilder)
	scb_map.SupplyChainNodes[entityDTO] = supplyTest
	template := scb_map.Create()
	if assert.Equal(1, len(template)) {
		assert.Equal(supplyTest.entityTemplate, template[0])
	}
}
