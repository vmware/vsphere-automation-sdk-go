// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPISecurityServicesFirewall.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRuleInSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULE_IN_SECTION_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRuleInSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULE_IN_SECTION_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRuleInSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULE_IN_SECTION_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRuleInSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULE_IN_SECTION_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addRulesInSectionCreateMultiple.
const ManagementPlaneAPISecurityServicesFirewall_ADD_RULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSection.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#addSectionWithRulesCreateWithRules.
const ManagementPlaneAPISecurityServicesFirewall_ADD_SECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_BEFORE = "insert_before"

// Possible value for ``objectType`` of method ManagementPlaneAPISecurityServicesFirewall#checkMemberIfExistsCheckIfExists.
const ManagementPlaneAPISecurityServicesFirewall_CHECK_MEMBER_IF_EXISTS_CHECK_IF_EXISTS_OBJECT_TYPE_NSGROUP = "NSGroup"

// Possible value for ``objectType`` of method ManagementPlaneAPISecurityServicesFirewall#checkMemberIfExistsCheckIfExists.
const ManagementPlaneAPISecurityServicesFirewall_CHECK_MEMBER_IF_EXISTS_CHECK_IF_EXISTS_OBJECT_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``objectType`` of method ManagementPlaneAPISecurityServicesFirewall#checkMemberIfExistsCheckIfExists.
const ManagementPlaneAPISecurityServicesFirewall_CHECK_MEMBER_IF_EXISTS_CHECK_IF_EXISTS_OBJECT_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesFirewall#getFirewallSectionStats.
const ManagementPlaneAPISecurityServicesFirewall_GET_FIREWALL_SECTION_STATS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesFirewall#getFirewallSectionStats.
const ManagementPlaneAPISecurityServicesFirewall_GET_FIREWALL_SECTION_STATS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesFirewall#getFirewallStats.
const ManagementPlaneAPISecurityServicesFirewall_GET_FIREWALL_STATS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesFirewall#getFirewallStats.
const ManagementPlaneAPISecurityServicesFirewall_GET_FIREWALL_STATS_SOURCE_CACHED = "cached"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesFirewall#getRules.
const ManagementPlaneAPISecurityServicesFirewall_GET_RULES_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesFirewall#getRules.
const ManagementPlaneAPISecurityServicesFirewall_GET_RULES_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesFirewall#getSectionsSummary.
const ManagementPlaneAPISecurityServicesFirewall_GET_SECTIONS_SUMMARY_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPISecurityServicesFirewall#getSectionsSummary.
const ManagementPlaneAPISecurityServicesFirewall_GET_SECTIONS_SUMMARY_SOURCE_CACHED = "cached"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_VIF = "VIF"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_LOGICALROUTER = "LOGICALROUTER"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_BRIDGEENDPOINT = "BRIDGEENDPOINT"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_DHCP_SERVICE = "DHCP_SERVICE"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_METADATA_PROXY = "METADATA_PROXY"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_L2VPN_SESSION = "L2VPN_SESSION"

// Possible value for ``enforcedOn`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_ENFORCED_ON_NONE = "NONE"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_INCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``type`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_TYPE_LAYER2 = "LAYER2"

// Possible value for ``type`` of method ManagementPlaneAPISecurityServicesFirewall#listSections.
const ManagementPlaneAPISecurityServicesFirewall_LIST_SECTIONS_TYPE_LAYER3 = "LAYER3"

// Possible value for ``objectType`` of method ManagementPlaneAPISecurityServicesFirewall#removeMemberRemoveMember.
const ManagementPlaneAPISecurityServicesFirewall_REMOVE_MEMBER_REMOVE_MEMBER_OBJECT_TYPE_NSGROUP = "NSGroup"

// Possible value for ``objectType`` of method ManagementPlaneAPISecurityServicesFirewall#removeMemberRemoveMember.
const ManagementPlaneAPISecurityServicesFirewall_REMOVE_MEMBER_REMOVE_MEMBER_OBJECT_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``objectType`` of method ManagementPlaneAPISecurityServicesFirewall#removeMemberRemoveMember.
const ManagementPlaneAPISecurityServicesFirewall_REMOVE_MEMBER_REMOVE_MEMBER_OBJECT_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``category`` of method ManagementPlaneAPISecurityServicesFirewall#resetFirewallRuleStatsReset.
const ManagementPlaneAPISecurityServicesFirewall_RESET_FIREWALL_RULE_STATS_RESET_CATEGORY_L3DFW = "L3DFW"

