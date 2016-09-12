package builder

import (
	"fmt"

	"github.com/vmturbo/vmturbo-go-sdk/pkg/common"
	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

type EntityDTOBuilder struct {
	entity          *proto.EntityDTO
	commodity       *proto.CommodityDTO
	currentProvider *common.ProviderDTO

	err error
}

func NewEntityDTOBuilder(eType proto.EntityDTO_EntityType, id string) *EntityDTOBuilder {
	entity := new(proto.EntityDTO)
	entity.EntityType = &eType
	entity.Id = &id
	var commoditiesBought []*proto.EntityDTO_CommodityBought
	var commoditiesSold []*proto.CommodityDTO
	entity.CommoditiesBought = commoditiesBought
	entity.CommoditiesSold = commoditiesSold
	return &EntityDTOBuilder{
		entity: entity,
	}
}

func (eb *EntityDTOBuilder) Create() (*proto.EntityDTO, error) {
	if eb.err != nil {
		return nil, eb.err
	}
	return eb.entity, nil
}

func (eb *EntityDTOBuilder) DisplayName(disp string) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	eb.entity.DisplayName = &disp

	return eb
}

func (eb *EntityDTOBuilder) SellsCommodities(commDTOs []*proto.CommodityDTO) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	commSold := eb.entity.CommoditiesSold
	commSold = append(commSold, commDTOs...)
	eb.entity.CommoditiesSold = commSold

	return eb
}

func (eb *EntityDTOBuilder) Sells(commodityType proto.CommodityDTO_CommodityType, key string) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	commDTO := new(proto.CommodityDTO)
	commDTO.CommodityType = &commodityType
	commDTO.Key = &key

	commSold := eb.entity.CommoditiesSold
	commSold = append(commSold, commDTO)
	eb.entity.CommoditiesSold = commSold
	eb.commodity = commDTO
	return eb
}

func (eb *EntityDTOBuilder) Used(used float64) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	hasCommodity := eb.requireCommodity()
	if hasCommodity {
		// TODO use map might be better
		for _, commDto := range eb.entity.CommoditiesSold {
			if eb.commodity.GetCommodityType() == commDto.GetCommodityType() {
				commDto.Used = &used
			}
		}
	}
	return eb
}

func (eb *EntityDTOBuilder) Capacity(capacity float64) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	hasCommodity := eb.requireCommodity()
	if hasCommodity {
		for _, commDto := range eb.entity.CommoditiesSold {
			if eb.commodity.GetCommodityType() == commDto.GetCommodityType() {
				commDto.Capacity = &capacity
			}
		}
	}
	return eb
}

func (eb *EntityDTOBuilder) requireCommodity() bool {
	if eb.commodity == nil {
		return false
	}
	return true
}

func (eb *EntityDTOBuilder) SetProvider(provider *common.ProviderDTO) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	eb.currentProvider = provider
	return eb
}

func (eb *EntityDTOBuilder) SetProviderWithTypeAndID(pType proto.EntityDTO_EntityType, id string) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	eb.currentProvider = common.CreateProvider(pType, id)
	return eb
}

// Add an commodity which buys from the current provider.
func (eb *EntityDTOBuilder) Buys(commodityType proto.CommodityDTO_CommodityType, key string, used float64) *EntityDTOBuilder {
	if eb.currentProvider == nil {
		// TODO should have error message. Notify set current provider first
		eb.err = fmt.Errorf("Provider has not been set for %v", commodityType)
		return eb
	}
	commDTO := new(proto.CommodityDTO)
	commDTO.CommodityType = &commodityType
	commDTO.Key = &key
	commDTO.Used = &used
	eb.BuysCommodity(commDTO)
	return eb
}

func (eb *EntityDTOBuilder) BuysCommodities(commDTOs []*proto.CommodityDTO) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	for _, commDTO := range commDTOs {
		eb.BuysCommodity(commDTO)
	}
	return eb
}

func (eb *EntityDTOBuilder) BuysCommodity(commDTO *proto.CommodityDTO) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	if eb.currentProvider == nil {
		// TODO should have error message. Notify set current provider first
		eb.err = fmt.Errorf("Porvider has not been set for %v", commDTO)
		return eb
	}

	// add commodity bought
	commBought := eb.entity.GetCommoditiesBought()
	providerId := eb.currentProvider.Id
	commBoughtFromProvider, find := eb.findCommBoughtProvider(providerId)
	if !find {
		// Create an EntityDTO_CommodityBought with empty CommodityDTO list
		commBoughtFromProvider = new(proto.EntityDTO_CommodityBought)
		var commBoughtList []*proto.CommodityDTO
		commBoughtFromProvider.Bought = commBoughtList
		commBoughtFromProvider.ProviderId = providerId

		//add to commBoughtFromProvider the CommoditiesBought list of entity
		commBought = append(commBought, commBoughtFromProvider)

		eb.entity.CommoditiesBought = commBought
	}
	commodities := commBoughtFromProvider.GetBought()
	commodities = append(commodities, commDTO)
	commBoughtFromProvider.Bought = commodities

	return eb
}

// Find if this the current provider has already been in the map.
// If found, return the commodityDTO list bought from the provider.
// TODO this should belongs to entityDTO
func (eb *EntityDTOBuilder) findCommBoughtProvider(providerId *string) (*proto.EntityDTO_CommodityBought, bool) {
	for _, commBougthProvider := range eb.entity.CommoditiesBought {
		if commBougthProvider.GetProviderId() == *providerId {
			return commBougthProvider, true
		}
	}
	return nil, false
}

// Add property to an entity
func (eb *EntityDTOBuilder) SetProperty(name, value string) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	if eb.entity.GetEntityProperties() == nil {
		var entityProps []*proto.EntityDTO_EntityProperty
		eb.entity.EntityProperties = entityProps
	}
	prop := &proto.EntityDTO_EntityProperty{
		Name:  &name,
		Value: &value,
	}
	eb.entity.EntityProperties = append(eb.entity.EntityProperties, prop)
	return eb
}

// Set the ReplacementEntityMetadata that will contain the information about the external entity
// that this entity will patch with the metrics data it collected.
func (eb *EntityDTOBuilder) ReplacedBy(replacementEntityMetaData *proto.EntityDTO_ReplacementEntityMetaData) *EntityDTOBuilder {
	if eb.err != nil {
		return eb
	}
	origin := proto.EntityDTO_PROXY
	eb.entity.Origin = &origin
	eb.entity.ReplacementEntityData = replacementEntityMetaData
	return eb
}
