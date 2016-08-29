package builder

import "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"

type ExternalEntityLinkBuilder struct {
	entityLink *proto.ExternalEntityLink
}

func NewExternalEntityLinkBuilder() *ExternalEntityLinkBuilder {
	link := &proto.ExternalEntityLink{}
	return &ExternalEntityLinkBuilder{
		entityLink: link,
	}
}

// Initialize the buyer/seller external link that you're building.
// This method sets the entity types for the buyer and seller, as well as the type of provider
// relationship HOSTING or code LAYEREDOVER.
func (this *ExternalEntityLinkBuilder) Link(buyer, seller proto.EntityDTO_EntityType, relationship proto.Provider_ProviderType) *ExternalEntityLinkBuilder {
	this.entityLink.BuyerRef = &buyer
	this.entityLink.SellerRef = &seller
	this.entityLink.Relationship = &relationship

	return this
}

// Add a single bought commodity to the link.
func (this *ExternalEntityLinkBuilder) Commodity(comm proto.CommodityDTO_CommodityType, hasKey bool) *ExternalEntityLinkBuilder {
	commodityDefs := this.entityLink.GetCommodityDefs()
	commodityDef := &proto.ExternalEntityLink_CommodityDef{
		Type:   &comm,
		HasKey: &hasKey,
	}
	commodityDefs = append(commodityDefs, commodityDef)

	this.entityLink.CommodityDefs = commodityDefs
	return this
}

// Set a property of the discovered entity to the link. Operations Manager will use this property to
// stitch the discovered entity into the Operations Manager topology. This setting includes the property name
// and an arbitrary description.
func (this *ExternalEntityLinkBuilder) ProbeEntityPropertyDef(name, description string) *ExternalEntityLinkBuilder {
	entityProperty := &proto.ExternalEntityLink_EntityPropertyDef{
		Name:        &name,
		Description: &description,
	}
	currentProps := this.entityLink.GetProbeEntityPropertyDef()
	currentProps = append(currentProps, entityProperty)
	this.entityLink.ProbeEntityPropertyDef = currentProps

	return this
}

// Set an ServerEntityPropertyDef to the link you're building.
// The ServerEntityPropertyDef includes metadata for the properties of the
// external entity. Operations Manager can use the metadata
// to stitch entities discovered by the probe together with external entities.
// An external entity is one that exists in the Operations Manager topology, but has
// not been discovered by the probe.
func (this *ExternalEntityLinkBuilder) ExternalEntityPropertyDef(propertyDef *proto.ExternalEntityLink_ServerEntityPropDef) *ExternalEntityLinkBuilder {
	currentExtProps := this.entityLink.GetExternalEntityPropertyDefs()
	currentExtProps = append(currentExtProps, propertyDef)

	this.entityLink.ExternalEntityPropertyDefs = currentExtProps

	return this
}

// Get the ExternalEntityLink that you have built.
func (this *ExternalEntityLinkBuilder) Build() *proto.ExternalEntityLink {
	return this.entityLink
}
