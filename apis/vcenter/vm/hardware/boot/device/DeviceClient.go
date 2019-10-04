/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Device
 * Used by client-side stubs.
 */

package device

import (
)

// The ``Device`` interface provides methods for configuring the device order used when booting a virtual machine. 
//
//  The boot order may be specified using a mixture of device classes and device instances, chosen from among the following: 
//
// * Type#Type_CDROM: Boot from a virtual CD-ROM drive; the device instance(s) will be chosen by the BIOS subsystem.
// * Type#Type_FLOPPY: Boot from a virtual floppy drive; the device instance(s) will be chosen by the BIOS subsystem.
// * Type#Type_DISK: Boot from a virtual disk device; the device instance is specified explicitly in Entry#disks list, and multiple instances may be specified in the list.
// * Type#Type_ETHERNET: Boot from a virtual Ethernet adapter; the device instance is specified explicitly as Entry#nic, and multiple adapters may be specified in the boot order list.
type DeviceClient interface {


    // Returns an ordered list of boot devices for the virtual machine. If the array is empty, the virtual machine uses a default boot sequence.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Ordered list of configured boot devices.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string) ([]Entry, error) 


    // Sets the virtual devices that will be used to boot the virtual machine. The virtual machine will check the devices in order, attempting to boot from each, until the virtual machine boots successfully. If the array is empty, the virtual machine will use a default boot sequence. There should be no more than one instance of Entry for a given device type except Type#Type_ETHERNET in the array.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param devicesParam Ordered list of boot devices.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found, or if any of the specified virtual devices is not found.
    // @throws InvalidArgument if a any of the CDROM, DISK, ETHERNET, FLOPPY values appears in more than one ``Entry`` with the exception of Type#Type_ETHERNET, which may appear multiple times if the virtual machine has been configured with multiple Ethernet adapters.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Set(vmParam string, devicesParam []Entry) error 

}
