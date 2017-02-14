package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type TestProbe struct{}
type TestProbeDiscoveryClient struct {}
type TestProbeRegistrationClient struct {}
func (handler *TestProbeDiscoveryClient) GetAccountValues() *TurboTarget {
	return nil
}
func (handler *TestProbeDiscoveryClient) Validate(accountValues[] *proto.AccountValue) *proto.ValidationResponse {
	return nil
}

func (handler *TestProbeDiscoveryClient) Discover(accountValues[] *proto.AccountValue) *proto.DiscoveryResponse {
	return nil
}

func (registrationClient *TestProbeRegistrationClient) GetSupplyChainDefinition() []*proto.TemplateDTO {
	return nil
}
func (registrationClient *TestProbeRegistrationClient) GetAccountDefinition() []*proto.AccountDefEntry {
	return nil
}
func (registrationClient *TestProbeRegistrationClient) GetIdentifyingFields() string {
	return ""
}

