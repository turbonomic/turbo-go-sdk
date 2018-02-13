package builder

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

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

// Set up the action policies for the specified entity type.
// Takes the EntityActionPolicyBuilder as an argument to conveniently set up the policies for
// different action types for this entity type.
func (builder *ActionPolicyBuilder) ForEntity(entityActionBuilder *EntityActionPolicyBuilder) *ActionPolicyBuilder {

	builder.ActionPolicyMap[entityActionBuilder.EntityType] = entityActionBuilder.Build()

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

// Builder for action policies for an entity type
type EntityActionPolicyBuilder struct {
	EntityType      proto.EntityDTO_EntityType
	ActionPolicyMap map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability
}

// Create a new instance for an entity type to set up the action policies
func NewEntityActionPolicyBuilder(entityType proto.EntityDTO_EntityType) *EntityActionPolicyBuilder {
	return &EntityActionPolicyBuilder{
		EntityType:      entityType,
		ActionPolicyMap: make(map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability),
	}
}

// Lists the action types that can be supported and executed for the entity type
func (builder *EntityActionPolicyBuilder) Supports(actionTypes ...proto.ActionItemDTO_ActionType) *EntityActionPolicyBuilder {
	for _, actionType := range actionTypes {
		builder.ActionPolicyMap[actionType] = proto.ActionPolicyDTO_SUPPORTED
	}
	return builder
}

// Lists the action types that cannot be supported and executed for the entity type
func (builder *EntityActionPolicyBuilder) DoesNotSupport(actionTypes ...proto.ActionItemDTO_ActionType) *EntityActionPolicyBuilder {
	for _, actionType := range actionTypes {
		builder.ActionPolicyMap[actionType] = proto.ActionPolicyDTO_NOT_SUPPORTED
	}
	return builder
}

// Lists the action types that cannot be executed currently for the entity type
func (builder *EntityActionPolicyBuilder) RecommendOnly(actionTypes ...proto.ActionItemDTO_ActionType) *EntityActionPolicyBuilder {
	for _, actionType := range actionTypes {
		builder.ActionPolicyMap[actionType] = proto.ActionPolicyDTO_NOT_EXECUTABLE
	}
	return builder
}

// Specifies if the entity can be moved or not.
func (builder *EntityActionPolicyBuilder) CanMove(move bool) *EntityActionPolicyBuilder {
	if move {
		builder.ActionPolicyMap[proto.ActionItemDTO_MOVE] = proto.ActionPolicyDTO_SUPPORTED
	} else {
		builder.ActionPolicyMap[proto.ActionItemDTO_MOVE] = proto.ActionPolicyDTO_NOT_SUPPORTED
	}
	return builder
}

// Specifies if the entity can be resized or not.
func (builder *EntityActionPolicyBuilder) CanResize(resize bool) *EntityActionPolicyBuilder {
	if resize {
		builder.ActionPolicyMap[proto.ActionItemDTO_RESIZE] = proto.ActionPolicyDTO_SUPPORTED
	} else {
		builder.ActionPolicyMap[proto.ActionItemDTO_RESIZE] = proto.ActionPolicyDTO_NOT_SUPPORTED
	}
	return builder
}

// Specifies if the entity can be cloned or not.
func (builder *EntityActionPolicyBuilder) CanClone(clone bool) *EntityActionPolicyBuilder {
	if clone {
		builder.ActionPolicyMap[proto.ActionItemDTO_PROVISION] = proto.ActionPolicyDTO_SUPPORTED
	} else {
		builder.ActionPolicyMap[proto.ActionItemDTO_PROVISION] = proto.ActionPolicyDTO_NOT_SUPPORTED
	}
	return builder
}

func (builder *EntityActionPolicyBuilder) Build() map[proto.ActionItemDTO_ActionType]proto.ActionPolicyDTO_ActionCapability {
	return builder.ActionPolicyMap
}

// A ProbeInfoBuilder builds a ProbeInfo instance.
// ProbeInfo structure stores the data necessary to register the Probe with the Turbonomic server.
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

	probeInfo := &proto.ProbeInfo{
		ProbeType:     &probeType,
		ProbeCategory: &probeCat,
	}
	return &ProbeInfoBuilder{
		probeInfo: probeInfo,
	}
}

