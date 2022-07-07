// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FileStore.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package node

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func fileStoreCopyfromremotefileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewStringType()
	fields["copy_from_remote_file_properties"] = bindings.NewReferenceType(model.CopyFromRemoteFilePropertiesBindingType)
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["copy_from_remote_file_properties"] = "CopyFromRemoteFileProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileStoreCopyfromremotefileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FilePropertiesBindingType)
}

func fileStoreCopyfromremotefileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_name"] = bindings.NewStringType()
	fields["copy_from_remote_file_properties"] = bindings.NewReferenceType(model.CopyFromRemoteFilePropertiesBindingType)
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["copy_from_remote_file_properties"] = "CopyFromRemoteFileProperties"
	paramsTypeMap["file_name"] = bindings.NewStringType()
	paramsTypeMap["copy_from_remote_file_properties"] = bindings.NewReferenceType(model.CopyFromRemoteFilePropertiesBindingType)
	paramsTypeMap["fileName"] = bindings.NewStringType()
	pathParams["file_name"] = "fileName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=copy_from_remote_file",
		"copy_from_remote_file_properties",
		"POST",
		"/api/v1/node/file-store/{fileName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func fileStoreCopytoremotefileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewStringType()
	fields["copy_to_remote_file_properties"] = bindings.NewReferenceType(model.CopyToRemoteFilePropertiesBindingType)
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["copy_to_remote_file_properties"] = "CopyToRemoteFileProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileStoreCopytoremotefileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func fileStoreCopytoremotefileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_name"] = bindings.NewStringType()
	fields["copy_to_remote_file_properties"] = bindings.NewReferenceType(model.CopyToRemoteFilePropertiesBindingType)
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["copy_to_remote_file_properties"] = "CopyToRemoteFileProperties"
	paramsTypeMap["copy_to_remote_file_properties"] = bindings.NewReferenceType(model.CopyToRemoteFilePropertiesBindingType)
	paramsTypeMap["file_name"] = bindings.NewStringType()
	paramsTypeMap["fileName"] = bindings.NewStringType()
	pathParams["file_name"] = "fileName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=copy_to_remote_file",
		"copy_to_remote_file_properties",
		"POST",
		"/api/v1/node/file-store/{fileName}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func fileStoreCreateremotedirectoryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["create_remote_directory_properties"] = bindings.NewReferenceType(model.CreateRemoteDirectoryPropertiesBindingType)
	fieldNameMap["create_remote_directory_properties"] = "CreateRemoteDirectoryProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileStoreCreateremotedirectoryOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func fileStoreCreateremotedirectoryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["create_remote_directory_properties"] = bindings.NewReferenceType(model.CreateRemoteDirectoryPropertiesBindingType)
	fieldNameMap["create_remote_directory_properties"] = "CreateRemoteDirectoryProperties"
	paramsTypeMap["create_remote_directory_properties"] = bindings.NewReferenceType(model.CreateRemoteDirectoryPropertiesBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=create_remote_directory",
		"create_remote_directory_properties",
		"POST",
		"/api/v1/node/file-store",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func fileStoreDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileStoreDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func fileStoreDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["file_name"] = "FileName"
	paramsTypeMap["file_name"] = bindings.NewStringType()
	paramsTypeMap["fileName"] = bindings.NewStringType()
	pathParams["file_name"] = "fileName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/node/file-store/{fileName}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func fileStoreGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileStoreGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FilePropertiesBindingType)
}

func fileStoreGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["file_name"] = "FileName"
	paramsTypeMap["file_name"] = bindings.NewStringType()
	paramsTypeMap["fileName"] = bindings.NewStringType()
	pathParams["file_name"] = "fileName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/node/file-store/{fileName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func fileStoreListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileStoreListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FilePropertiesListResultBindingType)
}

func fileStoreListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/node/file-store",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
