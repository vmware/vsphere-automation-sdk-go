/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Namespaces
 * Used by client-side stubs.
 */

package namespaces

import (
)

// The ``Namespaces`` provides methods to create, read, update, and delete settings store namespaces. A namespace is a logical container that holds a group of configuration items. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NamespacesClient interface {


    // Lists all the namespaces in Settings Service. System.Read privilege is required. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return a list of namespaces and their corresponding infos
    // @throws Unauthorized if the user is unauthorized to perform this operation.
    // @throws Error if the system reports an error while responding to the request.
    List() ([]Info, error) 


    // Retrieves the details of the namespace. System.Read privilege is required. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam identifier of the namespace
    // @return specific Info a namespace
    // @throws Unauthorized if the user is unauthorized to perform this operation.
    // @throws Error if the system reports an error while responding to the request.
    Get(namespaceParam string) (Info, error) 


    // Creates a namespace. Requires the SettingsStore.Manage privilege. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam the specification used for creating a namespace.
    // @throws AlreadyExists if the namespace with the specified name already exists
    // @throws Unauthorized if user is unauthorized to perform this operation
    // @throws InvalidArgument if any of privileges specified in the createspec does not exist
    // @throws Error if the system reports an error while responding to the request
    Create(specParam CreateSpec) error 


    // Deletes a namespace along with its associated item data. Requires the SettingsStore.Manage privilege. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam namespace identifier
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.settings.namespaces``.
    // @throws Unauthorized if user is unauthorized to perform this operation
    // @throws Error if the system reports an error while responding to the request
    Delete(namespaceParam string) error 


    // Update a namespace with a changing configuration. Requires the SettingsStore.Manage privilege. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam namespace identifier
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.settings.namespaces``.
    // @param specParam specification to update the namespace
    // @throws NotFound if the namespace is not found
    // @throws Unauthorized if user is unauthorized to perform this operation
    // @throws InvalidArgument if any of privileges specified in the createspec does not exist
    // @throws Error if the system reports an error while responding to the request
    Update(namespaceParam string, specParam UpdateSpec) error 

}
