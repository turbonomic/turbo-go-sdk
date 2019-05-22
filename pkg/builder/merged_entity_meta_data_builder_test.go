package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestCreateMatchingData(t *testing.T) {
	empty := ""
	propertyName := "prop1"
	fieldName := "field1"
	fieldPaths := []string{}
	delimiter := ","

	matching := createMatchingData(propertyName, empty, fieldPaths, delimiter)
	assert.Equal(t, matching.delimiter, delimiter)
	assert.Equal(t, matching.propertyName, propertyName)

	matching = createMatchingData(propertyName, empty, fieldPaths, empty)
	assert.Equal(t, matching.delimiter, empty)
	assert.Equal(t, matching.propertyName, propertyName)

	matching2 := createMatchingData(empty, fieldName, fieldPaths, delimiter)
	assert.Equal(t, matching2.delimiter, delimiter)
	assert.Equal(t, matching2.propertyName, "")
}

// matchingData tests
func TestNewMatchingData(t *testing.T) {
	emptyStr := ""
	emptyList := []string{}
	propertyName := "prop1"
	fieldName := "field1"
	fieldPaths := []string{}
	delimiter := ","

	table := []struct {
		md       *matchingData
		expected *matchingData
	}{
		{
			md: &matchingData{},
			expected: &matchingData{propertyName: emptyStr, delimiter: emptyStr,
				fieldName: emptyStr, fieldPaths: emptyList,
				useEntityOid: false},
		},
		{
			md: &matchingData{propertyName: propertyName},
			expected: &matchingData{propertyName: propertyName, delimiter: emptyStr,
				fieldName: emptyStr, fieldPaths: emptyList,
				useEntityOid: false},
		},
		{
			md: &matchingData{propertyName: propertyName, delimiter: delimiter},
			expected: &matchingData{propertyName: propertyName, delimiter: delimiter,
				fieldName: emptyStr, fieldPaths: emptyList,
				useEntityOid: false},
		},
		{
			md: &matchingData{fieldName: fieldName, fieldPaths: fieldPaths},
			expected: &matchingData{propertyName: emptyStr, delimiter: emptyStr,
				fieldName: fieldName, fieldPaths: fieldPaths,
				useEntityOid: false},
		},
		{
			md: &matchingData{fieldName: fieldName, fieldPaths: fieldPaths, delimiter: delimiter},
			expected: &matchingData{propertyName: emptyStr, delimiter: delimiter,
				fieldName: fieldName, fieldPaths: fieldPaths,
				useEntityOid: false},
		},
		{
			md: &matchingData{useEntityOid: true},
			expected: &matchingData{propertyName: emptyStr, delimiter: emptyStr,
				fieldName: emptyStr, fieldPaths: emptyList,
				useEntityOid: true},
		},
	}

	for _, item := range table {
		matchingData := newMatchingData(item.md)
		expectedData := item.expected

		if expectedData.propertyName == emptyStr {
			assert.Nil(t, matchingData.GetMatchingProperty())
		} else {
			assert.Equal(t, expectedData.propertyName, matchingData.GetMatchingProperty().GetPropertyName())
		}

		assert.Equal(t, expectedData.delimiter, matchingData.GetDelimiter())

		if expectedData.fieldName == emptyStr {
			assert.Nil(t, matchingData.GetMatchingField())
		} else {
			assert.Equal(t, expectedData.fieldName, matchingData.GetMatchingField().GetFieldName())
		}

		if !expectedData.useEntityOid {
			assert.Nil(t, matchingData.GetMatchingEntityOid())
		} else {
			assert.NotNil(t, matchingData.GetMatchingEntityOid())
		}
	}
}

// tests MergedEntityMetadataBuilder
func TestMergedEntityMetadataBuilderInternalMatching(t *testing.T) {
	builder := NewMergedEntityMetadataBuilder()

	prop1 := "IP"
	prop2 := "Port"

	var props []string
	props = append(props, prop1)
	props = append(props, prop2)

	builder.InternalMatchingType(MergedEntityMetadata_STRING)
	builder.ExternalMatchingType(MergedEntityMetadata_STRING)
	builder.InternalMatchingProperty(prop2)
	builder.InternalMatchingProperty(prop1)

	md, err := builder.Build()

	assert.Nil(t, err)

	assert.Equal(t, proto.MergedEntityMetadata_STRING, md.GetMatchingMetadata().GetReturnType())
	assert.Equal(t, 2, len(md.GetMatchingMetadata().GetMatchingData()))
	assert.NotNil(t, md.GetMatchingMetadata().GetMatchingData())

	for _, data := range md.GetMatchingMetadata().GetMatchingData() {
		assert.NotNil(t, data.GetMatchingProperty())
		assert.Nil(t, data.GetMatchingField())
		assert.Nil(t, data.GetMatchingEntityOid())

		assert.Contains(t, props, data.GetMatchingProperty().GetPropertyName())
		assert.Equal(t, "", data.GetDelimiter())
	}
}

func TestMergedEntityMetadataBuilderExternalMatching(t *testing.T) {
	builder := NewMergedEntityMetadataBuilder()

	prop1 := "IP"

	builder.InternalMatchingType(MergedEntityMetadata_STRING)
	builder.ExternalMatchingType(MergedEntityMetadata_STRING)
	builder.ExternalMatchingPropertyWithDelimiter(prop1, ",")

	md, err := builder.Build()

	assert.Nil(t, err)

	assert.Equal(t, proto.MergedEntityMetadata_STRING, md.GetMatchingMetadata().GetExternalEntityReturnType())
	assert.Equal(t, 1, len(md.GetMatchingMetadata().GetExternalEntityMatchingProperty()))
	assert.NotNil(t, md.GetMatchingMetadata().GetExternalEntityMatchingProperty())

	for _, data := range md.GetMatchingMetadata().GetExternalEntityMatchingProperty() {
		assert.NotNil(t, data.GetMatchingProperty())
		assert.Nil(t, data.GetMatchingField())
		assert.Nil(t, data.GetMatchingEntityOid())

		assert.Equal(t, prop1, data.GetMatchingProperty().GetPropertyName())
		assert.Equal(t, ",", data.GetDelimiter())
	}
}
