// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingServicesPolicyBasedRouting.
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

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRuleInSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULE_IN_SECTION_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRuleInSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULE_IN_SECTION_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRuleInSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULE_IN_SECTION_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRuleInSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULE_IN_SECTION_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRulesInSectionCreateMultiple.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRulesInSectionCreateMultiple.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRulesInSectionCreateMultiple.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRRulesInSectionCreateMultiple.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRRULES_IN_SECTION_CREATE_MULTIPLE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSection.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSectionWithRulesCreateWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSectionWithRulesCreateWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSectionWithRulesCreateWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#addPBRSectionWithRulesCreateWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_ADD_PBRSECTION_WITH_RULES_CREATE_WITH_RULES_OPERATION_BEFORE = "insert_before"

// Possible value for ``filterType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#getPBRRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_GET_PBRRULES_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#getPBRRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_GET_PBRRULES_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_EXCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``excludeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_EXCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``filterType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_FILTER_TYPE_FILTER = "FILTER"

// Possible value for ``filterType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_FILTER_TYPE_SEARCH = "SEARCH"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_INCLUDE_APPLIED_TO_TYPE_NSGROUP = "NSGroup"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALROUTER = "LogicalRouter"

// Possible value for ``includeAppliedToType`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#listPBRSections.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_LIST_PBRSECTIONS_INCLUDE_APPLIED_TO_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRRuleRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRRULE_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRRuleRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRRULE_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRRuleRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRRULE_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRRuleRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRRULE_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionRevise.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_REVISE_OPERATION_BEFORE = "insert_before"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionWithRulesReviseWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionWithRulesReviseWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionWithRulesReviseWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method ManagementPlaneAPINetworkingServicesPolicyBasedRouting#revisePBRSectionWithRulesReviseWithRules.
const ManagementPlaneAPINetworkingServicesPolicyBasedRouting_REVISE_PBRSECTION_WITH_RULES_REVISE_WITH_RULES_OPERATION_BEFORE = "insert_before"

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRRuleInSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_rule"] = "PBRRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRRuleInSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRRuleBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRRuleInSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_rule"] = "PBRRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
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
		"p_BR_rule",
		"POST",
		"/api/v1/pbr/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRRulesInSectionCreateMultipleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_rule_list"] = bindings.NewReferenceType(model.PBRRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_rule_list"] = "PBRRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRRulesInSectionCreateMultipleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRRuleListBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRRulesInSectionCreateMultipleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_rule_list"] = bindings.NewReferenceType(model.PBRRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_rule_list"] = "PBRRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_rule_list"] = bindings.NewReferenceType(model.PBRRuleListBindingType)
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
		"p_BR_rule_list",
		"POST",
		"/api/v1/pbr/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["p_BR_section"] = "PBRSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["p_BR_section"] = "PBRSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
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
		"p_BR_section",
		"POST",
		"/api/v1/pbr/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRSectionWithRulesCreateWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["p_BR_section_rule_list"] = "PBRSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRSectionWithRulesCreateWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingAddPBRSectionWithRulesCreateWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["p_BR_section_rule_list"] = "PBRSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
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
		"p_BR_section_rule_list",
		"POST",
		"/api/v1/pbr/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingDeletePBRRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingDeletePBRRuleOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingDeletePBRRuleRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingDeletePBRSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingDeletePBRSectionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingDeletePBRSectionRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRRuleBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRuleRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRuleStatsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRuleStatsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRStatsBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRuleStatsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}/rules/{ruleId}/stats",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRRuleListResultBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRRulesRestMetadata() protocol.OperationRestMetadata {
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
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["applied_tos"] = "AppliedTos"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["destinations"] = "Destinations"
	fieldNameMap["filter_type"] = "FilterType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["services"] = "Services"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["sources"] = "Sources"
	paramsTypeMap["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["services"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sectionId"] = bindings.NewStringType()
	pathParams["section_id"] = "sectionId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["sources"] = "sources"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["destinations"] = "destinations"
	queryParams["applied_tos"] = "applied_tos"
	queryParams["services"] = "services"
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
		"/api/v1/pbr/sections/{sectionId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionStatsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionStatsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRStatsListBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionStatsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}/rules/stats",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionWithRulesListWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fieldNameMap["section_id"] = "SectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionWithRulesListWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingGetPBRSectionWithRulesListWithRulesRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingListPBRSectionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
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
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingListPBRSectionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionListResultBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingListPBRSectionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["services"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sources"] = bindings.NewOptionalType(bindings.NewStringType())
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
	paramsTypeMap["sources"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["exclude_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_tos"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["include_applied_to_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["services"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["destinations"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["sources"] = "sources"
	queryParams["filter_type"] = "filter_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["destinations"] = "destinations"
	queryParams["applied_tos"] = "applied_tos"
	queryParams["exclude_applied_to_type"] = "exclude_applied_to_type"
	queryParams["include_applied_to_type"] = "include_applied_to_type"
	queryParams["services"] = "services"
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
		"/api/v1/pbr/sections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRRuleReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["p_BR_rule"] = "PBRRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRRuleReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRRuleBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRRuleReviseRestMetadata() protocol.OperationRestMetadata {
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
	fields["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["p_BR_rule"] = "PBRRule"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
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
		"p_BR_rule",
		"POST",
		"/api/v1/pbr/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRSectionReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section"] = "PBRSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRSectionReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRSectionReviseRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section"] = "PBRSection"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
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
		"p_BR_section",
		"POST",
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRSectionWithRulesReviseWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section_rule_list"] = "PBRSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRSectionWithRulesReviseWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingRevisePBRSectionWithRulesReviseWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section_rule_list"] = "PBRSectionRuleList"
	fieldNameMap["id"] = "Id"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
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
		"p_BR_section_rule_list",
		"POST",
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["p_BR_rule"] = "PBRRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRRuleBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRRuleRestMetadata() protocol.OperationRestMetadata {
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
	fields["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["p_BR_rule"] = "PBRRule"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["p_BR_rule"] = bindings.NewReferenceType(model.PBRRuleBindingType)
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
		"p_BR_rule",
		"PUT",
		"/api/v1/pbr/sections/{sectionId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRSectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section"] = "PBRSection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRSectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRSectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section"] = "PBRSection"
	paramsTypeMap["section_id"] = bindings.NewStringType()
	paramsTypeMap["p_BR_section"] = bindings.NewReferenceType(model.PBRSectionBindingType)
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
		"p_BR_section",
		"PUT",
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRSectionWithRulesUpdateWithRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section_rule_list"] = "PBRSectionRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRSectionWithRulesUpdateWithRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
}

func managementPlaneAPINetworkingServicesPolicyBasedRoutingUpdatePBRSectionWithRulesUpdateWithRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["section_id"] = bindings.NewStringType()
	fields["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
	fieldNameMap["section_id"] = "SectionId"
	fieldNameMap["p_BR_section_rule_list"] = "PBRSectionRuleList"
	paramsTypeMap["p_BR_section_rule_list"] = bindings.NewReferenceType(model.PBRSectionRuleListBindingType)
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
		"p_BR_section_rule_list",
		"POST",
		"/api/v1/pbr/sections/{sectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
