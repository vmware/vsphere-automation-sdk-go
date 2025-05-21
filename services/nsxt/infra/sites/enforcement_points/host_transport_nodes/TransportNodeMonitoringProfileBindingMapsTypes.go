// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TransportNodeMonitoringProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package host_transport_nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func transportNodeMonitoringProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TransportNodeMonitoringProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func transportNodeMonitoringProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	paramsTypeMap["host_transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["hostTransportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingMapId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["host_transport_node_id"] = "hostTransportNodeId"
	pathParams["binding_map_id"] = "bindingMapId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/host-transport-nodes/{hostTransportNodeId}/transport-node-monitoring-profile-binding-maps/{bindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeMonitoringProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TransportNodeMonitoringProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
}

func transportNodeMonitoringProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	paramsTypeMap["host_transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["hostTransportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingMapId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["host_transport_node_id"] = "hostTransportNodeId"
	pathParams["binding_map_id"] = "bindingMapId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/host-transport-nodes/{hostTransportNodeId}/transport-node-monitoring-profile-binding-maps/{bindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeMonitoringProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TransportNodeMonitoringProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapListResultBindingType)
}

func transportNodeMonitoringProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["host_transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["hostTransportNodeId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["host_transport_node_id"] = "hostTransportNodeId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/host-transport-nodes/{hostTransportNodeId}/transport-node-monitoring-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeMonitoringProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fields["transport_node_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	fieldNameMap["transport_node_monitoring_profile_binding_map"] = "TransportNodeMonitoringProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TransportNodeMonitoringProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func transportNodeMonitoringProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fields["transport_node_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	fieldNameMap["transport_node_monitoring_profile_binding_map"] = "TransportNodeMonitoringProfileBindingMap"
	paramsTypeMap["host_transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transport_node_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["hostTransportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingMapId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["host_transport_node_id"] = "hostTransportNodeId"
	pathParams["binding_map_id"] = "bindingMapId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"transport_node_monitoring_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/host-transport-nodes/{hostTransportNodeId}/transport-node-monitoring-profile-binding-maps/{bindingMapId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeMonitoringProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fields["transport_node_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	fieldNameMap["transport_node_monitoring_profile_binding_map"] = "TransportNodeMonitoringProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TransportNodeMonitoringProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
}

func transportNodeMonitoringProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["host_transport_node_id"] = vapiBindings_.NewStringType()
	fields["binding_map_id"] = vapiBindings_.NewStringType()
	fields["transport_node_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["host_transport_node_id"] = "HostTransportNodeId"
	fieldNameMap["binding_map_id"] = "BindingMapId"
	fieldNameMap["transport_node_monitoring_profile_binding_map"] = "TransportNodeMonitoringProfileBindingMap"
	paramsTypeMap["host_transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transport_node_monitoring_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeMonitoringProfileBindingMapBindingType)
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["hostTransportNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingMapId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["host_transport_node_id"] = "hostTransportNodeId"
	pathParams["binding_map_id"] = "bindingMapId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"transport_node_monitoring_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/host-transport-nodes/{hostTransportNodeId}/transport-node-monitoring-profile-binding-maps/{bindingMapId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
