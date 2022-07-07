// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: LogicalRouterPorts.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERUPLINKPORT = "LogicalRouterUpLinkPort"

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERDOWNLINKPORT = "LogicalRouterDownLinkPort"

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERLINKPORTONTIER0 = "LogicalRouterLinkPortOnTIER0"

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERLINKPORTONTIER1 = "LogicalRouterLinkPortOnTIER1"

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERLOOPBACKPORT = "LogicalRouterLoopbackPort"

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERIPTUNNELPORT = "LogicalRouterIPTunnelPort"

// Possible value for ``resourceType`` of method LogicalRouterPorts#list.
const LogicalRouterPorts_LIST_RESOURCE_TYPE_LOGICALROUTERCENTRALIZEDSERVICEPORT = "LogicalRouterCentralizedServicePort"

func logicalRouterPortsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalRouterPortsCreateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
}

func logicalRouterPortsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	paramsTypeMap["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
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
		"logical_router_port",
		"POST",
		"/api/v1/logical-router-ports",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalRouterPortsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["cascade_delete_linked_ports"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["cascade_delete_linked_ports"] = "CascadeDeleteLinkedPorts"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalRouterPortsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func logicalRouterPortsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["cascade_delete_linked_ports"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["cascade_delete_linked_ports"] = "CascadeDeleteLinkedPorts"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["cascade_delete_linked_ports"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["cascade_delete_linked_ports"] = "cascade_delete_linked_ports"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalRouterPortsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalRouterPortsGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
}

func logicalRouterPortsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalRouterPortsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalRouterPortsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalRouterPortListResultBindingType)
}

func logicalRouterPortsListRestMetadata() protocol.OperationRestMetadata {
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
	fields["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["logical_router_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["logical_router_id"] = "logical_router_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["resource_type"] = "resource_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["logical_switch_id"] = "logical_switch_id"
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
		"/api/v1/logical-router-ports",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func logicalRouterPortsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalRouterPortsUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
}

func logicalRouterPortsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = bindings.NewStringType()
	fields["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["logical_router_port"] = "LogicalRouterPort"
	paramsTypeMap["logical_router_port"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LogicalRouterPortBindingType)}, bindings.REST)
	paramsTypeMap["logical_router_port_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterPortId"] = bindings.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
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
		"logical_router_port",
		"PUT",
		"/api/v1/logical-router-ports/{logicalRouterPortId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
