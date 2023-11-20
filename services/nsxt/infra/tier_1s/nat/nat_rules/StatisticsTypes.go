// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Statistics.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nat_rules

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``action`` of method Statistics#list.
const Statistics_LIST_ACTION_AGGREGATE = "aggregate"

func statisticsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["container_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["container_cluster_path"] = "ContainerClusterPath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatisticsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.PolicyNatRuleStatisticsListResultBindingType)
}

func statisticsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["nat_id"] = vapiBindings_.NewStringType()
	fields["nat_rule_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["container_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["nat_id"] = "NatId"
	fieldNameMap["nat_rule_id"] = "NatRuleId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["container_cluster_path"] = "ContainerClusterPath"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["container_cluster_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["nat_rule_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["nat_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["natId"] = vapiBindings_.NewStringType()
	paramsTypeMap["natRuleId"] = vapiBindings_.NewStringType()
	pathParams["nat_rule_id"] = "natRuleId"
	pathParams["nat_id"] = "natId"
	pathParams["tier1_id"] = "tier1Id"
	queryParams["container_cluster_path"] = "container_cluster_path"
	queryParams["action"] = "action"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/nat/{natId}/nat-rules/{natRuleId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
