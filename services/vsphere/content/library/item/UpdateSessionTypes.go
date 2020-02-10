/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: UpdateSession.
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

// Resource type for an update session.
const UpdateSession_RESOURCE_TYPE = "com.vmware.content.library.item.UpdateSession"




func updateSessionCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(UpdateSessionModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
}

func updateSessionCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(UpdateSessionModelBindingType)
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
		map[string]int{"InvalidArgument": 400,"InvalidElementType": 400,"NotFound": 404,"ResourceBusy": 500})
}

func updateSessionGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(UpdateSessionModelBindingType)
}

func updateSessionGetRestMetadata() protocol.OperationRestMetadata {
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

func updateSessionListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, ""))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, ""), reflect.TypeOf([]string{}))
}

func updateSessionListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, ""))
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

func updateSessionCompleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionCompleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSessionCompleteRestMetadata() protocol.OperationRestMetadata {
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

func updateSessionKeepAliveInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["client_progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["client_progress"] = "ClientProgress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionKeepAliveOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSessionKeepAliveRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["client_progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["client_progress"] = "ClientProgress"
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

func updateSessionCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSessionCancelRestMetadata() protocol.OperationRestMetadata {
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

func updateSessionFailInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["client_error_message"] = bindings.NewStringType()
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["client_error_message"] = "ClientErrorMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionFailOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSessionFailRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["client_error_message"] = bindings.NewStringType()
	fieldNameMap["update_session_id"] = "UpdateSessionId"
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

func updateSessionDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSessionDeleteRestMetadata() protocol.OperationRestMetadata {
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

func updateSessionUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["update_spec"] = bindings.NewReferenceType(UpdateSessionModelBindingType)
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["update_spec"] = "UpdateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSessionUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSessionUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["update_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.UpdateSession"}, "")
	fields["update_spec"] = bindings.NewReferenceType(UpdateSessionModelBindingType)
	fieldNameMap["update_session_id"] = "UpdateSessionId"
	fieldNameMap["update_spec"] = "UpdateSpec"
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"InvalidArgument": 400})
}


