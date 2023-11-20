// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Rules.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package sections

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``operation`` of method Rules#create.
const Rules_CREATE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Rules#create.
const Rules_CREATE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Rules#create.
const Rules_CREATE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Rules#create.
const Rules_CREATE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method Rules#createmultiple.
const Rules_CREATEMULTIPLE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Rules#createmultiple.
const Rules_CREATEMULTIPLE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Rules#createmultiple.
const Rules_CREATEMULTIPLE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Rules#createmultiple.
const Rules_CREATEMULTIPLE_OPERATION_BEFORE = "insert_before"

// Possible value for ``filterType`` of method Rules#list.
const Rules_LIST_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method Rules#list.
const Rules_LIST_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_BEFORE = "insert_before"

func rulesCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
}

func rulesCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
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
		"",
		"firewall_rule",
		"POST",
		"/api/v1/firewall/sections/{sectionId}/rules",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesCreatemultipleInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule_list"] = "FirewallRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesCreatemultipleOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallRuleListBindingType)
}

func rulesCreatemultipleRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule_list"] = "FirewallRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleListBindingType)
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
		"action=create_multiple",
		"firewall_rule_list",
		"POST",
		"/api/v1/firewall/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func rulesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ruleId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
}

func rulesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ruleId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["context_profiles"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["deep_search"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["extended_sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["search_invalid_references"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["context_profiles"] = "ContextProfiles"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["deep_search"] = "DeepSearch"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["extended_sources"] = "ExtendedSources"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["search_invalid_references"] = "SearchInvalidReferences"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallRuleListResultBindingType)
}

func rulesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["context_profiles"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["deep_search"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["extended_sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["search_invalid_references"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["context_profiles"] = "ContextProfiles"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["deep_search"] = "DeepSearch"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["extended_sources"] = "ExtendedSources"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["search_invalid_references"] = "SearchInvalidReferences"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["deep_search"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["context_profiles"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["extended_sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["search_invalid_references"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["cursor"] = "cursor"
	queryParams["sources"] = "sources"
	queryParams["deep_search"] = "deep_search"
	queryParams["destinations"] = "destinations"
	queryParams["context_profiles"] = "context_profiles"
	queryParams["services"] = "services"
	queryParams["sort_by"] = "sort_by"
	queryParams["extended_sources"] = "extended_sources"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["applied_tos"] = "applied_tos"
	queryParams["search_invalid_references"] = "search_invalid_references"
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
		"/api/v1/firewall/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesReviseInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesReviseOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
}

func rulesReviseRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	paramsTypeMap["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ruleId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"firewall_rule",
		"POST",
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RulesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
}

func rulesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["rule_id"] = vapiBindings_.NewStringType()
	fields["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	paramsTypeMap["rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_rule"] = vapiBindings_.NewReferenceType(nsxModel.FirewallRuleBindingType)
	paramsTypeMap["sectionId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ruleId"] = vapiBindings_.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"firewall_rule",
		"PUT",
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
