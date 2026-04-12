// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DvsTeamingPolicies.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package enforcement_points

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func dvsTeamingPoliciesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DvsTeamingPoliciesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func dvsTeamingPoliciesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	paramsTypeMap["teaming_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["teamingPolicyId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	pathParams["teaming_policy_id"] = "teamingPolicyId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/dvs-teaming-policies/{teamingPolicyId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dvsTeamingPoliciesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DvsTeamingPoliciesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
}

func dvsTeamingPoliciesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	paramsTypeMap["teaming_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["teamingPolicyId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	pathParams["teaming_policy_id"] = "teamingPolicyId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/dvs-teaming-policies/{teamingPolicyId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dvsTeamingPoliciesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["dvs_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["dvs_id"] = "DvsId"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DvsTeamingPoliciesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingListResultBindingType)
}

func dvsTeamingPoliciesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["dvs_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["dvs_id"] = "DvsId"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["dvs_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	queryParams["cursor"] = "cursor"
	queryParams["dvs_id"] = "dvs_id"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/dvs-teaming-policies",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dvsTeamingPoliciesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fields["distributed_virtual_switch_named_teaming"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	fieldNameMap["distributed_virtual_switch_named_teaming"] = "DistributedVirtualSwitchNamedTeaming"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DvsTeamingPoliciesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
}

func dvsTeamingPoliciesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fields["distributed_virtual_switch_named_teaming"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	fieldNameMap["distributed_virtual_switch_named_teaming"] = "DistributedVirtualSwitchNamedTeaming"
	paramsTypeMap["distributed_virtual_switch_named_teaming"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
	paramsTypeMap["teaming_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["teamingPolicyId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	pathParams["teaming_policy_id"] = "teamingPolicyId"
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
		"distributed_virtual_switch_named_teaming",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/dvs-teaming-policies/{teamingPolicyId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dvsTeamingPoliciesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fields["distributed_virtual_switch_named_teaming"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	fieldNameMap["distributed_virtual_switch_named_teaming"] = "DistributedVirtualSwitchNamedTeaming"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DvsTeamingPoliciesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
}

func dvsTeamingPoliciesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["teaming_policy_id"] = vapiBindings_.NewStringType()
	fields["distributed_virtual_switch_named_teaming"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["teaming_policy_id"] = "TeamingPolicyId"
	fieldNameMap["distributed_virtual_switch_named_teaming"] = "DistributedVirtualSwitchNamedTeaming"
	paramsTypeMap["distributed_virtual_switch_named_teaming"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVirtualSwitchNamedTeamingBindingType)
	paramsTypeMap["teaming_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["teamingPolicyId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	pathParams["teaming_policy_id"] = "teamingPolicyId"
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
		"distributed_virtual_switch_named_teaming",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/dvs-teaming-policies/{teamingPolicyId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
