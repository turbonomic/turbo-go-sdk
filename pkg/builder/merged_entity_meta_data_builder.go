package builder

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type ReturnType string

const (
	MergedEntityMetadata_STRING      ReturnType = "String"
	MergedEntityMetadata_LIST_STRING ReturnType = "List"
)

var (
	returnTypeMapping = map[ReturnType]proto.MergedEntityMetadata_ReturnType{
		MergedEntityMetadata_STRING:      proto.MergedEntityMetadata_STRING,
		MergedEntityMetadata_LIST_STRING: proto.MergedEntityMetadata_LIST_STRING,
	}
)

// ============================ MergedEntityMetadata_MatchingMetadata ==============================
type matchingMetadataBuilder struct {
	internalReturnType   ReturnType
	externalReturnType   ReturnType
	internalMatchingData []*matchingData
	externalMatchingData []*matchingData
	commoditiesSold      []proto.CommodityDTO_CommodityType
	commoditiesBought    []*commodityBoughtMetadata
}

type matchingData struct {
	propertyName string
	delimiter    string
	fieldName    string
	fieldPaths   []string
	entityOid    string
}

func newMatchingMetadataBuilder() *matchingMetadataBuilder {

	builder := &matchingMetadataBuilder{
		internalMatchingData: []*matchingData{},
		externalMatchingData: []*matchingData{},
	}
	return builder
}

func (builder *matchingMetadataBuilder) build() (*proto.MergedEntityMetadata_MatchingMetadata, error) {
	matchingMetadata := &proto.MergedEntityMetadata_MatchingMetadata{}
	matchingMetadata.MatchingData = []*proto.MergedEntityMetadata_MatchingData{}
	matchingMetadata.ExternalEntityMatchingProperty = []*proto.MergedEntityMetadata_MatchingData{}

	// external property return type
	if builder.externalReturnType != "" {
		rtype, exists := returnTypeMapping[builder.externalReturnType]
		if exists {
			matchingMetadata.ExternalEntityReturnType = &rtype
		} else {
			return nil, fmt.Errorf("Unknown external entity metadata return type")
		}
	} else {
		return nil, fmt.Errorf("External entity metadata return type not set")
	}

	// internal property return type
	if builder.internalReturnType != "" {
		rtype, exists := returnTypeMapping[builder.internalReturnType]
		if exists {
			matchingMetadata.ReturnType = &rtype
		} else {
			return nil, fmt.Errorf("Unknown internal entity metadata return type")
		}
	} else {
		return nil, fmt.Errorf("Internal entity metadata return type not set")
	}

	// create internal property matching data
	var internalMatchingDataList []*proto.MergedEntityMetadata_MatchingData
	for _, internalData := range builder.internalMatchingData {
		internalMatchingDataList = append(internalMatchingDataList, newMatchingData(internalData))
	}

	// create external property matching data
	var externalMatchingDataList []*proto.MergedEntityMetadata_MatchingData
	for _, externalData := range builder.externalMatchingData {
		externalMatchingDataList = append(externalMatchingDataList, newMatchingData(externalData))
	}

	//
	matchingMetadata.MatchingData = internalMatchingDataList
	matchingMetadata.ExternalEntityMatchingProperty = externalMatchingDataList
	return matchingMetadata, nil
}

func (builder *matchingMetadataBuilder) addInternalMatchingData(internal *matchingData) *matchingMetadataBuilder {

	builder.internalMatchingData = append(builder.internalMatchingData, internal)
	return builder
}

func (builder *matchingMetadataBuilder) addExternalMatchingData(external *matchingData) *matchingMetadataBuilder {

	builder.externalMatchingData = append(builder.externalMatchingData, external)
	return builder
}

