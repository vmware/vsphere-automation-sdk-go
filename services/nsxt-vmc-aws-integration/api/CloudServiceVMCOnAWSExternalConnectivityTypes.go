// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CloudServiceVMCOnAWSExternalConnectivity.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"reflect"
)

// Possible value for ``action`` of method CloudServiceVMCOnAWSExternalConnectivity#attachVif.
const CloudServiceVMCOnAWSExternalConnectivity_ATTACH_VIF_ACTION_ATTACH = "ATTACH"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listAdvertisedExternalRoutes.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_ADVERTISED_EXTERNAL_ROUTES_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listAdvertisedExternalRoutes.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_ADVERTISED_EXTERNAL_ROUTES_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listAdvertisedExternalRoutesInCsvFormatCsv.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_ADVERTISED_EXTERNAL_ROUTES_IN_CSV_FORMAT_CSV_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listAdvertisedExternalRoutesInCsvFormatCsv.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_ADVERTISED_EXTERNAL_ROUTES_IN_CSV_FORMAT_CSV_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listLearnedExternalRoutes.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_LEARNED_EXTERNAL_ROUTES_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listLearnedExternalRoutes.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_LEARNED_EXTERNAL_ROUTES_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listLearnedExternalRoutesInCsvFormatCsv.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_LEARNED_EXTERNAL_ROUTES_IN_CSV_FORMAT_CSV_CONNECTIVITY_TYPE_DIRECT_CONNECT = "DIRECT_CONNECT"

// Possible value for ``connectivityType`` of method CloudServiceVMCOnAWSExternalConnectivity#listLearnedExternalRoutesInCsvFormatCsv.
const CloudServiceVMCOnAWSExternalConnectivity_LIST_LEARNED_EXTERNAL_ROUTES_IN_CSV_FORMAT_CSV_CONNECTIVITY_TYPE_DEPLOYMENT_CONNECTIVITY_GROUP = "DEPLOYMENT_CONNECTIVITY_GROUP"

func cloudServiceVMCOnAWSExternalConnectivityAttachVifInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vif_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewStringType()
	fieldNameMap["vif_id"] = "VifId"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityAttachVifOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSExternalConnectivityAttachVifRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["vif_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewStringType()
	fieldNameMap["vif_id"] = "VifId"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["vif_id"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["vifId"] = bindings.NewStringType()
	pathParams["vif_id"] = "vifId"
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
		"/cloud-service/api/v1/infra/direct-connect/vifs/{vifId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityCreateAssociatedGroupConnectionInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_group_id"] = bindings.NewStringType()
	fields["associated_base_group_connection_info"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.AssociatedBaseGroupConnectionInfoBindingType)}, bindings.REST)
	fieldNameMap["sddc_group_id"] = "SddcGroupId"
	fieldNameMap["associated_base_group_connection_info"] = "AssociatedBaseGroupConnectionInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityCreateAssociatedGroupConnectionInfoOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.AssociatedBaseGroupConnectionInfoBindingType)}, bindings.REST)
}

func cloudServiceVMCOnAWSExternalConnectivityCreateAssociatedGroupConnectionInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["sddc_group_id"] = bindings.NewStringType()
	fields["associated_base_group_connection_info"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.AssociatedBaseGroupConnectionInfoBindingType)}, bindings.REST)
	fieldNameMap["sddc_group_id"] = "SddcGroupId"
	fieldNameMap["associated_base_group_connection_info"] = "AssociatedBaseGroupConnectionInfo"
	paramsTypeMap["sddc_group_id"] = bindings.NewStringType()
	paramsTypeMap["associated_base_group_connection_info"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.AssociatedBaseGroupConnectionInfoBindingType)}, bindings.REST)
	paramsTypeMap["sddcGroupId"] = bindings.NewStringType()
	pathParams["sddc_group_id"] = "sddcGroupId"
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
		"associated_base_group_connection_info",
		"PUT",
		"/cloud-service/api/v1/infra/associated-groups/{sddcGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityCreateDxBgpInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["direct_connect_bgp_info"] = bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
	fieldNameMap["direct_connect_bgp_info"] = "DirectConnectBgpInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityCreateDxBgpOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityCreateDxBgpRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["direct_connect_bgp_info"] = bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
	fieldNameMap["direct_connect_bgp_info"] = "DirectConnectBgpInfo"
	paramsTypeMap["direct_connect_bgp_info"] = bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
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
		"direct_connect_bgp_info",
		"PUT",
		"/cloud-service/api/v1/infra/direct-connect/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityDeleteAssociatedGroupConnectionInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_group_id"] = bindings.NewStringType()
	fieldNameMap["sddc_group_id"] = "SddcGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityDeleteAssociatedGroupConnectionInfoOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSExternalConnectivityDeleteAssociatedGroupConnectionInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["sddc_group_id"] = bindings.NewStringType()
	fieldNameMap["sddc_group_id"] = "SddcGroupId"
	paramsTypeMap["sddc_group_id"] = bindings.NewStringType()
	paramsTypeMap["sddcGroupId"] = bindings.NewStringType()
	pathParams["sddc_group_id"] = "sddcGroupId"
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
		"/cloud-service/api/v1/infra/associated-groups/{sddcGroupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityDeleteVifByIdInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vif_id"] = bindings.NewStringType()
	fieldNameMap["vif_id"] = "VifId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityDeleteVifByIdOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSExternalConnectivityDeleteVifByIdRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["vif_id"] = bindings.NewStringType()
	fieldNameMap["vif_id"] = "VifId"
	paramsTypeMap["vif_id"] = bindings.NewStringType()
	paramsTypeMap["vifId"] = bindings.NewStringType()
	pathParams["vif_id"] = "vifId"
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
		"/cloud-service/api/v1/infra/direct-connect/vifs/{vifId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityGetAssociatedGroupConnectionInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["sddc_group_id"] = bindings.NewStringType()
	fieldNameMap["sddc_group_id"] = "SddcGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityGetAssociatedGroupConnectionInfoOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.AssociatedBaseGroupConnectionInfoBindingType)}, bindings.REST)
}

