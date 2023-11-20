// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Observations.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package traceflows

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_PHYSICAL = "PHYSICAL"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_LR = "LR"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_LS = "LS"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_DFW = "DFW"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_BRIDGE = "BRIDGE"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_EDGE_TUNNEL = "EDGE_TUNNEL"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_EDGE_HOSTSWITCH = "EDGE_HOSTSWITCH"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_FW_BRIDGE = "FW_BRIDGE"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_EDGE_RTEP_TUNNEL = "EDGE_RTEP_TUNNEL"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_LOAD_BALANCER = "LOAD_BALANCER"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_NAT = "NAT"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_IPSEC = "IPSEC"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_SERVICE_INSERTION = "SERVICE_INSERTION"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_VMC = "VMC"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_SPOOFGUARD = "SPOOFGUARD"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_EDGE_FW = "EDGE_FW"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_DLB = "DLB"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_ANTREA_SPOOFGUARD = "ANTREA_SPOOFGUARD"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_ANTREA_LB = "ANTREA_LB"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_ANTREA_ROUTING = "ANTREA_ROUTING"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_ANTREA_DFW = "ANTREA_DFW"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_ANTREA_FORWARDING = "ANTREA_FORWARDING"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_HOST_SWITCH = "HOST_SWITCH"

// Possible value for ``componentType`` of method Observations#list.
const Observations_LIST_COMPONENT_TYPE_UNKNOWN = "UNKNOWN"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONFORWARDED = "TraceflowObservationForwarded"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONDROPPED = "TraceflowObservationDropped"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONDELIVERED = "TraceflowObservationDelivered"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONRECEIVED = "TraceflowObservationReceived"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONFORWARDEDLOGICAL = "TraceflowObservationForwardedLogical"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONDROPPEDLOGICAL = "TraceflowObservationDroppedLogical"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONRECEIVEDLOGICAL = "TraceflowObservationReceivedLogical"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONREPLICATIONLOGICAL = "TraceflowObservationReplicationLogical"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONRELAYEDLOGICAL = "TraceflowObservationRelayedLogical"

// Possible value for ``resourceType`` of method Observations#list.
const Observations_LIST_RESOURCE_TYPE_TRACEFLOWOBSERVATIONPROTECTED = "TraceflowObservationProtected"

func observationsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traceflow_id"] = vapiBindings_.NewStringType()
	fields["component_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["traceflow_id"] = "TraceflowId"
	fieldNameMap["component_name"] = "ComponentName"
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_node_name"] = "TransportNodeName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ObservationsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.TraceflowObservationListResultBindingType)
}

func observationsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traceflow_id"] = vapiBindings_.NewStringType()
	fields["component_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["traceflow_id"] = "TraceflowId"
	fieldNameMap["component_name"] = "ComponentName"
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["transport_node_name"] = "TransportNodeName"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["traceflow_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["component_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_node_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["traceflowId"] = vapiBindings_.NewStringType()
	pathParams["traceflow_id"] = "traceflowId"
	queryParams["cursor"] = "cursor"
	queryParams["component_type"] = "component_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["component_name"] = "component_name"
	queryParams["resource_type"] = "resource_type"
	queryParams["sort_by"] = "sort_by"
	queryParams["transport_node_name"] = "transport_node_name"
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
		"/api/v1/traceflows/{traceflowId}/observations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
