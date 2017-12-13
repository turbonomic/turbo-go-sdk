package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/golang/glog"
)

type ISupplyChainProvider interface {
	GetSupplyChainDefinition() []*proto.TemplateDTO
}

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

// Default implementation for the providing the full, incremental, performance
// discovery intervals for the probe.
type DiscoveryMetadata struct {
	fullDiscovery int32
	incrementalDiscovery int32
	performanceDiscovery int32
}

func NewDiscoveryMetadata() *DiscoveryMetadata {
	return &DiscoveryMetadata{
		fullDiscovery: 600,
		incrementalDiscovery: -1,
		performanceDiscovery: -1,
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

func (dMetadata *DiscoveryMetadata) SetIncrementalRediscoveryIntervalSeconds(incrementalDiscovery int32) {
	if incrementalDiscovery >= 60 {
		dMetadata.incrementalDiscovery = incrementalDiscovery
	} else {
		glog.Errorf("Invalid discovery interval %d", incrementalDiscovery)
	}
}

func (dMetadata *DiscoveryMetadata) SetPerformanceRediscoveryIntervalSeconds(performanceDiscovery int32) {
	if performanceDiscovery > -1 {
		dMetadata.fullDiscovery = performanceDiscovery
	} else {
		glog.Errorf("Invalid performance discovery interval %d", performanceDiscovery)
	}
}

func (dMetadata *DiscoveryMetadata) SetFullRediscoveryIntervalSeconds(fullDiscovery int32) {
	if fullDiscovery > -1 {
		dMetadata.fullDiscovery = fullDiscovery
	} else {
		glog.Errorf("Invalid incremental discovery interval %d", fullDiscovery)
	}
}

type IActionPolicyProvider interface {
	GetActionPolicy() []*proto.ActionPolicyDTO
}

type IEntityMetadataProvider interface {
	GetEntityMetadata() []*proto.EntityIdentityMetadata
}


type IIncrementalDiscovery interface {
	DiscoverIncremental(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error)
}

type IPerformanceDiscovery interface {
	DiscoverPerformance(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error)
}

