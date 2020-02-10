/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Boot
 * Used by client-side stubs.
 */

package hardware


// The ``Boot`` interface provides methods for configuring the settings used when booting a virtual machine.
type BootClient interface {

    // Returns the boot-related settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Boot-related settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string) (BootInfo, error)

    // Updates the boot-related settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for updating the boot-related settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws InvalidArgument if one of the provided settings is not permitted; for example, specifying a negative value for ``delay``.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Update(vmParam string, specParam BootUpdateSpec) error
}
