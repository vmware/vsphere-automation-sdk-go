// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Interfaces.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package network

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Interfaces#get.
const Interfaces_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Interfaces#get.
const Interfaces_GET_SOURCE_CACHED = "cached"

// Possible value for ``adminStatus`` of method Interfaces#list.
const Interfaces_LIST_ADMIN_STATUS_UP = "UP"

// Possible value for ``adminStatus`` of method Interfaces#list.
const Interfaces_LIST_ADMIN_STATUS_DOWN = "DOWN"

// Possible value for ``source`` of method Interfaces#list.
const Interfaces_LIST_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Interfaces#list.
const Interfaces_LIST_SOURCE_CACHED = "cached"

func interfacesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeInterfacePropertiesBindingType)
}

func interfacesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fields["interface_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["interface_id"] = bindings.NewStringType()
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	paramsTypeMap["interfaceId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	pathParams["interface_id"] = "interfaceId"
	queryParams["source"] = "source"
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
		"/api/v1/transport-nodes/{transportNodeId}/network/interfaces/{interfaceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func interfacesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fields["admin_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["admin_status"] = "AdminStatus"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeInterfacePropertiesListResultBindingType)
}

func interfacesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fields["admin_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["admin_status"] = "AdminStatus"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["admin_status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	queryParams["admin_status"] = "admin_status"
	queryParams["source"] = "source"
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
		"/api/v1/transport-nodes/{transportNodeId}/network/interfaces",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
