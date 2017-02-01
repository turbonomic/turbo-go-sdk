package service

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"os"

	restclient "github.com/turbonomic/turbo-api/pkg/client"
	"github.com/turbonomic/turbo-go-sdk/pkg/vmtapi"

	"github.com/turbonomic/turbo-go-sdk/pkg/communication"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
)

type TAPService struct {
	// Interface to the Turbo Server
	*communication.MediationContainer
	*probe.TurboProbe
	*vmtapi.TurboAPIHandler // TODO: use vmtapi.Client
	*restclient.Client
}

func (tapService *TAPService) ConnectToTurbo() {
	IsRegistered := make(chan bool, 1)

	// Connect to the Turbo server
	go tapService.MediationContainer.Init(IsRegistered)

	// start a separate go routine to listen for probe registration and create targets in turbo server
	tapService.createTurboTargets(IsRegistered)
}

// Invokes the Turbo Rest API to create VMTTarget representing the target environment
// that is being controlled by the TAP service.
// Targets are created only after the service is notified of successful registration with the server
func (tapService *TAPService) createTurboTargets(IsRegistered chan bool) {
	fmt.Printf("[******************* TAPService *****************] Waiting for registration complete .... %s\n", IsRegistered)
	// Block till a message arrives on the channel
	status := <-IsRegistered
	if !status {
		fmt.Println("[TAPService] Probe " + tapService.ProbeCategory + "::" + tapService.ProbeType + " should be registered before adding Targets")
		return
	}
	fmt.Println("[TAPService] Probe " + tapService.ProbeCategory + "::" + tapService.ProbeType + " Registered : ============ Add Targets ========")
	var targets []*probe.TurboTarget
	targets = tapService.GetProbeTargets()
	for _, targetInfo := range targets {
		fmt.Println("[TAPService] Adding target %s", targetInfo)
		tapService.AddTurboTarget(targetInfo)
		// TODO: use this api call
		//targetData := &vmttype.Target{
		//	Category: tapService.ProbeCategory,
		//	Type:     targetInfo.GetTargetType(),
		//	InputFields: []*vmttype.InputField{
		//		{
		//			Value:           &targetInfo.GetNameOrAddress(),
		//			Name:            "nameOrAddress",
		//			GroupProperties: []*vmttype.List{},
		//		},
		//		{
		//			Value:           "username",
		//			Name:            "VC-username",
		//			GroupProperties: []*vmttype.List{},
		//		},
		//		{
		//			Value:           "password",
		//			Name:            "VC-password",
		//			GroupProperties: []*vmttype.List{},
		//		},
		//	},
		//}
		//tapService.AddTarget(targetData)
	}
}

// ==============================================================================

// Configuration parameters for communicating with the Turbo server
type TurboCommunicationConfig struct {
	// Config for the Rest API client
	*vmtapi.TurboAPIConfig
	// Config for RemoteMediation client that communicates using websocket
	*communication.ContainerConfig
}

func ParseTurboCommunicationConfig(configFile string) *TurboCommunicationConfig {
	// load the config
	turboCommConfig := readTurboCommunicationConfig(configFile)
	if turboCommConfig == nil {
		os.Exit(1)
	}
	fmt.Println("WebscoketContainer Config : ", turboCommConfig.ContainerConfig)
	fmt.Println("RestAPI Config: ", turboCommConfig.TurboAPIConfig)

	// validate the config
	// TODO: return validation errors
	turboCommConfig.ValidateContainerConfig()
	turboCommConfig.ValidateTurboAPIConfig()
	fmt.Println("---------- Loaded Turbo Communication Config ---------")
	return turboCommConfig
}

func readTurboCommunicationConfig(path string) *TurboCommunicationConfig {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		glog.Errorf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Println(string(file))
	var config TurboCommunicationConfig

	err := json.Unmarshal(file, &config)

	if err != nil {
		fmt.Printf("[TurboCommunicationConfig] Unmarshall error :%v\n", err)
		return nil
	}
	fmt.Printf("[TurboCommunicationConfig] Results: %+v\n", config)
	return &config
}

// ==============================================================================
// Convenience builder for building a TAPService
type TAPServiceBuilder struct {
	tapService *TAPService
}

