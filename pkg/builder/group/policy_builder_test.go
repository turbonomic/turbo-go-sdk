package group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestPlacePolicy(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER
	pType := proto.EntityDTO_VIRTUAL_MACHINE

	containers := []string{"container1", "container2"}
	vms := []string{"vm1", "vm2"}
	var maxBuyers int32 = 4

	placePolicyBuilder := Place(id).
		WithBuyers(StaticBuyers(containers).OfType(eType).AtMost(maxBuyers)).
		OnSellers(StaticSellers(vms).OfType(pType))

	groupDTOList, _ := assertPlacePolicyConditions(t, placePolicyBuilder)

	for _, groupDTO := range groupDTOList {
		fmt.Printf("PLACE: %++v\n", groupDTO)
	}
}

func TestPlacePolicyDynamicGroups(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER
	pType := proto.EntityDTO_VIRTUAL_MACHINE

	var maxBuyers int32 = 4

	selectionSpec1 := StringProperty().
		Name("Name").
		Expression(proto.GroupDTO_SelectionSpec_CONTAINS).
		SetProperty("app_")

	selectionSpec2 := StringProperty().
		Name("Name").
		Expression(proto.GroupDTO_SelectionSpec_CONTAINS).
		SetProperty("vm_")

	placePolicyBuilder := Place(id).
		WithBuyers(DynamicBuyers(SelectedBy(selectionSpec1)).OfType(eType).AtMost(maxBuyers)).
		OnSellers(DynamicSellers(SelectedBy(selectionSpec2)).OfType(pType))
	groupDTOList, _ := assertPlacePolicyConditions(t, placePolicyBuilder)

	var buyerGroup, sellerGroup *proto.GroupDTO
	for _, groupDTO := range groupDTOList {
		if groupDTO.GetConstraintInfo().GetIsBuyer() {
			buyerGroup = groupDTO
		} else {
			sellerGroup = groupDTO
		}
	}
	assertDynamicGroupDTO(t, buyerGroup, eType, []*GenericSelectionSpecBuilder{selectionSpec1})
	assertDynamicGroupDTO(t, sellerGroup, pType, []*GenericSelectionSpecBuilder{selectionSpec2})

	for _, groupDTO := range groupDTOList {
		fmt.Printf("PLACE: %++v\n", groupDTO)
	}
}

func TestDoNotPlacePolicy(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER
	pType := proto.EntityDTO_VIRTUAL_MACHINE
	containers := []string{"container1", "container2"}

	selectionSpec1 := StringProperty().
		Name("Name").
		Expression(proto.GroupDTO_SelectionSpec_CONTAINS).
		SetProperty("vm_")

	placePolicyBuilder := DoNotPlace(id).
		WithBuyers(StaticBuyers(containers).OfType(eType).AtMost(4)).
		OnSellers(DynamicSellers(SelectedBy(selectionSpec1)).OfType(pType))

	groupDTOList, _ := assertDoNotPlacePolicyConditions(t, placePolicyBuilder)

	var buyerGroup, sellerGroup *proto.GroupDTO
	for _, groupDTO := range groupDTOList {
		if groupDTO.GetConstraintInfo().GetIsBuyer() {
			buyerGroup = groupDTO
		} else {
			sellerGroup = groupDTO
		}
	}
	assertStaticGroupDTO(t, buyerGroup, eType, containers)
	assertDynamicGroupDTO(t, sellerGroup, pType, []*GenericSelectionSpecBuilder{selectionSpec1})

	for _, groupDTO := range groupDTOList {
		fmt.Printf("DO NOT PLACE: %++v\n", groupDTO)
	}
}

func TestPlacePolicyInvalidGroups(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER
	pType := proto.EntityDTO_VIRTUAL_MACHINE

	containers := []string{"container1", "container2"}
	vms := []string{"vm1", "vm2"}

	// Missing seller group
	placePolicyBuilder := Place(id).
		WithBuyers(StaticBuyers(containers).OfType(eType))

	assertPlacePolicyConditions(t, placePolicyBuilder)

	// Missing buyer group
	placePolicyBuilder = Place(id).
		OnSellers(StaticSellers(vms).OfType(pType))

	assertPlacePolicyConditions(t, placePolicyBuilder)

	// Invalid buyer group
	placePolicyBuilder = Place(id).
		WithBuyers(StaticBuyers(containers)).
		OnSellers(StaticSellers(vms).OfType(pType))

	assertPlacePolicyConditions(t, placePolicyBuilder)

	// Invalid seller group
	placePolicyBuilder = Place(id).
		WithBuyers(StaticBuyers(containers).OfType(eType)).
		OnSellers(StaticSellers(vms))

	assertPlacePolicyConditions(t, placePolicyBuilder)
}

