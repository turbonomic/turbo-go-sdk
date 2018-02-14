package probe

import (
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg"
	"testing"
)

func TestInvalidProbeConf(t *testing.T) {
	pc, err := NewProbeConfig("", "")
	assert.NotNil(t, err)
	assert.Nil(t, pc)

	pc, err = NewProbeConfig("Type1", "")
	assert.NotNil(t, err)
	assert.Nil(t, pc)

	pc, err = NewProbeConfig("", "Category1")
	assert.NotNil(t, err)
	assert.Nil(t, pc)
}

func TestProbeConf(t *testing.T) {
	pc, err := NewProbeConfig("Type1", "Category1")
	assert.Nil(t, err)
	assert.NotNil(t, pc)
	assert.EqualValues(t, "Type1", pc.ProbeType)
	assert.EqualValues(t, "Category1", pc.ProbeCategory)
	assert.NotNil(t, pc.discoveryMetadata)
}

func TestValidateProbeConfInvalid(t *testing.T) {
	pc := &ProbeConfig{}

	err := pc.Validate()
	assert.NotNil(t, err)

	pc = &ProbeConfig{
		ProbeCategory: "Category1",
	}
	err = pc.Validate()
	assert.NotNil(t, err)

	pc = &ProbeConfig{
		ProbeType: "Type1",
	}
	err = pc.Validate()
	assert.NotNil(t, err)
}

func TestValidateProbeConf(t *testing.T) {
	pc := &ProbeConfig{
		ProbeCategory: "Category1",
		ProbeType:     "Type1",
	}
	err := pc.Validate()
	assert.Nil(t, err)
	assert.EqualValues(t, "Type1", pc.ProbeType)
	assert.EqualValues(t, "Category1", pc.ProbeCategory)
	assert.NotNil(t, pc.discoveryMetadata)
}

func TestSetDiscoveryMetadata(t *testing.T) {
	pc, _ := NewProbeConfig("Type1", "Category1")
	assert.NotNil(t, pc.discoveryMetadata)

	dm := pc.discoveryMetadata
	assert.EqualValues(t, pkg.DEFAULT_FULL_DISCOVERY_IN_SECS, pc.discoveryMetadata.fullDiscovery)
	assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED, pc.discoveryMetadata.incrementalDiscovery)
	assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED, pc.discoveryMetadata.performanceDiscovery)

	dm.SetFullRediscoveryIntervalSeconds(900)
	pc.SetDiscoveryMetadata(dm)
	assert.EqualValues(t, 900, pc.discoveryMetadata.fullDiscovery)

	dm.SetIncrementalRediscoveryIntervalSeconds(10)
	dm.SetPerformanceRediscoveryIntervalSeconds(10)
	pc.SetDiscoveryMetadata(dm)

	assert.EqualValues(t, pkg.DEFAULT_MIN_DISCOVERY_IN_SECS, pc.discoveryMetadata.incrementalDiscovery)
	assert.EqualValues(t, pkg.DEFAULT_MIN_DISCOVERY_IN_SECS, pc.discoveryMetadata.performanceDiscovery)

	dm.SetFullRediscoveryIntervalSeconds(100)
	dm.SetPerformanceRediscoveryIntervalSeconds(60)
	dm.SetPerformanceRediscoveryIntervalSeconds(90)
	pc.SetDiscoveryMetadata(dm)
	assert.EqualValues(t, 100, pc.discoveryMetadata.fullDiscovery)
	assert.EqualValues(t, 60, pc.discoveryMetadata.incrementalDiscovery)
	assert.EqualValues(t, 90, pc.discoveryMetadata.performanceDiscovery)
}
