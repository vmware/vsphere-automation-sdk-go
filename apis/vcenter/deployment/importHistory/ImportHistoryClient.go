/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ImportHistory
 * Used by client-side stubs.
 */

package importHistory

import (
)

// The ``ImportHistory`` interface provides methods for managing the import of vCenter historical data, e.g. Tasks, Events and Statistics, when is is imported separately from the upgrade or migration process.
type ImportHistoryClient interface {


    // Get the current status of the vCenter historical data import.
    // @return Info structure containing the status information about the historical data import status.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized to perform the method.
    // @throws Error Generic error
    Get() (Info, error) 


    // Creates and starts task for importing vCenter historical data.
    //
    // @param specParam An optional ``CreateSpec`` info that can be passed for creating a new historical data import task and starts it.
    // If null, default value will be: 
    //
    // * name : vcenter.deployment.history.import
    // * description : vCenter Server history import
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized to perform the method.
    // @throws NotAllowedInCurrentState if vCenter historical data import task cannot be started at this time. This can happen in the following cases: 
    //
    // * If historical data import has already been canceled because a canceled task cannot be re-started
    // * If historical data import has already been completed because a completed task cannot be re-started
    // * If historical data import has already been paused because a paused task can only be resumed or canceled
    // @throws AlreadyInDesiredState if vCenter historical data import task has already being started.
    // @throws Error Generic error
    Start(specParam *CreateSpec) error 


    // Pauses the task for importing vCenter historical data.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized to perform the method.
    // @throws NotAllowedInCurrentState if vCenter historical data import task cannot be paused at this time. Pause can be accepted only in task.Status#Status_RUNNING state.
    // @throws AlreadyInDesiredState if vCenter historical data import task is already paused
    // @throws Error Generic error
    Pause() error 


    // Resumes the task for importing vCenter historical data.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized to perform the method.
    // @throws NotAllowedInCurrentState if vCenter historical data import task cannot be resumed at this state. Resume can be accepted only in task.Status#Status_BLOCKED state
    // @throws AlreadyInDesiredState if vCenter historical data import task is already resumed.
    // @throws Error Generic error.
    Resume() error 


    // Cancels the task for importing vCenter historical data.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized to perform the method.
    // @throws NotAllowedInCurrentState if vCenter historical data import task cannot be canceled at this state. This can happen in the following cases: 
    //
    // * If historical data import has not been started yet because a not running task cannot be canceled
    // * If historical data import has already been completed because a completed task cannot be canceled
    // @throws AlreadyInDesiredState if vCenter historical data import task is already canceled.
    // @throws Error Generic error.
    Cancel() error 

}
