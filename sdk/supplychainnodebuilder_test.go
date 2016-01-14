package sdk

import (
	"github.com/stretchr/testify/assert"
	//	"math"
	"testing"
)

// Tests that the function NewSupplyChainBuilder() returns an initialized SupplyChainNodeBuilder struct
func TestNewSupplyChainNodeBuilder(t *testing.T) {
	assert := assert.New(t)
	newscNodeBuilder := NewSupplyChainNodeBuilder()
	assert.Equal((*TemplateDTO)(nil), newscNodeBuilder.entityTemplate)
	assert.Equal((*Provider)(nil), newscNodeBuilder.currentProvider)
}

// Tests that Create() method returns a pointer to a TemplateDTO struct
func TestSupplyChainNodeBuilderCreate(t *testing.T) {
	assert := assert.New(t)
	entityTemplate := new(TemplateDTO)
	supplycnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: entityTemplate,
	}
	entityT := supplycnbuilder.Create()
	assert.Equal(entityTemplate, entityT)
}

// Tests that Entity() method creates a new TemplateDTO struct and assigns it to the calling
// SuppyChainBuilder's entityTemplate variable
// Tests that the new TemplateDTO has member variable TemplateClass set to the EntityDTO_EntityType
// passed to the Entity() method
func TestEntity(t *testing.T) {
	assert := assert.New(t)
	scnbuilder := &SupplyChainNodeBuilder{}
	entityType := new(EntityDTO_EntityType)
	scnb := scnbuilder.Entity(*entityType)
	if assert.NotNil(scnb.entityTemplate) {
		assert.Equal(*entityType, *scnb.entityTemplate.TemplateClass)
		assert.Equal(entityType, scnb.entityTemplate.TemplateClass)
		assert.Equal(TemplateDTO_BASE, *scnb.entityTemplate.TemplateType)
		//	assert.Equal(&TemplateDTO_BASE, scnb.entityTemplate.TemplateType)
		assert.Equal(int32(0), *scnb.entityTemplate.TemplatePriority)
		assert.Equal(false, *scnb.entityTemplate.AutoCreate)
		if assert.NotNil(t, scnb.entityTemplate.CommoditySold) {
			assert.Equal(0, len(scnb.entityTemplate.CommoditySold))
		}
		if assert.NotNil(t, scnb.entityTemplate.CommodityBought) {
			assert.Equal(0, len(scnb.entityTemplate.CommodityBought))
		}
	}

}
