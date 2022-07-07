// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Rules.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package sections

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func rulesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func rulesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
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

func rulesCreatemultipleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_rule_list"] = bindings.NewReferenceType(model.FirewallRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule_list"] = "FirewallRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesCreatemultipleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleListBindingType)
}

func rulesCreatemultipleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_rule_list"] = bindings.NewReferenceType(model.FirewallRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_rule_list"] = "FirewallRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["firewall_rule_list"] = bindings.NewReferenceType(model.FirewallRuleListBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["id"] = "id"
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

func rulesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func rulesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func rulesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["context_profiles"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["deep_search"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["extended_sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["search_invalid_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
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
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleListResultBindingType)
}

func rulesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["context_profiles"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["deep_search"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["extended_sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["search_invalid_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
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
	paramsTypeMap["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["context_profiles"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["extended_sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["search_invalid_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["services"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["deep_search"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
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
		"/api/v1/firewall/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func rulesReviseRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
	queryParams["id"] = "id"
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

func rulesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func rulesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["firewall_rule"] = "FirewallRule"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_rule"] = bindings.NewReferenceType(model.FirewallRuleBindingType)
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
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
