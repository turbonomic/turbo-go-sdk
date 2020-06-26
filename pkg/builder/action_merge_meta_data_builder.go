package builder

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type ActionAggregationTargetBuilder struct {
	entityType        proto.EntityDTO_EntityType
	relatedEntityType proto.EntityDTO_EntityType
	relatedBy         proto.ConnectedEntity_ConnectionType
	//mergeType proto.ActionMergeTargetData_MergeType
}

func NewActionAggregationTargetBuilder(entityType proto.EntityDTO_EntityType,
	relatedEntityType proto.EntityDTO_EntityType,
	relatedBy proto.ConnectedEntity_ConnectionType) *ActionAggregationTargetBuilder {
	return &ActionAggregationTargetBuilder{
		entityType:        entityType,
		relatedEntityType: relatedEntityType,
		relatedBy:         relatedBy,
		//mergeType: proto.ActionMergeTargetData_PASS_THROUGH,
	}
}

func (builder *ActionAggregationTargetBuilder) Create() *proto.ActionMergeTargetData {
	target := &proto.ActionMergeTargetData{
		//EntityType: &builder.entityType,
		RelatedTo:  &builder.relatedEntityType,
		RelatedBy: &proto.ActionMergeTargetData_EntityRelationship{
			EntityRelationship: &proto.ActionMergeTargetData_EntityRelationship_ConnectionType{
				ConnectionType: builder.relatedBy,
			},
		},
		//MergeType: &builder.mergeType,
	}
	return target
}

type ActionDeDuplicateAndAggregationTargetBuilder struct {
	deDuplicationTarget *ActionAggregationTargetBuilder
	aggregationTarget   *ActionAggregationTargetBuilder
}

func NewActionDeDuplicateAndAggregationTargetBuilder() *ActionDeDuplicateAndAggregationTargetBuilder {
	return &ActionDeDuplicateAndAggregationTargetBuilder{}
}

func (builder *ActionDeDuplicateAndAggregationTargetBuilder) DeDuplicatedBy(
	deDuplicationTarget *ActionAggregationTargetBuilder) *ActionDeDuplicateAndAggregationTargetBuilder {
	//deDuplicationTarget.mergeType = proto.ActionMergeTargetData_DE_DUPLICATE
	builder.deDuplicationTarget = deDuplicationTarget
	return builder
}

func (builder *ActionDeDuplicateAndAggregationTargetBuilder) AggregatedBy(
	aggregationTarget *ActionAggregationTargetBuilder) *ActionDeDuplicateAndAggregationTargetBuilder {
	//aggregationTarget.mergeType = proto.ActionMergeTargetData_AGGREGATE
	builder.aggregationTarget = aggregationTarget
	return builder
}
func (builder *ActionDeDuplicateAndAggregationTargetBuilder) Create() *proto.ChainedActionMergeTargetData {

	chainedMergeTarget := &proto.ChainedActionMergeTargetData{}
	true_flag := true
	deDuplicationTargetLink := &proto.ChainedActionMergeTargetData_TargetDataLink{
		MergeTarget:    builder.deDuplicationTarget.Create(),
		DeDuplicate:    &true_flag,
	}
	chainedMergeTarget.TargetLinks = append(chainedMergeTarget.TargetLinks, deDuplicationTargetLink)

	false_flag := false
	aggregationTargetLink := &proto.ChainedActionMergeTargetData_TargetDataLink{
		MergeTarget:    builder.aggregationTarget.Create(),
		DeDuplicate:    &false_flag,
	}
	chainedMergeTarget.TargetLinks = append(chainedMergeTarget.TargetLinks, aggregationTargetLink)

	return chainedMergeTarget
}

type ResizeMergePolicyBuilder struct {
	entityType                *proto.EntityDTO_EntityType
	aggregationTargets        []*ActionAggregationTargetBuilder
	chainedAggregationTargets []*ActionDeDuplicateAndAggregationTargetBuilder
	commTypes                 []*CommodityMergeData
}

