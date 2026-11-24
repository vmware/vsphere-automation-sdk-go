// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: State.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package distributed_vlan_connections

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_UNKNOWN = "VALIDATION_UNKNOWN"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_IN_PROGRESS = "VALIDATION_IN_PROGRESS"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_SUCCESS = "VALIDATION_SUCCESS"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_TIMED_OUT = "VALIDATION_TIMED_OUT"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_FAILED = "VALIDATION_FAILED"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_REALIZATION_FAILURE_HOST_SWITCH_NOT_FOUND = "VALIDATION_REALIZATION_FAILURE_HOST_SWITCH_NOT_FOUND"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_REALIZATION_FAILURE_NO_CONNECTED_UPLINK = "VALIDATION_REALIZATION_FAILURE_NO_CONNECTED_UPLINK"

// Possible value for ``status`` of method State#list.
const State_LIST_STATUS_REALIZATION_FAILURE_INTERNAL_ERROR = "VALIDATION_REALIZATION_FAILURE_INTERNAL_ERROR"

func stateListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["distributed_vlan_connection_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["host_transport_node_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_node_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_span_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["zone_external_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["distributed_vlan_connection_id"] = "DistributedVlanConnectionId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["host_transport_node_path"] = "HostTransportNodePath"
	fieldNameMap["include_node_details"] = "IncludeNodeDetails"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_span_path"] = "NetworkSpanPath"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["status"] = "Status"
	fieldNameMap["zone_external_id"] = "ZoneExternalId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StateListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVlanConnectionStateBindingType)
}

func stateListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["distributed_vlan_connection_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["host_transport_node_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_node_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["network_span_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["zone_external_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["distributed_vlan_connection_id"] = "DistributedVlanConnectionId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["host_transport_node_path"] = "HostTransportNodePath"
	fieldNameMap["include_node_details"] = "IncludeNodeDetails"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_span_path"] = "NetworkSpanPath"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["status"] = "Status"
	fieldNameMap["zone_external_id"] = "ZoneExternalId"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["host_transport_node_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["zone_external_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["network_span_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["distributed_vlan_connection_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_node_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["distributedVlanConnectionId"] = vapiBindings_.NewStringType()
	pathParams["distributed_vlan_connection_id"] = "distributedVlanConnectionId"
	queryParams["cursor"] = "cursor"
	queryParams["host_transport_node_path"] = "host_transport_node_path"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["zone_external_id"] = "zone_external_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["network_span_path"] = "network_span_path"
	queryParams["include_node_details"] = "include_node_details"
	queryParams["page_size"] = "page_size"
	queryParams["status"] = "status"
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
		"/policy/api/v1/infra/distributed-vlan-connections/{distributedVlanConnectionId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