// Possible value for ``category`` of method ManagementPlaneAPISecurityServicesFirewall#resetFirewallRuleStatsReset.
const ManagementPlaneAPISecurityServicesFirewall_RESET_FIREWALL_RULE_STATS_RESET_CATEGORY_L3EDGE = "L3EDGE"

// Possible value for ``category`` of method ManagementPlaneAPISecurityServicesFirewall#resetFirewallRuleStatsReset.
const ManagementPlaneAPISecurityServicesFirewall_RESET_FIREWALL_RULE_STATS_RESET_CATEGORY_L3BRIDGEPORT = "L3BRIDGEPORT"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseRuleRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_RULE_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseRuleRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_RULE_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseRuleRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_RULE_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseRuleRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_RULE_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionRevise.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPISecurityServicesFirewall#reviseSectionWithRulesReviseWithRules.
const ManagementPlaneAPISecurityServicesFirewall_REVISE_SECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_BEFORE = "insert_before"

func managementPlaneAPISecurityServicesFirewallAddMemberAddMemberInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallAddMemberAddMemberOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func managementPlaneAPISecurityServicesFirewallAddMemberAddMemberRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	paramsTypeMap["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
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
		"action=add_member",
		"resource_reference",
		"POST",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallAddRuleInSectionInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallAddRuleInSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func managementPlaneAPISecurityServicesFirewallAddRuleInSectionRestMetadata() protocol.OperationRestMetadata {
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
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallAddRulesInSectionCreateMultipleInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallAddRulesInSectionCreateMultipleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleListBindingType)
}

