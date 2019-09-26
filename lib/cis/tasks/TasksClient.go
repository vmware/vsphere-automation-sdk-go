/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tasks
 * Used by client-side stubs.
 */

package tasks

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/cis/task"
)

// The ``Tasks`` interface provides methods for managing the task related to a long running operation.
type TasksClient interface {


    // Returns information about a task.
    //
    // @param taskParam Task identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
    // @param specParam Specification on what to get for a task.
    // If null, the behavior is equivalent to a GetSpec with all properties null which means only the data described in task.Info will be returned and the result of the operation will be return.
    // @return Information about the specified task.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the task is not found.
    // @throws ResourceInaccessible if the task's state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(taskParam string, specParam *GetSpec) (task.Info, error) 


    // Returns information about at most 1000 visible (subject to permission checks) tasks matching the FilterSpec. All tasks must be in the same provider.
    //
    // @param filterSpecParam Specification of matching tasks.
    // This is currently required. In the future, if it is null, the behavior is equivalent to a FilterSpec with all properties null which means all tasks match the filter.
    // @param resultSpecParam Specification of what to return for a task.
    // If null, the behavior is equivalent to a GetSpec with all properties null which means only the data describe in task.Info will be returned and the result of the operation will be return.
    // @return Map of task identifier to information about the task.
    // The key in the return value map will be an identifier for the resource type: ``com.vmware.cis.task``.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws ResourceInaccessible if a task's state cannot be accessed or over 1000 tasks matching the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterSpecParam *FilterSpec, resultSpecParam *GetSpec) (map[string]task.Info, error) 


    // Cancel a running operation associated with the task. This is the best effort attempt. Operation may not be cancelled anymore once it reaches certain stage.
    //
    // @param taskParam Task identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the task is already canceled or completed.
    // @throws NotFound if the task is not found.
    // @throws ResourceInaccessible if the task's state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the task is not cancelable.
    Cancel(taskParam string) error 

}
