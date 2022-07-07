// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: BfdPeers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package static_routes

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func bfdPeersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["static_hop_bfd_peer"] = bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["static_hop_bfd_peer"] = "StaticHopBfdPeer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bfdPeersCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
}

func bfdPeersCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["static_hop_bfd_peer"] = bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["static_hop_bfd_peer"] = "StaticHopBfdPeer"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["static_hop_bfd_peer"] = bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
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
		"static_hop_bfd_peer",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/bfd-peers",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bfdPeersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bfd_peer_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bfd_peer_id"] = "BfdPeerId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bfdPeersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func bfdPeersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bfd_peer_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bfd_peer_id"] = "BfdPeerId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["bfd_peer_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["bfdPeerId"] = bindings.NewStringType()
	pathParams["bfd_peer_id"] = "bfdPeerId"
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/bfd-peers/{bfdPeerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bfdPeersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bfd_peer_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bfd_peer_id"] = "BfdPeerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bfdPeersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
}

func bfdPeersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bfd_peer_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bfd_peer_id"] = "BfdPeerId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["bfd_peer_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["bfdPeerId"] = bindings.NewStringType()
	pathParams["bfd_peer_id"] = "bfdPeerId"
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
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/bfd-peers/{bfdPeerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bfdPeersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bfdPeersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticHopBfdPeerListResultBindingType)
}

func bfdPeersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
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
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/bfd-peers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bfdPeersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bfd_peer_id"] = bindings.NewStringType()
	fields["static_hop_bfd_peer"] = bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bfd_peer_id"] = "BfdPeerId"
	fieldNameMap["static_hop_bfd_peer"] = "StaticHopBfdPeer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bfdPeersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
}

func bfdPeersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bfd_peer_id"] = bindings.NewStringType()
	fields["static_hop_bfd_peer"] = bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bfd_peer_id"] = "BfdPeerId"
	fieldNameMap["static_hop_bfd_peer"] = "StaticHopBfdPeer"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["static_hop_bfd_peer"] = bindings.NewReferenceType(model.StaticHopBfdPeerBindingType)
	paramsTypeMap["bfd_peer_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["bfdPeerId"] = bindings.NewStringType()
	pathParams["bfd_peer_id"] = "bfdPeerId"
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
		"static_hop_bfd_peer",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/bfd-peers/{bfdPeerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
