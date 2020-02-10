/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Tools
 * Used by client-side stubs.
 */

package vm


// The ``Tools`` interface provides methods for managing VMware Tools in the guest operating system. This interface was added in vSphere API 7.0.0.
type ToolsClient interface {

    // Get the properties of VMware Tools. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return VMware Tools properties.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
	Get(vmParam string) (ToolsInfo, error)

    // Update the properties of VMware Tools. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam The new values.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the ToolsUpdateSpec#upgradePolicy property contains a value that is not supported by the server.
    // @throws NotFound if the virtual machine is not found.
	Update(vmParam string, specParam ToolsUpdateSpec) error

    // Begins the Tools upgrade process. To monitor the status of the Tools upgrade, clients should check the Tools status by calling Tools#get and examining ``versionStatus`` and ``runState``. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param commandLineOptionsParam Command line options passed to the installer to modify the installation procedure for Tools.
    // Set if any additional options are desired.
    // @throws NotFound if the virtual machine is not found.
    // @throws ServiceUnavailable if the VMware Tools are not running.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws AlreadyInDesiredState is an upgrade is already in progress.
    // @throws Error if the upgrade process fails inside the guest operating system.
	Upgrade(vmParam string, commandLineOptionsParam *string) error
}
