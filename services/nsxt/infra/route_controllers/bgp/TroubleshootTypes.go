// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Troubleshoot.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package bgp

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func troubleshootDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func troubleshootDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	paramsTypeMap["route_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routeControllerId"] = vapiBindings_.NewStringType()
	pathParams["route_controller_id"] = "routeControllerId"
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
		"/policy/api/v1/infra/route-controllers/{routeControllerId}/bgp/troubleshoot",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func troubleshootGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
}

func troubleshootGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	paramsTypeMap["route_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routeControllerId"] = vapiBindings_.NewStringType()
	pathParams["route_controller_id"] = "routeControllerId"
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
		"/policy/api/v1/infra/route-controllers/{routeControllerId}/bgp/troubleshoot",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func troubleshootPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fields["bgp_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	fieldNameMap["bgp_troubleshoot_config"] = "BgpTroubleshootConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func troubleshootPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fields["bgp_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	fieldNameMap["bgp_troubleshoot_config"] = "BgpTroubleshootConfig"
	paramsTypeMap["bgp_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
	paramsTypeMap["route_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routeControllerId"] = vapiBindings_.NewStringType()
	pathParams["route_controller_id"] = "routeControllerId"
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
		"bgp_troubleshoot_config",
		"PATCH",
		"/policy/api/v1/infra/route-controllers/{routeControllerId}/bgp/troubleshoot",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func troubleshootUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fields["bgp_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	fieldNameMap["bgp_troubleshoot_config"] = "BgpTroubleshootConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
}

func troubleshootUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["route_controller_id"] = vapiBindings_.NewStringType()
	fields["bgp_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
	fieldNameMap["route_controller_id"] = "RouteControllerId"
	fieldNameMap["bgp_troubleshoot_config"] = "BgpTroubleshootConfig"
	paramsTypeMap["bgp_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.BgpTroubleshootConfigBindingType)
	paramsTypeMap["route_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routeControllerId"] = vapiBindings_.NewStringType()
	pathParams["route_controller_id"] = "routeControllerId"
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
		"bgp_troubleshoot_config",
		"PUT",
		"/policy/api/v1/infra/route-controllers/{routeControllerId}/bgp/troubleshoot",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
