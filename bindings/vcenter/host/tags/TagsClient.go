/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tags
 * Used by client-side stubs.
 */

package tags

import (
)

// The ``Tags`` interface provides methods to manage tag associations of a host. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TagsClient interface {


    // Attaches tags to a host. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Identifier of the host to which the tags will be assigned.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @param tagsParam Set of identifiers of the tags to be assigned.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @return For which tags this attachment succeeded or failed.
    // @throws NotFound if the host is not registered on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Add(hostParam string, tagsParam map[string]bool) (Status, error) 


    // Detaches tags from a host. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Identifier of the host from which the tags will be removed.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @param tagsParam The set of identifiers of tags to be removed.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @return For which tags this detachment succeeded or failed.
    // @throws NotFound if the host is not registered on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Remove(hostParam string, tagsParam map[string]bool) (Status, error) 


    // Lists all tags attached to the host. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param hostParam Identifier of the host to be queried for its assigned tags.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @return The list of tags associated with the host.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @throws NotFound if the host is not registered on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(hostParam string) (map[string]bool, error) 

}
