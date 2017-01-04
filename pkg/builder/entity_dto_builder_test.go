package builder

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

func TestEntityDTOBuilder_NewEntityDTOBuilder(t *testing.T) {
	table := []struct {
		eType proto.EntityDTO_EntityType
		id    string
	}{
		{
			randomEntityType(),
			rand.String(5),
		},
	}
	for _, item := range table {
		builder := NewEntityDTOBuilder(item.eType, item.id)
		expectedBuilder := &EntityDTOBuilder{
			entityType: &item.eType,
			id:         &item.id,
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("Expect builder %++v, got %++v", expectedBuilder, builder)
		}
	}
}

// Tests the method Create() , which returns the entity member of the EntityDTOBuilder that
// called this method.
func TestCreate(t *testing.T) {
	table := []struct {
		eType                        proto.EntityDTO_EntityType
		id                           string
		powerState                   proto.EntityDTO_PowerState
		origin                       proto.EntityDTO_EntityOrigin
		commoditiesBoughtProviderMap map[string][]*proto.CommodityDTO
		commoditiesSold              []*proto.CommodityDTO
		err                          error

		expectsError bool
	}{
		{
			eType:      randomEntityType(),
			id:         rand.String(5),
			powerState: randomPowerState(),
			origin:     randomOrigin(),
			commoditiesBoughtProviderMap: map[string][]*proto.CommodityDTO{
				rand.String(5): []*proto.CommodityDTO{
					randomCommodityDTOBought(),
				},
			},
			commoditiesSold: []*proto.CommodityDTO{
				randomCommodityDTOSold(),
			},
			expectsError: false,
		},
		{
			err:          fmt.Errorf("Fake Error"),
			expectsError: true,
		},
	}
	for _, item := range table {
		builder := &EntityDTOBuilder{
			entityType: &item.eType,
			id:         &item.id,
			powerState: &item.powerState,
			origin:     &item.origin,
			commoditiesBoughtProviderMap: item.commoditiesBoughtProviderMap,
			commoditiesSold:              item.commoditiesSold,
			err:                          item.err,
		}
		entityDTO, err := builder.Create()

		if gotError := err != nil; item.expectsError != gotError {
			t.Errorf("Expect error? %t, but got hasError? %t", item.expectsError, gotError)
		}
		if !item.expectsError {
			expectedEntityDTO := &proto.EntityDTO{
				EntityType:        &item.eType,
				Id:                &item.id,
				PowerState:        &item.powerState,
				Origin:            &item.origin,
				CommoditiesSold:   item.commoditiesSold,
				CommoditiesBought: buildCommodityBoughtFromMap(item.commoditiesBoughtProviderMap),
			}
			if !reflect.DeepEqual(expectedEntityDTO, entityDTO) {
				t.Errorf("\nExpect\t %++v, \ngot\t %++v", expectedEntityDTO, entityDTO)
			}
		}
	}
}

// Tests method DisplayName() which sets the DisplayName of the entity member of the
// EntityDTOBuilder that calls DisplayName()
func TestDisplayName(t *testing.T) {
	table := []struct {
		displayName string
		err         error
	}{
		{
			displayName: rand.String(10),
			err:         nil,
		},
		{
			err: fmt.Errorf("Fake error"),
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		if item.err != nil {
			base.err = item.err
		}
		builder := base.DisplayName(item.displayName)

		var displayName *string
		if item.displayName != "" {
			displayName = &item.displayName
		}
		expectedBuilder := &EntityDTOBuilder{
			entityType:  base.entityType,
			id:          base.id,
			displayName: displayName,
			err:         item.err,
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected: %++v, \ngot %++v", expectedBuilder, builder)
		}
	}
}

func TestEntityDTOBuilder_SellsCommodities(t *testing.T) {
	table := []struct {
		commDTOs []*proto.CommodityDTO
		err      error
	}{
		{
			commDTOs: []*proto.CommodityDTO{
				randomCommodityDTOSold(),
				randomCommodityDTOSold(),
			},
			err: nil,
		},
		{
			err: fmt.Errorf("Fake error"),
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		if item.err != nil {
			base.err = item.err
		}
		builder := base.SellsCommodities(item.commDTOs)
		expectedBuilder := &EntityDTOBuilder{
			entityType:      base.entityType,
			id:              base.id,
			commoditiesSold: item.commDTOs,
			err:             item.err,
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected: %++v, \ngot %++v", expectedBuilder, builder)
		}
	}
}

func TestEntityDTOBuilder_SellsCommodity(t *testing.T) {
	table := []struct {
		commDTO *proto.CommodityDTO
		err      error
	}{
		{
			commDTO: randomCommodityDTOSold(),
			err: nil,
		},
		{
			err: fmt.Errorf("Fake error"),
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		if item.err != nil {
			base.err = item.err
		}
		builder := base.SellsCommodity(item.commDTO)
		var comms []*proto.CommodityDTO
		if item.commDTO != nil {
			comms = append(comms, item.commDTO)
		}
		expectedBuilder := &EntityDTOBuilder{
			entityType:      base.entityType,
			id:              base.id,
			commoditiesSold: comms,
			err:             item.err,
		}
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected:\n %++v, \ngot\n %++v", expectedBuilder, builder)
		}
	}
}

// Create a random entity type.
func randomEntityType() proto.EntityDTO_EntityType {
	return proto.EntityDTO_EntityType(mathrand.Int31n(42))
}

// Create a random commodity type.
func randomCommodityType() proto.CommodityDTO_CommodityType {
	return proto.CommodityDTO_CommodityType(mathrand.Int31n(77))
}

// Create a random power state value, range from 1 to 4.
func randomPowerState() proto.EntityDTO_PowerState {
	return proto.EntityDTO_PowerState(mathrand.Int31n(4) + 1)
}

// Create a random entity origin, range from 1 to 2.
func randomOrigin() proto.EntityDTO_EntityOrigin {
	return proto.EntityDTO_EntityOrigin(mathrand.Int31n(2) + 1)
}

// Create a random commodityDTO bought.
func randomCommodityDTOBought() *proto.CommodityDTO {
	// a random commodity type.
	cType := randomCommodityType()
	// a random key
	key := rand.String(5)
	// a random used
	used := mathrand.Float64()
	return &proto.CommodityDTO{
		CommodityType: &cType,
		Key:           &key,
		Used:          &used,
	}

}

// Create a random CommodityDTO sold.
func randomCommodityDTOSold() *proto.CommodityDTO {
	// a random commodity type.
	cType := randomCommodityType()
	// a random key
	key := rand.String(5)
	// a random capacity
	capacity := mathrand.Float64()
	// a random used
	used := mathrand.Float64()
	return &proto.CommodityDTO{
		CommodityType: &cType,
		Key:           &key,
		Capacity:      &capacity,
		Used:          &used,
	}

}

// Create a random EntityDTOBuilder.
func randomBaseEntityDTOBuilder() *EntityDTOBuilder {
	return NewEntityDTOBuilder(randomEntityType(), rand.String(5))
}
