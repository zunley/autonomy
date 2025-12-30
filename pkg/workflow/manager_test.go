package workflow

import (
	"github.com/zunley/autonomy/pkg/types"
	"testing"
)

func TestManagerAddWorkflow(t *testing.T) {

	wff := &types.Workflow{
		Name:       "TestWorkflowRun",
		Schedule:   "@every 2s",
		WorkingDir: "/home",
		Steps: []types.Step{
			{
				Name:  "ls",
				Shell: "ls -l",
			},
		},
	}

	agentConfig := &types.AgentConfig{
		AgentID: "macbook",
		ControlNode: &types.ControlNode{
			URL:   "https://127.0.0.1/api/v1/uploader",
			Token: "a1b2c3d4",
		},
	}
	wfm := NewManager(agentConfig)
	wfm.Run()
	if err := wfm.AddWorkflow(wff); err != nil {
		t.Errorf("%s", err.Error())
	}

	c := make(chan int)
	c <- 1
}