type CommodityMergeData struct {
	commType    proto.CommodityDTO_CommodityType
	changedAttr proto.ActionItemDTO_CommodityAttribute
}

func NewResizeMergeSpecBuilder() *ResizeMergePolicyBuilder {
	return &ResizeMergePolicyBuilder{
		//entityType: &entityType,
	}
}

func (rb *ResizeMergePolicyBuilder) ForEntityType(entityType proto.EntityDTO_EntityType) *ResizeMergePolicyBuilder {
	rb.entityType = &entityType
	return rb
}

func (rb *ResizeMergePolicyBuilder) AggregateBy(mergeTarget *ActionAggregationTargetBuilder) *ResizeMergePolicyBuilder {
	rb.aggregationTargets = append(rb.aggregationTargets, mergeTarget)
	return rb
}

func (rb *ResizeMergePolicyBuilder) DeDuplicateAndAggregateBy(mergeTarget *ActionDeDuplicateAndAggregationTargetBuilder) *ResizeMergePolicyBuilder {
	rb.chainedAggregationTargets = append(rb.chainedAggregationTargets, mergeTarget)
	return rb
}

func (rb *ResizeMergePolicyBuilder) ForCommodity(commType proto.CommodityDTO_CommodityType) *ResizeMergePolicyBuilder {
	comm := &CommodityMergeData{
		commType: commType,
	}
	rb.commTypes = append(rb.commTypes, comm)
	return rb
}

func (rb *ResizeMergePolicyBuilder) ForCommodityAndAttribute(commType proto.CommodityDTO_CommodityType,
	changedAttr proto.ActionItemDTO_CommodityAttribute) *ResizeMergePolicyBuilder {
	comm := &CommodityMergeData{
		commType:    commType,
		changedAttr: changedAttr,
	}
	rb.commTypes = append(rb.commTypes, comm)
	return rb
}

func (rb *ResizeMergePolicyBuilder) Build() (*proto.ActionMergePolicyDTO, error) {
	if rb.entityType == nil {
		return nil, fmt.Errorf("Entity type required for action merge policy")
	}

	if len(rb.aggregationTargets) == 0 && len(rb.chainedAggregationTargets) == 0 {
		return nil, fmt.Errorf("Target type required for action merge policy")
	}

	if len(rb.commTypes) == 0 {
		return nil, fmt.Errorf("Commodity types required for resize merge policy")
	}

	commMergeDataList := []*proto.ResizeMergeSpec_CommodityMergeData{}
	for _, commData := range rb.commTypes {
		commMergeData := &proto.ResizeMergeSpec_CommodityMergeData{
			CommodityType: &commData.commType,
			ChangedAttr:   &commData.changedAttr,
		}
		commMergeDataList = append(commMergeDataList, commMergeData)
	}
	resizeSpec := &proto.ResizeMergeSpec{
		CommodityData: commMergeDataList,
	}

	mergeSpec := &proto.ActionMergePolicyDTO{
		EntityType: rb.entityType,

		ActionSpec: &proto.ActionMergePolicyDTO_ResizeSpec{
			ResizeSpec: resizeSpec,
		},
	}

	var executionTargetList []*proto.ActionMergeExecutionTarget
	for _, targetData := range rb.aggregationTargets {
		executionTarget := &proto.ActionMergeExecutionTarget{
			ExecutionTarget: &proto.ActionMergeExecutionTarget_MergeTarget{
				MergeTarget: targetData.Create(),
			},
		}
		executionTargetList = append(executionTargetList, executionTarget)
	}

	for _, targetData := range rb.chainedAggregationTargets {
		chainedTarget := targetData.Create()
		fmt.Printf("chainedTarget: %++v\n", chainedTarget)
		executionTarget := &proto.ActionMergeExecutionTarget{
			ExecutionTarget: &proto.ActionMergeExecutionTarget_ChainedMergeTarget{
				ChainedMergeTarget: chainedTarget,
			},
		}
		executionTargetList = append(executionTargetList, executionTarget)
	}

	mergeSpec.ExecutionTargets = executionTargetList
	return mergeSpec, nil
}
