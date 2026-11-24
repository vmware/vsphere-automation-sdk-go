// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
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
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func segmentRealTimeEthProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
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
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
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
		"/global-manager/api/v1/global-infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentRealTimeEthProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
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
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	paramsTypeMap["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
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
		"/global-manager/api/v1/global-infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentRealTimeEthProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SegmentRealTimeEthProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
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
	fields["segment_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["segment_real_time_eth_profile_binding_map_id"] = "SegmentRealTimeEthProfileBindingMapId"
	fieldNameMap["segment_real_time_eth_profile_binding_map"] = "SegmentRealTimeEthProfileBindingMap"
	paramsTypeMap["segment_real_time_eth_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.SegmentRealTimeEthProfileBindingMapBindingType)
	paramsTypeMap["segment_real_time_eth_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segment_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentId"] = vapiBindings_.NewStringType()
	paramsTypeMap["segmentRealTimeEthProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["segment_real_time_eth_profile_binding_map_id"] = "segmentRealTimeEthProfileBindingMapId"
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
		"/global-manager/api/v1/global-infra/segments/{segmentId}/segment-real-time-eth-profile-binding-maps/{segmentRealTimeEthProfileBindingMapId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
