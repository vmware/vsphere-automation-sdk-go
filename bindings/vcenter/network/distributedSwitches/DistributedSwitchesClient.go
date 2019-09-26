/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: DistributedSwitches
 * Used by client-side stubs.
 */

package distributedSwitches

import (
)

// The DistributedSwitches interface provides methods for manipulating a vSphere Distributed Switch (VDS). **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DistributedSwitchesClient interface {


    // Returns information about at most 1000 visible (subject to permission checks) distributed switches in vCenter matching the FilterSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param filterParam Specification of matching distributed switches for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all distributed switches match the filter.
    // @return Commonly used information about the distributed switches matching the FilterSpec.
    // @throws InvalidArgument if the FilterSpec#clusters property contains a value that is not a cluster in the inventory, the FilterSpec#minMtu property is negative, or the FilterSpec#minVersion property is ill-formed.
    // @throws UnableToAllocateResource if more than 1000 distributed switches match the FilterSpec.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
