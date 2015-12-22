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

func TestProviderDTOGetProviderType(t *testing.T) {
	assert := assert.New(t)
	fmt.Println("in TestProviderDTOGetProviderType")
	// TODO. Not sure if this is a good way to generate an EntityType. Or we hardcode one particular type here.
	pType := new(EntityDTO_EntityType)

	providerDto := &ProviderDTO{
		providerType: pType,
	}

	assert.Equal(pType, providerDto.getProviderType())
}

/*
 * Tests the method NewEntityDTOBuilder() , which should return a pointer to a EntityDTOBuilder 
 * instance containing only its EntityDTOBuilder.entity member instantiated.
 */
func Test_EntityDTOBuilder(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(5)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	if assert.NotNil(t, entityDTOBuilder.entity){
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

// check with DongYi is this method supposed to be get entity?
func TestEntityDTOBuilder_Create(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(5)
	entity := new(EntityDTO)
	entity.EntityType = pType
	entity.Id = &idstr
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}
	assert.Equal(*entityDTOBuilder.entity, *entityDTOBuilder.Create())
	assert.Equal(entityDTOBuilder.entity, entityDTOBuilder.Create())
}

func TestEntityDTOBuilder_DisplayName(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(5)
	entity := new(EntityDTO)
	entity.EntityType = pType
	entity.Id = &idstr
	entityDTOBuilder := &EntityDTOBuilder{
		entity: entity,
	}

	dispName := rand.String(6)
	entityDTOBuilder.DisplayName(dispName)
	assert.Equal(&dispName, entityDTOBuilder.entity.DisplayName)
	assert.Equal(dispName, *entityDTOBuilder.entity.DisplayName)
}

func TestEntityDTOBuilder_Sells(t *testing.T) {
	assert := assert.New(t)
	commType := new(CommodityDTO_CommodityType)
	keystr := rand.String(6)

	pType := new(EntityDTO_EntityType)
	idstr := rand.String(5)
	entity := new(EntityDTO)
	entity.EntityType = pType
	entity.Id = &idstr
	var commoditiesSold []*CommodityDTO
	entity.CommoditiesSold = commoditiesSold
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
	//	assert.Equal(t, entityDTOBuilder, entityDTOBuilder_ptr)
}

func TestEntityDTOBuilder_Used_True(t *testing.T) {
	r := mathrand.New(mathrand.NewSource(99))
	used := r.Float64()
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(6)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	if assert.NotNil(t, entityDTOBuilder.entity.CommoditiesSold) {
		assert.Equal(0, len(entityDTOBuilder.entity.CommoditiesSold))
	}
	commType := new(CommodityDTO_CommodityType)
	keystr := rand.String(6)
	entityDTOBuilder.Sells(*commType, keystr)
	assert.Equal(1, len(entityDTOBuilder.entity.CommoditiesSold))
	assert.Equal(commType, entityDTOBuilder.entity.CommoditiesSold[0].CommodityType)
	assert.Equal(*commType, *entityDTOBuilder.entity.CommoditiesSold[0].CommodityType)
	assert.Equal(keystr, *entityDTOBuilder.entity.CommoditiesSold[0].Key)
	assert.Equal(&keystr, entityDTOBuilder.entity.CommoditiesSold[0].Key)
	entityDTOBuilder.Used(used)
	assert.Equal(used, *entityDTOBuilder.entity.CommoditiesSold[0].Used)
}

func TestEntityDTOBuilder_Used_False(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(6)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	r := mathrand.New(mathrand.NewSource(99))
	used := r.Float64()
	assert.Equal(0, len(entityDTOBuilder.Used(used).entity.CommoditiesSold))
}

func TestEntityDTOBuilder_Capacity_True(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(6)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
	r := mathrand.New(mathrand.NewSource(99))
	cap := r.Float64()
	commType := new(CommodityDTO_CommodityType)
	keystr := rand.String(6)
	entityDTOBuilder.Sells(*commType, keystr)
	entityDTOBuilder.Capacity(cap)
	assert.Equal(cap, *entityDTOBuilder.entity.CommoditiesSold[0].Capacity)
}

func TestEntityDTOBuilder_Capacity_False(t *testing.T) {
	assert := assert.New(t)
	pType := new(EntityDTO_EntityType)
	idstr := rand.String(6)
	entityDTOBuilder := NewEntityDTOBuilder(*pType, idstr)
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
