/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Update
 * Used by client-side stubs.
 */

package update

import (
)

// The ``Update`` interface provides methods to get the status of the appliance update.
type UpdateClient interface {


    // Gets the current status of the appliance update.
    // @return Info structure containing the status information about the appliance.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    Get() (Info, error) 


    // Request the cancellation the update operation that is currently in progress.
    // @throws Error Generic error
    // @throws NotAllowedInCurrentState Current task is not cancellable
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    Cancel() error 

}
