package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

// TODO:
type IActionExecutor interface {
	executeAction(actionExecutionDTO proto.ActionExecutionDTO,
		accountValues []*proto.AccountValue,
		progressTracker IProgressTracker)
}

type IProgressTracker interface {
}
