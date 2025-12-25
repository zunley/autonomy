package workflow

import (
	"os/exec"
	"time"

	"github.com/google/uuid"
	"github.com/zunley/autonomy/pkg/types"
)

type Workflow interface {
	Run() (types.RunResult, error)
}

func NewWorkflow(wf *types.Workflow) Workflow {
	return nil
}

type workflow struct {
	wf *types.Workflow
}

func (w *workflow) Run() (*types.RunResult, error) {
	startTime := time.Now()
	// TODO
	runID := uuid.New().String()
	workDir := wf.WorkingDir

	status := "success"

	for i, step := range wf.Steps {
		cmd := exec.Command("sh", "-e", "-c", step.Shell)
		cmd.Dir = workDir

		output, err := cmd.Combined

		// TODO log 日志
		fmt.Println(string(output))

		if err != nil {
			status = "failed"
			break
		}
	}
	completedTime := time.Now()

	return &types.RunResult{
		AgentID:      agentID,
		WorkflowName: wf.Name,
		RunID:        runID,
		StartedAt:    startTime,
		CompletedAt:  completedTime,
		Status:       status,
	}, nil
}
