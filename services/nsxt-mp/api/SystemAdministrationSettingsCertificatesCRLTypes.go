// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationSettingsCertificatesCRL.
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

func systemAdministrationSettingsCertificatesCRLCreateCrlDistributionPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crl_distribution_point"] = bindings.NewReferenceType(model.CrlDistributionPointBindingType)
	fieldNameMap["crl_distribution_point"] = "CrlDistributionPoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsCertificatesCRLCreateCrlDistributionPointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CrlDistributionPointBindingType)
}

func systemAdministrationSettingsCertificatesCRLCreateCrlDistributionPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["crl_distribution_point"] = bindings.NewReferenceType(model.CrlDistributionPointBindingType)
	fieldNameMap["crl_distribution_point"] = "CrlDistributionPoint"
	paramsTypeMap["crl_distribution_point"] = bindings.NewReferenceType(model.CrlDistributionPointBindingType)
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
		"crl_distribution_point",
		"POST",
		"/api/v1/trust-management/crl-distribution-points",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsCertificatesCRLDeleteCrlDistributionPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsCertificatesCRLDeleteCrlDistributionPointOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationSettingsCertificatesCRLDeleteCrlDistributionPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	paramsTypeMap["crl_distribution_point_id"] = bindings.NewStringType()
	paramsTypeMap["crlDistributionPointId"] = bindings.NewStringType()
	pathParams["crl_distribution_point_id"] = "crlDistributionPointId"
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
		"/api/v1/trust-management/crl-distribution-points/{crlDistributionPointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CrlDistributionPointBindingType)
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	paramsTypeMap["crl_distribution_point_id"] = bindings.NewStringType()
	paramsTypeMap["crlDistributionPointId"] = bindings.NewStringType()
	pathParams["crl_distribution_point_id"] = "crlDistributionPointId"
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
		"/api/v1/trust-management/crl-distribution-points/{crlDistributionPointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointPemInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crl_pem_request_type"] = bindings.NewReferenceType(model.CrlPemRequestTypeBindingType)
	fieldNameMap["crl_pem_request_type"] = "CrlPemRequestType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointPemOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointPemRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["crl_pem_request_type"] = bindings.NewReferenceType(model.CrlPemRequestTypeBindingType)
	fieldNameMap["crl_pem_request_type"] = "CrlPemRequestType"
	paramsTypeMap["crl_pem_request_type"] = bindings.NewReferenceType(model.CrlPemRequestTypeBindingType)
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
		"crl_pem_request_type",
		"POST",
		"/api/v1/trust-management/crl-distribution-points/pem-file",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CrlDistributionPointStatusBindingType)
}

func systemAdministrationSettingsCertificatesCRLGetCrlDistributionPointStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	paramsTypeMap["crl_distribution_point_id"] = bindings.NewStringType()
	paramsTypeMap["crlDistributionPointId"] = bindings.NewStringType()
	pathParams["crl_distribution_point_id"] = "crlDistributionPointId"
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
		"/api/v1/trust-management/crl-distribution-points/{crlDistributionPointId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsCertificatesCRLListCrlDistributionPointsInputType() bindings.StructType {
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

func systemAdministrationSettingsCertificatesCRLListCrlDistributionPointsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CrlDistributionPointListBindingType)
}

func systemAdministrationSettingsCertificatesCRLListCrlDistributionPointsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/trust-management/crl-distribution-points",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsCertificatesCRLUpdateCrlDistributionPointInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fields["crl_distribution_point"] = bindings.NewReferenceType(model.CrlDistributionPointBindingType)
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	fieldNameMap["crl_distribution_point"] = "CrlDistributionPoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsCertificatesCRLUpdateCrlDistributionPointOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CrlDistributionPointBindingType)
}

func systemAdministrationSettingsCertificatesCRLUpdateCrlDistributionPointRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["crl_distribution_point_id"] = bindings.NewStringType()
	fields["crl_distribution_point"] = bindings.NewReferenceType(model.CrlDistributionPointBindingType)
	fieldNameMap["crl_distribution_point_id"] = "CrlDistributionPointId"
	fieldNameMap["crl_distribution_point"] = "CrlDistributionPoint"
	paramsTypeMap["crl_distribution_point"] = bindings.NewReferenceType(model.CrlDistributionPointBindingType)
	paramsTypeMap["crl_distribution_point_id"] = bindings.NewStringType()
	paramsTypeMap["crlDistributionPointId"] = bindings.NewStringType()
	pathParams["crl_distribution_point_id"] = "crlDistributionPointId"
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
		"crl_distribution_point",
		"PUT",
		"/api/v1/trust-management/crl-distribution-points/{crlDistributionPointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
