package group

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// ConstraintInfoBuilder is the builder for the constraint info data in a Group DTO
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

// Build the ConstraintInfo DTO
func (constraintInfoBuilder *ConstraintInfoBuilder) Build() (*proto.GroupDTO_ConstraintInfo, error) {
	constraintInfo := &proto.GroupDTO_ConstraintInfo{
		ConstraintId:          &constraintInfoBuilder.constraintId,
		ConstraintType:        &constraintInfoBuilder.constraintType,
		ConstraintName:        &constraintInfoBuilder.constraintName,
		ConstraintDisplayName: &constraintInfoBuilder.constraintDisplayName,
	}

	if constraintInfoBuilder.isBuyer {
		// buyer group specific metadata
		constraintInfo.BuyerMetaData = &proto.GroupDTO_BuyerMetaData{}
		boolVal := true
		constraintInfo.IsBuyer = &boolVal
		// seller entity type should be provided for buyer seller policies
		setProvider := (constraintInfoBuilder.constraintType == proto.GroupDTO_BUYER_SELLER_AFFINITY) ||
			(constraintInfoBuilder.constraintType == proto.GroupDTO_BUYER_SELLER_ANTI_AFFINITY)
		if setProvider && constraintInfoBuilder.providerTypePtr == nil {
			return nil, fmt.Errorf("seller type required")
		}
		constraintInfo.BuyerMetaData.SellerType = constraintInfoBuilder.providerTypePtr
		// max buyers allowed
		if constraintInfoBuilder.maxBuyers > 0 {
			constraintInfo.BuyerMetaData.AtMost = &constraintInfoBuilder.maxBuyers
		}
	}

	// seller group specific metadata
	needsComplementary := constraintInfoBuilder.constraintType == proto.GroupDTO_BUYER_SELLER_ANTI_AFFINITY
	constraintInfo.NeedComplementary = &needsComplementary

	return constraintInfo, nil
}

// AbstractConstraintGroupBuilder is the builder of a Group with constraint
type AbstractConstraintGroupBuilder struct {
	*AbstractBuilder
	*ConstraintInfoBuilder
}

// Build policy group with Constraint Info
func (groupBuilder *AbstractConstraintGroupBuilder) Build() (*proto.GroupDTO, error) {

	groupDTO, err := groupBuilder.AbstractBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build group: %v", err)
	}

	var constraintInfo *proto.GroupDTO_ConstraintInfo
	constraintInfo, err = groupBuilder.ConstraintInfoBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build group constraint: %v", err)
	}

	if groupBuilder.constraintType == proto.GroupDTO_CLUSTER {
		entityType := *groupBuilder.entityTypePtr
		if entityType != proto.EntityDTO_VIRTUAL_MACHINE &&
			entityType != proto.EntityDTO_PHYSICAL_MACHINE &&
			entityType != proto.EntityDTO_STORAGE {
			return nil, fmt.Errorf("failed to build cluster: unsupported entity type %v", entityType)
		}
	}

	// set constraint info in the group DTO
	info := &proto.GroupDTO_ConstraintInfo_{
		ConstraintInfo: constraintInfo,
	}

	groupDTO.Info = info

	return groupDTO, nil
}
