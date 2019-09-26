/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Items
 * Used by client-side stubs.
 */

package items

import (
)

// The ``Items`` provides methods to create, list, update, and delete configuration items that are stored in a namespace in the Settings Store of HVC. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ItemsClient interface {


    // Returns a map of item values corresponding to the keys in the namespace. The readPrivilege defined in the namespace is checked before the operation is initiated. If keys is not passed in, the entire map of KV pairs are returned. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam the targetted namespace
    // @param keysParam 
    // This parameter is the keys of items to be listed. If not passed in, the entire map of KV items is returned
    // @return the map of key value items in the namespace
    // @throws Unauthorized If the user is not authorized.
    // @throws NotFound if the namespace is not found
    // @throws Error if the system reports an error while responding to the request.
    List(namespaceParam string, keysParam []string) (map[string]string, error) 


    // Put a map of key value pairs in the namespace. This overrides any existing key/value if they already exist in the namespace. The writePrivilege defined in the namespace is checked before the operation is initiated. When the value is not passed in the map items, the item in the namespace is deleted. When an empty (0 length) String, "", is passed in as the value, the value is set to "". **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param namespaceParam the targetted namespace
    // @param itemsParam the input list of key value items. When the value is not passed in the map items, the item in the namespace is deleted. When an empty (0 length) String, "", is passed in as the value, the value is set to ""
    // @throws Unauthorized if the user is not authorized
    // @throws NotFound if the namespace is not found
    // @throws Error if the system reports an error while responding to the request
    Put(namespaceParam string, itemsParam map[string]*string) error 

}
