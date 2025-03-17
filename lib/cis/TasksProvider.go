// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Tasks
// Used by service-side to provide implementations.

package cis

import (
	cisTask_ "github.com/vmware/vsphere-automation-sdk-go/lib/cis/task"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
)

// The “Tasks“ interface provides methods for managing the task related to a long running operation.
type TasksProvider interface {

	// Returns information about a task.
	//
	// @param taskParam Task identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
	// @param specParam Specification on what to get for a task.
	// If null, the behavior is equivalent to a TasksGetSpec with all properties null which means only the data described in cisTask_.Info will be returned and the result of the operation will be return.
	// @return Information about the specified task.
	//
	// @throws Error if the system reports an error while responding to the request.
	// @throws NotFound if the task is not found.
	// @throws ResourceInaccessible if the task's state cannot be accessed.
	// @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
	// @throws Unauthenticated if the user can not be authenticated.
	// @throws Unauthorized if the user doesn't have the required privileges.
	Get(taskParam string, specParam *TasksGetSpec, ctx *vapiCore_.ExecutionContext) (cisTask_.Info, error)
	// Returns information about at most 1000 visible (subject to permission checks) tasks matching the TasksFilterSpec. All tasks must be in the same provider.
	//
	// @param filterSpecParam Specification of matching tasks.
	// This is currently required. In the future, if it is null, the behavior is equivalent to a TasksFilterSpec with all properties null which means all tasks match the filter.
	// @param resultSpecParam Specification of what to return for a task.
	// If null, the behavior is equivalent to a TasksGetSpec with all properties null which means only the data describe in cisTask_.Info will be returned and the result of the operation will be return.
	// @return Map of task identifier to information about the task.
	// The key in the return value map will be an identifier for the resource type: ``com.vmware.cis.task``.
	//
	// @throws InvalidArgument if any of the specified parameters are invalid.
	// @throws ResourceInaccessible if a task's state cannot be accessed or over 1000 tasks matching the TasksFilterSpec.
	// @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
	// @throws Unauthenticated if the user can not be authenticated.
	// @throws Unauthorized if the user doesn't have the required privileges.
	List(filterSpecParam *TasksFilterSpec, resultSpecParam *TasksGetSpec, ctx *vapiCore_.ExecutionContext) (map[string]cisTask_.Info, error)
	// Cancel a running operation associated with the task. This is the best effort attempt. Operation may not be cancelled anymore once it reaches certain stage.
	//
	// @param taskParam Task identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
	//
	// @throws Error if the system reports an error while responding to the request.
	// @throws NotAllowedInCurrentState if the task is already canceled or completed.
	// @throws NotFound if the task is not found.
	// @throws ResourceInaccessible if the task's state cannot be accessed.
	// @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
	// @throws Unauthenticated if the user can not be authenticated.
	// @throws Unauthorized if the user doesn't have the required privileges.
	// @throws Unsupported if the task is not cancelable.
	Cancel(taskParam string, ctx *vapiCore_.ExecutionContext) error
}