func cloudServiceVMCOnAWSExternalConnectivityGetAssociatedGroupConnectionInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["sddc_group_id"] = bindings.NewStringType()
	fieldNameMap["sddc_group_id"] = "SddcGroupId"
	paramsTypeMap["sddc_group_id"] = bindings.NewStringType()
	paramsTypeMap["sddcGroupId"] = bindings.NewStringType()
	pathParams["sddc_group_id"] = "sddcGroupId"
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
		"/cloud-service/api/v1/infra/associated-groups/{sddcGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityGetDxBgpInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityGetDxBgpInfoOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityGetDxBgpInfoRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/direct-connect/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExternalSddcRoutesListResultBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["connectivity_type"] = "connectivity_type"
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
		"/cloud-service/api/v1/infra/external/routes/advertised",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInCsvFormatCsvInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInCsvFormatCsvOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExternalSddcRoutesListResultInCsvFormatBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedExternalRoutesInCsvFormatCsvRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["connectivity_type"] = "connectivity_type"
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
		"format=csv",
		"",
		"GET",
		"/cloud-service/api/v1/infra/external/routes/advertised",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPAdvertisedRoutesBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListAdvertisedRoutesRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/direct-connect/routes/advertised",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListAssociatedGroupConnectionInfosInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListAssociatedGroupConnectionInfosOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AssociatedGroupConnectionInfosListResultBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListAssociatedGroupConnectionInfosRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/associated-groups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListDirectConnectVifsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListDirectConnectVifsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VifsListResultBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListDirectConnectVifsRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/direct-connect/vifs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExternalSddcRoutesListResultBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["connectivity_type"] = "connectivity_type"
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
		"/cloud-service/api/v1/infra/external/routes/learned",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInCsvFormatCsvInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInCsvFormatCsvOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExternalSddcRoutesListResultInCsvFormatBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedExternalRoutesInCsvFormatCsvRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["connectivity_type"] = "ConnectivityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["connectivity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["connectivity_type"] = "connectivity_type"
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
		"format=csv",
		"",
		"GET",
		"/cloud-service/api/v1/infra/external/routes/learned",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPLearnedRoutesBindingType)
}

func cloudServiceVMCOnAWSExternalConnectivityListLearnedRoutesRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/direct-connect/routes/learned",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSExternalConnectivityUpdateDxBgpInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["direct_connect_bgp_info"] = bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
	fieldNameMap["direct_connect_bgp_info"] = "DirectConnectBgpInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSExternalConnectivityUpdateDxBgpInfoOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSExternalConnectivityUpdateDxBgpInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["direct_connect_bgp_info"] = bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
	fieldNameMap["direct_connect_bgp_info"] = "DirectConnectBgpInfo"
	paramsTypeMap["direct_connect_bgp_info"] = bindings.NewReferenceType(model.DirectConnectBgpInfoBindingType)
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
		"direct_connect_bgp_info",
		"PATCH",
		"/cloud-service/api/v1/infra/direct-connect/bgp",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
