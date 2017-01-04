package registration

import (
	comm "github.com/turbonomic/turbo-go-sdk/pkg/communication"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"github.com/turbonomic/turbo-go-sdk/example/sdk-example-probe/pkg/registration/supplychain"

	"github.com/golang/glog"
)

type MediationContainerInfoBuilder struct {
	targetType string
}

func NewMediationContainerInfoBuilder(targetType string) *MediationContainerInfoBuilder {
	return &MediationContainerInfoBuilder{
		targetType: targetType,
	}
}

func (this *MediationContainerInfoBuilder) Build() (*proto.ContainerInfo, error) {
	p, err := buildProbes(this.targetType)
	if err != nil {
		return nil, err
	}

	return &proto.ContainerInfo{
		Probes: p,
	}, nil
}

func createAccountDef() []*proto.AccountDefEntry {
	var acctDefProps []*proto.AccountDefEntry

	// target id
	targetIDAcctDefEntry := comm.NewAccountDefEntryBuilder("targetIdentifier", "Address",
		"IP address of the probe", ".*", proto.AccountDefEntry_OPTIONAL, false).Create()
	acctDefProps = append(acctDefProps, targetIDAcctDefEntry)

	// username
	usernameAcctDefEntry := comm.NewAccountDefEntryBuilder("username", "Username",
		"Username of the probe", ".*", proto.AccountDefEntry_OPTIONAL, false).Create()
	acctDefProps = append(acctDefProps, usernameAcctDefEntry)

	// password
	passwdAcctDefEntry := comm.NewAccountDefEntryBuilder("password", "Password",
		"Password of the probe", ".*", proto.AccountDefEntry_OPTIONAL, true).Create()
	acctDefProps = append(acctDefProps, passwdAcctDefEntry)

	return acctDefProps
}

func buildProbes(probeType string) ([]*proto.ProbeInfo, error) {
	// 1. Construct the account definition for exmaple probe.
	acctDefProps := createAccountDef()

	// 2. Build supply chain.
	supplyChainFactory := &supplychain.SupplyChainFactory{}
	templateDtos, err := supplyChainFactory.CreateSupplyChain()
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("Supply chain for the example probe is created.")

	// 3. construct the example probe info.
	probeCat := "Container"
	exampleProbe := comm.NewProbeInfoBuilder(probeType, probeCat, templateDtos, acctDefProps).Create()

	// 4. Add example probe to probeInfo list, here it is the only probe supported.
	var probes []*proto.ProbeInfo
	probes = append(probes, exampleProbe)

	return probes, nil
}
