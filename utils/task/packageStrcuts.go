package task

import (
	client "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

type TaskFunc func(connector client.Connector) error

type taskImpl struct {
	task TaskFunc
}

func CreateNewTask(function TaskFunc) *Task {
	var task Task = &taskImpl{task: function}
	return &task
}
