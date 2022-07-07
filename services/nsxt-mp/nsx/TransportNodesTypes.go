// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TransportNodes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``action`` of method TransportNodes#updatemaintenancemode.
const TransportNodes_UPDATEMAINTENANCEMODE_ACTION_ENTER_MAINTENANCE_MODE = "enter_maintenance_mode"

// Possible value for ``action`` of method TransportNodes#updatemaintenancemode.
const TransportNodes_UPDATEMAINTENANCEMODE_ACTION_FORCED_ENTER_MAINTENANCE_MODE = "forced_enter_maintenance_mode"

// Possible value for ``action`` of method TransportNodes#updatemaintenancemode.
const TransportNodes_UPDATEMAINTENANCEMODE_ACTION_EXIT_MAINTENANCE_MODE = "exit_maintenance_mode"

func transportNodesCleanstaleentriesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesCleanstaleentriesOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesCleanstaleentriesRestMetadata() protocol.OperationRestMetadata {
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
		"action=clean_stale_entries",
		"",
		"POST",
		"/api/v1/transport-nodes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fieldNameMap["transport_node"] = "TransportNode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeBindingType)
}

func transportNodesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fieldNameMap["transport_node"] = "TransportNode"
	paramsTypeMap["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
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
		"transport_node",
		"POST",
		"/api/v1/transport-nodes",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["unprepare_host"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["force"] = "Force"
	fieldNameMap["unprepare_host"] = "UnprepareHost"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["unprepare_host"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["force"] = "Force"
	fieldNameMap["unprepare_host"] = "UnprepareHost"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["unprepare_host"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["unprepare_host"] = "unprepare_host"
	queryParams["force"] = "force"
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
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesDeleteontransportnodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesDeleteontransportnodeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesDeleteontransportnodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"/api/v1/transport-nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesDisableflowcacheInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesDisableflowcacheOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesDisableflowcacheRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=disable_flow_cache",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesEnableflowcacheInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesEnableflowcacheOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesEnableflowcacheRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=enable_flow_cache",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeBindingType)
}

func transportNodesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesGetontransportnodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesGetontransportnodeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesGetontransportnodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"/api/v1/transport-nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["in_maintenance_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_ip"] = "NodeIp"
	fieldNameMap["node_types"] = "NodeTypes"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeListResultBindingType)
}

func transportNodesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["in_maintenance_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_ip"] = "NodeIp"
	fieldNameMap["node_types"] = "NodeTypes"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["in_maintenance_mode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["node_types"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["node_types"] = "node_types"
	queryParams["node_ip"] = "node_ip"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["transport_zone_id"] = "transport_zone_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["in_maintenance_mode"] = "in_maintenance_mode"
	queryParams["node_id"] = "node_id"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/transport-nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesMigratetovdsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fields["skip_maintmode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["skip_maintmode"] = "SkipMaintmode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesMigratetovdsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesMigratetovdsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fields["skip_maintmode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["skip_maintmode"] = "SkipMaintmode"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["skip_maintmode"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["skip_maintmode"] = "skip_maintmode"
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
		"action=migrate_to_vds",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesPostontransportnodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesPostontransportnodeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesPostontransportnodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"POST",
		"/api/v1/transport-nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesPutontransportnodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesPutontransportnodeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesPutontransportnodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"PUT",
		"/api/v1/transport-nodes/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesRedeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["transport_node"] = "TransportNode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesRedeployOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeBindingType)
}

func transportNodesRedeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["transport_node"] = "TransportNode"
	paramsTypeMap["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
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
		"action=redeploy",
		"transport_node",
		"POST",
		"/api/v1/transport-nodes/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesRefreshnodeconfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fields["read_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["read_only"] = "ReadOnly"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesRefreshnodeconfigurationOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesRefreshnodeconfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fields["read_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["read_only"] = "ReadOnly"
	paramsTypeMap["read_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["read_only"] = "read_only"
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
		"action=refresh_node_configuration&resource_type=EdgeNode",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesRestartinventorysyncInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesRestartinventorysyncOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesRestartinventorysyncRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=restart_inventory_sync",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesRestoreclusterconfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesRestoreclusterconfigOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesRestoreclusterconfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=restore_cluster_config",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesResynchostconfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transportnode_id"] = bindings.NewStringType()
	fieldNameMap["transportnode_id"] = "TransportnodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesResynchostconfigOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesResynchostconfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transportnode_id"] = bindings.NewStringType()
	fieldNameMap["transportnode_id"] = "TransportnodeId"
	paramsTypeMap["transportnode_id"] = bindings.NewStringType()
	paramsTypeMap["transportnodeId"] = bindings.NewStringType()
	pathParams["transportnode_id"] = "transportnodeId"
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
		"action=resync_host_config",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportnodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fields["esx_mgmt_if_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["if_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ping_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["skip_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vnic_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["transport_node"] = "TransportNode"
	fieldNameMap["esx_mgmt_if_migration_dest"] = "EsxMgmtIfMigrationDest"
	fieldNameMap["if_id"] = "IfId"
	fieldNameMap["ping_ip"] = "PingIp"
	fieldNameMap["skip_validation"] = "SkipValidation"
	fieldNameMap["vnic"] = "Vnic"
	fieldNameMap["vnic_migration_dest"] = "VnicMigrationDest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeBindingType)
}

func transportNodesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fields["esx_mgmt_if_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["if_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ping_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["skip_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vnic_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["transport_node"] = "TransportNode"
	fieldNameMap["esx_mgmt_if_migration_dest"] = "EsxMgmtIfMigrationDest"
	fieldNameMap["if_id"] = "IfId"
	fieldNameMap["ping_ip"] = "PingIp"
	fieldNameMap["skip_validation"] = "SkipValidation"
	fieldNameMap["vnic"] = "Vnic"
	fieldNameMap["vnic_migration_dest"] = "VnicMigrationDest"
	paramsTypeMap["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["skip_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["if_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["esx_mgmt_if_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["ping_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vnic_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["ping_ip"] = "ping_ip"
	queryParams["vnic"] = "vnic"
	queryParams["skip_validation"] = "skip_validation"
	queryParams["esx_mgmt_if_migration_dest"] = "esx_mgmt_if_migration_dest"
	queryParams["if_id"] = "if_id"
	queryParams["vnic_migration_dest"] = "vnic_migration_dest"
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
		"transport_node",
		"PUT",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesUpdatemaintenancemodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transportnode_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transportnode_id"] = "TransportnodeId"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesUpdatemaintenancemodeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesUpdatemaintenancemodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transportnode_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transportnode_id"] = "TransportnodeId"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transportnode_id"] = bindings.NewStringType()
	paramsTypeMap["transportnodeId"] = bindings.NewStringType()
	pathParams["transportnode_id"] = "transportnodeId"
	queryParams["action"] = "action"
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
		"POST",
		"/api/v1/transport-nodes/{transportnodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
