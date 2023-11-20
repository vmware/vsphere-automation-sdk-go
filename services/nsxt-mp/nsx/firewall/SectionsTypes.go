// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Sections.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

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

func sectionsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
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
	fields["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
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

func sectionsCreatewithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsCreatewithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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
	fields["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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
		"/api/v1/firewall/sections/{sectionId}",
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
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
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
		"/api/v1/firewall/sections/{sectionId}",
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
	fields["context_profiles"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["deep_search"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforced_on"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["exclude_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["extended_sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["locked"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["search_invalid_references"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["search_scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
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
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionListResultBindingType)
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
	fields["context_profiles"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["deep_search"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforced_on"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["exclude_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["extended_sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["locked"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["search_invalid_references"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["search_scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
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
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["deep_search"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["destinations"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["exclude_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_applied_to_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["context_profiles"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["services"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["search_scope"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["extended_sources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["filter_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["applied_tos"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforced_on"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["locked"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["search_invalid_references"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
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
		"/api/v1/firewall/sections",
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
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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
		"/api/v1/firewall/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sectionsLockInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section_lock"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsLockOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
}

func sectionsLockRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section_lock"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_section_lock"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionLockBindingType)
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

func sectionsReviseInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsReviseOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
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
	fields["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
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

func sectionsRevisewithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsRevisewithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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
	fields["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["operation"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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

func sectionsUnlockInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section_lock"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsUnlockOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
}

func sectionsUnlockRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section_lock"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_section_lock"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionLockBindingType)
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

func sectionsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
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
	fields["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	paramsTypeMap["firewall_section"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionBindingType)
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

func sectionsUpdatewithrulesInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = vapiBindings_.NewStringType()
	fields["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SectionsUpdatewithrulesOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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
	fields["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	paramsTypeMap["section_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_section_rule_list"] = vapiBindings_.NewReferenceType(nsxModel.FirewallSectionRuleListBindingType)
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
