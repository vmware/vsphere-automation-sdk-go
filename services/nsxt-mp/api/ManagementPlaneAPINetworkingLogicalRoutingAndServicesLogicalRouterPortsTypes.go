// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts.
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

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortArpTable.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_ARP_TABLE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortArpTable.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_ARP_TABLE_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortArpTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_ARP_TABLE_IN_CSV_FORMAT_CSV_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortArpTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_ARP_TABLE_IN_CSV_FORMAT_CSV_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortStatistics.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortStatistics.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_STATISTICS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortStatisticsSummary.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_STATISTICS_SUMMARY_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#getLogicalRouterPortStatisticsSummary.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_GET_LOGICAL_ROUTER_PORT_STATISTICS_SUMMARY_SOURCE_CACHED = "cached"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERUPLINKPORT = "LogicalRouterUpLinkPort"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERDOWNLINKPORT = "LogicalRouterDownLinkPort"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERLINKPORTONTIER0 = "LogicalRouterLinkPortOnTIER0"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERLINKPORTONTIER1 = "LogicalRouterLinkPortOnTIER1"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERLOOPBACKPORT = "LogicalRouterLoopbackPort"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERIPTUNNELPORT = "LogicalRouterIPTunnelPort"

// Possible value for ``resourceType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts#listLogicalRouterPorts.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPorts_LIST_LOGICAL_ROUTER_PORTS_RESOURCE_TYPE_LOGICALROUTERCENTRALIZEDSERVICEPORT = "LogicalRouterCentralizedServicePort"

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsCreateLogicalRouterPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsCreateLogicalRouterPortOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsCreateLogicalRouterPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	paramsTypeMap["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
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
		"logical_router_port",
		"POST",
		"/api/v1/logical-router-ports",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsDeleteLogicalRouterPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["cascade_delete_linked_ports"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["cascade_delete_linked_ports"] = "CascadeDeleteLinkedPorts"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsDeleteLogicalRouterPortOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsDeleteLogicalRouterPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["cascade_delete_linked_ports"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["cascade_delete_linked_ports"] = "CascadeDeleteLinkedPorts"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["cascade_delete_linked_ports"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["cascade_delete_linked_ports"] = "cascade_delete_linked_ports"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortArpTableInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortArpTableOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortArpTableBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortArpTableRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}/arp-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortArpTableInCsvFormatCsvInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortArpTableInCsvFormatCsvOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortArpTableInCsvFormatBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortArpTableInCsvFormatCsvRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"format=csv",
		"",
		"GET",
		"/api/v1/logical-router-ports/{logicalRouterPortId}/arp-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortStateBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	paramsTypeMap["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["barrier_id"] = "barrier_id"
	queryParams["request_id"] = "request_id"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortStatisticsBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStatisticsSummaryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStatisticsSummaryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortStatisticsSummaryBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsGetLogicalRouterPortStatisticsSummaryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}/statistics/summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsListLogicalRouterPortsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsListLogicalRouterPortsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsListLogicalRouterPortsRestMetadata() protocol.OperationRestMetadata {
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
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["logical_router_id"] = "logical_router_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["resource_type"] = "resource_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["logical_switch_id"] = "logical_switch_id"
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
		"/api/v1/logical-router-ports",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsReadLogicalRouterPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsReadLogicalRouterPortOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsReadLogicalRouterPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsUpdateLogicalRouterPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsUpdateLogicalRouterPortOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouterPortsUpdateLogicalRouterPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	paramsTypeMap["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
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
		"logical_router_port",
		"PUT",
		"/api/v1/logical-router-ports/{logicalRouterPortId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
