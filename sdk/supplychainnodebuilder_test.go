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

// Tests that the getEntity() method returns the TemplateClass member variable in the
// struct point at by this.entityTemplate
func TestgetEntity(t *testing.T) {
	assert := assert.New(t)
	class := new(EntityDTO_EntityType)
	templateDTO := &TemplateDTO{
		TemplateClass: class,
	}
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: templateDTO,
	}
	entityDTO_EntityType := scnbuilder.getEntity()
	assert.Equal(class, &entityDTO_EntityType)
	assert.Equal(*class, entityDTO_EntityType)
}

// Tests that requireEntityTemplate returns true for the case when
// this.entityTemplate is not nil
func TestrequireEntityTemplate_notnil(t *testing.T) {
	assert := assert.New(t)
	templateDTO := new(TemplateDTO)
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: templateDTO,
	}
	entityTemplate := scnbuilder.requireEntityTemplate()
	assert.Equal(true, entityTemplate)
}

// Tests that requireEntityTemplate returns false for the case when
// this.entityTemplate is nil
func TestrequireEntityTemplate_nil(t *testing.T) {
	assert := assert.New(t)
	scnbuilder := &SupplyChainNodeBuilder{
	//entityTemplate is nil
	}
	entityTemplate := scnbuilder.requireEntityTemplate()
	assert.Equal(false, entityTemplate)
}
