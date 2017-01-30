package communication

import (
	"fmt"
	probe "github.com/turbonomic/turbo-go-sdk/pkg/probe"
)

type ContainerConfig struct  {
	VmtServerAddress string
	VmtServerPort 	 string
	VmtUserName	 string
	VmtPassword	 string
	ConnectionRetry	 int16		//
	IsSecure	 bool
	ApplicationBase  string
	ProbesDir        string		//TODO: dont need until we can package like in the java sdk
}

// TODO:
func (containerConfig *ContainerConfig) ValidateContainerConfig() bool {
	fmt.Println("========== Container Config =============")
	fmt.Println("VmtServerAddress : " + string(containerConfig.VmtServerAddress))
	fmt.Println("VmtUsername : " + containerConfig.VmtUserName)
	fmt.Println("VmtPassword : " + containerConfig.VmtPassword)
	fmt.Println("isSecure : " , containerConfig.IsSecure)
	fmt.Println("ApplicationBase : " + containerConfig.ApplicationBase)
	return true
}

// ======================================================================

type ProbeSignature struct {
	ProbeType	string
	ProbeCategory	string
}

type ProbeProperties struct {
	ProbeSignature *ProbeSignature
	Probe          *probe.TurboProbe
}

// ===========================================================================================================

type MediationContainer struct {
	// map of probes
	allProbes                map[string]*ProbeProperties
	theRemoteMediationClient *RemoteMediationClient
	containerConfig          *ContainerConfig
}

// Static method to create an instance of the Mediation Container
func CreateMediationContainer(containerConfig *ContainerConfig) *MediationContainer {
	fmt.Println("---------- Created MediationContainer ----------")
	theContainer := &MediationContainer{}	// TODO: make a singleton instance

	//  Load the main container configuration file and validate it
	theContainer.containerConfig = containerConfig

	// Map for the Probes
	theContainer.allProbes = make(map[string]*ProbeProperties)

	// Create the RemoteMediationClient to start the session with the server
	theContainer.theRemoteMediationClient = CreateRemoteMediationClient(theContainer.allProbes, theContainer.containerConfig)

	return theContainer
}

//func (theContainer *MediationContainer) GetRemoteMediationClient() *RemoteMediationClient {
//	return theContainer.theRemoteMediationClient
//}


// Start the RemoteMediationClient
func (theContainer *MediationContainer) Init(probeRegisteredMsg chan bool) {
	fmt.Println("[MediationContainer] Initializing Mediation Container .....")
	// Assert that the probes are registered before starting the handshake
	if len(theContainer.allProbes) == 0 {
		fmt.Println("[MediationContainer] No probes are registered with the container")
		return
	}
	// Open connection to the server and start server handshake to register probes
	fmt.Println("[MediationContainer] Registering ", len(theContainer.allProbes) , " probes")

	remoteMediationClient := theContainer.theRemoteMediationClient
	remoteMediationClient.Init(probeRegisteredMsg)
}

func (theContainer *MediationContainer) Close() {
	theContainer.theRemoteMediationClient.Stop()
	// TODO: clear probe map ?
}

// ============================= Probe Management ==================
func (theContainer *MediationContainer) LoadProbe(probe *probe.TurboProbe) {
	// load the probe config
	config := &ProbeSignature{
		ProbeCategory: probe.ProbeCategory,
		ProbeType: probe.ProbeType,
	}

	probeProp := &ProbeProperties{
		ProbeSignature: config,
		Probe: probe,
	}

	// TODO: check if the probe type already exists and warn before overwriting
	theContainer.allProbes[config.ProbeType] = probeProp //createProbeFunc(configFile)
	fmt.Println("[MediationContainer] Registered  " + config.ProbeCategory + "::" + config.ProbeType)
}

func (theContainer *MediationContainer) GetProbe(probeType string) *probe.TurboProbe {
	probeProps := theContainer.allProbes[probeType]

	if probeProps != nil {
		probe := probeProps.Probe
		registrationClient := probe.RegistrationClient
		acctDefProps := registrationClient.GetAccountDefinition()
		fmt.Println("[MediationContainer] Found " + probeProps.ProbeSignature.ProbeCategory + "::" + probeProps.ProbeSignature.ProbeType + " ==> " , acctDefProps)
		return probe
	}
	fmt.Println("[MediationContainer] Cannot find Probe of type " + probeType)
	return nil
}

