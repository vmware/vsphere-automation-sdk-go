// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: LdapServers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package domains

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``action`` of method LdapServers#create0.
const LdapServers_CREATE_0_ACTION_CONNECTIVITY = "CONNECTIVITY"

func ldapServersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["directory_ldap_server"] = "DirectoryLdapServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapServersCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
}

func ldapServersCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["directory_ldap_server"] = "DirectoryLdapServer"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	paramsTypeMap["domainId"] = bindings.NewStringType()
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
		"directory_ldap_server",
		"POST",
		"/api/v1/directory/domains/{domainId}/ldap-servers",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersCreate0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapServersCreate0OutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ldapServersCreate0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
	pathParams["domain_id"] = "domainId"
	queryParams["action"] = "action"
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
		"POST",
		"/api/v1/directory/domains/{domainId}/ldap-servers/{serverId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapServersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ldapServersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
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
		"/api/v1/directory/domains/{domainId}/ldap-servers/{serverId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapServersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
}

func ldapServersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
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
		"/api/v1/directory/domains/{domainId}/ldap-servers/{serverId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapServersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectoryLdapServerListResultsBindingType)
}

func ldapServersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	pathParams["domain_id"] = "domainId"
	queryParams["cursor"] = "cursor"
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
		"/api/v1/directory/domains/{domainId}/ldap-servers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapServersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fields["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["directory_ldap_server"] = "DirectoryLdapServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapServersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
}

func ldapServersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["server_id"] = bindings.NewStringType()
	fields["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["directory_ldap_server"] = "DirectoryLdapServer"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	paramsTypeMap["server_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["serverId"] = bindings.NewStringType()
	pathParams["server_id"] = "serverId"
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
		"directory_ldap_server",
		"PUT",
		"/api/v1/directory/domains/{domainId}/ldap-servers/{serverId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
