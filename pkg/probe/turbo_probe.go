package probe

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"github.com/golang/glog"
)

// Turbo Probe Abstraction
// Consists of clients that handle probe registration metadata, discovery and action execution for different probe targets
type TurboProbe struct {
	ProbeType          string
	ProbeCategory      string
	RegistrationClient TurboRegistrationClient
	DiscoveryClientMap map[string]TurboDiscoveryClient

	ActionClient TurboActionExecutorClient
}

type TurboRegistrationClient interface {
	GetAccountDefinition() []*proto.AccountDefEntry
	GetSupplyChainDefinition() []*proto.TemplateDTO
	GetIdentifyingFields() string
	// TODO: - add methods to get entity metadata, action policy data
}

type TurboDiscoveryClient interface {
	Discover(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error)
	Validate(accountValues []*proto.AccountValue) (*proto.ValidationResponse, error)
	GetAccountValues() *TurboTargetInfo
}

type ProbeConfig struct {
	ProbeType     string
	ProbeCategory string
}

type TurboTargetConf interface {
	GetAccountValues() []*proto.AccountValue
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
		ProbeType:          probeConf.ProbeType,
		ProbeCategory:      probeConf.ProbeCategory,
		DiscoveryClientMap: make(map[string]TurboDiscoveryClient),
	}

	glog.V(2).Infof("[NewTurboProbe] Created TurboProbe: %s", myProbe)
	return myProbe, nil
}

func (theProbe *TurboProbe) getDiscoveryClient(targetIdentifier string) TurboDiscoveryClient {
	return theProbe.DiscoveryClientMap[targetIdentifier]
}

// TODO: this method should be synchronized
func (theProbe *TurboProbe) GetTurboDiscoveryClient(accountValues []*proto.AccountValue) TurboDiscoveryClient {
	var address string
	identifyingField := theProbe.RegistrationClient.GetIdentifyingFields()

	for _, accVal := range accountValues {
		if accVal.GetKey() == identifyingField {
			address = accVal.GetStringValue()
		}
	}
	target := theProbe.getDiscoveryClient(address)

	if target == nil {
		glog.Errorf("[GetTurboDiscoveryClient] Cannot find Target for address: " + address)
		//TODO: CreateDiscoveryClient(address, accountValues, )
		return nil
	}
	glog.V(2).Infof("[GetTurboDiscoveryClient] Found Target for address: %s", address)
	return target
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

func (theProbe *TurboProbe) ExecuteAction(actionExecutionDTO *proto.ActionExecutionDTO, accountValues []*proto.AccountValue,
	progressTracker ActionProgressTracker) *proto.ActionResult {
	if theProbe.ActionClient == nil {
		glog.V(3).Infof("ActionClient not defined for Probe %s", theProbe.ProbeType)
		return theProbe.createActionErrorDTO("ActionClient not defined for Probe " + theProbe.ProbeType)
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
		targetInfo.targetType = theProbe.ProbeType
		targetInfo.targetIdentifierField = targetId

		targets = append(targets, targetInfo)
	}
	return targets
}

// The ProbeInfo for the probe
func (theProbe *TurboProbe) GetProbeInfo() (*proto.ProbeInfo, error) {
	// 1. Get the account definition for probe
	var acctDefProps []*proto.AccountDefEntry
	var templateDtos []*proto.TemplateDTO

	acctDefProps = theProbe.RegistrationClient.GetAccountDefinition()

	// 2. Get the supply chain.
	templateDtos = theProbe.RegistrationClient.GetSupplyChainDefinition()

	// 3. construct the example probe info.
	probeCat := theProbe.ProbeCategory
	probeType := theProbe.ProbeType
	probeInfo := builder.NewProbeInfoBuilder(probeType, probeCat, templateDtos, acctDefProps).Create()

	// Fields that serve to uniquely identify a target
	id := theProbe.RegistrationClient.GetIdentifyingFields()

	probeInfo.TargetIdentifierField = &id

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
