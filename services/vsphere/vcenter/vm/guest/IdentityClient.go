/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Identity
 * Used by client-side stubs.
 */

package guest


// The ``Identity`` interface provides methods for retrieving guest operating system identification information. This interface was added in vSphere API 6.7.
type IdentityClient interface {

    // Return information about the guest. This method was added in vSphere API 6.7.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return guest identification information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws ServiceUnavailable if VMware Tools has not provided any data.
	Get(vmParam string) (IdentityInfo, error)
}
