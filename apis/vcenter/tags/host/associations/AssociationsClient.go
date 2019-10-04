/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Associations
 * Used by client-side stubs.
 */

package associations

import (
)

// The ``Associations`` interface provides methods to manage associations of a single tag to multiple hosts. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AssociationsClient interface {


    // Attaches a tag to multiple hosts. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag to be assigned to the specified hosts.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @param hostsParam Set of identifiers of hosts to which the tag will be assigned.
    // The parameter must contain identifiers for the resource type: ``HostSystem``.
    // @return For which hosts this tag attachment succeeded or failed.
    // @throws NotFound if the tag is not known to this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Add(tagParam string, hostsParam map[string]bool) (Status, error) 


    // Detaches a tag from multiple hosts. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag to be removed from the specified hosts.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @param hostsParam Set of identifiers of hosts from which the tag will be removed.
    // The parameter must contain identifiers for the resource type: ``HostSystem``.
    // @return For which hosts this tag detachment succeeded or failed.
    // @throws NotFound if the tag is not known to this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Remove(tagParam string, hostsParam map[string]bool) (Status, error) 


    // Lists all hosts that have this tag attached. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag to be queried for its associated hosts.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:HostSystem``.
    // @return The set of hosts associated with this tag.
    // The return value will contain identifiers for the resource type: ``HostSystem``.
    // @throws NotFound if the tag is not known to this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(tagParam string) (map[string]bool, error) 

}
