package workflow

import "errors"

type Manager interface {
	Run() error
	Stop()
}

func NewManager() Manager {
	return nil
}
