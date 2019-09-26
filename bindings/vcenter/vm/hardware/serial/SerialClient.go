/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Serial
 * Used by client-side stubs.
 */

package serial

import (
)

// The ``Serial`` interface provides methods for configuring the virtual serial ports of a virtual machine.
type SerialClient interface {


    // Returns commonly used information about the virtual serial ports belonging to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return List of commonly used information about virtual serial ports.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(vmParam string) ([]Summary, error) 


    // Returns information about a virtual serial port.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param portParam Virtual serial port identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SerialPort``.
    // @return Information about the specified virtual serial port.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual serial port is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string, portParam string) (Info, error) 


    // Adds a virtual serial port to the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the new virtual serial port.
    // @return Virtual serial port identifier.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SerialPort``.
    // @throws Error if the system reported that the serial port device was created but was unable to confirm the creation because the identifier of the new device could not be determined.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered off.
    // @throws UnableToAllocateResource if the virtual machine already has the maximum number of supported serial ports.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Create(vmParam string, specParam CreateSpec) (string, error) 


    // Updates the configuration of a virtual serial port.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param portParam Virtual serial port identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SerialPort``.
    // @param specParam Specification for updating the virtual serial port.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual serial port is not found.
    // @throws NotAllowedInCurrentState if one or more of the properties specified in the ``spec`` parameter cannot be modified due to the current power state of the virtual machine or the connection state of the virtual serial port.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Update(vmParam string, portParam string, specParam UpdateSpec) error 


    // Removes a virtual serial port from the virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param portParam Virtual serial port identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SerialPort``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual serial port is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered off.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(vmParam string, portParam string) error 


    // Connects a virtual serial port of a powered-on virtual machine to its backing. Connecting the virtual device makes the backing accessible from the perspective of the guest operating system. 
    //
    //  For a powered-off virtual machine, the Serial#update method may be used to configure the virtual serial port to start in the connected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param portParam Virtual serial port identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SerialPort``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual serial port is not found.
    // @throws AlreadyInDesiredState if the virtual serial port is already connected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Connect(vmParam string, portParam string) error 


    // Disconnects a virtual serial port of a powered-on virtual machine from its backing. The virtual device is still present and its backing configuration is unchanged, but from the perspective of the guest operating system, the serial port is not connected to its backing. 
    //
    //  For a powered-off virtual machine, the Serial#update method may be used to configure the virtual serial port to start in the disconnected state when the virtual machine is powered on.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param portParam Virtual serial port identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vm.hardware.SerialPort``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine or virtual serial port is not found.
    // @throws AlreadyInDesiredState if the virtual serial port is already disconnected.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Disconnect(vmParam string, portParam string) error 

}
