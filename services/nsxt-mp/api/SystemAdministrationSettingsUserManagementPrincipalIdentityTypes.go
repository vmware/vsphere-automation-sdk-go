// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationSettingsUserManagementPrincipalIdentity.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func systemAdministrationSettingsUserManagementPrincipalIdentityAddOidcEndPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["oidc_end_point"] = bindings.NewReferenceType(model.OidcEndPointBindingType)
	fieldNameMap["oidc_end_point"] = "OidcEndPoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityAddOidcEndPointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.OidcEndPointBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityAddOidcEndPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["oidc_end_point"] = bindings.NewReferenceType(model.OidcEndPointBindingType)
	fieldNameMap["oidc_end_point"] = "OidcEndPoint"
	paramsTypeMap["oidc_end_point"] = bindings.NewReferenceType(model.OidcEndPointBindingType)
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
		"oidc_end_point",
		"POST",
		"/api/v1/trust-management/oidc-uris",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityDeletePrincipalIdentityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityDeletePrincipalIdentityOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationSettingsUserManagementPrincipalIdentityDeletePrincipalIdentityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	paramsTypeMap["principal_identity_id"] = bindings.NewStringType()
	paramsTypeMap["principalIdentityId"] = bindings.NewStringType()
	pathParams["principal_identity_id"] = "principalIdentityId"
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
		"/api/v1/trust-management/principal-identities/{principalIdentityId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityDeleteTokenBasedPrincipalIdentityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityDeleteTokenBasedPrincipalIdentityOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationSettingsUserManagementPrincipalIdentityDeleteTokenBasedPrincipalIdentityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	paramsTypeMap["principal_identity_id"] = bindings.NewStringType()
	paramsTypeMap["principalIdentityId"] = bindings.NewStringType()
	pathParams["principal_identity_id"] = "principalIdentityId"
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
		"/api/v1/trust-management/token-principal-identities/{principalIdentityId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetOidcEndPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewStringType()
	fields["refresh"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["id"] = "Id"
	fieldNameMap["refresh"] = "Refresh"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetOidcEndPointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.OidcEndPointBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetOidcEndPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["id"] = bindings.NewStringType()
	fields["refresh"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["id"] = "Id"
	fieldNameMap["refresh"] = "Refresh"
	paramsTypeMap["refresh"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	queryParams["refresh"] = "refresh"
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
		"/api/v1/trust-management/oidc-uris/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentitiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentitiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PrincipalIdentityListBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentitiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/trust-management/principal-identities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentityOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PrincipalIdentityBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetPrincipalIdentityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	paramsTypeMap["principal_identity_id"] = bindings.NewStringType()
	paramsTypeMap["principalIdentityId"] = bindings.NewStringType()
	pathParams["principal_identity_id"] = "principalIdentityId"
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
		"/api/v1/trust-management/principal-identities/{principalIdentityId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetTokenBasedPrincipalIdentityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetTokenBasedPrincipalIdentityOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TokenBasedPrincipalIdentityBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityGetTokenBasedPrincipalIdentityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["principal_identity_id"] = bindings.NewStringType()
	fieldNameMap["principal_identity_id"] = "PrincipalIdentityId"
	paramsTypeMap["principal_identity_id"] = bindings.NewStringType()
	paramsTypeMap["principalIdentityId"] = bindings.NewStringType()
	pathParams["principal_identity_id"] = "principalIdentityId"
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
		"/api/v1/trust-management/token-principal-identities/{principalIdentityId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityListOidcEndPointsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityListOidcEndPointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.OidcEndPointListResultBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityListOidcEndPointsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/trust-management/oidc-uris",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityListTokenBasedPrincipalIdentitiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityListTokenBasedPrincipalIdentitiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TokenBasedPrincipalIdentityListResultBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityListTokenBasedPrincipalIdentitiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/trust-management/token-principal-identities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal_identity"] = bindings.NewReferenceType(model.PrincipalIdentityBindingType)
	fieldNameMap["principal_identity"] = "PrincipalIdentity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PrincipalIdentityBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["principal_identity"] = bindings.NewReferenceType(model.PrincipalIdentityBindingType)
	fieldNameMap["principal_identity"] = "PrincipalIdentity"
	paramsTypeMap["principal_identity"] = bindings.NewReferenceType(model.PrincipalIdentityBindingType)
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
		"principal_identity",
		"POST",
		"/api/v1/trust-management/principal-identities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityWithCertificateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal_identity_with_certificate"] = bindings.NewReferenceType(model.PrincipalIdentityWithCertificateBindingType)
	fieldNameMap["principal_identity_with_certificate"] = "PrincipalIdentityWithCertificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityWithCertificateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PrincipalIdentityBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterPrincipalIdentityWithCertificateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["principal_identity_with_certificate"] = bindings.NewReferenceType(model.PrincipalIdentityWithCertificateBindingType)
	fieldNameMap["principal_identity_with_certificate"] = "PrincipalIdentityWithCertificate"
	paramsTypeMap["principal_identity_with_certificate"] = bindings.NewReferenceType(model.PrincipalIdentityWithCertificateBindingType)
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
		"principal_identity_with_certificate",
		"POST",
		"/api/v1/trust-management/principal-identities/with-certificate",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterTokenBasedPrincipalIdentityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["token_based_principal_identity"] = bindings.NewReferenceType(model.TokenBasedPrincipalIdentityBindingType)
	fieldNameMap["token_based_principal_identity"] = "TokenBasedPrincipalIdentity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterTokenBasedPrincipalIdentityOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TokenBasedPrincipalIdentityBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityRegisterTokenBasedPrincipalIdentityRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["token_based_principal_identity"] = bindings.NewReferenceType(model.TokenBasedPrincipalIdentityBindingType)
	fieldNameMap["token_based_principal_identity"] = "TokenBasedPrincipalIdentity"
	paramsTypeMap["token_based_principal_identity"] = bindings.NewReferenceType(model.TokenBasedPrincipalIdentityBindingType)
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
		"token_based_principal_identity",
		"POST",
		"/api/v1/trust-management/token-principal-identities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityUpdateOidcEndPointThumbprintUpdateThumbprintInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_oidc_end_point_thumbprint_request"] = bindings.NewReferenceType(model.UpdateOidcEndPointThumbprintRequestBindingType)
	fieldNameMap["update_oidc_end_point_thumbprint_request"] = "UpdateOidcEndPointThumbprintRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityUpdateOidcEndPointThumbprintUpdateThumbprintOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.OidcEndPointBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityUpdateOidcEndPointThumbprintUpdateThumbprintRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["update_oidc_end_point_thumbprint_request"] = bindings.NewReferenceType(model.UpdateOidcEndPointThumbprintRequestBindingType)
	fieldNameMap["update_oidc_end_point_thumbprint_request"] = "UpdateOidcEndPointThumbprintRequest"
	paramsTypeMap["update_oidc_end_point_thumbprint_request"] = bindings.NewReferenceType(model.UpdateOidcEndPointThumbprintRequestBindingType)
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
		"action=update_thumbprint",
		"update_oidc_end_point_thumbprint_request",
		"POST",
		"/api/v1/trust-management/oidc-uris",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsUserManagementPrincipalIdentityUpdatePrincipalIdentityCertificateUpdateCertificateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_principal_identity_certificate_request"] = bindings.NewReferenceType(model.UpdatePrincipalIdentityCertificateRequestBindingType)
	fieldNameMap["update_principal_identity_certificate_request"] = "UpdatePrincipalIdentityCertificateRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityUpdatePrincipalIdentityCertificateUpdateCertificateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PrincipalIdentityBindingType)
}

func systemAdministrationSettingsUserManagementPrincipalIdentityUpdatePrincipalIdentityCertificateUpdateCertificateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["update_principal_identity_certificate_request"] = bindings.NewReferenceType(model.UpdatePrincipalIdentityCertificateRequestBindingType)
	fieldNameMap["update_principal_identity_certificate_request"] = "UpdatePrincipalIdentityCertificateRequest"
	paramsTypeMap["update_principal_identity_certificate_request"] = bindings.NewReferenceType(model.UpdatePrincipalIdentityCertificateRequestBindingType)
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
		"action=update_certificate",
		"update_principal_identity_certificate_request",
		"POST",
		"/api/v1/trust-management/principal-identities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
