package builder

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

func TestNewCommodityDTOBuilder(t *testing.T) {
	commodityType := rand.RandomCommodityType()
	expectedBuilder := &CommodityDTOBuilder{
		commodityType: &commodityType,
	}
	builder := NewCommodityDTOBuilder(commodityType)
	if !reflect.DeepEqual(expectedBuilder, builder) {
		t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
	}
}

func TestCommodityDTOBuilder_Create(t *testing.T) {
	table := []struct {
		existingErr error
		expectError bool
	}{
		{
			existingErr: fmt.Errorf("Fake"),
			expectError: true,
		},
		{
			expectError: false,
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		if item.existingErr != nil {
			base.err = item.existingErr
		}
		commodityDTO, err := base.Create()
		if item.expectError {
			if err == nil {
				t.Errorf("Expected erro, got no error.")
			}
		} else {
			expectedCommodityDTO := &proto.CommodityDTO{
				CommodityType: base.commodityType,
			}
			if !reflect.DeepEqual(expectedCommodityDTO, commodityDTO) {
				t.Errorf("Expected %+v, got %+v", expectedCommodityDTO, commodityDTO)
			}
		}
	}
}

func TestCommodityDTOBuilder_Key(t *testing.T) {
	table := []struct {
		key         string
		existingErr error
	}{
		{
			key:         rand.String(5),
			existingErr: fmt.Errorf("Fake"),
		},
		{
			key: rand.String(5),
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		expectedBuilder := &CommodityDTOBuilder{commodityType: base.commodityType}
		if item.existingErr != nil {
			base.err = item.existingErr
			expectedBuilder.err = item.existingErr
		} else {
			var key *string
			if item.key != "" {
				key = &item.key
			}
			expectedBuilder.commodityType = base.commodityType
			expectedBuilder.key = key
		}
		builder := base.Key(item.key)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestCommodityDTOBuilder_Used(t *testing.T) {
	table := []struct {
		used        float64
		existingErr error
	}{
		{
			used:        mathrand.Float64(),
			existingErr: fmt.Errorf("Fake"),
		},
		{
			used: mathrand.Float64(),
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		var used *float64
		if item.existingErr != nil {
			base.err = item.existingErr
		} else {
			used = &item.used
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType: base.commodityType,
			used:          used,
			err:           item.existingErr,
		}
		builder := base.Used(item.used)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestCommodityDTOBuilder_Reservation(t *testing.T) {
	table := []struct {
		reservation        float64
		existingErr error
	}{
		{
			reservation:        mathrand.Float64(),
			existingErr: fmt.Errorf("Fake"),
		},
		{
			reservation: mathrand.Float64(),
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		var reservation *float64
		if item.existingErr != nil {
			base.err = item.existingErr
		} else {
			reservation = &item.reservation
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType: base.commodityType,
			reservation:          reservation,
			err:           item.existingErr,
		}
		builder := base.Reservation(item.reservation)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestCommodityDTOBuilder_Capacity(t *testing.T) {
	table := []struct {
		capacity    float64
		existingErr error
	}{
		{
			existingErr: fmt.Errorf("Fake"),
			capacity:    mathrand.Float64(),
		},
		{
			capacity: mathrand.Float64(),
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		var capacity *float64
		if item.existingErr != nil {
			base.err = item.existingErr
		} else {
			capacity = &item.capacity
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType: base.commodityType,
			capacity:      capacity,
			err:           item.existingErr,
		}
		builder := base.Capacity(item.capacity)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestCommodityDTOBuilder_Resizable(t *testing.T) {
	table := []struct {
		resizable   bool
		existingErr error
	}{
		{
			existingErr: fmt.Errorf("Fake"),
			resizable:   mathrand.Int31n(2) == 1,
		},
		{
			resizable: true,
		},
		{
			resizable: false,
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		var resizable *bool
		if item.existingErr != nil {
			base.err = item.existingErr
		} else {
			resizable = &item.resizable
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType: base.commodityType,
			resizable:     resizable,
			err:           item.existingErr,
		}
		builder := base.Resizable(item.resizable)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func randomBaseCommodityDTOBuilder() *CommodityDTOBuilder {
	cType := rand.RandomCommodityType()
	return &CommodityDTOBuilder{
		commodityType: &cType,
	}
}
