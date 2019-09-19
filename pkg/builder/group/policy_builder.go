package group

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// Build the constraint for the buyer group of the policy
func buyerGroupConstraint(constraintType proto.GroupDTO_ConstraintType, constraintId string) *ConstraintInfoBuilder {
	constraintInfoBuilder := newConstraintBuilder(constraintType, constraintId) //.BuyerGroup()
	constraintInfoBuilder.isBuyer = true

	return constraintInfoBuilder
}

// Build the constraint for the seller group of the policy
func sellerGroupConstraint(constraintType proto.GroupDTO_ConstraintType, constraintId string) *ConstraintInfoBuilder {
	constraintInfoBuilder := newConstraintBuilder(constraintType, constraintId)
	constraintInfoBuilder.isBuyer = false

	return constraintInfoBuilder
}

type buyerSellerPolicyData struct {
	policyId       string
	displayName    string
	constraintType proto.GroupDTO_ConstraintType
	buyerData      *BuyerPolicyData
	sellerData     *SellerPolicyData
}

type buyerBuyerPolicyData struct {
	policyId       string
	displayName    string
	constraintType proto.GroupDTO_ConstraintType
	buyerData      *BuyerPolicyData
	sellerTypePtr  *proto.EntityDTO_EntityType
}

type PlacePolicyBuilder struct {
	*buyerSellerPolicyData
}

type DoNotPlacePolicyBuilder struct {
	*buyerSellerPolicyData
}

type PlaceTogetherPolicyBuilder struct {
	*buyerBuyerPolicyData
}

type DoNotPlaceTogetherPolicyBuilder struct {
	*buyerBuyerPolicyData
}

////========================================================================

func Place(policyId string) *PlacePolicyBuilder {
	placePolicy := &PlacePolicyBuilder{
		&buyerSellerPolicyData{
			policyId:       policyId,
			constraintType: proto.GroupDTO_BUYER_SELLER_AFFINITY,
		},
	}
	return placePolicy
}

func (place *PlacePolicyBuilder) WithDisplayName(displayName string) *PlacePolicyBuilder {
	place.displayName = displayName
	return place
}

func (place *PlacePolicyBuilder) WithBuyers(buyers *BuyerPolicyData) *PlacePolicyBuilder {
	place.buyerData = buyers
	return place
}

func (place *PlacePolicyBuilder) OnSellers(sellers *SellerPolicyData) *PlacePolicyBuilder {
	place.sellerData = sellers
	return place
}

func (place *PlacePolicyBuilder) Build() ([]*proto.GroupDTO, error) {
	return buildBuyerSellerPolicyGroup(place.buyerSellerPolicyData)
}

func DoNotPlace(policyId string) *DoNotPlacePolicyBuilder {
	doNotPlace := &DoNotPlacePolicyBuilder{
		&buyerSellerPolicyData{
			policyId:       policyId,
			constraintType: proto.GroupDTO_BUYER_SELLER_ANTI_AFFINITY,
		},
	}
	return doNotPlace
}

func (doNotPlace *DoNotPlacePolicyBuilder) WithDisplayName(displayName string) *DoNotPlacePolicyBuilder {
	doNotPlace.displayName = displayName
	return doNotPlace
}

func (doNotPlace *DoNotPlacePolicyBuilder) WithBuyers(buyers *BuyerPolicyData) *DoNotPlacePolicyBuilder {
	doNotPlace.buyerData = buyers
	return doNotPlace
}

func (doNotPlace *DoNotPlacePolicyBuilder) OnSellers(sellers *SellerPolicyData) *DoNotPlacePolicyBuilder {
	doNotPlace.sellerData = sellers
	return doNotPlace
}

func (doNotPlace *DoNotPlacePolicyBuilder) Build() ([]*proto.GroupDTO, error) {
	return buildBuyerSellerPolicyGroup(doNotPlace.buyerSellerPolicyData)
}

////========================================================================

func PlaceTogether(policyId string) *PlaceTogetherPolicyBuilder {
	placeTogetherPolicy := &PlaceTogetherPolicyBuilder{
		&buyerBuyerPolicyData{
			policyId:       policyId,
			constraintType: proto.GroupDTO_BUYER_BUYER_AFFINITY,
		},
	}
	return placeTogetherPolicy
}

