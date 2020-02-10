/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Cdrom
 * Used by client-side stubs.
 */

package hardware


// The ``Cdrom`` interface provides methods for configuring the virtual CD-ROM devices of a virtual machine.
type CdromClient interface {

    // Returns commonly used information about the virtual CD-ROM devices belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual CD-ROM devices.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(vmParam string) ([]CdromSummary, error)

    // Returns information about a virtual CD-ROM device.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param cdromParam Virtual CD-ROM device identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @return Information about the specified virtual CD-ROM device.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual CD-ROM device is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(vmParam string, cdromParam string) (CdromInfo, error)

    // Adds a virtual CD-ROM device to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual CD-ROM device.
    // @return Virtual CD-ROM device identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @throws Error if the system reported that the CD-ROM device was created but was unable to confirm the creation because the identifier of the new device could not be determined.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended or if the virtual machine is powered on and virtual CD-ROM type is IDE.
    // @throws UnableToAllocateResource if the specified storage address is unavailable; for example, if the SCSI adapter requested does not exist.
    // @throws ResourceInUse if the specified storage address is in use.
    // @throws InvalidArgument if the specified storage address is out of bounds.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the guest operating system of the virtual machine is not supported and spec includes null properties that default to guest-specific values.
	Create(vmParam string, specParam CdromCreateSpec) (string, error)

    // Updates the configuration of a virtual CD-ROM device.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param cdromParam Virtual CD-ROM device identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @param specParam Specification for updating the virtual CD-ROM device.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual CD-ROM device is not found.
    // @throws NotAllowedInCurrentState if one or more of the properties specified in the ``spec`` parameter cannot be modified due to the current power state of the virtual machine or the connection state of the virtual CD-ROM device.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Update(vmParam string, cdromParam string, specParam CdromUpdateSpec) error

    // Removes a virtual CD-ROM device from the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param cdromParam Virtual CD-ROM device identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual CD-ROM device is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended or if the virtual machine is powered on and virtual CD-ROM type is IDE.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Delete(vmParam string, cdromParam string) error

    // Connects a virtual CD-ROM device of a powered-on virtual machine to its backing. Connecting the virtual device makes the backing accessible from the perspective of the guest operating system. 
    //
    //  For a powered-off virtual machine, the Cdrom#update method may be used to configure the virtual CD-ROM device to start in the connected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param cdromParam Virtual CD-ROM device identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual CD-ROM device is not found.
    // @throws AlreadyInDesiredState if the virtual CD-ROM device is already connected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Connect(vmParam string, cdromParam string) error

    // Disconnects a virtual CD-ROM device of a powered-on virtual machine from its backing. The virtual device is still present and its backing configuration is unchanged, but from the perspective of the guest operating system, the CD-ROM device is not connected to its backing resource. 
    //
    //  For a powered-off virtual machine, the Cdrom#update method may be used to configure the virtual CD-ROM device to start in the disconnected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param cdromParam Virtual CD-ROM device identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual CD-ROM device is not found.
    // @throws AlreadyInDesiredState if the virtual CD-ROM device is already disconnected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Disconnect(vmParam string, cdromParam string) error
}
