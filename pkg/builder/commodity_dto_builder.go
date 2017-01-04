package builder

import "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"

type CommodityDTOBuilder struct {
	commodityType      *proto.CommodityDTO_CommodityType
	key                *string
	used               *float64
	reservation        *float64
	capacity           *float64
	limit              *float64
	peak               *float64
	active             *bool
	resizable          *bool
	displayName        *string
	thin               *bool
	computedUsed       *bool
	usedIncrement      *float64
	propMap            map[string][]string
	storageLatencyData *proto.CommodityDTO_StorageLatencyData
	storageAccessData  *proto.CommodityDTO_StorageAccessData

	err error
}

func NewCommodityDTOBuilder(commodityType proto.CommodityDTO_CommodityType) *CommodityDTOBuilder {
	return &CommodityDTOBuilder{
		commodityType: &commodityType,
	}
}

func (cb *CommodityDTOBuilder) Create() (*proto.CommodityDTO, error) {
	if cb.err != nil {
		return nil, cb.err
	}
	return &proto.CommodityDTO{
		CommodityType:      cb.commodityType,
		Key:                cb.key,
		Used:               cb.used,
		Reservation:        cb.reservation,
		Capacity:           cb.capacity,
		Limit:              cb.limit,
		Peak:               cb.peak,
		Active:             cb.active,
		Resizable:          cb.active,
		DisplayName:        cb.displayName,
		Thin:               cb.thin,
		ComputedUsed:       cb.computedUsed,
		UsedIncrement:      cb.usedIncrement,
		PropMap:            buildPropertyMap(cb.propMap),
		StorageLatencyData: cb.storageLatencyData,
		StorageAccessData:  cb.storageAccessData,
	}, nil
}

func (cb *CommodityDTOBuilder) Key(key string) *CommodityDTOBuilder {
	if cb.err != nil {
		return cb
	}
	cb.key = &key
	return cb
}

func (cb *CommodityDTOBuilder) Capacity(capacity float64) *CommodityDTOBuilder {
	if cb.err != nil {
		return cb
	}
	cb.capacity = &capacity
	return cb
}

func (cb *CommodityDTOBuilder) Used(used float64) *CommodityDTOBuilder {
	if cb.err != nil {
		return cb
	}
	cb.used = &used
	return cb
}

func buildPropertyMap(propMap map[string][]string) []*proto.CommodityDTO_PropertiesList {
	if propMap == nil {
		return nil
	}
	propList := []*proto.CommodityDTO_PropertiesList{}
	for name, values := range propMap {
		propList = append(propList, &proto.CommodityDTO_PropertiesList{
			Name:   &name,
			Values: values,
		})
	}
	return propList
}
