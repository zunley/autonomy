package agent

import "github.com/zunley/autonomy/pkg/types"

func main() {

	var err error
	var workflow_dir string
	var config types.AgentConfig

	m := workflow.NewManager()
	_ = m.Run()
	defer m.Stop()

	// TODO 阻塞
}
