/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Policies
 * Used by client-side stubs.
 */

package storage


// The ``Policies`` interface provides methods for managing the storage policies. This interface was added in vSphere API 6.7.
type PoliciesClient interface {

    // Returns information about at most 1024 visible (subject to permission checks) storage solicies availabe in vCenter. These storage policies can be used for provisioning virtual machines or disks. This method was added in vSphere API 6.7.
    //
    // @param filterParam Specification of matching storage policies for which information should be returned.
    // If null, the behavior is equivalent to a PoliciesFilterSpec with all properties null which means all storage policies match the filter
    // @return Commonly used Information about the storage policies.
    // @throws InvalidArgument if the PoliciesFilterSpec contains a value that is not supported by the server.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws UnableToAllocateResource if more than 1024 storage policies exist.
	List(filterParam *PoliciesFilterSpec) ([]PoliciesSummary, error)

    // Returns datastore compatibility summary about a specific storage policy. This method was added in vSphere API 6.7.
    //
    // @param policyParam The storage policy identifier
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.StoragePolicy``.
    // @param datastoresParam Datastores used to check compatibility against a storage policy. The number of datastores is limited to 1024.
    // The parameter must contain identifiers for the resource type: ``Datastore``.
    // @return datastore compatibility summary about a specific storage policy.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the storage policy specified does not exist.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws UnableToAllocateResource if input more than 1024 datastores.
	CheckCompatibility(policyParam string, datastoresParam map[string]bool) (PoliciesCompatibilityInfo, error)
}
