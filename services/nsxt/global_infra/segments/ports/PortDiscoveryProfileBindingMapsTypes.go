// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PortDiscoveryProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package ports

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func portDiscoveryProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortDiscoveryProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portDiscoveryProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	paramsTypeMap["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraSegmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraPortId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portDiscoveryProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_discovery_profile_binding_map_id"] = "portDiscoveryProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-discovery-profile-binding-maps/{portDiscoveryProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portDiscoveryProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortDiscoveryProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
}

func portDiscoveryProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	paramsTypeMap["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraSegmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraPortId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portDiscoveryProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_discovery_profile_binding_map_id"] = "portDiscoveryProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-discovery-profile-binding-maps/{portDiscoveryProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portDiscoveryProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortDiscoveryProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapListResultBindingType)
}

func portDiscoveryProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["infra_segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["infra_port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["infraSegmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraPortId"] = vapiBindings_.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-discovery-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portDiscoveryProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	fieldNameMap["port_discovery_profile_binding_map"] = "PortDiscoveryProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortDiscoveryProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func portDiscoveryProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	fieldNameMap["port_discovery_profile_binding_map"] = "PortDiscoveryProfileBindingMap"
	paramsTypeMap["port_discovery_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
	paramsTypeMap["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraSegmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraPortId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portDiscoveryProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_discovery_profile_binding_map_id"] = "portDiscoveryProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"port_discovery_profile_binding_map",
		"PATCH",
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-discovery-profile-binding-maps/{portDiscoveryProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portDiscoveryProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	fieldNameMap["port_discovery_profile_binding_map"] = "PortDiscoveryProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PortDiscoveryProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
}

func portDiscoveryProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = vapiBindings_.NewStringType()
	fields["infra_port_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["port_discovery_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_discovery_profile_binding_map_id"] = "PortDiscoveryProfileBindingMapId"
	fieldNameMap["port_discovery_profile_binding_map"] = "PortDiscoveryProfileBindingMap"
	paramsTypeMap["port_discovery_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PortDiscoveryProfileBindingMapBindingType)
	paramsTypeMap["port_discovery_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infra_port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraSegmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["infraPortId"] = vapiBindings_.NewStringType()
	paramsTypeMap["portDiscoveryProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_discovery_profile_binding_map_id"] = "portDiscoveryProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"port_discovery_profile_binding_map",
		"PUT",
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-discovery-profile-binding-maps/{portDiscoveryProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
