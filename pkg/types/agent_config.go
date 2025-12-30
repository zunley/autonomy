// pkg/types/config.go
package types

type AgentConfig struct {
	AgentID     string       `yaml:"agent_id"`
	ControlNode *ControlNode `yaml:"control_node,omitempty"`
	LogLevel    string       `yaml:"log_level,omitempty"` // "debug", "info", "warn"
}

type ControlNode struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}
