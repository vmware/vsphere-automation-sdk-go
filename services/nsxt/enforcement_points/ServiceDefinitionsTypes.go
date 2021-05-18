// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ServiceDefinitions.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package enforcement_points

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func serviceDefinitionsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition"] = "ServiceDefinition"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceDefinitionsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDefinitionBindingType)
}

func serviceDefinitionsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition"] = "ServiceDefinition"
	paramsTypeMap["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	paramsTypeMap["enforcement_point_id"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointId"] = bindings.NewStringType()
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"service_definition",
		"POST",
		"/policy/api/v1/enforcement-points/{enforcementPointId}/service-definitions",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceDefinitionsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition_id"] = "ServiceDefinitionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceDefinitionsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceDefinitionsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition_id"] = "ServiceDefinitionId"
	paramsTypeMap["enforcement_point_id"] = bindings.NewStringType()
	paramsTypeMap["service_definition_id"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointId"] = bindings.NewStringType()
	paramsTypeMap["serviceDefinitionId"] = bindings.NewStringType()
	pathParams["service_definition_id"] = "serviceDefinitionId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"/policy/api/v1/enforcement-points/{enforcementPointId}/service-definitions/{serviceDefinitionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceDefinitionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition_id"] = "ServiceDefinitionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceDefinitionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDefinitionBindingType)
}

func serviceDefinitionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition_id"] = "ServiceDefinitionId"
	paramsTypeMap["enforcement_point_id"] = bindings.NewStringType()
	paramsTypeMap["service_definition_id"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointId"] = bindings.NewStringType()
	paramsTypeMap["serviceDefinitionId"] = bindings.NewStringType()
	pathParams["service_definition_id"] = "serviceDefinitionId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"/policy/api/v1/enforcement-points/{enforcementPointId}/service-definitions/{serviceDefinitionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceDefinitionsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enforcement_point_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceDefinitionsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInsertionServiceListResultBindingType)
}

func serviceDefinitionsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["enforcement_point_id"] = bindings.NewStringType()
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	paramsTypeMap["enforcement_point_id"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointId"] = bindings.NewStringType()
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"/policy/api/v1/enforcement-points/{enforcementPointId}/service-definitions",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceDefinitionsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition_id"] = bindings.NewStringType()
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition_id"] = "ServiceDefinitionId"
	fieldNameMap["service_definition"] = "ServiceDefinition"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceDefinitionsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceDefinitionBindingType)
}

func serviceDefinitionsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["service_definition_id"] = bindings.NewStringType()
	fields["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["service_definition_id"] = "ServiceDefinitionId"
	fieldNameMap["service_definition"] = "ServiceDefinition"
	paramsTypeMap["service_definition"] = bindings.NewReferenceType(model.ServiceDefinitionBindingType)
	paramsTypeMap["enforcement_point_id"] = bindings.NewStringType()
	paramsTypeMap["service_definition_id"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointId"] = bindings.NewStringType()
	paramsTypeMap["serviceDefinitionId"] = bindings.NewStringType()
	pathParams["service_definition_id"] = "serviceDefinitionId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
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
		"service_definition",
		"PUT",
		"/policy/api/v1/enforcement-points/{enforcementPointId}/service-definitions/{serviceDefinitionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
