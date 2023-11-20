// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Resources.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package shares

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func resourcesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ResourcesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func resourcesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	paramsTypeMap["share_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shared_resource_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shareId"] = vapiBindings_.NewStringType()
	paramsTypeMap["sharedResourceId"] = vapiBindings_.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
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
		"/policy/api/v1/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ResourcesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
}

func resourcesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	paramsTypeMap["share_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shared_resource_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shareId"] = vapiBindings_.NewStringType()
	paramsTypeMap["sharedResourceId"] = vapiBindings_.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
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
		"/policy/api/v1/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ResourcesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceListResultBindingType)
}

func resourcesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["resource_type"] = "ResourceType"
	paramsTypeMap["share_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["resource_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["shareId"] = vapiBindings_.NewStringType()
	pathParams["share_id"] = "shareId"
	queryParams["resource_type"] = "resource_type"
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
		"/policy/api/v1/infra/shares/{shareId}/resources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fields["shared_resource"] = vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["shared_resource"] = "SharedResource"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ResourcesPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func resourcesPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fields["shared_resource"] = vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["shared_resource"] = "SharedResource"
	paramsTypeMap["share_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shared_resource_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shared_resource"] = vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
	paramsTypeMap["shareId"] = vapiBindings_.NewStringType()
	paramsTypeMap["sharedResourceId"] = vapiBindings_.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
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
		"shared_resource",
		"PATCH",
		"/policy/api/v1/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fields["shared_resource"] = vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["shared_resource"] = "SharedResource"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ResourcesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
}

func resourcesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = vapiBindings_.NewStringType()
	fields["shared_resource_id"] = vapiBindings_.NewStringType()
	fields["shared_resource"] = vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["shared_resource"] = "SharedResource"
	paramsTypeMap["share_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shared_resource_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["shared_resource"] = vapiBindings_.NewReferenceType(nsx_policyModel.SharedResourceBindingType)
	paramsTypeMap["shareId"] = vapiBindings_.NewStringType()
	paramsTypeMap["sharedResourceId"] = vapiBindings_.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
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
		"shared_resource",
		"PUT",
		"/policy/api/v1/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
