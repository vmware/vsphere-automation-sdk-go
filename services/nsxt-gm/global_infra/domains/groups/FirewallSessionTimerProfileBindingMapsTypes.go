// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FirewallSessionTimerProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package groups

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func firewallSessionTimerProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func firewallSessionTimerProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileBindingMapId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["firewall_session_timer_profile_binding_map_id"] = "firewallSessionTimerProfileBindingMapId"
	pathParams["domain_id"] = "domainId"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-session-timer-profile-binding-maps/{firewallSessionTimerProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
}

func firewallSessionTimerProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileBindingMapId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["firewall_session_timer_profile_binding_map_id"] = "firewallSessionTimerProfileBindingMapId"
	pathParams["domain_id"] = "domainId"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-session-timer-profile-binding-maps/{firewallSessionTimerProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapListResultBindingType)
}

func firewallSessionTimerProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-session-timer-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	fieldNameMap["policy_firewall_session_timer_profile_binding_map"] = "PolicyFirewallSessionTimerProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func firewallSessionTimerProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	fieldNameMap["policy_firewall_session_timer_profile_binding_map"] = "PolicyFirewallSessionTimerProfileBindingMap"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["policy_firewall_session_timer_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
	paramsTypeMap["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileBindingMapId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["firewall_session_timer_profile_binding_map_id"] = "firewallSessionTimerProfileBindingMapId"
	pathParams["domain_id"] = "domainId"
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
		"policy_firewall_session_timer_profile_binding_map",
		"PATCH",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-session-timer-profile-binding-maps/{firewallSessionTimerProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallSessionTimerProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	fieldNameMap["policy_firewall_session_timer_profile_binding_map"] = "PolicyFirewallSessionTimerProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallSessionTimerProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
}

func firewallSessionTimerProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_session_timer_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_session_timer_profile_binding_map_id"] = "FirewallSessionTimerProfileBindingMapId"
	fieldNameMap["policy_firewall_session_timer_profile_binding_map"] = "PolicyFirewallSessionTimerProfileBindingMap"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["policy_firewall_session_timer_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallSessionTimerProfileBindingMapBindingType)
	paramsTypeMap["firewall_session_timer_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallSessionTimerProfileBindingMapId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["firewall_session_timer_profile_binding_map_id"] = "firewallSessionTimerProfileBindingMapId"
	pathParams["domain_id"] = "domainId"
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
		"policy_firewall_session_timer_profile_binding_map",
		"PUT",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-session-timer-profile-binding-maps/{firewallSessionTimerProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
