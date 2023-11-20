// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SupportBundles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package administration

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func supportBundlesCollectInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["support_bundle_request"] = vapiBindings_.NewReferenceType(nsxModel.SupportBundleRequestBindingType)
	fields["override_async_response"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["require_delete_or_override_async_response"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["support_bundle_request"] = "SupportBundleRequest"
	fieldNameMap["override_async_response"] = "OverrideAsyncResponse"
	fieldNameMap["require_delete_or_override_async_response"] = "RequireDeleteOrOverrideAsyncResponse"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SupportBundlesCollectOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.SupportBundleResultBindingType)
}

func supportBundlesCollectRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["support_bundle_request"] = vapiBindings_.NewReferenceType(nsxModel.SupportBundleRequestBindingType)
	fields["override_async_response"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["require_delete_or_override_async_response"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["support_bundle_request"] = "SupportBundleRequest"
	fieldNameMap["override_async_response"] = "OverrideAsyncResponse"
	fieldNameMap["require_delete_or_override_async_response"] = "RequireDeleteOrOverrideAsyncResponse"
	paramsTypeMap["override_async_response"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["require_delete_or_override_async_response"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["support_bundle_request"] = vapiBindings_.NewReferenceType(nsxModel.SupportBundleRequestBindingType)
	queryParams["override_async_response"] = "override_async_response"
	queryParams["require_delete_or_override_async_response"] = "require_delete_or_override_async_response"
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

func supportBundlesDeleteasyncresponseInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SupportBundlesDeleteasyncresponseOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func supportBundlesDeleteasyncresponseRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
