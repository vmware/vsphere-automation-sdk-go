// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Diagnosis.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package logical_routers

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_SERVICE_ROUTER_TIER0 = "SERVICE_ROUTER_TIER0"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_VRF_SERVICE_ROUTER_TIER0 = "VRF_SERVICE_ROUTER_TIER0"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_DISTRIBUTED_ROUTER_TIER0 = "DISTRIBUTED_ROUTER_TIER0"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_VRF_DISTRIBUTED_ROUTER_TIER0 = "VRF_DISTRIBUTED_ROUTER_TIER0"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_SERVICE_ROUTER_TIER1 = "SERVICE_ROUTER_TIER1"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_VRF_SERVICE_ROUTER_TIER1 = "VRF_SERVICE_ROUTER_TIER1"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_DISTRIBUTED_ROUTER_TIER1 = "DISTRIBUTED_ROUTER_TIER1"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_VRF_DISTRIBUTED_ROUTER_TIER1 = "VRF_DISTRIBUTED_ROUTER_TIER1"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_TUNNEL = "TUNNEL"

// Possible value for ``routerType`` of method Diagnosis#get.
const Diagnosis_GET_ROUTER_TYPE_RTEP_TUNNEL_VRF = "RTEP_TUNNEL_VRF"

func diagnosisGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["router_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["router_type"] = "RouterType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DiagnosisGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LogicalRoutersInfoBindingType)
}

func diagnosisGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fields["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["router_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["router_type"] = "RouterType"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["router_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["fields"] = "fields"
	queryParams["router_type"] = "router_type"
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
		"/api/v1/transport-nodes/{transportNodeId}/node/logical-routers/diagnosis",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func diagnosisGet0InputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DiagnosisGet0OutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LogicalRoutersInfoBindingType)
}

func diagnosisGet0RestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = vapiBindings_.NewStringType()
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["logical_router_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["logicalRouterId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/transport-nodes/{transportNodeId}/node/logical-routers/{logicalRouterId}/diagnosis",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
