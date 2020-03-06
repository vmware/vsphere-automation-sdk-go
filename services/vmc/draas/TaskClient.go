/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Task
 * Used by client-side stubs.
 */

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

type TaskClient interface {

    // Retrieve details of a task.
    //
    // @param orgParam Organization identifier (required)
    // @param taskParam task identifier (required)
    // @return com.vmware.vmc.draas.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the task with given identifier
	Get(orgParam string, taskParam string) (model.Task, error)
}
