package task

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
)

type TaskFunc func(sessionManager session.Manager) error

func CreateNewTask(function TaskFunc) *Task {
	var task Task = &task{task: function}
	return &task
}
