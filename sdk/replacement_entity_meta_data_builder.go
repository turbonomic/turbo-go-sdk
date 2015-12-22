package sdk

type ReplacementEntityMetaDataBuilder struct {
	metaData *EntityDTO_ReplacementEntityMetaData
}

func NewReplacementEntityMetaDataBuilder() *ReplacementEntityMetaDataBuilder {
	var identifyingProp []string
	var buyingCommTypes []CommodityDTO_CommodityType
	var sellingCommTypes []CommodityDTO_CommodityType
	replacementEntityMetaData := &EntityDTO_ReplacementEntityMetaData{
		IdentifyingProp:  identifyingProp,
		BuyingCommTypes:  buyingCommTypes,
		SellingCommTypes: sellingCommTypes,
	}
	return &ReplacementEntityMetaDataBuilder{
		metaData: replacementEntityMetaData,
	}
}

func (this *ReplacementEntityMetaDataBuilder) Build() *EntityDTO_ReplacementEntityMetaData {
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
func (this *ReplacementEntityMetaDataBuilder) PatchBuying(commType CommodityDTO_CommodityType) *ReplacementEntityMetaDataBuilder {
	this.metaData.BuyingCommTypes = append(this.metaData.GetBuyingCommTypes(), commType)
	return this
}

// Set the commodity type whose metric values will be transferred to the entity
//  this DTO will be replaced by.
func (this *ReplacementEntityMetaDataBuilder) PatchSelling(commType CommodityDTO_CommodityType) *ReplacementEntityMetaDataBuilder {
	this.metaData.SellingCommTypes = append(this.metaData.GetSellingCommTypes(), commType)
	return this
}
