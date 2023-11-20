// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package logical_switches

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func statusGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["diagnostic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["switch_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["switching_profile_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["uplink_teaming_policy_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["vni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
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
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatusGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LogicalSwitchStatusSummaryBindingType)
}

func statusGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["diagnostic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["switch_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["switching_profile_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["uplink_teaming_policy_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["vni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
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
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["switching_profile_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["switch_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["uplink_teaming_policy_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vni"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["vlan"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["diagnostic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
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
		"/api/v1/logical-switches/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
