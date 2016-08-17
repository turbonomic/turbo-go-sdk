package sdk

type EntityDTOBuilder struct {
	entity          *EntityDTO
	commodity       *CommodityDTO
	currentProvider *ProviderDTO
}

type ProviderDTO struct {
	providerType *EntityDTO_EntityType
	Id           *string
}

func CreateProvider(pType EntityDTO_EntityType, id string) *ProviderDTO {
	return &ProviderDTO{
		providerType: &pType,
		Id:           &id,
	}
}

func (pDto *ProviderDTO) getProviderType() *EntityDTO_EntityType {
	return pDto.providerType
}

func (pDto *ProviderDTO) getId() *string {
	return pDto.Id
}

func NewEntityDTOBuilder(eType EntityDTO_EntityType, id string) *EntityDTOBuilder {
	entity := new(EntityDTO)
	entity.EntityType = &eType
	entity.Id = &id
	var commoditiesBought []*EntityDTO_CommodityBought
	var commoditiesSold []*CommodityDTO
	entity.CommoditiesBought = commoditiesBought
	entity.CommoditiesSold = commoditiesSold
	return &EntityDTOBuilder{
		entity: entity,
	}
}

func (eb *EntityDTOBuilder) Create() *EntityDTO {
	return eb.entity
}

func (eb *EntityDTOBuilder) DisplayName(disp string) *EntityDTOBuilder {
	eb.entity.DisplayName = &disp

	return eb
}

func (eb *EntityDTOBuilder) SellsCommodities(commDTOs []*CommodityDTO) {
	commSold := eb.entity.CommoditiesSold
	commSold = append(commSold, commDTOs...)
	eb.entity.CommoditiesSold = commSold
}

func (eb *EntityDTOBuilder) Sells(commodityType CommodityDTO_CommodityType, key string) *EntityDTOBuilder {
	commDTO := new(CommodityDTO)
	commDTO.CommodityType = &commodityType
	commDTO.Key = &key

	commSold := eb.entity.CommoditiesSold
	commSold = append(commSold, commDTO)
	eb.entity.CommoditiesSold = commSold
	eb.commodity = commDTO
	return eb
}

func (eb *EntityDTOBuilder) Used(used float64) *EntityDTOBuilder {
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

func (eb *EntityDTOBuilder) SetProvider(provider *ProviderDTO) {
	eb.currentProvider = provider
}

func (eb *EntityDTOBuilder) SetProviderWithTypeAndID(pType EntityDTO_EntityType, id string) *EntityDTOBuilder {
	eb.currentProvider = &ProviderDTO{
		providerType: &pType,
		Id:           &id,
	}
	return eb
}

// Add an commodity which buys from the current provider.
func (eb *EntityDTOBuilder) Buys(commodityType CommodityDTO_CommodityType, key string, used float64) *EntityDTOBuilder {
	if eb.currentProvider == nil {
		// TODO should have error message. Notify set current provider first
		return eb
	}
	// defaultCapacity := float64(0.0)
	commDTO := new(CommodityDTO)
	commDTO.CommodityType = &commodityType
	commDTO.Key = &key
	// commDTO.Capacity = &defaultCapacity
	commDTO.Used = &used
	eb.BuysCommodity(commDTO)
	return eb
}

func (eb *EntityDTOBuilder) BuysCommodities(commDTOs []*CommodityDTO) {
	for _, commDTO := range commDTOs {
		eb.BuysCommodity(commDTO)
	}
}

func (eb *EntityDTOBuilder) BuysCommodity(commDTO *CommodityDTO) *EntityDTOBuilder {
	if eb.currentProvider == nil {
		// TODO should have error message. Notify set current provider first
		return eb
	}

	// add commodity bought
	commBought := eb.entity.GetCommoditiesBought()
	providerId := eb.currentProvider.Id
	commBoughtFromProvider, find := eb.findCommBoughtProvider(providerId)
	if !find {
		// Create an EntityDTO_CommodityBought with empty CommodityDTO list
		commBoughtFromProvider = new(EntityDTO_CommodityBought)
		var commBoughtList []*CommodityDTO
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
func (eb *EntityDTOBuilder) findCommBoughtProvider(providerId *string) (*EntityDTO_CommodityBought, bool) {
	for _, commBougthProvider := range eb.entity.CommoditiesBought {
		if commBougthProvider.GetProviderId() == *providerId {
			return commBougthProvider, true
		}
	}
	return nil, false
}

// Add property to an entity
func (eb *EntityDTOBuilder) SetProperty(name, value string) *EntityDTOBuilder {
	if eb.entity.GetEntityProperties() == nil {
		var entityProps []*EntityDTO_EntityProperty
		eb.entity.EntityProperties = entityProps
	}
	prop := &EntityDTO_EntityProperty{
		Name:  &name,
		Value: &value,
	}
	eb.entity.EntityProperties = append(eb.entity.EntityProperties, prop)
	return eb
}

// Set the ReplacementEntityMetadata that will contain the information about the external entity
// that this entity will patch with the metrics data it collected.
func (eb *EntityDTOBuilder) ReplacedBy(replacementEntityMetaData *EntityDTO_ReplacementEntityMetaData) *EntityDTOBuilder {
	origin := EntityDTO_PROXY
	eb.entity.Origin = &origin
	eb.entity.ReplacementEntityData = replacementEntityMetaData
	return eb
}
