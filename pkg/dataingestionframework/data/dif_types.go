package data

import (
	"fmt"
	set "github.com/deckarep/golang-set"
)

//Data ingestion framework topology entity
type DIFEntity struct {
	UID                 string                     `json:"uniqueId"`
	Type                string                     `json:"type"`
	Name                string                     `json:"name"`
	HostedOn            *DIFHostedOn               `json:"hostedOn"`
	MatchingIdentifiers *DIFMatchingIdentifiers    `json:"matchIdentifiers"`
	PartOf              []*DIFPartOf               `json:"partOf"`
	Metrics             map[string][]*DIFMetricVal `json:"metrics"`
	namespace           string
	partOfSet           set.Set
	hostTypeSet         set.Set
}

type DIFMatchingIdentifiers struct {
	IPAddress string `json:"ipAddress"`
}

type DIFHostedOn struct {
	HostType  []DIFHostType `json:"hostType"`
	IPAddress string        `json:"ipAddress"`
	HostUuid  string        `json:"hostUuid"`
}

type DIFPartOf struct {
	ParentEntity string `json:"entity"`
	UniqueId     string `json:"uniqueId"`
	Label        string `json:"label,omitempty"`
}

func NewDIFEntity(uid, eType string) *DIFEntity {
	return &DIFEntity{
		UID:         uid,
		Type:        eType,
		Name:        uid,
		partOfSet:   set.NewSet(),
		hostTypeSet: set.NewSet(),
		Metrics:     make(map[string][]*DIFMetricVal),
	}
}

func (e *DIFEntity) WithName(name string) *DIFEntity {
	e.Name = name
	return e
}

func (e *DIFEntity) WithNamespace(namespace string) *DIFEntity {
	e.namespace = namespace
	return e
}

func (e *DIFEntity) GetNamespace() string {
	return e.namespace
}

func (e *DIFEntity) PartOfEntity(entity, id, label string) *DIFEntity {
	if e.partOfSet.Contains(id) {
		return e
	}
	e.partOfSet.Add(id)
	e.PartOf = append(e.PartOf, &DIFPartOf{entity, id, label})
	return e
}

func (e *DIFEntity) HostedOnType(hostType DIFHostType) *DIFEntity {
	if e.hostTypeSet.Contains(hostType) {
		return e
	}
	if e.HostedOn == nil {
		e.HostedOn = &DIFHostedOn{}
	}
	e.HostedOn.HostType = append(e.HostedOn.HostType, hostType)
	e.hostTypeSet.Add(hostType)
	return e
}

func (e *DIFEntity) GetHostedOnType() []DIFHostType {
	var hostTypes []DIFHostType
	for _, hostType := range e.hostTypeSet.ToSlice() {
		hostTypes = append(hostTypes, hostType.(DIFHostType))
	}
	return hostTypes
}

func (e *DIFEntity) HostedOnIP(ip string) *DIFEntity {
	if e.HostedOn == nil {
		e.HostedOn = &DIFHostedOn{}
	}
	e.HostedOn.IPAddress = ip
	return e
}

func (e *DIFEntity) HostedOnUID(uid string) *DIFEntity {
	if e.HostedOn == nil {
		e.HostedOn = &DIFHostedOn{}
	}
	e.HostedOn.HostUuid = uid
	return e
}

func (e *DIFEntity) Matching(id string) *DIFEntity {
	if e.MatchingIdentifiers == nil {
		e.MatchingIdentifiers = &DIFMatchingIdentifiers{id}
		return e
	}
	// Overwrite
	e.MatchingIdentifiers.IPAddress = id
	return e
}

func (e *DIFEntity) AddMetric(metricType string, kind DIFMetricValKind, value float64, key string) {
	var metricVal *DIFMetricVal
	var metricKey *string
	// Only set non-empty key
	if key != "" {
		metricKey = &key
	}
	meList, found := e.Metrics[metricType]
	if !found {
		// This is a new metric type
		metricVal = &DIFMetricVal{Key: metricKey}
		e.Metrics[metricType] = append(e.Metrics[metricType], metricVal)
	} else if metricKey != nil {
		// We have seen this type, check if we have a metric with the same non-empty key
		for _, me := range meList {
			if me.Key != nil && *me.Key == *metricKey {
				metricVal = me
				break
			}
		}
		if metricVal == nil {
			// This is a metric of the same type but a new key
			metricVal = &DIFMetricVal{Key: metricKey}
			e.Metrics[metricType] = append(e.Metrics[metricType], metricVal)
		}
	} else {
		if len(meList) < 1 {
			return
		}
		metricVal = meList[0]
	}
	if kind == AVERAGE {
		metricVal.Average = &value
	} else if kind == CAPACITY {
		metricVal.Capacity = &value
	}
}

func (e *DIFEntity) AddMetrics(metricType string, metricVals []*DIFMetricVal) {
	e.Metrics[metricType] = append(e.Metrics[metricType], metricVals...)
}

func (e *DIFEntity) String() string {
	s := fmt.Sprintf("%s[%s:%s]", e.Type, e.UID, e.Name)
	if e.MatchingIdentifiers != nil {
		s += fmt.Sprintf(" IP[%s]", e.MatchingIdentifiers.IPAddress)
	}
	if e.PartOf != nil {
		s += fmt.Sprintf(" PartOf")
		for _, partOf := range e.PartOf {
			s += fmt.Sprintf("[%s:%s]", partOf.ParentEntity, partOf.UniqueId)
		}
	}
	if e.HostedOn != nil {
		s += fmt.Sprintf(" HostedOn")
		s += fmt.Sprintf("[%s:%s]",
			e.HostedOn.HostUuid, e.HostedOn.IPAddress)
	}
	for metricName, metricList := range e.Metrics {
		for _, metric := range metricList {
			s += fmt.Sprintf(" Metric %s:[%v]", metricName, metric)
		}
	}
	return s
}
