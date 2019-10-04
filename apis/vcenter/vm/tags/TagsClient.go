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

// The ``Tags`` interface provides methods to manage tag associations of a virtual machine. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TagsClient interface {


    // Attaches tags to a virtual machine. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Identifier of the virtual machine to which the tags will be assigned.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param tagsParam The set of identifiers of tags to be assigned.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @return For which tags this attachment succeeded or failed.
    // @throws NotFound if the virtual machine is not registered on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Add(vmParam string, tagsParam map[string]bool) (Status, error) 


    // Detaches tags from a virtual machine. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Identifier of the virtual machine from which the tags will be removed.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param tagsParam The set of identifiers of tags to be removed.
    // The parameter must contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @return For which tags this detachment succeeded or failed.
    // @throws NotFound if the virtual machine is not registered on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Remove(vmParam string, tagsParam map[string]bool) (Status, error) 


    // Lists all tags attached to the virtual machine. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Identifier of the virtual machine to be queried for its assigned tags.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return The set of tags associated with the virtual machine.
    // The return value will contain identifiers for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @throws NotFound if the virtual machine is not registered on this vCenter server.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string) (map[string]bool, error) 

}
