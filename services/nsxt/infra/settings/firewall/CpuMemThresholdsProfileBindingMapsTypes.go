// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CpuMemThresholdsProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func cpuMemThresholdsProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CpuMemThresholdsProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func cpuMemThresholdsProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CpuMemThresholdsProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
}

func cpuMemThresholdsProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CpuMemThresholdsProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapListResultBindingType)
}

func cpuMemThresholdsProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
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
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CpuMemThresholdsProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func cpuMemThresholdsProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	paramsTypeMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"policy_firewall_CPU_mem_thresholds_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CpuMemThresholdsProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
}

func cpuMemThresholdsProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	paramsTypeMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"policy_firewall_CPU_mem_thresholds_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
