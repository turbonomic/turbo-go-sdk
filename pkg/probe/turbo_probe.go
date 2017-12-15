package probe

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"github.com/golang/glog"
)

// Turbo Probe Abstraction
// Consists of clients that handle probe registration metadata,
// and the discovery and action execution for different probe targets
type TurboProbe struct {
	ProbeConfiguration *ProbeConfig
	RegistrationClient *ProbeRegistrationAgent
	DiscoveryClientMap map[string]*TargetDiscoveryAgent

	ActionClient TurboActionExecutorClient
}

type ProbeRegistrationAgent struct {
	ISupplyChainProvider
	IAccountDefinitionProvider
	IActionPolicyProvider
	IEntityMetadataProvider
	discoveryMetadata *DiscoveryMetadata
}

type TurboRegistrationClient interface {
	ISupplyChainProvider
	IAccountDefinitionProvider
}

func NewProbeRegistrator() *ProbeRegistrationAgent {
	registrator := &ProbeRegistrationAgent{}
	registrator.discoveryMetadata = NewDiscoveryMetadata()
	return registrator
}

type TargetDiscoveryAgent struct {
	TargetId string
	TurboDiscoveryClient
	IIncrementalDiscovery
	IPerformanceDiscovery
}

func NewTargetDiscoveryAgent(targetId string) *TargetDiscoveryAgent {
	targetAgent := &TargetDiscoveryAgent{TargetId: targetId}
	return targetAgent
}

type TurboDiscoveryClient interface {
	// Discovers the target, creating an EntityDTO representation of every object
	// discovered by the  probe.
	// @param accountValues object, holding all the account values
	// @return discovery response, consisting of both entities and errors, if any
	Discover(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error)
	Validate(accountValues []*proto.AccountValue) (*proto.ValidationResponse, error)
	GetAccountValues() *TurboTargetInfo
}

type ProbeConfig struct {
	ProbeType            string
	ProbeCategory        string
	FullDiscovery        int32
	IncrementalDiscovery int32
	PerformanceDiscovery int32
}

// ===========================================    New Probe ==========================================================

func newTurboProbe(probeConf *ProbeConfig) (*TurboProbe, error) {
	if probeConf.ProbeType == "" {
		return nil, ErrorInvalidProbeType()
	}

	if probeConf.ProbeCategory == "" {
		return nil, ErrorInvalidProbeCategory()
	}

	myProbe := &TurboProbe{
		ProbeConfiguration: probeConf,
		DiscoveryClientMap: make(map[string]*TargetDiscoveryAgent),
	}

	// registration client defaults
	registrationClient := NewProbeRegistrator()
	registrationClient.discoveryMetadata.SetFullRediscoveryIntervalSeconds(probeConf.FullDiscovery)
	registrationClient.discoveryMetadata.SetIncrementalRediscoveryIntervalSeconds(probeConf.IncrementalDiscovery)
	registrationClient.discoveryMetadata.SetPerformanceRediscoveryIntervalSeconds(probeConf.PerformanceDiscovery)

	myProbe.RegistrationClient = registrationClient

	glog.V(2).Infof("[NewTurboProbe] Created TurboProbe: %s", myProbe)
	return myProbe, nil
}

func (theProbe *TurboProbe) getDiscoveryClient(targetIdentifier string) TurboDiscoveryClient {
	target, exists := theProbe.DiscoveryClientMap[targetIdentifier]

	if !exists {
		glog.Errorf("[GetTurboDiscoveryClient] Cannot find Target for address: %s", targetIdentifier)
		return nil
	}
	return target.TurboDiscoveryClient
}

// TODO: this method should be synchronized
func (theProbe *TurboProbe) GetTurboDiscoveryClient(accountValues []*proto.AccountValue) TurboDiscoveryClient {
	var address string
	identifyingField := theProbe.RegistrationClient.GetIdentifyingFields()

	address = findTargetId(accountValues, identifyingField)
	target, exists := theProbe.DiscoveryClientMap[address]

	if !exists {
		glog.Errorf("[GetTurboDiscoveryClient] Cannot find Target for address: %s", address)
		//TODO: CreateDiscoveryClient(address, accountValues, )
		return nil
	}
	glog.V(2).Infof("[GetTurboDiscoveryClient] Found Target for address: %s", address)
	return target.TurboDiscoveryClient
}

func findTargetId(accountValues []*proto.AccountValue, identifyingField string) string {
	var address string

	for _, accVal := range accountValues {
		if accVal.GetKey() == identifyingField {
			address = accVal.GetStringValue()
			return address
		}
	}
	return address
}

