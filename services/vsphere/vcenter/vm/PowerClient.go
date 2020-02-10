/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Power
 * Used by client-side stubs.
 */

package vm


// The ``Power`` interface provides methods for managing the power state of a virtual machine.
type PowerClient interface {

    // Returns the power state information of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Power state information for the specified virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration or execution state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``System.Read``.
	Get(vmParam string) (PowerInfo, error)

    // Powers on a powered-off or suspended virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws AlreadyInDesiredState if the virtual machine is already powered on.
    // @throws Unsupported if the virtual machine does not support being powered on (e.g. marked as a template, serving as a fault-tolerance secondary virtual machine).
    // @throws UnableToAllocateResource if resources cannot be allocated for the virtual machine (e.g. physical resource allocation policy cannot be satisfied, insufficient licenses are available to run the virtual machine).
    // @throws ResourceInaccessible if resources required by the virtual machine are not accessible (e.g. virtual machine configuration files or virtual disks are on inaccessible storage, no hosts are available to run the virtual machine).
    // @throws ResourceInUse if resources required by the virtual machine are in use (e.g. virtual machine configuration files or virtual disks are locked, host containing the virtual machine is an HA failover host).
    // @throws ResourceBusy if the virtual machine is performing another operation.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Interact.PowerOn``.
	Start(vmParam string) error

    // Powers off a powered-on or suspended virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws AlreadyInDesiredState if the virtual machine is already powered off.
    // @throws ResourceBusy if the virtual machine is performing another operation.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Interact.PowerOff``.
	Stop(vmParam string) error

    // Suspends a powered-on virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws AlreadyInDesiredState if the virtual machine is already suspended.
    // @throws NotAllowedInCurrentState if the virtual machine is powered off.
    // @throws ResourceBusy if the virtual machine is performing another operation.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Interact.Suspend``.
	Suspend(vmParam string) error

    // Resets a powered-on virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is powered off or suspended.
    // @throws ResourceBusy if the virtual machine is performing another operation
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``VirtualMachine`` referenced by the parameter ``vm`` requires ``VirtualMachine.Interact.Reset``.
	Reset(vmParam string) error
}
