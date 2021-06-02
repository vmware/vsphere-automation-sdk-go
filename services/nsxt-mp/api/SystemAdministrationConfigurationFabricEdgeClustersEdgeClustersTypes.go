// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricEdgeClustersEdgeClusters.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricEdgeClustersEdgeClusters#getEdgeClusterStatus.
const SystemAdministrationConfigurationFabricEdgeClustersEdgeClusters_GET_EDGE_CLUSTER_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method SystemAdministrationConfigurationFabricEdgeClustersEdgeClusters#getEdgeClusterStatus.
const SystemAdministrationConfigurationFabricEdgeClustersEdgeClusters_GET_EDGE_CLUSTER_STATUS_SOURCE_CACHED = "cached"

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersCreateEdgeClusterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersCreateEdgeClusterOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersCreateEdgeClusterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	paramsTypeMap["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
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
		"edge_cluster",
		"POST",
		"/api/v1/edge-clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersDeleteEdgeClusterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersDeleteEdgeClusterOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersDeleteEdgeClusterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterAllocationStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterAllocationStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterAllocationStatusBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterAllocationStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}/allocation-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterStateBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterStatusBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersListEdgeClustersInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersListEdgeClustersOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterListResultBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersListEdgeClustersRestMetadata() protocol.OperationRestMetadata {
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
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/edge-clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReadEdgeClusterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReadEdgeClusterOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReadEdgeClusterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReplaceEdgeClusterMemberTransportNodeReplaceTransportNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster_member_transport_node"] = bindings.NewReferenceType(model.EdgeClusterMemberTransportNodeBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster_member_transport_node"] = "EdgeClusterMemberTransportNode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReplaceEdgeClusterMemberTransportNodeReplaceTransportNodeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReplaceEdgeClusterMemberTransportNodeReplaceTransportNodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster_member_transport_node"] = bindings.NewReferenceType(model.EdgeClusterMemberTransportNodeBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster_member_transport_node"] = "EdgeClusterMemberTransportNode"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edge_cluster_member_transport_node"] = bindings.NewReferenceType(model.EdgeClusterMemberTransportNodeBindingType)
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"action=replace_transport_node",
		"edge_cluster_member_transport_node",
		"POST",
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersUpdateEdgeClusterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersUpdateEdgeClusterOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func systemAdministrationConfigurationFabricEdgeClustersEdgeClustersUpdateEdgeClusterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"edge_cluster",
		"PUT",
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
