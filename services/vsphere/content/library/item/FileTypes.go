/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: File.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``ChecksumAlgorithm`` enumeration class defines the valid checksum algorithms.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FileChecksumAlgorithm string

const (
    // Checksum algorithm: SHA-1
	FileChecksumAlgorithm_SHA1 FileChecksumAlgorithm = "SHA1"
    // Checksum algorithm: MD5
	FileChecksumAlgorithm_MD5 FileChecksumAlgorithm = "MD5"
    // Checksum algorithm: SHA-256. This constant field was added in vSphere API 6.8.
	FileChecksumAlgorithm_SHA256 FileChecksumAlgorithm = "SHA256"
    // Checksum algorithm: SHA-512. This constant field was added in vSphere API 6.8.
	FileChecksumAlgorithm_SHA512 FileChecksumAlgorithm = "SHA512"
)

func (c FileChecksumAlgorithm) FileChecksumAlgorithm() bool {
	switch c {
	case FileChecksumAlgorithm_SHA1:
		return true
	case FileChecksumAlgorithm_MD5:
		return true
	case FileChecksumAlgorithm_SHA256:
		return true
	case FileChecksumAlgorithm_SHA512:
		return true
	default:
		return false
	}
}


// Provides checksums for a FileInfo object.
type FileChecksumInfo struct {
    // The checksum algorithm (SHA1, MD5, SHA256, SHA512) used to calculate the checksum.
	Algorithm *FileChecksumAlgorithm
    // The checksum value calculated with FileChecksumInfo#algorithm.
	Checksum string
}

// The ``Info`` class provides information about a file in Content Library Service storage. 
//
//  A file is an actual stored object for a library item. An item will have zero files initially, but one or more can be uploaded to the item.
type FileInfo struct {
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



func fileGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["name"] = bindings.NewStringType()
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FileInfoBindingType)
}

func fileGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["name"] = bindings.NewStringType()
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["name"] = "Name"
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

func fileListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(FileInfoBindingType), reflect.TypeOf([]FileInfo{}))
}

func fileListRestMetadata() protocol.OperationRestMetadata {
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


func FileChecksumInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["algorithm"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.file.checksum_algorithm", reflect.TypeOf(FileChecksumAlgorithm(FileChecksumAlgorithm_SHA1))))
	fieldNameMap["algorithm"] = "Algorithm"
	fields["checksum"] = bindings.NewStringType()
	fieldNameMap["checksum"] = "Checksum"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.file.checksum_info", fields, reflect.TypeOf(FileChecksumInfo{}), fieldNameMap, validators)
}

func FileInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
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
	return bindings.NewStructType("com.vmware.content.library.item.file.info", fields, reflect.TypeOf(FileInfo{}), fieldNameMap, validators)
}

