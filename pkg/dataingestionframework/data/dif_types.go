package data

import (
	"fmt"

	set "github.com/deckarep/golang-set"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// USING the common DIF Data
var DIFEntityType = map[proto.EntityDTO_EntityType]string{
	proto.EntityDTO_VIRTUAL_MACHINE:       "virtualMachine",
	proto.EntityDTO_APPLICATION_COMPONENT: "application",
	proto.EntityDTO_BUSINESS_APPLICATION:  "businessApplication",
	proto.EntityDTO_BUSINESS_TRANSACTION:  "businessTransaction",
	proto.EntityDTO_DATABASE_SERVER:       "databaseServer",
	proto.EntityDTO_SERVICE:               "service",
}

type DIFHostType string

const (
	VM        DIFHostType = "virtualMachine"
	CONTAINER DIFHostType = "container"
)

var DIFMetricType = map[proto.CommodityDTO_CommodityType]string{
	proto.CommodityDTO_RESPONSE_TIME:     "responseTime",
	proto.CommodityDTO_TRANSACTION:       "transaction",
	proto.CommodityDTO_VCPU:              "cpu",
	proto.CommodityDTO_VMEM:              "memory",
	proto.CommodityDTO_THREADS:           "threads",
	proto.CommodityDTO_HEAP:              "heap",
	proto.CommodityDTO_COLLECTION_TIME:   "collectionTime",
	proto.CommodityDTO_DB_MEM:            "dbMem",
	proto.CommodityDTO_DB_CACHE_HIT_RATE: "dbCacheHitRate",
	proto.CommodityDTO_CONNECTION:        "connection",
}

//Data ingestion framework topology entity
type DIFEntity struct {
	UID                 string                     `json:"uniqueId"`
	Type                string                     `json:"type"`
	Name                string                     `json:"name"`
	HostedOn            *DIFHostedOn               `json:"hostedOn"`
	MatchingIdentifiers *DIFMatchingIdentifiers    `json:"matchIdentifiers"`
	PartOf              []*DIFPartOf               `json:"partOf"`
	Metrics             map[string][]*DIFMetricVal `json:"metrics"`
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

func (e *DIFEntity) PartOfEntity(entity, id string) *DIFEntity {
	if e.partOfSet.Contains(id) {
		return e
	}
	e.partOfSet.Add(id)
	e.PartOf = append(e.PartOf, &DIFPartOf{entity, id})
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

func (e *DIFEntity) Matching(id string) {
	if e.MatchingIdentifiers == nil {
		e.MatchingIdentifiers = &DIFMatchingIdentifiers{id}
		return
	}
	// Overwrite
	e.MatchingIdentifiers.IPAddress = id
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
	return s
}
