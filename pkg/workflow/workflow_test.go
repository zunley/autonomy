package workflow

import (
	"github.com/zunley/autonomy/pkg/types"
	"testing"
)

func TestWorkflowRun(t *testing.T) {

	wff := &types.Workflow{
		Name:       "TestWorkflowRun",
		Schedule:   "* * 0/6 * *",
		WorkingDir: "/home",
		Steps: []types.Step{
			{
				Name:  "ls",
				Shell: "ls -l",
			},
		},
	}

	wf := NewWorkflow(wff)
	if rst, err := wf.Run(); err != nil {
		t.Errorf("%+v", *rst)
	}
}
