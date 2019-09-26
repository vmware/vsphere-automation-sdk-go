/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Members
 * Used by client-side stubs.
 */

package members

import (
)

// The ``Members`` interface provides methods to manage the membership of virtual machines in a cluster module. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MembersClient interface {


    // Adds virtual machines to the module. These virtual machines are required to be in the same vCenter cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param moduleParam Identifier of the module to which the specified virtual machines are added.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.cluster.modules``.
    // @param vmsParam Set of identifiers of virtual machines that will be added to the module.
    // The parameter must contain identifiers for the resource type: ``VirtualMachine``.
    // @return Whether the addition of members to the module succeeded or failed.
    // @throws NotFound if the module is not known to this vCenter server.
    Add(moduleParam string, vmsParam map[string]bool) (Status, error) 


    // Removes virtual machines from the module. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param moduleParam Identifier of the module from which the specified virtual machines are removed.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.cluster.modules``.
    // @param vmsParam Set of identifiers of virtual machines that will be removed from the module.
    // The parameter must contain identifiers for the resource type: ``VirtualMachine``.
    // @return Whether the removal of members from the module succeeded or failed.
    // @throws NotFound if the module is not known to this vCenter server.
    Remove(moduleParam string, vmsParam map[string]bool) (Status, error) 


    // Returns the virtual machines that are members of the module. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param moduleParam Identifier of the module to be queried for its virtual machines.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.cluster.modules``.
    // @return The virtual machines that are members of the module.
    // @throws NotFound if the module is not known to this vCenter server.
    Get(moduleParam string) (GetResult, error) 

}
