/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Counters
 * Used by client-side stubs.
 */

package counters

import (
)

// The ``Counters`` interface provides methods to perform various Counter related operations. Counter is derived from metric. It applies the metric to a particular class of a resource.
type CountersClient interface {


    // Returns information about all counters matching the filter parameters.
    //
    // @param filterParam Filters the returned records.
    // When null no filtering will be applied.
    // @return List of Counters.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    List(filterParam *FilterSpec) ([]Info, error) 


    // Returns information about a specific Counter.
    //
    // @param cidParam Counter ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vstats.model.Counter``.
    // @return Information about the requested counter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if ``cid`` is invalid.
    // @throws NotFound if Counter could not be located.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
    Get(cidParam string) (Info, error) 

}
