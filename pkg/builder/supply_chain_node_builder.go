package builder

import (
	"fmt"
	"math"

	"github.com/golang/glog"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

type SupplyChainNodeBuilder struct {
	entityTemplate  *proto.TemplateDTO
	currentProvider *proto.Provider

	err error
}

// Create a new SupplyChainNode Builder
func NewSupplyChainNodeBuilder() *SupplyChainNodeBuilder {
	return new(SupplyChainNodeBuilder)
}

// Create a SupplyChainNode
func (scnb *SupplyChainNodeBuilder) Create() (*proto.TemplateDTO, error) {
	if scnb.err != nil {
		return nil, fmt.Errorf("Cannot create supply chain node because of error: %v", scnb.err)
	}
	return scnb.entityTemplate, nil
}

// Build the entity of the SupplyChainNode
func (scnb *SupplyChainNodeBuilder) Entity(entityType proto.EntityDTO_EntityType) *SupplyChainNodeBuilder {
	var commSold []*proto.TemplateCommodity
	var commBought []*proto.TemplateDTO_CommBoughtProviderProp
	templateType := proto.TemplateDTO_BASE
	priority := int32(0)
	scnb.entityTemplate = &proto.TemplateDTO{
		TemplateClass:    &entityType,
		TemplateType:     &templateType,
		TemplatePriority: &priority,
		CommoditySold:    commSold,
		CommodityBought:  commBought,
	}
	return scnb
}

// The very basic selling method. If want others, use other names
func (scnb *SupplyChainNodeBuilder) Sells(templateComm proto.TemplateCommodity) *SupplyChainNodeBuilder {
	if hasEntityTemplate := scnb.requireEntityTemplate(); !hasEntityTemplate {
		//TODO should give error
		return scnb
	}

	commSold := scnb.entityTemplate.CommoditySold
	commSold = append(commSold, &templateComm)
	scnb.entityTemplate.CommoditySold = commSold
	return scnb
}

// set the provider of the SupplyChainNode
func (scnb *SupplyChainNodeBuilder) Provider(provider proto.EntityDTO_EntityType, pType proto.Provider_ProviderType) *SupplyChainNodeBuilder {
	if scnb.err != nil {
		return scnb
	}
	if hasTemplate := scnb.requireEntityTemplate(); !hasTemplate {
		scnb.err = fmt.Errorf("EntityTemplate is not found. Must set before call Provider().")
		return scnb
	}

	if pType == proto.Provider_LAYERED_OVER {
		maxCardinality := int32(math.MaxInt32)
		minCardinality := int32(0)
		scnb.currentProvider = &proto.Provider{
			TemplateClass:  &provider,
			ProviderType:   &pType,
			CardinalityMax: &maxCardinality,
			CardinalityMin: &minCardinality,
		}
	} else {
		hostCardinality := int32(1)
		scnb.currentProvider = &proto.Provider{
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
func (scnb *SupplyChainNodeBuilder) Buys(templateComm proto.TemplateCommodity) *SupplyChainNodeBuilder {
	if scnb.err != nil {
		return scnb
	}
	if hasEntityTemplate := scnb.requireEntityTemplate(); !hasEntityTemplate {
		scnb.err = fmt.Errorf("EntityTemplate is not found. Must set before call Buys().")
		return scnb
	}

	if hasProvider := scnb.requireProvider(); !hasProvider {
		scnb.err = fmt.Errorf("Provider is not found. Must set before call Buys().")
		return scnb
	}

	boughtMap := scnb.entityTemplate.GetCommodityBought()
	// 1. Check if the current provider is already added to the CommodityBought map of current templateDTO.
	providerProp, exist := findProviderInCommBoughtMap(boughtMap, scnb.currentProvider)
	// 2. If not exist, put current provider into CommodityBought map.
	if !exist {

		providerProp = new(proto.TemplateDTO_CommBoughtProviderProp)
		providerProp.Key = scnb.currentProvider
		var value []*proto.TemplateCommodity
		providerProp.Value = value

		boughtMap = append(boughtMap, providerProp)
		scnb.entityTemplate.CommodityBought = boughtMap
	}
	// 3. Add current commodity into commodityBought map.
	providerPropValue := providerProp.GetValue()
	providerPropValue = append(providerPropValue, &templateComm)
	providerProp.Value = providerPropValue

	return scnb
}

// Check if current provider exists in CommodityBoughtProvider map of the templateDTO.
func findProviderInCommBoughtMap(commBoughtProviders []*proto.TemplateDTO_CommBoughtProviderProp,
	provider *proto.Provider) (*proto.TemplateDTO_CommBoughtProviderProp, bool) {
	for _, pp := range commBoughtProviders {
		if pp.GetKey() == provider {
			return pp, true
		}
	}
	return nil, false
}

//Adds an external entity link to the current node.
func (scnb *SupplyChainNodeBuilder) Link(extEntityLink *proto.ExternalEntityLink) *SupplyChainNodeBuilder {
	if set := scnb.requireEntityTemplate(); !set {
		return scnb
	}

	linkProp := &proto.TemplateDTO_ExternalEntityLinkProp{}
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

func (scnb *SupplyChainNodeBuilder) addExternalLinkPropToTemplateEntity(extEntityLinkProp *proto.TemplateDTO_ExternalEntityLinkProp) {
	currentLinks := scnb.entityTemplate.GetExternalLink()
	currentLinks = append(currentLinks, extEntityLinkProp)
	scnb.entityTemplate.ExternalLink = currentLinks
}

// Get the entityType of the TemplateDTO
func (scnb *SupplyChainNodeBuilder) getType() (*proto.EntityDTO_EntityType, error) {
	if hasEntityTemplate := scnb.requireEntityTemplate(); !hasEntityTemplate {
		return nil, fmt.Errorf("EntityTemplate has no been set. Call Entity() first.")
	}
	entityType := scnb.entityTemplate.GetTemplateClass()
	return &entityType, nil
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
