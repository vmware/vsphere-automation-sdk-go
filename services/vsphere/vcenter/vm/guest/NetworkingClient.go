/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Networking
 * Used by client-side stubs.
 */

package guest


// The ``Networking`` interface provides methods for retrieving guest operating system network information. This interface was added in vSphere API 7.0.0.
type NetworkingClient interface {

    // Returns information about the network configuration in the guest operating system. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Virtual machine ID
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the networking configuration in the guest operating system.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
	Get(vmParam string) (NetworkingInfo, error)
}
