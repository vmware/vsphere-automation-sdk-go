// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: GlobalConfigs.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func globalConfigsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_type"] = bindings.NewStringType()
	fieldNameMap["config_type"] = "ConfigType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalConfigsGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
}

func globalConfigsGetRestMetadata() protocol.OperationRestMetadata {
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

func globalConfigsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalConfigsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GlobalConfigsListResultBindingType)
}

func globalConfigsListRestMetadata() protocol.OperationRestMetadata {
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

func globalConfigsResyncconfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_type"] = bindings.NewStringType()
	fields["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	fieldNameMap["config_type"] = "ConfigType"
	fieldNameMap["global_configs"] = "GlobalConfigs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalConfigsResyncconfigOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
}

func globalConfigsResyncconfigRestMetadata() protocol.OperationRestMetadata {
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

func globalConfigsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_type"] = bindings.NewStringType()
	fields["global_configs"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
	fieldNameMap["config_type"] = "ConfigType"
	fieldNameMap["global_configs"] = "GlobalConfigs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func globalConfigsUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.GlobalConfigsBindingType)}, bindings.REST)
}

func globalConfigsUpdateRestMetadata() protocol.OperationRestMetadata {
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
