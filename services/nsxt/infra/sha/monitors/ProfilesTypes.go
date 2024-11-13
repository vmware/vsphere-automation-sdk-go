// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Profiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package monitors

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func profilesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ProfilesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func profilesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	paramsTypeMap["monitor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitor_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitorId"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitorProfileId"] = vapiBindings_.NewStringType()
	pathParams["monitor_profile_id"] = "monitorProfileId"
	pathParams["monitor_id"] = "monitorId"
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
		"/policy/api/v1/infra/sha/monitors/{monitorId}/profiles/{monitorProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func profilesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ProfilesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
}

func profilesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	paramsTypeMap["monitor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitor_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitorId"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitorProfileId"] = vapiBindings_.NewStringType()
	pathParams["monitor_profile_id"] = "monitorProfileId"
	pathParams["monitor_id"] = "monitorId"
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
		"/policy/api/v1/infra/sha/monitors/{monitorId}/profiles/{monitorProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func profilesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ProfilesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileListResultBindingType)
}

func profilesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["monitor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["monitorId"] = vapiBindings_.NewStringType()
	pathParams["monitor_id"] = "monitorId"
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
		"/policy/api/v1/infra/sha/monitors/{monitorId}/profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func profilesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fields["sha_monitor_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	fieldNameMap["sha_monitor_profile"] = "ShaMonitorProfile"
	fieldNameMap["force"] = "Force"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ProfilesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
}

func profilesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fields["sha_monitor_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	fieldNameMap["sha_monitor_profile"] = "ShaMonitorProfile"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["monitor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitor_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sha_monitor_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
	paramsTypeMap["monitorId"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitorProfileId"] = vapiBindings_.NewStringType()
	pathParams["monitor_profile_id"] = "monitorProfileId"
	pathParams["monitor_id"] = "monitorId"
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
		"sha_monitor_profile",
		"PATCH",
		"/policy/api/v1/infra/sha/monitors/{monitorId}/profiles/{monitorProfileId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func profilesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fields["sha_monitor_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	fieldNameMap["sha_monitor_profile"] = "ShaMonitorProfile"
	fieldNameMap["force"] = "Force"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ProfilesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
}

func profilesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["monitor_id"] = vapiBindings_.NewStringType()
	fields["monitor_profile_id"] = vapiBindings_.NewStringType()
	fields["sha_monitor_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["monitor_id"] = "MonitorId"
	fieldNameMap["monitor_profile_id"] = "MonitorProfileId"
	fieldNameMap["sha_monitor_profile"] = "ShaMonitorProfile"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["monitor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitor_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sha_monitor_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.ShaMonitorProfileBindingType)
	paramsTypeMap["monitorId"] = vapiBindings_.NewStringType()
	paramsTypeMap["monitorProfileId"] = vapiBindings_.NewStringType()
	pathParams["monitor_profile_id"] = "monitorProfileId"
	pathParams["monitor_id"] = "monitorId"
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
		"sha_monitor_profile",
		"PUT",
		"/policy/api/v1/infra/sha/monitors/{monitorId}/profiles/{monitorProfileId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
