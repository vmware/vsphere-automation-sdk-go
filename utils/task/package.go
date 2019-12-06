/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Package task manages the task(s) to be executed on the Host server.
package task

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
)

// Function is function prototype, implementation of which can be executed as Task on the Host server.
type Function func(sessionManager session.Manager) error

// CreateNewTask creates a Task with function, that can be executed on the Host server.
func CreateNewTask(function Function) *Task {
	var task Task = &task{task: function}
	return &task
}
