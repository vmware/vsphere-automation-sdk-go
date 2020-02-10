/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Storage
 * Used by client-side stubs.
 */

package item

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item/Storage"
)

// ``Storage`` is a resource that represents a specific instance of a file stored on a storage backing. Unlike File, which is abstract, storage represents concrete files on the various storage backings. A file is only represented once in File, but will be represented multiple times (once for each storage backing) in ``Storage``. The ``Storage`` interface provides information on the storage backing and the specific location of the file in that backing to privileged users who want direct access to the file on the storage medium.
type StorageClient interface {

    // Retrieves the storage information for a specific file in a library item.
    //
    // @param libraryItemIdParam  Identifier of the library item whose storage information should be retrieved.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param fileNameParam  Name of the file for which the storage information should be listed.
    // @return The array of all the storage items for the given file within the given library item.
    // @throws NotFound  if the specified library item does not exist.
    // @throws NotFound  if the specified file does not exist in the given library item.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.ReadStorage``.
	Get(libraryItemIdParam string, fileNameParam string) ([]Storage.StorageInfo, error)

    // Lists all storage items for a given library item.
    //
    // @param libraryItemIdParam  Identifier of the library item whose storage information should be listed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return The array of all storage items for a given library item.
    // @throws NotFound  if the specified library item does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.ReadStorage``.
	List(libraryItemIdParam string) ([]Storage.StorageInfo, error)
}
