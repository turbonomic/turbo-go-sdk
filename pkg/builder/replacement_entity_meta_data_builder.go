package builder

import "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"

type ReplacementEntityMetaDataBuilder struct {
	metaData *proto.EntityDTO_ReplacementEntityMetaData
}

func NewReplacementEntityMetaDataBuilder() *ReplacementEntityMetaDataBuilder {
	var identifyingProp []string
	var buyingCommTypes []proto.CommodityDTO_CommodityType
	var sellingCommTypes []proto.CommodityDTO_CommodityType
	replacementEntityMetaData := &proto.EntityDTO_ReplacementEntityMetaData{
		IdentifyingProp:  identifyingProp,
		BuyingCommTypes:  buyingCommTypes,
		SellingCommTypes: sellingCommTypes,
	}
	return &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMetaData,
	}
}

func (this *ReplacementEntityMetaDataBuilder) Build() *proto.EntityDTO_ReplacementEntityMetaData {
	return this.metaData
}

// Specifies the name of the property whose value will be used to find the server entity
// for which this entity is a proxy. The value for the property must be set while building the
// entity.
// Specific properties are pre-defined for some entity types. See the constants defined in
// supply_chain_constants for the names of the specific properties.
func (this *ReplacementEntityMetaDataBuilder) Matching(property string) *ReplacementEntityMetaDataBuilder {
	this.metaData.IdentifyingProp = append(this.metaData.GetIdentifyingProp(), property)
	return this
}

// Set the commodity type whose metric values will be transferred to the entity
//  this DTO will be replaced by.
func (this *ReplacementEntityMetaDataBuilder) PatchBuying(commType proto.CommodityDTO_CommodityType) *ReplacementEntityMetaDataBuilder {
	this.metaData.BuyingCommTypes = append(this.metaData.GetBuyingCommTypes(), commType)
	return this
}

// Set the commodity type whose metric values will be transferred to the entity
//  this DTO will be replaced by.
func (this *ReplacementEntityMetaDataBuilder) PatchSelling(commType proto.CommodityDTO_CommodityType) *ReplacementEntityMetaDataBuilder {
	this.metaData.SellingCommTypes = append(this.metaData.GetSellingCommTypes(), commType)
	return this
}
