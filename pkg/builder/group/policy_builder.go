package group

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// -------------------------------------------------------------------------------------------------
// Builder for the constraint info data in a Group DTO
type ConstraintInfoBuilder struct {
	constraintType        proto.GroupDTO_ConstraintType
	constraintId          string
	constraintName        string
	constraintDisplayName string
	providerTypePtr       *proto.EntityDTO_EntityType
	isBuyer               bool

	maxBuyers int32
}

func newConstraintBuilder(constraintType proto.GroupDTO_ConstraintType, constraintId string) *ConstraintInfoBuilder {

	return &ConstraintInfoBuilder{
		constraintName: constraintId,
		constraintType: constraintType,
	}
}

func (constraintInfoBuilder *ConstraintInfoBuilder) WithName(constraintName string) *ConstraintInfoBuilder {
	constraintInfoBuilder.constraintName = constraintName
	return constraintInfoBuilder
}

func (constraintInfoBuilder *ConstraintInfoBuilder) WithDisplayName(constraintDisplayName string) *ConstraintInfoBuilder {
	constraintInfoBuilder.constraintDisplayName = constraintDisplayName
	return constraintInfoBuilder
}

// Set the entity type of the seller group for the buyer entities
func (constraintInfoBuilder *ConstraintInfoBuilder) WithSellerType(providerType proto.EntityDTO_EntityType) *ConstraintInfoBuilder {
	constraintInfoBuilder.providerTypePtr = &providerType
	return constraintInfoBuilder
}

// Set the maximum number of buyer entities allowed in the policy
func (constraintInfoBuilder *ConstraintInfoBuilder) AtMostBuyers(maxBuyers int32) *ConstraintInfoBuilder {
	constraintInfoBuilder.maxBuyers = maxBuyers
	return constraintInfoBuilder
}

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

// Build the ConstraintInfo DTO
func (constraintInfoBuilder *ConstraintInfoBuilder) Build() (*proto.GroupDTO_ConstraintInfo, error) {
	constraintInfo := &proto.GroupDTO_ConstraintInfo{
		ConstraintId:   &constraintInfoBuilder.constraintId,
		ConstraintType: &constraintInfoBuilder.constraintType,
	}
	constraintInfo.ConstraintName = &constraintInfoBuilder.constraintName
	constraintInfo.ConstraintDisplayName = &constraintInfoBuilder.constraintDisplayName

	if constraintInfoBuilder.isBuyer {
		// buyer group specific metadata
		constraintInfo.BuyerMetaData = &proto.GroupDTO_BuyerMetaData{}
		bool := true
		constraintInfo.IsBuyer = &bool
		// seller entity type should be provided for buyer seller policies
		setProvider := (constraintInfoBuilder.constraintType == proto.GroupDTO_BUYER_SELLER_AFFINITY) ||
			(constraintInfoBuilder.constraintType == proto.GroupDTO_BUYER_SELLER_ANTI_AFFINITY)
		if setProvider && constraintInfoBuilder.providerTypePtr == nil {
			return nil, fmt.Errorf("Seller type required")
		}
		constraintInfo.BuyerMetaData.SellerType = constraintInfoBuilder.providerTypePtr
		// max buyers allowed
		if constraintInfoBuilder.maxBuyers > 0 {
			constraintInfo.BuyerMetaData.AtMost = &constraintInfoBuilder.maxBuyers
		}
	} else {
		// seller group specific metadata
		needsComplementary := constraintInfoBuilder.constraintType == proto.GroupDTO_BUYER_SELLER_ANTI_AFFINITY
		constraintInfo.NeedComplementary = &needsComplementary
	}

	return constraintInfo, nil
}

// -------------------------------------------------------------------------------------------------
// Group belonging to a Policy
type AbstractConstraintGroupBuilder struct {
	*AbstractBuilder
	*ConstraintInfoBuilder
}

// Build policy group with Constraint Info
func (groupBuilder *AbstractConstraintGroupBuilder) Build() (*proto.GroupDTO, error) {

	groupDTO, err := groupBuilder.AbstractBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("[AbstractConstraintGroupBuilder] Error building group")
	}

	var constraintInfo *proto.GroupDTO_ConstraintInfo
	constraintInfo, err = groupBuilder.ConstraintInfoBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("[AbstractConstraintGroupBuilder] Error building constraint")
	}

	// set constraint info in the group DTO
	info := &proto.GroupDTO_ConstraintInfo_{
		ConstraintInfo: constraintInfo,
	}

	groupDTO.Info = info

	return groupDTO, nil
}

