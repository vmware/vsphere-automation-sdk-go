/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Type
 * Used by client-side stubs.
 */

package type_

import (
)

// The ``Type`` interface exposes the library.ItemModel types that this Content Library Service supports. 
//
//  A library item has an optional type which can be specified with the library.ItemModel#type property. For items with a type that is supported by a plugin, the Content Library Service may understand the files which are part of the library item and can produce metadata for the item. 
//
//  In other cases, uploads may require a process in which one upload implies subsequent uploads. For example, an Open Virtualization Format (OVF) package is composed of an OVF descriptor file and the associated virtual disk files. Uploading an OVF descriptor can enable the Content Library Service to understand that the complete OVF package requires additional disk files, and it can set up the transfers for the disks automatically by adding the file entries for the disks when the OVF descriptor is uploaded. 
//
//  When a type is not supported by a plugin, or the type is not specified, the Content Library Service can handle a library item in a default way, without adding metadata to the item or guiding the upload process.
type TypeClient interface {


    // Returns a array of Info instances which describe the type support plugins in this Content Library.
    // @return The array of Info instances which describe the type support plugins in this Content Library.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``ContentLibrary.TypeIntrospection``.
    List() ([]Info, error) 

}
