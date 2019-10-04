/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Sata
 * Used by client-side stubs.
 */

package sata

import (
)

// The ``Sata`` interface provides methods for configuring the virtual SATA adapters of a virtual machine.
type SataClient interface {


    // Returns commonly used information about the virtual SATA adapters belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual SATA adapters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(vmParam string) ([]Summary, error) 


    // Returns information about a virtual SATA adapter.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual SATA adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SataAdapter``.
    // @return Information about the specified virtual SATA adapter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual SATA adapter is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string, adapterParam string) (Info, error) 


    // Adds a virtual SATA adapter to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual SATA adapter.
    // @return Virtual SATA adapter identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SataAdapter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Error if the system reported that the SATA adapter was created but was unable to confirm the creation because the identifier of the new adapter could not be determined.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended
    // @throws NotFound if the virtual machine is not found.
    // @throws UnableToAllocateResource if there are no more available SATA buses on the virtual machine.
    // @throws ResourceInUse if the specified SATA bus or PCI address is in use.
    // @throws InvalidArgument if the specified SATA bus or PCI address is out of bounds.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the guest operating system of the virtual machine is not supported and spec includes null properties that default to guest-specific values.
    Create(vmParam string, specParam CreateSpec) (string, error) 


    // Removes a virtual SATA adapter from the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param adapterParam Virtual SATA adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SataAdapter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended
    // @throws NotFound if the virtual machine or virtual SATA adapter is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(vmParam string, adapterParam string) error 

}
