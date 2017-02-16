package probe

import (
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
	"reflect"
	"github.com/turbonomic/turbo-api/pkg/api"
)

func TestNewTurboTargetInfoBuilder(t *testing.T) {
	table := []struct {
		category              string
		targetType            string
		targetIdentifierField string
		accountValues         []*proto.AccountValue
	}{
		{
			category:              rand.String(5),
			targetType:            rand.String(5),
			targetIdentifierField: rand.String(5),
			accountValues:         []*proto.AccountValue{rand.RandomAccoutValue()},
		},
	}

	for _, item := range table {
		expectedBuilder := &TurboTargetInfoBuilder{
			targetCategory:        item.category,
			targetType:            item.targetType,
			targetIdentifierField: item.targetIdentifierField,
			accountValues:         item.accountValues,
		}
		builder := NewTurboTargetInfoBuilder(item.category, item.targetType,
			item.targetIdentifierField, item.accountValues)
		if !reflect.DeepEqual(expectedBuilder, builder) {
			t.Errorf("Expected %v, got %v", expectedBuilder, builder)
		}
	}
}

func TestTurboTargetInfoBuilder_Create(t *testing.T) {
	table := []struct {
		category              string
		targetType            string
		targetIdentifierField string
		accountValues         []*proto.AccountValue
	}{
		{
			category:              rand.String(5),
			targetType:            rand.String(5),
			targetIdentifierField: rand.String(5),
			accountValues:         []*proto.AccountValue{rand.RandomAccoutValue()},
		},
	}
	for _, item := range table {
		expectedInfo := &TurboTargetInfo{
			targetCategory:        item.category,
			targetType:            item.targetType,
			targetIdentifierField: item.targetIdentifierField,
			accountValues:         item.accountValues,
		}
		info := NewTurboTargetInfoBuilder(item.category, item.targetType,
			item.targetIdentifierField, item.accountValues).Create()
		if !reflect.DeepEqual(expectedInfo, info) {
			t.Errorf("Expected %v, got %v", expectedInfo, info)
		}
	}
}

func TestTurboTargetInfo_TargetIdentifierField(t *testing.T) {
	field := rand.String(5)
	info := &TurboTargetInfo{
		targetIdentifierField: field,
	}
	if !reflect.DeepEqual(field, info.TargetIdentifierField()) {
		t.Errorf("Expect %s, got %s", field, info.TargetIdentifierField())
	}
}

func TestTurboTargetInfo_TargetType(t *testing.T) {
	targetType := rand.String(5)
	info := &TurboTargetInfo{
		targetType: targetType,
	}
	if !reflect.DeepEqual(targetType, info.TargetType()) {
		t.Errorf("Expect %s, got %s", targetType, info.TargetType())
	}
}

func TestTurboTargetInfo_TargetCategory(t *testing.T) {
	targetCategory := rand.String(5)
	info := &TurboTargetInfo{
		targetCategory: targetCategory,
	}
	if !reflect.DeepEqual(targetCategory, info.TargetCategory()) {
		t.Errorf("Expect %s, got %s", targetCategory, info.TargetCategory())
	}
}

func TestTurboTargetInfo_GetTargetInstance(t *testing.T) {
	table := []struct {
		category   string
		targetType string
		idField    string

		acctValueKey         string
		acctValueStringValue string

		expectedTarget *api.Target
	}{
		{
			category:             rand.String(5),
			targetType:           rand.String(5),
			idField:              rand.String(5),
			acctValueKey:         rand.String(5),
			acctValueStringValue: rand.String(5),
		},
	}
	for _, item := range table {
		info := &TurboTargetInfo{
			targetCategory:        item.category,
			targetType:            item.targetType,
			targetIdentifierField: item.idField,
			accountValues:         []*proto.AccountValue{
				{
					Key: &item.acctValueKey,
					StringValue: &item.acctValueStringValue,
				},
			},
		}
		expectedTarget := &api.Target{
			Category: item.category,
			Type:     item.targetType,
			InputFields: []*api.InputField{
				{
					Name:            item.acctValueKey,
					Value:           item.acctValueStringValue,
					GroupProperties: []*api.List{},
				},
			},
		}
		target := info.GetTargetInstance()
		if !reflect.DeepEqual(expectedTarget, target) {
			t.Errorf("Expected %v, got %v", expectedTarget, target)
		}
	}
}
