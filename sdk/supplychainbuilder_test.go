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

// Tests that the Create function instantiates a new array of pointers to TemplateDTO
// Tests that the values in the map in the SupplyChainBuilder struct that calls this method
// are copied correctly to the array instantiated in Create() and the array is returned
// The array is empty if the map in this SupplyChainBuilder is nil
// The array contains one TemplateDTO pointer if the map in this had been instatiated and
// set with one key and value
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

func TestTop(t *testing.T) {
	assert := assert.New(t)
	scb := &SupplyChainBuilder{}
	class := new(EntityDTO_EntityType)
	template := &TemplateDTO{
		TemplateClass: class,
	}
	scnBuilder := &SupplyChainNodeBuilder{
		entityTemplate: template,
	}
	scBuilder := scb.Top(scnBuilder)
	assert.Equal(scBuilder.currentNode, scnBuilder)
	topNode_entity := scBuilder.currentNode.getEntity()
	// getEntity returns a EntityDTO_EntityType which is same as *class
	assert.Equal(topNode_entity, scnBuilder.getEntity())
	assert.Equal(1, len(scBuilder.SupplyChainNodes))
	// TODO : Fix below, the key is giving a nil value
	//assert.Equal(scBuilder.SupplyChainNodes[class_ran], scnBuilder)
}
