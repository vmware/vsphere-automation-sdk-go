// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SegmentRealTimeEthProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package segments

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func segmentRealTimeEthProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentRealTimeEthProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func segmentRealTimeEthProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentRealTimeEthProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentRealTimeEthProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
}

func segmentRealTimeEthProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentRealTimeEthProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentRealTimeEthProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapListResultBindingType)
}

func segmentRealTimeEthProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentRealTimeEthProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentRealTimeEthProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func segmentRealTimeEthProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	paramsTypeMap["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"segment_real_time_eth_profile_binding_map",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentRealTimeEthProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentRealTimeEthProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
}

func segmentRealTimeEthProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	paramsTypeMap["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"segment_real_time_eth_profile_binding_map",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
