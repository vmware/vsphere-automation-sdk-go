// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Task.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
	"reflect"
)

func taskGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["task"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["task"] = "Task"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func taskGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TaskBindingType)
}

func taskGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["task"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["task"] = "Task"
	paramsTypeMap["task"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["task"] = bindings.NewStringType()
	pathParams["task"] = "task"
	pathParams["org"] = "org"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/vmc/draas/api/orgs/{org}/tasks/{task}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}

func taskListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func taskListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(model.TaskBindingType), reflect.TypeOf([]model.Task{}))
}

func taskListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["filter"] = "Filter"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	pathParams["org"] = "org"
	queryParams["filter"] = "$filter"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/vmc/draas/api/orgs/{org}/tasks",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func taskUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["task"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["task"] = "Task"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func taskUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TaskBindingType)
}

func taskUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["task"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["task"] = "Task"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["task"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["task"] = bindings.NewStringType()
	pathParams["task"] = "task"
	pathParams["org"] = "org"
	queryParams["action"] = "action"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"POST",
		"/vmc/draas/api/orgs/{org}/tasks/{task}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}
