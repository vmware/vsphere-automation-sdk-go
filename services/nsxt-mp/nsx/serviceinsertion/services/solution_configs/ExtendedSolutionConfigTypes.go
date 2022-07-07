// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ExtendedSolutionConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package solution_configs

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func extendedSolutionConfigCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fields["extended_solution_config"] = bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func extendedSolutionConfigCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
}

func extendedSolutionConfigCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fields["extended_solution_config"] = bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["extended_solution_config"] = bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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

func extendedSolutionConfigDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func extendedSolutionConfigDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func extendedSolutionConfigDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}/extended-solution-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func extendedSolutionConfigGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func extendedSolutionConfigGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
}

func extendedSolutionConfigGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/solution-configs/{solutionConfigId}/extended-solution-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func extendedSolutionConfigUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fields["extended_solution_config"] = bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func extendedSolutionConfigUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
}

func extendedSolutionConfigUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["solution_config_id"] = bindings.NewStringType()
	fields["extended_solution_config"] = bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["solution_config_id"] = "SolutionConfigId"
	fieldNameMap["extended_solution_config"] = "ExtendedSolutionConfig"
	paramsTypeMap["solution_config_id"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["extended_solution_config"] = bindings.NewReferenceType(model.ExtendedSolutionConfigBindingType)
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["solutionConfigId"] = bindings.NewStringType()
	pathParams["solution_config_id"] = "solutionConfigId"
	pathParams["service_id"] = "serviceId"
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
