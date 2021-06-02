// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingVPNIPSECPeerEndpoints.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNLocalEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_sec_VPN_local_endpoint"] = bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
	fieldNameMap["ip_sec_VPN_local_endpoint"] = "IpSecVPNLocalEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNLocalEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNLocalEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_sec_VPN_local_endpoint"] = bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
	fieldNameMap["ip_sec_VPN_local_endpoint"] = "IpSecVPNLocalEndpoint"
	paramsTypeMap["ip_sec_VPN_local_endpoint"] = bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
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
		"ip_sec_VPN_local_endpoint",
		"POST",
		"/api/v1/vpn/ipsec/local-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNPeerEndPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_sec_VPN_peer_endpoint"] = bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
	fieldNameMap["ip_sec_VPN_peer_endpoint"] = "IpSecVPNPeerEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNPeerEndPointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsCreateIPSecVPNPeerEndPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_sec_VPN_peer_endpoint"] = bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
	fieldNameMap["ip_sec_VPN_peer_endpoint"] = "IpSecVPNPeerEndpoint"
	paramsTypeMap["ip_sec_VPN_peer_endpoint"] = bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
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
		"ip_sec_VPN_peer_endpoint",
		"POST",
		"/api/v1/vpn/ipsec/peer-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNLocalEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_local_endpoint_id"] = "IpsecVpnLocalEndpointId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNLocalEndpointOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNLocalEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_local_endpoint_id"] = "IpsecVpnLocalEndpointId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnLocalEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_local_endpoint_id"] = "ipsecVpnLocalEndpointId"
	queryParams["force"] = "force"
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
		"/api/v1/vpn/ipsec/local-endpoints/{ipsecVpnLocalEndpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNPeerEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNPeerEndpointOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsDeleteIPSecVPNPeerEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnPeerEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_peer_endpoint_id"] = "ipsecVpnPeerEndpointId"
	queryParams["force"] = "force"
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
		"/api/v1/vpn/ipsec/peer-endpoints/{ipsecVpnPeerEndpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNLocalEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_local_endpoint_id"] = "IpsecVpnLocalEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNLocalEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNLocalEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_local_endpoint_id"] = "IpsecVpnLocalEndpointId"
	paramsTypeMap["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnLocalEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_local_endpoint_id"] = "ipsecVpnLocalEndpointId"
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
		"/api/v1/vpn/ipsec/local-endpoints/{ipsecVpnLocalEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	paramsTypeMap["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnPeerEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_peer_endpoint_id"] = "ipsecVpnPeerEndpointId"
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
		"/api/v1/vpn/ipsec/peer-endpoints/{ipsecVpnPeerEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointWithPSKShowSensitiveDataInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointWithPSKShowSensitiveDataOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsGetIPSecVPNPeerEndpointWithPSKShowSensitiveDataRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	paramsTypeMap["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnPeerEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_peer_endpoint_id"] = "ipsecVpnPeerEndpointId"
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
		"action=show-sensitive-data",
		"",
		"GET",
		"/api/v1/vpn/ipsec/peer-endpoints/{ipsecVpnPeerEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNLocalEndpointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ipsec_vpn_service_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ipsec_vpn_service_id"] = "IpsecVpnServiceId"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNLocalEndpointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNLocalEndpointListResultBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNLocalEndpointsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ipsec_vpn_service_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ipsec_vpn_service_id"] = "IpsecVpnServiceId"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["ipsec_vpn_service_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["logical_router_id"] = "logical_router_id"
	queryParams["ipsec_vpn_service_id"] = "ipsec_vpn_service_id"
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
		"/api/v1/vpn/ipsec/local-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNPeerEndpointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNPeerEndpointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNPeerEndpointListResultBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsListIPSecVPNPeerEndpointsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
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
		"/api/v1/vpn/ipsec/peer-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNLocalEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	fields["ip_sec_VPN_local_endpoint"] = bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
	fieldNameMap["ipsec_vpn_local_endpoint_id"] = "IpsecVpnLocalEndpointId"
	fieldNameMap["ip_sec_VPN_local_endpoint"] = "IpSecVPNLocalEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNLocalEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNLocalEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	fields["ip_sec_VPN_local_endpoint"] = bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
	fieldNameMap["ipsec_vpn_local_endpoint_id"] = "IpsecVpnLocalEndpointId"
	fieldNameMap["ip_sec_VPN_local_endpoint"] = "IpSecVPNLocalEndpoint"
	paramsTypeMap["ip_sec_VPN_local_endpoint"] = bindings.NewReferenceType(model.IPSecVPNLocalEndpointBindingType)
	paramsTypeMap["ipsec_vpn_local_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnLocalEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_local_endpoint_id"] = "ipsecVpnLocalEndpointId"
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
		"ip_sec_VPN_local_endpoint",
		"PUT",
		"/api/v1/vpn/ipsec/local-endpoints/{ipsecVpnLocalEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNPeerEndpointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fields["ip_sec_VPN_peer_endpoint"] = bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	fieldNameMap["ip_sec_VPN_peer_endpoint"] = "IpSecVPNPeerEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNPeerEndpointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
}

func managementPlaneAPINetworkingVPNIPSECPeerEndpointsUpdateIPSecVPNPeerEndpointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	fields["ip_sec_VPN_peer_endpoint"] = bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
	fieldNameMap["ipsec_vpn_peer_endpoint_id"] = "IpsecVpnPeerEndpointId"
	fieldNameMap["ip_sec_VPN_peer_endpoint"] = "IpSecVPNPeerEndpoint"
	paramsTypeMap["ip_sec_VPN_peer_endpoint"] = bindings.NewReferenceType(model.IPSecVPNPeerEndpointBindingType)
	paramsTypeMap["ipsec_vpn_peer_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["ipsecVpnPeerEndpointId"] = bindings.NewStringType()
	pathParams["ipsec_vpn_peer_endpoint_id"] = "ipsecVpnPeerEndpointId"
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
		"ip_sec_VPN_peer_endpoint",
		"PUT",
		"/api/v1/vpn/ipsec/peer-endpoints/{ipsecVpnPeerEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
