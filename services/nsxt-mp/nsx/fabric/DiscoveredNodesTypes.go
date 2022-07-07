// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DiscoveredNodes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package fabric

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``hasParent`` of method DiscoveredNodes#list.
const DiscoveredNodes_LIST_HAS_PARENT_TRUE = "true"

// Possible value for ``hasParent`` of method DiscoveredNodes#list.
const DiscoveredNodes_LIST_HAS_PARENT_FALSE = "false"

func discoveredNodesCreatetransportnodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_ext_id"] = bindings.NewStringType()
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fieldNameMap["node_ext_id"] = "NodeExtId"
	fieldNameMap["transport_node"] = "TransportNode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func discoveredNodesCreatetransportnodeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeBindingType)
}

func discoveredNodesCreatetransportnodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_ext_id"] = bindings.NewStringType()
	fields["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	fieldNameMap["node_ext_id"] = "NodeExtId"
	fieldNameMap["transport_node"] = "TransportNode"
	paramsTypeMap["transport_node"] = bindings.NewReferenceType(model.TransportNodeBindingType)
	paramsTypeMap["node_ext_id"] = bindings.NewStringType()
	paramsTypeMap["nodeExtId"] = bindings.NewStringType()
	pathParams["node_ext_id"] = "nodeExtId"
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
		"action=create_transport_node",
		"transport_node",
		"POST",
		"/api/v1/fabric/discovered-nodes/{nodeExtId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func discoveredNodesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_ext_id"] = bindings.NewStringType()
	fieldNameMap["node_ext_id"] = "NodeExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func discoveredNodesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DiscoveredNodeBindingType)
}

func discoveredNodesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_ext_id"] = bindings.NewStringType()
	fieldNameMap["node_ext_id"] = "NodeExtId"
	paramsTypeMap["node_ext_id"] = bindings.NewStringType()
	paramsTypeMap["nodeExtId"] = bindings.NewStringType()
	pathParams["node_ext_id"] = "nodeExtId"
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
		"/api/v1/fabric/discovered-nodes/{nodeExtId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func discoveredNodesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cm_local_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["has_parent"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["origin_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["parent_compute_collection"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cm_local_id"] = "CmLocalId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["has_parent"] = "HasParent"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ip_address"] = "IpAddress"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_type"] = "NodeType"
	fieldNameMap["origin_id"] = "OriginId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["parent_compute_collection"] = "ParentComputeCollection"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func discoveredNodesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DiscoveredNodeListResultBindingType)
}

func discoveredNodesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cm_local_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["has_parent"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["origin_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["parent_compute_collection"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cm_local_id"] = "CmLocalId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["has_parent"] = "HasParent"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["ip_address"] = "IpAddress"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_type"] = "NodeType"
	fieldNameMap["origin_id"] = "OriginId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["parent_compute_collection"] = "ParentComputeCollection"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["origin_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["parent_compute_collection"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cm_local_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["has_parent"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["cm_local_id"] = "cm_local_id"
	queryParams["has_parent"] = "has_parent"
	queryParams["external_id"] = "external_id"
	queryParams["origin_id"] = "origin_id"
	queryParams["ip_address"] = "ip_address"
	queryParams["sort_by"] = "sort_by"
	queryParams["display_name"] = "display_name"
	queryParams["node_type"] = "node_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["parent_compute_collection"] = "parent_compute_collection"
	queryParams["node_id"] = "node_id"
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
		"/api/v1/fabric/discovered-nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func discoveredNodesReapplyclusterconfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_ext_id"] = bindings.NewStringType()
	fieldNameMap["node_ext_id"] = "NodeExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func discoveredNodesReapplyclusterconfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeBindingType)
}

func discoveredNodesReapplyclusterconfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_ext_id"] = bindings.NewStringType()
	fieldNameMap["node_ext_id"] = "NodeExtId"
	paramsTypeMap["node_ext_id"] = bindings.NewStringType()
	paramsTypeMap["nodeExtId"] = bindings.NewStringType()
	pathParams["node_ext_id"] = "nodeExtId"
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
		"action=reapply_cluster_config",
		"",
		"POST",
		"/api/v1/fabric/discovered-nodes/{nodeExtId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