func TestDoNotPlacePolicyInvalidGroups(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER
	pType := proto.EntityDTO_VIRTUAL_MACHINE

	containers := []string{"container1", "container2"}
	vms := []string{"vm1", "vm2"}

	// Missing seller group
	doNotPlacePolicyBuilder := DoNotPlace(id).
		WithBuyers(StaticBuyers(containers).OfType(eType))

	assertDoNotPlacePolicyConditions(t, doNotPlacePolicyBuilder)

	// Missing buyer group
	doNotPlacePolicyBuilder = DoNotPlace(id).
		OnSellers(StaticSellers(vms).OfType(pType))

	assertDoNotPlacePolicyConditions(t, doNotPlacePolicyBuilder)

	// Invalid buyer group
	doNotPlacePolicyBuilder = DoNotPlace(id).
		WithBuyers(StaticBuyers(containers)).
		OnSellers(StaticSellers(vms).OfType(pType))

	assertDoNotPlacePolicyConditions(t, doNotPlacePolicyBuilder)

	// Invalid seller group
	doNotPlacePolicyBuilder = DoNotPlace(id).
		WithBuyers(StaticBuyers(containers).OfType(eType)).
		OnSellers(StaticSellers(vms))

	assertDoNotPlacePolicyConditions(t, doNotPlacePolicyBuilder)
}

func assertPlacePolicyConditions(t *testing.T, placePolicyBuilder *PlacePolicyBuilder) ([]*proto.GroupDTO, error) {
	assert.NotNil(t, placePolicyBuilder)

	var isValidBuyerData, isValidSellerData bool
	isValidBuyerData = assertValidBuyerData(t, placePolicyBuilder.buyerData)
	isValidSellerData = assertValidSellerData(t, placePolicyBuilder.sellerData)

	groupDTOList, err := placePolicyBuilder.Build()
	if isValidBuyerData && isValidSellerData {
		assert.Nil(t, err)
		assert.EqualValues(t, 2, len(groupDTOList))
	} else {
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, len(groupDTOList))
	}
	return groupDTOList, err
}

func assertDoNotPlacePolicyConditions(t *testing.T, doNotPlacePolicyBuilder *DoNotPlacePolicyBuilder) ([]*proto.GroupDTO, error) {
	assert.NotNil(t, doNotPlacePolicyBuilder)
	var isValidBuyerData, isValidSellerData bool
	isValidBuyerData = assertValidBuyerData(t, doNotPlacePolicyBuilder.buyerData)
	isValidSellerData = assertValidSellerData(t, doNotPlacePolicyBuilder.sellerData)

	groupDTOList, err := doNotPlacePolicyBuilder.Build()
	if isValidBuyerData && isValidSellerData {
		assert.Nil(t, err)
		assert.EqualValues(t, 2, len(groupDTOList))
	} else {
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, len(groupDTOList))
	}
	return groupDTOList, err
}

func assertPlaceTogetherPolicyConditions(t *testing.T, placeTogetherPolicyBuilder *PlaceTogetherPolicyBuilder) ([]*proto.GroupDTO, error) {
	assert.NotNil(t, placeTogetherPolicyBuilder)

	var isValidBuyerData bool
	isValidBuyerData = assertValidBuyerData(t, placeTogetherPolicyBuilder.buyerData)

	groupDTOList, err := placeTogetherPolicyBuilder.Build()
	if isValidBuyerData {
		assert.Nil(t, err)
		assert.EqualValues(t, 1, len(groupDTOList))
	} else {
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, len(groupDTOList))
	}
	return groupDTOList, err
}

func assertDoNotPlaceTogetherPolicyConditions(t *testing.T, doNotPlaceTogetherPolicyBuilder *DoNotPlaceTogetherPolicyBuilder) ([]*proto.GroupDTO, error) {
	assert.NotNil(t, doNotPlaceTogetherPolicyBuilder)

	var isValidBuyerData bool
	isValidBuyerData = assertValidBuyerData(t, doNotPlaceTogetherPolicyBuilder.buyerData)

	fmt.Printf("### [assertDoNotPlaceTogetherPolicyConditions] isValidBuyerData:%v \n", isValidBuyerData)

	groupDTOList, err := doNotPlaceTogetherPolicyBuilder.Build()
	if isValidBuyerData {
		assert.Nil(t, err)
		assert.EqualValues(t, 1, len(groupDTOList))
	} else {
		assert.NotNil(t, err)
		assert.EqualValues(t, 0, len(groupDTOList))
	}
	return groupDTOList, err
}

