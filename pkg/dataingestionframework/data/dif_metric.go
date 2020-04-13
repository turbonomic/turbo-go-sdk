package data

type DIFMetric struct {
	MetricMap map[string]*DIFMetricVal
}

type DIFMetricVal struct {
	Average     *float64       `json:"average"`
	Min         *float64       `json:"min"`
	Max         *float64       `json:"min"`
	Capacity    *float64       `json:"min"`
	Unit        *DIFMetricUnit `json:"unit"`
	Key         *string        `json:"key"`
	Description *string        `json:"description"`
	RawMetrics  interface{}    `json:"rawData"`
}

type DIFMetricUnit string

const (
	COUNT DIFMetricUnit = "count"
	TPS   DIFMetricUnit = "tps"
	MS    DIFMetricUnit = "ms"
	MB    DIFMetricUnit = "mb"
	MHZ   DIFMetricUnit = "mhz"
	PCT   DIFMetricUnit = "pct"
)

type DIFMetricValKey string

const (
	KEY         DIFMetricValKey = "key"
	DESCRIPTION DIFMetricValKey = "description"
	RAWDATA     DIFMetricValKey = "rawData"
	AVERAGE     DIFMetricValKey = "average"
	MAX         DIFMetricValKey = "max"
	MIN         DIFMetricValKey = "min"
	CAPACITY    DIFMetricValKey = "capacity"
	UNIT        DIFMetricValKey = "unit"
)

const (
	UNSET_FLOAT  = 0.0
	UNSET_STRING = ""
)
