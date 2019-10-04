/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Pending
 * Used by client-side stubs.
 */

package pending

import (
)

// The ``Pending`` interface provides method for listing pending minor or major updates of vCenter Server.
type PendingClient interface {


    // Lists all available minor and major updates.
    // @return Information about the pending patch/updates for the given vCenter server
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    List() (ListResult, error) 


    // Gets detailed update information.
    //
    // @param versionParam A version identified the update
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.update.pending``.
    // @return A detailed information about the particular vCenter patch/update
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws NotFound if there is no pending update assosiated with the ``version`` in the system.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    Get(versionParam string) (Info, error) 

}
