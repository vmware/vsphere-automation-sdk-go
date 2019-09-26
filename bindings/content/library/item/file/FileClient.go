/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: File
 * Used by client-side stubs.
 */

package file

import (
)

// The ``File`` interface can be used to query for information on the files within a library item. Files are objects which are added to a library item through the UpdateSession and File interfaces.
type FileClient interface {


    // Retrieves the information for a single file in a library item by its name.
    //
    // @param libraryItemIdParam  Identifier of the library item whose file information should be returned.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param nameParam  Name of the file in the library item whose information should be returned.
    // @return The Info object with information on the specified file.
    // @throws NotFound  if ``library_item_id`` refers to a library item that does not exist.
    // @throws NotFound  if ``name`` refers to a file that does not exist in the library item.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``System.Read``.
    Get(libraryItemIdParam string, nameParam string) (Info, error) 


    // Lists all of the files that are stored within a given library item.
    //
    // @param libraryItemIdParam  Identifier of the library item whose files should be listed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return The array of all of the files that are stored within the given library item.
    // @throws NotFound  if ``library_item_id`` refers to a library item that does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``System.Read``.
    List(libraryItemIdParam string) ([]Info, error) 

}
