// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CommunicationEntries.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package communication_maps

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
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

func communicationEntriesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommunicationEntriesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func communicationEntriesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationMapId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationEntryId"] = vapiBindings_.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommunicationEntriesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
}

func communicationEntriesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationMapId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationEntryId"] = vapiBindings_.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommunicationEntriesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryListResultBindingType)
}

func communicationEntriesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["communication_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationMapId"] = vapiBindings_.NewStringType()
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
		"/policy/api/v1/infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fields["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommunicationEntriesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func communicationEntriesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fields["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	paramsTypeMap["communication_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationMapId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationEntryId"] = vapiBindings_.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"communication_entry",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesReviseInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fields["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	fields["anchor_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommunicationEntriesReviseOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
}

func communicationEntriesReviseRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fields["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	fields["anchor_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	paramsTypeMap["communication_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["anchor_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationMapId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationEntryId"] = vapiBindings_.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
	queryParams["anchor_path"] = "anchor_path"
	queryParams["operation"] = "operation"
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
		"action=revise",
		"communication_entry",
		"POST",
		"/policy/api/v1/infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func communicationEntriesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fields["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommunicationEntriesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
}

func communicationEntriesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["communication_map_id"] = vapiBindings_.NewStringType()
	fields["communication_entry_id"] = vapiBindings_.NewStringType()
	fields["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["communication_map_id"] = "CommunicationMapId"
	fieldNameMap["communication_entry_id"] = "CommunicationEntryId"
	fieldNameMap["communication_entry"] = "CommunicationEntry"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry"] = vapiBindings_.NewReferenceType(nsx_policyModel.CommunicationEntryBindingType)
	paramsTypeMap["communication_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["communication_entry_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationMapId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communicationEntryId"] = vapiBindings_.NewStringType()
	pathParams["communication_entry_id"] = "communicationEntryId"
	pathParams["communication_map_id"] = "communicationMapId"
	pathParams["domain_id"] = "domainId"
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
		"communication_entry",
		"PUT",
		"/policy/api/v1/infra/domains/{domainId}/communication-maps/{communicationMapId}/communication-entries/{communicationEntryId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
