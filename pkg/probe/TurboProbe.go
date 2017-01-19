package probe

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"

	"github.com/golang/glog"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type TurboProbe struct {
	RegistrationClient  	TurboRegistrationClient
	ActionExecutor 		IActionExecutor
	ProbeType		string
	ProbeCategory		string
	DiscoveryClientMap 	map[string]TurboDiscoveryClient
	TurboAPIClient		*TurboAPIHandler

	// TODO: state with respect to the server
	IsRegistered		chan bool
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

func NewTurboProbe(probeConf *ProbeConfig) *TurboProbe {
	fmt.Println("[TurboProbe] : ", probeConf)
	// load the probe config
	//config := parseProbeConfig(configFile)

	myProbe := &TurboProbe{
		ProbeType: probeConf.ProbeType,
		ProbeCategory: probeConf.ProbeCategory,
		DiscoveryClientMap: make(map[string]TurboDiscoveryClient),
		IsRegistered: make(chan bool, 1),	// buffered channel so the send does not block
	}

	fmt.Printf("[TurboProbe] : Created TurboProbe %s\n", myProbe)
	// Read config file to get the probe category and client
	return myProbe
}

func parseProbeConfig(configFile string) *ProbeConfig {
	// load the config
	probeConfig := readProbeConfig(configFile)
	fmt.Println("ProbeCategory : " + string(probeConfig.ProbeCategory))
	fmt.Println("ProbeType : " + probeConfig.ProbeType)

	// validate the config
	probeConfig.validate()
	return probeConfig
}

// TODO:
func (probeConfig *ProbeConfig) validate() bool {
	return false
}

func readProbeConfig(path string) *ProbeConfig {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		glog.Errorf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Println(string(file))
	var config ProbeConfig

	json.Unmarshal(file, &config)

	glog.V(4).Infof("Results: %+v\n", config)
	return &config
}

// ==============================================================================================================

func (theProbe *TurboProbe) SetTurboAPIHandler(turboApiClient *TurboAPIHandler) {
	theProbe.TurboAPIClient = turboApiClient
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

func (theProbe *TurboProbe) AddTargets()  {
	isRegistered := <- theProbe.IsRegistered
	if !isRegistered {
		fmt.Println("[TurboProbe] Probe " + theProbe.ProbeCategory + "::" + theProbe.ProbeType + " should be registered before adding Targets")
		return
	}
	fmt.Println("[TurboProbe] Probe " + theProbe.ProbeCategory + "::" + theProbe.ProbeType + " Registered : ============ Add Targets ========")
	var targets []*TurboTarget
	targets = theProbe.GetProbeTargets()
	for _, targetInfo := range targets {
		theProbe.TurboAPIClient.AddTarget(targetInfo)
	}
}


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