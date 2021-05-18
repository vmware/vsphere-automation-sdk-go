// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemadministrationconfigurationfabricedgeclustersinterSites.
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

func systemadministrationconfigurationfabricedgeclustersinterSitesGetEdgeClusterInterSiteStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetEdgeClusterInterSiteStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterInterSiteStatusBindingType)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetEdgeClusterInterSiteStatusRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/edge-clusters/{edgeClusterId}/inter-site/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborAdvertisedRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = bindings.NewStringType()
	fields["neighbor_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborAdvertisedRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborRouteDetailsBindingType)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborAdvertisedRoutesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = bindings.NewStringType()
	fields["neighbor_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	paramsTypeMap["neighbor_id"] = bindings.NewStringType()
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	paramsTypeMap["neighborId"] = bindings.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["edge_node_id"] = "edgeNodeId"
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/bgp/neighbors/{neighborId}/advertised-routes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = bindings.NewStringType()
	fields["neighbor_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborRouteDetailsBindingType)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborRoutesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = bindings.NewStringType()
	fields["neighbor_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	paramsTypeMap["neighbor_id"] = bindings.NewStringType()
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	paramsTypeMap["neighborId"] = bindings.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["edge_node_id"] = "edgeNodeId"
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/bgp/neighbors/{neighborId}/routes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpSummaryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpSummaryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InterSiteBgpSummaryBindingType)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpSummaryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	pathParams["edge_node_id"] = "edgeNodeId"
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/bgp/summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeInterSiteStatisticsBindingType)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	pathParams["edge_node_id"] = "edgeNodeId"
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemadministrationconfigurationfabricedgeclustersinterSitesListInterSiteEdgeNodeBgpNeighborsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesListInterSiteEdgeNodeBgpNeighborsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborListResultBindingType)
}

func systemadministrationconfigurationfabricedgeclustersinterSitesListInterSiteEdgeNodeBgpNeighborsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	pathParams["edge_node_id"] = "edgeNodeId"
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/bgp/neighbors",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
