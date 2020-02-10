/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: DownloadSession.
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

// Resource type for a download session.
const DownloadSession_RESOURCE_TYPE = "com.vmware.content.library.item.DownloadSession"




func downloadSessionCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(DownloadSessionModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
}

func downloadSessionCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(DownloadSessionModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
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
		map[string]int{"InvalidArgument": 400,"NotFound": 404})
}

func downloadSessionGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(DownloadSessionModelBindingType)
}

func downloadSessionGetRestMetadata() protocol.OperationRestMetadata {
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

func downloadSessionListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, ""))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, ""), reflect.TypeOf([]string{}))
}

func downloadSessionListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, ""))
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

func downloadSessionKeepAliveInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["progress"] = "Progress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionKeepAliveOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func downloadSessionKeepAliveRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["progress"] = "Progress"
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

func downloadSessionCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func downloadSessionCancelRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400})
}

func downloadSessionDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func downloadSessionDeleteRestMetadata() protocol.OperationRestMetadata {
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

func downloadSessionFailInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["client_error_message"] = bindings.NewStringType()
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["client_error_message"] = "ClientErrorMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func downloadSessionFailOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func downloadSessionFailRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["client_error_message"] = bindings.NewStringType()
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["client_error_message"] = "ClientErrorMessage"
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


