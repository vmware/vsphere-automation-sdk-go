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

// The ``VM`` interface provides methods to manage tag associations that have a virtual machine identifier. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMClient interface {


    // Returns tag associations that match the specified filter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param iterationParam The specification of a page to be retrieved.
    // If null, the first page will be retrieved.
    // @param filterParam The specification of matching tag associations.
    // If null, the behavior is equivalent to a FilterSpec with all properties null, which means all tag associations match the filter.
    // @return A page of the tag associations matching the FilterSpec.
    // @throws InvalidArgument if IterationSpec#marker is not a marker returned from an earlier invocation of this {\\\\@term operation).
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(iterationParam *IterationSpec, filterParam *FilterSpec) (ListResult, error) 

}
