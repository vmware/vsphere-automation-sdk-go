// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Neighbors.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package bgp

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func neighborsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NeighborsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func neighborsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	paramsTypeMap["neighbor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	paramsTypeMap["neighborId"] = vapiBindings_.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["router_controller_id"] = "routerControllerId"
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
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp/neighbors/{neighborId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func neighborsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NeighborsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
}

func neighborsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	paramsTypeMap["neighbor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	paramsTypeMap["neighborId"] = vapiBindings_.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["router_controller_id"] = "routerControllerId"
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
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp/neighbors/{neighborId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func neighborsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_conflicts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_conflicts"] = "IncludeConflicts"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NeighborsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigListResultBindingType)
}

func neighborsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_conflicts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_conflicts"] = "IncludeConflicts"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_conflicts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	pathParams["router_controller_id"] = "routerControllerId"
	queryParams["cursor"] = "cursor"
	queryParams["include_conflicts"] = "include_conflicts"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp/neighbors",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func neighborsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_neighbor_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	fieldNameMap["route_controller_bgp_neighbor_config"] = "RouteControllerBgpNeighborConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NeighborsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func neighborsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_neighbor_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	fieldNameMap["route_controller_bgp_neighbor_config"] = "RouteControllerBgpNeighborConfig"
	paramsTypeMap["neighbor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["route_controller_bgp_neighbor_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	paramsTypeMap["neighborId"] = vapiBindings_.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["router_controller_id"] = "routerControllerId"
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
		"route_controller_bgp_neighbor_config",
		"PATCH",
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp/neighbors/{neighborId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func neighborsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_neighbor_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	fieldNameMap["route_controller_bgp_neighbor_config"] = "RouteControllerBgpNeighborConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NeighborsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
}

func neighborsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["router_controller_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fields["route_controller_bgp_neighbor_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
	fieldNameMap["router_controller_id"] = "RouterControllerId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	fieldNameMap["route_controller_bgp_neighbor_config"] = "RouteControllerBgpNeighborConfig"
	paramsTypeMap["neighbor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["route_controller_bgp_neighbor_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.RouteControllerBgpNeighborConfigBindingType)
	paramsTypeMap["router_controller_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["routerControllerId"] = vapiBindings_.NewStringType()
	paramsTypeMap["neighborId"] = vapiBindings_.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["router_controller_id"] = "routerControllerId"
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
		"route_controller_bgp_neighbor_config",
		"PUT",
		"/policy/api/v1/infra/route-controllers/{routerControllerId}/bgp/neighbors/{neighborId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