// Get an instance of TAPServiceBuilder
func NewTAPServiceBuilder() *TAPServiceBuilder {
	serviceBuilder := &TAPServiceBuilder{}
	service := &TAPService{
	//IsRegistered: make(chan bool, 1),	// buffered channel so the send does not block
	}
	serviceBuilder.tapService = service
	return serviceBuilder
}

// Build a new instance of TAPService.
func (pb *TAPServiceBuilder) Create() *TAPService {
	if &pb.tapService.TurboProbe == nil {
		fmt.Println("[TAPServiceBuilder] Null turbo probe") //TODO: throw exception
		return nil
	}

	return pb.tapService
}

func (pb *TAPServiceBuilder) WithTurboCommunicator(commConfig *TurboCommunicationConfig) *TAPServiceBuilder {
	// The Webscoket Container
	theContainer := communication.CreateMediationContainer(commConfig.ContainerConfig)
	pb.tapService.MediationContainer = theContainer
	// The RestAPI Handler
	// TODO: if rest api config has validation errors or not specified, do not create the handler
	turboApiHandler := vmtapi.NewTurboAPIHandler(commConfig.TurboAPIConfig)
	pb.tapService.TurboAPIHandler = turboApiHandler

	//config := restclient.NewConfigBuilder(commConfig.TurboAPIConfig.VmtRestServerAddress).
	//	APIPath("/vmturbo/rest").
	//	BasicAuthentication(commConfig.TurboAPIConfig.VmtRestUser, commConfig.TurboAPIConfig.VmtRestPassword).//"<UI-username>", "UI-password").
	//	Create()
	//client, err := restclient.NewAPIClientWithBA(config)
	//if err != nil {
	//	fmt.Errorf("[TAPServiceBuilder] Error creating  Turbo Rest API client: %s", err)
	//}
	//pb.tapService.Client = client

	return pb
}

// The Communication Layer to communicate with the Turbo server
// Uses Websocket to listen to server requests for discovery and actions
// Uses Rest API to send requests to server for deployment
func (pb *TAPServiceBuilder) WithTurboCommunicator2(commConfFile string) *TAPServiceBuilder {
	//  Load the main communication configuration file and validate it
	fmt.Println("[TAPServiceBuilder] TurboCommunicator configuration from %s", commConfFile)
	commConfig := ParseTurboCommunicationConfig(commConfFile)

	// The Webscoket Container
	theContainer := communication.CreateMediationContainer(commConfig.ContainerConfig)
	pb.tapService.MediationContainer = theContainer
	// The RestAPI Handler
	// TODO: if rest api config has validation errors or not specified, do not create the handler
	turboApiHandler := vmtapi.NewTurboAPIHandler(commConfig.TurboAPIConfig)
	pb.tapService.TurboAPIHandler = turboApiHandler

	//config := restclient.NewConfigBuilder(commConfig.TurboAPIConfig.VmtRestServerAddress).
	//	APIPath("/vmturbo/rest").
	//	BasicAuthentication(commConfig.TurboAPIConfig.VmtRestUser, commConfig.TurboAPIConfig.VmtRestPassword).//"<UI-username>", "UI-password").
	//	Create()
	//client, err := restclient.NewAPIClientWithBA(config)
	//if err != nil {
	//	fmt.Errorf("[TAPServiceBuilder] Error creating  Turbo Rest API client: %s", err)
	//}
	//pb.tapService.Client = client

	return pb
}

// The TurboProbe representing the service in the Turbo server
func (pb *TAPServiceBuilder) WithTurboProbe(probeBuilder *probe.ProbeBuilder) *TAPServiceBuilder {
	// Check that the MediationContainer has been created
	if pb.tapService.MediationContainer == nil {
		pb.tapService.TurboProbe = nil
		fmt.Println("[TAPServiceBuilder] Null Mediation Container") // TODO: throw exception
		return nil
	}
	turboProbe := probeBuilder.Create()
	pb.tapService.TurboProbe = turboProbe //TODO: throws exception

	// Load the probe in the container
	theContainer := pb.tapService.MediationContainer
	theContainer.LoadProbe(turboProbe)
	theContainer.GetProbe(turboProbe.ProbeType)
	return pb
}
