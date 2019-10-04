/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Network
 * Used by client-side stubs.
 */

package network

import (
)

// The Network interface provides methods for manipulating a vCenter Server network.
type NetworkClient interface {


    // Returns information about at most 1000 visible (subject to permission checks) networks in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching networks for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all networks match the filter.
    // @return Commonly used information about the networks matching the FilterSpec.
    // @throws InvalidArgument if the FilterSpec#types property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 1000 networks match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
