package service

import (
	"fmt"
	"net/url"

	restclient "github.com/turbonomic/turbo-api/pkg/client"

	"github.com/turbonomic/turbo-go-sdk/pkg/mediationcontainer"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"

	"github.com/golang/glog"
)

const (
	defaultTurboAPIPath = "/vmturbo/rest"
)

type TAPService struct {
	// Interface to the Turbo Server
	*probe.TurboProbe
	*restclient.Client
}

func (tapService *TAPService) ConnectToTurbo() {
	IsRegistered := make(chan bool, 1)
	defer close(IsRegistered)

	// start a separate go routine to connect to the Turbo server
	go mediationcontainer.InitMediationContainer(IsRegistered)
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
	targetInfos := tapService.GetProbeTargets()
	for _, targetInfo := range targetInfos {
		target := targetInfo.GetTargetInstance()
		glog.V(4).Infof("Now adding target %v", target)
		resp, err := tapService.AddTarget(target)
		if err != nil {
			glog.Errorf("Error during add target %v: %s", targetInfo, err)
			// TODO, do we want to return an error?
			continue
		}
		glog.V(3).Infof("Successfully add target: %v", resp)
	}
}

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
func (builder *TAPServiceBuilder) Create() (*TAPService, error) {
	if builder.err != nil {
		return nil, builder.err
	}

	return builder.tapService, nil
}

// Build the mediation container and Turbo API client.
func (builder *TAPServiceBuilder) WithTurboCommunicator(commConfig *TurboCommunicationConfig) *TAPServiceBuilder {
	if builder.err != nil {
		return builder
	}

	// Create the mediation container. This is a singleton.
	containerConfig := &mediationcontainer.MediationContainerConfig{
		ServerMeta:      commConfig.ServerMeta,
		WebSocketConfig: commConfig.WebSocketConfig,
	}
	mediationcontainer.CreateMediationContainer(containerConfig)

	// The RestAPI Handler
	serverAddress, err := url.Parse(commConfig.TurboServer)
	if err != nil {
		builder.err = fmt.Errorf("Error during create Turbo API client config: Incorrect URL: %s\n", err)
		return builder
	}
	config := restclient.NewConfigBuilder(serverAddress).
		APIPath(defaultTurboAPIPath).
		BasicAuthentication(commConfig.OpsManagerUsername, commConfig.OpsManagerPassword).
		Create()
	glog.V(4).Infof("The Turbo API client config authentication is: %s, %s", commConfig.OpsManagerUsername, commConfig.OpsManagerPassword)
	glog.V(4).Infof("The Turbo API client config is create successfully: %v", config)
	apiClient, err := restclient.NewAPIClientWithBA(config)
	if err != nil {
		builder.err = fmt.Errorf("Error during create Turbo API client: %s\n", err)
		return builder
	}
	glog.V(4).Infof("The Turbo API client is create successfully: %v", apiClient)
	builder.tapService.Client = apiClient

	return builder
}

// The TurboProbe representing the service in the Turbo server
func (builder *TAPServiceBuilder) WithTurboProbe(probeBuilder *probe.ProbeBuilder) *TAPServiceBuilder {
	if builder.err != nil {
		return builder
	}
	// Create the probe
	turboProbe, err := probeBuilder.Create()
	if err != nil {
		builder.err = err
		return builder
	}

	builder.tapService.TurboProbe = turboProbe
	// Register the probe
	err = mediationcontainer.LoadProbe(turboProbe)
	if err != nil {
		builder.err = err
		return builder
	}

	return builder
}
