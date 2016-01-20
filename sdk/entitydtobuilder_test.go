package sdk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	mathrand "math/rand"
	"testing"
)

// Tests that getProviderType() returns the correct pointer to the
// providerType member variable that the getProviderType method is called on
func TestGetProviderType(t *testing.T) {
	assert := assert.New(t)
	fmt.Println("in TestProviderDTOGetProviderType")
	pType := new(EntityDTO_EntityType)

	providerDto := &ProviderDTO{
		providerType: pType,
	}

	assert.Equal(pType, providerDto.getProviderType())
}

// Tests that the getId() method returns a string pointer to
// the Id member variable of the ProviderDTO struct that getId() is called on
func TestGetId(t *testing.T) {
	assert := assert.New(t)
	fmt.Println("in TestProviderDTOProviderID")
	id := rand.String(5)
	providerDto := &ProviderDTO{
		Id: &id,
	}
	assert.Equal(&id, providerDto.getId())
}

//Tests the method NewEntityDTOBuilder() , which should return a pointer to a EntityDTOBuilder
//instance containing only its EntityDTOBuilder.entity member instantiated.
func TestNewEntityDTOBuilder(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(5)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	if assert.NotNil(t, entityDTOBuilder.entity) {
		assert.Equal(pType, entityDTOBuilder.entity.EntityType)
		assert.Equal(&idstr, entityDTOBuilder.entity.Id)
		if assert.NotNil(t, entityDTOBuilder.entity.CommoditiesBought) {
			assert.Equal(0, len(entityDTOBuilder.entity.CommoditiesBought))
		}
		if assert.NotNil(t, entityDTOBuilder.entity.CommoditiesSold) {
			assert.Equal(0, len(entityDTOBuilder.entity.CommoditiesSold))
		}
	}
}

// Tests the method Create() , which returns the entity member of the EntityDTOBuilder that
// called this method.
func TestCreate(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	assert.Equal(*entity, *entityDTOBuilder.Create())
	assert.Equal(entity, entityDTOBuilder.Create())
}

// Tests method DisplayName() which sets the DisplayName of the entity member of the
// EntityDTOBuilder that calls DisplayName()
func TestDisplayName(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}

	dispName := rand.String(6)
	entityDTOBuilder.DisplayName(dispName)
	assert.Equal(&dispName, entityDTOBuilder.entity.DisplayName)
	assert.Equal(dispName, *entityDTOBuilder.entity.DisplayName)
}

// Tests Sells() method which sets the CommodityType and key members of a new CommodityDTO instance
// and appends the new CommodityDTO instance to the CommoditiesSold member array of the entity memb// er of the EntityDTOBuilder that calls this method.
func TestSells(t *testing.T) {
	assert := assert.New(t)
	commType := new(CommodityDTO_CommodityType)
	keystr := rand.String(6)

	entity := new(EntityDTO)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}

	if assert.NotNil(t, entityDTOBuilder.entity.CommoditiesSold) {
		assert.Equal(0, len(entityDTOBuilder.entity.CommoditiesSold))
	}
	entityDTOBuilder.Sells(*commType, keystr)
	assert.Equal(1, len(entityDTOBuilder.entity.CommoditiesSold))
	assert.Equal(commType, entityDTOBuilder.entity.CommoditiesSold[0].CommodityType)
	assert.Equal(*commType, *entityDTOBuilder.entity.CommoditiesSold[0].CommodityType)
	assert.Equal(keystr, *entityDTOBuilder.entity.CommoditiesSold[0].Key)
	assert.Equal(&keystr, entityDTOBuilder.entity.CommoditiesSold[0].Key)
}

