/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: File.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package downloadsession

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``PrepareStatus`` enumeration class defines the state of the file in preparation for download.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FilePrepareStatus string

const (
    // The file hasn't been requested for preparation.
	FilePrepareStatus_UNPREPARED FilePrepareStatus = "UNPREPARED"
    // A prepare has been requested, however the server hasn't started the preparation yet.
	FilePrepareStatus_PREPARE_REQUESTED FilePrepareStatus = "PREPARE_REQUESTED"
    // A prepare has been requested and the file is in the process of being prepared.
	FilePrepareStatus_PREPARING FilePrepareStatus = "PREPARING"
    // Prepare succeeded. The file is ready for download.
	FilePrepareStatus_PREPARED FilePrepareStatus = "PREPARED"
    // Prepare failed.
	FilePrepareStatus_ERROR FilePrepareStatus = "ERROR"
)

func (p FilePrepareStatus) FilePrepareStatus() bool {
	switch p {
	case FilePrepareStatus_UNPREPARED:
		return true
	case FilePrepareStatus_PREPARE_REQUESTED:
		return true
	case FilePrepareStatus_PREPARING:
		return true
	case FilePrepareStatus_PREPARED:
		return true
	case FilePrepareStatus_ERROR:
		return true
	default:
		return false
	}
}


// The ``EndpointType`` enumeration class defines the types of endpoints used to download the file.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FileEndpointType string

const (
    // An https download endpoint.
	FileEndpointType_HTTPS FileEndpointType = "HTTPS"
    // A direct download endpoint indicating the location of the file on storage. The caller is responsible for retrieving the file from the storage location directly.
	FileEndpointType_DIRECT FileEndpointType = "DIRECT"
)

func (e FileEndpointType) FileEndpointType() bool {
	switch e {
	case FileEndpointType_HTTPS:
		return true
	case FileEndpointType_DIRECT:
		return true
	default:
		return false
	}
}


// The ``Info`` class defines the downloaded file.
type FileInfo struct {
    // The name of the file.
	Name string
    // The file size, in bytes.
	Size *int64
    // The number of bytes that have been transferred by the server so far for making this file prepared for download. This value may stay at zero till the client starts downloading the file.
	BytesTransferred int64
    // The preparation status (UNPREPARED, PREPARE_REQUESTED, PREPARING, PREPARED, ERROR) of the file.
	Status FilePrepareStatus
    // Endpoint at which the file is available for download. The value is valid only when the FileInfo#status is FilePrepareStatus#FilePrepareStatus_PREPARED.
	DownloadEndpoint *item.TransferEndpoint
    // The checksum information of the file. When the download is complete, you can retrieve the checksum from the File#get method to verify the checksum for the downloaded file.
	ChecksumInfo *item.FileChecksumInfo
    // Error message for a failed preparation when the prepare status is FilePrepareStatus#FilePrepareStatus_ERROR.
	ErrorMessage *std.LocalizableMessage
}



func fileListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
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
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
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

func filePrepareInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fields["endpoint_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.endpoint_type", reflect.TypeOf(FileEndpointType(FileEndpointType_HTTPS))))
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["endpoint_type"] = "EndpointType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func filePrepareOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FileInfoBindingType)
}

func filePrepareRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fields["endpoint_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.endpoint_type", reflect.TypeOf(FileEndpointType(FileEndpointType_HTTPS))))
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["endpoint_type"] = "EndpointType"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthorized": 403})
}

func fileGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["file_name"] = "FileName"
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
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["download_session_id"] = "DownloadSessionId"
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400})
}


func FileInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["bytes_transferred"] = bindings.NewIntegerType()
	fieldNameMap["bytes_transferred"] = "BytesTransferred"
	fields["status"] = bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.prepare_status", reflect.TypeOf(FilePrepareStatus(FilePrepareStatus_UNPREPARED)))
	fieldNameMap["status"] = "Status"
	fields["download_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
	fieldNameMap["download_endpoint"] = "DownloadEndpoint"
	fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(item.FileChecksumInfoBindingType))
	fieldNameMap["checksum_info"] = "ChecksumInfo"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["error_message"] = "ErrorMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.downloadsession.file.info", fields, reflect.TypeOf(FileInfo{}), fieldNameMap, validators)
}

