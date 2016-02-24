package sdk

import (
	"math"

	"github.com/golang/glog"
)

type SupplyChainNodeBuilder struct {
	entityTemplate  *TemplateDTO
	currentProvider *Provider
}

// Create a new SupplyChainNode Builder
func NewSupplyChainNodeBuilder() *SupplyChainNodeBuilder {
	return new(SupplyChainNodeBuilder)
}

// Create a SupplyChainNode
func (scnb *SupplyChainNodeBuilder) Create() *TemplateDTO {
	return scnb.entityTemplate
}

// Build the entity of the SupplyChainNode
func (scnb *SupplyChainNodeBuilder) Entity(entityType EntityDTO_EntityType) *SupplyChainNodeBuilder {
	var commSold []*TemplateCommodity
	var commBought []*TemplateDTO_CommBoughtProviderProp
	templateType := TemplateDTO_BASE
	priority := int32(0)
	scnb.entityTemplate = &TemplateDTO{
		TemplateClass:    &entityType,
		TemplateType:     &templateType,
		TemplatePriority: &priority,
		CommoditySold:    commSold,
		CommodityBought:  commBought,
	}
	return scnb
}

// Get the entityType of the TemplateDTO
func (scnb *SupplyChainNodeBuilder) getEntity() EntityDTO_EntityType {
	if hasEntityTemplate := scnb.requireEntityTemplate(); !hasEntityTemplate {
		//TODO!! should give error
		// return
	}
	return scnb.entityTemplate.GetTemplateClass()
}

// Check if the entityTemplate has been set.
func (scnb *SupplyChainNodeBuilder) requireEntityTemplate() bool {
	if scnb.entityTemplate == nil {
		return false
	}

	return true
}

// Check if the provider has been set.
func (scnb *SupplyChainNodeBuilder) requireProvider() bool {
	if scnb.currentProvider == nil {
		return false
	}
	return true
}

// The very basic selling method. If want others, use other names
func (scnb *SupplyChainNodeBuilder) Selling(comm CommodityDTO_CommodityType, key string) *SupplyChainNodeBuilder {
	if hasEntityTemplate := scnb.requireEntityTemplate(); !hasEntityTemplate {
		//TODO should give error
		return scnb
	}

	// Add commodity sold
	templateComm := &TemplateCommodity{
		Key:           &key,
		CommodityType: &comm,
	}
	commSold := scnb.entityTemplate.CommoditySold
	commSold = append(commSold, templateComm)
	scnb.entityTemplate.CommoditySold = commSold
	return scnb
}

// set the provider of the SupplyChainNode
func (scnb *SupplyChainNodeBuilder) Provider(provider EntityDTO_EntityType, pType Provider_ProviderType) *SupplyChainNodeBuilder {
	if hasTemplate := scnb.requireEntityTemplate(); !hasTemplate {
		//TODO should give error

		return scnb
	}

	if pType == Provider_LAYERED_OVER {
		maxCardinality := int32(math.MaxInt32)
		minCardinality := int32(0)
		scnb.currentProvider = &Provider{
			TemplateClass:  &provider,
			ProviderType:   &pType,
			CardinalityMax: &maxCardinality,
			CardinalityMin: &minCardinality,
		}
	} else {
		hostCardinality := int32(1)
		scnb.currentProvider = &Provider{
			TemplateClass:  &provider,
			ProviderType:   &pType,
			CardinalityMax: &hostCardinality,
			CardinalityMin: &hostCardinality,
		}
	}

	return scnb
}

// Add a commodity this node buys from the current provider. The provider must already been specified.
// If there is no provider for this node, does not add the commodity.
func (scnb *SupplyChainNodeBuilder) Buys(templateComm TemplateCommodity) *SupplyChainNodeBuilder {
	if hasEntityTemplate := scnb.requireEntityTemplate(); !hasEntityTemplate {
		//TODO should give error
		//glog.V(3).Infof("---------- Error! No entity found! ----------")
		return scnb
	}

	if hasProvider := scnb.requireProvider(); !hasProvider {
		//TODO should give error
		//glog.V(3).Infof("---------- Error! No provider found! ----------")
		return scnb
	}

	boughtMap := scnb.entityTemplate.GetCommodityBought()
	providerProp, exist := scnb.findCommBoughtProvider(scnb.currentProvider)
	if !exist {

		providerProp = new(TemplateDTO_CommBoughtProviderProp)
		providerProp.Key = scnb.currentProvider
		var value []*TemplateCommodity
		providerProp.Value = value

		boughtMap = append(boughtMap, providerProp)
		scnb.entityTemplate.CommodityBought = boughtMap

	}

	providerPropValue := providerProp.GetValue()
	providerPropValue = append(providerPropValue, &templateComm)
	providerProp.Value = providerPropValue

	return scnb
}

// Check if current provider exists in BoughtMap of the templateDTO.
// TODO, this should be a method in templateDTO?
func (scnb *SupplyChainNodeBuilder) findCommBoughtProvider(provider *Provider) (*TemplateDTO_CommBoughtProviderProp, bool) {
	for _, pp := range scnb.entityTemplate.GetCommodityBought() {

		if pp.GetKey() == provider {
			return pp, true
		}
	}
	return nil, false
}

//Adds an external entity link to the current node.
func (scnb *SupplyChainNodeBuilder) Link(extEntityLink *ExternalEntityLink) *SupplyChainNodeBuilder {
	if set := scnb.requireEntityTemplate(); !set {
		return scnb
	}

	linkProp := &TemplateDTO_ExternalEntityLinkProp{}
	if extEntityLink.GetBuyerRef() == scnb.entityTemplate.GetTemplateClass() {
		seller := extEntityLink.GetSellerRef()
		linkProp.Key = &seller
	} else if extEntityLink.GetSellerRef() == scnb.entityTemplate.GetTemplateClass() {
		buyer := extEntityLink.GetBuyerRef()
		linkProp.Key = &buyer
	} else {
		glog.Errorf("Template entity is not one of the entity in this external link")
		return scnb
	}
	linkProp.Value = extEntityLink
	scnb.addExternalLinkPropToTemplateEntity(linkProp)
	return scnb

}

func (scnb *SupplyChainNodeBuilder) addExternalLinkPropToTemplateEntity(extEntityLinkProp *TemplateDTO_ExternalEntityLinkProp) {
	currentLinks := scnb.entityTemplate.GetExternalLink()
	currentLinks = append(currentLinks, extEntityLinkProp)
	scnb.entityTemplate.ExternalLink = currentLinks
}
