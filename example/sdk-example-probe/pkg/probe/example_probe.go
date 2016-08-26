package probe

import (
	"github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

type ExampleProbe struct{}

func NewExampleProbe() *ExampleProbe {
	return &ExampleProbe{}
}

func (this *ExampleProbe) Discover() ([]*proto.EntityDTO, error) {
	return nil, nil
}
