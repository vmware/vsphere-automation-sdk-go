// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: HostSwitchProfiles.
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

// Possible value for ``deploymentType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_DEPLOYMENT_TYPE_VIRTUAL_MACHINE = "VIRTUAL_MACHINE"

// Possible value for ``deploymentType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_DEPLOYMENT_TYPE_PHYSICAL_MACHINE = "PHYSICAL_MACHINE"

// Possible value for ``deploymentType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_DEPLOYMENT_TYPE_UNKNOWN = "UNKNOWN"

// Possible value for ``hostswitchProfileType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_HOSTSWITCH_PROFILE_TYPE_UPLINKHOSTSWITCHPROFILE = "UplinkHostSwitchProfile"

// Possible value for ``hostswitchProfileType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_HOSTSWITCH_PROFILE_TYPE_LLDPHOSTSWITCHPROFILE = "LldpHostSwitchProfile"

// Possible value for ``hostswitchProfileType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_HOSTSWITCH_PROFILE_TYPE_NIOCPROFILE = "NiocProfile"

// Possible value for ``hostswitchProfileType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_HOSTSWITCH_PROFILE_TYPE_EXTRACONFIGHOSTSWITCHPROFILE = "ExtraConfigHostSwitchProfile"

// Possible value for ``nodeType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_NODE_TYPE_EDGENODE = "EdgeNode"

// Possible value for ``nodeType`` of method HostSwitchProfiles#list.
const HostSwitchProfiles_LIST_NODE_TYPE_PUBLICCLOUDGATEWAYNODE = "PublicCloudGatewayNode"

func hostSwitchProfilesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["base_host_switch_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
	fieldNameMap["base_host_switch_profile"] = "BaseHostSwitchProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostSwitchProfilesCreateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
}

func hostSwitchProfilesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["base_host_switch_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
	fieldNameMap["base_host_switch_profile"] = "BaseHostSwitchProfile"
	paramsTypeMap["base_host_switch_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
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
		"base_host_switch_profile",
		"POST",
		"/api/v1/host-switch-profiles",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func hostSwitchProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_switch_profile_id"] = bindings.NewStringType()
	fieldNameMap["host_switch_profile_id"] = "HostSwitchProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostSwitchProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hostSwitchProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["host_switch_profile_id"] = bindings.NewStringType()
	fieldNameMap["host_switch_profile_id"] = "HostSwitchProfileId"
	paramsTypeMap["host_switch_profile_id"] = bindings.NewStringType()
	paramsTypeMap["hostSwitchProfileId"] = bindings.NewStringType()
	pathParams["host_switch_profile_id"] = "hostSwitchProfileId"
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
		"/api/v1/host-switch-profiles/{hostSwitchProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func hostSwitchProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_switch_profile_id"] = bindings.NewStringType()
	fieldNameMap["host_switch_profile_id"] = "HostSwitchProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostSwitchProfilesGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
}

func hostSwitchProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["host_switch_profile_id"] = bindings.NewStringType()
	fieldNameMap["host_switch_profile_id"] = "HostSwitchProfileId"
	paramsTypeMap["host_switch_profile_id"] = bindings.NewStringType()
	paramsTypeMap["hostSwitchProfileId"] = bindings.NewStringType()
	pathParams["host_switch_profile_id"] = "hostSwitchProfileId"
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
		"/api/v1/host-switch-profiles/{hostSwitchProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func hostSwitchProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hostswitch_profile_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["deployment_type"] = "DeploymentType"
	fieldNameMap["hostswitch_profile_type"] = "HostswitchProfileType"
	fieldNameMap["include_system_owned"] = "IncludeSystemOwned"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_type"] = "NodeType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostSwitchProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.HostSwitchProfilesListResultBindingType)
}

func hostSwitchProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hostswitch_profile_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["deployment_type"] = "DeploymentType"
	fieldNameMap["hostswitch_profile_type"] = "HostswitchProfileType"
	fieldNameMap["include_system_owned"] = "IncludeSystemOwned"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_type"] = "NodeType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["uplink_teaming_policy_name"] = "UplinkTeamingPolicyName"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["hostswitch_profile_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["deployment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["uplink_teaming_policy_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["include_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["hostswitch_profile_type"] = "hostswitch_profile_type"
	queryParams["uplink_teaming_policy_name"] = "uplink_teaming_policy_name"
	queryParams["deployment_type"] = "deployment_type"
	queryParams["include_system_owned"] = "include_system_owned"
	queryParams["node_type"] = "node_type"
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
		"/api/v1/host-switch-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func hostSwitchProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_switch_profile_id"] = bindings.NewStringType()
	fields["base_host_switch_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
	fieldNameMap["host_switch_profile_id"] = "HostSwitchProfileId"
	fieldNameMap["base_host_switch_profile"] = "BaseHostSwitchProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostSwitchProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
}

func hostSwitchProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["host_switch_profile_id"] = bindings.NewStringType()
	fields["base_host_switch_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
	fieldNameMap["host_switch_profile_id"] = "HostSwitchProfileId"
	fieldNameMap["base_host_switch_profile"] = "BaseHostSwitchProfile"
	paramsTypeMap["host_switch_profile_id"] = bindings.NewStringType()
	paramsTypeMap["base_host_switch_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseHostSwitchProfileBindingType)}, bindings.REST)
	paramsTypeMap["hostSwitchProfileId"] = bindings.NewStringType()
	pathParams["host_switch_profile_id"] = "hostSwitchProfileId"
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
		"base_host_switch_profile",
		"PUT",
		"/api/v1/host-switch-profiles/{hostSwitchProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
