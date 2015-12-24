package sdk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vmturbo/vmturbo-go-sdk/util/rand"
	mathrand "math/rand"
	"testing"
)

func TestProviderDTOGetProviderID(t *testing.T) {
	assert := assert.New(t)
	fmt.Println("in TestProviderDTOProviderID")
	id := rand.String(5)

	providerDto := &ProviderDTO{
		Id: &id,
	}

	assert.Equal(&id, providerDto.getId())
}

// Tests that getProviderType() returns the correct pointer
func TestProviderDTOGetProviderType(t *testing.T) {
	assert := assert.New(t)
	fmt.Println("in TestProviderDTOGetProviderType")
	pType := new(EntityDTO_EntityType)

	providerDto := &ProviderDTO{
		providerType: pType,
	}

	assert.Equal(pType, providerDto.getProviderType())
}

//Tests the method NewEntityDTOBuilder() , which should return a pointer to a EntityDTOBuilder
//instance containing only its EntityDTOBuilder.entity member instantiated.
func Test_EntityDTOBuilder(t *testing.T) {
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
func TestEntityDTOBuilder_Create(t *testing.T) {
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
func TestEntityDTOBuilder_DisplayName(t *testing.T) {
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
func TestEntityDTOBuilder_Sells(t *testing.T) {
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
func TestEntityDTOBuilder_Used_True(t *testing.T) {
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
	assert.Equal(used, *entityDTOBuilder.entity.CommoditiesSold[0].Used)
}

// Tests the method Used(used float64) to not set the CommodityDTO in the
// this.entity.CommoditiesSold array with the used float64 variable passed as argument to Used.
func TestEntityDTOBuilder_Used_False(t *testing.T) {
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
func TestEntityDTOBuilder_Capacity_True(t *testing.T) {
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
func TestEntityDTOBuilder_Capacity_False(t *testing.T) {
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
// NewEntityDTOBuilder() constructor creates a DTO builder and only initializes the
// entity member object of the EntityDTOBuilder it returns.
func TestEntityDTOBuilder_requireCommodity_True(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(6)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	commDTO := new(CommodityDTO)
	commType := new(CommodityDTO_CommodityType)
	keystr := rand.String(6)

	commDTO.CommodityType = commType
	commDTO.Key = &keystr
	entityDTOBuilder.commodity = commDTO
	assert.Equal(true, entityDTOBuilder.requireCommodity())
}

// test to see if the EntityDTOBuilder calling object's member commodity is indeed null
// NewEntityDTOBuilder() constructor creates a DTO builder and only initializes the
// entity member object of the EntityDTOBuilder it returns.
func TestEntityDTOBuilder_requireCommodity_False(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(6)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	assert.Equal(true, entityDTOBuilder.requireCommodity())
}
