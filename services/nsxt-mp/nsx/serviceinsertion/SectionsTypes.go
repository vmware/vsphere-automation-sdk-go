// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Sections.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package serviceinsertion

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``operation`` of method Sections#create.
const Sections_CREATE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Sections#create.
const Sections_CREATE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Sections#create.
const Sections_CREATE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Sections#create.
const Sections_CREATE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method Sections#createwithrules.
const Sections_CREATEWITHRULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Sections#createwithrules.
const Sections_CREATEWITHRULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Sections#createwithrules.
const Sections_CREATEWITHRULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Sections#createwithrules.
const Sections_CREATEWITHRULES_OPERATION_BEFORE = "insert_before"

// Possible value for ``excludeAppliedToType`` of method Sections#list.
const Sections_LIST_EXCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``excludeAppliedToType`` of method Sections#list.
const Sections_LIST_EXCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``excludeAppliedToType`` of method Sections#list.
const Sections_LIST_EXCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``excludeAppliedToType`` of method Sections#list.
const Sections_LIST_EXCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``filterType`` of method Sections#list.
const Sections_LIST_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method Sections#list.
const Sections_LIST_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``includeAppliedToType`` of method Sections#list.
const Sections_LIST_INCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``includeAppliedToType`` of method Sections#list.
const Sections_LIST_INCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``includeAppliedToType`` of method Sections#list.
const Sections_LIST_INCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``includeAppliedToType`` of method Sections#list.
const Sections_LIST_INCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``type`` of method Sections#list.
const Sections_LIST_TYPE_L3REDIRECT = "L3REDIRECT"

// Possible value for ``operation`` of method Sections#revise.
const Sections_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Sections#revise.
const Sections_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Sections#revise.
const Sections_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Sections#revise.
const Sections_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method Sections#revisewithrules.
const Sections_REVISEWITHRULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Sections#revisewithrules.
const Sections_REVISEWITHRULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Sections#revisewithrules.
const Sections_REVISEWITHRULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Sections#revisewithrules.
const Sections_REVISEWITHRULES_OPERATION_BEFORE = "insert_before"

func sectionsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
}

func sectionsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["id"] = "id"
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
		"",
		"service_insertion_section",
		"POST",
		"/api/v1/serviceinsertion/sections",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsCreatewithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsCreatewithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
}

func sectionsCreatewithrulesRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	paramsTypeMap["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["id"] = "id"
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
		"action=create_with_rules",
		"service_insertion_section_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["cascade"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sectionsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["cascade"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cascade"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["cascade"] = "cascade"
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
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
}

func sectionsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["exclude_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["exclude_applied_to_type"] = "ExcludeAppliedToType"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["include_applied_to_type"] = "IncludeAppliedToType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionListResultBindingType)
}

func sectionsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["exclude_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["exclude_applied_to_type"] = "ExcludeAppliedToType"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["include_applied_to_type"] = "IncludeAppliedToType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["exclude_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["sources"] = "sources"
	queryParams["destinations"] = "destinations"
	queryParams["exclude_applied_to_type"] = "exclude_applied_to_type"
	queryParams["include_applied_to_type"] = "include_applied_to_type"
	queryParams["services"] = "services"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["applied_tos"] = "applied_tos"
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
		"/api/v1/serviceinsertion/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsListwithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsListwithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
}

func sectionsListwithrulesRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=list_with_rules",
		"",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsReviseInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsReviseOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
}

func sectionsReviseRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
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
		"service_insertion_section",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsRevisewithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsRevisewithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
}

func sectionsRevisewithrulesRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	paramsTypeMap["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
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
		"action=revise_with_rules",
		"service_insertion_section_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
}

func sectionsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section"] = "ServiceInsertionSection"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_insertion_section"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionBindingType)
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"service_insertion_section",
		"PUT",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsUpdatewithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsUpdatewithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
}

func sectionsUpdatewithrulesRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["service_insertion_section_rule_list"] = "ServiceInsertionSectionRuleList"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_insertion_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.ServiceInsertionSectionRuleListBindingType)
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=update_with_rules",
		"service_insertion_section_rule_list",
		"POST",
		"/api/v1/serviceinsertion/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