func (placeTogether *PlaceTogetherPolicyBuilder) WithDisplayName(
	displayName string) *PlaceTogetherPolicyBuilder {
	placeTogether.displayName = displayName
	return placeTogether
}

func (placeTogether *PlaceTogetherPolicyBuilder) WithBuyers(
	buyers *BuyerPolicyData) *PlaceTogetherPolicyBuilder {
	placeTogether.buyerData = buyers
	return placeTogether
}

func (placeTogether *PlaceTogetherPolicyBuilder) OnSellerType(
	sellerType proto.EntityDTO_EntityType) *PlaceTogetherPolicyBuilder {
	placeTogether.sellerTypePtr = &sellerType
	return placeTogether
}

func (placeTogether *PlaceTogetherPolicyBuilder) Build() ([]*proto.GroupDTO, error) {
	return buildBuyerBuyerPolicyGroup(placeTogether.buyerBuyerPolicyData)
}

func DoNotPlaceTogether(policyId string) *DoNotPlaceTogetherPolicyBuilder {
	doNotPlaceTogether := &DoNotPlaceTogetherPolicyBuilder{
		&buyerBuyerPolicyData{
			policyId:       policyId,
			constraintType: proto.GroupDTO_BUYER_BUYER_ANTI_AFFINITY,
		},
	}
	return doNotPlaceTogether
}

func (doNotPlaceTogether *DoNotPlaceTogetherPolicyBuilder) WithDisplayName(
	displayName string) *DoNotPlaceTogetherPolicyBuilder {
	doNotPlaceTogether.displayName = displayName
	return doNotPlaceTogether
}

func (doNotPlaceTogether *DoNotPlaceTogetherPolicyBuilder) WithBuyers(
	buyers *BuyerPolicyData) *DoNotPlaceTogetherPolicyBuilder {
	doNotPlaceTogether.buyerData = buyers
	return doNotPlaceTogether
}

func (doNotPlaceTogether *DoNotPlaceTogetherPolicyBuilder) OnSellerType(
	sellerType proto.EntityDTO_EntityType) *DoNotPlaceTogetherPolicyBuilder {
	doNotPlaceTogether.sellerTypePtr = &sellerType
	return doNotPlaceTogether
}

func (doNotPlaceTogether *DoNotPlaceTogetherPolicyBuilder) Build() ([]*proto.GroupDTO, error) {
	return buildBuyerBuyerPolicyGroup(doNotPlaceTogether.buyerBuyerPolicyData)
}

////========================================================================
func buildBuyerSellerPolicyGroup(policyData *buyerSellerPolicyData) ([]*proto.GroupDTO, error) {
	if policyData.buyerData == nil {
		return nil, fmt.Errorf("[buildBuyerSellerPolicyGroup] Buyer group data not set")
	}
	if policyData.sellerData == nil {
		return nil, fmt.Errorf("[buildBuyerSellerPolicyGroup] Seller group data not set")
	}

	var groupDTOs []*proto.GroupDTO

	// Buyer group and constraints
	buyerGroup, err := createPolicyBuyerGroup(
		policyData.policyId,
		policyData.displayName,
		policyData.constraintType,
		policyData.buyerData,
		policyData.sellerData)
	if err != nil {
		return []*proto.GroupDTO{}, err
	} else {
		groupDTO, err := buyerGroup.Build()
		if err != nil {
			return []*proto.GroupDTO{}, err
		} else {
			groupDTOs = append(groupDTOs, groupDTO)
		}
	}

	// Seller group and constraints
	sellerGroup, err := createPolicySellerGroup(
		policyData.policyId,
		policyData.displayName,
		policyData.constraintType,
		policyData.sellerData)
	if err != nil {
		return []*proto.GroupDTO{}, err
	} else {
		groupDTO, err := sellerGroup.Build()
		if err != nil {
			return []*proto.GroupDTO{}, err
		} else {
			groupDTOs = append(groupDTOs, groupDTO)
		}
	}

	return groupDTOs, nil
}

