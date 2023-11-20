// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ExtendedSolutionConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package solution_configs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func extendedSolutionConfigCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fields["extended_solution_config"] = vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExtendedSolutionConfigCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
}

func extendedSolutionConfigCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fields["extended_solution_config"] = vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	paramsTypeMap["extended_solution_config"] = vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
	paramsTypeMap["solution_config_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["solutionConfigId"] = vapiBindings_.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"extended_solution_config",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}/extended-solution-config",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func extendedSolutionConfigDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExtendedSolutionConfigDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func extendedSolutionConfigDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	paramsTypeMap["solution_config_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["solutionConfigId"] = vapiBindings_.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}/extended-solution-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func extendedSolutionConfigGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExtendedSolutionConfigGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
}

func extendedSolutionConfigGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	paramsTypeMap["solution_config_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["solutionConfigId"] = vapiBindings_.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}/extended-solution-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func extendedSolutionConfigUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fields["extended_solution_config"] = vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExtendedSolutionConfigUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
}

func extendedSolutionConfigUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["solution_config_id"] = vapiBindings_.NewStringType()
	fields["extended_solution_config"] = vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	paramsTypeMap["extended_solution_config"] = vapiBindings_.NewReferenceType(nsxModel.ExtendedSolutionConfigBindingType)
	paramsTypeMap["solution_config_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["solutionConfigId"] = vapiBindings_.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"extended_solution_config",
		"PUT",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}/extended-solution-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
