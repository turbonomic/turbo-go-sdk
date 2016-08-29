package builder

import "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"

type CommodityDTOBuilder struct {
	commDTO *proto.CommodityDTO
}

func NewCommodityDTOBuilder(commodityType proto.CommodityDTO_CommodityType) *CommodityDTOBuilder {
	commodityDTO := new(proto.CommodityDTO)
	commodityDTO.CommodityType = &commodityType
	return &CommodityDTOBuilder{
		commDTO: commodityDTO,
	}
}

func (this *CommodityDTOBuilder) Create() *proto.CommodityDTO {
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
