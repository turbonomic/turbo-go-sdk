package builder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
		reservation float64
		existingErr error
	}{
		{
			reservation: mathrand.Float64(),
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
			reservation:   reservation,
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

func TestCommodityDTOBuilder_Active(t *testing.T) {
	table := []struct {
		active      bool
		existingErr error
	}{
		{
			existingErr: fmt.Errorf("Fake"),
			active:      true,
		},
		{
			active: true,
		},
		{
			active: false,
		},
	}
	cType := proto.CommodityDTO_CommodityType(106)
	for _, item := range table {
		base := &CommodityDTOBuilder{
			commodityType: &cType,
		}
		var active *bool
		if item.existingErr != nil {
			base.err = item.existingErr
		} else {
			active = &item.active
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType: base.commodityType,
			active:        active,
			err:           item.existingErr,
		}
		builder := base.Active(item.active)
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

func TestCommodityDTOBuilder_UtilizationData(t *testing.T) {
	testPoints := []float64{10.0, 11.0}
	testLastPointTimestampMs := int64(1585853340000)
	testIntervalMs := int32(0)
	testUtilizationData := &proto.CommodityDTO_UtilizationData{
		Point:                testPoints,
		LastPointTimestampMs: &testLastPointTimestampMs,
		IntervalMs:           &testIntervalMs,
	}
	table := []struct {
		utilizationData *proto.CommodityDTO_UtilizationData
		existingErr     error
	}{
		{
			existingErr:     fmt.Errorf("Fake"),
			utilizationData: testUtilizationData,
		},
		{
			utilizationData: testUtilizationData,
		},
	}
	cType := proto.CommodityDTO_CommodityType(106)
	for _, item := range table {
		base := &CommodityDTOBuilder{
			commodityType: &cType,
		}
		var utilizationData *proto.CommodityDTO_UtilizationData
		if item.existingErr != nil {
			base.err = item.existingErr
		} else {
			utilizationData = testUtilizationData
		}
		expectedBuilder := &CommodityDTOBuilder{
			commodityType:   &cType,
			utilizationData: utilizationData,
			err:             item.existingErr,
		}
		builder := base.UtilizationData(testPoints, testLastPointTimestampMs, testIntervalMs)
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("\nExpected %+v, \ngot      %+v", expectedBuilder, builder)
		}
	}
}

func TestCommodityDTOBuilder_RemainingGCCapacity(t *testing.T) {
	builder := NewCommodityDTOBuilder(proto.CommodityDTO_REMAINING_GC_CAPACITY).Used(10)
	commodity, err := builder.Create()
	assert.NoError(t, err)
	assert.NotNil(t, commodity)
	assert.Equal(t, commodity.GetCapacity(), 100.0)
	assert.Equal(t, commodity.GetUsed(), 90.0)
}

func TestCommodityDTOBuilder_NoCapacity(t *testing.T) {
	builder := NewCommodityDTOBuilder(proto.CommodityDTO_CONNECTION).Used(10)
	commodity, err := builder.Create()
	assert.Nil(t, commodity)
	assert.EqualError(t, err, "commodity CONNECTION has nil capacity")
}

func TestCommodityDTOBuilder_NegativeUsed(t *testing.T) {
	builder := randomBaseCommodityDTOBuilder().Used(-1).Capacity(200)
	commodity, err := builder.Create()
	assert.Nil(t, commodity)
	assert.EqualError(t, err, fmt.Errorf("commodity %v has negative used value", builder.commodityType).Error())
}

func randomBaseCommodityDTOBuilder() *CommodityDTOBuilder {
	cType := rand.RandomCommodityType()
	return &CommodityDTOBuilder{
		commodityType: &cType,
	}
}
