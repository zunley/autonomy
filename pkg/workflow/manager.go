package workflow

import (
	"github.com/zunley/autonomy/pkg/types"
	"log"
)

type Manager interface {
	Run() error
	Stop()
	AddWorkflow(*types.Workflow) error
	RemoveWorkflow(name string) error
}

func NewManager(agentConfig *types.AgentConfig) Manager {

	scheduler := NewCronScheduler()
	uploader := NewHTTPUploader(agentConfig.ControlNode)
	return &workflowManager{
		scheduler: scheduler,
		uploader:  uploader,
		workflows: make(map[string]Workflow),
	}
}

type workflowManager struct {
	scheduler Scheduler
	// map[WorkflowName]Workflow
	workflows map[string]Workflow
	uploader  Uploader
}

func (wfm *workflowManager) Run() error {
	if wfm.scheduler != nil {
		wfm.scheduler.Start()
	}
	return nil
}

func (wfm *workflowManager) Stop() {
	if wfm.scheduler != nil {
		wfm.scheduler.Stop()
	}
}

func (wfm *workflowManager) AddWorkflow(wff *types.Workflow) error {
	if _, ok := wfm.workflows[wff.Name]; ok {
		return nil
	}

	wf := NewWorkflow(wff)
	if err := wfm.scheduler.Schedule(wff.Name, wff.Schedule, func() {
		rst, err := wf.Run()
		log.Printf("%+v\n", rst)
		if err != nil {
			log.Printf("Error RunWorkflow %s: %s", wff.Name, err.Error())
		}
		if err := wfm.uploader.Upload(rst); err != nil {
			log.Printf("Error Upload: %s", err.Error())
		}
	}); err != nil {
		return err
	}
	wfm.workflows[wff.Name] = wf
	log.Printf("Add Workflow: %s\n", wff.Name)
	return nil
}

func (wfm *workflowManager) RemoveWorkflow(name string) error {
	log.Fatal("workflowManager:RemoveWorkflow not Implement.")
	return nil
}
