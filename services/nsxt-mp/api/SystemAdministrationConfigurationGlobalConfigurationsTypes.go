// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationGlobalConfigurations.
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

func systemAdministrationConfigurationGlobalConfigurationsGetGlobalConfigsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_type"] = bindings.NewStringType()
	fieldNameMap["config_type"] = "ConfigType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsGetGlobalConfigsOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationGlobalConfigurationsGetGlobalConfigsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["config_type"] = bindings.NewStringType()
	fieldNameMap["config_type"] = "ConfigType"
	paramsTypeMap["config_type"] = bindings.NewStringType()
	paramsTypeMap["configType"] = bindings.NewStringType()
	pathParams["config_type"] = "configType"
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
		"/api/v1/global-configs/{configType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationGlobalConfigurationsListCentralNodeConfigProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsListCentralNodeConfigProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CentralNodeConfigProfileListResultBindingType)
}

func systemAdministrationConfigurationGlobalConfigurationsListCentralNodeConfigProfilesRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/configs/central-config/node-config-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationGlobalConfigurationsListGlobalConfigsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsListGlobalConfigsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GlobalConfigsListResultBindingType)
}

func systemAdministrationConfigurationGlobalConfigurationsListGlobalConfigsRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/global-configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationGlobalConfigurationsReadCentralNodeConfigProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile_id"] = bindings.NewStringType()
	fields["show_sensitive_data"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["profile_id"] = "ProfileId"
	fieldNameMap["show_sensitive_data"] = "ShowSensitiveData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsReadCentralNodeConfigProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
}

func systemAdministrationConfigurationGlobalConfigurationsReadCentralNodeConfigProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["profile_id"] = bindings.NewStringType()
	fields["show_sensitive_data"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["profile_id"] = "ProfileId"
	fieldNameMap["show_sensitive_data"] = "ShowSensitiveData"
	paramsTypeMap["profile_id"] = bindings.NewStringType()
	paramsTypeMap["show_sensitive_data"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["profileId"] = bindings.NewStringType()
	pathParams["profile_id"] = "profileId"
	queryParams["show_sensitive_data"] = "show_sensitive_data"
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
		"/api/v1/configs/central-config/node-config-profiles/{profileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationGlobalConfigurationsResyncGlobalConfigsResyncConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_type"] = bindings.NewStringType()
	fields["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	fieldNameMap["config_type"] = "ConfigType"
	fieldNameMap["global_configs"] = "GlobalConfigs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsResyncGlobalConfigsResyncConfigOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationGlobalConfigurationsResyncGlobalConfigsResyncConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["config_type"] = bindings.NewStringType()
	fields["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	fieldNameMap["config_type"] = "ConfigType"
	fieldNameMap["global_configs"] = "GlobalConfigs"
	paramsTypeMap["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	paramsTypeMap["config_type"] = bindings.NewStringType()
	paramsTypeMap["configType"] = bindings.NewStringType()
	pathParams["config_type"] = "configType"
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
		"action=resync_config",
		"global_configs",
		"PUT",
		"/api/v1/global-configs/{configType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationGlobalConfigurationsUpdateCentralNodeConfigProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_config_profile_id"] = bindings.NewStringType()
	fields["central_node_config_profile"] = bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
	fieldNameMap["node_config_profile_id"] = "NodeConfigProfileId"
	fieldNameMap["central_node_config_profile"] = "CentralNodeConfigProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsUpdateCentralNodeConfigProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
}

func systemAdministrationConfigurationGlobalConfigurationsUpdateCentralNodeConfigProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_config_profile_id"] = bindings.NewStringType()
	fields["central_node_config_profile"] = bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
	fieldNameMap["node_config_profile_id"] = "NodeConfigProfileId"
	fieldNameMap["central_node_config_profile"] = "CentralNodeConfigProfile"
	paramsTypeMap["central_node_config_profile"] = bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
	paramsTypeMap["node_config_profile_id"] = bindings.NewStringType()
	paramsTypeMap["nodeConfigProfileId"] = bindings.NewStringType()
	pathParams["node_config_profile_id"] = "nodeConfigProfileId"
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
		"central_node_config_profile",
		"PUT",
		"/api/v1/configs/central-config/node-config-profiles/{nodeConfigProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationGlobalConfigurationsUpdateGlobalConfigsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_type"] = bindings.NewStringType()
	fields["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	fieldNameMap["config_type"] = "ConfigType"
	fieldNameMap["global_configs"] = "GlobalConfigs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationGlobalConfigurationsUpdateGlobalConfigsOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
}

func systemAdministrationConfigurationGlobalConfigurationsUpdateGlobalConfigsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["config_type"] = bindings.NewStringType()
	fields["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	fieldNameMap["config_type"] = "ConfigType"
	fieldNameMap["global_configs"] = "GlobalConfigs"
	paramsTypeMap["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	paramsTypeMap["config_type"] = bindings.NewStringType()
	paramsTypeMap["configType"] = bindings.NewStringType()
	pathParams["config_type"] = "configType"
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
		"global_configs",
		"PUT",
		"/api/v1/global-configs/{configType}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
