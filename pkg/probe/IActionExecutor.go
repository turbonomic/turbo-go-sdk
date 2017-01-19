package probe

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)
type IActionExecutor interface {
	executeAction(actionExecutionDTO proto.ActionExecutionDTO,
	                accountValues[] *proto.AccountValue,
			progressTracker IProgressTracker)
}

type IProgressTracker interface {

}