func (theProbe *TurboProbe) DiscoverTarget(accountValues []*proto.AccountValue) *proto.DiscoveryResponse {
	glog.V(2).Infof("Discover Target: %s", accountValues)
	var handler TurboDiscoveryClient
	handler = theProbe.GetTurboDiscoveryClient(accountValues)
	if handler == nil {
		description := "Non existent target"
		return theProbe.createDiscoveryErrorDTO(description, proto.ErrorDTO_CRITICAL)
	}
	var discoveryResponse *proto.DiscoveryResponse

	glog.V(4).Infof("Send discovery request to handler %v", handler)
	discoveryResponse, err := handler.Discover(accountValues)

	if err != nil {
		description := fmt.Sprintf("Error discovering target %s", err)
		severity := proto.ErrorDTO_CRITICAL

		discoveryResponse = theProbe.createDiscoveryErrorDTO(description, severity)
		glog.Errorf("Error discovering target %s", discoveryResponse)
	}
	glog.V(3).Infof("Discovery response: %s", discoveryResponse)
	return discoveryResponse
}

func (theProbe *TurboProbe) ValidateTarget(accountValues []*proto.AccountValue) *proto.ValidationResponse {
	glog.V(2).Infof("Validate Target: %++v", accountValues)
	var handler TurboDiscoveryClient
	handler = theProbe.GetTurboDiscoveryClient(accountValues)
	if handler == nil {
		description := "Target not found"
		severity := proto.ErrorDTO_CRITICAL
		return theProbe.createValidationErrorDTO(description, severity)
	}

	var validationResponse *proto.ValidationResponse
	glog.V(3).Infof("Send validation request to handler %s", handler)
	validationResponse, err := handler.Validate(accountValues)

	// TODO: if the handler is nil, implies the target is added from the UI
	// Create a new discovery client for this target and add it to the map of discovery clients
	// allow to pass a func to instantiate a default discovery client
	if err != nil {
		description := fmt.Sprintf("Error validating target %s", err)
		severity := proto.ErrorDTO_CRITICAL

		validationResponse = theProbe.createValidationErrorDTO(description, severity)
		glog.Errorf("Error validating target %s", validationResponse)
	}
	glog.V(3).Infof("Validation response: %s", validationResponse)
	return validationResponse
}

func (theProbe *TurboProbe) DiscoverTargetIncremental(accountValues []*proto.AccountValue) *proto.DiscoveryResponse {
	glog.V(2).Infof("Incremental discovery for Target: %s", accountValues)
	targetId := findTargetId(accountValues, theProbe.RegistrationClient.GetIdentifyingFields())
	var handler IIncrementalDiscovery
	target, exists := theProbe.DiscoveryClientMap[targetId]

	if !exists {
		description := "Non existent target"
		return theProbe.createDiscoveryErrorDTO(description, proto.ErrorDTO_CRITICAL)
	}

	handler = target.IIncrementalDiscovery
	if handler == nil {
		description := "Incremental discovery not supported"
		return theProbe.createDiscoveryErrorDTO(description, proto.ErrorDTO_WARNING)
	}
	var discoveryResponse *proto.DiscoveryResponse

	glog.V(4).Infof("Send incremental discovery request to handler %v", handler)
	discoveryResponse, err := handler.DiscoverIncremental(accountValues)

	if err != nil {
		description := fmt.Sprintf("Error during incremental discovering target %s", err)
		severity := proto.ErrorDTO_CRITICAL

		discoveryResponse = theProbe.createDiscoveryErrorDTO(description, severity)
		glog.Errorf("Error during incremental discovery of target %s", discoveryResponse)
	}
	glog.V(3).Infof("Incremental Discovery response: %s", discoveryResponse)
	return discoveryResponse
}

func (theProbe *TurboProbe) DiscoverTargetPerformance(accountValues []*proto.AccountValue) *proto.DiscoveryResponse {
	glog.V(2).Infof("Performance discovery for Target: %s", accountValues)
	targetId := findTargetId(accountValues, theProbe.RegistrationClient.GetIdentifyingFields())
	var handler IPerformanceDiscovery
	target, exists := theProbe.DiscoveryClientMap[targetId]

	if !exists {
		description := "Non existent target"
		return theProbe.createDiscoveryErrorDTO(description, proto.ErrorDTO_CRITICAL)
	}

	handler = target.IPerformanceDiscovery
	if handler == nil {
		description := "Performance discovery not supported"
		return theProbe.createDiscoveryErrorDTO(description, proto.ErrorDTO_WARNING)
	}
	var discoveryResponse *proto.DiscoveryResponse

	glog.V(4).Infof("Send performance discovery request to handler %v", handler)
	discoveryResponse, err := handler.DiscoverPerformance(accountValues)

	if err != nil {
		description := fmt.Sprintf("Error during performance discovery of target %s", err)
		severity := proto.ErrorDTO_CRITICAL

		discoveryResponse = theProbe.createDiscoveryErrorDTO(description, severity)
		glog.Errorf("Error during performance discovery of target %s", discoveryResponse)
	}
	glog.V(3).Infof("Performance Discovery response: %s", discoveryResponse)
	return discoveryResponse
}

