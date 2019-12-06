/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package task

import (
	"fmt"
	"log"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils/session"
)

type task struct {
	task Function
}

func (t *task) Execute(sessionManager session.Manager) error {
	if sessionManager == nil {
		return fmt.Errorf("Task Execution Error: Invalid Session")
	}
	err := sessionManager.Login()
	if err != nil {
		return err
	}

	defer sessionManager.Logout()

	err = t.task(sessionManager)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
