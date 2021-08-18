// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TlsInspectionConfigProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package security

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func tlsInspectionConfigProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsInspectionConfigProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tlsInspectionConfigProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["tls_inspection_config_profile"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["tlsInspectionConfigProfile"] = bindings.NewStringType()
	pathParams["tls_inspection_config_profile"] = "tlsInspectionConfigProfile"
	queryParams["override"] = "override"
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
		"/policy/api/v1/infra/security/tls-inspection-config-profiles/{tlsInspectionConfigProfile}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsInspectionConfigProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
}

func tlsInspectionConfigProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	paramsTypeMap["tls_inspection_config_profile"] = bindings.NewStringType()
	paramsTypeMap["tlsInspectionConfigProfile"] = bindings.NewStringType()
	pathParams["tls_inspection_config_profile"] = "tlsInspectionConfigProfile"
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
		"/policy/api/v1/infra/security/tls-inspection-config-profiles/{tlsInspectionConfigProfile}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsInspectionConfigProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyTlsConfigProfileListResultBindingType)
}

func tlsInspectionConfigProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
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
		"/policy/api/v1/infra/security/tls-inspection-config-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfilesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fields["policy_tls_config_profile"] = bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	fieldNameMap["policy_tls_config_profile"] = "PolicyTlsConfigProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsInspectionConfigProfilesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tlsInspectionConfigProfilesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fields["policy_tls_config_profile"] = bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	fieldNameMap["policy_tls_config_profile"] = "PolicyTlsConfigProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["tls_inspection_config_profile"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["policy_tls_config_profile"] = bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
	paramsTypeMap["tlsInspectionConfigProfile"] = bindings.NewStringType()
	pathParams["tls_inspection_config_profile"] = "tlsInspectionConfigProfile"
	queryParams["override"] = "override"
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
		"policy_tls_config_profile",
		"PATCH",
		"/policy/api/v1/infra/security/tls-inspection-config-profiles/{tlsInspectionConfigProfile}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tlsInspectionConfigProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fields["policy_tls_config_profile"] = bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	fieldNameMap["policy_tls_config_profile"] = "PolicyTlsConfigProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsInspectionConfigProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
}

func tlsInspectionConfigProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tls_inspection_config_profile"] = bindings.NewStringType()
	fields["policy_tls_config_profile"] = bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["tls_inspection_config_profile"] = "TlsInspectionConfigProfile"
	fieldNameMap["policy_tls_config_profile"] = "PolicyTlsConfigProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["tls_inspection_config_profile"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["policy_tls_config_profile"] = bindings.NewReferenceType(model.PolicyTlsConfigProfileBindingType)
	paramsTypeMap["tlsInspectionConfigProfile"] = bindings.NewStringType()
	pathParams["tls_inspection_config_profile"] = "tlsInspectionConfigProfile"
	queryParams["override"] = "override"
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
		"policy_tls_config_profile",
		"PUT",
		"/policy/api/v1/infra/security/tls-inspection-config-profiles/{tlsInspectionConfigProfile}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
