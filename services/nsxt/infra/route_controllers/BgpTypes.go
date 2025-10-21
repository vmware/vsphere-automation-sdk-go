// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Bgp.
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

func bgpDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BgpDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bgpDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	pathParams["router_controller_id"] = "routerControllerId"
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
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bgpGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BgpGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
}

func bgpGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	pathParams["router_controller_id"] = "routerControllerId"
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
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bgpPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_routing_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["route_controller_bgp_routing_config"] = "RouteControllerBgpRoutingConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BgpPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bgpPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_routing_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["route_controller_bgp_routing_config"] = "RouteControllerBgpRoutingConfig"
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["route_controller_bgp_routing_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	pathParams["router_controller_id"] = "routerControllerId"
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
		"route_controller_bgp_routing_config",
		"PATCH",
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bgpUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_routing_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["route_controller_bgp_routing_config"] = "RouteControllerBgpRoutingConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BgpUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
}

func bgpUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_routing_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["route_controller_bgp_routing_config"] = "RouteControllerBgpRoutingConfig"
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["route_controller_bgp_routing_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpRoutingConfigBindingType)
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	pathParams["router_controller_id"] = "routerControllerId"
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
		"route_controller_bgp_routing_config",
		"PUT",
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
