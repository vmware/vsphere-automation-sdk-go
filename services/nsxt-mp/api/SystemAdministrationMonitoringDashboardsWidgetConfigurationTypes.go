// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationMonitoringDashboardsWidgetConfiguration.
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

func systemAdministrationMonitoringDashboardsWidgetConfigurationCreateWidgetConfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = bindings.NewStringType()
	fields["widget_configuration"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationCreateWidgetConfigurationOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationCreateWidgetConfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = bindings.NewStringType()
	fields["widget_configuration"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	paramsTypeMap["view_id"] = bindings.NewStringType()
	paramsTypeMap["widget_configuration"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
	paramsTypeMap["viewId"] = bindings.NewStringType()
	pathParams["view_id"] = "viewId"
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
		"widget_configuration",
		"POST",
		"/api/v1/ui-views/{viewId}/widgetconfigurations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationDeleteWidgetConfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = bindings.NewStringType()
	fields["widgetconfiguration_id"] = bindings.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationDeleteWidgetConfigurationOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationDeleteWidgetConfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = bindings.NewStringType()
	fields["widgetconfiguration_id"] = bindings.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	paramsTypeMap["widgetconfiguration_id"] = bindings.NewStringType()
	paramsTypeMap["view_id"] = bindings.NewStringType()
	paramsTypeMap["viewId"] = bindings.NewStringType()
	paramsTypeMap["widgetconfigurationId"] = bindings.NewStringType()
	pathParams["view_id"] = "viewId"
	pathParams["widgetconfiguration_id"] = "widgetconfigurationId"
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
		"/api/v1/ui-views/{viewId}/widgetconfigurations/{widgetconfigurationId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationGetWidgetConfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = bindings.NewStringType()
	fields["widgetconfiguration_id"] = bindings.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationGetWidgetConfigurationOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationGetWidgetConfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = bindings.NewStringType()
	fields["widgetconfiguration_id"] = bindings.NewStringType()
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	paramsTypeMap["widgetconfiguration_id"] = bindings.NewStringType()
	paramsTypeMap["view_id"] = bindings.NewStringType()
	paramsTypeMap["viewId"] = bindings.NewStringType()
	paramsTypeMap["widgetconfigurationId"] = bindings.NewStringType()
	pathParams["view_id"] = "viewId"
	pathParams["widgetconfiguration_id"] = "widgetconfigurationId"
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
		"/api/v1/ui-views/{viewId}/widgetconfigurations/{widgetconfigurationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationListWidgetConfigurationsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = bindings.NewStringType()
	fields["container"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["widget_ids"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["container"] = "Container"
	fieldNameMap["widget_ids"] = "WidgetIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationListWidgetConfigurationsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.WidgetConfigurationListBindingType)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationListWidgetConfigurationsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = bindings.NewStringType()
	fields["container"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["widget_ids"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["container"] = "Container"
	fieldNameMap["widget_ids"] = "WidgetIds"
	paramsTypeMap["view_id"] = bindings.NewStringType()
	paramsTypeMap["container"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["widget_ids"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["viewId"] = bindings.NewStringType()
	pathParams["view_id"] = "viewId"
	queryParams["container"] = "container"
	queryParams["widget_ids"] = "widget_ids"
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
		"/api/v1/ui-views/{viewId}/widgetconfigurations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationUpdateWidgetConfigurationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["view_id"] = bindings.NewStringType()
	fields["widgetconfiguration_id"] = bindings.NewStringType()
	fields["widget_configuration"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationUpdateWidgetConfigurationOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
}

func systemAdministrationMonitoringDashboardsWidgetConfigurationUpdateWidgetConfigurationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["view_id"] = bindings.NewStringType()
	fields["widgetconfiguration_id"] = bindings.NewStringType()
	fields["widget_configuration"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
	fieldNameMap["view_id"] = "ViewId"
	fieldNameMap["widgetconfiguration_id"] = "WidgetconfigurationId"
	fieldNameMap["widget_configuration"] = "WidgetConfiguration"
	paramsTypeMap["widgetconfiguration_id"] = bindings.NewStringType()
	paramsTypeMap["view_id"] = bindings.NewStringType()
	paramsTypeMap["widget_configuration"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.WidgetConfigurationBindingType)}, bindings.REST)
	paramsTypeMap["viewId"] = bindings.NewStringType()
	paramsTypeMap["widgetconfigurationId"] = bindings.NewStringType()
	pathParams["view_id"] = "viewId"
	pathParams["widgetconfiguration_id"] = "widgetconfigurationId"
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
		"widget_configuration",
		"PUT",
		"/api/v1/ui-views/{viewId}/widgetconfigurations/{widgetconfigurationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
