package workflow

import "github.com/zunley/autonomy/pkg/types"

type Workflow interface {
	Run() (types.RunResult, error)
}

func NewWorkflow() Workflow {
	return nil
}
