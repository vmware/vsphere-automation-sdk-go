// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationLifecycleManagementUpgradeBundles.
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

func systemAdministrationLifecycleManagementUpgradeBundlesCancelUpgradeBundleUploadCancelUploadInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementUpgradeBundlesCancelUpgradeBundleUploadCancelUploadOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationLifecycleManagementUpgradeBundlesCancelUpgradeBundleUploadCancelUploadRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	paramsTypeMap["bundle_id"] = bindings.NewStringType()
	paramsTypeMap["bundleId"] = bindings.NewStringType()
	pathParams["bundle_id"] = "bundleId"
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
		"/api/v1/upgrade/bundles/{bundleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementUpgradeBundlesFetchUpgradeBundleFromUrlInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upgrade_bundle_fetch_request"] = bindings.NewReferenceType(model.UpgradeBundleFetchRequestBindingType)
	fieldNameMap["upgrade_bundle_fetch_request"] = "UpgradeBundleFetchRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementUpgradeBundlesFetchUpgradeBundleFromUrlOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeBundleIdBindingType)
}

func systemAdministrationLifecycleManagementUpgradeBundlesFetchUpgradeBundleFromUrlRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["upgrade_bundle_fetch_request"] = bindings.NewReferenceType(model.UpgradeBundleFetchRequestBindingType)
	fieldNameMap["upgrade_bundle_fetch_request"] = "UpgradeBundleFetchRequest"
	paramsTypeMap["upgrade_bundle_fetch_request"] = bindings.NewReferenceType(model.UpgradeBundleFetchRequestBindingType)
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
		"upgrade_bundle_fetch_request",
		"POST",
		"/api/v1/upgrade/bundles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleInfoOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeBundleInfoBindingType)
}

func systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	paramsTypeMap["bundle_id"] = bindings.NewStringType()
	paramsTypeMap["bundleId"] = bindings.NewStringType()
	pathParams["bundle_id"] = "bundleId"
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
		"/api/v1/upgrade/bundles/{bundleId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleUploadStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleUploadStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeBundleUploadStatusBindingType)
}

func systemAdministrationLifecycleManagementUpgradeBundlesGetUpgradeBundleUploadStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	paramsTypeMap["bundle_id"] = bindings.NewStringType()
	paramsTypeMap["bundleId"] = bindings.NewStringType()
	pathParams["bundle_id"] = "bundleId"
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
		"/api/v1/upgrade/bundles/{bundleId}/upload-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationLifecycleManagementUpgradeBundlesTriggerUcUpgradeUpgradeUcInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationLifecycleManagementUpgradeBundlesTriggerUcUpgradeUpgradeUcOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationLifecycleManagementUpgradeBundlesTriggerUcUpgradeUpgradeUcRestMetadata() protocol.OperationRestMetadata {
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
		"action=upgrade_uc",
		"",
		"POST",
		"/api/v1/upgrade",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
