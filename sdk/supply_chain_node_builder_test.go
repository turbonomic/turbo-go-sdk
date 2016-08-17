package sdk

import (
	"github.com/stretchr/testify/assert"
	"math"
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
func TestGetEntity(t *testing.T) {
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
func TestRequireEntityTemplate_notnil(t *testing.T) {
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
func TestRequireEntityTemplate_nil(t *testing.T) {
	assert := assert.New(t)
	scnbuilder := &SupplyChainNodeBuilder{
	//entityTemplate is nil
	}
	entityTemplate := scnbuilder.requireEntityTemplate()
	assert.Equal(false, entityTemplate)
}

// Tests that when the member variable this.currentProvider is not nil
// then requireProvider() method returns true
func TestRequireProvider_notnil(t *testing.T) {
	assert := assert.New(t)
	provider := new(Provider)
	scnbuilder := &SupplyChainNodeBuilder{
		currentProvider: provider,
	}
	requireProvider := scnbuilder.requireProvider()
	assert.Equal(true, requireProvider)
}

// Tests that when the member variable this.currentProvider is not nil
// then requireProvider() method returns true
func TestRequireProvider_nil(t *testing.T) {
	assert := assert.New(t)
	scnbuilder := &SupplyChainNodeBuilder{}
	requireProvider := scnbuilder.requireProvider()
	assert.Equal(false, requireProvider)
}

// Tests that a new TemplateCommodity is created  and that its member variables Key
// CommodityType are set to "" and the pointer to CommodityDTO_CommodityType passed
// to Selling method
func TestSelling(t *testing.T) {
	assert := assert.New(t)
	templateDTO := &TemplateDTO{}
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: templateDTO,
	}
	comm := new(CommodityDTO_CommodityType)
	scnb := scnbuilder.Selling(*comm)
	// make sure the CommoditySold array is not nil
	assert.NotEqual(([]*TemplateCommodity)(nil), scnb.entityTemplate.CommoditySold)
	if assert.Equal(1, len(scnb.entityTemplate.CommoditySold)) {
		assert.Equal(comm, scnb.entityTemplate.CommoditySold[0].CommodityType)
		assert.Equal(*comm, *scnb.entityTemplate.CommoditySold[0].CommodityType)
	}
}

// Tests that the Provider function creates a new Provider struct
// Tests that the new Provider struct sets its member variables TemplateClass and
// ProviderType to the pointer of the arguments passed to Provider method
// Tests the CadinalityMax and CardinalityMin for the pType = Provider_LAYERED_OVER case
func TestProvider_LAYERED_OVER(t *testing.T) {
	assert := assert.New(t)
	template := &TemplateDTO{}
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: template,
	}
	provider := new(EntityDTO_EntityType)
	pType := Provider_LAYERED_OVER
	scnb := scnbuilder.Provider(*provider, pType)
	assert.Equal(provider, scnb.currentProvider.TemplateClass)
	assert.Equal(pType, *scnb.currentProvider.ProviderType)
	assert.Equal(int32(math.MaxInt32), *scnb.currentProvider.CardinalityMax)
	assert.Equal(int32(0), *scnb.currentProvider.CardinalityMin)
}

//Tests that the Provider function creates a new Provider struct
// Tests that the new Provider struct sets its member variables TemplateClass and
// ProviderType to the pointer of the arguments passed to Provider method
// Tests the CadinalityMax and CardinalityMin for the pType = Provider_HOSTING case
func TestProvider_HOSTING(t *testing.T) {
	assert := assert.New(t)
	template := &TemplateDTO{}
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: template,
	}
	provider := new(EntityDTO_EntityType)
	pType := Provider_HOSTING
	scnb := scnbuilder.Provider(*provider, pType)
	assert.Equal(provider, scnb.currentProvider.TemplateClass)
	assert.Equal(pType, *scnb.currentProvider.ProviderType)
	assert.Equal(int32(1), *scnb.currentProvider.CardinalityMax)
	assert.Equal(int32(1), *scnb.currentProvider.CardinalityMin)
}

