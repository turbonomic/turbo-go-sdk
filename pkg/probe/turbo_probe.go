package probe

import (
	"fmt"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	builder "github.com/turbonomic/turbo-go-sdk/pkg/builder"
)

// Turbo Probe Abstraction
// Consists of clients that handle probe registration and discovery for different probe targets
type TurboProbe struct {
	ProbeType		string
	ProbeCategory		string
	RegistrationClient  	TurboRegistrationClient
	DiscoveryClientMap 	map[string]TurboDiscoveryClient
	ActionExecutor 		IActionExecutor		//TODO:
}

type TurboRegistrationClient interface {
	GetAccountDefinition() []*proto.AccountDefEntry
	GetSupplyChainDefinition() []*proto.TemplateDTO
	GetIdentifyingFields() string
	// TODO: - add methods to get entity metadata, action policy data
}

type TurboDiscoveryClient interface {
	Discover(accountValues[] *proto.AccountValue) *proto.DiscoveryResponse
	Validate(accountValues[] *proto.AccountValue) *proto.ValidationResponse
	GetAccountValues() *TurboTarget
}

type DiscoveryClientInstantiateFunc func(accountValues[] *proto.AccountValue) TurboDiscoveryClient

// ==============================================================================================================
type ProbeConfig struct {
	ProbeType	string
	ProbeCategory	string
}

func (probeConfig *ProbeConfig) validate() bool {
	// Validate probe type and category
	if &probeConfig.ProbeType == nil {
		fmt.Println("[ProbeConfig] Null Probe type")	//TODO: throw exception
		return false
	}

	if &probeConfig.ProbeCategory == nil {
		fmt.Println("[ProbeConfig] Null probe category")	//TODO: throw exception
		return false
	}
	return true
}

// ==============================================================================================================

func NewTurboProbe(probeConf *ProbeConfig) *TurboProbe {
	fmt.Println("[TurboProbe] : ", probeConf)
	if !probeConf.validate() {
		fmt.Println("[NewTurboProbe] Errors creating new TurboProbe")	//TODO: throw exception
		return nil
	}
	myProbe := &TurboProbe{
		ProbeType: probeConf.ProbeType,
		ProbeCategory: probeConf.ProbeCategory,
		DiscoveryClientMap: make(map[string]TurboDiscoveryClient),
	}

	fmt.Printf("[TurboProbe] : Created TurboProbe %s\n", myProbe)
	return myProbe
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
func (theProbe *TurboProbe) GetTurboDiscoveryClient(accountValues[] *proto.AccountValue) TurboDiscoveryClient {
	var address string
	identifyingField := theProbe.RegistrationClient.GetIdentifyingFields()

	for _, accVal  := range accountValues {

		if *accVal.Key == identifyingField {
			address = *accVal.StringValue
		}
	}
	target := theProbe.getDiscoveryClient(address)

	if target == nil {
		fmt.Println("****** [TurboProbe][DiscoveryTarget] Cannot find Target for address : " + address)
		//TODO: CreateDiscoveryClient(address, accountValues, )
		return nil
	}
	fmt.Println("[TurboProbe][DiscoveryTarget] Found Target for address : " + address)
	return target
}


func (theProbe *TurboProbe) DiscoverTarget(accountValues[] *proto.AccountValue) *proto.DiscoveryResponse {
	fmt.Println("[TurboProbe] ============ Discover Target ========", accountValues)
	var handler TurboDiscoveryClient
	handler = theProbe.GetTurboDiscoveryClient(accountValues)
	if handler != nil {
		return handler.Discover(accountValues)
	}
	fmt.Println("[TurboProbe] Error discovering target ", accountValues)
	return nil
}

func (theProbe *TurboProbe) ValidateTarget(accountValues[] *proto.AccountValue) *proto.ValidationResponse {
	fmt.Println("[TurboProbe] ============ Validate Target ========", accountValues)
	var handler TurboDiscoveryClient
	handler = theProbe.GetTurboDiscoveryClient(accountValues)

	if handler != nil {
		fmt.Printf("[TurboProbe] Send validation request to handler %s\n", handler)
		return handler.Validate(accountValues)
	}

	// TODO: if the handler is nil, implies the target is added from the UI
	// Create a new discovery client for this target and add it to the map of discovery clients
	// allow to pass a func to instantiate a default discovery client

	fmt.Println("[TurboProbe] Error validating target ", accountValues)
	return nil
}

// ==============================================================================================================
// The Targets associated with this probe type
func (theProbe *TurboProbe) GetProbeTargets() []*TurboTarget {
	// Iterate over the discovery client map and send requests to the server
	var targets []*TurboTarget
	for targetId, discoveryClient := range theProbe.DiscoveryClientMap {

		targetInfo := discoveryClient.GetAccountValues()
		targetInfo.targetType = theProbe.ProbeType
		targetInfo.targetIdentifier = targetId

		targets = append(targets, targetInfo)
	}
	return targets
}

// The ProbeInfo for the probe
func (theProbe *TurboProbe) GetProbeInfo() (*proto.ProbeInfo, error) {
	// 1. Get the account definition for probe
	var acctDefProps []*proto.AccountDefEntry
	var templateDtos  []*proto.TemplateDTO

	acctDefProps = theProbe.RegistrationClient.GetAccountDefinition()

	// 2. Get the supply chain.
	templateDtos = theProbe.RegistrationClient.GetSupplyChainDefinition()

	// 3. construct the example probe info.
	probeCat := theProbe.ProbeCategory
	probeType := theProbe.ProbeType
	probeInfo := builder.NewProbeInfoBuilder(probeType, probeCat, templateDtos, acctDefProps).Create()
	id := theProbe.RegistrationClient.GetIdentifyingFields()

	// /id := "targetIdentifier"		// TODO: parameterize this for different probes using AccountInfo struct
	probeInfo.TargetIdentifierField = &id

	return probeInfo, nil
}