func newMatchingData(matchingData *matchingData) *proto.MergedEntityMetadata_MatchingData {
	// Create MergedEntityMetadata/MatchingMetadata/MatchingData for the internal property
	matchingDataBuilder := &proto.MergedEntityMetadata_MatchingData{}

	propertyName := matchingData.propertyName
	if propertyName != "" {
		entityPropertyNameBuilder := &proto.MergedEntityMetadata_EntityPropertyName{}
		entityPropertyNameBuilder.PropertyName = &propertyName

		matchingDataProperty := &proto.MergedEntityMetadata_MatchingData_MatchingProperty{}
		matchingDataProperty.MatchingProperty = entityPropertyNameBuilder

		matchingDataBuilder.MatchingData = matchingDataProperty
	}

	fieldName := matchingData.fieldName
	if fieldName != "" {
		entityFieldBuilder := &proto.MergedEntityMetadata_EntityField{}
		entityFieldBuilder.FieldName = &fieldName
		entityFieldBuilder.MessagePath = matchingData.fieldPaths

		matchingDataField := &proto.MergedEntityMetadata_MatchingData_MatchingField{}
		matchingDataField.MatchingField = entityFieldBuilder

		matchingDataBuilder.MatchingData = matchingDataField
	}

	return matchingDataBuilder
}

// ============================= MergedEntityMetadata_EntityPropertyName ===========================
type propertyBuilder struct {
	propertyName string
}

func newPropertyBuilder(propertyName string) *propertyBuilder {
	return &propertyBuilder{
		propertyName: propertyName,
	}
}

func (builder *propertyBuilder) build() *proto.MergedEntityMetadata_EntityPropertyName {
	if builder.propertyName != "" {
		entityPropertyNameBuilder := &proto.MergedEntityMetadata_EntityPropertyName{}
		entityPropertyNameBuilder.PropertyName = &builder.propertyName
		return entityPropertyNameBuilder
	}
	return nil
}

// =========================== MergedEntityMetadata_EntityField ===================================
type fieldBuilder struct {
	fieldName string
	paths     []string
}

func newFieldBuilder(fieldName string) *fieldBuilder {
	return &fieldBuilder{
		fieldName: fieldName,
	}
}

func (builder *fieldBuilder) build() *proto.MergedEntityMetadata_EntityField {
	fieldName := builder.fieldName
	if fieldName != "" {
		entityFieldBuilder := &proto.MergedEntityMetadata_EntityField{}
		entityFieldBuilder.FieldName = &fieldName
		entityFieldBuilder.MessagePath = builder.paths
		return entityFieldBuilder
	}
	return nil
}

// ============================== MergedEntityMetadata_CommodityBoughtMetadata =====================
type commodityBoughtMetadata struct {
	providerType     *proto.EntityDTO_EntityType
	replacesProvider *proto.EntityDTO_EntityType
	commBought       []proto.CommodityDTO_CommodityType
}

func newCommodityBoughtMetada(commBought []proto.CommodityDTO_CommodityType) *commodityBoughtMetadata {
	return &commodityBoughtMetadata{
		commBought: commBought,
	}
}

// ============================== MergedEntityMetadataBuilder ======================================

type MergedEntityMetadataBuilder struct {
	// MergedEntityMetadata - the main protobuf structure to return
	metadata *proto.MergedEntityMetadata

	// MergedEntityMetadata consists of  MatchingMetadata for proxy and external entity
	*matchingMetadataBuilder
	propertyBuilders   []*propertyBuilder
	fieldBuilders      []*fieldBuilder
	commBoughtMetadata []*commodityBoughtMetadata
}

func NewMergedEntityMetadataBuilder() *MergedEntityMetadataBuilder {
	builder := &MergedEntityMetadataBuilder{
		metadata:                &proto.MergedEntityMetadata{},
		matchingMetadataBuilder: newMatchingMetadataBuilder(),
	}

	return builder
}

func (builder *MergedEntityMetadataBuilder) build() (*proto.MergedEntityMetadata, error) {
	metadata := &proto.MergedEntityMetadata{}

	// Add the internal and external property matching metadata
	matchingMetadata, err := builder.matchingMetadataBuilder.build()
	if err != nil {
		return nil, err
	}
	metadata.MatchingMetadata = matchingMetadata

	if len(builder.commoditiesSold) > 0 {
		metadata.CommoditiesSold = append(metadata.CommoditiesSold, builder.commoditiesSold...)
	}

	if len(builder.commBoughtMetadata) > 0 {
		for _, commBought := range builder.commBoughtMetadata {
			commMetadata := &proto.MergedEntityMetadata_CommodityBoughtMetadata{
				CommodityMetadata: commBought.commBought,
				ProviderType:      commBought.providerType,
				ReplacesProvider:  commBought.replacesProvider,
			}
			metadata.CommoditiesBought = append(metadata.CommoditiesBought, commMetadata)
		}
	}
	return metadata, nil
}

