/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package VM

import (
)

// The ``VM`` class provides methods related to query virtual machines of given compliance statuses.
type VMClient interface {


    // Returns compliance information about at most 1000 virtual machines matching the filter FilterSpec. If there are no virtual machines matching the FilterSpec an empty List is returned. Virtual machines without storage policy association are not returned.
    //
    // @param filterParam compliance status of matching virtual machines for which information should be returned.
    // @return compliance information about virtual machines matching the filter FilterSpec.
    // The key in the return value map will be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the FilterSpec#status property contains a value that is not supported by the server.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have the required privileges.
    // @throws Unsupported if the API is invoked against vCenter Server version is less than 6.5
    // @throws UnableToAllocateResource If more than 1000 results match the FilterSpec
    List(filterParam FilterSpec) (map[string]Info, error) 

}
