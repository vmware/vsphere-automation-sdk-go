/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: IsoPath
 * Used by client-side stubs.
 */

package isoPath

import (
)

// Provides an interface to perform datastore path related operations on a library item. 
//
//  This is an API that will let its client lookup the datastore path of a .iso file within an ISO content library item. 
type IsoPathClient interface {


    // Returns the datastore path of a .iso file within an ISO content library item.
    //
    // @param libraryItemIdParam The identifier of the ISO item that has a .iso file.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return The datastore path of the .iso file in the library item identified by ``library_item_id``.
    // @throws NotFound If the item identified by ``library_item_id`` is not found.
    // @throws InvalidArgument If the item identified by ``library_item_id`` is not a valid library item of type 'iso'.
    // @throws InternalServerError If the Iso Service encounters an unexpected or unknown error.
    // @throws Error If the Iso Service is disabled.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``System.Read``.
    GetDatastorePath(libraryItemIdParam string) (string, error) 

}
