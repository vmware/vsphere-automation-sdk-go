/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Interfaces
 * Used by client-side stubs.
 */

package networking


// The ``Interfaces`` interface provides methods for retrieving guest operating system network interface information. This interface was added in vSphere API 7.0.0.
type InterfacesClient interface {

    // Returns information about the networking interfaces in the guest operating system. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Virtual machine ID
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the interfaces configured in the guest operating system. Interfaces are ordered in a guest operating system specific determined order.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
	List(vmParam string) ([]InterfacesInfo, error)
}