// Tests the method Used(used float64) to set the CommodityDTO in the
// this.entity.CommoditiesSold array with the used float64 variable passed as argument to Used.
// Tests case: hasCommodity == true
func TestUsed_True(t *testing.T) {
	r := mathrand.New(mathrand.NewSource(99))
	used := r.Float64()
	assert := assert.New(t)
	entity := new(EntityDTO)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	if assert.NotNil(t, entityDTOBuilder.entity.CommoditiesSold) {
		assert.Equal(0, len(entityDTOBuilder.entity.CommoditiesSold))
	}
	commType := new(CommodityDTO_CommodityType)
	commDTO := new(CommodityDTO)
	commDTO.CommodityType = commType
	commSold := append(entity.CommoditiesSold, commDTO)
	entity.CommoditiesSold = commSold
	entityDTOBuilder.commodity = commDTO
	assert.Equal(1, len(entityDTOBuilder.entity.CommoditiesSold))
	assert.Equal(commType, entityDTOBuilder.entity.CommoditiesSold[0].CommodityType)
	assert.Equal(*commType, *entityDTOBuilder.entity.CommoditiesSold[0].CommodityType)
	entityDTOBuilder.Used(used)
	assert.Equal(&used, entityDTOBuilder.entity.CommoditiesSold[0].Used)
	assert.Equal(used, *entityDTOBuilder.entity.CommoditiesSold[0].Used)
}

// Tests the method Used(used float64) to not set the CommodityDTO in the
// this.entity.CommoditiesSold array with the used float64 variable passed as argument to Used.
func TestUsed_False(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	// creates an EntityDTOBuilder with commodity = null
	// so that used is not set
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	r := mathrand.New(mathrand.NewSource(99))
	used := r.Float64()
	assert.Equal(0, len(entityDTOBuilder.Used(used).entity.CommoditiesSold))
}

// Tests that the correct capacity is set when hasCommodity = true
func TestCapacity_True(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	commDTO := new(CommodityDTO)
	commType := new(CommodityDTO_CommodityType)
	commDTO.CommodityType = commType
	entityDTOBuilder := &EntityDTOBuilder{
		entity:    entity,
		commodity: commDTO,
	}
	r := mathrand.New(mathrand.NewSource(99))
	cap := r.Float64()
	commSold := append(entityDTOBuilder.entity.CommoditiesSold, commDTO)
	entityDTOBuilder.entity.CommoditiesSold = commSold
	entityDTOBuilder.Capacity(cap)
	assert.Equal(cap, *entityDTOBuilder.entity.CommoditiesSold[0].Capacity)
}

// Tests that the method Capacity does not set commDTO.Capacity when hasCommodity = false
func TestCapacity_False(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	r := mathrand.New(mathrand.NewSource(99))
	cap := r.Float64()
	assert.Equal(0, len(entityDTOBuilder.Capacity(cap).entity.CommoditiesSold))
}

// test to see if the EntityDTOBuilder calling object's member commodity is indeed not null
// Creates an EntityDTOBuilder and only initializes the
// commodity member object of the EntityDTOBuilder.
func TestRequireCommodity_True(t *testing.T) {
	assert := assert.New(t)
	commDTO := new(CommodityDTO)
	entityDTOBuilder := &EntityDTOBuilder{
		commodity: commDTO,
	}
	assert.Equal(true, entityDTOBuilder.requireCommodity())
	assert.Equal(commDTO, entityDTOBuilder.commodity)
}

// test to see if the EntityDTOBuilder calling object's member commodity is indeed null
// Creates a EntityDTOBuilder and only initializes the
// commodity member object of the EntityDTOBuilder.
func TestRequireCommodity_False(t *testing.T) {
	assert := assert.New(t)
	entityDTOBuilder := &EntityDTOBuilder{}
	assert.Equal(false, entityDTOBuilder.requireCommodity())
}