func assertValidBuyerData(t *testing.T, buyerData *BuyerPolicyData) bool {
	if buyerData == nil {
		return false
	}
	if buyerData.entityTypePtr == nil {
		return false
	}

	if buyerData.entities == nil && buyerData.matchingBuyers == nil {
		return false
	}
	return true
}

func assertValidSellerData(t *testing.T, sellerData *SellerPolicyData) bool {
	if sellerData == nil {
		return false
	}
	if sellerData.entityTypePtr == nil {
		return false
	}

	if sellerData.entities == nil && sellerData.matchingBuyers == nil {
		return false
	}
	return true
}

func assertStaticGroupDTO(t *testing.T, group *proto.GroupDTO, eType proto.EntityDTO_EntityType, members []string) {
	assert.EqualValues(t, group.EntityType, &eType)
	assert.EqualValues(t, len(members), len(group.GetMemberList().GetMember()))
	assert.EqualValues(t, members, group.GetMemberList().GetMember())
	assert.Nil(t, group.GetSelectionSpecList())
}

func assertDynamicGroupDTO(t *testing.T, group *proto.GroupDTO, eType proto.EntityDTO_EntityType, members []*GenericSelectionSpecBuilder) {
	assert.EqualValues(t, group.EntityType, &eType)
	assert.EqualValues(t, len(members), len(group.GetSelectionSpecList().GetSelectionSpec()))
	assert.Nil(t, group.GetMemberList())
}

func TestPlaceTogetherPolicy(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER

	selectionSpec1 := StringProperty().
		Name("Name").
		Expression(proto.GroupDTO_SelectionSpec_CONTAINS).
		SetProperty("vm_")
	selectionSpec2 := StringProperty().
		Name("Id").
		Expression(proto.GroupDTO_SelectionSpec_CONTAINS).
		SetProperty("3333-")

	placeTogetherPolicyBuilder := PlaceTogether(id).
		WithBuyers(DynamicBuyers(SelectedBy(selectionSpec1).and(selectionSpec2)).OfType(eType))

	groupDTOList, _ := assertPlaceTogetherPolicyConditions(t, placeTogetherPolicyBuilder)
	for _, groupDTO := range groupDTOList {
		fmt.Printf("PLACE TOGETHER: %++v\n", groupDTO)
	}

	// Invalid buyer group
	placeTogetherPolicyBuilder = PlaceTogether(id).
		WithBuyers(DynamicBuyers(SelectedBy(selectionSpec1).and(selectionSpec2)))

	isValidBuyerData := assertValidBuyerData(t, placeTogetherPolicyBuilder.buyerData)
	assert.Equal(t, false, isValidBuyerData)
	assertPlaceTogetherPolicyConditions(t, placeTogetherPolicyBuilder)
}

func TestDoNotPlaceTogetherPolicy(t *testing.T) {
	id := "policy1"
	eType := proto.EntityDTO_CONTAINER

	buyerData := StaticBuyers([]string{"container1", "container2"}).AtMost(4).OfType(eType)
	fmt.Printf("%++v\n", buyerData)

	doNotPlaceTogetherPolicyBuilder := DoNotPlaceTogether(id).
		WithBuyers(StaticBuyers([]string{"container1", "container2"}).OfType(eType))

	groupDTOList, _ := assertDoNotPlaceTogetherPolicyConditions(t, doNotPlaceTogetherPolicyBuilder)
	for _, groupDTO := range groupDTOList {
		fmt.Printf("DO NOT PLACE TOGETHER: %++v\n", groupDTO)
	}

	// Invalid buyer group
	doNotPlaceTogetherPolicyBuilder = DoNotPlaceTogether(id).
		WithBuyers(StaticBuyers([]string{"container1", "container2"}))

	isValidBuyerData := assertValidBuyerData(t, doNotPlaceTogetherPolicyBuilder.buyerData)
	assert.Equal(t, false, isValidBuyerData)
	assertDoNotPlaceTogetherPolicyConditions(t, doNotPlaceTogetherPolicyBuilder)
}
