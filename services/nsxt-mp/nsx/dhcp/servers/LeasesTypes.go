// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Leases.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package servers

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Leases#list.
const Leases_LIST_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Leases#list.
const Leases_LIST_SOURCE_CACHED = "cached"

func leasesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["ip"] = vapiBindings_.NewStringType()
	fields["mac"] = vapiBindings_.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["ip"] = "Ip"
	fieldNameMap["mac"] = "Mac"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LeasesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func leasesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["ip"] = vapiBindings_.NewStringType()
	fields["mac"] = vapiBindings_.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["ip"] = "Ip"
	fieldNameMap["mac"] = "Mac"
	paramsTypeMap["ip"] = vapiBindings_.NewStringType()
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["mac"] = vapiBindings_.NewStringType()
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["ip"] = "ip"
	queryParams["mac"] = "mac"
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
		"/api/v1/dhcp/servers/{serverId}/leases",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func leasesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["pool_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["pool_id"] = "PoolId"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LeasesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.DhcpLeasesBindingType)
}

func leasesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["pool_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["pool_id"] = "PoolId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["address"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["pool_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["address"] = "address"
	queryParams["source"] = "source"
	queryParams["pool_id"] = "pool_id"
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
		"/api/v1/dhcp/servers/{serverId}/leases",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
