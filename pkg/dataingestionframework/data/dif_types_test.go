package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddMetricWithNoKey(t *testing.T) {
	average := float64(300)
	capacity := float64(500)
	mType := "responseTime"
	difEntity := NewDIFEntity("uid", "application")
	difEntity.AddMetric(mType, AVERAGE, average, "")
	difEntity.AddMetric(mType, CAPACITY, capacity, "")
	metricMap := difEntity.Metrics
	assert.Equal(t, 1, len(metricMap))
	metrics, _ := metricMap[mType]
	assert.Equal(t, 1, len(metrics))
	metric := metrics[0]
	assert.Equal(t, *metric.Average, average)
	assert.Equal(t, *metric.Capacity, capacity)
	assert.Nil(t, metric.Key)
}

func TestAddMetricWithSameKey(t *testing.T) {
	average := float64(300)
	capacity := float64(500)
	mType := "kpi"
	key := "key"
	difEntity := NewDIFEntity("uid", "application")
	difEntity.AddMetric(mType, AVERAGE, average, key)
	difEntity.AddMetric(mType, CAPACITY, capacity, key)
	metricMap := difEntity.Metrics
	assert.Equal(t, 1, len(metricMap))
	metrics, _ := metricMap[mType]
	assert.Equal(t, 1, len(metrics))
	metric := metrics[0]
	assert.Equal(t, *metric.Average, average)
	assert.Equal(t, *metric.Capacity, capacity)
	assert.EqualValues(t, *metric.Key, key)
}

func TestAddMetricWithDifferentKey(t *testing.T) {
	average1 := float64(300)
	average2 := float64(500)
	mType := "kpi"
	key1 := "key1"
	key2 := "key2"
	difEntity := NewDIFEntity("uid", "application")
	difEntity.AddMetric(mType, AVERAGE, average1, key1)
	difEntity.AddMetric(mType, AVERAGE, average2, key2)
	metricMap := difEntity.Metrics
	assert.Equal(t, 1, len(metricMap))
	metrics, _ := metricMap[mType]
	assert.Equal(t, 2, len(metrics))
	total := float64(0)
	for _, metric := range metrics {
		total += *metric.Average
	}
	assert.EqualValues(t, average1+average2, total)
}

func TestAddMetricWithOverwritingValues(t *testing.T) {
	average1 := float64(300)
	average2 := float64(500)
	average3 := float64(777)
	capacity3 := float64(10000)
	mType := "kpi"
	key1 := "key1"
	key2 := "key2"
	key3 := key2
	difEntity := NewDIFEntity("uid", "application")
	difEntity.AddMetric(mType, AVERAGE, average1, key1)
	difEntity.AddMetric(mType, AVERAGE, average2, key2)
	// average3 should overwrite average2 for key2
	difEntity.AddMetric(mType, AVERAGE, average3, key3)
	difEntity.AddMetric(mType, CAPACITY, capacity3, key3)
	metricMap := difEntity.Metrics
	assert.Equal(t, 1, len(metricMap))
	metrics, _ := metricMap[mType]
	assert.Equal(t, 2, len(metrics))
	assert.EqualValues(t, average1, *metrics[0].Average)
	assert.Equal(t, key1, *metrics[0].Key)
	assert.EqualValues(t, average3, *metrics[1].Average)
	assert.EqualValues(t, capacity3, *metrics[1].Capacity)
	assert.Equal(t, key2, *metrics[1].Key)
}

func TestAddMetricWithKeyAndNoKey(t *testing.T) {
	// This should be an edge case
	average1 := float64(300)
	average2 := float64(500)
	average3 := float64(777)
	capacity3 := float64(10000)
	mType := "kpi"
	key1 := "key1"
	key2 := ""
	difEntity := NewDIFEntity("uid", "application")
	difEntity.AddMetric(mType, AVERAGE, average1, key1)
	difEntity.AddMetric(mType, AVERAGE, average2, key2)
	difEntity.AddMetric(mType, AVERAGE, average3, key2)
	difEntity.AddMetric(mType, CAPACITY, capacity3, key2)
	metricMap := difEntity.Metrics
	assert.Equal(t, 1, len(metricMap))
	metrics, _ := metricMap[mType]
	assert.Equal(t, 2, len(metrics))
	assert.EqualValues(t, average1, *metrics[0].Average)
	assert.Equal(t, key1, *metrics[0].Key)
	assert.EqualValues(t, average3, *metrics[1].Average)
	assert.EqualValues(t, capacity3, *metrics[1].Capacity)
	assert.Nil(t, metrics[1].Key)
}
