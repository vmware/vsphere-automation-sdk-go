// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: BridgePolicies.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package domains

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func bridgePoliciesCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	fieldNameMap["bridge_firewall_policy"] = "BridgeFirewallPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgePoliciesCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
}

func bridgePoliciesCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	fieldNameMap["bridge_firewall_policy"] = "BridgeFirewallPolicy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridgeFirewallPolicyId"] = vapiBindings_.NewStringType()
	pathParams["bridge_firewall_policy_id"] = "bridgeFirewallPolicyId"
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
		"bridge_firewall_policy",
		"POST",
		"/policy/api/v1/infra/domains/{domainId}/bridge-policies/{bridgeFirewallPolicyId}",
		"application/json",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgePoliciesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgePoliciesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bridgePoliciesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridgeFirewallPolicyId"] = vapiBindings_.NewStringType()
	pathParams["bridge_firewall_policy_id"] = "bridgeFirewallPolicyId"
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
		"/policy/api/v1/infra/domains/{domainId}/bridge-policies/{bridgeFirewallPolicyId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgePoliciesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgePoliciesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
}

func bridgePoliciesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridgeFirewallPolicyId"] = vapiBindings_.NewStringType()
	pathParams["bridge_firewall_policy_id"] = "bridgeFirewallPolicyId"
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
		"/policy/api/v1/infra/domains/{domainId}/bridge-policies/{bridgeFirewallPolicyId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgePoliciesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_rule_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["include_rule_count"] = "IncludeRuleCount"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgePoliciesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyListResultBindingType)
}

func bridgePoliciesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_rule_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["include_rule_count"] = "IncludeRuleCount"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_rule_count"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	pathParams["domain_id"] = "domainId"
	queryParams["cursor"] = "cursor"
	queryParams["include_rule_count"] = "include_rule_count"
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
		"/policy/api/v1/infra/domains/{domainId}/bridge-policies",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgePoliciesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	fieldNameMap["bridge_firewall_policy"] = "BridgeFirewallPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgePoliciesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bridgePoliciesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	fieldNameMap["bridge_firewall_policy"] = "BridgeFirewallPolicy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridgeFirewallPolicyId"] = vapiBindings_.NewStringType()
	pathParams["bridge_firewall_policy_id"] = "bridgeFirewallPolicyId"
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
		"bridge_firewall_policy",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/bridge-policies/{bridgeFirewallPolicyId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bridgePoliciesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	fieldNameMap["bridge_firewall_policy"] = "BridgeFirewallPolicy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BridgePoliciesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
}

func bridgePoliciesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	fields["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["bridge_firewall_policy_id"] = "BridgeFirewallPolicyId"
	fieldNameMap["bridge_firewall_policy"] = "BridgeFirewallPolicy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridge_firewall_policy"] = vapiBindings_.NewReferenceType(nsx_policyModel.BridgeFirewallPolicyBindingType)
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bridgeFirewallPolicyId"] = vapiBindings_.NewStringType()
	pathParams["bridge_firewall_policy_id"] = "bridgeFirewallPolicyId"
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
		"bridge_firewall_policy",
		"PUT",
		"/policy/api/v1/infra/domains/{domainId}/bridge-policies/{bridgeFirewallPolicyId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
