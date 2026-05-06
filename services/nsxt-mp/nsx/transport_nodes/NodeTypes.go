// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Node.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package transport_nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func nodeGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NodeGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodePropertiesBindingType)
}

func nodeGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"/api/v1/transport-nodes/{transportNodeId}/node",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeRestartInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NodeRestartOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func nodeRestartRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=restart",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}/node",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeSetcoredumpconfigInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["core_dump_config"] = vapiBindings_.NewReferenceType(nsxModel.CoreDumpConfigBindingType)
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["core_dump_config"] = "CoreDumpConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NodeSetcoredumpconfigOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func nodeSetcoredumpconfigRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["core_dump_config"] = vapiBindings_.NewReferenceType(nsxModel.CoreDumpConfigBindingType)
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["core_dump_config"] = "CoreDumpConfig"
	paramsTypeMap["core_dump_config"] = vapiBindings_.NewReferenceType(nsxModel.CoreDumpConfigBindingType)
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=set_coredump_config",
		"core_dump_config",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}/node",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeSetsystemtimeInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_time"] = vapiBindings_.NewReferenceType(nsxModel.NodeTimeBindingType)
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_time"] = "NodeTime"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NodeSetsystemtimeOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func nodeSetsystemtimeRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_time"] = vapiBindings_.NewReferenceType(nsxModel.NodeTimeBindingType)
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_time"] = "NodeTime"
	paramsTypeMap["node_time"] = vapiBindings_.NewReferenceType(nsxModel.NodeTimeBindingType)
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=set_system_time",
		"node_time",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}/node",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeShutdownInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NodeShutdownOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func nodeShutdownRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=shutdown",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}/node",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodePropertiesBindingType)
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_properties"] = "NodeProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NodeUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodePropertiesBindingType)
}

func nodeUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["node_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodePropertiesBindingType)
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["node_properties"] = "NodeProperties"
	paramsTypeMap["node_properties"] = vapiBindings_.NewReferenceType(nsxModel.NodePropertiesBindingType)
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"node_properties",
		"PUT",
		"/api/v1/transport-nodes/{transportNodeId}/node",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