// Return the instance of the ProbeInfo DTO created by the builder
func (builder *ProbeInfoBuilder) Create() *proto.ProbeInfo {
	checkFullDiscoveryInterval(builder.probeInfo)
	return builder.probeInfo
}

// Set the field name whose value is used to uniquely identify the target for this probe
func (builder *ProbeInfoBuilder) WithIdentifyingField(idField string) *ProbeInfoBuilder {
	builder.probeInfo.TargetIdentifierField = append(builder.probeInfo.TargetIdentifierField,
		idField)
	return builder
}

// Set the supply chain for the probe
func (builder *ProbeInfoBuilder) WithSupplyChain(supplyChainSet []*proto.TemplateDTO,
) *ProbeInfoBuilder {
	builder.probeInfo.SupplyChainDefinitionSet = supplyChainSet
	return builder
}

// Set the account definition for creating targets for this probe
func (builder *ProbeInfoBuilder) WithAccountDefinition(acctDefSet []*proto.AccountDefEntry,
) *ProbeInfoBuilder {
	builder.probeInfo.AccountDefinition = acctDefSet
	return builder
}

// Set the interval in seconds for running the full discovery of the probe
func (builder *ProbeInfoBuilder) WithFullDiscoveryInterval(fullDiscoveryInSecs int32,
) *ProbeInfoBuilder {
	// Ignore if the interval is less than DEFAULT_MIN_DISCOVERY_IN_SECS
	if fullDiscoveryInSecs < pkg.DEFAULT_MIN_DISCOVERY_IN_SECS {
		return builder
	}
	builder.probeInfo.FullRediscoveryIntervalSeconds = &fullDiscoveryInSecs
	return builder
}

// Set the interval in seconds for executing the incremental discovery of the probe
func (builder *ProbeInfoBuilder) WithIncrementalDiscoveryInterval(incrementalDiscoveryInSecs int32,
) *ProbeInfoBuilder {
	// Ignore if the interval implies the DISCOVERY_NOT_SUPPORTED value
	if incrementalDiscoveryInSecs <= pkg.DISCOVERY_NOT_SUPPORTED {
		return builder
	}
	builder.probeInfo.IncrementalRediscoveryIntervalSeconds = &incrementalDiscoveryInSecs
	return builder
}

// Set the interval in seconds for executing the performance or metrics discovery of the probe
func (builder *ProbeInfoBuilder) WithPerformanceDiscoveryInterval(performanceDiscoveryInSecs int32,
) *ProbeInfoBuilder {
	// Ignore if the interval implies the DISCOVERY_NOT_SUPPORTED value
	if performanceDiscoveryInSecs <= pkg.DISCOVERY_NOT_SUPPORTED {
		return builder
	}
	builder.probeInfo.PerformanceRediscoveryIntervalSeconds = &performanceDiscoveryInSecs
	return builder
}

func (builder *ProbeInfoBuilder) WithActionPolicySet(actionPolicySet []*proto.ActionPolicyDTO,
) *ProbeInfoBuilder {
	builder.probeInfo.ActionPolicy = actionPolicySet
	return builder
}

func (builder *ProbeInfoBuilder) WithEntityMetadata(entityMetadataSet []*proto.EntityIdentityMetadata,
) *ProbeInfoBuilder {
	builder.probeInfo.EntityMetadata = entityMetadataSet
	return builder
}

// Assert that the full discovery interval is set
func checkFullDiscoveryInterval(probeInfo *proto.ProbeInfo) {
	var interval int32
	interval = pkg.DEFAULT_MIN_DISCOVERY_IN_SECS

	defaultFullDiscoveryIntervalMessage := func() {
		glog.V(2).Infof("No rediscovery interval specified. "+
			"	Using a default value of %d seconds", interval)
	}

	if (probeInfo.FullRediscoveryIntervalSeconds == nil) ||
		(*probeInfo.FullRediscoveryIntervalSeconds <= 0) {
		probeInfo.FullRediscoveryIntervalSeconds = &interval
		defaultFullDiscoveryIntervalMessage()
	}
	//if *probeInfo.FullRediscoveryIntervalSeconds <= 0 {
	//	probeInfo.FullRediscoveryIntervalSeconds = &interval
	//	defaultFullDiscoveryIntervalMessage()
	//}
}
