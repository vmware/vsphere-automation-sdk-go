/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Datacenter
 * Used by client-side stubs.
 */

package datacenter

import (
)

// The ``Datacenter`` interface provides methods to manage datacenters in the vCenter Server.
type DatacenterClient interface {


    // Create a new datacenter in the vCenter inventory
    //
    // @param specParam Specification for the new datacenter to be created.
    // @return The identifier of the newly created datacenter
    // The return value will be an identifier for the resource type: ``Datacenter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws AlreadyExists if the datacenter with the same name is already present.
    // @throws InvalidArgument if the datacenter name is empty or invalid as per the underlying implementation.
    // @throws InvalidArgument if the folder is not specified and the system cannot choose a suitable one.
    // @throws NotFound if the datacenter folder cannot be found.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Create(specParam CreateSpec) (string, error) 


    // Delete an empty datacenter from the vCenter Server
    //
    // @param datacenterParam Identifier of the datacenter to be deleted.
    // The parameter must be an identifier for the resource type: ``Datacenter``.
    // @param forceParam If true, delete the datacenter even if it is not empty.
    // If null a errors.ResourceInUse exception will be reported if the datacenter is not empty. This is the equivalent of passing the value false.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if there is no datacenter associated with ``datacenter`` in the system.
    // @throws ResourceInUse if the datacenter associated with ``datacenter`` is not empty.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(datacenterParam string, forceParam *bool) error 


    // Returns information about at most 1000 visible (subject to permission checks) datacenters in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching datacenters for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all datacenters match the filter.
    // @return Commonly used information about the datacenters matching the FilterSpec.
    // @throws UnableToAllocateResource if more than 1000 datacenters match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 


    // Retrieves information about the datacenter corresponding to ``datacenter``.
    //
    // @param datacenterParam Identifier of the datacenter.
    // The parameter must be an identifier for the resource type: ``Datacenter``.
    // @return The Info instances that corresponds to the ``datacenter``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if there is no datacenter associated with ``datacenter`` in the system.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(datacenterParam string) (Info, error) 

}
