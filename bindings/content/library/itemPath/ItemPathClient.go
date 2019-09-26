/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ItemPath
 * Used by client-side stubs.
 */

package itemPath

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library"
)

// The ``ItemPath`` interface provides methods for managing library items.
type ItemPathClient interface {


    // Returns the library.ItemModel for an item that contains the file specified by the given datastore path.
    //
    // @param datastorePathParam  Path of a content library file on a datastore.
    // @return The library.ItemModel instance of an item that contains the file with the given ``datastore_path``.
    // @throws NotFound  if no file with the given ``datastore_path`` exists in a content library.
    // @throws InvalidArgument  if the given ``datastore_path`` does not point to a content library file.
    // @throws Error  If the Iso Service is disabled. (This method is added to support the lookup needed by the Iso Service UI)
    GetByDatastorePath(datastorePathParam string) (library.ItemModel, error) 

}
