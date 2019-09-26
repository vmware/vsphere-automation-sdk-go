/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Identity
 * Used by client-side stubs.
 */

package identity

import (
)

// The ``Identity`` interface provides methods for retrieving guest operating system identification information.
type IdentityClient interface {


    // Return information about the guest.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return guest identification information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws ServiceUnavailable if VMware Tools has not provided any data.
    Get(vmParam string) (Info, error) 

}
