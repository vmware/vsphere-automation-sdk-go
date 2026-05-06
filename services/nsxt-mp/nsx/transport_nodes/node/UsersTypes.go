// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Users.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package node

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``filter`` of method Users#list.
const Users_LIST_FILTER_ALL = "all"

// Possible value for ``filter`` of method Users#list.
const Users_LIST_FILTER_GUEST = "guest"

// Possible value for ``filter`` of method Users#list.
const Users_LIST_FILTER_SERVICE = "service"

func usersDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func usersDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	pathParams["userid"] = "userid"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"/api/v1/transport-nodes/{transportNodeId}/node/users/{userid}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	pathParams["userid"] = "userid"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"/api/v1/transport-nodes/{transportNodeId}/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["filter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["internal"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["filter"] = "Filter"
	fieldNameMap["internal"] = "Internal"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesListResultBindingType)
}

func usersListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["filter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["internal"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["filter"] = "Filter"
	fieldNameMap["internal"] = "Internal"
	paramsTypeMap["filter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["internal"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["filter"] = "filter"
	queryParams["internal"] = "internal"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"/api/v1/transport-nodes/{transportNodeId}/node/users",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersResetpasswordInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersResetpasswordOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func usersResetpasswordRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	pathParams["userid"] = "userid"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"/api/v1/transport-nodes/{transportNodeId}/node/users/{userid}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	pathParams["userid"] = "userid"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"/api/v1/transport-nodes/{transportNodeId}/node/users/{userid}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
