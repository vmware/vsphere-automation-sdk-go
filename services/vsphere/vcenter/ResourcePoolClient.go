/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ResourcePool
 * Used by client-side stubs.
 */

package vcenter


// The ResourcePool interface provides methods for manipulating a vCenter Server resource pool. 
//
//  This interface does not include virtual appliances in the inventory of resource pools even though part of the behavior of a virtual appliance is to act like a resource pool.
type ResourcePoolClient interface {

    // Retrieves information about the resource pool indicated by ``resource_pool``.
    //
    // @param resourcePoolParam Identifier of the resource pool for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``ResourcePool``.
    // @return Information about the resource pool.
    // @throws NotFound If the resource pool indicated by ``resource_pool`` does not exist.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``ResourcePool`` referenced by the parameter ``resource_pool`` requires ``System.Read``.
	Get(resourcePoolParam string) (ResourcePoolInfo, error)

    // Returns information about at most 1000 visible (subject to permission checks) resource pools in vCenter matching the ResourcePoolFilterSpec.
    //
    // @param filterParam Specification of matching resource pools for which information should be returned.
    // If null, the behavior is equivalent to a ResourcePoolFilterSpec with all properties null which means all resource pools match the filter.
    // @return Commonly used information about the resource pools matching the ResourcePoolFilterSpec.
    // @throws UnableToAllocateResource If more than 1000 resource pools match the ResourcePoolFilterSpec.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
	List(filterParam *ResourcePoolFilterSpec) ([]ResourcePoolSummary, error)

    // Creates a resource pool. This method was added in vSphere API 7.0.0.
    //
    // @param specParam Specification of the new resource pool to be created, see ResourcePoolCreateSpec.
    // @return The identifier of the newly created resource pool.
    // The return value will be an identifier for the resource type: ``ResourcePool``.
    // @throws Error If the system reports an error while responding to the request.
    // @throws InvalidArgument If a parameter passed in the spec is invalid.
    // @throws NotFound If the resource specified in parent could not be found
    // @throws ResourceInaccessible If the specified resource in parent is not accessible.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to create the resource pool could not be allocated.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``ResourcePool`` referenced by the property ResourcePoolCreateSpec#parent requires ``Resource.CreatePool``.
	Create(specParam ResourcePoolCreateSpec) (string, error)

    // Deletes a resource pool. This method was added in vSphere API 7.0.0.
    //
    // @param resourcePoolParam Identifier of the resource pool to be deleted.
    // The parameter must be an identifier for the resource type: ``ResourcePool``.
    // @throws Error If the system reports an error while responding to the request.
    // @throws NotFound If the resource pool is not found.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
    // @throws Unsupported If the resource pool is a root resource pool.
	Delete(resourcePoolParam string) error

    // Updates the configuration of a resource pool. This method was added in vSphere API 7.0.0.
    //
    // @param resourcePoolParam Identifier of the resource pool.
    // The parameter must be an identifier for the resource type: ``ResourcePool``.
    // @param specParam Specification for updating the configuration of the resource pool.
    // @throws Error If the system reports an error while responding to the request.
    // @throws InvalidArgument If any of the specified parameters is invalid.
    // @throws NotFound If the resource pool is not found.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource If any of the resources needed to reconfigure the resource pool could not be allocated.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``ResourcePool`` referenced by the parameter ``resource_pool`` requires ``Resource.EditPool``.
	Update(resourcePoolParam string, specParam ResourcePoolUpdateSpec) error
}
