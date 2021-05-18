// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Certificates.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package trust_management

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_MGMT_CLUSTER = "MGMT_CLUSTER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_MGMT_PLANE = "MGMT_PLANE"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_API = "API"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_NOTIFICATION_COLLECTOR = "NOTIFICATION_COLLECTOR"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_SYSLOG_SERVER = "SYSLOG_SERVER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_RSYSLOG_CLIENT = "RSYSLOG_CLIENT"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_APH = "APH"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CLIENT_AUTH = "CLIENT_AUTH"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_RMQ = "RMQ"

// Possible value for ``type`` of method Certificates#list.
const Certificates_LIST_TYPE_CERTIFICATE = "cluster_api_certificate"

// Possible value for ``usage`` of method Certificates#validate.
const Certificates_VALIDATE_USAGE_SERVER = "SERVER"

// Possible value for ``usage`` of method Certificates#validate.
const Certificates_VALIDATE_USAGE_CLIENT = "CLIENT"

func certificatesApplycertificateInputType() bindings.StructType {
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

func certificatesApplycertificateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func certificatesApplycertificateRestMetadata() protocol.OperationRestMetadata {
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
		"/global-manager/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = bindings.NewStringType()
	fieldNameMap["cert_id"] = "CertId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func certificatesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = bindings.NewStringType()
	fieldNameMap["cert_id"] = "CertId"
	paramsTypeMap["cert_id"] = bindings.NewStringType()
	paramsTypeMap["certId"] = bindings.NewStringType()
	pathParams["cert_id"] = "certId"
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
		"/global-manager/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = bindings.NewStringType()
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CertificateBindingType)
}

func certificatesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = bindings.NewStringType()
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["details"] = "Details"
	paramsTypeMap["cert_id"] = bindings.NewStringType()
	paramsTypeMap["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["certId"] = bindings.NewStringType()
	pathParams["cert_id"] = "certId"
	queryParams["details"] = "details"
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
		"/global-manager/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesImportcertificateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["trust_object_data"] = bindings.NewReferenceType(model.TrustObjectDataBindingType)
	fieldNameMap["trust_object_data"] = "TrustObjectData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesImportcertificateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CertificateListBindingType)
}

func certificatesImportcertificateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["trust_object_data"] = bindings.NewReferenceType(model.TrustObjectDataBindingType)
	fieldNameMap["trust_object_data"] = "TrustObjectData"
	paramsTypeMap["trust_object_data"] = bindings.NewReferenceType(model.TrustObjectDataBindingType)
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
		"action=import",
		"trust_object_data",
		"POST",
		"/global-manager/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["details"] = "Details"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CertificateListBindingType)
}

func certificatesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["details"] = "Details"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["details"] = "details"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
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
		"/global-manager/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesSetapplianceproxycertificateforintersitecommunicationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["set_inter_site_aph_certificate_request"] = bindings.NewReferenceType(model.SetInterSiteAphCertificateRequestBindingType)
	fieldNameMap["set_inter_site_aph_certificate_request"] = "SetInterSiteAphCertificateRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesSetapplianceproxycertificateforintersitecommunicationOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func certificatesSetapplianceproxycertificateforintersitecommunicationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["set_inter_site_aph_certificate_request"] = bindings.NewReferenceType(model.SetInterSiteAphCertificateRequestBindingType)
	fieldNameMap["set_inter_site_aph_certificate_request"] = "SetInterSiteAphCertificateRequest"
	paramsTypeMap["set_inter_site_aph_certificate_request"] = bindings.NewReferenceType(model.SetInterSiteAphCertificateRequestBindingType)
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
		"action=set_appliance_proxy_certificate_for_inter_site_communication",
		"set_inter_site_aph_certificate_request",
		"POST",
		"/global-manager/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesSetpicertificateforfederationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["set_principal_identity_certificate_for_federation_request"] = bindings.NewReferenceType(model.SetPrincipalIdentityCertificateForFederationRequestBindingType)
	fieldNameMap["set_principal_identity_certificate_for_federation_request"] = "SetPrincipalIdentityCertificateForFederationRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesSetpicertificateforfederationOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func certificatesSetpicertificateforfederationRestMetadata() protocol.OperationRestMetadata {
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
		"/global-manager/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = bindings.NewStringType()
	fields["usage"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["usage"] = "Usage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CertificateCheckingStatusBindingType)
}

func certificatesValidateRestMetadata() protocol.OperationRestMetadata {
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
		"/global-manager/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
