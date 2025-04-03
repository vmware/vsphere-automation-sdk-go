// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ServiceNodes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service_clusters

import (
	"reflect"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func serviceNodesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["relocate_and_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	fieldNameMap["force"] = "Force"
	fieldNameMap["relocate_and_delete"] = "RelocateAndDelete"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceNodesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func serviceNodesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["relocate_and_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	fieldNameMap["force"] = "Force"
	fieldNameMap["relocate_and_delete"] = "RelocateAndDelete"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["relocate_and_delete"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["service_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceClusterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceNodeId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["service_node_id"] = "serviceNodeId"
	pathParams["service_cluster_id"] = "serviceClusterId"
	pathParams["site_id"] = "siteId"
	queryParams["relocate_and_delete"] = "relocate_and_delete"
	queryParams["force"] = "force"
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
		"DELETE",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/service-clusters/{serviceClusterId}/service-nodes/{serviceNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func serviceNodesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceNodesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
}

func serviceNodesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceClusterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceNodeId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["service_node_id"] = "serviceNodeId"
	pathParams["service_cluster_id"] = "serviceClusterId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/service-clusters/{serviceClusterId}/service-nodes/{serviceNodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func serviceNodesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["in_maintenance_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_zone_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["management_ip"] = "ManagementIp"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_zone_path"] = "TransportZonePath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceNodesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeListResultBindingType)
}

func serviceNodesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["in_maintenance_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_zone_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["in_maintenance_mode"] = "InMaintenanceMode"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["management_ip"] = "ManagementIp"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_zone_path"] = "TransportZonePath"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transport_zone_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["service_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["management_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["in_maintenance_mode"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceClusterId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["service_cluster_id"] = "serviceClusterId"
	pathParams["site_id"] = "siteId"
	queryParams["cursor"] = "cursor"
	queryParams["transport_zone_path"] = "transport_zone_path"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["management_ip"] = "management_ip"
	queryParams["sort_by"] = "sort_by"
	queryParams["in_maintenance_mode"] = "in_maintenance_mode"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/service-clusters/{serviceClusterId}/service-nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func serviceNodesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fields["policy_service_node"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	fieldNameMap["policy_service_node"] = "PolicyServiceNode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceNodesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func serviceNodesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fields["policy_service_node"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	fieldNameMap["policy_service_node"] = "PolicyServiceNode"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_service_node"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
	paramsTypeMap["service_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceClusterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceNodeId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["service_node_id"] = "serviceNodeId"
	pathParams["service_cluster_id"] = "serviceClusterId"
	pathParams["site_id"] = "siteId"
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
		"policy_service_node",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/service-clusters/{serviceClusterId}/service-nodes/{serviceNodeId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func serviceNodesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fields["policy_service_node"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	fieldNameMap["policy_service_node"] = "PolicyServiceNode"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceNodesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
}

func serviceNodesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["service_cluster_id"] = vapiBindings_.NewStringType()
	fields["service_node_id"] = vapiBindings_.NewStringType()
	fields["policy_service_node"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["service_cluster_id"] = "ServiceClusterId"
	fieldNameMap["service_node_id"] = "ServiceNodeId"
	fieldNameMap["policy_service_node"] = "PolicyServiceNode"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_service_node"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceNodeBindingType)
	paramsTypeMap["service_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceClusterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceNodeId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["service_node_id"] = "serviceNodeId"
	pathParams["service_cluster_id"] = "serviceClusterId"
	pathParams["site_id"] = "siteId"
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
		"policy_service_node",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/service-clusters/{serviceClusterId}/service-nodes/{serviceNodeId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


