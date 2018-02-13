package probe

import (
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg"
	"testing"
)

func TestNewDiscoveryMetadata(t *testing.T) {
	dm := NewDiscoveryMetadata()
	assert.EqualValues(t, 600, dm.GetFullRediscoveryIntervalSeconds())
	assert.EqualValues(t, -1, dm.GetIncrementalRediscoveryIntervalSeconds())
	assert.EqualValues(t, -1, dm.GetPerformanceRediscoveryIntervalSeconds())
}

func TestSetDicoveryIntervals(t *testing.T) {
	table := []struct {
		full        int32
		incremental int32
		performance int32
	}{
		{full: 0, incremental: 0, performance: 0},
		{full: -1, incremental: -1, performance: -1},
		{full: 60, incremental: 120, performance: 300},
		{full: 30, incremental: 20, performance: 30},
		{full: 30},
		{full: -1},
		{full: 1200},
		{incremental: 20, performance: 30},
		{incremental: 60, performance: 60},
	}

	for _, item := range table {
		dm := NewDiscoveryMetadata()
		dm.SetFullRediscoveryIntervalSeconds(item.full)
		dm.SetIncrementalRediscoveryIntervalSeconds(item.incremental)
		dm.SetPerformanceRediscoveryIntervalSeconds(item.performance)
		checkDiscoveryMetadata(t, item.full, dm, pkg.FULL_DISCOVERY)
		checkDiscoveryMetadata(t, item.incremental, dm, pkg.INCREMENTAL_DISCOVERY)
		checkDiscoveryMetadata(t, item.performance, dm, pkg.PERFORMANCE_DISCOVERY)
	}
}

func checkDiscoveryMetadata(t *testing.T, expected int32, dm *DiscoveryMetadata, discoveryType pkg.DiscoveryType) {
	actual := func(dm *DiscoveryMetadata, discoveryType pkg.DiscoveryType) int32 {
		if discoveryType == pkg.FULL_DISCOVERY {
			return dm.GetFullRediscoveryIntervalSeconds()
		} else if discoveryType == pkg.INCREMENTAL_DISCOVERY {
			return dm.GetIncrementalRediscoveryIntervalSeconds()
		} else if discoveryType == pkg.PERFORMANCE_DISCOVERY {
			return dm.GetPerformanceRediscoveryIntervalSeconds()
		}
		return 0
	}

	if discoveryType == pkg.FULL_DISCOVERY {
		if expected <= 0 {
			assert.EqualValues(t, pkg.DEFAULT_FULL_DISCOVERY_IN_SECS, actual(dm, discoveryType))
		} else if expected >= 60 {
			assert.EqualValues(t, expected, actual(dm, discoveryType))
		} else {
			assert.EqualValues(t, pkg.DEFAULT_MIN_DISCOVERY_IN_SECS, actual(dm, discoveryType))
		}
	}

	if discoveryType == pkg.INCREMENTAL_DISCOVERY || discoveryType == pkg.PERFORMANCE_DISCOVERY {
		if expected <= 0 {
			assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED, actual(dm, discoveryType))
		} else if expected >= 60 {
			assert.EqualValues(t, expected, actual(dm, discoveryType))
		} else {
			assert.EqualValues(t, pkg.DEFAULT_MIN_DISCOVERY_IN_SECS, actual(dm, discoveryType))
		}
	}
}
