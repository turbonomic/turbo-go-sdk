package data

import "time"

//Data injection framework topology
type Topology struct {
	Version    string       `json:"version"`
	Updatetime int64        `json:"updateTime"`
	Scope      string       `json:"scope"`
	Entities   []*DIFEntity `json:"topology"`
}

func NewTopology() *Topology {
	return &Topology{
		Version:    "v1",
		Updatetime: 0,
		Entities:   []*DIFEntity{},
		Scope:      "",
	}
}

func (r *Topology) SetUpdateTime() {
	t := time.Now()
	r.Updatetime = t.Unix()
}
