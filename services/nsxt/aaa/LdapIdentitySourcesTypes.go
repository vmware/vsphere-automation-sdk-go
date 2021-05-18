// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: LdapIdentitySources.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package aaa

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func ldapIdentitySourcesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ldapIdentitySourcesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	paramsTypeMap["ldap_identity_source_id"] = bindings.NewStringType()
	paramsTypeMap["ldapIdentitySourceId"] = bindings.NewStringType()
	pathParams["ldap_identity_source_id"] = "ldapIdentitySourceId"
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
		"/policy/api/v1/aaa/ldap-identity-sources/{ldapIdentitySourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesFetchcertificateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_source_ldap_server_endpoint"] = bindings.NewReferenceType(model.IdentitySourceLdapServerEndpointBindingType)
	fieldNameMap["identity_source_ldap_server_endpoint"] = "IdentitySourceLdapServerEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesFetchcertificateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PeerCertificateChainBindingType)
}

func ldapIdentitySourcesFetchcertificateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_source_ldap_server_endpoint"] = bindings.NewReferenceType(model.IdentitySourceLdapServerEndpointBindingType)
	fieldNameMap["identity_source_ldap_server_endpoint"] = "IdentitySourceLdapServerEndpoint"
	paramsTypeMap["identity_source_ldap_server_endpoint"] = bindings.NewReferenceType(model.IdentitySourceLdapServerEndpointBindingType)
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
		"action=fetch_certificate",
		"identity_source_ldap_server_endpoint",
		"POST",
		"/policy/api/v1/aaa/ldap-identity-sources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
}

func ldapIdentitySourcesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	paramsTypeMap["ldap_identity_source_id"] = bindings.NewStringType()
	paramsTypeMap["ldapIdentitySourceId"] = bindings.NewStringType()
	pathParams["ldap_identity_source_id"] = "ldapIdentitySourceId"
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
		"/policy/api/v1/aaa/ldap-identity-sources/{ldapIdentitySourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LdapIdentitySourceListResultBindingType)
}

func ldapIdentitySourcesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
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
		"/policy/api/v1/aaa/ldap-identity-sources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesProbeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesProbeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LdapIdentitySourceProbeResultsBindingType)
}

func ldapIdentitySourcesProbeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	paramsTypeMap["ldap_identity_source_id"] = bindings.NewStringType()
	paramsTypeMap["ldapIdentitySourceId"] = bindings.NewStringType()
	pathParams["ldap_identity_source_id"] = "ldapIdentitySourceId"
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
		"action=probe",
		"",
		"POST",
		"/policy/api/v1/aaa/ldap-identity-sources/{ldapIdentitySourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesProbeidentitysourceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ldap_identity_source"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
	fieldNameMap["ldap_identity_source"] = "LdapIdentitySource"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesProbeidentitysourceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LdapIdentitySourceProbeResultsBindingType)
}

func ldapIdentitySourcesProbeidentitysourceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ldap_identity_source"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
	fieldNameMap["ldap_identity_source"] = "LdapIdentitySource"
	paramsTypeMap["ldap_identity_source"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
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
		"action=probe_identity_source",
		"ldap_identity_source",
		"POST",
		"/policy/api/v1/aaa/ldap-identity-sources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesProbeldapserverInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_source_ldap_server"] = bindings.NewReferenceType(model.IdentitySourceLdapServerBindingType)
	fieldNameMap["identity_source_ldap_server"] = "IdentitySourceLdapServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesProbeldapserverOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdentitySourceLdapServerProbeResultBindingType)
}

func ldapIdentitySourcesProbeldapserverRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_source_ldap_server"] = bindings.NewReferenceType(model.IdentitySourceLdapServerBindingType)
	fieldNameMap["identity_source_ldap_server"] = "IdentitySourceLdapServer"
	paramsTypeMap["identity_source_ldap_server"] = bindings.NewReferenceType(model.IdentitySourceLdapServerBindingType)
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
		"action=probe_ldap_server",
		"identity_source_ldap_server",
		"POST",
		"/policy/api/v1/aaa/ldap-identity-sources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ldapIdentitySourcesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fields["ldap_identity_source"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	fieldNameMap["ldap_identity_source"] = "LdapIdentitySource"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ldapIdentitySourcesUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
}

func ldapIdentitySourcesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ldap_identity_source_id"] = bindings.NewStringType()
	fields["ldap_identity_source"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
	fieldNameMap["ldap_identity_source_id"] = "LdapIdentitySourceId"
	fieldNameMap["ldap_identity_source"] = "LdapIdentitySource"
	paramsTypeMap["ldap_identity_source_id"] = bindings.NewStringType()
	paramsTypeMap["ldap_identity_source"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.LdapIdentitySourceBindingType)}, bindings.REST)
	paramsTypeMap["ldapIdentitySourceId"] = bindings.NewStringType()
	pathParams["ldap_identity_source_id"] = "ldapIdentitySourceId"
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
		"ldap_identity_source",
		"PUT",
		"/policy/api/v1/aaa/ldap-identity-sources/{ldapIdentitySourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
