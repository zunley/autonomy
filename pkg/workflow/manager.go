package workflow

import (
	"github.com/robfig/cron/v3"
	"github.com/zunley/autonomy/pkg/types"
	"errors"
	"fmt"
)

type Manager interface {
	Run() error
	Stop()
	AddWorkflow(*types.Workflow) error
}

func NewManager() Manager {

	return &workflowManager {
		schedule: cron.New(),
		workflows: make(map[cron.EntryID]Workflow),
	}
}

type workflowManager struct {
	schedule *cron.Cron
	workflows map[cron.EntryID]Workflow
}

func (wfm *workflowManager) Run() error {
	if wfm.schedule != nil {
		wfm.schedule.Start()
	}
	return nil
}

func (wfm *workflowManager) Stop() {
	if wfm.schedule != nil {
		wfm.schedule.Stop()
	}
}

func (wfm *workflowManager) AddWorkflow(wff *types.Workflow) error {
	
	if _, err := cron.ParseStandard(wff.Schedule); err != nil {
		s := fmt.Sprintf("Invalid schedule for workflow %s: %s", wff.Name, wff.Schedule)
		return errors.New(s)
	}

	wf := NewWorkflow(wff)
	entryID, err := wfm.schedule.AddFunc(wff.Schedule, func() {
		rst, _ := wf.Run()
		// TODO upload
		fmt.Printf("%+v", rst)
	})
	if err != nil {
		return err
	}

	wfm.workflows[entryID] = wf

	return nil
}