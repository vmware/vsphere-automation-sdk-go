// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PortMonitoringProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package ports

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func portMonitoringProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portMonitoringProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["infra_port_id"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["infraPortId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
}

func portMonitoringProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["infra_port_id"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["infraPortId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMonitoringProfileBindingMapListResultBindingType)
}

func portMonitoringProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["infra_port_id"] = bindings.NewStringType()
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["infraPortId"] = bindings.NewStringType()
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
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-monitoring-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portMonitoringProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	paramsTypeMap["infra_port_id"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["infraPortId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"port_monitoring_profile_binding_map",
		"PATCH",
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func portMonitoringProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMonitoringProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
}

func portMonitoringProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["infra_port_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["infra_port_id"] = "InfraPortId"
	fieldNameMap["port_monitoring_profile_binding_map_id"] = "PortMonitoringProfileBindingMapId"
	fieldNameMap["port_monitoring_profile_binding_map"] = "PortMonitoringProfileBindingMap"
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.PortMonitoringProfileBindingMapBindingType)
	paramsTypeMap["infra_port_id"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["infraPortId"] = bindings.NewStringType()
	paramsTypeMap["portMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["infra_port_id"] = "infraPortId"
	pathParams["port_monitoring_profile_binding_map_id"] = "portMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"port_monitoring_profile_binding_map",
		"PUT",
		"/policy/api/v1/global-infra/segments/{infraSegmentId}/ports/{infraPortId}/port-monitoring-profile-binding-maps/{portMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
