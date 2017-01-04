package builder

import "github.com/turbonomic/turbo-go-sdk/pkg/proto"

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
// relationship HOSTING or code LAYERED_OVER.
func (builder *ExternalEntityLinkBuilder) Link(buyer, seller proto.EntityDTO_EntityType,
	relationship proto.Provider_ProviderType) *ExternalEntityLinkBuilder {

	builder.entityLink.BuyerRef = &buyer
	builder.entityLink.SellerRef = &seller
	builder.entityLink.Relationship = &relationship

	return builder
}

// Add a single bought commodity to the link.
func (builder *ExternalEntityLinkBuilder) Commodity(comm proto.CommodityDTO_CommodityType, hasKey bool) *ExternalEntityLinkBuilder {
	commodityDefs := builder.entityLink.GetCommodityDefs()
	commodityDef := &proto.ExternalEntityLink_CommodityDef{
		Type:   &comm,
		HasKey: &hasKey,
	}
	commodityDefs = append(commodityDefs, commodityDef)

	builder.entityLink.CommodityDefs = commodityDefs
	return builder
}

// Set a property of the discovered entity to the link. Operations Manager will use builder property to
// stitch the discovered entity into the Operations Manager topology. This setting includes the property name
// and an arbitrary description.
func (builder *ExternalEntityLinkBuilder) ProbeEntityPropertyDef(name, description string) *ExternalEntityLinkBuilder {
	entityProperty := &proto.ExternalEntityLink_EntityPropertyDef{
		Name:        &name,
		Description: &description,
	}
	currentProps := builder.entityLink.GetProbeEntityPropertyDef()
	currentProps = append(currentProps, entityProperty)
	builder.entityLink.ProbeEntityPropertyDef = currentProps

	return builder
}

// Set an ServerEntityPropertyDef to the link you're building.
// The ServerEntityPropertyDef includes metadata for the properties of the  external entity. Operations Manager can
// use the metadata to stitch entities discovered by the probe together with external entities.
// An external entity is one that exists in the Operations Manager topology, but has not been discovered by the probe.
func (builder *ExternalEntityLinkBuilder) ExternalEntityPropertyDef(propertyDef *proto.ExternalEntityLink_ServerEntityPropDef) *ExternalEntityLinkBuilder {
	currentExtProps := builder.entityLink.GetExternalEntityPropertyDefs()
	currentExtProps = append(currentExtProps, propertyDef)

	builder.entityLink.ExternalEntityPropertyDefs = currentExtProps

	return builder
}

// Get the ExternalEntityLink that you have built.
func (builder *ExternalEntityLinkBuilder) Build() *proto.ExternalEntityLink {
	return builder.entityLink
}