// test that the SetProvider method creates a ProviderDTO and sets its providerType and id to the
// passed arguments
func TestSetProvider(t *testing.T) {
	assert := assert.New(t)
	entityDTOBuilder := &EntityDTOBuilder{}
	pType := new(EntityDTO_EntityType)
	id := rand.String(6)
	eb := entityDTOBuilder.SetProvider(*pType, id)
	assert.Equal(pType, eb.currentProvider.providerType)
	assert.Equal(*pType, *eb.currentProvider.providerType)
	assert.Equal(&id, eb.currentProvider.Id)
	assert.Equal(id, *eb.currentProvider.Id)
}

// Tests that a CommodityDTO is created with the arguments commodityType, key and used arguments
// as its members . Tests that the created CommodityDTO is added to the Bought array in the
// existing EntityDTO_CommodityBought struct with *ProviderId = * eb.currentProvider.Id
// When find = true
func TestBuys_True(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	providerIdStr := rand.String(5)
	entityDTO_CommodityBought := new(EntityDTO_CommodityBought)
	entityDTO_CommodityBought.ProviderId = &providerIdStr
	// Bought array is len = 0 for this EntityDTO_CommodityBought
	entity.CommoditiesBought = append(entity.CommoditiesBought, entityDTO_CommodityBought)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	providerDTO := new(ProviderDTO)
	providerDTO.Id = &providerIdStr
	entityDTOBuilder.currentProvider = providerDTO
	commType := new(CommodityDTO_CommodityType)
	key := rand.String(6)
	r := mathrand.New(mathrand.NewSource(99))
	used := r.Float64()
	eb := entityDTOBuilder.Buys(*commType, key, used)
	// eb.GetCommoditiesBought  => eb.CommoditiesBought  type []*EntityDTO_CommodityBought
	assert.Equal(commType, eb.entity.CommoditiesBought[0].Bought[0].CommodityType)
	assert.Equal(*commType, *eb.entity.CommoditiesBought[0].Bought[0].CommodityType)
	assert.Equal(&key, eb.entity.CommoditiesBought[0].Bought[0].Key)
	assert.Equal(key, *eb.entity.CommoditiesBought[0].Bought[0].Key)
	assert.Equal(&used, eb.entity.CommoditiesBought[0].Bought[0].Used)
	assert.Equal(used, *eb.entity.CommoditiesBought[0].Bought[0].Used)
}

// Tests that a CommodityDTO is created with the arguments commodityType, key and used arguments
// as its members . Tests that the created CommodityDTO is added to the Bought array in the
// existing EntityDTO_CommodityBought struct with *ProviderId = * eb.currentProvider.Id
// When find = false
func TestBuys_False(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	providerIdStr := rand.String(5)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	providerDTO := new(ProviderDTO)
	providerDTO.Id = &providerIdStr
	entityDTOBuilder.currentProvider = providerDTO
	commType := new(CommodityDTO_CommodityType)
	key := rand.String(6)
	r := mathrand.New(mathrand.NewSource(99))
	used := r.Float64()
	assert.Equal(0, len(entityDTOBuilder.entity.CommoditiesBought))
	eb := entityDTOBuilder.Buys(*commType, key, used)
	// eb.GetCommoditiesBought  => eb.CommoditiesBought  type []*EntityDTO_CommodityBought
	assert.Equal(1, len(eb.entity.CommoditiesBought))
	assert.Equal(commType, eb.entity.CommoditiesBought[0].Bought[0].CommodityType)
	assert.Equal(*commType, *eb.entity.CommoditiesBought[0].Bought[0].CommodityType)
	assert.Equal(&key, eb.entity.CommoditiesBought[0].Bought[0].Key)
	assert.Equal(key, *eb.entity.CommoditiesBought[0].Bought[0].Key)
	assert.Equal(&used, eb.entity.CommoditiesBought[0].Bought[0].Used)
	assert.Equal(used, *eb.entity.CommoditiesBought[0].Bought[0].Used)
}

