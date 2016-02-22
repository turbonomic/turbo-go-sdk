package sdk

type CommodtiyDTOBuilder struct {
	commDTO *CommodityDTO
}

func NewCommodtiyDTOBuilder(commodityType CommodityDTO_CommodityType) *CommodtiyDTOBuilder {
	commodityDTO := new(CommodityDTO)
	commodityDTO.CommodityType = &commodityType
	return &CommodtiyDTOBuilder{
		commDTO: commodityDTO,
	}
}

func (this *CommodtiyDTOBuilder) Create() *CommodityDTO {
	return this.commDTO
}

func (this *CommodtiyDTOBuilder) Key(key string) *CommodtiyDTOBuilder {
	this.commDTO.Key = &key
	return this
}

func (this *CommodtiyDTOBuilder) Capacity(capcity float64) *CommodtiyDTOBuilder {
	this.commDTO.Capacity = &capcity
	return this
}

func (this *CommodtiyDTOBuilder) Used(used float64) *CommodtiyDTOBuilder {
	this.commDTO.Used = &used
	return this
}
