package mediationcontainerinfobuilder

import (
	"github.com/vmturbo/vmturbo-go-sdk/pkg/comm"
	sdkproto "github.com/vmturbo/vmturbo-go-sdk/pkg/proto"
)

type MediationContainerInfoBuilder struct {
	targetType string
}

func NewMediationContainerInfoBuilder(targetType string) {
	return &MediationContainerInfoBuilder{
		targetType: targetType,
	}
}

func (this *MediationContainerInfoBuilder) Build() *sdkproto.ContainerInfo {
	return &sdkproto.ContainerInfo{
		Probes: buildProbes(this.targetType),
	}
}

func createAccountDef() []*sdkproto.AccountDefEntry {
	var acctDefProps []*sdkproto.AccountDefEntry

	// target id
	targetIDAcctDefEntry := comm.NewAccountDefEntryBuilder("targetIdentifier", "Address",
		"IP address of the probe", ".*", sdkproto.AccountDefEntry_OPTIONAL, false).Create()
	acctDefProps = append(acctDefProps, targetIDAcctDefEntry)

	// username
	usernameAcctDefEntry := comm.NewAccountDefEntryBuilder("username", "Username",
		"Username of the probe", ".*", sdkproto.AccountDefEntry_OPTIONAL, false).Create()
	acctDefProps = append(acctDefProps, usernameAcctDefEntry)

	// password
	passwdAcctDefEntry := comm.NewAccountDefEntryBuilder("password", "Password",
		"Password of the probe", ".*", comm.AccountDefEntry_OPTIONAL, true).Create()
	acctDefProps = append(acctDefProps, passwdAcctDefEntry)

	return acctDefProps
}

func buildProbes(probeType string) []*sdkproto.ProbeInfo {
	// 1. Construct the account definition for exmaple probe.
	acctDefProps := this.createAccountDef()

	// 2. Build supply chain.
	supplyChainFactory := &SupplyChainFactory{}
	templateDtos := supplyChainFactory.CreateSupplyChain()
	glog.V(2).Infof("Supply chain for the example probe is created.")

	// 3. construct the example probe info.
	probeCat := "Container"
	exampleProbe := comm.NewProbeInfoBuilder(probeType, probeCat, templateDtos, acctDefProps).Create()

	// 4. Add example probe to probeInfo list, here it is the only probe supported.
	var probes []*sdkproto.ProbeInfo
	probes = append(probes, exampleProbe)

	return probes
}
