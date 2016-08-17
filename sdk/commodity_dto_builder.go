package sdk

type CommodityDTOBuilder struct {
	commDTO *CommodityDTO
}

func NewCommodityDTOBuilder(commodityType CommodityDTO_CommodityType) *CommodityDTOBuilder {
	commodityDTO := new(CommodityDTO)
	commodityDTO.CommodityType = &commodityType
	return &CommodityDTOBuilder{
		commDTO: commodityDTO,
	}
}

func (this *CommodityDTOBuilder) Create() *CommodityDTO {
	return this.commDTO
}

func (this *CommodityDTOBuilder) Key(key string) *CommodityDTOBuilder {
	this.commDTO.Key = &key
	return this
}

func (this *CommodityDTOBuilder) Capacity(capcity float64) *CommodityDTOBuilder {
	this.commDTO.Capacity = &capcity
	return this
}

func (this *CommodityDTOBuilder) Used(used float64) *CommodityDTOBuilder {
	this.commDTO.Used = &used
	return this
}
