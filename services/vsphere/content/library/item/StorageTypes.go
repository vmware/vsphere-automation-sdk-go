/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Storage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item/Storage"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"net/url"
)



// The ``Info`` class is the expanded form of FileInfo that includes details about the storage backing for a file in a library item.
type StorageInfo struct {
    // The storage backing on which this object resides. This might not be the same as the default storage backing associated with the library.
	StorageBacking library.StorageBacking
    // URIs that identify the file on the storage backing. 
    //
    //  These URIs may be specific to the backing and may need interpretation by the client. A client that understands a URI scheme in this list may use that URI to directly access the file on the storage backing. This can provide high-performance support for file manipulation.
	StorageUris []url.URL
    // A checksum for validating the content of the file. 
    //
    //  This value can be used to verify that a transfer was completed without errors.
	ChecksumInfo *FileChecksumInfo
    // The name of the file. 
    //
    //  This value will be unique within the library item for each file. It cannot be an empty string.
	Name string
    // The file size, in bytes. The file size is the storage used and not the uploaded or provisioned size. For example, when uploading a disk to a datastore, the amount of storage that the disk consumes may be different from the disk file size. When the file is not cached, the size is 0.
	Size int64
    // Indicates whether the file is on disk or not.
	Cached bool
    // The version of this file; incremented when a new copy of the file is uploaded.
	Version string
}



func storageGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageGetOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(Storage.StorageInfoBindingType), reflect.TypeOf([]Storage.StorageInfo{}))
}

func storageGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["file_name"] = "FileName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}

func storageListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(Storage.StorageInfoBindingType), reflect.TypeOf([]Storage.StorageInfo{}))
}

func storageListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}


func StorageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_backing"] = bindings.NewReferenceType(library.StorageBackingBindingType)
	fieldNameMap["storage_backing"] = "StorageBacking"
	fields["storage_uris"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
	fieldNameMap["storage_uris"] = "StorageUris"
	fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(FileChecksumInfoBindingType))
	fieldNameMap["checksum_info"] = "ChecksumInfo"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	fields["cached"] = bindings.NewBooleanType()
	fieldNameMap["cached"] = "Cached"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.storage.info", fields, reflect.TypeOf(Storage.StorageInfo{}), fieldNameMap, validators)
}

