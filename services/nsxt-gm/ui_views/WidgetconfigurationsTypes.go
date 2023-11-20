// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Widgetconfigurations.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package ui_views

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func widgetconfigurationsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widget_configuration"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func WidgetconfigurationsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
}

func widgetconfigurationsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widget_configuration"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	paramsTypeMap["widget_configuration"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
	paramsTypeMap["view_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["viewId"] = vapiBindings_.NewStringType()
	pathParams["view_id"] = "viewId"
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
		"widget_configuration",
		"POST",
		"/global-manager/api/v1/ui-views/{viewId}/widgetconfigurations",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func widgetconfigurationsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func WidgetconfigurationsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func widgetconfigurationsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	paramsTypeMap["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["view_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["viewId"] = vapiBindings_.NewStringType()
	paramsTypeMap["widgetconfigurationId"] = vapiBindings_.NewStringType()
	pathParams["view_id"] = "viewId"
	pathParams["widgetconfiguration_id"] = "widgetconfigurationId"
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
		"/global-manager/api/v1/ui-views/{viewId}/widgetconfigurations/{widgetconfigurationId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func widgetconfigurationsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["container"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["widget_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["container"] = "Container"
	fieldNameMap["widget_ids"] = "WidgetIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func WidgetconfigurationsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationListBindingType)
}

func widgetconfigurationsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["container"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["widget_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["container"] = "Container"
	fieldNameMap["widget_ids"] = "WidgetIds"
	paramsTypeMap["container"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["view_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["widget_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["viewId"] = vapiBindings_.NewStringType()
	pathParams["view_id"] = "viewId"
	queryParams["container"] = "container"
	queryParams["widget_ids"] = "widget_ids"
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
		"/global-manager/api/v1/ui-views/{viewId}/widgetconfigurations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func widgetconfigurationsGet0InputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func WidgetconfigurationsGet0OutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
}

func widgetconfigurationsGet0RestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	paramsTypeMap["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["view_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["viewId"] = vapiBindings_.NewStringType()
	paramsTypeMap["widgetconfigurationId"] = vapiBindings_.NewStringType()
	pathParams["view_id"] = "viewId"
	pathParams["widgetconfiguration_id"] = "widgetconfigurationId"
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
		"/global-manager/api/v1/ui-views/{viewId}/widgetconfigurations/{widgetconfigurationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func widgetconfigurationsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	fields["widget_configuration"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func WidgetconfigurationsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
}

func widgetconfigurationsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = vapiBindings_.NewStringType()
	fields["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	fields["widget_configuration"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	paramsTypeMap["widget_configuration"] = vapiBindings_.NewDynamicStructType([]vapiBindings_.ReferenceType{vapiBindings_.NewReferenceType(nsx_global_policyModel.WidgetConfigurationBindingType)})
	paramsTypeMap["widgetconfiguration_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["view_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["viewId"] = vapiBindings_.NewStringType()
	paramsTypeMap["widgetconfigurationId"] = vapiBindings_.NewStringType()
	pathParams["view_id"] = "viewId"
	pathParams["widgetconfiguration_id"] = "widgetconfigurationId"
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
		"widget_configuration",
		"PUT",
		"/global-manager/api/v1/ui-views/{viewId}/widgetconfigurations/{widgetconfigurationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
