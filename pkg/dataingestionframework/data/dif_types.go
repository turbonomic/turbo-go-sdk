package data

import "fmt"

//Data injection framework topoloty entity
type DIFEntity struct {
	UID  string `json:"uniqueId"`
	Type string `json:"type"`
	Name string `json:"name"`

	HostedOn            *DIFHostedOn             `json:"hostedOn"`
	MatchingIdentifiers *DIFMatchingIdentifiers  `json:"matchIdentifiers"`
	PartOf              []*DIFPartOf             `json:"partOf"`
	Metrics             []map[string][]*DIFMetricVal`json:"metrics"`
}


type DIFMatchingIdentifiers struct {
	IPAddress string `json:"ipAddress"`
}

type DIFHostedOn struct {
	HostType  []string `json:"hostType"`
	IPAddress string   `json:"ipAddress"`
	HostUuid  string   `json:"hostUuid"`
}

type DIFPartOf struct {
	ParentEntity string `json:"entity"`
	UniqueId     string `json:"uniqueId"`
}

func DIFEntityToString(entity *DIFEntity) string {
	var s string
	s = fmt.Sprintf("[%s]%s:%s\n", entity.Type, entity.UID, entity.Name)

	if entity.MatchingIdentifiers != nil {
		s += fmt.Sprintf("		ip:%s\n", entity.MatchingIdentifiers.IPAddress)
	}

	if entity.PartOf != nil {
		s += fmt.Sprintf("	PartOf:\n")
		for _, partOf := range entity.PartOf {
			s += fmt.Sprintf("		%s:%s\n", partOf.ParentEntity, partOf.UniqueId)
		}
	}

	if entity.HostedOn != nil {
		s += fmt.Sprintf("		%s:%s\n", entity.HostedOn.HostUuid, entity.HostedOn.IPAddress)
	}

	return s
}
