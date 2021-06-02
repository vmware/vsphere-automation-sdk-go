// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches.
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

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchMacTable.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_MAC_TABLE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchMacTable.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_MAC_TABLE_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchMacTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_MAC_TABLE_IN_CSV_FORMAT_CSV_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchMacTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_MAC_TABLE_IN_CSV_FORMAT_CSV_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchStatistics.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchStatistics.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_STATISTICS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_STATUS_SUMMARY_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_STATUS_SUMMARY_SOURCE_CACHED = "cached"

// Possible value for ``transportType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_STATUS_SUMMARY_TRANSPORT_TYPE_OVERLAY = "OVERLAY"

// Possible value for ``transportType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_STATUS_SUMMARY_TRANSPORT_TYPE_VLAN = "VLAN"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchVtepTable.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_VTEP_TABLE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchVtepTable.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_VTEP_TABLE_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchVtepTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_VTEP_TABLE_IN_CSV_FORMAT_CSV_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#getLogicalSwitchVtepTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_GET_LOGICAL_SWITCH_VTEP_TABLE_IN_CSV_FORMAT_CSV_SOURCE_CACHED = "cached"

// Possible value for ``transportType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#listLogicalSwitches.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_LIST_LOGICAL_SWITCHES_TRANSPORT_TYPE_OVERLAY = "OVERLAY"

// Possible value for ``transportType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#listLogicalSwitches.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_LIST_LOGICAL_SWITCHES_TRANSPORT_TYPE_VLAN = "VLAN"

// Possible value for ``status`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#listLogicalSwitchesByState.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_LIST_LOGICAL_SWITCHES_BY_STATE_STATUS_PENDING = "PENDING"

// Possible value for ``status`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#listLogicalSwitchesByState.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_LIST_LOGICAL_SWITCHES_BY_STATE_STATUS_IN_PROGRESS = "IN_PROGRESS"

// Possible value for ``status`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#listLogicalSwitchesByState.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_LIST_LOGICAL_SWITCHES_BY_STATE_STATUS_PARTIAL_SUCCESS = "PARTIAL_SUCCESS"

// Possible value for ``status`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches#listLogicalSwitchesByState.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches_LIST_LOGICAL_SWITCHES_BY_STATE_STATUS_SUCCESS = "SUCCESS"

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesCreateLogicalSwitchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesCreateLogicalSwitchOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesCreateLogicalSwitchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	paramsTypeMap["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
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
		"logical_switch",
		"POST",
		"/api/v1/logical-switches",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesDeleteLogicalSwitchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cascade"] = "Cascade"
	fieldNameMap["detach"] = "Detach"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesDeleteLogicalSwitchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesDeleteLogicalSwitchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cascade"] = "Cascade"
	fieldNameMap["detach"] = "Detach"
	paramsTypeMap["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
	queryParams["cascade"] = "cascade"
	queryParams["detach"] = "detach"
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
		"/api/v1/logical-switches/{lswitchId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"/api/v1/logical-switches/{lswitchId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MacAddressListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"/api/v1/logical-switches/{lswitchId}/mac-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInCsvFormatCsvInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInCsvFormatCsvOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MacAddressCsvListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInCsvFormatCsvRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"format=csv",
		"",
		"GET",
		"/api/v1/logical-switches/{lswitchId}/mac-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchStateBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"/api/v1/logical-switches/{lswitchId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchStatisticsBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"/api/v1/logical-switches/{lswitchId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchStatusBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fieldNameMap["lswitch_id"] = "LswitchId"
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"/api/v1/logical-switches/{lswitchId}/summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusSummaryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusSummaryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchStatusSummaryBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusSummaryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	paramsTypeMap["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["switching_profile_id"] = "switching_profile_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["uplink_teaming_policy_name"] = "uplink_teaming_policy_name"
	queryParams["vni"] = "vni"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["vlan"] = "vlan"
	queryParams["included_fields"] = "included_fields"
	queryParams["transport_type"] = "transport_type"
	queryParams["transport_zone_id"] = "transport_zone_id"
	queryParams["diagnostic"] = "diagnostic"
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
		"/api/v1/logical-switches/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VtepListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"/api/v1/logical-switches/{lswitchId}/vtep-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInCsvFormatCsvInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInCsvFormatCsvOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VtepCsvListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInCsvFormatCsvRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"format=csv",
		"",
		"GET",
		"/api/v1/logical-switches/{lswitchId}/vtep-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	paramsTypeMap["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transport_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vni"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["vlan"] = bindings.NewOptionalType(bindings.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["uplink_teaming_policy_name"] = "uplink_teaming_policy_name"
	queryParams["vni"] = "vni"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["vlan"] = "vlan"
	queryParams["included_fields"] = "included_fields"
	queryParams["switching_profile_id"] = "switching_profile_id"
	queryParams["transport_type"] = "transport_type"
	queryParams["transport_zone_id"] = "transport_zone_id"
	queryParams["diagnostic"] = "diagnostic"
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
		"/api/v1/logical-switches",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesByStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesByStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchStateListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesByStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["status"] = "Status"
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["status"] = "status"
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
		"/api/v1/logical-switches/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesUpdateLogicalSwitchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lswitch_id"] = bindings.NewStringType()
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesUpdateLogicalSwitchOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesUpdateLogicalSwitchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lswitch_id"] = bindings.NewStringType()
	fields["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	fieldNameMap["lswitch_id"] = "LswitchId"
	fieldNameMap["logical_switch"] = "LogicalSwitch"
	paramsTypeMap["lswitch_id"] = bindings.NewStringType()
	paramsTypeMap["logical_switch"] = bindings.NewReferenceType(model.LogicalSwitchBindingType)
	paramsTypeMap["lswitchId"] = bindings.NewStringType()
	pathParams["lswitch_id"] = "lswitchId"
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
		"logical_switch",
		"PUT",
		"/api/v1/logical-switches/{lswitchId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
