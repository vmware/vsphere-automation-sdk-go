// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CommunicationEntries.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package communication_maps

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``operation`` of method CommunicationEntries#revise.
const CommunicationEntries_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method CommunicationEntries#revise.
const CommunicationEntries_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method CommunicationEntries#revise.
const CommunicationEntries_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method CommunicationEntries#revise.
const CommunicationEntries_REVISE_OPERATION_BEFORE = "insert_before"

func communicationEntriesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communicationEntriesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func communicationEntriesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["communication_map_id"] = bindings.NewStringType()
	paramsTypeMap["communication_entry_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["communicationMapId"] = bindings.NewStringType()
	paramsTypeMap["communicationEntryId"] = bindings.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/global-infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communicationEntriesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CommunicationEntryBindingType)
}

func communicationEntriesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["communication_map_id"] = bindings.NewStringType()
	paramsTypeMap["communication_entry_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["communicationMapId"] = bindings.NewStringType()
	paramsTypeMap["communicationEntryId"] = bindings.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/global-infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communicationEntriesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CommunicationEntryListResultBindingType)
}

func communicationEntriesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["communication_map_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["communicationMapId"] = bindings.NewStringType()
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/global-infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fields["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communicationEntriesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func communicationEntriesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fields["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["communication_map_id"] = bindings.NewStringType()
	paramsTypeMap["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	paramsTypeMap["communication_entry_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["communicationMapId"] = bindings.NewStringType()
	paramsTypeMap["communicationEntryId"] = bindings.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"communication_entry",
		"PATCH",
		"/policy/api/v1/global-infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fields["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	fields["anchor_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communicationEntriesReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CommunicationEntryBindingType)
}

func communicationEntriesReviseRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fields["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	fields["anchor_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["communication_map_id"] = bindings.NewStringType()
	paramsTypeMap["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	paramsTypeMap["anchor_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["communication_entry_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["communicationMapId"] = bindings.NewStringType()
	paramsTypeMap["communicationEntryId"] = bindings.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
	queryParams["anchor_path"] = "anchor_path"
	queryParams["operation"] = "operation"
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
		"action=revise",
		"communication_entry",
		"POST",
		"/policy/api/v1/global-infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fields["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communicationEntriesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CommunicationEntryBindingType)
}

func communicationEntriesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["communication_map_id"] = bindings.NewStringType()
	fields["communication_entry_id"] = bindings.NewStringType()
	fields["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["communication_map_id"] = bindings.NewStringType()
	paramsTypeMap["communication_entry"] = bindings.NewReferenceType(model.CommunicationEntryBindingType)
	paramsTypeMap["communication_entry_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["communicationMapId"] = bindings.NewStringType()
	paramsTypeMap["communicationEntryId"] = bindings.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"communication_entry",
		"PUT",
		"/policy/api/v1/global-infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
