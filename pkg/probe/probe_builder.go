package probe

import "fmt"

type ProbeBuilder struct {
	turboProbe *TurboProbe
}

// Get an instance of ProbeBuilder
func NewProbeBuilder (probeType string, probeCategory string) *ProbeBuilder {
	// Validate probe type and category
	if &probeType == nil {
		fmt.Println("[ProbeBuilder] Null Probe type")	//TODO: throw exception
		return nil
	}

	if &probeCategory == nil {
		fmt.Println("[ProbeBuilder] Null probe category")	//TODO: throw exception
		return nil
	}

	probeConf := &ProbeConfig {
		ProbeCategory: probeCategory,
		ProbeType: probeType,
	}
	probe := NewTurboProbe(probeConf)

	return &ProbeBuilder{
		turboProbe: probe,
	}
}

// Build an instance of TurboProbe.
func (pb *ProbeBuilder) Create() *TurboProbe {
	if pb.turboProbe.RegistrationClient == nil {
		fmt.Println("[ProbeBuilder] Null registration client")	//TODO: throw exception
		return nil
	}

	return pb.turboProbe
}

// Build an instance of TurboProbe.
func (pb *ProbeBuilder) RegisteredBy(registrationClient TurboRegistrationClient) *ProbeBuilder {
	pb.turboProbe.SetProbeRegistrationClient(registrationClient)
	return pb
}

// Build an instance of TurboProbe.
func (pb *ProbeBuilder) DiscoversTarget(targetId string, discoveryClient TurboDiscoveryClient) *ProbeBuilder {
	pb.turboProbe.SetDiscoveryClient(targetId, discoveryClient)
	return pb
}
