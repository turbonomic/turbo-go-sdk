package group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestBuyerDataStaticBuyers(t *testing.T) {
	eType := proto.EntityDTO_CONTAINER

	containers := []string{"container1", "container2"}

	buyerData := StaticBuyers(containers).OfType(eType)
	fmt.Printf("%++v\n", buyerData)
	assert.EqualValues(t, len(containers), len(buyerData.entities))
	assert.EqualValues(t, containers, buyerData.entities)
	assert.Nil(t, buyerData.matchingBuyers)
	assert.EqualValues(t, eType, *buyerData.entityTypePtr)
	assert.EqualValues(t, 0, buyerData.atMost)
}

func TestBuyerDataDynamicBuyers(t *testing.T) {
	eType := proto.EntityDTO_CONTAINER

	selectionSpec1 := StringProperty().
		Name("Name").
		Expression(proto.GroupDTO_SelectionSpec_CONTAINS).
		SetProperty("vm_")

	buyerData := DynamicBuyers(SelectedBy(selectionSpec1)).OfType(eType)
	fmt.Printf("%++v\n", buyerData)

	buyerSpec := buyerData.matchingBuyers.selectionSpecBuilderList
	for _, spec := range buyerSpec {
		fmt.Printf("%++v\n", spec.Build())
	}
	assert.EqualValues(t, eType, *buyerData.entityTypePtr)
	assert.EqualValues(t, 0, buyerData.atMost)
	assert.NotNil(t, buyerData.matchingBuyers)
	assert.EqualValues(t, 0, len(buyerData.entities))
}

func TestBuyerDataWithAtMost(t *testing.T) {
	eType := proto.EntityDTO_CONTAINER
	containers := []string{"container1", "container2"}

	var maxBuyers int32 = 4
	buyerData := StaticBuyers(containers).AtMost(maxBuyers).OfType(eType)
	assert.EqualValues(t, maxBuyers, buyerData.atMost)

	buyerData = StaticBuyers(containers).OfType(eType)
	assert.EqualValues(t, 0, buyerData.atMost)

	buyerData = buyerData.AtMost(10)
	assert.EqualValues(t, 10, buyerData.atMost)
}
