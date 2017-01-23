package probe

import (
	"fmt"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
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
	// TODO: - add methods to get entity metadata, action policy data
}

type TurboDiscoveryClient interface {
	Discover(accountValues[] *proto.AccountValue) *proto.DiscoveryResponse
	Validate(accountValues[] *proto.AccountValue) *proto.ValidationResponse
	GetAccountValues() *TurboTarget
}

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
		//IsRegistered: make(chan bool, 1),	// buffered channel so the send does not block
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

func (theProbe *TurboProbe) getDiscoveryClient(targetIdentifier string) TurboDiscoveryClient {
	return theProbe.DiscoveryClientMap[targetIdentifier]
}

func (theProbe *TurboProbe) GetTurboDiscoveryClient(accountValues[] *proto.AccountValue) TurboDiscoveryClient {
	var address string

	for _, accVal  := range accountValues {

		if *accVal.Key == "targetIdentifier" {
			address = *accVal.StringValue
		}
	}
	target := theProbe.getDiscoveryClient(address)

	if target == nil {
		fmt.Println("****** [TurboProbe][DiscoveryTarget] Cannot find Target for address : " + address)
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
		return handler.Validate(accountValues)
	}

	fmt.Println("[TurboProbe] Error validating target ", accountValues)
	return nil
}
//
//func (theProbe *TurboProbe) AddTargets()  {
//	isRegistered := <- theProbe.IsRegistered
//	if !isRegistered {
//		fmt.Println("[TurboProbe] Probe " + theProbe.ProbeCategory + "::" + theProbe.ProbeType + " should be registered before adding Targets")
//		return
//	}
//	fmt.Println("[TurboProbe] Probe " + theProbe.ProbeCategory + "::" + theProbe.ProbeType + " Registered : ============ Add Targets ========")
//	var targets []*TurboTarget
//	targets = theProbe.GetProbeTargets()
//	for _, targetInfo := range targets {
//		theProbe.TurboAPIClient.AddTarget(targetInfo)
//	}
//}


// ==============================================================================================================
func (theProbe *TurboProbe) GetProbeTargets() []*TurboTarget {
	// Iterate over the discovery client map and send requests to the server
	var targets []*TurboTarget
	for targetId, discoveryClient := range theProbe.DiscoveryClientMap {
		//targetInfo := &TurboTarget{
		//	targetType: theProbe.ProbeType,
		//	targetIdentifier: targetId,
		//}

		targetInfo := discoveryClient.GetAccountValues()
		targetInfo.targetType = theProbe.ProbeType
		targetInfo.targetIdentifier = targetId

		targets = append(targets, targetInfo)
	}
	return targets
}

func (theProbe *TurboProbe) GetProbeInfo() (*proto.ProbeInfo, error) {
	// TODO:
	// 1. Get the account definition for probe
	var acctDefProps []*proto.AccountDefEntry
	var templateDtos  []*proto.TemplateDTO

	acctDefProps = theProbe.RegistrationClient.GetAccountDefinition()

	// 2. Get the supply chain.
	templateDtos = theProbe.RegistrationClient.GetSupplyChainDefinition()

	// 3. construct the example probe info.
	probeCat := theProbe.ProbeCategory
	probeType := theProbe.ProbeType
	probeInfo := NewProbeInfoBuilder(probeType, probeCat, templateDtos, acctDefProps).Create()
	id := "targetIdentifier"		// TODO: parameterize this for different probes using AccountInfo struct
	probeInfo.TargetIdentifierField = &id

	// 4. Add example probe to probeInfo list, here it is the only probe supported.
	//var probes []*proto.ProbeInfo
	//probes = append(probes, exampleProbe)

	return probeInfo, nil
}