func (theProbe *TurboProbe) ExecuteAction(actionExecutionDTO *proto.ActionExecutionDTO, accountValues []*proto.AccountValue,
	progressTracker ActionProgressTracker) *proto.ActionResult {
	if theProbe.ActionClient == nil {
		glog.V(3).Infof("ActionClient not defined for Probe %s", theProbe.ProbeConfiguration.ProbeType)
		return theProbe.createActionErrorDTO("ActionClient not defined for Probe " + theProbe.ProbeConfiguration.ProbeType)
	}
	glog.V(3).Infof("Execute Action for Target: %s", accountValues)
	response, err := theProbe.ActionClient.ExecuteAction(actionExecutionDTO, accountValues, progressTracker)

	if err != nil {
		description := fmt.Sprintf("Error executing action %s", err)
		glog.Errorf("Error executing action")
		return theProbe.createActionErrorDTO(description)
	}
	return response
}

// ==============================================================================================================
// The Targets associated with this probe type
func (theProbe *TurboProbe) GetProbeTargets() []*TurboTargetInfo {
	// Iterate over the discovery client map and send requests to the server
	var targets []*TurboTargetInfo
	for targetId, discoveryClient := range theProbe.DiscoveryClientMap {

		targetInfo := discoveryClient.GetAccountValues()
		targetInfo.targetType = theProbe.ProbeConfiguration.ProbeType
		targetInfo.targetIdentifierField = targetId

		targets = append(targets, targetInfo)
	}
	return targets
}

// The ProbeInfo for the probe
func (theProbe *TurboProbe) GetProbeInfo() (*proto.ProbeInfo, error) {
	// 1. construct the basic probe info.
	probeCat := theProbe.ProbeConfiguration.ProbeCategory
	probeType := theProbe.ProbeConfiguration.ProbeType
	probeInfoBuilder := builder.NewBasicProbeInfoBuilder(probeType, probeCat)

	registrationClient := theProbe.RegistrationClient
	// 2. Get the supply chain.
	if registrationClient.ISupplyChainProvider != nil {
		probeInfoBuilder.WithSupplyChain(registrationClient.GetSupplyChainDefinition())
	}

	// 3. Get the account definition
	if registrationClient.IAccountDefinitionProvider != nil {
		probeInfoBuilder.WithAccountDefinition(registrationClient.GetAccountDefinition())
	}

	// 4. Fields that serve to uniquely identify a target
	probeInfoBuilder = probeInfoBuilder.WithIdentifyingField(registrationClient.GetIdentifyingFields())

	// 5. discovery intervals metadata
	probeInfoBuilder.WithFullDiscoveryInterval(registrationClient.discoveryMetadata.GetFullRediscoveryIntervalSeconds())
	probeInfoBuilder.WithIncrementalDiscoveryInterval(registrationClient.discoveryMetadata.GetIncrementalRediscoveryIntervalSeconds())
	probeInfoBuilder.WithPerformanceDiscoveryInterval(registrationClient.discoveryMetadata.GetPerformanceRediscoveryIntervalSeconds())

	// 6. action policy metadata
	if registrationClient.IActionPolicyProvider != nil {
		probeInfoBuilder.WithActionPolicySet(registrationClient.GetActionPolicy())
	}

	// 7. entity metadata
	if registrationClient.IEntityMetadataProvider != nil {
		probeInfoBuilder.WithEntityMetadata(registrationClient.GetEntityMetadata())
	}

	probeInfo := probeInfoBuilder.Create()
	glog.V(2).Infof("ProbeInfo %++v\n", probeInfo)

	return probeInfo, nil
}

func (theProbe *TurboProbe) createValidationErrorDTO(errMsg string, severity proto.ErrorDTO_ErrorSeverity) *proto.ValidationResponse {
	errorDTO := &proto.ErrorDTO{
		Severity:    &severity,
		Description: &errMsg,
	}
	var errorDtoList []*proto.ErrorDTO
	errorDtoList = append(errorDtoList, errorDTO)

	validationResponse := &proto.ValidationResponse{
		ErrorDTO: errorDtoList,
	}
	return validationResponse
}

func (theProbe *TurboProbe) createDiscoveryErrorDTO(errMsg string, severity proto.ErrorDTO_ErrorSeverity) *proto.DiscoveryResponse {
	errorDTO := &proto.ErrorDTO{
		Severity:    &severity,
		Description: &errMsg,
	}
	var errorDtoList []*proto.ErrorDTO
	errorDtoList = append(errorDtoList, errorDTO)

	discoveryResponse := &proto.DiscoveryResponse{
		ErrorDTO: errorDtoList,
	}
	return discoveryResponse
}

func (theProbe *TurboProbe) createActionErrorDTO(errMsg string) *proto.ActionResult {
	var progress int32
	progress = 100
	state := proto.ActionResponseState_FAILED
	// build ActionResponse
	actionResponse := &proto.ActionResponse{
		ActionResponseState: &state,
		ResponseDescription: &errMsg,
		Progress:            &progress,
	}

	response := &proto.ActionResult{
		Response: actionResponse,
	}

	return response
}
