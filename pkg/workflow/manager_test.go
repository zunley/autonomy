package workflow

import (
	"testing"
	"github.com/zunley/autonomy/pkg/types"
)

func TestManagerAddWorkflow(t *testing.T) {

	wff := &types.Workflow{
		Name: "TestWorkflowRun",
		Schedule: "* * 0/6 * *",
		WorkingDir: "/home",
		Steps: []types.Step{
			{
				Name: "ls",
				Shell: "ls -l",
			},
		},
	}

	wfm := NewManager()
	if err := wfm.AddWorkflow(wff); err != nil {
		t.Errorf("%s", err.Error())
	}
}
