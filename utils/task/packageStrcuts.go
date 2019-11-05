package task

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
)

type TaskFunc func(sessionManager session.SessionManager) error

type taskImpl struct {
	task TaskFunc
}

func CreateNewTask(function TaskFunc) *Task {
	var task Task = &taskImpl{task: function}
	return &task
}
