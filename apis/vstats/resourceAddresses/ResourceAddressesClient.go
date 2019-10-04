/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ResourceAddresses
 * Used by client-side stubs.
 */

package resourceAddresses

import (
)

// The ``ResourceAddresses`` interface provides methods to perform resource addressing queries.
type ResourceAddressesClient interface {


    // Returns the list of Resource Addresses matching the filter parameters.
    //
    // @param filterParam Criteria for selecting records to return.
    // If null all records will be returned.
    // @return Matching resource addresses.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    List(filterParam *FilterSpec) (ListResult, error) 


    // Returns information about a specific Resource Address.
    //
    // @param idParam Resource Address ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.RsrcAddr``.
    // @return Information about the desired Resource Address.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``id`` is invalid.
    // @throws NotFound if Resource Address could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Get(idParam string) (Info, error) 

}
