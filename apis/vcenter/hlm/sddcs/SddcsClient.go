/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Sddcs
 * Used by client-side stubs.
 */

package sddcs

import (
)

// The ``Sddcs`` interface provides the method to list the Sddcs under an organization. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SddcsClient interface {


    // Lists all Sddcs that are under the same organization as the local node. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param filterParam Specification on how to filter the Sddcs
    // If null, the behavior is equivalent to a listing all Sddcs.
    // @return the list of Sddcs
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if the system reports an error while responding to the request.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
