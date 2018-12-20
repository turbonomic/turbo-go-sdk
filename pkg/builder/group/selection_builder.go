package group

import "github.com/turbonomic/turbo-go-sdk/pkg/proto"

type PropertyType string

const (
	STRING_PROP      PropertyType = "String"
	DOUBLE_PROP      PropertyType = "Double"
	STRING_LIST_PROP PropertyType = "StringList"
	DOUBLE_LIST_PROP PropertyType = "DoubleList"
)

type Matching struct {
	selectionSpecBuilderList []SelectionSpecBuilder
}

func SelectedBy(entitySpec SelectionSpecBuilder) *Matching {
	var selectionSpecBuilderList []SelectionSpecBuilder
	selectionSpecBuilderList = append(selectionSpecBuilderList, entitySpec)
	matching := &Matching{
		selectionSpecBuilderList: selectionSpecBuilderList,
	}
	return matching
}
func (matching *Matching) and(entitySpec SelectionSpecBuilder) *Matching {
	matching.selectionSpecBuilderList = append(matching.selectionSpecBuilderList, entitySpec)
	return matching
}

// ------------------------------------------------------------------------------------------------

type SelectionSpecBuilder interface {
	isSelectionSpecBuilder()
	Build() *proto.GroupDTO_SelectionSpec
}

type GenericSelectionSpecBuilder struct {
	selectionSpec *proto.GroupDTO_SelectionSpec
	propertyType  PropertyType
}

func newGenericSelectionSpecBuilder() *GenericSelectionSpecBuilder {
	builder := &GenericSelectionSpecBuilder{
		selectionSpec: &proto.GroupDTO_SelectionSpec{},
	}
	return builder
}

func StringProperty() *GenericSelectionSpecBuilder {
	builder := newGenericSelectionSpecBuilder()
	builder.propertyType = STRING_PROP

	return builder
}

func StringListProperty() *GenericSelectionSpecBuilder {
	builder := newGenericSelectionSpecBuilder()
	builder.propertyType = STRING_LIST_PROP

	return builder
}

func DoubleProperty() *GenericSelectionSpecBuilder {
	builder := newGenericSelectionSpecBuilder()
	builder.propertyType = DOUBLE_PROP

	return builder
}

func DoubleListProperty() *GenericSelectionSpecBuilder {
	builder := newGenericSelectionSpecBuilder()
	builder.propertyType = DOUBLE_LIST_PROP

	return builder
}

func (builder *GenericSelectionSpecBuilder) Name(propertyName string) *GenericSelectionSpecBuilder {
	builder.selectionSpec.Property = &propertyName
	return builder
}

func (builder *GenericSelectionSpecBuilder) Expression(expression proto.GroupDTO_SelectionSpec_ExpressionType) *GenericSelectionSpecBuilder {
	builder.selectionSpec.ExpressionType = &expression
	return builder
}

func (builder *GenericSelectionSpecBuilder) SetProperty(property interface{}) *GenericSelectionSpecBuilder {

	switch builder.propertyType {
	case STRING_PROP:
		//var propVal string
		propVal, ok := property.(string)
		if ok {
			propValString := &proto.GroupDTO_SelectionSpec_PropertyValueString{}
			propValString.PropertyValueString = propVal
			builder.selectionSpec.PropertyValue = propValString
		}
		break

	case STRING_LIST_PROP:
		propVal, ok := property.([]string)
		if ok {
			propValStringList := &proto.GroupDTO_SelectionSpec_PropertyValueStringList{
				PropertyValueStringList: &proto.GroupDTO_SelectionSpec_PropertyStringList{
					PropertyValue: propVal,
				},
			}

			builder.selectionSpec.PropertyValue = propValStringList
		}
		break

	case DOUBLE_PROP:
		propVal, ok := property.(float64)
		if ok {
			propValDouble := &proto.GroupDTO_SelectionSpec_PropertyValueDouble{
				PropertyValueDouble: propVal,
			}
			builder.selectionSpec.PropertyValue = propValDouble
		}
		break
	case DOUBLE_LIST_PROP: //TODO: unit test
		propVal, ok := property.([]float64)
		if ok {
			propValDoubleList := &proto.GroupDTO_SelectionSpec_PropertyValueDoubleList{
				PropertyValueDoubleList: &proto.GroupDTO_SelectionSpec_PropertyDoubleList{
					PropertyValue: propVal,
				},
			}
			builder.selectionSpec.PropertyValue = propValDoubleList
		}
		break
	}

	return builder
}

func (builder *GenericSelectionSpecBuilder) Build() *proto.GroupDTO_SelectionSpec {
	return builder.selectionSpec
}

func (builder *GenericSelectionSpecBuilder) isSelectionSpecBuilder() {}
