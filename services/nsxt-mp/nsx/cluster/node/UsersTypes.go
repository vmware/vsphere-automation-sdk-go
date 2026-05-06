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

func usersActivateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersActivateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersActivateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"action=activate",
		"node_user_password_property",
		"POST",
		"/api/v1/cluster/{clusterNodeId}/node/users/{userid}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersCreateaudituserInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersCreateaudituserOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersCreateaudituserRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"action=create_audit_user",
		"node_user_properties",
		"POST",
		"/api/v1/cluster/{clusterNodeId}/node/users",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersCreateserviceuserInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersCreateserviceuserOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersCreateserviceuserRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"action=create_service_user",
		"node_user_properties",
		"POST",
		"/api/v1/cluster/{clusterNodeId}/node/users",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersCreateuserInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersCreateuserOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersCreateuserRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"action=create_user",
		"node_user_properties",
		"POST",
		"/api/v1/cluster/{clusterNodeId}/node/users",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersDeactivateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersDeactivateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
}

func usersDeactivateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"action=deactivate",
		"",
		"POST",
		"/api/v1/cluster/{clusterNodeId}/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"/api/v1/cluster/{clusterNodeId}/node/users/{userid}",
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"/api/v1/cluster/{clusterNodeId}/node/users/{userid}",
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["filter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["internal"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["filter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["internal"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["filter"] = "Filter"
	fieldNameMap["internal"] = "Internal"
	paramsTypeMap["filter"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["internal"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"/api/v1/cluster/{clusterNodeId}/node/users",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersResetownpasswordInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["reset_node_user_own_password_properties"] = vapiBindings_.NewReferenceType(nsxModel.ResetNodeUserOwnPasswordPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["reset_node_user_own_password_properties"] = "ResetNodeUserOwnPasswordProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UsersResetownpasswordOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func usersResetownpasswordRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["reset_node_user_own_password_properties"] = vapiBindings_.NewReferenceType(nsxModel.ResetNodeUserOwnPasswordPropertiesBindingType)
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["reset_node_user_own_password_properties"] = "ResetNodeUserOwnPasswordProperties"
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["reset_node_user_own_password_properties"] = vapiBindings_.NewReferenceType(nsxModel.ResetNodeUserOwnPasswordPropertiesBindingType)
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"action=reset_own_password",
		"reset_node_user_own_password_properties",
		"POST",
		"/api/v1/cluster/{clusterNodeId}/node/users",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func usersResetpasswordInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"/api/v1/cluster/{clusterNodeId}/node/users/{userid}",
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
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
	fields["cluster_node_id"] = vapiBindings_.NewStringType()
	fields["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["cluster_node_id"] = "ClusterNodeId"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodeUserPropertiesBindingType)
	paramsTypeMap["cluster_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["clusterNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["cluster_node_id"] = "clusterNodeId"
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
		"/api/v1/cluster/{clusterNodeId}/node/users/{userid}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
