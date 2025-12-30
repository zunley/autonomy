package workflow

import (
	"os/exec"
	"time"

	"github.com/google/uuid"
	"github.com/zunley/autonomy/pkg/types"
)

type Workflow interface {
	Run() (*types.RunResult, error)
}

func NewWorkflow(wff *types.Workflow) Workflow {
	return &workflow{
		wff: wff,
	}
}

type workflow struct {
	wff *types.Workflow
}

func (wf *workflow) Run() (*types.RunResult, error) {
	wff := wf.wff
	startTime := time.Now()
	runID := uuid.New().String()
	workingDir := wff.WorkingDir
	stepsOutput := make([]types.StepOutput, len(wff.Steps))

	status := "success"

	for i, step := range wff.Steps {
		cmd := exec.Command("sh", "-e", "-c", step.Shell)
		cmd.Dir = workingDir
		output, err := cmd.CombinedOutput()

		stepsOutput[i].Name = step.Name
		stepsOutput[i].Shell = step.Shell
		stepsOutput[i].Output = string(output)

		if err != nil {
			status = "failed"
			break
		}

	}
	completedTime := time.Now()

	return &types.RunResult{
		WorkflowName: wff.Name,
		RunID:        runID,
		StartedAt:    startTime,
		CompletedAt:  completedTime,
		Status:       status,
		StepsOutput:  stepsOutput,
	}, nil
}