func (builder *MergedEntityMetadataBuilder) keepStandAlone(bool_val bool) *MergedEntityMetadataBuilder {
	builder.metadata.KeepStandalone = &bool_val
	return builder
}

func (builder *MergedEntityMetadataBuilder) internalMatchingType(returnType ReturnType) *MergedEntityMetadataBuilder {

	builder.matchingMetadataBuilder.internalReturnType = returnType
	return builder
}

func (builder *MergedEntityMetadataBuilder) internalMatchingProperty(propertyName string) *MergedEntityMetadataBuilder {
	internal := &matchingData{
		propertyName: propertyName,
	}
	builder.matchingMetadataBuilder.addInternalMatchingData(internal)

	return builder
}

func (builder *MergedEntityMetadataBuilder) internalMatchingField(fieldName string, fieldPaths []string) *MergedEntityMetadataBuilder {
	internal := &matchingData{
		fieldName:  fieldName,
		fieldPaths: fieldPaths,
	}
	builder.matchingMetadataBuilder.addInternalMatchingData(internal)

	return builder
}

func (builder *MergedEntityMetadataBuilder) externalMatchingType(returnType ReturnType) *MergedEntityMetadataBuilder {

	builder.matchingMetadataBuilder.externalReturnType = returnType
	return builder
}

func (builder *MergedEntityMetadataBuilder) externalMatchingProperty(propertyName string) *MergedEntityMetadataBuilder {
	external := &matchingData{
		propertyName: propertyName,
	}
	builder.matchingMetadataBuilder.addExternalMatchingData(external)

	return builder
}

func (builder *MergedEntityMetadataBuilder) externalMatchingField(fieldName string, fieldPaths []string) *MergedEntityMetadataBuilder {
	external := &matchingData{
		fieldName:  fieldName,
		fieldPaths: fieldPaths,
	}
	builder.matchingMetadataBuilder.addInternalMatchingData(external)

	return builder
}

func (builder *MergedEntityMetadataBuilder) patchProperty(propertyName string) *MergedEntityMetadataBuilder {
	builder.propertyBuilders = append(builder.propertyBuilders,
		newPropertyBuilder(propertyName))
	return builder
}

func (builder *MergedEntityMetadataBuilder) patchFields(fieldName string) *MergedEntityMetadataBuilder {
	builder.fieldBuilders = append(builder.fieldBuilders,
		newFieldBuilder(fieldName))
	return builder
}

func (builder *MergedEntityMetadataBuilder) patchSold(commType proto.CommodityDTO_CommodityType) *MergedEntityMetadataBuilder {
	builder.commoditiesSold = append(builder.commoditiesSold, commType)
	return builder
}

func (builder *MergedEntityMetadataBuilder) patchSoldList(commType []proto.CommodityDTO_CommodityType) *MergedEntityMetadataBuilder {
	builder.commoditiesSold = append(builder.commoditiesSold, commType...)
	return builder
}

func (builder *MergedEntityMetadataBuilder) patchBought(commType []proto.CommodityDTO_CommodityType,
	providerType *proto.EntityDTO_EntityType) *MergedEntityMetadataBuilder {

	commBoughtMetadata := newCommodityBoughtMetada(commType)
	commBoughtMetadata.providerType = providerType

	builder.commBoughtMetadata = append(builder.commBoughtMetadata, commBoughtMetadata)
	return builder
}

func (builder *MergedEntityMetadataBuilder) patchBoughtAndReplaceProvider(
	commType []proto.CommodityDTO_CommodityType,
	providerType *proto.EntityDTO_EntityType,
	replacesProvider *proto.EntityDTO_EntityType) *MergedEntityMetadataBuilder {

	commBoughtMetadata := newCommodityBoughtMetada(commType)
	commBoughtMetadata.providerType = providerType
	commBoughtMetadata.replacesProvider = replacesProvider

	builder.commBoughtMetadata = append(builder.commBoughtMetadata, commBoughtMetadata)
	return builder
}
