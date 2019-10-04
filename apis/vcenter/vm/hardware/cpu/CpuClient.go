/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Cpu
 * Used by client-side stubs.
 */

package cpu

import (
)

// The ``Cpu`` interface provides methods for configuring the CPU settings of a virtual machine.
type CpuClient interface {


    // Returns the CPU-related settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return CPU-related settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string) (Info, error) 


    // Updates the CPU-related settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for updating the CPU-related settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws InvalidArgument if one of the provided settings is not permitted; for example, specifying a negative value for ``count``.
    // @throws NotAllowedInCurrentState if ``hotAddEnabled`` or ``hotRemoveEnabled`` is specified and the virtual machine is not powered off.
    // @throws NotAllowedInCurrentState if ``count`` is specified and is greater than ``count``, ``hotAddEnabled`` is false, and the virtual machine is not powered off.
    // @throws NotAllowedInCurrentState if ``count`` is specified and is less than ``count``, ``hotRemoveEnabled`` is false, and the virtual machine is not powered off.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Update(vmParam string, specParam UpdateSpec) error 

}
