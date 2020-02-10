/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Installer
 * Used by client-side stubs.
 */

package tools


// The ``Installer`` (\\\\@term service} provides methods to install VMware Tools in the guest operating system. This interface was added in vSphere API 7.0.0.
type InstallerClient interface {

    // Get information about the VMware Tools installer. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return information about the VMware Tools installer.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
	Get(vmParam string) (InstallerInfo, error)

    // Connects the VMware Tools CD installer as a CD-ROM for the guest operating system. On Windows guest operating systems with autorun, this should cause the installer to initiate the Tools installation which will need user input to complete. On other (non-Windows) guest operating systems this will make the Tools installation available, and a a user will need to do guest-specific actions. On Linux, this includes opening an archive and running the installer. To monitor the status of the Tools install, clients should check the ``versionStatus`` and ``runState`` from Tools#get. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Virtual machine ID
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
    // @throws AlreadyInDesiredState if the VMware Tools CD is already connected.
    // @throws Error if the Tools installation fails in the guest operating system.
	Connect(vmParam string) error

    // Disconnect the VMware Tools installer CD image. This method was added in vSphere API 7.0.0.
    //
    // @param vmParam Virtual machine ID
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws NotAllowedInCurrentState if the virtual machine is not powered on.
	Disconnect(vmParam string) error
}