// Test to assert that the correct *EntityDTO_CommodityBought is returned if the
// providerId for this struct is the same as the providerId passed to the method findCommBoughtProvider
func TestFindCommBoughtProvider_True_inrange(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	providerIdStr := rand.String(5)
	entityDTO_CommodityBought := new(EntityDTO_CommodityBought)
	entityDTO_CommodityBought.ProviderId = &providerIdStr
	// Bought array is len = 0 for this EntityDTO_CommodityBought
	entity.CommoditiesBought = append(entity.CommoditiesBought, entityDTO_CommodityBought)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	commBoughtProvider, wasFound := entityDTOBuilder.findCommBoughtProvider(&providerIdStr)
	assert.Equal(true, wasFound)
	assert.Equal(entityDTO_CommodityBought, commBoughtProvider)
	assert.Equal(*entityDTO_CommodityBought, *commBoughtProvider)
}

// Test to assert that the correct *EntityDTO_CommodityBought is returned if the
// providerId for this struct is the same as the providerId passed to the method findCommBoughtProvider
func TestFindCommBoughtProvider_False_inrange(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	providerIdStr := rand.String(5)
	currentProviderIdStr := rand.String(6)
	entityDTO_CommodityBought := new(EntityDTO_CommodityBought)
	entityDTO_CommodityBought.ProviderId = &providerIdStr
	// Bought array is len = 0 for this EntityDTO_CommodityBought
	entity.CommoditiesBought = append(entity.CommoditiesBought, entityDTO_CommodityBought)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	commBoughtProvider, wasFound := entityDTOBuilder.findCommBoughtProvider(&currentProviderIdStr)
	assert.Equal(false, wasFound)
	assert.Equal((*EntityDTO_CommodityBought)(nil), commBoughtProvider)
}

// Tests that the name and value passed as arguments to SetProperty are
// appended to the array EntityProperties of eb.entity
// Test case when eb.entity.GetEntityProperties != nil
func TestSetProperty_notnil(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	name := rand.String(5)
	value := rand.String(6)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	eb := entityDTOBuilder.SetProperty(name, value)
	assert.Equal(&name, eb.entity.EntityProperties[0].Name)
	assert.Equal(name, *eb.entity.EntityProperties[0].Name)
	assert.Equal(&value, eb.entity.EntityProperties[0].Value)
	assert.Equal(value, *eb.entity.EntityProperties[0].Value)
}

// Tests that the name and value passed as arguments to SetProperty are
// appended to the array EntityProperties of eb.entity
// Test case when eb.entity.GetEntityProperties = nil
func TestSetProperty_nil(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	name := rand.String(5)
	value := rand.String(6)
	entity.EntityProperties = nil
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	assert.Equal([]*EntityDTO_EntityProperty(nil), entityDTOBuilder.entity.EntityProperties)
	//assert.Equal(([]*EntityDTO_EntityProperty)(nil), entityDTOBuilder.entity.EntityProperties)
	//	assert.Nil(t, entityDTOBuilder.entity.EntityProperties)
	eb := entityDTOBuilder.SetProperty(name, value)
	assert.Equal(&name, eb.entity.EntityProperties[0].Name)
	assert.Equal(name, *eb.entity.EntityProperties[0].Name)
	assert.Equal(&value, eb.entity.EntityProperties[0].Value)
	assert.Equal(value, *eb.entity.EntityProperties[0].Value)
}

// Tests that the *EntityDTO_ReplacementEntityMetaData passed to ReplacedBy is used to set
// this.entity.ReplacementEntityData
func TestReplacedBy(t *testing.T) {
	assert := assert.New(t)
	entity := new(EntityDTO)
	replacementEntityMetaData := new(EntityDTO_ReplacementEntityMetaData)
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	assert.Equal((*EntityDTO_EntityOrigin)(nil), entityDTOBuilder.entity.Origin)
	eb := entityDTOBuilder.ReplacedBy(replacementEntityMetaData)
	assert.Equal(replacementEntityMetaData, eb.entity.ReplacementEntityData)
	assert.Equal(*replacementEntityMetaData, *eb.entity.ReplacementEntityData)

}
