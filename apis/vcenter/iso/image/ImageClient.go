/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Image
 * Used by client-side stubs.
 */

package image

import (
)

// Provides an interface to mount and unmount an ISO image on a virtual machine. 
//
//  This is an API that will let its client mount or unmount an ISO image on a virtual machine as a CD-ROM. 
type ImageClient interface {


    // Mounts an ISO image from a content library on a virtual machine.
    //
    // @param libraryItemParam The identifier of the library item having the ISO image to mount on the virtual machine.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param vmParam The identifier of the virtual machine where the specified ISO image will be mounted.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return The identifier of the newly created virtual CD-ROM backed by the specified ISO image.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @throws NotFound If either ``vm`` or the ``library_item`` is not found.
    // @throws InvalidArgument If no .iso file is present on the library item.
    // @throws NotAllowedInCurrentState When the operation is not allowed on the virtual machine in its current state.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Config.AddRemoveDevice``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item`` requires ``ContentLibrary.DownloadSession``.
    Mount(libraryItemParam string, vmParam string) (string, error) 


    // Unmounts a previously mounted CD-ROM using an ISO image as a backing.
    //
    // @param vmParam The identifier of the virtual machine from which to unmount the virtual CD-ROM.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param cdromParam The device identifier of the CD-ROM.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.Cdrom``.
    // @throws NotFound If the virtual machine identified by ``vm`` is not found or the ``cdrom`` does not identify a virtual CD-ROM in the virtual machine.
    // @throws NotAllowedInCurrentState When the operation is not allowed on the virtual machine in its current state.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Config.AddRemoveDevice``.
    // * The resource ``com.vmware.vcenter.vm.hardware.Cdrom`` referenced by the parameter ``cdrom`` requires ``System.Read``.
    Unmount(vmParam string, cdromParam string) error 

}
