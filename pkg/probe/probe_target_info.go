package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type ProbeTargetInfo struct {
	communicationBindingChannel string
	*TurboTargetInfo
}

// Build the default target instance based on information from ProbeTargetInfo.
func (targetInfo *ProbeTargetInfo) GetTargetInstance() *proto.ProbeTargetInfo {
	//add target identifier?
	return &proto.ProbeTargetInfo{
		InputValues: targetInfo.accountValues,
		CommunicationBindingChannel: &targetInfo.communicationBindingChannel,
	}
}