package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	restclient "github.com/turbonomic/turbo-api/pkg/client"

	"github.com/turbonomic/turbo-go-sdk/pkg/vmtapi"
	"github.com/turbonomic/turbo-go-sdk/pkg/communication"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"

	"github.com/golang/glog"
)

type TAPService struct {
	// Interface to the Turbo Server
	*probe.TurboProbe
	*vmtapi.TurboAPIHandler // TODO: use vmtapi.Client
	*restclient.Client
}

func (tapService *TAPService) ConnectToTurbo() {
	IsRegistered := make(chan bool, 1)
	defer close(IsRegistered)

	// start a separate go routine to connect to the Turbo server
	go communication.InitMediationContainer(IsRegistered)
	//go tapService.mediationContainer.Init(IsRegistered)

	// Wait for probe registration complete to create targets in turbo server
	tapService.createTurboTargets(IsRegistered)
}

// Invokes the Turbo Rest API to create VMTTarget representing the target environment
// that is being controlled by the TAP service.
// Targets are created only after the service is notified of successful registration with the server
func (tapService *TAPService) createTurboTargets(IsRegistered chan bool) {
	glog.Infof("********* Waiting for registration complete .... %s\n", IsRegistered)
	// Block till a message arrives on the channel
	status := <-IsRegistered
	if !status {
		glog.Infof("Probe " + tapService.ProbeCategory + "::" + tapService.ProbeType + " is not registered")
		return
	}
	glog.Infof("Probe " + tapService.ProbeCategory + "::" + tapService.ProbeType + " Registered : === Add Targets ===")
	var targets []*probe.TurboTarget
	targets = tapService.GetProbeTargets()
	for _, targetInfo := range targets {
		glog.Infof("Adding target %s", targetInfo)
		tapService.AddTurboTarget(targetInfo)
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

func ParseTurboCommunicationConfig(configFile string) (*TurboCommunicationConfig, error) {
	// load the config
	turboCommConfig, err := readTurboCommunicationConfig(configFile)
	if turboCommConfig == nil {
		return nil, err
	}
	glog.Infof("WebscoketContainer Config : ", turboCommConfig.ContainerConfig)
	glog.Infof("RestAPI Config: ", turboCommConfig.TurboAPIConfig)

	// validate the config
	// TODO: return validation errors
	_, err = turboCommConfig.ValidateContainerConfig()
	if err != nil {
		return nil, err
	}
	_, err = turboCommConfig.ValidateTurboAPIConfig()
	if err != nil {
		return nil, err
	}
	glog.Infof("---------- Loaded Turbo Communication Config ---------")
	return turboCommConfig, nil
}

func readTurboCommunicationConfig(path string) (*TurboCommunicationConfig, error) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, fmt.Errorf("File error: %v\n" + e.Error())
	}
	//fmt.Println(string(file))
	var config TurboCommunicationConfig

	err := json.Unmarshal(file, &config)

	if err != nil {
		return nil, fmt.Errorf("Unmarshall error :%v\n" + err.Error())
	}
	glog.Infof("Results: %+v\n", config)
	return &config, nil
}

// ==============================================================================
// Convenience builder for building a TAPService
type TAPServiceBuilder struct {
	tapService *TAPService

	err error
}

// Get an instance of TAPServiceBuilder
func NewTAPServiceBuilder() *TAPServiceBuilder {
	serviceBuilder := &TAPServiceBuilder{}
	service := &TAPService{}
	serviceBuilder.tapService = service
	return serviceBuilder
}

// Build a new instance of TAPService.
func (pb *TAPServiceBuilder) Create() (*TAPService, error) {
	if pb.err != nil {
		return nil, pb.err
	}

	return pb.tapService, nil
}

func (pb *TAPServiceBuilder) WithTurboCommunicator(commConfig *TurboCommunicationConfig) *TAPServiceBuilder {
	if pb.err != nil {
		return pb
	}
	// The WebSocket Container
	communication.CreateMediationContainer(commConfig.ContainerConfig)
	// The RestAPI Handler
	// TODO: if rest api config has validation errors or not specified, do not create the handler
	turboApiHandler := vmtapi.NewTurboAPIHandler(commConfig.TurboAPIConfig)
	pb.tapService.TurboAPIHandler = turboApiHandler

	return pb
}

// The TurboProbe representing the service in the Turbo server
func (pb *TAPServiceBuilder) WithTurboProbe(probeBuilder *probe.ProbeBuilder) *TAPServiceBuilder {
	if pb.err != nil {
		return pb
	}
	// Create the probe
	turboProbe, err := probeBuilder.Create()
	if err != nil {
		pb.err = err
		return pb
	}

	pb.tapService.TurboProbe = turboProbe
	// Register the probe
	err = communication.LoadProbe(turboProbe)
	if err != nil {
		pb.err = err
		return pb
	}

	return pb
}
