package probe

import (
	"github.com/golang/glog"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

const (
	DEFAULT_FULL_DISCOVERY_IN_SECS int32 = 600
	DEFAULT_MIN_DISCOVERY_IN_SECS  int32 = 60
	DISCOVERY_NOT_SUPPORTED        int32 = -1
)

// ISupplyChainProvider provides the entities defined in the supply chain for a probe
type ISupplyChainProvider interface {
	GetSupplyChainDefinition() []*proto.TemplateDTO
}

// IAccountDefinitionProvider provides the definitions used to uniquely create a target
// for the probe
type IAccountDefinitionProvider interface {
	GetAccountDefinition() []*proto.AccountDefEntry
	GetIdentifyingFields() string
}

// IFullDiscoveryMetadata specifies the interval at which discoveries will be executed
// for the probe that supplies this interface during registration.
// The value is specified in seconds. If the interface implementation is not provided, a default
// of 600 seconds (10 minutes) will be used.
// The minimum value allowed for this field is 60 seconds (1 minute).
type IFullDiscoveryMetadata interface {
	GetFullRediscoveryIntervalSeconds() int32
}

// IIncrementalDiscoveryMetadata specifies the interval at which incremental discoveries
// will be executed for the probe that supplies this interface during registration.
// The value is specified in seconds. If the interface implementation is not provided,
// the probe does not support incremental discovery.
type IIncrementalDiscoveryMetadata interface {
	GetIncrementalRediscoveryIntervalSeconds() int32
}

// IPerformanceDiscoveryMetadata specifies the interval at which performance discoveries
// will be executed for the probe that supplues this interface during registration.
// The value is specified in seconds. If the interface implementation is not provided,
// the probe does not support performance discovery
type IPerformanceDiscoveryMetadata interface {
	GetPerformanceRediscoveryIntervalSeconds() int32
}

// Implementation for the providing the full, incremental, performance
// discovery intervals for the probe.
type DiscoveryMetadata struct {
	fullDiscovery        int32
	incrementalDiscovery int32
	performanceDiscovery int32
}

// Create a DiscoveryMetadata structure with default values for the different discovery intervals.
// Full discovery default is set to 600 seconds which means full discovery will occur every 10 minutes.
// Incremental discovery is set to -1 implies that incremental discovery is not supported by the probe.
// Performance discovery is set to -1 implies that the performance discovery is not supported by the probe.
func NewDiscoveryMetadata() *DiscoveryMetadata {
	return &DiscoveryMetadata{
		fullDiscovery:        DEFAULT_FULL_DISCOVERY_IN_SECS,
		incrementalDiscovery: DISCOVERY_NOT_SUPPORTED,
		performanceDiscovery: DISCOVERY_NOT_SUPPORTED,
	}
}

func (dMetadata *DiscoveryMetadata) GetFullRediscoveryIntervalSeconds() int32 {
	return dMetadata.fullDiscovery
}

func (dMetadata *DiscoveryMetadata) GetIncrementalRediscoveryIntervalSeconds() int32 {
	return dMetadata.incrementalDiscovery
}

func (dMetadata *DiscoveryMetadata) GetPerformanceRediscoveryIntervalSeconds() int32 {
	return dMetadata.performanceDiscovery
}

// Set the discovery interval at which the incremental discovery request will be sent to the probe
func (dMetadata *DiscoveryMetadata) SetIncrementalRediscoveryIntervalSeconds(incrementalDiscovery int32) {
	if incrementalDiscovery >= DEFAULT_MIN_DISCOVERY_IN_SECS {
		dMetadata.incrementalDiscovery = incrementalDiscovery
	} else {
		glog.Errorf("Invalid incremental discovery interval %d", incrementalDiscovery)
	}
}

// Set the discovery interval at which the performance discovery request will be sent to the probe
func (dMetadata *DiscoveryMetadata) SetPerformanceRediscoveryIntervalSeconds(performanceDiscovery int32) {
	if performanceDiscovery >= DEFAULT_MIN_DISCOVERY_IN_SECS {
		dMetadata.performanceDiscovery = performanceDiscovery
	} else {
		glog.Errorf("Invalid performance discovery interval %d", performanceDiscovery)
	}
}

// Set the discovery interval at which the full discovery request will be sent to the probe
func (dMetadata *DiscoveryMetadata) SetFullRediscoveryIntervalSeconds(fullDiscovery int32) {
	if fullDiscovery >= DEFAULT_MIN_DISCOVERY_IN_SECS {
		dMetadata.fullDiscovery = fullDiscovery
	} else {
		glog.Errorf("Invalid  discovery interval %d", fullDiscovery)
	}
}

// IActionPolicyProvider provides the policies for action types supported for
// different entity types by the probe
type IActionPolicyProvider interface {
	GetActionPolicy() []*proto.ActionPolicyDTO
}

// IEntityMetadataProvider provides the metadata used to generate the unique identifier for
// entities discovered by a probe
type IEntityMetadataProvider interface {
	GetEntityMetadata() []*proto.EntityIdentityMetadata
}
