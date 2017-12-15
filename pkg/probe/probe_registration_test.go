package probe

import (
	"github.com/stretchr/testify/assert"
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
		{full: 60, incremental: 120, performance: 300},
		{full: 30, incremental: 20, performance: 30},
	}

	for _, item := range table {
		dm := NewDiscoveryMetadata()
		dm.SetFullRediscoveryIntervalSeconds(item.full)
		dm.SetIncrementalRediscoveryIntervalSeconds(item.incremental)
		dm.SetPerformanceRediscoveryIntervalSeconds(item.performance)

		if item.full >= 60 {
			assert.EqualValues(t, item.full, dm.GetFullRediscoveryIntervalSeconds())
		} else {
			assert.EqualValues(t, 600, dm.GetFullRediscoveryIntervalSeconds())
		}

		if item.incremental >= 60 {
			assert.EqualValues(t, item.incremental, dm.GetIncrementalRediscoveryIntervalSeconds())
		} else {
			assert.EqualValues(t, -1, dm.GetIncrementalRediscoveryIntervalSeconds())
		}

		if item.performance >= 60 {
			assert.EqualValues(t, item.performance, dm.GetPerformanceRediscoveryIntervalSeconds())
		} else {
			assert.EqualValues(t, -1, dm.GetPerformanceRediscoveryIntervalSeconds())
		}
	}
}
