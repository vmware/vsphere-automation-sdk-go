// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Tasks.
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

func tasksCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task_id"] = bindings.NewStringType()
	fieldNameMap["task_id"] = "TaskId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tasksCancelRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task_id"] = bindings.NewStringType()
	fieldNameMap["task_id"] = "TaskId"
	paramsTypeMap["task_id"] = bindings.NewStringType()
	paramsTypeMap["taskId"] = bindings.NewStringType()
	pathParams["task_id"] = "taskId"
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
		"action=cancel",
		"",
		"POST",
		"/api/v1/node/tasks/{taskId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tasksDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task_id"] = bindings.NewStringType()
	fieldNameMap["task_id"] = "TaskId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tasksDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task_id"] = bindings.NewStringType()
	fieldNameMap["task_id"] = "TaskId"
	paramsTypeMap["task_id"] = bindings.NewStringType()
	paramsTypeMap["taskId"] = bindings.NewStringType()
	pathParams["task_id"] = "taskId"
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
		"/api/v1/node/tasks/{taskId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tasksGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task_id"] = bindings.NewStringType()
	fields["suppress_redirect"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["task_id"] = "TaskId"
	fieldNameMap["suppress_redirect"] = "SuppressRedirect"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ApplianceManagementTaskPropertiesBindingType)
}

func tasksGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["task_id"] = bindings.NewStringType()
	fields["suppress_redirect"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["task_id"] = "TaskId"
	fieldNameMap["suppress_redirect"] = "SuppressRedirect"
	paramsTypeMap["suppress_redirect"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["task_id"] = bindings.NewStringType()
	paramsTypeMap["taskId"] = bindings.NewStringType()
	pathParams["task_id"] = "taskId"
	queryParams["suppress_redirect"] = "suppress_redirect"
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
		"/api/v1/node/tasks/{taskId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tasksListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["request_method"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["request_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["request_uri"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["request_method"] = "RequestMethod"
	fieldNameMap["request_path"] = "RequestPath"
	fieldNameMap["request_uri"] = "RequestUri"
	fieldNameMap["status"] = "Status"
	fieldNameMap["user"] = "User"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ApplianceManagementTaskListResultBindingType)
}

func tasksListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["request_method"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["request_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["request_uri"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["request_method"] = "RequestMethod"
	fieldNameMap["request_path"] = "RequestPath"
	fieldNameMap["request_uri"] = "RequestUri"
	fieldNameMap["status"] = "Status"
	fieldNameMap["user"] = "User"
	paramsTypeMap["request_uri"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["request_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["user"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["request_method"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["request_path"] = "request_path"
	queryParams["request_method"] = "request_method"
	queryParams["fields"] = "fields"
	queryParams["user"] = "user"
	queryParams["request_uri"] = "request_uri"
	queryParams["status"] = "status"
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
		"/api/v1/node/tasks",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
