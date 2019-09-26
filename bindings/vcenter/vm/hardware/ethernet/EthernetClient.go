/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ethernet
 * Used by client-side stubs.
 */

package ethernet

import (
)

// The ``Ethernet`` interface provides methods for configuring the virtual Ethernet adapters of a virtual machine.
type EthernetClient interface {


    // Returns commonly used information about the virtual Ethernet adapters belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual Ethernet adapters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(vmParam string) ([]Summary, error) 


    // Returns information about a virtual Ethernet adapter.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param nicParam Virtual Ethernet adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Ethernet``.
    // @return Information about the specified virtual Ethernet adapter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual Ethernet adapter is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string, nicParam string) (Info, error) 


    // Adds a virtual Ethernet adapter to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual Ethernet adapter.
    // @return Virtual Ethernet adapter identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Ethernet``.
    // @throws Error if the system reported that the Ethernet adapter was created but was unable to confirm the creation because the identifier of the new adapter could not be determined.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or network backing is not found.
    // @throws UnableToAllocateResource if the virtual machine already has the maximum number of supported Ethernet adapters.
    // @throws InvalidArgument if the specified PCI address is out of bounds, HOST_DEVICE is specified as the type, or a backing cannot be found in the case that backing is left null.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the guest operating system of the virtual machine is not supported and spec includes null properties that default to guest-specific values.
    Create(vmParam string, specParam CreateSpec) (string, error) 


    // Updates the configuration of a virtual Ethernet adapter.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param nicParam Virtual Ethernet adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Ethernet``.
    // @param specParam Specification for updating the virtual Ethernet adapter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if HOST_DEVICE is specified as the type.
    // @throws NotFound if the virtual machine, virtual Ethernet adapter, or backing network is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Update(vmParam string, nicParam string, specParam UpdateSpec) error 


    // Removes a virtual Ethernet adapter from the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param nicParam Virtual Ethernet adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Ethernet``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual Ethernet adapter is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(vmParam string, nicParam string) error 


    // Connects a virtual Ethernet adapter of a powered-on virtual machine to its backing. Connecting the virtual device makes the backing accessible from the perspective of the guest operating system. 
    //
    //  For a powered-off virtual machine, the Ethernet#update method may be used to configure the virtual Ethernet adapter to start in the connected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param nicParam Virtual Ethernet adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Ethernet``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual Ethernet adapter is not found.
    // @throws AlreadyInDesiredState if the virtual Ethernet adapter is already connected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Connect(vmParam string, nicParam string) error 


    // Disconnects a virtual Ethernet adapter of a powered-on virtual machine from its backing. The virtual device is still present and its backing configuration is unchanged, but from the perspective of the guest operating system, the Ethernet adapter is not connected to its backing resource. 
    //
    //  For a powered-off virtual machine, the Ethernet#update method may be used to configure the virtual Ethernet adapter to start in the disconnected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param nicParam Virtual Ethernet adapter identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Ethernet``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual Ethernet adapter is not found.
    // @throws AlreadyInDesiredState if the virtual Ethernet adapter is already disconnected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Disconnect(vmParam string, nicParam string) error 

}
