package builder

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"testing"

	"errors"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

func TestEntityDTOBuilder_NewEntityDTOBuilder(t *testing.T) {
	table := []struct {
		eType proto.EntityDTO_EntityType
		id    string
	}{
		{
			rand.RandomEntityType(),
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
			eType:      rand.RandomEntityType(),
			id:         rand.String(5),
			powerState: rand.RandomPowerState(),
			origin:     rand.RandomOrigin(),
			commoditiesBoughtProviderMap: map[string][]*proto.CommodityDTO{
				rand.String(5): []*proto.CommodityDTO{
					rand.RandomCommodityDTOBought(),
				},
			},
			commoditiesSold: []*proto.CommodityDTO{
				rand.RandomCommodityDTOSold(),
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
				rand.RandomCommodityDTOSold(),
				rand.RandomCommodityDTOSold(),
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
		err     error
	}{
		{
			commDTO: rand.RandomCommodityDTOSold(),
			err:     nil,
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

func TestEntityDTOBuilder_WithPowerState(t *testing.T) {
	table := []struct {
		powerState  proto.EntityDTO_PowerState
		existingErr error
	}{
		{
			powerState:  rand.RandomPowerState(),
			existingErr: fmt.Errorf("Error"),
		},
		{
			powerState: rand.RandomPowerState(),
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		expectedBuilder := &EntityDTOBuilder{
			entityType: base.entityType,
			id:         base.id,
		}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			expectedBuilder.powerState = &item.powerState
		}
		builder := base.WithPowerState(item.powerState)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
		}
	}
}

func TestEntityDTOBuilder_Monitored(t *testing.T) {
	table := []struct {
		monitored   bool
		existingErr error
	}{
		{
			monitored:   mathrand.Int31n(2) == 1,
			existingErr: fmt.Errorf("Error"),
		},
		{
			monitored: mathrand.Int31n(2) == 1,
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		expectedBuilder := &EntityDTOBuilder{
			entityType: base.entityType,
			id:         base.id,
		}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			expectedBuilder.monitored = &item.monitored
		}
		builder := base.Monitored(item.monitored)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
		}
	}
}

func TestEntityDTOBuilder_ApplicationData(t *testing.T) {
	table := []struct {
		appData *proto.EntityDTO_ApplicationData

		entityDataHasSetFlag bool
		existingErr          error
	}{
		{
			appData:     rand.RandomApplicationData(),
			existingErr: fmt.Errorf("Error"),
		},
		{
			appData:              rand.RandomApplicationData(),
			entityDataHasSetFlag: false,
		},
		{
			appData:              rand.RandomApplicationData(),
			entityDataHasSetFlag: true,
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		base.entityDataHasSet = item.entityDataHasSetFlag
		expectedBuilder := &EntityDTOBuilder{
			entityType:       base.entityType,
			id:               base.id,
			entityDataHasSet: base.entityDataHasSet,
		}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			if item.entityDataHasSetFlag {
				expectedBuilder.err = fmt.Errorf("EntityData has already been set. Cannot use %v as entity data.", item.appData)

			} else {
				expectedBuilder.applicationData = item.appData
				expectedBuilder.entityDataHasSet = true
			}
		}
		builder := base.ApplicationData(item.appData)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestEntityDTOBuilder_VirtualMachineData(t *testing.T) {
	table := []struct {
		vmData *proto.EntityDTO_VirtualMachineData

		entityDataHasSetFlag bool
		existingErr          error
	}{
		{
			vmData:      rand.RandomVirtualMachineData(),
			existingErr: fmt.Errorf("Error"),
		},
		{
			vmData:               rand.RandomVirtualMachineData(),
			entityDataHasSetFlag: false,
		},
		{
			vmData:               rand.RandomVirtualMachineData(),
			entityDataHasSetFlag: true,
		},
	}
	for _, item := range table {
		base := randomBaseEntityDTOBuilder()
		base.entityDataHasSet = item.entityDataHasSetFlag
		expectedBuilder := &EntityDTOBuilder{
			entityType:       base.entityType,
			id:               base.id,
			entityDataHasSet: base.entityDataHasSet,
		}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			if item.entityDataHasSetFlag {
				expectedBuilder.err = fmt.Errorf("EntityData has already been set. Cannot use %v as entity data.", item.vmData)

			} else {
				expectedBuilder.virtualMachineData = item.vmData
				expectedBuilder.entityDataHasSet = true
			}
		}
		builder := base.VirtualMachineData(item.vmData)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestEntityDTOBuilder_VirtualApplicationData(t *testing.T) {
	table := []struct {
		vAppData *proto.EntityDTO_VirtualApplicationData

		entityDataHasSetFlag bool
		existingErr          error
	}{
		{
			vAppData:    rand.RandomVirtualApplicationData(),
			existingErr: errors.New("Error"),
		},
		{
			vAppData:             rand.RandomVirtualApplicationData(),
			entityDataHasSetFlag: false,
		},
		{
			vAppData:             rand.RandomVirtualApplicationData(),
			entityDataHasSetFlag: true,
		},
	}
	for i, item := range table {
		base := randomBaseEntityDTOBuilder()
		base.entityDataHasSet = item.entityDataHasSetFlag
		expectedBuilder := &EntityDTOBuilder{
			entityType:       base.entityType,
			id:               base.id,
			entityDataHasSet: base.entityDataHasSet,
		}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			if item.entityDataHasSetFlag {
				expectedBuilder.err = fmt.Errorf("EntityData has already been set. Cannot use %v as entity data.", item.vAppData)

			} else {
				expectedBuilder.virtualApplicationData = item.vAppData
				expectedBuilder.entityDataHasSet = true
			}
		}
		builder := base.VirtualApplicationData(item.vAppData)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("Test case %d failed. Expected %+v, \ngot      %+v", i, expectedBuilder, builder)
		}
	}
}

func TestBuildCommodityBoughtFromMap(t *testing.T) {
	table := []struct {
		providerCount int

		provider1  string
		commodity1 []*proto.CommodityDTO

		provider2  string
		commodity2 []*proto.CommodityDTO
	}{
		{
			providerCount: 0,
		},
		{
			providerCount: 1,

			provider1: rand.String(5),
			commodity1: []*proto.CommodityDTO{
				rand.RandomCommodityDTOSold(),
				rand.RandomCommodityDTOSold(),
			},
		},
		{
			providerCount: 2,

			provider1: rand.String(5),
			commodity1: []*proto.CommodityDTO{
				rand.RandomCommodityDTOSold(),
			},

			provider2: rand.String(5),
			commodity2: []*proto.CommodityDTO{
				rand.RandomCommodityDTOSold(),
			},
		},
	}
	for i, item := range table {
		inputMap := make(map[string][]*proto.CommodityDTO)
		if item.providerCount > 0 {
			if item.provider1 != "" {
				inputMap[item.provider1] = item.commodity1
			}
			if item.provider2 != "" {
				inputMap[item.provider2] = item.commodity2
			}
		}

		expectedCommoditiesBought := make(map[string]*proto.EntityDTO_CommodityBought)
		if item.providerCount > 0 {
			if item.provider1 != "" {
				expectedCommoditiesBought[item.provider1] =
					&proto.EntityDTO_CommodityBought{
						ProviderId: &item.provider1,
						Bought:     item.commodity1,
					}
			}
			if item.provider2 != "" {
				expectedCommoditiesBought[item.provider2] =
					&proto.EntityDTO_CommodityBought{
						ProviderId: &item.provider2,
						Bought:     item.commodity2,
					}
			}
		}

		gotCommoditiesBought := buildCommodityBoughtFromMap(inputMap)
		for _, commBought := range gotCommoditiesBought {
			found := false
			if expectedComm, exists := expectedCommoditiesBought[commBought.GetProviderId()]; exists {
				if !reflect.DeepEqual(expectedComm, commBought) {
					t.Errorf("Test case %d failed. Expected %++v, got %++v", i,
						expectedComm, commBought)
					continue
				}
				found = true
				delete(expectedCommoditiesBought, commBought.GetProviderId())
			}
			if !found {
				t.Errorf("Test case %d failed. Unexpected commodity bought %++v", i, commBought)
			}
		}
		if len(expectedCommoditiesBought) != 0 {
			t.Errorf("Test case %d failed. Expected commodities bought %++v NOT found.", i,
				expectedCommoditiesBought)
		}
	}
}

// Create a random EntityDTOBuilder.
func randomBaseEntityDTOBuilder() *EntityDTOBuilder {
	return NewEntityDTOBuilder(rand.RandomEntityType(), rand.String(5))
}
