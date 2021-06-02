// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Default.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_MGMT_CLUSTER = "MGMT_CLUSTER"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_MGMT_PLANE = "MGMT_PLANE"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_API = "API"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_NOTIFICATION_COLLECTOR = "NOTIFICATION_COLLECTOR"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_SYSLOG_SERVER = "SYSLOG_SERVER"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_RSYSLOG_CLIENT = "RSYSLOG_CLIENT"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_APH = "APH"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_CLIENT_AUTH = "CLIENT_AUTH"

// Possible value for ``serviceType`` of method Default#applyCertificateApplyCertificate.
const Default_APPLY_CERTIFICATE_APPLY_CERTIFICATE_SERVICE_TYPE_RMQ = "RMQ"

// Possible value for ``usage`` of method Default#validateCertificateValidate.
const Default_VALIDATE_CERTIFICATE_VALIDATE_USAGE_SERVER = "SERVER"

// Possible value for ``usage`` of method Default#validateCertificateValidate.
const Default_VALIDATE_CERTIFICATE_VALIDATE_USAGE_CLIENT = "CLIENT"

func defaultApplyCertificateApplyCertificateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = bindings.NewStringType()
	fields["service_type"] = bindings.NewStringType()
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["service_type"] = "ServiceType"
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func defaultApplyCertificateApplyCertificateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func defaultApplyCertificateApplyCertificateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = bindings.NewStringType()
	fields["service_type"] = bindings.NewStringType()
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["service_type"] = "ServiceType"
	fieldNameMap["node_id"] = "NodeId"
	paramsTypeMap["cert_id"] = bindings.NewStringType()
	paramsTypeMap["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_type"] = bindings.NewStringType()
	paramsTypeMap["certId"] = bindings.NewStringType()
	pathParams["cert_id"] = "certId"
	queryParams["service_type"] = "service_type"
	queryParams["node_id"] = "node_id"
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
		"action=apply_certificate",
		"",
		"POST",
		"/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func defaultSetPrincipalIdentityCertificateForFederationSetPiCertificateForFederationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["set_principal_identity_certificate_for_federation_request"] = bindings.NewReferenceType(model.SetPrincipalIdentityCertificateForFederationRequestBindingType)
	fieldNameMap["set_principal_identity_certificate_for_federation_request"] = "SetPrincipalIdentityCertificateForFederationRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func defaultSetPrincipalIdentityCertificateForFederationSetPiCertificateForFederationOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func defaultSetPrincipalIdentityCertificateForFederationSetPiCertificateForFederationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["set_principal_identity_certificate_for_federation_request"] = bindings.NewReferenceType(model.SetPrincipalIdentityCertificateForFederationRequestBindingType)
	fieldNameMap["set_principal_identity_certificate_for_federation_request"] = "SetPrincipalIdentityCertificateForFederationRequest"
	paramsTypeMap["set_principal_identity_certificate_for_federation_request"] = bindings.NewReferenceType(model.SetPrincipalIdentityCertificateForFederationRequestBindingType)
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
		"action=set_pi_certificate_for_federation",
		"set_principal_identity_certificate_for_federation_request",
		"POST",
		"/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func defaultValidateCertificateValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = bindings.NewStringType()
	fields["usage"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["usage"] = "Usage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func defaultValidateCertificateValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CertificateCheckingStatusBindingType)
}

func defaultValidateCertificateValidateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = bindings.NewStringType()
	fields["usage"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["usage"] = "Usage"
	paramsTypeMap["usage"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cert_id"] = bindings.NewStringType()
	paramsTypeMap["certId"] = bindings.NewStringType()
	pathParams["cert_id"] = "certId"
	queryParams["usage"] = "usage"
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
		"action=validate",
		"",
		"GET",
		"/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
