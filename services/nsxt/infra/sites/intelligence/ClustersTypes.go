// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Clusters.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package intelligence

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``clusterType`` of method Clusters#list.
const Clusters_LIST_CLUSTER_TYPE_CLUSTER = "HOST_CLUSTER"

func clustersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_id"] = "ClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
}

func clustersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_id"] = "ClusterId"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["cluster_id"] = "clusterId"
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
		"/policy/api/v1/infra/sites/{siteId}/intelligence/clusters/{clusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clustersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_type"] = "ClusterType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enabled"] = "Enabled"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoListBindingType)
}

func clustersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_type"] = "ClusterType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enabled"] = "Enabled"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cluster_type"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["cluster_type"] = "cluster_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["enabled"] = "enabled"
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
		"/policy/api/v1/infra/sites/{siteId}/intelligence/clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clustersPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_cluster_info_list"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoListBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["intelligence_transport_node_cluster_info_list"] = "IntelligenceTransportNodeClusterInfoList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersPatchOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoListBindingType)
}

func clustersPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_cluster_info_list"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoListBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["intelligence_transport_node_cluster_info_list"] = "IntelligenceTransportNodeClusterInfoList"
	paramsTypeMap["intelligence_transport_node_cluster_info_list"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoListBindingType)
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
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
		"intelligence_transport_node_cluster_info_list",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/intelligence/clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clustersPatch0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_cluster_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["intelligence_transport_node_cluster_info"] = "IntelligenceTransportNodeClusterInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersPatch0OutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
}

func clustersPatch0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_cluster_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["intelligence_transport_node_cluster_info"] = "IntelligenceTransportNodeClusterInfo"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["intelligence_transport_node_cluster_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["cluster_id"] = "clusterId"
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
		"intelligence_transport_node_cluster_info",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/intelligence/clusters/{clusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clustersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_cluster_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["intelligence_transport_node_cluster_info"] = "IntelligenceTransportNodeClusterInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
}

func clustersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_cluster_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["intelligence_transport_node_cluster_info"] = "IntelligenceTransportNodeClusterInfo"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["intelligence_transport_node_cluster_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeClusterInfoBindingType)
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["cluster_id"] = "clusterId"
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
		"intelligence_transport_node_cluster_info",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/intelligence/clusters/{clusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
