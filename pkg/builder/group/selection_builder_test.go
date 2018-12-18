package group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestGenericSelectionSpecBuilder(t *testing.T) {
	selectionSpec := StringProperty().
		Name("Propert1").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("Hello").Build()

	fmt.Printf("Spec %++v\n", selectionSpec)
}

func TestStringSelectionSpecBuilder(t *testing.T) {
	propVal := "Hello"
	selectionSpec := StringProperty().SetProperty(propVal).Build()
	assert.EqualValues(t, propVal, selectionSpec.GetPropertyValueString())

	fmt.Printf("Spec %++v\n", selectionSpec)
}

func TestStringListSelectionSpecBuilder(t *testing.T) {
	propVal := []string{"Hello, World"}
	selectionSpec := StringListProperty().SetProperty(propVal).Build()
	assert.EqualValues(t, propVal, selectionSpec.GetPropertyValueStringList().PropertyValue)

	fmt.Printf("Spec %++v\n", selectionSpec)
}

func TestDoubleSelectionSpecBuilder(t *testing.T) {
	propVal := 500.0
	selectionSpec := DoubleProperty().SetProperty(propVal).Build()
	assert.EqualValues(t, propVal, selectionSpec.GetPropertyValueDouble())

	fmt.Printf("Spec %++v\n", selectionSpec)
}

func TestDoubleListSelectionSpecBuilder(t *testing.T) {
	propVal := []float64{500.0, 600.0}
	selectionSpec := DoubleListProperty().SetProperty(propVal).Build()
	assert.EqualValues(t, propVal, selectionSpec.GetPropertyValueDoubleList().PropertyValue)

	fmt.Printf("Spec %++v\n", selectionSpec)
}
