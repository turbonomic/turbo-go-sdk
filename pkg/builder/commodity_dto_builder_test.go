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
	commodityType := randomCommodityType()
	expectedBuilder := &CommodityDTOBuilder{
		commodityType: &commodityType,
	}
	builder := NewCommodityDTOBuilder(commodityType)
	if !reflect.DeepEqual(expectedBuilder, builder) {
		t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
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
			existingErr: fmt.Errorf("Fake"),
		},
		{
			key: rand.String(5),
		},
	}

	for _, item := range table {
		base := randomBaseCommodityDTOBuilder()
		if item.existingErr != nil {
			base.err = item.existingErr
		}
		var key *string
		if item.key != "" {
			key = &item.key
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType: base.commodityType,
			key:           key,
			err:           item.existingErr,
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

func TestCommodityDTOBuilder_Capacity(t *testing.T) {
	table := []struct {
		capacity        float64
		existingErr error
	}{
		{
			existingErr: fmt.Errorf("Fake"),
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
			capacity:          capacity,
			err:           item.existingErr,
		}
		builder := base.Capacity(item.capacity)
		if !reflect.DeepEqual(builder, expectedBuilder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func randomBaseCommodityDTOBuilder() *CommodityDTOBuilder {
	cType := randomCommodityType()
	return &CommodityDTOBuilder{
		commodityType: &cType,
	}
}
