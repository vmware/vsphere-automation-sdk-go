// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricNodesTransportNodeStatus.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``nodeType`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getAllTransportNodesStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_ALL_TRANSPORT_NODES_STATUS_NODE_TYPE_HOST = "HOST"

// Possible value for ``nodeType`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getAllTransportNodesStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_ALL_TRANSPORT_NODES_STATUS_NODE_TYPE_EDGE = "EDGE"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getPnicStatusesForTransportNode.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_PNIC_STATUSES_FOR_TRANSPORT_NODE_STATUS_UNKNOWN = "UNKNOWN"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getPnicStatusesForTransportNode.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_PNIC_STATUSES_FOR_TRANSPORT_NODE_STATUS_UP = "UP"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getPnicStatusesForTransportNode.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_PNIC_STATUSES_FOR_TRANSPORT_NODE_STATUS_DOWN = "DOWN"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getPnicStatusesForTransportNode.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_PNIC_STATUSES_FOR_TRANSPORT_NODE_STATUS_DEGRADED = "DEGRADED"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReport.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReport.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_SOURCE_CACHED = "cached"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReport.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_STATUS_UP = "UP"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReport.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_STATUS_DOWN = "DOWN"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReport.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_STATUS_DEGRADED = "DEGRADED"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReportForaTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_FOR_ATRANSPORT_ZONE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReportForaTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_FOR_ATRANSPORT_ZONE_SOURCE_CACHED = "cached"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReportForaTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_FOR_ATRANSPORT_ZONE_STATUS_UP = "UP"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReportForaTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_FOR_ATRANSPORT_ZONE_STATUS_DOWN = "DOWN"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeReportForaTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_REPORT_FOR_ATRANSPORT_ZONE_STATUS_DEGRADED = "DEGRADED"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#getTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_GET_TRANSPORT_NODE_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_0 = "0"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_NO_DIAGNOSTIC = "NO_DIAGNOSTIC"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_1 = "1"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_CONTROL_DETECTION_TIME_EXPIRED = "CONTROL_DETECTION_TIME_EXPIRED"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_2 = "2"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_ECHO_FUNCTION_FAILED = "ECHO_FUNCTION_FAILED"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_3 = "3"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_NEIGHBOR_SIGNALED_SESSION_DOWN = "NEIGHBOR_SIGNALED_SESSION_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_4 = "4"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_FORWARDING_PLANE_RESET = "FORWARDING_PLANE_RESET"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_5 = "5"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_PATH_DOWN = "PATH_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_6 = "6"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_CONCATENATED_PATH_DOWN = "CONCATENATED_PATH_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_7 = "7"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_ADMINISTRATIVELY_DOWN = "ADMINISTRATIVELY_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_8 = "8"

// Possible value for ``bfdDiagnosticCode`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_BFD_DIAGNOSTIC_CODE_REVERSE_CONCATENATED_PATH_DOWN = "REVERSE_CONCATENATED_PATH_DOWN"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``tunnelStatus`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_TUNNEL_STATUS_UP = "UP"

// Possible value for ``tunnelStatus`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listRemoteTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_REMOTE_TRANSPORT_NODE_STATUS_TUNNEL_STATUS_DOWN = "DOWN"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_STATUS_UP = "UP"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_STATUS_DOWN = "DOWN"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_STATUS_DEGRADED = "DEGRADED"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatus.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_STATUS_UNKNOWN = "UNKNOWN"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatusForTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_FOR_TRANSPORT_ZONE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatusForTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_FOR_TRANSPORT_ZONE_SOURCE_CACHED = "cached"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatusForTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_FOR_TRANSPORT_ZONE_STATUS_UP = "UP"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatusForTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_FOR_TRANSPORT_ZONE_STATUS_DOWN = "DOWN"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatusForTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_FOR_TRANSPORT_ZONE_STATUS_DEGRADED = "DEGRADED"

// Possible value for ``status`` of method SystemAdministrationConfigurationFabricNodesTransportNodeStatus#listTransportNodeStatusForTransportZone.
const SystemAdministrationConfigurationFabricNodesTransportNodeStatus_LIST_TRANSPORT_NODE_STATUS_FOR_TRANSPORT_ZONE_STATUS_UNKNOWN = "UNKNOWN"

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetAllTransportNodesStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_type"] = "NodeType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetAllTransportNodesStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.HeatMapTransportZoneStatusBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetAllTransportNodesStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_type"] = "NodeType"
	paramsTypeMap["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["node_type"] = "node_type"
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
		"/api/v1/transport-nodes/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetPnicStatusesForTransportNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetPnicStatusesForTransportNodeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PnicBondStatusListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetPnicStatusesForTransportNodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
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
		"/api/v1/transport-nodes/{nodeId}/pnic-bond-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeReportInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeReportOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeReportRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["source"] = "source"
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
		"/api/v1/transport-zones/transport-node-status-report",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeReportForaTransportZoneInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["zone_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["zone_id"] = "ZoneId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeReportForaTransportZoneOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeReportForaTransportZoneRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["zone_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["zone_id"] = "ZoneId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["zone_id"] = bindings.NewStringType()
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["zoneId"] = bindings.NewStringType()
	pathParams["zone_id"] = "zoneId"
	queryParams["source"] = "source"
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
		"/api/v1/transport-zones/{zoneId}/transport-node-status-report",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeStatusBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusGetTransportNodeStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["source"] = "source"
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
		"/api/v1/transport-nodes/{nodeId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListRemoteTransportNodeStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["bfd_diagnostic_code"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["bfd_diagnostic_code"] = "BfdDiagnosticCode"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListRemoteTransportNodeStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeStatusListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListRemoteTransportNodeStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["bfd_diagnostic_code"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["bfd_diagnostic_code"] = "BfdDiagnosticCode"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["tunnel_status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["bfd_diagnostic_code"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["tunnel_status"] = "tunnel_status"
	queryParams["bfd_diagnostic_code"] = "bfd_diagnostic_code"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
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
		"/api/v1/transport-nodes/{nodeId}/remote-transport-node-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListTransportNodeStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListTransportNodeStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeStatusListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListTransportNodeStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/transport-zones/transport-node-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListTransportNodeStatusForTransportZoneInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["zone_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["zone_id"] = "ZoneId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListTransportNodeStatusForTransportZoneOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeStatusListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesTransportNodeStatusListTransportNodeStatusForTransportZoneRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["zone_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["zone_id"] = "ZoneId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["zone_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["zoneId"] = bindings.NewStringType()
	pathParams["zone_id"] = "zoneId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/transport-zones/{zoneId}/transport-node-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
