/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Policy
 * Used by client-side stubs.
 */

package policy

import (
)

// The ``Policy`` interface provides methods to configure the storage policies associated with the virtual machine home and/or its virtual disks.
type PolicyClient interface {


    // Updates the storage policy configuration of a virtual machine and/or its associated virtual hard disks.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Storage Policy Specification for updating the virtual machine and virtual disks.
    // @throws Error if the system reports an error while responding to the request.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user is not authenticated.
    // @throws Unauthorized if the user doesn't have the required priveleges.
    // @throws InvalidArgument if the storage policy specified is invalid.
    // @throws ResourceBusy if the virtual machine or disk is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine or disk's configuration state cannot be accessed.
    Update(vmParam string, specParam UpdateSpec) error 


    // Returns Information about Storage Policy associated with a virtual machine's home directory and/or its virtual hard disks.
    //
    // @param vmParam Virtual machine identifier
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Overview of Storage Policy associated with a virtual machine's home directory and/or its associated virtual hard disks.
    // @throws Error if the system reports an error while responding to the request.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
    Get(vmParam string) (Info, error) 

}
