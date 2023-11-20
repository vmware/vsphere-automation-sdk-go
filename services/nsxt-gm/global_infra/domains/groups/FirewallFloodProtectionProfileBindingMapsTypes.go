// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FirewallFloodProtectionProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package groups

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func firewallFloodProtectionProfileBindingMapsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FirewallFloodProtectionProfileBindingMapsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func firewallFloodProtectionProfileBindingMapsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FirewallFloodProtectionProfileBindingMapsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
}

func firewallFloodProtectionProfileBindingMapsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FirewallFloodProtectionProfileBindingMapsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapListResultBindingType)
}

func firewallFloodProtectionProfileBindingMapsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FirewallFloodProtectionProfileBindingMapsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func firewallFloodProtectionProfileBindingMapsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_firewall_flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"policy_firewall_flood_protection_profile_binding_map",
		"PATCH",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FirewallFloodProtectionProfileBindingMapsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
}

func firewallFloodProtectionProfileBindingMapsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["policy_firewall_flood_protection_profile_binding_map"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = vapiBindings_.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"policy_firewall_flood_protection_profile_binding_map",
		"PUT",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
