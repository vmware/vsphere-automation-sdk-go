// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Users.
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

func usersActivateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersActivateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func usersActivateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=activate",
		"node_user_password_property",
		"POST",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersDeactivateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersDeactivateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func usersDeactivateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=deactivate",
		"",
		"POST",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func usersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesListResultBindingType)
}

func usersListRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/node/users",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersResetownpasswordInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reset_node_user_own_password_properties"] = bindings.NewReferenceType(model.ResetNodeUserOwnPasswordPropertiesBindingType)
	fieldNameMap["reset_node_user_own_password_properties"] = "ResetNodeUserOwnPasswordProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersResetownpasswordOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func usersResetownpasswordRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["reset_node_user_own_password_properties"] = bindings.NewReferenceType(model.ResetNodeUserOwnPasswordPropertiesBindingType)
	fieldNameMap["reset_node_user_own_password_properties"] = "ResetNodeUserOwnPasswordProperties"
	paramsTypeMap["reset_node_user_own_password_properties"] = bindings.NewReferenceType(model.ResetNodeUserOwnPasswordPropertiesBindingType)
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
		"action=reset_own_password",
		"reset_node_user_own_password_properties",
		"POST",
		"/api/v1/node/users",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersResetpasswordInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersResetpasswordOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func usersResetpasswordRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=reset_password",
		"node_user_password_property",
		"POST",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["node_user_properties"] = bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func usersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func usersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["node_user_properties"] = bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"node_user_properties",
		"PUT",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
