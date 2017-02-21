package probe

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// Turbo Probe Abstraction
// Consists of clients that handle probe registration and discovery for different probe targets
type TurboProbe struct {
	ProbeType          string
	ProbeCategory      string
	RegistrationClient TurboRegistrationClient
	DiscoveryClientMap map[string]TurboDiscoveryClient
	ActionExecutor     ActionExecutorClient //TODO:
}

type TurboRegistrationClient interface {
	GetAccountDefinition() []*proto.AccountDefEntry
	GetSupplyChainDefinition() []*proto.TemplateDTO
	GetIdentifyingFields() string
	// TODO: - add methods to get entity metadata, action policy data
}

type TurboDiscoveryClient interface {
	Discover(accountValues []*proto.AccountValue) *proto.DiscoveryResponse
	Validate(accountValues []*proto.AccountValue) *proto.ValidationResponse
	GetAccountValues() *TurboTargetInfo
}

type DiscoveryClientInstantiateFunc func(accountValues []*proto.AccountValue) TurboDiscoveryClient

type ProbeConfig struct {
	ProbeType     string
	ProbeCategory string
}

// ==============================================================================================================

func NewTurboProbe(probeConf *ProbeConfig) (*TurboProbe, error) {
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

	glog.Infof("[NewTurboProbe] Created TurboProbe %s\n", myProbe)
	return myProbe, nil
}

func (theProbe *TurboProbe) SetProbeRegistrationClient(registrationClient TurboRegistrationClient) {
	theProbe.RegistrationClient = registrationClient
}

func (theProbe *TurboProbe) SetDiscoveryClient(targetIdentifier string, discoveryClient TurboDiscoveryClient) {
	theProbe.DiscoveryClientMap[targetIdentifier] = discoveryClient
}

//
//// TODO: CreateClient func as input
//func (theProbe *TurboProbe) SetDiscoveryClientFunc(targetIdentifier string, accountValues[] *proto.AccountValue),
//							createClientFunc DiscoveryClientInstantiateFunc) {
//	discoveryClient := createClientFunc(accountValues)
//	theProbe.DiscoveryClientMap[targetIdentifier] = discoveryClient
//}

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
		glog.Errorf("****** [GetTurboDiscoveryClient] Cannot find Target for address : " + address)
		//TODO: CreateDiscoveryClient(address, accountValues, )
		return nil
	}
	glog.Infof("[GetTurboDiscoveryClient] Found Target for address : " + address)
	return target
}

func (theProbe *TurboProbe) DiscoverTarget(accountValues []*proto.AccountValue) *proto.DiscoveryResponse {
	glog.Infof("Discover Target : ", accountValues)
	var handler TurboDiscoveryClient
	handler = theProbe.GetTurboDiscoveryClient(accountValues)
	var discoveryResponse *proto.DiscoveryResponse
	if handler != nil {
		discoveryResponse = handler.Discover(accountValues)
	}
	if discoveryResponse == nil {
		description := "Null Entity DTOs"
		severity := proto.ErrorDTO_CRITICAL
		errorDTO := &proto.ErrorDTO{
			Severity:    &severity,
			Description: &description,
		}
		var errorDtoList []*proto.ErrorDTO
		errorDtoList = append(errorDtoList, errorDTO)

		discoveryResponse = &proto.DiscoveryResponse{
			ErrorDTO: errorDtoList,
		}
	}
	glog.Infof("Error discovering target %s", discoveryResponse)
	return discoveryResponse
}

func (theProbe *TurboProbe) ValidateTarget(accountValues []*proto.AccountValue) *proto.ValidationResponse {
	glog.Infof("Validate Target : ", accountValues)
	var handler TurboDiscoveryClient
	handler = theProbe.GetTurboDiscoveryClient(accountValues)

	if handler != nil {
		glog.Infof("Send validation request to handler %s\n", handler)
		return handler.Validate(accountValues)
	}

	// TODO: if the handler is nil, implies the target is added from the UI
	// Create a new discovery client for this target and add it to the map of discovery clients
	// allow to pass a func to instantiate a default discovery client

	description := "Target not found"
	severity := proto.ErrorDTO_CRITICAL

	errorDTO := &proto.ErrorDTO{
		Severity:    &severity,
		Description: &description,
	}
	var errorDtoList []*proto.ErrorDTO
	errorDtoList = append(errorDtoList, errorDTO)

	validationResponse := &proto.ValidationResponse{
		ErrorDTO: errorDtoList,
	}
	glog.Infof("Error validating target %s", validationResponse)
	return validationResponse
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
