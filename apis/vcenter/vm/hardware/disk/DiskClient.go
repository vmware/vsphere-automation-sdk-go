/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Disk
 * Used by client-side stubs.
 */

package disk

import (
)

// The ``Disk`` interface provides methods for configuring the virtual disks of a virtual machine. A virtual disk has a backing such as a VMDK file. The backing has an independent lifecycle from the virtual machine when it is detached from the virtual machine. The Disk#create method provides the ability to create a new virtual disk. When creating a virtual disk, a new VMDK file may be created or an existing VMDK file may used as a backing. Once a VMDK file is associated with a virtual machine, its lifecycle will be bound to the virtual machine. In other words, it will be deleted when the virtual machine is deleted. The Disk#delete method provides the ability to detach a VMDK file from the virtual machine. The Disk#delete method does not delete the VMDK file that backs the virtual disk. Once detached, the VMDK file will not be destroyed when the virtual machine to which it was associated is deleted.
type DiskClient interface {


    // Returns commonly used information about the virtual disks belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about the virtual disks.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(vmParam string) ([]Summary, error) 


    // Returns information about a virtual disk.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param diskParam Virtual disk identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Disk``.
    // @return Information about the specified virtual disk.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual disk is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string, diskParam string) (Info, error) 


    // Adds a virtual disk to the virtual machine. While adding the virtual disk, a new VMDK file may be created or an existing VMDK file may be used to back the virtual disk.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual disk.
    // @return Virtual disk identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Disk``.
    // @throws Error if system reported that the disk device was created but was unable to confirm the creation because the identifier of the new device could not be determined.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended or if the virtual machine is powered on and virtual disk type is IDE.
    // @throws UnableToAllocateResource if the specified storage address is unavailable; for example, if the SCSI adapter requested does not exist.
    // @throws ResourceInUse if the specified storage address is in use.
    // @throws InvalidArgument if the specified storage address is out of bounds or if the specified storage policy is invalid.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if the guest operating system of the virtual machine is not supported and spec includes null properties that default to guest-specific values.
    Create(vmParam string, specParam CreateSpec) (string, error) 


    // Updates the configuration of a virtual disk. An update method can be used to detach the existing VMDK file and attach another VMDK file to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param diskParam Virtual disk identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Disk``.
    // @param specParam Specification for updating the virtual disk.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual disk is not found.
    // @throws NotAllowedInCurrentState if one or more of the properties specified in the ``spec`` parameter cannot be modified due to the current power state of the virtual machine or the connection state of the virtual disk.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Update(vmParam string, diskParam string, specParam UpdateSpec) error 


    // Removes a virtual disk from the virtual machine. This method does not destroy the VMDK file that backs the virtual disk. It only detaches the VMDK file from the virtual machine. Once detached, the VMDK file will not be destroyed when the virtual machine to which it was associated is deleted.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param diskParam Virtual disk identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Disk``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual disk is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended or if the virtual machine is powered on and virtual disk type is IDE.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(vmParam string, diskParam string) error 

}
