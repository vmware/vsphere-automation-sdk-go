// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles.
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

// Possible value for ``product`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#cancelBundleUploadCancelUpload.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_CANCEL_BUNDLE_UPLOAD_CANCEL_UPLOAD_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``fileType`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#getBundleIds.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_GET_BUNDLE_IDS_FILE_TYPE_OVA = "OVA"

// Possible value for ``product`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#getBundleIds.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_GET_BUNDLE_IDS_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``product`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#getBundleUploadPermissions.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_GET_BUNDLE_UPLOAD_PERMISSIONS_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``product`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#getBundleUploadStatus.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_GET_BUNDLE_UPLOAD_STATUS_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``product`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#getOvfDeployInfo.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_GET_OVF_DEPLOY_INFO_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``fileType`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#uploadBundleViaRemoteFile.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_UPLOAD_BUNDLE_VIA_REMOTE_FILE_FILE_TYPE_OVA = "OVA"

// Possible value for ``product`` of method SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles#uploadBundleViaRemoteFile.
const SystemAdministrationConfigurationNSXIntelligenceRepositoryBundles_UPLOAD_BUNDLE_VIA_REMOTE_FILE_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesCancelBundleUploadCancelUploadInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesCancelBundleUploadCancelUploadOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesCancelBundleUploadCancelUploadRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["bundle_id"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	paramsTypeMap["bundleId"] = bindings.NewStringType()
	pathParams["bundle_id"] = "bundleId"
	queryParams["product"] = "product"
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
		"action=cancel_upload",
		"",
		"POST",
		"/api/v1/repository/bundles/{bundleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleIdsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleIdsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BundleIdsBindingType)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleIdsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["file_type"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	queryParams["product"] = "product"
	queryParams["file_type"] = "file_type"
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
		"/api/v1/repository/bundles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadPermissionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadPermissionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BundleUploadPermissionBindingType)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadPermissionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	paramsTypeMap["product"] = bindings.NewStringType()
	queryParams["product"] = "product"
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
		"/api/v1/repository/bundles/upload-allowed",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BundleUploadStatusBindingType)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetBundleUploadStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["bundle_id"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	paramsTypeMap["bundleId"] = bindings.NewStringType()
	pathParams["bundle_id"] = "bundleId"
	queryParams["product"] = "product"
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
		"/api/v1/repository/bundles/{bundleId}/upload-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetOvfDeployInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetOvfDeployInfoOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.OvfInfoBindingType)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesGetOvfDeployInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["product"] = bindings.NewStringType()
	fieldNameMap["product"] = "Product"
	paramsTypeMap["product"] = bindings.NewStringType()
	queryParams["product"] = "product"
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
		"/api/v1/repository/bundles/ovf-deploy-info",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesUploadBundleViaRemoteFileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["remote_bundle_url"] = bindings.NewReferenceType(model.RemoteBundleUrlBindingType)
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["remote_bundle_url"] = "RemoteBundleUrl"
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesUploadBundleViaRemoteFileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BundleIdBindingType)
}

func systemAdministrationConfigurationNSXIntelligenceRepositoryBundlesUploadBundleViaRemoteFileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["remote_bundle_url"] = bindings.NewReferenceType(model.RemoteBundleUrlBindingType)
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["remote_bundle_url"] = "RemoteBundleUrl"
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["remote_bundle_url"] = bindings.NewReferenceType(model.RemoteBundleUrlBindingType)
	paramsTypeMap["file_type"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	queryParams["product"] = "product"
	queryParams["file_type"] = "file_type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"remote_bundle_url",
		"POST",
		"/api/v1/repository/bundles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
