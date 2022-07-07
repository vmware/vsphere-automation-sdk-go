// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Leases.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package servers

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Leases#list.
const Leases_LIST_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Leases#list.
const Leases_LIST_SOURCE_CACHED = "cached"

func leasesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["ip"] = bindings.NewStringType()
	fields["mac"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["ip"] = "Ip"
	fieldNameMap["mac"] = "Mac"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func leasesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func leasesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["ip"] = bindings.NewStringType()
	fields["mac"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["ip"] = "Ip"
	fieldNameMap["mac"] = "Mac"
	paramsTypeMap["mac"] = bindings.NewStringType()
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["ip"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["ip"] = "ip"
	queryParams["mac"] = "mac"
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
		"/api/v1/dhcp/servers/{serverId}/leases",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func leasesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["pool_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["pool_id"] = "PoolId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func leasesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpLeasesBindingType)
}

func leasesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["pool_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["pool_id"] = "PoolId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["pool_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["address"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["address"] = "address"
	queryParams["source"] = "source"
	queryParams["pool_id"] = "pool_id"
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
		"/api/v1/dhcp/servers/{serverId}/leases",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
