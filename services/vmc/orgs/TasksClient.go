/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Tasks
 * Used by client-side stubs.
 */

package orgs

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

type TasksClient interface {

    // Retrieve details of a task.
    //
    // @param orgParam Organization identifier (required)
    // @param taskParam Task identifier (required)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the task with given identifier
	Get(orgParam string, taskParam string) (model.Task, error)

    // List all tasks with optional filtering.
    //
    // @param orgParam Organization identifier (required)
    // @param filterParam Filter expression Binary Operators: 'eq', 'ne', 'lt', 'gt', 'le', 'ge', 'mul', 'div', 'mod', 'sub', 'add' Unary Operators: 'not', '-' (minus) String Operators: 'startswith', 'endswith', 'length', 'contains', 'tolower', 'toupper', Nested attributes are composed using '.' Dates must be formatted as yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss[.SSS]Z Strings should enclosed in single quotes, escape single quote with two single quotes The special literal 'created' will be mapped to the time the resource was first created. Examples: - $filter=(updated gt 2016-08-09T13:00:00Z) and (org_id eq 278710ff4e-6b6d-4d4e-aefb-ca637f38609e) - $filter=(created eq 2016-08-09) - $filter=(created gt 2016-08-09) and (sddc.status eq 'READY') (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	List(orgParam string, filterParam *string) ([]model.Task, error)

    // Request that a running task be canceled. This is advisory only, some tasks may not be cancelable, and some tasks might take an arbitrary amount of time to respond to a cancelation request. The task must be monitored to determine subsequent status.
    //
    // @param orgParam Organization identifier (required)
    // @param taskParam Task identifier (required)
    // @param actionParam If = 'cancel', task will be canceled (optional)
    // @return com.vmware.vmc.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the task with given identifier
	Update(orgParam string, taskParam string, actionParam *string) (model.Task, error)
}
