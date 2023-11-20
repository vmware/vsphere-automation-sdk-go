// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Certificates.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package trust_management

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_APH_TN = "APH_TN"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CLIENT_AUTH = "CLIENT_AUTH"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_RMQ = "RMQ"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_K8S_MSG_CLIENT = "K8S_MSG_CLIENT"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_WEB_PROXY = "WEB_PROXY"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_API = "CBM_API"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_CCP = "CBM_CCP"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_CSM = "CBM_CSM"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_MP = "CBM_MP"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_GM = "CBM_GM"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_AR = "CBM_AR"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_MONITORING = "CBM_MONITORING"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_IDPS_REPORTING = "CBM_IDPS_REPORTING"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_CM_INVENTORY = "CBM_CM_INVENTORY"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_MESSAGING_MANAGER = "CBM_MESSAGING_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_UPGRADE_COORDINATOR = "CBM_UPGRADE_COORDINATOR"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_SITE_MANAGER = "CBM_SITE_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_CLUSTER_MANAGER = "CBM_CLUSTER_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_CORFU = "CBM_CORFU"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CBM_SITE_PROXY_CLIENT = "CBM_SITE_PROXY_CLIENT"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_COMPUTE_MANAGER = "COMPUTE_MANAGER"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_CCP = "CCP"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_ANALYTICS_AGENT = "ANALYTICS_AGENT"

// Possible value for ``serviceType`` of method Certificates#applycertificate.
const Certificates_APPLYCERTIFICATE_SERVICE_TYPE_ANALYTICS_KAFKA = "ANALYTICS_KAFKA"

// Possible value for ``type`` of method Certificates#list.
const Certificates_LIST_TYPE_CLUSTER_API_CERTIFICATE = "cluster_api_certificate"

// Possible value for ``type`` of method Certificates#list.
const Certificates_LIST_TYPE_API_CERTIFICATE = "api_certificate"

// Possible value for ``usage`` of method Certificates#validate.
const Certificates_VALIDATE_USAGE_SERVER = "SERVER"

// Possible value for ``usage`` of method Certificates#validate.
const Certificates_VALIDATE_USAGE_CLIENT = "CLIENT"

func certificatesApplycertificateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = vapiBindings_.NewStringType()
	fields["service_type"] = vapiBindings_.NewStringType()
	fields["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["service_type"] = "ServiceType"
	fieldNameMap["node_id"] = "NodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesApplycertificateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func certificatesApplycertificateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = vapiBindings_.NewStringType()
	fields["service_type"] = vapiBindings_.NewStringType()
	fields["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["service_type"] = "ServiceType"
	fieldNameMap["node_id"] = "NodeId"
	paramsTypeMap["service_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["cert_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["certId"] = vapiBindings_.NewStringType()
	pathParams["cert_id"] = "certId"
	queryParams["service_type"] = "service_type"
	queryParams["node_id"] = "node_id"
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

func certificatesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cert_id"] = "CertId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func certificatesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = vapiBindings_.NewStringType()
	fieldNameMap["cert_id"] = "CertId"
	paramsTypeMap["cert_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["certId"] = vapiBindings_.NewStringType()
	pathParams["cert_id"] = "certId"
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
		"/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesFetchpeercertificatechainInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tls_service_endpoint"] = vapiBindings_.NewReferenceType(nsxModel.TlsServiceEndpointBindingType)
	fieldNameMap["tls_service_endpoint"] = "TlsServiceEndpoint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesFetchpeercertificatechainOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.PeerCertificateChainBindingType)
}

func certificatesFetchpeercertificatechainRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tls_service_endpoint"] = vapiBindings_.NewReferenceType(nsxModel.TlsServiceEndpointBindingType)
	fieldNameMap["tls_service_endpoint"] = "TlsServiceEndpoint"
	paramsTypeMap["tls_service_endpoint"] = vapiBindings_.NewReferenceType(nsxModel.TlsServiceEndpointBindingType)
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
		"action=fetch_peer_certificate_chain",
		"tls_service_endpoint",
		"POST",
		"/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = vapiBindings_.NewStringType()
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["details"] = "Details"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.CertificateBindingType)
}

func certificatesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = vapiBindings_.NewStringType()
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["details"] = "Details"
	paramsTypeMap["cert_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["certId"] = vapiBindings_.NewStringType()
	pathParams["cert_id"] = "certId"
	queryParams["details"] = "details"
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
		"/api/v1/trust-management/certificates/{certId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesImportcertificateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["trust_object_data"] = vapiBindings_.NewReferenceType(nsxModel.TrustObjectDataBindingType)
	fieldNameMap["trust_object_data"] = "TrustObjectData"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesImportcertificateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.CertificateListBindingType)
}

func certificatesImportcertificateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["trust_object_data"] = vapiBindings_.NewReferenceType(nsxModel.TrustObjectDataBindingType)
	fieldNameMap["trust_object_data"] = "TrustObjectData"
	paramsTypeMap["trust_object_data"] = vapiBindings_.NewReferenceType(nsxModel.TrustObjectDataBindingType)
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
		"action=import",
		"trust_object_data",
		"POST",
		"/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesImporttrustedcaInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alias"] = vapiBindings_.NewStringType()
	fields["trust_object_data"] = vapiBindings_.NewReferenceType(nsxModel.TrustObjectDataBindingType)
	fieldNameMap["alias"] = "Alias"
	fieldNameMap["trust_object_data"] = "TrustObjectData"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesImporttrustedcaOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func certificatesImporttrustedcaRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alias"] = vapiBindings_.NewStringType()
	fields["trust_object_data"] = vapiBindings_.NewReferenceType(nsxModel.TrustObjectDataBindingType)
	fieldNameMap["alias"] = "Alias"
	fieldNameMap["trust_object_data"] = "TrustObjectData"
	paramsTypeMap["alias"] = vapiBindings_.NewStringType()
	paramsTypeMap["trust_object_data"] = vapiBindings_.NewReferenceType(nsxModel.TrustObjectDataBindingType)
	paramsTypeMap["alias"] = vapiBindings_.NewStringType()
	pathParams["alias"] = "alias"
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
		"action=import_trusted_ca",
		"trust_object_data",
		"POST",
		"/api/v1/trust-management/certificates/{alias}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["details"] = "Details"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.CertificateListBindingType)
}

func certificatesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["details"] = "Details"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["details"] = "details"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["node_id"] = "node_id"
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
		"/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesSetapplianceproxycertificateforintersitecommunicationInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["set_inter_site_aph_certificate_request"] = vapiBindings_.NewReferenceType(nsxModel.SetInterSiteAphCertificateRequestBindingType)
	fieldNameMap["set_inter_site_aph_certificate_request"] = "SetInterSiteAphCertificateRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesSetapplianceproxycertificateforintersitecommunicationOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func certificatesSetapplianceproxycertificateforintersitecommunicationRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["set_inter_site_aph_certificate_request"] = vapiBindings_.NewReferenceType(nsxModel.SetInterSiteAphCertificateRequestBindingType)
	fieldNameMap["set_inter_site_aph_certificate_request"] = "SetInterSiteAphCertificateRequest"
	paramsTypeMap["set_inter_site_aph_certificate_request"] = vapiBindings_.NewReferenceType(nsxModel.SetInterSiteAphCertificateRequestBindingType)
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
		"action=set_appliance_proxy_certificate_for_inter_site_communication",
		"set_inter_site_aph_certificate_request",
		"POST",
		"/api/v1/trust-management/certificates",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func certificatesSetpicertificateforfederationInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["set_principal_identity_certificate_for_federation_request"] = vapiBindings_.NewReferenceType(nsxModel.SetPrincipalIdentityCertificateForFederationRequestBindingType)
	fieldNameMap["set_principal_identity_certificate_for_federation_request"] = "SetPrincipalIdentityCertificateForFederationRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesSetpicertificateforfederationOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func certificatesSetpicertificateforfederationRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["set_principal_identity_certificate_for_federation_request"] = vapiBindings_.NewReferenceType(nsxModel.SetPrincipalIdentityCertificateForFederationRequestBindingType)
	fieldNameMap["set_principal_identity_certificate_for_federation_request"] = "SetPrincipalIdentityCertificateForFederationRequest"
	paramsTypeMap["set_principal_identity_certificate_for_federation_request"] = vapiBindings_.NewReferenceType(nsxModel.SetPrincipalIdentityCertificateForFederationRequestBindingType)
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

func certificatesValidateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_id"] = vapiBindings_.NewStringType()
	fields["usage"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["usage"] = "Usage"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CertificatesValidateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.CertificateCheckingStatusBindingType)
}

func certificatesValidateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cert_id"] = vapiBindings_.NewStringType()
	fields["usage"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["cert_id"] = "CertId"
	fieldNameMap["usage"] = "Usage"
	paramsTypeMap["usage"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["cert_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["certId"] = vapiBindings_.NewStringType()
	pathParams["cert_id"] = "certId"
	queryParams["usage"] = "usage"
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
