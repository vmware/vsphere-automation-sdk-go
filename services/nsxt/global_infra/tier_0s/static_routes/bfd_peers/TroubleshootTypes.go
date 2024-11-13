// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Troubleshoot.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package bfd_peers

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func troubleshootGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
}

func troubleshootGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
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
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/static-routes/bfd-peers/troubleshoot",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func troubleshootPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["static_route_bfd_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
	fields["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["static_route_bfd_troubleshoot_config"] = "StaticRouteBfdTroubleshootConfig"
	fieldNameMap["override"] = "Override"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func troubleshootPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["static_route_bfd_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
	fields["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["static_route_bfd_troubleshoot_config"] = "StaticRouteBfdTroubleshootConfig"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["static_route_bfd_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	queryParams["override"] = "override"
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
		"static_route_bfd_troubleshoot_config",
		"PATCH",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/static-routes/bfd-peers/troubleshoot",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func troubleshootUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["static_route_bfd_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
	fields["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["static_route_bfd_troubleshoot_config"] = "StaticRouteBfdTroubleshootConfig"
	fieldNameMap["override"] = "Override"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TroubleshootUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
}

func troubleshootUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["static_route_bfd_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
	fields["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["static_route_bfd_troubleshoot_config"] = "StaticRouteBfdTroubleshootConfig"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["override"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["static_route_bfd_troubleshoot_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.StaticRouteBfdTroubleshootConfigBindingType)
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	queryParams["override"] = "override"
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
		"static_route_bfd_troubleshoot_config",
		"PUT",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/static-routes/bfd-peers/troubleshoot",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
