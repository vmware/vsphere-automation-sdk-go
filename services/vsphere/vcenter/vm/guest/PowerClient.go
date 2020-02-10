/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Power
 * Used by client-side stubs.
 */

package guest


// The ``Power`` interface provides methods for managing the guest operating system power state of a virtual machine. This interface was added in vSphere API 6.7.
type PowerClient interface {

    // Returns information about the guest operating system power state. This method was added in vSphere API 6.7.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Guest OS powerstate information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
	Get(vmParam string) (PowerInfo, error)

    // Issues a request to the guest operating system asking it to perform a clean shutdown of all services. This request returns immediately and does not wait for the guest operating system to complete the operation. This method was added in vSphere API 6.7.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws AlreadyInDesiredState if the virtual machine is not powered on.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws NotAllowedInCurrentState if the virtual machine is suspended.
    // @throws ResourceBusy if the virtual machine is performing another operation.
    // @throws Unsupported if the virtual machine does not support being powered on (e.g. marked as a template, serving as a fault-tolerance secondary virtual machine).
	Shutdown(vmParam string) error

    // Issues a request to the guest operating system asking it to perform a reboot. This request returns immediately and does not wait for the guest operating system to complete the operation. This method was added in vSphere API 6.7.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws ResourceBusy if the virtual machine is performing another operation.
    // @throws Unsupported if the virtual machine does not support being powered on (e.g. marked as a template, serving as a fault-tolerance secondary virtual machine).
	Reboot(vmParam string) error

    // Issues a request to the guest operating system asking it to perform a suspend operation. This method was added in vSphere API 6.7.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws AlreadyInDesiredState if the virtual machine is suspended.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if VMware Tools is not running.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws ResourceBusy if the virtual machine is performing another operation.
    // @throws Unsupported if the virtual machine does not support being powered on (e.g. marked as a template, serving as a fault-tolerance secondary virtual machine).
	Standby(vmParam string) error
}
