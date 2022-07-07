// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: IpPools.
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

func ipPoolsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["dhcp_ip_pool"] = bindings.NewReferenceType(model.DhcpIpPoolBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["dhcp_ip_pool"] = "DhcpIpPool"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpIpPoolBindingType)
}

func ipPoolsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["dhcp_ip_pool"] = bindings.NewReferenceType(model.DhcpIpPoolBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["dhcp_ip_pool"] = "DhcpIpPool"
	paramsTypeMap["dhcp_ip_pool"] = bindings.NewReferenceType(model.DhcpIpPoolBindingType)
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
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
		"dhcp_ip_pool",
		"POST",
		"/api/v1/dhcp/servers/{serverId}/ip-pools",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipPoolsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["pool_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["pool_id"] = "PoolId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipPoolsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["pool_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["pool_id"] = "PoolId"
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["pool_id"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	paramsTypeMap["poolId"] = bindings.NewStringType()
	pathParams["pool_id"] = "poolId"
	pathParams["server_id"] = "serverId"
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
		"/api/v1/dhcp/servers/{serverId}/ip-pools/{poolId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipPoolsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["pool_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["pool_id"] = "PoolId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpIpPoolBindingType)
}

func ipPoolsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["pool_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["pool_id"] = "PoolId"
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["pool_id"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	paramsTypeMap["poolId"] = bindings.NewStringType()
	pathParams["pool_id"] = "poolId"
	pathParams["server_id"] = "serverId"
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
		"/api/v1/dhcp/servers/{serverId}/ip-pools/{poolId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipPoolsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpIpPoolListResultBindingType)
}

func ipPoolsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["cursor"] = "cursor"
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
		"/api/v1/dhcp/servers/{serverId}/ip-pools",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipPoolsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fields["pool_id"] = bindings.NewStringType()
	fields["dhcp_ip_pool"] = bindings.NewReferenceType(model.DhcpIpPoolBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["pool_id"] = "PoolId"
	fieldNameMap["dhcp_ip_pool"] = "DhcpIpPool"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpIpPoolBindingType)
}

func ipPoolsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fields["pool_id"] = bindings.NewStringType()
	fields["dhcp_ip_pool"] = bindings.NewReferenceType(model.DhcpIpPoolBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["pool_id"] = "PoolId"
	fieldNameMap["dhcp_ip_pool"] = "DhcpIpPool"
	paramsTypeMap["dhcp_ip_pool"] = bindings.NewReferenceType(model.DhcpIpPoolBindingType)
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["pool_id"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	paramsTypeMap["poolId"] = bindings.NewStringType()
	pathParams["pool_id"] = "poolId"
	pathParams["server_id"] = "serverId"
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
		"dhcp_ip_pool",
		"PUT",
		"/api/v1/dhcp/servers/{serverId}/ip-pools/{poolId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
