/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Compliance
 * Used by client-side stubs.
 */

package policy


// The Compliance interface provides methods that return the compliance status of virtual machine entities(virtual machine home directory and virtual disks) that specify storage policy requirements. This interface was added in vSphere API 6.7.
type ComplianceClient interface {

    // Returns the cached storage policy compliance information of a virtual machine. This method was added in vSphere API 6.7.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Virtual machine storage policy compliance Info ComplianceInfo.
    // If null, neither the virtual machine home directory nor any of it's virtual disks are associated with a storage policy.
    // @throws Error if the system reports an error while responding to the request.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
	Get(vmParam string) (*ComplianceInfo, error)

    // Returns the storage policy Compliance ComplianceInfo of a virtual machine after explicitly re-computing compliance check. This method was added in vSphere API 6.7.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param checkSpecParam Parameter specifies the entities on which storage policy compliance check is to be invoked. The storage compliance Info ComplianceInfo is returned.
    // If null, the behavior is equivalent to a ComplianceCheckSpec with CheckSpec#vmHome set to true and CheckSpec#disks populated with all disks attached to the virtual machine.
    // @return Virtual machine storage policy compliance ``Info`` class .
    // If null, neither the virtual machine home directory nor any of it's virtual disks are associated with a storage policy.
    // @throws Error if the system reports an error while responding to the request.
    // @throws ServiceUnavailable if the system is unable to communicate with a service necessary to complete the request.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
	Check(vmParam string, checkSpecParam *ComplianceCheckSpec) (*ComplianceInfo, error)
}
