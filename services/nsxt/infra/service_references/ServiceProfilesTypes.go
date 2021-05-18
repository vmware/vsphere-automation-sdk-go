// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ServiceProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service_references

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func serviceProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
}

func serviceProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles/{serviceProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyServiceProfileListResultBindingType)
}

func serviceProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
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
		"/policy/api/v1/infra/service-references/{serviceReferenceId}/service-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceProfilesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fields["policy_service_profile"] = bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceProfilesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceProfilesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fields["policy_service_profile"] = bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["policy_service_profile"] = bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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

func serviceProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fields["policy_service_profile"] = bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
}

func serviceProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_reference_id"] = bindings.NewStringType()
	fields["service_profile_id"] = bindings.NewStringType()
	fields["policy_service_profile"] = bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
	fieldNameMap["service_reference_id"] = "ServiceReferenceId"
	fieldNameMap["service_profile_id"] = "ServiceProfileId"
	fieldNameMap["policy_service_profile"] = "PolicyServiceProfile"
	paramsTypeMap["service_reference_id"] = bindings.NewStringType()
	paramsTypeMap["policy_service_profile"] = bindings.NewReferenceType(model.PolicyServiceProfileBindingType)
	paramsTypeMap["service_profile_id"] = bindings.NewStringType()
	paramsTypeMap["serviceReferenceId"] = bindings.NewStringType()
	paramsTypeMap["serviceProfileId"] = bindings.NewStringType()
	pathParams["service_reference_id"] = "serviceReferenceId"
	pathParams["service_profile_id"] = "serviceProfileId"
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
