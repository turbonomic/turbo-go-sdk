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
}
