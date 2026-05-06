// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FlowMon.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package dataplane

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func flowMonGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["core_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["timeout"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["core_id"] = "CoreId"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["timeout"] = "Timeout"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FlowMonGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.EdgeDataplaneTopkFlowsBindingType)
}

func flowMonGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["core_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["timeout"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["core_id"] = "CoreId"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["timeout"] = "Timeout"
	paramsTypeMap["core_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["timeout"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["core_id"] = "core_id"
	queryParams["fields"] = "fields"
	queryParams["timeout"] = "timeout"
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
		"/api/v1/transport-nodes/{transportNodeId}/node/services/nsxt-mp/dataplane/flow-mon",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func flowMonUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["edge_dataplane_flow_monitor_start_setting"] = vapiBindings_.NewReferenceType(nsxModel.EdgeDataplaneFlowMonitorStartSettingBindingType)
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["top10"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["edge_dataplane_flow_monitor_start_setting"] = "EdgeDataplaneFlowMonitorStartSetting"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["top10"] = "Top10"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FlowMonUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.EdgeDataplaneFlowMonitorMessageBindingType)
}

func flowMonUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["edge_dataplane_flow_monitor_start_setting"] = vapiBindings_.NewReferenceType(nsxModel.EdgeDataplaneFlowMonitorStartSettingBindingType)
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["top10"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["edge_dataplane_flow_monitor_start_setting"] = "EdgeDataplaneFlowMonitorStartSetting"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["top10"] = "Top10"
	paramsTypeMap["top10"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["edge_dataplane_flow_monitor_start_setting"] = vapiBindings_.NewReferenceType(nsxModel.EdgeDataplaneFlowMonitorStartSettingBindingType)
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["top10"] = "top_10"
	queryParams["fields"] = "fields"
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
		"edge_dataplane_flow_monitor_start_setting",
		"PUT",
		"/api/v1/transport-nodes/{transportNodeId}/node/services/nsxt-mp/dataplane/flow-mon",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
