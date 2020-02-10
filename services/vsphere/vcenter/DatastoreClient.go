/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Datastore
 * Used by client-side stubs.
 */

package vcenter


// The Datastore interface provides methods for manipulating a datastore.
type DatastoreClient interface {

    // Retrieves information about the datastore indicated by ``datastore``.
    //
    // @param datastoreParam Identifier of the datastore for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``Datastore``.
    // @return information about the datastore.
    // @throws NotFound if the datastore indicated by ``datastore`` does not exist.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Get(datastoreParam string) (DatastoreInfo, error)

    // Returns information about at most 2500 visible (subject to permission checks) datastores in vCenter matching the DatastoreFilterSpec.
    //
    // @param filterParam Specification of matching datastores for which information should be returned.
    // If null, the behavior is equivalent to a DatastoreFilterSpec with all properties null which means all datastores match the filter.
    // @return Commonly used information about the datastores matching the DatastoreFilterSpec.
    // @throws InvalidArgument if the DatastoreFilterSpec#types property contains a value that is not supported by the server.
    // @throws InvalidArgument if the DatastoreFilterSpec#types property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 2500 datastores match the DatastoreFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *DatastoreFilterSpec) ([]DatastoreSummary, error)
}
