/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Hardware
 * Used by client-side stubs.
 */

package hardware

import (
)

// The ``Hardware`` interface provides methods for configuring the virtual hardware of a virtual machine.
type HardwareClient interface {


    // Returns the virtual hardware settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Virtual hardware settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string) (Info, error) 


    // Updates the virtual hardware settings of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for updating the virtual hardware settings of the virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws AlreadyInDesiredState if the virtual machine is already configured for the desired hardware version.
    // @throws InvalidArgument if the requested virtual hardware version is not newer than the current version.
    // @throws Unsupported if the requested virtual hardware version is not supported by the server.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Update(vmParam string, specParam UpdateSpec) error 


    // Upgrades the virtual machine to a newer virtual hardware version.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param versionParam New virtual machine version.
    // If null, defaults to the most recent virtual hardware version supported by the server.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered off.
    // @throws AlreadyInDesiredState if the virtual machine is already configured for the desired hardware version.
    // @throws InvalidArgument if ``version`` is older than the current virtual hardware version.
    // @throws Unsupported if ``version`` is not supported by the server.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Upgrade(vmParam string, versionParam *Version) error 

}
