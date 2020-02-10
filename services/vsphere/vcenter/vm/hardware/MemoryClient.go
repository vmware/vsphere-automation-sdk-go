/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Memory
 * Used by client-side stubs.
 */

package hardware


// The ``Memory`` interface provides methods for configuring the memory settings of a virtual machine.
type MemoryClient interface {

    // Returns the memory-related settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Memory-related settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string) (MemoryInfo, error)

    // Updates the memory-related settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for updating the memory-related settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws InvalidArgument if one of the provided settings is not permitted; for example, specifying a negative value for ``sizeMiB``.
    // @throws NotAllowedInCurrentState if ``hotAddEnabled`` is specified and the virtual machine is not powered off.
    // @throws NotAllowedInCurrentState if ``sizeMiB`` is specified, ``hotAddEnabled`` is false, and the virtual machine is not powered off.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Update(vmParam string, specParam MemoryUpdateSpec) error
}