// For the case when exist = true
// Tests that the TemplateCommodity passed to the Buys() method is appended to the
// member array named Value in the TemplateDTO_CommBoughtProviderProp having the
//  Key = this.currentProvider
func TestBuys_exist_true(t *testing.T) {
	assert := assert.New(t)
	provider := &Provider{
	// nothing here for now
	}
	//	var value []*TemplateCommodity
	// this will be the case when exist = true
	templateDTO_CBPP := &TemplateDTO_CommBoughtProviderProp{
		Key: provider,
		//	Value:value,
	}
	template := new(TemplateDTO)
	commBought := append(template.CommodityBought, templateDTO_CBPP)
	template.CommodityBought = commBought
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate:  template,
		currentProvider: provider,
	}
	templateCommodity := new(TemplateCommodity)
	scnb := scnbuilder.Buys(*templateCommodity)
	commoditiesBought := scnb.entityTemplate.CommodityBought
	var providerProp *TemplateDTO_CommBoughtProviderProp
	for _, pp := range commoditiesBought {
		if pp.GetKey() == provider {
			providerProp = pp
		}
	}
	// provideProp is the  *TemplateDTO_CommBoughtProviderProp
	// containing the []*TemplateCommodity to which Buys appended templateCommodity
	assert.Equal(1, len(scnb.entityTemplate.CommodityBought))
	if assert.NotEqual(([]*TemplateCommodity)(nil), commoditiesBought[0].Value) {
		assert.Equal(1, len(providerProp.Value))
		assert.Equal(templateCommodity, providerProp.Value[0])

	}
}

// For the case when exist = false
// Tests that a new TemplateDTO_CommBoughtProviderProp struct is created
// with it Key variable = this.currentProvider
// Tests that the new TemplateDTO_CommBoughtProviderProp struct is appended
// to this.entityTemplate.CommodityBought
// Tests that the TemplateCommodity passed to the Buys() method is appended to the
// member array named Value in the newly created TemplateDTO_CommBoughtProviderProp
func TestBuys_exist_false(t *testing.T) {
	assert := assert.New(t)
	eTemplate := new(TemplateDTO)
	provider := new(Provider)
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate:  eTemplate,
		currentProvider: provider,
	}
	templateCommodity := new(TemplateCommodity)
	scnb := scnbuilder.Buys(*templateCommodity)
	var providerProp *TemplateDTO_CommBoughtProviderProp
	var commoditiesBought []*TemplateDTO_CommBoughtProviderProp
	if assert.NotEqual(([]*TemplateDTO_CommBoughtProviderProp)(nil), scnb.entityTemplate.CommodityBought) {
		assert.Equal(1, len(scnb.entityTemplate.CommodityBought))
		commoditiesBought = scnb.entityTemplate.CommodityBought
		for _, pp := range commoditiesBought {
			if pp.GetKey() == provider {
				providerProp = pp
			}
		}
	}
	// provideProp is the  *TemplateDTO_CommBoughtProviderProp
	// containing the []*TemplateCommodity to which Buys appended templateCommodity
	if assert.NotEqual(([]*TemplateCommodity)(nil), commoditiesBought[0].Value) {
		assert.Equal(1, len(providerProp.Value))
		assert.Equal(templateCommodity, providerProp.Value[0])
	}
}

// Tests that findCommBoughtProvider returns the TemplateDTO_CommBoughtProviderProp and true when
// the argument to it matches the memeber variable named
// Key of one of the elements in the CommodityBought array
func TestFindCommBoughtProvider_true(t *testing.T) {
	assert := assert.New(t)
	template := new(TemplateDTO)
	provider := new(Provider)
	providerprop := &TemplateDTO_CommBoughtProviderProp{
		Key: provider,
	}
	commbought := append(template.CommodityBought, providerprop)
	template.CommodityBought = commbought
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate:  template,
		currentProvider: provider,
	}
	commboughtProviderProp, isFound := scnbuilder.findCommBoughtProvider(provider)
	assert.Equal(true, isFound)
	assert.Equal(providerprop, commboughtProviderProp)
}