func managementPlaneAPISecurityServicesFirewallAddRulesInSectionCreateMultipleRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallAddSectionInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallAddSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func managementPlaneAPISecurityServicesFirewallAddSectionRestMetadata() protocol.OperationRestMetadata {
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
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallAddSectionWithRulesCreateWithRulesInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallAddSectionWithRulesCreateWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesFirewallAddSectionWithRulesCreateWithRulesRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallCheckMemberIfExistsCheckIfExistsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallCheckMemberIfExistsCheckIfExistsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func managementPlaneAPISecurityServicesFirewallCheckMemberIfExistsCheckIfExistsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	paramsTypeMap["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["object_id"] = bindings.NewStringType()
	paramsTypeMap["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["deep_check"] = "deep_check"
	queryParams["object_type"] = "object_type"
	queryParams["object_id"] = "object_id"
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
		"action=check_if_exists",
		"",
		"POST",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallCreateFirewallProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["base_firewall_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
	fieldNameMap["base_firewall_profile"] = "BaseFirewallProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallCreateFirewallProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesFirewallCreateFirewallProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["base_firewall_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
	fieldNameMap["base_firewall_profile"] = "BaseFirewallProfile"
	paramsTypeMap["base_firewall_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
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
		"base_firewall_profile",
		"POST",
		"/api/v1/firewall/profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallDeleteFirewallProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile_id"] = bindings.NewStringType()
	fieldNameMap["profile_id"] = "ProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallDeleteFirewallProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesFirewallDeleteFirewallProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["profile_id"] = bindings.NewStringType()
	fieldNameMap["profile_id"] = "ProfileId"
	paramsTypeMap["profile_id"] = bindings.NewStringType()
	paramsTypeMap["profileId"] = bindings.NewStringType()
	pathParams["profile_id"] = "profileId"
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
		"/api/v1/firewall/profiles/{profileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallDeleteRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallDeleteRuleOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesFirewallDeleteRuleRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallDeleteSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallDeleteSectionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesFirewallDeleteSectionRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallDisableFirewallOnTargetResourceDisableFirewallInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallDisableFirewallOnTargetResourceDisableFirewallOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TargetResourceStatusBindingType)
}

func managementPlaneAPISecurityServicesFirewallDisableFirewallOnTargetResourceDisableFirewallRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
	pathParams["id"] = "id"
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
		"action=disable_firewall",
		"",
		"POST",
		"/api/v1/firewall/status/{contextType}/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallEnableFirewallOnTargetResourceEnableFirewallInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallEnableFirewallOnTargetResourceEnableFirewallOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TargetResourceStatusBindingType)
}

func managementPlaneAPISecurityServicesFirewallEnableFirewallOnTargetResourceEnableFirewallRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
	pathParams["id"] = "id"
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
		"action=enable_firewall",
		"",
		"POST",
		"/api/v1/firewall/status/{contextType}/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetExcludeListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetExcludeListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExcludeListBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetExcludeListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetFirewallProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile_id"] = bindings.NewStringType()
	fieldNameMap["profile_id"] = "ProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["profile_id"] = bindings.NewStringType()
	fieldNameMap["profile_id"] = "ProfileId"
	paramsTypeMap["profile_id"] = bindings.NewStringType()
	paramsTypeMap["profileId"] = bindings.NewStringType()
	pathParams["profile_id"] = "profileId"
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
		"/api/v1/firewall/profiles/{profileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetFirewallSectionStatsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallSectionStatsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallStatsListBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallSectionStatsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["source"] = "source"
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
		"/api/v1/firewall/sections/{sectionId}/rules/stats",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallStatsBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatsRestMetadata() protocol.OperationRestMetadata {
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
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	pathParams["rule_id"] = "ruleId"
	queryParams["source"] = "source"
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
		"/api/v1/firewall/sections/{sectionId}/rules/{ruleId}/stats",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallStatusBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
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
		"/api/v1/firewall/status/{contextType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatusOnTargetResourceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatusOnTargetResourceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TargetResourceStatusBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetFirewallStatusOnTargetResourceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
	pathParams["id"] = "id"
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
		"/api/v1/firewall/status/{contextType}/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetRuleRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallGetRuleStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetRuleStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RuleStateBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetRuleStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["rule_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["rule_id"] = "ruleId"
	queryParams["barrier_id"] = "barrier_id"
	queryParams["request_id"] = "request_id"
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
		"/api/v1/firewall/rules/{ruleId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetRulesInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallGetRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleListResultBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetRulesRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallGetSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetSectionRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallGetSectionStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetSectionStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionStateBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetSectionStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	paramsTypeMap["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["barrier_id"] = "barrier_id"
	queryParams["request_id"] = "request_id"
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
		"/api/v1/firewall/sections/{sectionId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallGetSectionWithRulesListWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetSectionWithRulesListWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetSectionWithRulesListWithRulesRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallGetSectionsSummaryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallGetSectionsSummaryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionsSummaryListBindingType)
}

func managementPlaneAPISecurityServicesFirewallGetSectionsSummaryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["source"] = "source"
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
		"/api/v1/firewall/sections/summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallListFirewallProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallListFirewallProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallProfileListResultBindingType)
}

func managementPlaneAPISecurityServicesFirewallListFirewallProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["resource_type"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["resource_type"] = "resource_type"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/firewall/profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallListFirewallStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallListFirewallStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallStatusListResultBindingType)
}

func managementPlaneAPISecurityServicesFirewallListFirewallStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/firewall/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallListSectionsInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallListSectionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionListResultBindingType)
}

func managementPlaneAPISecurityServicesFirewallListSectionsRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallLockSectionLockInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallLockSectionLockOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func managementPlaneAPISecurityServicesFirewallLockSectionLockRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallReadFirewallRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallReadFirewallRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func managementPlaneAPISecurityServicesFirewallReadFirewallRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
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
		"/api/v1/firewall/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallRemoveMemberRemoveMemberInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallRemoveMemberRemoveMemberOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func managementPlaneAPISecurityServicesFirewallRemoveMemberRemoveMemberRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	paramsTypeMap["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["object_id"] = bindings.NewStringType()
	paramsTypeMap["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["deep_check"] = "deep_check"
	queryParams["object_type"] = "object_type"
	queryParams["object_id"] = "object_id"
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
		"action=remove_member",
		"",
		"POST",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallResetFirewallRuleStatsResetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewStringType()
	fieldNameMap["category"] = "Category"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallResetFirewallRuleStatsResetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityServicesFirewallResetFirewallRuleStatsResetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["category"] = bindings.NewStringType()
	fieldNameMap["category"] = "Category"
	paramsTypeMap["category"] = bindings.NewStringType()
	queryParams["category"] = "category"
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
		"action=reset",
		"",
		"POST",
		"/api/v1/firewall/stats",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallReviseRuleReviseInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallReviseRuleReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func managementPlaneAPISecurityServicesFirewallReviseRuleReviseRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallReviseSectionReviseInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallReviseSectionReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func managementPlaneAPISecurityServicesFirewallReviseSectionReviseRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallReviseSectionWithRulesReviseWithRulesInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallReviseSectionWithRulesReviseWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesFirewallReviseSectionWithRulesReviseWithRulesRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallUnlockSectionUnlockInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_lock"] = bindings.NewReferenceType(model.FirewallSectionLockBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_lock"] = "FirewallSectionLock"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallUnlockSectionUnlockOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func managementPlaneAPISecurityServicesFirewallUnlockSectionUnlockRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallUpdateExcludeListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["exclude_list"] = bindings.NewReferenceType(model.ExcludeListBindingType)
	fieldNameMap["exclude_list"] = "ExcludeList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallUpdateExcludeListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExcludeListBindingType)
}

