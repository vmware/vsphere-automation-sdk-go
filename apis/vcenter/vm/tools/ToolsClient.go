/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tools
 * Used by client-side stubs.
 */

package tools

import (
)

// The ``Tools`` interface provides methods for managing VMware Tools in the guest operating system.
type ToolsClient interface {


    // Get the properties of VMware Tools.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return VMware Tools properties.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    Get(vmParam string) (Info, error) 


    // Update the properties of VMware Tools.
    //
    // @param vmParam Identifier of the virtual machine.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam The new values.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the UpdateSpec#upgradePolicy property contains a value that is not supported by the server.
    // @throws NotFound if the virtual machine is not found.
    Update(vmParam string, specParam UpdateSpec) error 


    // Begins the Tools upgrade process. To monitor the status of the Tools upgrade, clients should check the Tools status by calling Tools#get and examining ``versionStatus`` and ``runState``.
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
