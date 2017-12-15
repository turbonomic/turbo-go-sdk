package builder

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	DEFAULT_FULL_DISCOVERY_IN_SECS int32 = 600
	DISCOVERY_NOT_SUPPORTED        int32 = -1
)

//// Helper methods to create AccountDefinition map for sub classes of the probe
// An AccountDefEntryBuilder builds an AccountDefEntry instance.
type AccountDefEntryBuilder struct {
	accountDefEntry *proto.AccountDefEntry
}

func NewAccountDefEntryBuilder(name, displayName, description, verificationRegex string,
	mandatory bool, isSecret bool) *AccountDefEntryBuilder {
	fieldType := &proto.CustomAccountDefEntry_PrimitiveValue_{
		PrimitiveValue: proto.CustomAccountDefEntry_STRING,
	}
	entry := &proto.CustomAccountDefEntry{
		Name:              &name,
		DisplayName:       &displayName,
		Description:       &description,
		VerificationRegex: &verificationRegex,
		IsSecret:          &isSecret,
		FieldType:         fieldType,
	}

	customDef := &proto.AccountDefEntry_CustomDefinition{
		CustomDefinition: entry,
	}

	accountDefEntry := &proto.AccountDefEntry{
		Mandatory:  &mandatory,
		Definition: customDef,
	}

	return &AccountDefEntryBuilder{
		accountDefEntry: accountDefEntry,
	}
}

func (builder *AccountDefEntryBuilder) Create() *proto.AccountDefEntry {
	return builder.accountDefEntry
}

// Action Policy Metadata
type ActionPolicyBuilder struct {
	ActionPolicyMap map[proto.EntityDTO_EntityType]map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability
}

func NewActionPolicyBuilder() *ActionPolicyBuilder {
	return &ActionPolicyBuilder{
		ActionPolicyMap: make(map[proto.EntityDTO_EntityType]map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability),
	}
}

func (builder *ActionPolicyBuilder) WithEntityActions(entityType proto.EntityDTO_EntityType,
	actionType proto.ActionItemDTO_ActionType,
	actionCapability proto.ActionPolicyDTO_ActionCapability) *ActionPolicyBuilder {

	_, exists := builder.ActionPolicyMap[entityType]
	if !exists {
		builder.ActionPolicyMap[entityType] =
			make(map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability)
	}
	entityPolicies, _ := builder.ActionPolicyMap[entityType]
	entityPolicies[actionType] = actionCapability

	return builder
}

func (builder *ActionPolicyBuilder) Create() []*proto.ActionPolicyDTO {
	var policies []*proto.ActionPolicyDTO

	for entityType, entityPolicies := range builder.ActionPolicyMap {
		policyElements := []*proto.ActionPolicyDTO_ActionPolicyElement{}

		for key, val := range entityPolicies {
			actionType := key
			actionCapability := val
			actionPolicy := &proto.ActionPolicyDTO_ActionPolicyElement{
				ActionType:       &actionType,
				ActionCapability: &actionCapability,
			}

			policyElements = append(policyElements, actionPolicy)
		}
		eType := entityType
		policyDto := &proto.ActionPolicyDTO{
			EntityType:    &eType,
			PolicyElement: policyElements,
		}

		policies = append(policies, policyDto)
	}
	return policies
}

// A ProbeInfoBuilder builds a ProbeInfo instance.
type ProbeInfoBuilder struct {
	probeInfo *proto.ProbeInfo
}

// NewProbeInfoBuilder builds the ProbeInfo DTO for the given probe
func NewProbeInfoBuilder(probeType, probeCat string,
	supplyChainSet []*proto.TemplateDTO,
	acctDef []*proto.AccountDefEntry) *ProbeInfoBuilder {
	// New ProbeInfo protobuf with this input
	probeInfo := &proto.ProbeInfo{
		ProbeType:                &probeType,
		ProbeCategory:            &probeCat,
		SupplyChainDefinitionSet: supplyChainSet,
		AccountDefinition:        acctDef,
	}
	return &ProbeInfoBuilder{
		probeInfo: probeInfo,
	}
}

// NewBasicProbeInfoBuilder builds the ProbeInfo DTO for the given probe
func NewBasicProbeInfoBuilder(probeType, probeCat string) *ProbeInfoBuilder {
	var full, other int32
	full = DEFAULT_FULL_DISCOVERY_IN_SECS
	other = DISCOVERY_NOT_SUPPORTED

	probeInfo := &proto.ProbeInfo{
		ProbeType:                             &probeType,
		ProbeCategory:                         &probeCat,
		FullRediscoveryIntervalSeconds:        &full,
		IncrementalRediscoveryIntervalSeconds: &other,
		PerformanceRediscoveryIntervalSeconds: &other,
	}
	return &ProbeInfoBuilder{
		probeInfo: probeInfo,
	}
}

func (builder *ProbeInfoBuilder) WithIdentifyingField(idField string) *ProbeInfoBuilder {
	builder.probeInfo.TargetIdentifierField = append(builder.probeInfo.TargetIdentifierField, idField)
	return builder
}

func (builder *ProbeInfoBuilder) WithSupplyChain(supplyChainSet []*proto.TemplateDTO) *ProbeInfoBuilder {
	builder.probeInfo.SupplyChainDefinitionSet = supplyChainSet
	return builder
}

func (builder *ProbeInfoBuilder) WithAccountDefinition(acctDefSet []*proto.AccountDefEntry) *ProbeInfoBuilder {
	builder.probeInfo.AccountDefinition = acctDefSet
	return builder
}

func (builder *ProbeInfoBuilder) WithFullDiscoveryInterval(fullDiscoveryInSecs int32) *ProbeInfoBuilder {
	builder.probeInfo.FullRediscoveryIntervalSeconds = &fullDiscoveryInSecs
	return builder
}

func (builder *ProbeInfoBuilder) WithIncrementalDiscoveryInterval(incrementalDiscoveryInSecs int32) *ProbeInfoBuilder {
	builder.probeInfo.IncrementalRediscoveryIntervalSeconds = &incrementalDiscoveryInSecs
	return builder
}

func (builder *ProbeInfoBuilder) WithPerformanceDiscoveryInterval(performanceDiscoveryInSecs int32) *ProbeInfoBuilder {
	builder.probeInfo.PerformanceRediscoveryIntervalSeconds = &performanceDiscoveryInSecs
	return builder
}

func (builder *ProbeInfoBuilder) WithActionPolicySet(actionPolicySet []*proto.ActionPolicyDTO) *ProbeInfoBuilder {
	builder.probeInfo.ActionPolicy = actionPolicySet
	return builder
}

func (builder *ProbeInfoBuilder) WithEntityMetadata(entityMetadataSet []*proto.EntityIdentityMetadata) *ProbeInfoBuilder {
	builder.probeInfo.EntityMetadata = entityMetadataSet
	return builder
}

func (builder *ProbeInfoBuilder) Create() *proto.ProbeInfo {
	return builder.probeInfo
}
