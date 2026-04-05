// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: BgpRouteTable.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package route_controllers

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func bgpRouteTableListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_appliance_path"] = vapiBindings_.NewStringType()
	fields["network_prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	fieldNameMap["virtual_network_appliance_path"] = "VirtualNetworkAppliancePath"
	fieldNameMap["network_prefix"] = "NetworkPrefix"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BgpRouteTableListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BgpRIBListResultBindingType)
}

func bgpRouteTableListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fields["virtual_network_appliance_path"] = vapiBindings_.NewStringType()
	fields["network_prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	fieldNameMap["virtual_network_appliance_path"] = "VirtualNetworkAppliancePath"
	fieldNameMap["network_prefix"] = "NetworkPrefix"
	paramsTypeMap["virtual_network_appliance_path"] = vapiBindings_.NewStringType()
	paramsTypeMap["network_prefix"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["route_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routeControllerId"] = vapiBindings_.NewStringType()
	pathParams["route_controller_id"] = "routeControllerId"
	queryParams["virtual_network_appliance_path"] = "virtual_network_appliance_path"
	queryParams["network_prefix"] = "network_prefix"
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
		"/policy/api/v1/infra/route-controllers/{routeControllerId}/bgp-route-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
