package data

type CDPMetric struct {
	MetricMap map[string]*CDPMetricVal
}

type CDPMetricVal struct {
	Average     float64       `json:"average"`
	Min         float64       `json:"min"`
	Max         float64       `json:"min"`
	Capacity    float64       `json:"min"`
	Unit        CDPMetricUnit `json:"unit"`
	Key         string        `json:"key"`
	Description string        `json:"description"`
	RawMetrics  interface{}   `json:"rawData"`
}

type CDPMetricUnit string

const (
	COUNT CDPMetricUnit = "count"
	TPS   CDPMetricUnit = "tps"
	MS    CDPMetricUnit = "ms"
	MB    CDPMetricUnit = "mb"
	MHZ   CDPMetricUnit = "mhz"
	PCT   CDPMetricUnit = "pct"
)

type CDPMetricValKey string

const (
	KEY         CDPMetricValKey = "key"
	DESCRIPTION CDPMetricValKey = "description"
	RAWDATA     CDPMetricValKey = "rawData"
	AVERAGE     CDPMetricValKey = "average"
	MAX         CDPMetricValKey = "max"
	MIN         CDPMetricValKey = "min"
	CAPACITY    CDPMetricValKey = "capacity"
	UNIT        CDPMetricValKey = "unit"
)

const (
	UNSET_FLOAT  = -100.0
	UNSET_STRING = ""
)
