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

// Tests that a map is initialized and assigned to this.SupplyChainNodes
// Tests that the SupplyChainNodeBuilder passed to the Top() method
// is assigned to an entry in the recently initialized map
// Tests that this.currentNode is set to the argument to Top() method
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
	assert.NotEqual(nil, scBuilder.SupplyChainNodes)
	topNode_entity := scBuilder.currentNode.getEntity()
	// getEntity returns a EntityDTO_EntityType which is same as *class
	assert.Equal(topNode_entity, scnBuilder.getEntity())
	assert.Equal(1, len(scBuilder.SupplyChainNodes))
	// TODO : Fix below, the key is giving a nil value
	//assert.Equal(scBuilder.SupplyChainNodes[class_ran], scnBuilder)
	assert.Equal(scBuilder.currentNode, scnBuilder)
}

// Tests that a map is already initialized and assigned to this.SupplyChainNodes
// before the method Entity is called
// Tests that the SupplyChainNodeBuilder passed to the Entity() method
// is assigned to an entry in the map this.SupplyChainNodes
// Tests that this.currentNode is set to the argument to Entity() method
func TestEntity_hasTopTrue(t *testing.T) {
	assert := assert.New(t)
	supplyChainMap := make(map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)
	supplyCB := &SupplyChainBuilder{SupplyChainNodes: supplyChainMap}
	//assert.Equal((map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)(nil), supplyCB.SupplyChainNodes)
	node := new(SupplyChainNodeBuilder)
	scb := supplyCB.Entity(node)
	assert.NotEqual(nil, scb.SupplyChainNodes)
	if assert.NotEqual(nil, scb.SupplyChainNodes) {
		assert.Equal(1, len(scb.SupplyChainNodes))
	}
	assert.Equal(scb.currentNode, node)
}

// Tests that a map is not initialized and assigned to this.SupplyChainNodes
// before the method Entity is called
func TestEntity_hasTopFalse(t *testing.T) {
	assert := assert.New(t)
	supplyCB := &SupplyChainBuilder{}
	assert.Equal((map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)(nil), supplyCB.SupplyChainNodes)
}

// Tests that method hasTopNode() returns true when the member variable this.SupplyChainNodes is not nil
func TestHasTopNode_truecase(t *testing.T) {
	assert := assert.New(t)
	supplyChainNodes := make(map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)
	supplyCB := &SupplyChainBuilder{
		SupplyChainNodes: supplyChainNodes,
	}
	hasSCNodes := supplyCB.hasTopNode()
	if assert.NotEqual((map[*EntityDTO_EntityType]*SupplyChainNodeBuilder)(nil), supplyCB.SupplyChainNodes) {
		assert.Equal(true, hasSCNodes)
	}
}

// Tests that method hasTopNode() returns false when the member variable this.SupplyChainNodes is nil
func TestHasTopNode_falsecase(t *testing.T) {
	assert := assert.New(t)
	supplyCB := &SupplyChainBuilder{}
	hasSCNodes := supplyCB.hasTopNode()
	assert.Equal(false, hasSCNodes)
}

// Tests that the ExternalEntityLink passed to ConnectsTo method is appended to
// this.currentNode.entityTemplate.ExternalLink array when error from scb.requireCurrentNode()
// is not nil
func TestConnectsTo_nilerror(t *testing.T) {
	assert := assert.New(t)
	entityTemplate := new(TemplateDTO)
	currentnode := &SupplyChainNodeBuilder{
		entityTemplate: entityTemplate,
	}
	scbuilder := &SupplyChainBuilder{
		currentNode: currentnode,
	}
	extentityLink := new(ExternalEntityLink)
	scb := scbuilder.ConnectsTo(extentityLink)
	// currentNode.Link(st)  => scnb.entityTemplate.ExternalLink[0].Value == st
	if assert.NotNil(scb.currentNode) {
		assert.Equal(scb.currentNode.entityTemplate.ExternalLink[0].Value, extentityLink)
	}
}

// Tests that scb.requireCurrentNode() returns a non nil error when this.currentNode is nil
func TestConnectsTo_error(t *testing.T) {
	assert := assert.New(t)
	scbuilder := &SupplyChainBuilder{}
	extentityLink := new(ExternalEntityLink)
	scb := scbuilder.ConnectsTo(extentityLink)
	assert.Equal((*SupplyChainNodeBuilder)(nil), scb.currentNode)
}

// Tests that requireCurrentNode() returns nil when currentNode is not nil
func TestRequireCurrentNode_nil(t *testing.T) {
	assert := assert.New(t)
	currentnode := new(SupplyChainNodeBuilder)
	scbuilder := &SupplyChainBuilder{
		currentNode: currentnode,
	}
	scb := scbuilder.requireCurrentNode()
	assert.Equal(nil, scb)
}

// Tests that requireCurrentNode() returns not nil when currentNode is nil
func TestRequireCurrentNode_notnil(t *testing.T) {
	assert := assert.New(t)
	scbuilder := &SupplyChainBuilder{}
	scb := scbuilder.requireCurrentNode()
	assert.NotEqual(nil, scb)
}
