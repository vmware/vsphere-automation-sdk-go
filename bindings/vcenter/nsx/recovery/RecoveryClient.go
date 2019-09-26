/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Recovery
 * Used by client-side stubs.
 */

package recovery

import (
)

// The ``Recovery`` interface represents all the operations of NSXi Infrastructure Recovery for vSphere clusters. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RecoveryClient interface {


    // Returns whether recovery is needed. If needed, then returns the type of loss that needs recovery. If recovery is not needed, then clears the recovery state to allow the WCP service to enter steady state. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Returns an Info
    // @throws Unauthorized if the user is not a member of the Administrators group.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws InternalServerError if the type of recovery is invalid.
    // @throws Error for any other unspecified error.
    Get() (Info, error) 


    // Recovers the system from the loss of WCP service and/or NSX. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Contains information related to recovery of NSX for WCP clusters.
    // @return The task identifier for the operation. The task is not cancellable.
    // The return value will be an identifier for the resource type: ``Task``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators group.
    // @throws NotAllowedInCurrentState if the operation is already in progress.
    // @throws AlreadyInDesiredState if the operation is not needed.
    // @throws Unsupported if NSX-I was not installed.
    // @throws InternalServerError if the type of recovery is invalid.
    // @throws Error for any other unspecified error.
    Execute(specParam ExecuteSpec) (string, error) 


    // Gets the recovery status. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Returns an ExecutionStatus.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators group.
    // @throws Error for any other unspecified error.
    ExecuteStatus() (ExecutionStatus, error) 

}
