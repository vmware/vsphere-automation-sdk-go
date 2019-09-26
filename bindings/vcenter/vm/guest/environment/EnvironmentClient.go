/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Environment
 * Used by client-side stubs.
 */

package environment

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/vm/guest"
)

// The ``Environment`` interface provides methods to manage environment variables in the guest operating system. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EnvironmentClient interface {


    // Reads a single environment variable from the guest operating system. 
    //
    //  If the authentication uses guest.Credentials#interactiveSession, then the environment being read will be that of the user logged into the desktop. Otherwise it's the environment of the system user. 
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine to perform the operation on.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param credentialsParam The guest authentication data. See guest.Credentials.
    // @param nameParam The name of the environment variable to be read.
    // @return The value of the ``name`` environment variable.
    // Unset if the ``name`` environment variable is not set.
    // @throws NotFound if the virtual machine ``vm`` is not found.
    // @throws NotFound if the environment variable ``name`` is not found.
    // @throws Unsupported if the operation is not supported by the VMware tools in the guest OS.
    // @throws Unsupported if the operation is disabled by the VMware tools in the guest OS.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not running.
    // @throws ResourceBusy if the virtual machine is busy.
    // @throws Unauthenticated if the ``credentials`` are not valid.
    // @throws ServiceUnavailable if the VMware tools are not running.
    Get(vmParam string, credentialsParam guest.Credentials, nameParam string) (string, error) 


    // Reads a list of environment variables from the guest operating system. 
    //
    //  If the authentication uses guest.Credentials#interactiveSession, then the environment being read will be that of the user logged into the desktop. Otherwise it's the environment of the system user. 
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Virtual machine to perform the operation on.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param credentialsParam The guest authentication data. See guest.Credentials.
    // @param namesParam The names of the variables to be read. If the map with bool value is empty, then all the environment variables are returned.
    // @return Mapping from environment variable names to environment variable values, or all environment variables if nothing is specified. If any specified environment variable contained in ``names`` is not set, then nothing is returned for that variable.
    // @throws NotFound if the virtual machine ``vm`` is not found.
    // @throws Unsupported if the operation is not supported by the VMware tools in the guest OS.
    // @throws Unsupported if the operation is disabled by the VMware tools in the guest OS.
    // @throws NotAllowedInCurrentState if the virtual machine ``vm`` is not running.
    // @throws ResourceBusy if the virtual machine is busy.
    // @throws Unauthenticated if the ``credentials`` are not valid.
    // @throws ServiceUnavailable if the VMware tools are not running.
    List(vmParam string, credentialsParam guest.Credentials, namesParam map[string]bool) (map[string]string, error) 

}
