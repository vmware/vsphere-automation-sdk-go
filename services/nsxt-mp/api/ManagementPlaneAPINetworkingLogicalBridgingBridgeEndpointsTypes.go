// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints.
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

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints#getBridgeEndpointStatistics.
const ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints_GET_BRIDGE_ENDPOINT_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints#getBridgeEndpointStatistics.
const ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints_GET_BRIDGE_ENDPOINT_STATISTICS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints#getBridgeEndpointStatus.
const ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints_GET_BRIDGE_ENDPOINT_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints#getBridgeEndpointStatus.
const ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints_GET_BRIDGE_ENDPOINT_STATUS_SOURCE_CACHED = "cached"

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsCreateBridgeEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bridge_endpoint"] = bindings.NewReferenceType(model.BridgeEndpointBindingType)
	fieldNameMap["bridge_endpoint"] = "BridgeEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsCreateBridgeEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BridgeEndpointBindingType)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsCreateBridgeEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bridge_endpoint"] = bindings.NewReferenceType(model.BridgeEndpointBindingType)
	fieldNameMap["bridge_endpoint"] = "BridgeEndpoint"
	paramsTypeMap["bridge_endpoint"] = bindings.NewReferenceType(model.BridgeEndpointBindingType)
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
		"bridge_endpoint",
		"POST",
		"/api/v1/bridge-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsDeleteBridgeEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bridgeendpoint_id"] = bindings.NewStringType()
	fieldNameMap["bridgeendpoint_id"] = "BridgeendpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsDeleteBridgeEndpointOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsDeleteBridgeEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bridgeendpoint_id"] = bindings.NewStringType()
	fieldNameMap["bridgeendpoint_id"] = "BridgeendpointId"
	paramsTypeMap["bridgeendpoint_id"] = bindings.NewStringType()
	paramsTypeMap["bridgeendpointId"] = bindings.NewStringType()
	pathParams["bridgeendpoint_id"] = "bridgeendpointId"
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
		"/api/v1/bridge-endpoints/{bridgeendpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bridgeendpoint_id"] = bindings.NewStringType()
	fieldNameMap["bridgeendpoint_id"] = "BridgeendpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BridgeEndpointBindingType)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bridgeendpoint_id"] = bindings.NewStringType()
	fieldNameMap["bridgeendpoint_id"] = "BridgeendpointId"
	paramsTypeMap["bridgeendpoint_id"] = bindings.NewStringType()
	paramsTypeMap["bridgeendpointId"] = bindings.NewStringType()
	pathParams["bridgeendpoint_id"] = "bridgeendpointId"
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
		"/api/v1/bridge-endpoints/{bridgeendpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["endpoint_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["endpoint_id"] = "EndpointId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BridgeEndpointStatisticsBindingType)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["endpoint_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["endpoint_id"] = "EndpointId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["endpointId"] = bindings.NewStringType()
	pathParams["endpoint_id"] = "endpointId"
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
		"/api/v1/bridge-endpoints/{endpointId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["endpoint_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["endpoint_id"] = "EndpointId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BridgeEndpointStatusBindingType)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["endpoint_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["endpoint_id"] = "EndpointId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["endpointId"] = bindings.NewStringType()
	pathParams["endpoint_id"] = "endpointId"
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
		"/api/v1/bridge-endpoints/{endpointId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsListBridgeEndpointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["bridge_endpoint_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan_transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bridge_cluster_id"] = "BridgeClusterId"
	fieldNameMap["bridge_endpoint_profile_id"] = "BridgeEndpointProfileId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["vlan_transport_zone_id"] = "VlanTransportZoneId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsListBridgeEndpointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BridgeEndpointListResultBindingType)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsListBridgeEndpointsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["bridge_endpoint_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan_transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bridge_cluster_id"] = "BridgeClusterId"
	fieldNameMap["bridge_endpoint_profile_id"] = "BridgeEndpointProfileId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["vlan_transport_zone_id"] = "VlanTransportZoneId"
	paramsTypeMap["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["bridge_endpoint_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vlan_transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["bridge_endpoint_profile_id"] = "bridge_endpoint_profile_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["vlan_transport_zone_id"] = "vlan_transport_zone_id"
	queryParams["bridge_cluster_id"] = "bridge_cluster_id"
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
		"/api/v1/bridge-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsUpdateBridgeEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bridgeendpoint_id"] = bindings.NewStringType()
	fields["bridge_endpoint"] = bindings.NewReferenceType(model.BridgeEndpointBindingType)
	fieldNameMap["bridgeendpoint_id"] = "BridgeendpointId"
	fieldNameMap["bridge_endpoint"] = "BridgeEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsUpdateBridgeEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BridgeEndpointBindingType)
}

func managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsUpdateBridgeEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bridgeendpoint_id"] = bindings.NewStringType()
	fields["bridge_endpoint"] = bindings.NewReferenceType(model.BridgeEndpointBindingType)
	fieldNameMap["bridgeendpoint_id"] = "BridgeendpointId"
	fieldNameMap["bridge_endpoint"] = "BridgeEndpoint"
	paramsTypeMap["bridgeendpoint_id"] = bindings.NewStringType()
	paramsTypeMap["bridge_endpoint"] = bindings.NewReferenceType(model.BridgeEndpointBindingType)
	paramsTypeMap["bridgeendpointId"] = bindings.NewStringType()
	pathParams["bridgeendpoint_id"] = "bridgeendpointId"
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
		"bridge_endpoint",
		"PUT",
		"/api/v1/bridge-endpoints/{bridgeendpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
