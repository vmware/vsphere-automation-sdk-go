// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AllocatedServices.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``networkServiceClusterType`` of method AllocatedServices#list.
const AllocatedServices_LIST_NETWORK_SERVICE_CLUSTER_TYPE_EDGE = "EDGE"

// Possible value for ``networkServiceClusterType`` of method AllocatedServices#list.
const AllocatedServices_LIST_NETWORK_SERVICE_CLUSTER_TYPE_VIRTUAL_NETWORK_APPLIANCE = "VIRTUAL_NETWORK_APPLIANCE"

// Possible value for ``networkServiceClusterType`` of method AllocatedServices#list.
const AllocatedServices_LIST_NETWORK_SERVICE_CLUSTER_TYPE_VIRTUAL_NETWORK_SERVICE = "VIRTUAL_NETWORK_SERVICE"

func allocatedServicesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_service_cluster_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_service_node_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_conflicts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_service_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_service_cluster_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transit_gateway_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vpc_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["virtual_network_service_cluster_id"] = "VirtualNetworkServiceClusterId"
	fieldNameMap["virtual_network_service_node_id"] = "VirtualNetworkServiceNodeId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_conflicts"] = "IncludeConflicts"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_service_cluster_path"] = "NetworkServiceClusterPath"
	fieldNameMap["network_service_cluster_type"] = "NetworkServiceClusterType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transit_gateway_path"] = "TransitGatewayPath"
	fieldNameMap["vpc_path"] = "VpcPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func AllocatedServicesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.NetworkServiceAllocationListResultBindingType)
}

func allocatedServicesListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["virtual_network_service_cluster_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_service_node_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_conflicts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_service_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_service_cluster_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transit_gateway_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vpc_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["virtual_network_service_cluster_id"] = "VirtualNetworkServiceClusterId"
	fieldNameMap["virtual_network_service_node_id"] = "VirtualNetworkServiceNodeId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_conflicts"] = "IncludeConflicts"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_service_cluster_path"] = "NetworkServiceClusterPath"
	fieldNameMap["network_service_cluster_type"] = "NetworkServiceClusterType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transit_gateway_path"] = "TransitGatewayPath"
	fieldNameMap["vpc_path"] = "VpcPath"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["virtual_network_service_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtual_network_service_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["network_service_cluster_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["network_service_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_conflicts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transit_gateway_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtualNetworkServiceClusterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtualNetworkServiceNodeId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["virtual_network_service_cluster_id"] = "virtualNetworkServiceClusterId"
	pathParams["virtual_network_service_node_id"] = "virtualNetworkServiceNodeId"
	pathParams["site_id"] = "siteId"
	queryParams["cursor"] = "cursor"
	queryParams["network_service_cluster_path"] = "network_service_cluster_path"
	queryParams["include_conflicts"] = "include_conflicts"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["vpc_path"] = "vpc_path"
	queryParams["transit_gateway_path"] = "transit_gateway_path"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["network_service_cluster_type"] = "network_service_cluster_type"
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
		"/policy/api/v1/global-infra/sites/{siteId}/enforcement-points/{enforcementpointId}/virtual-network-service-clusters/{virtualNetworkServiceClusterId}/state/nodes/{virtualNetworkServiceNodeId}/allocated-services",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
