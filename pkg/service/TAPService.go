package probe

import (
	"fmt"
	"io/ioutil"
	"github.com/golang/glog"
	"os"
	"encoding/json"

	"github.com/turbonomic/turbo-go-sdk/pkg/communication"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
)

type TAPService struct {
	// Interface to the Turbo Server
	*communication.MediationContainer
	*probe.TurboProbe
	*TurboAPIHandler
	IsRegistered		chan bool
}

// Invokes the Turbo Rest API to create VMTTarget representing the target environment
// that is being controlled by the TAP service.
// Targets are created only after the service is notified of successful registration with the server
func (tapService *TAPService) addTargets()  {
	// Block till a message arrives on the channel
	isRegistered := <- tapService.IsRegistered
	if !isRegistered {
		fmt.Println("[TAPService] Probe " + tapService.ProbeCategory + "::" + tapService.ProbeType + " should be registered before adding Targets")
		return
	}
	fmt.Println("[TAPService] Probe " + tapService.ProbeCategory + "::" + tapService.ProbeType + " Registered : ============ Add Targets ========")
	var targets []*probe.TurboTarget
	targets = tapService.GetProbeTargets()
	for _, targetInfo := range targets {
		tapService.AddTarget(targetInfo)
	}
}

func (tapService *TAPService) ConnectToTurbo() {
	// Connect to the Turbo server
	tapService.MediationContainer.Init(tapService.IsRegistered)

	// start a separate go routine to listen for probe registration and create targets in turbo server
	go tapService.addTargets()
}

// ==============================================================================

type TurboCommunicationConfig struct  {
	*TurboAPIConfig
	*communication.ContainerConfig
}

func parseTurboCommunicationConfig (configFile string) *TurboCommunicationConfig {
	// load the config
	turboCommConfig := readTurboCommunicationConfig (configFile)
	fmt.Println("WebscoketContainer Config : ", turboCommConfig.ContainerConfig)
	fmt.Println("RestAPI Config: ", turboCommConfig.TurboAPIConfig)

	// validate the config
	// TODO: return validation errors
	turboCommConfig.ValidateContainerConfig()
	turboCommConfig.ValidateTurboAPIConfig()
	return turboCommConfig
}

func readTurboCommunicationConfig (path string) *TurboCommunicationConfig {
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
	}
	fmt.Printf("[TurboCommunicationConfig] Results: %+v\n", config)
	return &config
}

// ==============================================================================
type TAPServiceBuilder struct {
	tapService *TAPService
}
// Get an instance of ClientMessageBuilder
func NewTAPServiceBuilder () *TAPServiceBuilder {
	serviceBuilder := &TAPServiceBuilder{}
	service := &TAPService {
		IsRegistered: make(chan bool, 1),	// buffered channel so the send does not block
	}
	serviceBuilder.tapService = service
	return serviceBuilder
}

// Build a new instance of TAPService.
func (pb *TAPServiceBuilder) Create() *TAPService {
	if &pb.tapService.TurboProbe == nil {
		fmt.Println("[TAPServiceBuilder] Null turbo probe")	//TODO: throw exception
		return nil
	}

	return pb.tapService
}

// The Communication Layer to communicate with the Turbo server
func (pb *TAPServiceBuilder) WithTurboCommunicator(commConfFile string) *TAPServiceBuilder {
	//  Load the main container configuration file and validate it
	fmt.Println("[TAPServiceBuilder] TurboCommunicator configuration from %s", commConfFile)
	commConfig := parseTurboCommunicationConfig(commConfFile)
	fmt.Println("---------- Loaded Turbo Communication Config ---------")
	// The Webscoket Container
	theContainer := communication.CreateMediationContainer(commConfig.ContainerConfig)
	pb.tapService.MediationContainer = theContainer
	// The RestAPI Handler
	// TODO: if rest api config has validation errors or not specified, do not create the handler
	turboApiHandler := NewTurboAPIHandler(commConfig.TurboAPIConfig)
	pb.tapService.TurboAPIHandler = turboApiHandler
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


