package types

import "time"

type Workflow struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Schedule    string `yaml:"schedule"`
	WorkingDir  string `yaml:"working_dir"`
	Steps       []Step `yaml:"steps"`
}

type Step struct {
	Name  string `yaml:"name"`
	Shell string `yaml:shell`
}

type RunResult struct {
	AgentID      string       `json:"agent_id"`
	WorkflowName string       `json:"workflow_name"`
	RunID        string       `json:"run_id"`
	StartedAt    time.Time    `json:"started_at"`
	CompletedAt  time.Time    `json:"completed_at"`
	Status       string       `json:"status"` // "success", "failed"
	StepsOutput  []StepOutput `json:"steps_output"`
}

type StepOutput struct {
	Step
	Output string `json:"output"`
}
