/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TagUsage
 * Used by client-side stubs.
 */

package tagUsage

import (
)

// The ``TagUsage`` interface provides methods to query which tags are used by policies. **Warning:** This interface is available as technical preview. It may be changed in a future release.
type TagUsageClient interface {


    // Returns information about the tags used by policies available in this vCenter server matching the FilterSpec. **Warning:** This method is available as technical preview. It may be changed in a future release.
    //
    // @param filterParam Specification for matching tags used by policies.
    // If null, the behavior is equivalent to a FilterSpec with all properties null, which means all tags used by policies match the filter.
    // @return The list of tags used by policies available on this vCenter server matching the FilterSpec.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
