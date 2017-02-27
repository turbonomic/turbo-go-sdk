package probe

import (
	"errors"
	"github.com/golang/glog"
)

type ProbeBuilder struct {
	probeType          string
	probeCategory      string
	registrationClient TurboRegistrationClient
	discoveryClientMap map[string]TurboDiscoveryClient
	actionClient	   TurboActionExecutorClient
	builderError       error
}

func ErrorInvalidTargetIdentifier() error {
	return errors.New("Null Target Identifier")
}

func ErrorInvalidProbeType() error {
	return errors.New("Null Probe type")
}

func ErrorInvalidProbeCategory() error {
	return errors.New("Null Probe category")
}

func ErrorInvalidRegistrationClient() error {
	return errors.New("Null registration client")
}
func ErrorInvalidActionClient() error {
	return errors.New("Null action client")
}

func ErrorInvalidDiscoveryClient(targetId string) error {
	return errors.New("Invalid discovery client for target [" + targetId + "]")
}

func ErrorUndefinedDiscoveryClient() error {
	return errors.New("No discovery clients defined")
}

func ErrorCreatingProbe(probeType string, probeCategory string) error {
	return errors.New("Error creating probe for " + probeCategory + "::" + probeType)
}

// Get an instance of ProbeBuilder
func NewProbeBuilder(probeType string, probeCategory string) *ProbeBuilder {
	// Validate probe type and category
	probeBuilder := &ProbeBuilder{}
	if probeType == "" {
		probeBuilder.builderError = ErrorInvalidProbeType()
		return probeBuilder
	}

	if probeCategory == "" {
		probeBuilder.builderError = ErrorInvalidProbeCategory()
		return probeBuilder
	}

	return &ProbeBuilder{
		probeCategory:      probeCategory,
		probeType:          probeType,
		discoveryClientMap: make(map[string]TurboDiscoveryClient),
	}
}

// Build an instance of TurboProbe.
func (pb *ProbeBuilder) Create() (*TurboProbe, error) {
	if pb.builderError != nil {
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}

	if len(pb.discoveryClientMap) == 0 {
		pb.builderError = ErrorUndefinedDiscoveryClient()
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}

	probeConf := &ProbeConfig{
		ProbeCategory: pb.probeCategory,
		ProbeType:     pb.probeType,
	}
	turboProbe, err := newTurboProbe(probeConf)
	if err != nil {
		pb.builderError = ErrorCreatingProbe(pb.probeType, pb.probeCategory)
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}

	turboProbe.RegistrationClient = pb.registrationClient
	turboProbe.ActionClient = pb.actionClient
	for targetId, discoveryClient := range pb.discoveryClientMap {
		turboProbe.DiscoveryClientMap[targetId] = discoveryClient
	}

	return turboProbe, nil
}

// Set the registration client for the probe
func (pb *ProbeBuilder) RegisteredBy(registrationClient TurboRegistrationClient) *ProbeBuilder {
	if registrationClient == nil {
		pb.builderError = ErrorInvalidRegistrationClient()
		return pb
	}
	pb.registrationClient = registrationClient

	return pb
}

// Set a target and discovery client for the probe
func (pb *ProbeBuilder) DiscoversTarget(targetId string, discoveryClient TurboDiscoveryClient) *ProbeBuilder {
	if targetId == "" {
		pb.builderError = ErrorInvalidTargetIdentifier()
		return pb
	}
	if discoveryClient == nil {
		pb.builderError = ErrorInvalidDiscoveryClient(targetId)
		return pb
	}

	pb.discoveryClientMap[targetId] = discoveryClient

	return pb
}

// Set the action client for the probe
func (pb *ProbeBuilder) ExecutesActionsBy(actionClient TurboActionExecutorClient) *ProbeBuilder {
	if actionClient == nil {
		pb.builderError = ErrorInvalidActionClient()
		return pb
	}
	pb.actionClient = actionClient

	return pb
}