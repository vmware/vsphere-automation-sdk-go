/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: File.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package updatesession

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``SourceType`` enumeration class defines how the file content is retrieved.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FileSourceType string

const (
    // No source type has been requested.
	FileSourceType_NONE FileSourceType = "NONE"
    // The client is uploading content using HTTP(S) PUT requests.
	FileSourceType_PUSH FileSourceType = "PUSH"
    // The server is pulling content from a URL. The URL scheme can be ``http``, ``https``, ``file``, or ``ds``.
	FileSourceType_PULL FileSourceType = "PULL"
)

func (s FileSourceType) FileSourceType() bool {
	switch s {
	case FileSourceType_NONE:
		return true
	case FileSourceType_PUSH:
		return true
	case FileSourceType_PULL:
		return true
	default:
		return false
	}
}


// The ``AddSpec`` class describes the properties of the file to be uploaded.
type FileAddSpec struct {
    // The name of the file being uploaded.
	Name string
    // The source type (NONE, PUSH, PULL) from which the file content will be retrieved.
	SourceType FileSourceType
    // Location from which the Content Library Service will fetch the file, rather than requiring a client to upload the file.
	SourceEndpoint *item.TransferEndpoint
    // The file size, in bytes.
	Size *int64
    // The checksum of the file. If specified, the server will verify the checksum once the file is received. If there is a mismatch, the upload will fail. For ova files, this value should not be set.
	ChecksumInfo *item.FileChecksumInfo
}

// The ``Info`` class defines the uploaded file.
type FileInfo struct {
    // The name of the file.
	Name string
    // The source type (NONE, PUSH, PULL) from which the file is being retrieved. This may be FileSourceType#FileSourceType_NONE if the file is not being changed.
	SourceType FileSourceType
    // The file size, in bytes as received by the server. This property is guaranteed to be set when the server has completely received the file.
	Size *int64
    // The checksum information of the file received by the server.
	ChecksumInfo *item.FileChecksumInfo
    // A source endpoint from which to retrieve the file.
	SourceEndpoint *item.TransferEndpoint
    // An upload endpoint to which the client can push the content.
	UploadEndpoint *item.TransferEndpoint
    // The number of bytes of this file that have been received by the server.
	BytesTransferred int64
    // The transfer status (WAITING_FOR_TRANSFER, TRANSFERRING, READY, VALIDATING, ERROR) of this file.
	Status item.TransferStatus
    // Details about the transfer error.
	ErrorMessage *std.LocalizableMessage
    // Whether or not the file will be kept in storage upon update session completion. The flag is true for most files, and false for metadata files such as manifest and certificate file of update session with library item type OVF. Any file with FileInfo#keepInStorage set to false will not show up in the list of files returned from File#list upon update session completion. This property was added in vSphere API 6.8.
	KeepInStorage *bool
}

// The ``ValidationError`` class defines the validation error of a file in the session.
type FileValidationError struct {
    // The name of the file.
	Name string
    // A message indicating why the file was considered invalid.
	ErrorMessage std.LocalizableMessage
}

// The ``ValidationResult`` class defines the result of validating the files in the session.
type FileValidationResult struct {
    // Whether the validation was succesful or not. In case of errors, the FileValidationResult#missingFiles and FileValidationResult#invalidFiles will contain at least one entry.
	HasErrors bool
    // A map with bool value containing the names of the files that are required but the client hasn't added.
	MissingFiles map[string]bool
    // A array containing the files that have been identified as invalid and details about the error.
	InvalidFiles []FileValidationError
}



func fileValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FileValidationResultBindingType)
}

func fileValidateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400})
}

func fileAddInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["file_spec"] = bindings.NewReferenceType(FileAddSpecBindingType)
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["file_spec"] = "FileSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileAddOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FileInfoBindingType)
}

func fileAddRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["file_spec"] = bindings.NewReferenceType(FileAddSpecBindingType)
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["file_spec"] = "FileSpec"
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
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}

func fileRemoveInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileRemoveOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func fileRemoveRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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

func fileListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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

func fileGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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


func FileAddSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["source_type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.file.source_type", reflect.TypeOf(FileSourceType(FileSourceType_NONE)))
	fieldNameMap["source_type"] = "SourceType"
	fields["source_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
	fieldNameMap["source_endpoint"] = "SourceEndpoint"
	fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(item.FileChecksumInfoBindingType))
	fieldNameMap["checksum_info"] = "ChecksumInfo"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("source_type",
		map[string][]bindings.FieldData{
			"PULL": []bindings.FieldData{
				bindings.NewFieldData("source_endpoint", true),
			},
			"NONE": []bindings.FieldData{},
			"PUSH": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.add_spec", fields, reflect.TypeOf(FileAddSpec{}), fieldNameMap, validators)
}

func FileInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["source_type"] = bindings.NewEnumType("com.vmware.content.library.item.updatesession.file.source_type", reflect.TypeOf(FileSourceType(FileSourceType_NONE)))
	fieldNameMap["source_type"] = "SourceType"
	fields["size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["size"] = "Size"
	fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(item.FileChecksumInfoBindingType))
	fieldNameMap["checksum_info"] = "ChecksumInfo"
	fields["source_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
	fieldNameMap["source_endpoint"] = "SourceEndpoint"
	fields["upload_endpoint"] = bindings.NewOptionalType(bindings.NewReferenceType(item.TransferEndpointBindingType))
	fieldNameMap["upload_endpoint"] = "UploadEndpoint"
	fields["bytes_transferred"] = bindings.NewIntegerType()
	fieldNameMap["bytes_transferred"] = "BytesTransferred"
	fields["status"] = bindings.NewEnumType("com.vmware.content.library.item.transfer_status", reflect.TypeOf(item.TransferStatus(item.TransferStatus_WAITING_FOR_TRANSFER)))
	fieldNameMap["status"] = "Status"
	fields["error_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["error_message"] = "ErrorMessage"
	fields["keep_in_storage"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["keep_in_storage"] = "KeepInStorage"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("source_type",
		map[string][]bindings.FieldData{
			"PULL": []bindings.FieldData{
				bindings.NewFieldData("source_endpoint", true),
			},
			"PUSH": []bindings.FieldData{
				bindings.NewFieldData("upload_endpoint", true),
			},
			"NONE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.info", fields, reflect.TypeOf(FileInfo{}), fieldNameMap, validators)
}

func FileValidationErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["error_message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["error_message"] = "ErrorMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.validation_error", fields, reflect.TypeOf(FileValidationError{}), fieldNameMap, validators)
}

func FileValidationResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["has_errors"] = bindings.NewBooleanType()
	fieldNameMap["has_errors"] = "HasErrors"
	fields["missing_files"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["missing_files"] = "MissingFiles"
	fields["invalid_files"] = bindings.NewListType(bindings.NewReferenceType(FileValidationErrorBindingType), reflect.TypeOf([]FileValidationError{}))
	fieldNameMap["invalid_files"] = "InvalidFiles"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.updatesession.file.validation_result", fields, reflect.TypeOf(FileValidationResult{}), fieldNameMap, validators)
}