// Set up buyer group for a policy
func createPolicyBuyerGroup(policyId string, displayName string, constraintType proto.GroupDTO_ConstraintType,
	buyerData *BuyerPolicyData, sellerData *SellerPolicyData) (*AbstractConstraintGroupBuilder, error) {
	if buyerData.entityTypePtr == nil {
		return nil, fmt.Errorf("buyer entity type is not set for policy %s", policyId)
	}
	if sellerData.entityTypePtr == nil {
		return nil, fmt.Errorf("seller entity type is not set for policy %s", policyId)
	}
	var buyerGroup *AbstractBuilder
	entityType := *buyerData.entityTypePtr
	if buyerData.entities != nil {
		buyerGroup = StaticGroup(policyId).
			OfType(entityType).
			WithEntities(buyerData.entities)
	} else if buyerData.matchingBuyers != nil {
		buyerGroup = DynamicGroup(policyId).
			OfType(entityType).
			MatchingEntities(buyerData.matchingBuyers)
	} else {
		return nil, fmt.Errorf("buyer group member data is missing for policy %s", policyId)
	}

	// constraint info
	var constraintInfoBuilder *ConstraintInfoBuilder
	constraintInfoBuilder = buyerGroupConstraint(constraintType, policyId)
	if buyerData.atMost > 0 {
		constraintInfoBuilder.AtMostBuyers(buyerData.atMost)
	}
	constraintInfoBuilder.WithSellerType(*sellerData.entityTypePtr)
	if displayName != "" {
		constraintInfoBuilder.WithDisplayName(displayName)
	}

	buyerConstraintGroup := &AbstractConstraintGroupBuilder{
		AbstractBuilder:       buyerGroup,
		ConstraintInfoBuilder: constraintInfoBuilder,
	}

	return buyerConstraintGroup, nil
}

// Set up seller group for a policy
func createPolicySellerGroup(policyId string, displayName string, constraintType proto.GroupDTO_ConstraintType,
	sellerData *SellerPolicyData) (*AbstractConstraintGroupBuilder, error) {

	var sellerGroup *AbstractBuilder
	if sellerData.entityTypePtr == nil {
		return nil, fmt.Errorf("seller entity type is not set for policy %s", policyId)
	}
	entityType := *sellerData.entityTypePtr
	if sellerData.entities != nil {
		sellerGroup = StaticGroup(policyId).
			OfType(entityType).
			WithEntities(sellerData.entities)
	} else if sellerData.matchingBuyers != nil {
		sellerGroup = DynamicGroup(policyId).
			OfType(entityType).
			MatchingEntities(sellerData.matchingBuyers)
	} else {
		return nil, fmt.Errorf("seller group member data is missing for policy %s", policyId)
	}

	// constraint info
	var constraintInfoBuilder *ConstraintInfoBuilder
	constraintInfoBuilder = sellerGroupConstraint(constraintType, policyId)
	if displayName != "" {
		constraintInfoBuilder.WithDisplayName(displayName)
	}
	sellerConstraintGroup := &AbstractConstraintGroupBuilder{
		AbstractBuilder:       sellerGroup,
		ConstraintInfoBuilder: constraintInfoBuilder,
	}

	return sellerConstraintGroup, nil
}

func buildBuyerBuyerPolicyGroup(policyData *buyerBuyerPolicyData) ([]*proto.GroupDTO, error) {
	if policyData.buyerData == nil {
		return nil, fmt.Errorf("[buildBuyerBuyerPolicyGroup] Buyer group data is not set")
	}
	if policyData.sellerTypePtr == nil {
		return nil, fmt.Errorf("[buildBuyerBuyerPolicyGroup] Seller entity type is not set")
	}
	var groupDTOs []*proto.GroupDTO

	// Buyer group and constraints
	sellerData := &SellerPolicyData{
		groupData: &groupData{
			entityTypePtr: policyData.sellerTypePtr,
		},
	}
	buyerGroup, err := createPolicyBuyerGroup(
		policyData.policyId,
		policyData.displayName,
		policyData.constraintType,
		policyData.buyerData,
		sellerData)
	if err != nil {
		return []*proto.GroupDTO{}, err
	} else {
		groupDTO, err := buyerGroup.Build()
		if err != nil {
			return []*proto.GroupDTO{}, err
		} else {
			groupDTOs = append(groupDTOs, groupDTO)
		}
	}

	return groupDTOs, nil
}
