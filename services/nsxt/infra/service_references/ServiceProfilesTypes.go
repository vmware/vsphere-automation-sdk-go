// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ServiceProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service_references

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func serviceProfilesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceProfilesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func serviceProfilesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_reference_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceReferenceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceProfileId"] = vapiBindings_.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceProfilesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
}

func serviceProfilesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_reference_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceReferenceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceProfileId"] = vapiBindings_.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceProfilesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileListResultBindingType)
}

func serviceProfilesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["service_reference_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["serviceReferenceId"] = vapiBindings_.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
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
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_service_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceProfilesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func serviceProfilesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_service_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	paramsTypeMap["service_reference_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_service_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
	paramsTypeMap["service_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceReferenceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceProfileId"] = vapiBindings_.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
		"policy_service_profile",
		"PATCH",
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_service_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ServiceProfilesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
}

func serviceProfilesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = vapiBindings_.NewStringType()
	fields["service_profile_id"] = vapiBindings_.NewStringType()
	fields["policy_service_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	paramsTypeMap["service_reference_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_service_profile"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyServiceProfileBindingType)
	paramsTypeMap["service_profile_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceReferenceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceProfileId"] = vapiBindings_.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
		"policy_service_profile",
		"PUT",
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
