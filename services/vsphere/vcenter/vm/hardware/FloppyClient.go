/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Floppy
 * Used by client-side stubs.
 */

package hardware


// The ``Floppy`` interface provides methods for configuring the virtual floppy drives of a virtual machine.
type FloppyClient interface {

    // Returns commonly used information about the virtual floppy drives belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual floppy drives.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(vmParam string) ([]FloppySummary, error)

    // Returns information about a virtual floppy drive.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param floppyParam Virtual floppy drive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Floppy``.
    // @return Information about the specified virtual floppy drive.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual floppy drive is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string, floppyParam string) (FloppyInfo, error)

    // Adds a virtual floppy drive to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual floppy drive.
    // @return Virtual floppy drive identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Floppy``.
    // @throws Error if the system reported that the floppy device was created but was unable to confirm the creation because the identifier of the new device could not be determined.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered off.
    // @throws UnableToAllocateResource if the virtual machine already has the maximum number of supported floppy drives.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Create(vmParam string, specParam FloppyCreateSpec) (string, error)

    // Updates the configuration of a virtual floppy drive.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param floppyParam Virtual floppy drive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Floppy``.
    // @param specParam Specification for updating the virtual floppy drive.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual floppy drive is not found.
    // @throws NotAllowedInCurrentState if one or more of the properties specified in the ``spec`` parameter cannot be modified due to the current power state of the virtual machine or the connection state of the virtual floppy drive.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Update(vmParam string, floppyParam string, specParam FloppyUpdateSpec) error

    // Removes a virtual floppy drive from the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param floppyParam Virtual floppy drive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Floppy``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual floppy drive is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered off.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Delete(vmParam string, floppyParam string) error

    // Connects a virtual floppy drive of a powered-on virtual machine to its backing. Connecting the virtual device makes the backing accessible from the perspective of the guest operating system. 
    //
    //  For a powered-off virtual machine, the Floppy#update method may be used to configure the virtual floppy drive to start in the connected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param floppyParam Virtual floppy drive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Floppy``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual floppy drive is not found.
    // @throws AlreadyInDesiredState if the virtual floppy drive is already connected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Connect(vmParam string, floppyParam string) error

    // Disconnects a virtual floppy drive of a powered-on virtual machine from its backing. The virtual device is still present and its backing configuration is unchanged, but from the perspective of the guest operating system, the floppy drive is not connected to its backing resource. 
    //
    //  For a powered-off virtual machine, the Floppy#update method may be used to configure the virtual floppy floppy to start in the disconnected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param floppyParam Virtual floppy drive identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Floppy``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual floppy drive is not found.
    // @throws AlreadyInDesiredState if the virtual floppy drive is already disconnected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Disconnect(vmParam string, floppyParam string) error
}
