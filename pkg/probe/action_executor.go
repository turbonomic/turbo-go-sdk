package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// TODO:
type ActionExecutorClient interface {
	executeAction(actionExecutionDTO proto.ActionExecutionDTO,
		accountValues []*proto.AccountValue,
		progressTracker ActionProgressTracker)
}

type ActionProgressTracker interface {
}
