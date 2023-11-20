// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: LdapServers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package identity_firewall_stores

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func ldapServersDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LdapServersDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func ldapServersDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["ldap_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ldapServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["ldap_server_id"] = "ldapServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/ldap-servers/{ldapServerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LdapServersGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
}

func ldapServersGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["ldap_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ldapServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["ldap_server_id"] = "ldapServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/ldap-servers/{ldapServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LdapServersListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerListResultsBindingType)
}

func ldapServersListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	queryParams["sort_by"] = "sort_by"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/ldap-servers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_ldap_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["identity_firewall_store_ldap_server"] = "IdentityFirewallStoreLdapServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LdapServersPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
}

func ldapServersPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_ldap_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["identity_firewall_store_ldap_server"] = "IdentityFirewallStoreLdapServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["ldap_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identity_firewall_store_ldap_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ldapServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["ldap_server_id"] = "ldapServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"identity_firewall_store_ldap_server",
		"PATCH",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/ldap-servers/{ldapServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_ldap_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["identity_firewall_store_ldap_server"] = "IdentityFirewallStoreLdapServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LdapServersUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
}

func ldapServersUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	fields["ldap_server_id"] = vapiBindings_.NewStringType()
	fields["identity_firewall_store_ldap_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["ldap_server_id"] = "LdapServerId"
	fieldNameMap["identity_firewall_store_ldap_server"] = "IdentityFirewallStoreLdapServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["ldap_server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identity_firewall_store_ldap_server"] = vapiBindings_.NewReferenceType(nsx_policyModel.IdentityFirewallStoreLdapServerBindingType)
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = vapiBindings_.NewStringType()
	paramsTypeMap["ldapServerId"] = vapiBindings_.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["ldap_server_id"] = "ldapServerId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"identity_firewall_store_ldap_server",
		"PUT",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/ldap-servers/{ldapServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
