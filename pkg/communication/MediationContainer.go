package communication

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"fmt"

	"github.com/golang/glog"

	probe "github.com/turbonomic/turbo-go-sdk/pkg/probe"
)

/**
Mediation Container
- load probe jars/files
- start mediation handshake
 */

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
func (containerConfig *ContainerConfig) validate() bool {
	return false
}

//ParseConfig loads and parses the container configuration file.
//The configuration properties will be validated before creating the config object
func ParseContainerConfig(configFile string) *ContainerConfig {
	// load the config
	containerConfig := ReadConfig(configFile)
	fmt.Println("VmtServerAddress : " + string(containerConfig.VmtServerAddress))
	fmt.Println("VmtUsername : " + containerConfig.VmtUserName)
	fmt.Println("VmtPassword : " + containerConfig.VmtPassword)
	fmt.Println("isSecure : " , containerConfig.IsSecure)
	fmt.Println("ApplicationBase : " + containerConfig.ApplicationBase)

	// validate the config
	containerConfig.validate()
	return containerConfig
}

// ======================================================================

type ProbeConfig struct {
	ProbeType	string
	ProbeCategory	string
}

type ProbeProperties struct {
	ProbeConfig *ProbeConfig
	Probe *probe.TurboProbe
}

// TODO:
func (probeConfig *ProbeConfig) validate() bool {
	return false
}

func ParseProbeConfig(configFile string) *ProbeConfig {
	// load the config
	probeConfig := ReadProbeConfig(configFile)
	fmt.Println("ProbeCategory : " + string(probeConfig.ProbeCategory))
	fmt.Println("ProbeType : " + probeConfig.ProbeType)

	// validate the config
	probeConfig.validate()
	return probeConfig
}

// ===========================================================================================================

type MediationContainer struct {
	// map of probes
	allProbes                map[string]*ProbeProperties
	theRemoteMediationClient *RemoteMediationClient
	containerConfig          *ContainerConfig
}

// Static method to create an instance of the Mediation Container
func CreateMediationContainer(containerConfigFile string) *MediationContainer {
	fmt.Println("---------- Started MediationContainer ----------")
	theContainer := &MediationContainer{}	// TODO: make a singleton instance

	//  Load the main container configuration file and validate it
	theContainer.containerConfig = ParseContainerConfig(containerConfigFile)
	fmt.Println("---------- Loaded Container Config ---------")

	// Map for the Probes
	theContainer.allProbes = make(map[string]*ProbeProperties)

	// Create the RemoteMediationClient to start the session with the server
	theContainer.theRemoteMediationClient = CreateRemoteMediationClient(theContainer.allProbes, theContainer.containerConfig)

	return theContainer
}

func (theContainer *MediationContainer) GetRemoteMediationClient() *RemoteMediationClient {
	return theContainer.theRemoteMediationClient
}


// Start the RemoteMediationClient
func (theContainer *MediationContainer) Init() {
	// Assert that the probes are registered before starting the handshake
	if len(theContainer.allProbes) == 0 {
		fmt.Println("[MediationContainer] No probes are registered with the container")
		return
	}
	// Open connection to the server and start server handshake to register probes
	fmt.Println("[MediationContainer] Registering ", len(theContainer.allProbes) , " probes")
	theContainer.theRemoteMediationClient.Init()
}

func (theContainer *MediationContainer) Close() {
	theContainer.theRemoteMediationClient.Stop()
	// TODO: clear probe map ?
}


// Get the container config from file.
func ReadConfig(path string) *ContainerConfig {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		glog.Errorf("[MediationContainer] Container conf file error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Println(string(file))
	var config ContainerConfig

	err := json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("[MediationContainer] Container conf unmarshall :%v\n", err)
	}

	glog.V(4).Infof("Results: %+v\n", config)
	return &config
}

func ReadProbeConfig(path string) *ProbeConfig {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		glog.Errorf("[MediationContainer] Probe conf file error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Println(string(file))
	var config ProbeConfig

	err := json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("[MediationContainer] Probe conf unmarshall :%v\n", err)
	}

	glog.V(4).Infof("Results: %+v\n", config)
	return &config
}

// ============================= Probe Management ==================
func (theContainer *MediationContainer) LoadProbe(probe *probe.TurboProbe) {
	// load the probe config
	config := &ProbeConfig{
		ProbeCategory: probe.ProbeCategory,
		ProbeType: probe.ProbeType,
	}

	probeProp := &ProbeProperties{
		ProbeConfig: config,
		//ProbeInterface: createProbeFunc(configFile),
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
		fmt.Println("[MediationContainer] Found " + probeProps.ProbeConfig.ProbeCategory + "::" + probeProps.ProbeConfig.ProbeType + " ==> " , acctDefProps)
		return probe
	}
	fmt.Println("[MediationContainer] Cannot find Probe of type " + probeType)
	return nil
}

