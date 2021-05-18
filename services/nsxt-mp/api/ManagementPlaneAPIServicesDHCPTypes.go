// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPIServicesDHCP.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``source`` of method ManagementPlaneAPIServicesDHCP#getDhcpLeaseInfo.
const ManagementPlaneAPIServicesDHCP_GET_DHCP_LEASE_INFO_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPIServicesDHCP#getDhcpLeaseInfo.
const ManagementPlaneAPIServicesDHCP_GET_DHCP_LEASE_INFO_SOURCE_CACHED = "cached"

func managementPlaneAPIServicesDHCPGetDhcpLeaseInfoInputType() bindings.StructType {
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

func managementPlaneAPIServicesDHCPGetDhcpLeaseInfoOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpLeasesBindingType)
}

func managementPlaneAPIServicesDHCPGetDhcpLeaseInfoRestMetadata() protocol.OperationRestMetadata {
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

func managementPlaneAPIServicesDHCPGetDhcpStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPIServicesDHCPGetDhcpStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpStatisticsBindingType)
}

func managementPlaneAPIServicesDHCPGetDhcpStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = bindings.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
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
		"",
		"GET",
		"/api/v1/dhcp/servers/{serverId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
