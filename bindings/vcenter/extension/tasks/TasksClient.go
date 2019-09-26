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
)

// The ``Tasks`` interface provides methods for creating tasks for an operation in vCenter extension and updating the description, state and progress of the task. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TasksClient interface {


    // Returns a taskId for the task using the CreateSpec. The taskId can be used for querying or updating the status of the task by the client. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification to be used by the service for creating a task.
    // @return Unique id of the task created, which can be used by client for querying and updating the task's status
    // The return value will be an identifier for the resource type: ``com.vmware.cis.task``.
    // @throws Error if the system reported an error while responding to the request.
    // @throws NotAllowedInCurrentState if attempted to update percentDone of a task that is not running or to change state after task is completed or errored or to set startTime of a non-running task or to set endTime of a non-completed task.
    // @throws InvalidArgument when the percentDone.completed of the task is less than 0 or greater than percentDone.total or if the error field is not populated for a failed task or if client passes a invalid status, startTime or endTime.
    Create(specParam CreateSpec) (string, error) 


    // Updates the information of the operation associated with a task on the ManagedObjects. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param taskIdParam Task-Id of the task whose information is to be updated.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
    // @param specParam Specification to be used by the service for updating a task.
    // @throws Error if the system reported an error while responding to the request.
    // @throws NotAllowedInCurrentState if attempted to update percentDone of a task that is not running or to change state after task is completed or errored or to set startTime of a non-running task or to set endTime of a non-completed task.
    // @throws InvalidArgument when the percentDone.completed of the task is less than 0 or greater than percentDone.total or if the error field is not populated for a failed task or if client passes a invalid status, startTime or endTime.
    Update(taskIdParam string, specParam UpdateSpec) error 

}
