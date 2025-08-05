// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Redeploy.
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

func redeployCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcementpoint_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_appliance_cluster_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_appliance_id"] = vapiBindings_.NewStringType()
	fields["base_network_appliance"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_policyModel.BaseNetworkApplianceBindingType)})
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["virtual_network_appliance_cluster_id"] = "VirtualNetworkApplianceClusterId"
	fieldNameMap["virtual_network_appliance_id"] = "VirtualNetworkApplianceId"
	fieldNameMap["base_network_appliance"] = "BaseNetworkAppliance"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RedeployCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_policyModel.BaseNetworkApplianceBindingType)})
}

func redeployCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["virtual_network_appliance_cluster_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_appliance_id"] = vapiBindings_.NewStringType()
	fields["base_network_appliance"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_policyModel.BaseNetworkApplianceBindingType)})
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["virtual_network_appliance_cluster_id"] = "VirtualNetworkApplianceClusterId"
	fieldNameMap["virtual_network_appliance_id"] = "VirtualNetworkApplianceId"
	fieldNameMap["base_network_appliance"] = "BaseNetworkAppliance"
	paramsTypeMap["enforcementpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtual_network_appliance_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtual_network_appliance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["base_network_appliance"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_policyModel.BaseNetworkApplianceBindingType)})
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementpointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtualNetworkApplianceClusterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["virtualNetworkApplianceId"] = vapiBindings_.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["virtual_network_appliance_id"] = "virtualNetworkApplianceId"
	pathParams["virtual_network_appliance_cluster_id"] = "virtualNetworkApplianceClusterId"
	pathParams["site_id"] = "siteId"
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
		"base_network_appliance",
		"POST",
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/virtual-network-appliance-clusters/{virtualNetworkApplianceClusterId}/virtual-network-appliances/{virtualNetworkApplianceId}/action/redeploy",
		"application/json",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
