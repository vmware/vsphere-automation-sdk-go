// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Entries.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package l7_access_profiles

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func entriesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func entriesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func entriesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["l7_access_entry_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["l7_access_profile_id"] = bindings.NewStringType()
	paramsTypeMap["l7AccessProfileId"] = bindings.NewStringType()
	paramsTypeMap["l7AccessEntryId"] = bindings.NewStringType()
	pathParams["l7_access_entry_id"] = "l7AccessEntryId"
	pathParams["l7_access_profile_id"] = "l7AccessProfileId"
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
		"/policy/api/v1/infra/l7-access-profiles/{l7AccessProfileId}/entries/{l7AccessEntryId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func entriesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func entriesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L7AccessEntryBindingType)
}

func entriesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	paramsTypeMap["l7_access_entry_id"] = bindings.NewStringType()
	paramsTypeMap["l7_access_profile_id"] = bindings.NewStringType()
	paramsTypeMap["l7AccessProfileId"] = bindings.NewStringType()
	paramsTypeMap["l7AccessEntryId"] = bindings.NewStringType()
	pathParams["l7_access_entry_id"] = "l7AccessEntryId"
	pathParams["l7_access_profile_id"] = "l7AccessProfileId"
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
		"/policy/api/v1/infra/l7-access-profiles/{l7AccessProfileId}/entries/{l7AccessEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func entriesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func entriesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L7AccessEntryListResultBindingType)
}

func entriesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["l7_access_profile_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["l7AccessProfileId"] = bindings.NewStringType()
	pathParams["l7_access_profile_id"] = "l7AccessProfileId"
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
		"/policy/api/v1/infra/l7-access-profiles/{l7AccessProfileId}/entries",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func entriesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fields["l7_access_entry"] = bindings.NewReferenceType(model.L7AccessEntryBindingType)
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	fieldNameMap["l7_access_entry"] = "L7AccessEntry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func entriesPatchOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L7AccessEntryBindingType)
}

func entriesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fields["l7_access_entry"] = bindings.NewReferenceType(model.L7AccessEntryBindingType)
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	fieldNameMap["l7_access_entry"] = "L7AccessEntry"
	paramsTypeMap["l7_access_entry"] = bindings.NewReferenceType(model.L7AccessEntryBindingType)
	paramsTypeMap["l7_access_entry_id"] = bindings.NewStringType()
	paramsTypeMap["l7_access_profile_id"] = bindings.NewStringType()
	paramsTypeMap["l7AccessProfileId"] = bindings.NewStringType()
	paramsTypeMap["l7AccessEntryId"] = bindings.NewStringType()
	pathParams["l7_access_entry_id"] = "l7AccessEntryId"
	pathParams["l7_access_profile_id"] = "l7AccessProfileId"
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
		"l7_access_entry",
		"PATCH",
		"/policy/api/v1/infra/l7-access-profiles/{l7AccessProfileId}/entries/{l7AccessEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func entriesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fields["l7_access_entry"] = bindings.NewReferenceType(model.L7AccessEntryBindingType)
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	fieldNameMap["l7_access_entry"] = "L7AccessEntry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func entriesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L7AccessEntryBindingType)
}

func entriesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["l7_access_profile_id"] = bindings.NewStringType()
	fields["l7_access_entry_id"] = bindings.NewStringType()
	fields["l7_access_entry"] = bindings.NewReferenceType(model.L7AccessEntryBindingType)
	fieldNameMap["l7_access_profile_id"] = "L7AccessProfileId"
	fieldNameMap["l7_access_entry_id"] = "L7AccessEntryId"
	fieldNameMap["l7_access_entry"] = "L7AccessEntry"
	paramsTypeMap["l7_access_entry"] = bindings.NewReferenceType(model.L7AccessEntryBindingType)
	paramsTypeMap["l7_access_entry_id"] = bindings.NewStringType()
	paramsTypeMap["l7_access_profile_id"] = bindings.NewStringType()
	paramsTypeMap["l7AccessProfileId"] = bindings.NewStringType()
	paramsTypeMap["l7AccessEntryId"] = bindings.NewStringType()
	pathParams["l7_access_entry_id"] = "l7AccessEntryId"
	pathParams["l7_access_profile_id"] = "l7AccessProfileId"
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
		"l7_access_entry",
		"PUT",
		"/policy/api/v1/infra/l7-access-profiles/{l7AccessProfileId}/entries/{l7AccessEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
