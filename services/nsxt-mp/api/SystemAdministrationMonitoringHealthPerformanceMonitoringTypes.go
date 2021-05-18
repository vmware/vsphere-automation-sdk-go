// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationMonitoringHealthPerformanceMonitoring.
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

func systemAdministrationMonitoringHealthPerformanceMonitoringCollectAlarmsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["page_size"] = "PageSize"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringCollectAlarmsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AlarmListResultBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringCollectAlarmsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["fields"] = "Fields"
	fieldNameMap["page_size"] = "PageSize"
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["fields"] = "fields"
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
		"/api/v1/hpm/alarms",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringHealthPerformanceMonitoringGetAggregationServiceGlobalConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringGetAggregationServiceGlobalConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GlobalCollectionConfigurationBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringGetAggregationServiceGlobalConfigRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/hpm/global-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringHealthPerformanceMonitoringGetFeatureStackConfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["feature_stack_name"] = bindings.NewStringType()
	fieldNameMap["feature_stack_name"] = "FeatureStackName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringGetFeatureStackConfigurationOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FeatureStackCollectionConfigurationBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringGetFeatureStackConfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["feature_stack_name"] = bindings.NewStringType()
	fieldNameMap["feature_stack_name"] = "FeatureStackName"
	paramsTypeMap["feature_stack_name"] = bindings.NewStringType()
	paramsTypeMap["featureStackName"] = bindings.NewStringType()
	pathParams["feature_stack_name"] = "featureStackName"
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
		"/api/v1/hpm/features/{featureStackName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringHealthPerformanceMonitoringListFeatureStackConfigurationsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringListFeatureStackConfigurationsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FeatureStackCollectionConfigurationListBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringListFeatureStackConfigurationsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/hpm/features",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringHealthPerformanceMonitoringResetAggregationServiceFeatureStackConfigurationResetCollectionFrequencyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["feature_stack_name"] = bindings.NewStringType()
	fieldNameMap["feature_stack_name"] = "FeatureStackName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringResetAggregationServiceFeatureStackConfigurationResetCollectionFrequencyOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FeatureStackCollectionConfigurationBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringResetAggregationServiceFeatureStackConfigurationResetCollectionFrequencyRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["feature_stack_name"] = bindings.NewStringType()
	fieldNameMap["feature_stack_name"] = "FeatureStackName"
	paramsTypeMap["feature_stack_name"] = bindings.NewStringType()
	paramsTypeMap["featureStackName"] = bindings.NewStringType()
	pathParams["feature_stack_name"] = "featureStackName"
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
		"action=reset_collection_frequency",
		"",
		"POST",
		"/api/v1/hpm/features/{featureStackName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringHealthPerformanceMonitoringUpdateAggregationServiceGlobalConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["global_collection_configuration"] = bindings.NewReferenceType(model.GlobalCollectionConfigurationBindingType)
	fieldNameMap["global_collection_configuration"] = "GlobalCollectionConfiguration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringUpdateAggregationServiceGlobalConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GlobalCollectionConfigurationBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringUpdateAggregationServiceGlobalConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["global_collection_configuration"] = bindings.NewReferenceType(model.GlobalCollectionConfigurationBindingType)
	fieldNameMap["global_collection_configuration"] = "GlobalCollectionConfiguration"
	paramsTypeMap["global_collection_configuration"] = bindings.NewReferenceType(model.GlobalCollectionConfigurationBindingType)
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
		"global_collection_configuration",
		"PUT",
		"/api/v1/hpm/global-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringHealthPerformanceMonitoringUpdateFeatureStackConfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["feature_stack_name"] = bindings.NewStringType()
	fields["feature_stack_collection_configuration"] = bindings.NewReferenceType(model.FeatureStackCollectionConfigurationBindingType)
	fieldNameMap["feature_stack_name"] = "FeatureStackName"
	fieldNameMap["feature_stack_collection_configuration"] = "FeatureStackCollectionConfiguration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringUpdateFeatureStackConfigurationOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FeatureStackCollectionConfigurationBindingType)
}

func systemAdministrationMonitoringHealthPerformanceMonitoringUpdateFeatureStackConfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["feature_stack_name"] = bindings.NewStringType()
	fields["feature_stack_collection_configuration"] = bindings.NewReferenceType(model.FeatureStackCollectionConfigurationBindingType)
	fieldNameMap["feature_stack_name"] = "FeatureStackName"
	fieldNameMap["feature_stack_collection_configuration"] = "FeatureStackCollectionConfiguration"
	paramsTypeMap["feature_stack_collection_configuration"] = bindings.NewReferenceType(model.FeatureStackCollectionConfigurationBindingType)
	paramsTypeMap["feature_stack_name"] = bindings.NewStringType()
	paramsTypeMap["featureStackName"] = bindings.NewStringType()
	pathParams["feature_stack_name"] = "featureStackName"
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
		"feature_stack_collection_configuration",
		"PUT",
		"/api/v1/hpm/features/{featureStackName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
