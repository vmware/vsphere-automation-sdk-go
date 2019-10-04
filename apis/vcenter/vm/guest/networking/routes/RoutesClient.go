/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Routes
 * Used by client-side stubs.
 */

package routes

import (
)

// The ``Routes`` interface provides methods for retrieving guest operating system network routing information.
type RoutesClient interface {


    // Returns information about network routing in the guest operating system.
    //
    // @param vmParam Virtual machine ID
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the network routes configured in the guest operating system.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    List(vmParam string) ([]Info, error) 

}