func managementPlaneAPISecurityServicesFirewallUpdateExcludeListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["exclude_list"] = bindings.NewReferenceType(model.ExcludeListBindingType)
	fieldNameMap["exclude_list"] = "ExcludeList"
	paramsTypeMap["exclude_list"] = bindings.NewReferenceType(model.ExcludeListBindingType)
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
		"exclude_list",
		"PUT",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallUpdateFirewallProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile_id"] = bindings.NewStringType()
	fields["base_firewall_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
	fieldNameMap["profile_id"] = "ProfileId"
	fieldNameMap["base_firewall_profile"] = "BaseFirewallProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallUpdateFirewallProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
}

func managementPlaneAPISecurityServicesFirewallUpdateFirewallProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["profile_id"] = bindings.NewStringType()
	fields["base_firewall_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
	fieldNameMap["profile_id"] = "ProfileId"
	fieldNameMap["base_firewall_profile"] = "BaseFirewallProfile"
	paramsTypeMap["base_firewall_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseFirewallProfileBindingType)}, bindings.REST)
	paramsTypeMap["profile_id"] = bindings.NewStringType()
	paramsTypeMap["profileId"] = bindings.NewStringType()
	pathParams["profile_id"] = "profileId"
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
		"base_firewall_profile",
		"PUT",
		"/api/v1/firewall/profiles/{profileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallUpdateFirewallStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["context_type"] = bindings.NewStringType()
	fields["firewall_status"] = bindings.NewReferenceType(model.FirewallStatusBindingType)
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["firewall_status"] = "FirewallStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallUpdateFirewallStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallStatusBindingType)
}

func managementPlaneAPISecurityServicesFirewallUpdateFirewallStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["context_type"] = bindings.NewStringType()
	fields["firewall_status"] = bindings.NewReferenceType(model.FirewallStatusBindingType)
	fieldNameMap["context_type"] = "ContextType"
	fieldNameMap["firewall_status"] = "FirewallStatus"
	paramsTypeMap["firewall_status"] = bindings.NewReferenceType(model.FirewallStatusBindingType)
	paramsTypeMap["context_type"] = bindings.NewStringType()
	paramsTypeMap["contextType"] = bindings.NewStringType()
	pathParams["context_type"] = "contextType"
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
		"firewall_status",
		"PUT",
		"/api/v1/firewall/status/{contextType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityServicesFirewallUpdateRuleInputType() bindings.StructType {
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

func managementPlaneAPISecurityServicesFirewallUpdateRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallRuleBindingType)
}

func managementPlaneAPISecurityServicesFirewallUpdateRuleRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallUpdateSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section"] = bindings.NewReferenceType(model.FirewallSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section"] = "FirewallSection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallUpdateSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionBindingType)
}

func managementPlaneAPISecurityServicesFirewallUpdateSectionRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPISecurityServicesFirewallUpdateSectionWithRulesUpdateWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["firewall_section_rule_list"] = bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["firewall_section_rule_list"] = "FirewallSectionRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityServicesFirewallUpdateSectionWithRulesUpdateWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FirewallSectionRuleListBindingType)
}

func managementPlaneAPISecurityServicesFirewallUpdateSectionWithRulesUpdateWithRulesRestMetadata() protocol.OperationRestMetadata {
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
