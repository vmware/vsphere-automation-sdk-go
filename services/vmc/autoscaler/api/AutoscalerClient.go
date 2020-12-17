/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Autoscaler
 * Used by client-side stubs.
 */

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
)

type AutoscalerClient interface {

    // Get cross-cluster load-balancer recommendations for the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @param clustersParam List of input cluster uuids. The minimum number of input clusters for cross-cluster load balancer is 2. The maximum number of input clusters for cross-cluster load balancer is 4. Example - \\\\`[\"e823eb1c-d79d-4763-ad6b-b447c14b6cd2\", \"71066e28-3a15-4670-8555-72068c9d5320\"]\\\\` (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Analysis(orgParam string, sddcParam string, clustersParam []string) (model.Task, error)

    // Retrieve details of an autoscaler task.
    //
    // @param orgParam org identifier (required)
    // @param taskParam task identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the task with given identifier
	Get(orgParam string, taskParam string) (model.Task, error)

    // List all tasks with optional filtering.
    //
    // @param orgParam org identifier (required)
    // @param filterParam Filter expression Binary Operators: 'eq', 'ne', 'lt', 'gt', 'le', 'ge', 'mul', 'div', 'mod', 'sub', 'add' Unary Operators: 'not', '-' (minus) String Operators: 'startswith', 'endswith', 'length', 'contains', 'tolower', 'toupper', Nested attributes are composed using '.' Dates must be formatted as yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss[.SSS]Z Strings should enclosed in single quotes, escape single quote with two single quotes The special literal 'created' will be mapped to the time the resource was first created. Examples: - $filter=(updated gt 2016-08-09T13:00:00Z) and (org_id eq 278710ff4e-6b6d-4d4e-aefb-ca637f38609e) - $filter=(created eq 2016-08-09) - $filter=(created gt 2016-08-09) and (sddc.status eq 'READY') (optional)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Forbidden
	List(orgParam string, filterParam *string) ([]model.Task, error)

    // Execute cross-cluster load balancing operations on the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Run(orgParam string, sddcParam string) (model.Task, error)

    // Stop cross-cluster load balancer initiated xvMotion operations on the sddc.
    //
    // @param orgParam org identifier (required)
    // @param sddcParam Sddc identifier (required)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
	Stop(orgParam string, sddcParam string) (model.Task, error)

    // Request that a running task be canceled. This is advisory only, some tasks may not be cancelable, and some tasks might take an arbitrary amount of time to respond to a cancelation request. The task must be monitored to determine subsequent status.
    //
    // @param orgParam org identifier (required)
    // @param taskParam task identifier (required)
    // @param actionParam If = 'cancel', task will be cancelled (optional)
    // @return com.vmware.model.Task
    // @throws Unauthenticated  Unauthorized
    // @throws InvalidRequest  Invalid action or bad argument
    // @throws Unauthorized  Forbidden
    // @throws NotFound  Cannot find the task with given identifier
	Update(orgParam string, taskParam string, actionParam *string) (model.Task, error)
}
