// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EndpointRules.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package endpoint_policies

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func endpointRulesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endpointRulesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func endpointRulesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_rule_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["endpointPolicyId"] = bindings.NewStringType()
	paramsTypeMap["endpointRuleId"] = bindings.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
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
		"/policy/api/v1/global-infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endpointRulesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EndpointRuleBindingType)
}

func endpointRulesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_rule_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["endpointPolicyId"] = bindings.NewStringType()
	paramsTypeMap["endpointRuleId"] = bindings.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
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
		"/policy/api/v1/global-infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endpointRulesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EndpointRuleListResultBindingType)
}

func endpointRulesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["endpoint_policy_id"] = bindings.NewStringType()
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["endpointPolicyId"] = bindings.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
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
		"/policy/api/v1/global-infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fields["endpoint_rule"] = bindings.NewReferenceType(model.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endpointRulesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func endpointRulesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fields["endpoint_rule"] = bindings.NewReferenceType(model.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_rule_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_rule"] = bindings.NewReferenceType(model.EndpointRuleBindingType)
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["endpointPolicyId"] = bindings.NewStringType()
	paramsTypeMap["endpointRuleId"] = bindings.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
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
		"endpoint_rule",
		"PATCH",
		"/policy/api/v1/global-infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func endpointRulesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fields["endpoint_rule"] = bindings.NewReferenceType(model.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endpointRulesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EndpointRuleBindingType)
}

func endpointRulesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["endpoint_policy_id"] = bindings.NewStringType()
	fields["endpoint_rule_id"] = bindings.NewStringType()
	fields["endpoint_rule"] = bindings.NewReferenceType(model.EndpointRuleBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["endpoint_policy_id"] = "EndpointPolicyId"
	fieldNameMap["endpoint_rule_id"] = "EndpointRuleId"
	fieldNameMap["endpoint_rule"] = "EndpointRule"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_rule_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_policy_id"] = bindings.NewStringType()
	paramsTypeMap["endpoint_rule"] = bindings.NewReferenceType(model.EndpointRuleBindingType)
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["endpointPolicyId"] = bindings.NewStringType()
	paramsTypeMap["endpointRuleId"] = bindings.NewStringType()
	pathParams["endpoint_policy_id"] = "endpointPolicyId"
	pathParams["endpoint_rule_id"] = "endpointRuleId"
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
		"endpoint_rule",
		"PUT",
		"/policy/api/v1/global-infra/domains/{domainId}/endpoint-policies/{endpointPolicyId}/endpoint-rules/{endpointRuleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