// Tests that findCommBoughtProvider returns nil and false when
// the argument to it does not match any of the member variables named
// Key of any of the elements in the this.entityTemplate.CommodityBought array
func TestFindCommBoughtProvider_false(t *testing.T) {
	assert := assert.New(t)
	template := new(TemplateDTO)
	provider1 := new(Provider)
	providerprop := &TemplateDTO_CommBoughtProviderProp{
		Key: provider1,
	}
	commbought := append(template.CommodityBought, providerprop)
	template.CommodityBought = commbought
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: template,
		//	currentProvider: provider,
	}
	provider2 := new(Provider)
	commboughtProviderProp, isFound := scnbuilder.findCommBoughtProvider(provider2)
	assert.Equal(false, isFound)
	assert.Equal((*TemplateDTO_CommBoughtProviderProp)(nil), commboughtProviderProp)
}

// Tests that the Link method creates a new struct of type TemplateDTO_ExternalEntityLinkProp
// and this struct's Key variable is set to the SellerRef value inside the ExternalEntityLink passed to
// the Link method. This is  *extEntityLink.BuyerRef = *this.entityTemplate.TemplateClass true
func TestLink_setTrue_sameBuyer(t *testing.T) {
	assert := assert.New(t)
	class1 := EntityDTO_STORAGE
	entityTemp := &TemplateDTO{
		TemplateClass: &class1,
	}
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: entityTemp,
	}
	class2 := EntityDTO_SWITCH
	extEntityLink := &ExternalEntityLink{
		BuyerRef:  &class1,
		SellerRef: &class2,
	}
	scnb := scnbuilder.Link(extEntityLink)
	assert.Equal(*extEntityLink.SellerRef, *scnb.entityTemplate.ExternalLink[0].Key)
	assert.Equal(extEntityLink, scnb.entityTemplate.ExternalLink[0].Value)
}

// Tests that the Link method creates a new struct of type TemplateDTO_ExternalEntityLinkProp
// and this struct's Key variable is set to the BuyerRef value inside the ExternalEntityLink passed to
// the Link method. This is  *extEntityLink.SellerRef = *this.entityTemplate.TemplateClass true
func TestLink_setTrue_sameSeller(t *testing.T) {
	assert := assert.New(t)
	class1 := EntityDTO_STORAGE
	entityTemp := &TemplateDTO{
		TemplateClass: &class1,
	}
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: entityTemp,
	}
	class2 := EntityDTO_SWITCH
	extEntityLink := &ExternalEntityLink{
		BuyerRef:  &class2,
		SellerRef: &class1,
	}
	scnb := scnbuilder.Link(extEntityLink)
	assert.Equal(*extEntityLink.BuyerRef, *scnb.entityTemplate.ExternalLink[0].Key)
	assert.Equal(extEntityLink, scnb.entityTemplate.ExternalLink[0].Value)
}

// Test that the pointer to TemplateDTO_ExternalEntityLinkProp is appended to the array
// this.entityTemplate.ExternalLink
func TestAddExternalLinkPropToTemplateEntity(t *testing.T) {
	assert := assert.New(t)
	templateDTO := new(TemplateDTO)
	scnbuilder := &SupplyChainNodeBuilder{
		entityTemplate: templateDTO,
	}
	extEntityLinkP := new(TemplateDTO_ExternalEntityLinkProp)
	scnbuilder.addExternalLinkPropToTemplateEntity(extEntityLinkP)
	if assert.NotEqual(([]*TemplateDTO_ExternalEntityLinkProp)(nil), scnbuilder.entityTemplate.ExternalLink) {
		assert.Equal(extEntityLinkP, scnbuilder.entityTemplate.ExternalLink[0])
	}
}
