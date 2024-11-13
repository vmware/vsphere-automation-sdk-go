// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RelocateAndRemoveEdgeTransportNode.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package action

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func relocateAndRemoveEdgeTransportNodeCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_id"] = vapiBindings_.NewStringType()
	fields["policy_edge_node_relocate_and_remove_member"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeNodeRelocateAndRemoveMemberBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["policy_edge_node_relocate_and_remove_member"] = "PolicyEdgeNodeRelocateAndRemoveMember"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RelocateAndRemoveEdgeTransportNodeCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func relocateAndRemoveEdgeTransportNodeCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["edge_cluster_id"] = vapiBindings_.NewStringType()
	fields["policy_edge_node_relocate_and_remove_member"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeNodeRelocateAndRemoveMemberBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["policy_edge_node_relocate_and_remove_member"] = "PolicyEdgeNodeRelocateAndRemoveMember"
	paramsTypeMap["edge_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_edge_node_relocate_and_remove_member"] = vapiBindings_.NewReferenceType(nsx_policyModel.PolicyEdgeNodeRelocateAndRemoveMemberBindingType)
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["edgeClusterId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"policy_edge_node_relocate_and_remove_member",
		"POST",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/edge-clusters/{edgeClusterId}/action/relocate-and-remove-edge-transport-node",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
