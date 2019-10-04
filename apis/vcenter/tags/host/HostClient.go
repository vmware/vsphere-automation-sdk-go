/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Host
 * Used by client-side stubs.
 */

package host

import (
)

// The ``Host`` interface provides methods to manage which tags can be associated with hosts. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostClient interface {


    // Returns information about the tags that are associated with hosts on this vCenter Server, where the tags need to match FilterSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param filterParam The specification of matching tag associations.
    // If null, the behavior is equivalent to a FilterSpec with all properties null, which means all tag associations match the filter.
    // @return ListResult with list of tag summaries, see Host.ListResult.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) (ListResult, error) 


    // Deletes any existing associations of this tag with hosts. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag of which the associations with hosts will be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @throws NotFound if a tag with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(tagParam string) error 

}
