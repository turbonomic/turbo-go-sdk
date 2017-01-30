package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// Abstraction for the VMTTarget object in the client
// TODO: reconcile with the definition in the turbo-api
type TurboTarget struct {
	targetType string
	targetIdentifier string
	user string
	password string
	AccountValues[] *proto.AccountValue
}

func (target *TurboTarget) GetTargetType() string {
	return target.targetType
}

func (target *TurboTarget) GetTargetId() string {
	return target.targetIdentifier
}

func (target *TurboTarget) GetNameOrAddress() string {
	return target.targetIdentifier
}

func (target *TurboTarget) GetUser() string {
	return target.user
}

func (target *TurboTarget) GetPassword() string {
	return target.password
}

func (target *TurboTarget) SetUser(user string) {
	target.user = user
}

func (target *TurboTarget) SetPassword(pwd string) {
	target.password = pwd
}




