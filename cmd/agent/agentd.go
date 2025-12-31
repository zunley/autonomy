package agent

import (
	"os"
	"github.com/zunley/autonomy/pkg/types"
)

const (
	DefaultConfigPath = "~/.autonomy/config.yaml"
)
func main() {

	var err error
	var workflow_dir string
	var config types.AgentConfig

	m := workflow.NewManager()
	_ = m.Run()
	defer m.Stop()

	// TODO 阻塞
}


func loadConfig(configPath string) (*types.AgentConfig, error) {
	data, err := os.ReadFile(configPath)
	
}