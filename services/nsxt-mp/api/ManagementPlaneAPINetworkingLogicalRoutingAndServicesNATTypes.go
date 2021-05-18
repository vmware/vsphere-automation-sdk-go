// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#getNatStatisticsPerLogicalRouter.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_GET_NAT_STATISTICS_PER_LOGICAL_ROUTER_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#getNatStatisticsPerLogicalRouter.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_GET_NAT_STATISTICS_PER_LOGICAL_ROUTER_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#getNatStatisticsPerRule.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_GET_NAT_STATISTICS_PER_RULE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#getNatStatisticsPerRule.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_GET_NAT_STATISTICS_PER_RULE_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#getNatStatisticsPerTransportNode.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_GET_NAT_STATISTICS_PER_TRANSPORT_NODE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#getNatStatisticsPerTransportNode.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_GET_NAT_STATISTICS_PER_TRANSPORT_NODE_SOURCE_CACHED = "cached"

// Possible value for ``ruleType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#listNatRules.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_LIST_NAT_RULES_RULE_TYPE_ALL = "ALL"

// Possible value for ``ruleType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#listNatRules.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_LIST_NAT_RULES_RULE_TYPE_NATV4 = "NATv4"

// Possible value for ``ruleType`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT#listNatRules.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesNAT_LIST_NAT_RULES_RULE_TYPE_NAT64 = "NAT64"

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["nat_rule"] = bindings.NewReferenceType(model.NatRuleBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["nat_rule"] = "NatRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatRuleBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["nat_rule"] = bindings.NewReferenceType(model.NatRuleBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["nat_rule"] = "NatRule"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["nat_rule"] = bindings.NewReferenceType(model.NatRuleBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"nat_rule",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRulesCreateMultipleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["nat_rule_list"] = bindings.NewReferenceType(model.NatRuleListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["nat_rule_list"] = "NatRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRulesCreateMultipleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatRuleListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATAddNatRulesCreateMultipleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["nat_rule_list"] = bindings.NewReferenceType(model.NatRuleListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["nat_rule_list"] = "NatRuleList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["nat_rule_list"] = bindings.NewReferenceType(model.NatRuleListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"nat_rule_list",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATDeleteNatRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATDeleteNatRuleOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATDeleteNatRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["rule_id"] = "ruleId"
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules/{ruleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatRuleBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["rule_id"] = "ruleId"
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerLogicalRouterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerLogicalRouterOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatStatisticsPerLogicalRouterBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerLogicalRouterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatStatisticsPerRuleBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["rule_id"] = "ruleId"
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules/{ruleId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerTransportNodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerTransportNodeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatStatisticsPerTransportNodeBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATGetNatStatisticsPerTransportNodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
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
		"/api/v1/transport-nodes/{nodeId}/statistics/nat-rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATListNatRulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["rule_type"] = "RuleType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATListNatRulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatRuleListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATListNatRulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["rule_type"] = "RuleType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["rule_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["cursor"] = "cursor"
	queryParams["rule_type"] = "rule_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
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
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATUpdateNatRuleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["nat_rule"] = bindings.NewReferenceType(model.NatRuleBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["nat_rule"] = "NatRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATUpdateNatRuleOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NatRuleBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesNATUpdateNatRuleRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["rule_id"] = bindings.NewStringType()
	fields["nat_rule"] = bindings.NewReferenceType(model.NatRuleBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["nat_rule"] = "NatRule"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["nat_rule"] = bindings.NewReferenceType(model.NatRuleBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["rule_id"] = "ruleId"
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"nat_rule",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/nat/rules/{ruleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
