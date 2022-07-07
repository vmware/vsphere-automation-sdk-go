// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Bgp.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package routing

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func bgpGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bgpGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpConfigBindingType)
}

func bgpGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
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
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bgpUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bgp_config"] = bindings.NewReferenceType(model.BgpConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bgp_config"] = "BgpConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bgpUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpConfigBindingType)
}

func bgpUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bgp_config"] = bindings.NewReferenceType(model.BgpConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bgp_config"] = "BgpConfig"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["bgp_config"] = bindings.NewReferenceType(model.BgpConfigBindingType)
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
		"bgp_config",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
