// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CreateGroupNetworkConnectivity.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package network_connectivity_configs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
	"reflect"
)

func createGroupNetworkConnectivityCreateGroupNetworkConnectivityInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["create_group_network_connectivity_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityRequestBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["create_group_network_connectivity_request"] = "CreateGroupNetworkConnectivityRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CreateGroupNetworkConnectivityCreateGroupNetworkConnectivityOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityResponseBindingType)
}

func createGroupNetworkConnectivityCreateGroupNetworkConnectivityRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["create_group_network_connectivity_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityRequestBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["create_group_network_connectivity_request"] = "CreateGroupNetworkConnectivityRequest"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["create_group_network_connectivity_request"] = vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityRequestBindingType)
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"create_group_network_connectivity_request",
		"POST",
		"/api/network/{orgId}/core/network-connectivity-configs/create-group-network-connectivity",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.concurrent_change": 409})
}
