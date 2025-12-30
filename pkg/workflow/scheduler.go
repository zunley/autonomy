package workflow

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
)

var defaultCronSpec = cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow

type Task func()

type Scheduler interface {
	/*
	 * Args:
	 * 	spec: cron string
	 * 	task: function
	 */
	Schedule(name, spec string, task Task) error
	Start() error
	Stop() error
}

type CronScheduler struct {
	tasks map[string]cron.EntryID
	cron  *cron.Cron
}

func NewCronScheduler() Scheduler {
	return &CronScheduler{
		tasks: make(map[string]cron.EntryID),
		cron:  cron.New(cron.WithSeconds()),
	}
}

func (cs *CronScheduler) Schedule(name, spec string, task Task) error {
	if _, err := cron.ParseStandard(spec); err != nil {
		s := fmt.Sprintf("Invalid schedule %s", spec)
		return errors.New(s)
	}

	entryID, err := cs.cron.AddFunc(spec, task)
	if err != nil {
		return err
	}
	cs.tasks[name] = entryID
	return nil
}

func (cs *CronScheduler) Start() error {
	cs.cron.Start()
	return nil
}

func (cs *CronScheduler) Stop() error {
	cs.cron.Stop()
	return nil
}
