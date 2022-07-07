// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Rules.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package intrusion_service_gateway_policies

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_TOP = "insert_top"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_BOTTOM = "insert_bottom"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_AFTER = "insert_after"

// Possible value for ``operation`` of method Rules#revise.
const Rules_REVISE_OPERATION_BEFORE = "insert_before"

func rulesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["policyId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["policy_id"] = "policyId"
	pathParams["rule_id"] = "ruleId"
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
		"/policy/api/v1/infra/domains/{domainId}/intrusion-service-gateway-policies/{policyId}/rules/{ruleId}",
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdsRuleBindingType)
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["policyId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["policy_id"] = "policyId"
	pathParams["rule_id"] = "ruleId"
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
		"/policy/api/v1/infra/domains/{domainId}/intrusion-service-gateway-policies/{policyId}/rules/{ruleId}",
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdsRuleListResultBindingType)
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["policy_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["policyId"] = bindings.NewStringType()
	pathParams["policy_id"] = "policyId"
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
		"/policy/api/v1/infra/domains/{domainId}/intrusion-service-gateway-policies/{policyId}/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["ids_rule"] = "IdsRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func rulesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["ids_rule"] = "IdsRule"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	paramsTypeMap["policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["policyId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["policy_id"] = "policyId"
	pathParams["rule_id"] = "ruleId"
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
		"ids_rule",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/intrusion-service-gateway-policies/{policyId}/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func rulesReviseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	fields["anchor_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["ids_rule"] = "IdsRule"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesReviseOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdsRuleBindingType)
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	fields["anchor_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["ids_rule"] = "IdsRule"
	fieldNameMap["anchor_path"] = "AnchorPath"
	fieldNameMap["operation"] = "Operation"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	paramsTypeMap["policy_id"] = bindings.NewStringType()
	paramsTypeMap["anchor_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["operation"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["policyId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["policy_id"] = "policyId"
	pathParams["rule_id"] = "ruleId"
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
		"ids_rule",
		"POST",
		"/policy/api/v1/infra/domains/{domainId}/intrusion-service-gateway-policies/{policyId}/rules/{ruleId}",
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["ids_rule"] = "IdsRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rulesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdsRuleBindingType)
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
	fields["domain_id"] = bindings.NewStringType()
	fields["policy_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["policy_id"] = "PolicyId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["ids_rule"] = "IdsRule"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["ids_rule"] = bindings.NewReferenceType(model.IdsRuleBindingType)
	paramsTypeMap["policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["policyId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["policy_id"] = "policyId"
	pathParams["rule_id"] = "ruleId"
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
		"ids_rule",
		"PUT",
		"/policy/api/v1/infra/domains/{domainId}/intrusion-service-gateway-policies/{policyId}/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
