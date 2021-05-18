// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationSettingsSupportBundle.
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

func systemAdministrationSettingsSupportBundleCollectSupportBundlesCollectInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["support_bundle_request"] = bindings.NewReferenceType(model.SupportBundleRequestBindingType)
	fields["override_async_response"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["require_delete_or_override_async_response"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["support_bundle_request"] = "SupportBundleRequest"
	fieldNameMap["override_async_response"] = "OverrideAsyncResponse"
	fieldNameMap["require_delete_or_override_async_response"] = "RequireDeleteOrOverrideAsyncResponse"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsSupportBundleCollectSupportBundlesCollectOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SupportBundleResultBindingType)
}

func systemAdministrationSettingsSupportBundleCollectSupportBundlesCollectRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["support_bundle_request"] = bindings.NewReferenceType(model.SupportBundleRequestBindingType)
	fields["override_async_response"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["require_delete_or_override_async_response"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["support_bundle_request"] = "SupportBundleRequest"
	fieldNameMap["override_async_response"] = "OverrideAsyncResponse"
	fieldNameMap["require_delete_or_override_async_response"] = "RequireDeleteOrOverrideAsyncResponse"
	paramsTypeMap["override_async_response"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["require_delete_or_override_async_response"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["support_bundle_request"] = bindings.NewReferenceType(model.SupportBundleRequestBindingType)
	queryParams["override_async_response"] = "override_async_response"
	queryParams["require_delete_or_override_async_response"] = "require_delete_or_override_async_response"
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
		"action=collect",
		"support_bundle_request",
		"POST",
		"/api/v1/administration/support-bundles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationSettingsSupportBundleDeleteSupportBundlesAsyncResponseDeleteAsyncResponseInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationSettingsSupportBundleDeleteSupportBundlesAsyncResponseDeleteAsyncResponseOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationSettingsSupportBundleDeleteSupportBundlesAsyncResponseDeleteAsyncResponseRestMetadata() protocol.OperationRestMetadata {
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
		"action=delete_async_response",
		"",
		"POST",
		"/api/v1/administration/support-bundles",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
