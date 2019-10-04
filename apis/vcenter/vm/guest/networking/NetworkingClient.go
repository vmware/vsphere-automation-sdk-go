/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Networking
 * Used by client-side stubs.
 */

package networking

import (
)

// The ``Networking`` interface provides methods for retrieving guest operating system network information.
type NetworkingClient interface {


    // Returns information about the network configuration in the guest operating system.
    //
    // @param vmParam Virtual machine ID
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the networking configuration in the guest operating system.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    Get(vmParam string) (Info, error) 

}
