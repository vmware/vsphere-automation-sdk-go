// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Sections.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_VIF = "VIF"

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_LOGICALROUTER = "LOGICALROUTER"

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_BRIDGEENDPOINT = "BRIDGEENDPOINT"

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_DHCP_SERVICE = "DHCP_SERVICE"

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_METADATA_PROXY = "METADATA_PROXY"

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_L2VPN_SESSION = "L2VPN_SESSION"

// Possible value for ``enforcedOn`` of method Sections#list.
const Sections_LIST_ENFORCED_ON_NONE = "NONE"

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
const Sections_LIST_TYPE_LAYER2 = "LAYER2"

// Possible value for ``type`` of method Sections#list.
const Sections_LIST_TYPE_LAYER3 = "LAYER3"

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

func sectionsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func sectionsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
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
		"firewall_section",
		"POST",
		"/api/v1/firewall/sections",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsCreatewithrulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsCreatewithrulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func sectionsCreatewithrulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
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
		"action=create_with_rules",
		"firewall_section_rule_list",
		"POST",
		"/api/v1/firewall/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sectionsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["cascade"] = "cascade"
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
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func sectionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["context_profiles"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["deep_search"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enforced_on"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["extended_sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["locked"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["search_invalid_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["search_scope"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["context_profiles"] = "ContextProfiles"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["deep_search"] = "DeepSearch"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["enforced_on"] = "EnforcedOn"
	fieldNameMap["exclude_applied_to_type"] = "ExcludeAppliedToType"
	fieldNameMap["extended_sources"] = "ExtendedSources"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["include_applied_to_type"] = "IncludeAppliedToType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["locked"] = "Locked"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["search_invalid_references"] = "SearchInvalidReferences"
	fieldNameMap["search_scope"] = "SearchScope"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionListResultBindingType)
}

func sectionsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["context_profiles"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["deep_search"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enforced_on"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["extended_sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["locked"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["search_invalid_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["search_scope"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["context_profiles"] = "ContextProfiles"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["deep_search"] = "DeepSearch"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["enforced_on"] = "EnforcedOn"
	fieldNameMap["exclude_applied_to_type"] = "ExcludeAppliedToType"
	fieldNameMap["extended_sources"] = "ExtendedSources"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["include_applied_to_type"] = "IncludeAppliedToType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["locked"] = "Locked"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["search_invalid_references"] = "SearchInvalidReferences"
	fieldNameMap["search_scope"] = "SearchScope"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["enforced_on"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["context_profiles"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["search_scope"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["extended_sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["search_invalid_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["services"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["locked"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["deep_search"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sources"] = "sources"
	queryParams["deep_search"] = "deep_search"
	queryParams["destinations"] = "destinations"
	queryParams["exclude_applied_to_type"] = "exclude_applied_to_type"
	queryParams["include_applied_to_type"] = "include_applied_to_type"
	queryParams["context_profiles"] = "context_profiles"
	queryParams["services"] = "services"
	queryParams["sort_by"] = "sort_by"
	queryParams["search_scope"] = "search_scope"
	queryParams["type"] = "type"
	queryParams["extended_sources"] = "extended_sources"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["applied_tos"] = "applied_tos"
	queryParams["enforced_on"] = "enforced_on"
	queryParams["locked"] = "locked"
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
		"/api/v1/firewall/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsListwithrulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsListwithrulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func sectionsListwithrulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=list_with_rules",
		"",
		"POST",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsLockInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsLockOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func sectionsLockRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=lock",
		"firewall_section_lock",
		"POST",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.resource_busy": 500, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func sectionsReviseRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
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
		"action=revise",
		"firewall_section",
		"POST",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsRevisewithrulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsRevisewithrulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func sectionsRevisewithrulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
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
		"action=revise_with_rules",
		"firewall_section_rule_list",
		"POST",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsUnlockInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsUnlockOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func sectionsUnlockRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=unlock",
		"firewall_section_lock",
		"POST",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.resource_busy": 500, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func sectionsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"firewall_section",
		"PUT",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsUpdatewithrulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sectionsUpdatewithrulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func sectionsUpdatewithrulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	paramsTypeMap["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
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
		"action=update_with_rules",
		"firewall_section_rule_list",
		"POST",
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
