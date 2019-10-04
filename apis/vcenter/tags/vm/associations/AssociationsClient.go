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

// The ``Associations`` interface provides methods to manage associations of a single tag to multiple virtual machines. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AssociationsClient interface {


    // Attaches a tag to multiple virtual machines. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag to be assigned to the specified virtual machines.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @param vmsParam Set of identifiers of virtual machines to which the tag will be assigned.
    // The parameter must contain identifiers for the resource type: ``VirtualMachine``.
    // @return For which vms this tag attachment succeeded or failed.
    // @throws NotFound if the tag is not known to this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Add(tagParam string, vmsParam map[string]bool) (Status, error) 


    // Detaches a tag from multiple virtual machines. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag to be removed from the specified virtual machines.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @param vmsParam Set of identifiers of virtual machines from which the tag will be removed.
    // The parameter must contain identifiers for the resource type: ``VirtualMachine``.
    // @return For which vms this tag detachment succeeded or failed.
    // @throws NotFound if the tag is not known to this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Remove(tagParam string, vmsParam map[string]bool) (Status, error) 


    // Lists all virtual machine that have this tag attached. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag to be queried for its associated virtual machines.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @return The set of virtual machines associated with this tag.
    // The return value will contain identifiers for the resource type: ``VirtualMachine``.
    // @throws NotFound if the tag is not known to this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(tagParam string) (map[string]bool, error) 

}