// -------------------------------------------------------------------------------------------------

type buyerSellerPolicyData struct {
	policyId       string
	constraintType proto.GroupDTO_ConstraintType
	buyerData      *BuyerPolicyData
	sellerData     *SellerPolicyData
}

type buyerBuyerPolicyData struct {
	policyId       string
	constraintType proto.GroupDTO_ConstraintType
	buyerData      *BuyerPolicyData
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

func (placeTogether *PlaceTogetherPolicyBuilder) WithBuyers(buyers *BuyerPolicyData) *PlaceTogetherPolicyBuilder {
	placeTogether.buyerData = buyers
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

func (doNoPlace *DoNotPlaceTogetherPolicyBuilder) WithBuyers(buyers *BuyerPolicyData) *DoNotPlaceTogetherPolicyBuilder {
	doNoPlace.buyerData = buyers
	return doNoPlace
}

func (doNoPlace *DoNotPlaceTogetherPolicyBuilder) Build() ([]*proto.GroupDTO, error) {
	return buildBuyerBuyerPolicyGroup(doNoPlace.buyerBuyerPolicyData)
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
	buyerGroup, err := createPolicyBuyerGroup(policyData.policyId, policyData.constraintType, policyData.buyerData, policyData.sellerData)
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
	sellerGroup, err := createPolicySellerGroup(policyData.policyId, policyData.constraintType, policyData.sellerData)
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
func createPolicyBuyerGroup(policyId string, constraintType proto.GroupDTO_ConstraintType,
	buyerData *BuyerPolicyData, sellerData *SellerPolicyData) (*AbstractConstraintGroupBuilder, error) {

	var buyerGroup *AbstractBuilder
	if buyerData.entityTypePtr == nil {
		return nil, fmt.Errorf("Buyer entity type is not set")
	}
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
		return nil, fmt.Errorf("Buyer group member data missing")
	}

	// constraint info
	var constraintInfoBuilder *ConstraintInfoBuilder
	constraintInfoBuilder = buyerGroupConstraint(constraintType, policyId)
	if buyerData.atMost != 0 {
		constraintInfoBuilder.AtMostBuyers(buyerData.atMost)
	}
	if sellerData != nil && sellerData.entityTypePtr != nil {
		constraintInfoBuilder.WithSellerType(*sellerData.entityTypePtr)
	}

	buyerConstraintGroup := &AbstractConstraintGroupBuilder{
		AbstractBuilder:       buyerGroup,
		ConstraintInfoBuilder: constraintInfoBuilder,
	}

	return buyerConstraintGroup, nil
}

// Set up seller group for a policy
func createPolicySellerGroup(policyId string, constraintType proto.GroupDTO_ConstraintType,
	sellerData *SellerPolicyData) (*AbstractConstraintGroupBuilder, error) {

	var sellerGroup *AbstractBuilder
	if sellerData.entityTypePtr == nil {
		return nil, fmt.Errorf("Seller entity type is not set")
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
		return nil, fmt.Errorf("Seller group member data missing")
	}

	// constraint info
	var constraintInfoBuilder *ConstraintInfoBuilder
	constraintInfoBuilder = sellerGroupConstraint(constraintType, policyId)

	sellerConstraintGroup := &AbstractConstraintGroupBuilder{
		AbstractBuilder:       sellerGroup,
		ConstraintInfoBuilder: constraintInfoBuilder,
	}

	return sellerConstraintGroup, nil
}

func buildBuyerBuyerPolicyGroup(policyData *buyerBuyerPolicyData) ([]*proto.GroupDTO, error) {

	if policyData.buyerData == nil {
		return nil, fmt.Errorf("[buildBuyerBuyerPolicyGroup] Buyer group data not set")
	}

	var groupDTOs []*proto.GroupDTO

	// Buyer group and constraints
	buyerGroup, err := createPolicyBuyerGroup(policyData.policyId, policyData.constraintType, policyData.buyerData, nil)
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
