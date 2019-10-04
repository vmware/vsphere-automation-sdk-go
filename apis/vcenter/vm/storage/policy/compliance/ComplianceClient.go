/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Compliance
 * Used by client-side stubs.
 */

package compliance

import (
)

// The Compliance interface provides methods that return the compliance status of virtual machine entities(virtual machine home directory and virtual disks) that specify storage policy requirements.
type ComplianceClient interface {


    // Returns the cached storage policy compliance information of a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Virtual machine storage policy compliance Info Info.
    // If null, neither the virtual machine home directory nor any of it's virtual disks are associated with a storage policy.
    // @throws Error if the system reports an error while responding to the request.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
    Get(vmParam string) (*Info, error) 


    // Returns the storage policy Compliance Info of a virtual machine after explicitly re-computing compliance check.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param checkSpecParam Parameter specifies the entities on which storage policy compliance check is to be invoked. The storage compliance Info Info is returned.
    // If null, the behavior is equivalent to a CheckSpec with CheckSpec#vmHome set to true and CheckSpec#disks populated with all disks attached to the virtual machine.
    // @return Virtual machine storage policy compliance ``Info`` class .
    // If null, neither the virtual machine home directory nor any of it's virtual disks are associated with a storage policy.
    // @throws Error if the system reports an error while responding to the request.
    // @throws ServiceUnavailable if the system is unable to communicate with a service necessary to complete the request.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
    Check(vmParam string, checkSpecParam *CheckSpec) (*Info, error) 

}
