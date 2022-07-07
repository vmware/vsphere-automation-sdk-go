// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package logical_switches

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Status#get.
const Status_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Status#get.
const Status_GET_SOURCE_CACHED = "cached"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_DEFAULT = "DEFAULT"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_SERVICE_PLANE = "SERVICE_PLANE"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_DHCP_RELAY = "DHCP_RELAY"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_GLOBAL = "GLOBAL"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_INTER_ROUTER = "INTER_ROUTER"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_EVPN = "EVPN"

// Possible value for ``switchType`` of method Status#get.
const Status_GET_SWITCH_TYPE_DVPG = "DVPG"

// Possible value for ``transportType`` of method Status#get.
const Status_GET_TRANSPORT_TYPE_OVERLAY = "OVERLAY"

// Possible value for ``transportType`` of method Status#get.
const Status_GET_TRANSPORT_TYPE_VLAN = "VLAN"

func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switch_type"] = bindings.NewOptionalType(bindings.NewStringType())
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
	fieldNameMap["switch_type"] = "SwitchType"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalSwitchStatusSummaryBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
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
	fields["switch_type"] = bindings.NewOptionalType(bindings.NewStringType())
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
	fieldNameMap["switch_type"] = "SwitchType"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_type"] = "TransportType"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	fieldNameMap["vlan"] = "Vlan"
	fieldNameMap["vni"] = "Vni"
	paramsTypeMap["switch_type"] = bindings.NewOptionalType(bindings.NewStringType())
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
	queryParams["switch_type"] = "switch_type"
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
