package communication

import (
	"fmt"
	"sync"
	"errors"
	"github.com/turbonomic/turbo-go-sdk/pkg/probe"
	"github.com/golang/glog"
)

type ContainerConfig struct {
	VmtServerAddress string
	VmtServerPort    string
	VmtUserName      string
	VmtPassword      string
	ConnectionRetry  int16 //
	IsSecure         bool
	BaseServerUrl    string
}

type ContainerConfigDefault string
const (
	VMT_Remote_Mediation_Server ContainerConfigDefault = "/vmturbo/remoteMediation"
	VMT_Remote_Mediation_Server_User ContainerConfigDefault = "vmtRemoteMediation"
	VMT_Remote_Mediation_Server_Pwd ContainerConfigDefault = "vmtRemoteMediation"
)


func (containerConfig *ContainerConfig) ValidateContainerConfig() (bool, error) {
	if containerConfig.VmtServerAddress == "" {
		return false, errors.New("Turbo Server Address IP is required "+ fmt.Sprint(containerConfig))
	}
	if containerConfig.BaseServerUrl == "" {
		containerConfig.BaseServerUrl = string(VMT_Remote_Mediation_Server)
	}
	if containerConfig.VmtUserName == "" {
		containerConfig.VmtUserName = string(VMT_Remote_Mediation_Server_User)
	}
	if containerConfig.VmtPassword == "" {
		containerConfig.VmtPassword = string(VMT_Remote_Mediation_Server_Pwd)
	}

	fmt.Println("========== Container Config =============")
	fmt.Println("VmtServerAddress : " + string(containerConfig.VmtServerAddress))
	fmt.Println("VmtUsername : " + containerConfig.VmtUserName)
	fmt.Println("VmtPassword : " + containerConfig.VmtPassword)
	fmt.Println("isSecure : ", containerConfig.IsSecure)
	fmt.Println("BaseServerUrl : " + containerConfig.BaseServerUrl)
	return true, nil
}

// ===========================================================================================================

type mediationContainer struct {
	// Configuration for making the transport connection
	containerConfig *ContainerConfig
	// Map of probes registered with the container
	allProbes map[string]*ProbeProperties
	// The Mediation client that will handle the messages from the server
	theRemoteMediationClient *remoteMediationClient
}

type ProbeSignature struct {
	ProbeType     string
	ProbeCategory string
}

type ProbeProperties struct {
	ProbeSignature *ProbeSignature
	Probe          *probe.TurboProbe
}

var (
	theInstance  *mediationContainer
	once sync.Once
)

func singletonMediationContainer() *mediationContainer {
	once.Do(func() {
		if theInstance == nil {
			theInstance = &mediationContainer{
				allProbes: make(map[string]*ProbeProperties),
				// TODO: create the probe store and mediation client here
			}
		}
	})
	return theInstance
}

// Static method to get the singleton instance of the Mediation Container
func CreateMediationContainer(containerConfig *ContainerConfig) *mediationContainer {
	// Validate the container config
	containerConfig.ValidateContainerConfig()
	glog.Infof("---------- Created MediationContainer ----------")
	theContainer := singletonMediationContainer()  //&mediationContainer {} // TODO: make a singleton instance

	//  Load the main container configuration file and validate it
	theContainer.containerConfig = containerConfig

	// Create the RemoteMediationClient to start the session with the server
	theContainer.theRemoteMediationClient = CreateRemoteMediationClient(theContainer.allProbes, theContainer.containerConfig)

	return theContainer
}

// Start the RemoteMediationClient
func InitMediationContainer(probeRegisteredMsg chan bool) {
	theContainer := singletonMediationContainer()
	//func (theContainer *mediationContainer) Init(probeRegisteredMsg chan bool) {
	glog.Infof("[MediationContainer] Initializing Mediation Container .....")
	// Assert that the probes are registered before starting the handshake
	if len(theContainer.allProbes) == 0 {
		glog.Errorf("No probes are registered with the container")
		return
	}
	// Open connection to the server and start server handshake to register probes
	glog.V(2).Infof("Registering ", len(theContainer.allProbes), " probes")

	remoteMediationClient := theContainer.theRemoteMediationClient
	remoteMediationClient.Init(probeRegisteredMsg)
}

func (theContainer *mediationContainer) Close() {
	theContainer.theRemoteMediationClient.Stop()
	// TODO: clear probe map ?
}

// ============================= Probe Management ==================
func LoadProbe(probe *probe.TurboProbe) error {
	//func (theContainer *mediationContainer) LoadProbe(probe *probe.TurboProbe)

	// load the probe config
	config := &ProbeSignature {
		ProbeCategory: probe.ProbeCategory,
		ProbeType:     probe.ProbeType,
	}

	probeProp := &ProbeProperties {
		ProbeSignature: config,
		Probe:          probe,
	}
	theContainer := singletonMediationContainer()
	if theContainer == nil {
		return errors.New("[LoadProbe] **** Null mediation container ****")
	}
	// TODO: check if the probe type already exists and warn before overwriting
	theContainer.allProbes[config.ProbeType] = probeProp
	glog.Infof("Registered " + config.ProbeCategory + "::" + config.ProbeType)
	return nil
}

func GetProbe(probeType string) (*probe.TurboProbe, error) {
	theContainer := singletonMediationContainer()
	if theContainer == nil {
		return nil, errors.New("[GetProbe] **** Null mediation container ****")
	}
	probeProps := theContainer.allProbes[probeType]

	if probeProps != nil {
		probe := probeProps.Probe
		registrationClient := probe.RegistrationClient
		acctDefProps := registrationClient.GetAccountDefinition()
		glog.V(2).Infof("Found "+probeProps.ProbeSignature.ProbeCategory+"::"+probeProps.ProbeSignature.ProbeType+" ==> ", acctDefProps)
		return probe, nil
	}
	return nil, errors.New("[GetProbe] Cannot find Probe of type " + probeType)

}
