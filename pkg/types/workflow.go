package types

type Workflow struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Schedule    string `yaml:"schedule"`
	Steps       []step `yaml:"steps"`
}

type Step struct {
	Name       string `yaml:"name"`
	Shell      string `yaml:shell`
	WorkingDir string `yaml:"steps"`
}

type RunResult struct {
	AgentID      string    `json:"agent_id"`
	WorkflowName string    `json:"workflow_name"`
	RunID        string    `json:"run_id"`
	StartedAt    time.Time `json:"started_at"`
	CompletedAt  time.Time `json:"completed_at"`
	Status       string    `json:"status"` // "success", "failed"
}
