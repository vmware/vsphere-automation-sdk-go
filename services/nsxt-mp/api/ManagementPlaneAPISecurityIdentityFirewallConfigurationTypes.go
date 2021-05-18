// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPISecurityIdentityFirewallConfiguration.
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

func managementPlaneAPISecurityIdentityFirewallConfigurationDeleteEnabledComputeCollectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationDeleteEnabledComputeCollectionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPISecurityIdentityFirewallConfigurationDeleteEnabledComputeCollectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
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
		"/api/v1/idfw/idfw-compute-collections/{ccExtId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetEnabledComputeCollectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetEnabledComputeCollectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwEnabledComputeCollectionBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetEnabledComputeCollectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
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
		"/api/v1/idfw/idfw-compute-collections/{ccExtId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetMasterSwitchSettingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetMasterSwitchSettingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwMasterSwitchSettingBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetMasterSwitchSettingRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/idfw/master-switch-setting",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetStandaloneHostsSwitchSettingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetStandaloneHostsSwitchSettingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwStandaloneHostsSwitchSettingBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationGetStandaloneHostsSwitchSettingRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/idfw/standalone-host-switch-setting",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationListEnabledComputeCollectionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationListEnabledComputeCollectionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwEnabledComputeCollectionListResultBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationListEnabledComputeCollectionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/idfw/idfw-compute-collections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateEnabledComputeCollectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fields["idfw_enabled_compute_collection"] = bindings.NewReferenceType(model.IdfwEnabledComputeCollectionBindingType)
	fieldNameMap["cc_ext_id"] = "CcExtId"
	fieldNameMap["idfw_enabled_compute_collection"] = "IdfwEnabledComputeCollection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateEnabledComputeCollectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwEnabledComputeCollectionBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateEnabledComputeCollectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fields["idfw_enabled_compute_collection"] = bindings.NewReferenceType(model.IdfwEnabledComputeCollectionBindingType)
	fieldNameMap["cc_ext_id"] = "CcExtId"
	fieldNameMap["idfw_enabled_compute_collection"] = "IdfwEnabledComputeCollection"
	paramsTypeMap["idfw_enabled_compute_collection"] = bindings.NewReferenceType(model.IdfwEnabledComputeCollectionBindingType)
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
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
		"idfw_enabled_compute_collection",
		"PUT",
		"/api/v1/idfw/idfw-compute-collections/{ccExtId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateMasterSwitchSettingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["idfw_master_switch_setting"] = bindings.NewReferenceType(model.IdfwMasterSwitchSettingBindingType)
	fieldNameMap["idfw_master_switch_setting"] = "IdfwMasterSwitchSetting"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateMasterSwitchSettingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwMasterSwitchSettingBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateMasterSwitchSettingRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["idfw_master_switch_setting"] = bindings.NewReferenceType(model.IdfwMasterSwitchSettingBindingType)
	fieldNameMap["idfw_master_switch_setting"] = "IdfwMasterSwitchSetting"
	paramsTypeMap["idfw_master_switch_setting"] = bindings.NewReferenceType(model.IdfwMasterSwitchSettingBindingType)
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
		"idfw_master_switch_setting",
		"PUT",
		"/api/v1/idfw/master-switch-setting",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateStandaloneHostsSwitchSettingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["idfw_standalone_hosts_switch_setting"] = bindings.NewReferenceType(model.IdfwStandaloneHostsSwitchSettingBindingType)
	fieldNameMap["idfw_standalone_hosts_switch_setting"] = "IdfwStandaloneHostsSwitchSetting"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateStandaloneHostsSwitchSettingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdfwStandaloneHostsSwitchSettingBindingType)
}

func managementPlaneAPISecurityIdentityFirewallConfigurationUpdateStandaloneHostsSwitchSettingRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["idfw_standalone_hosts_switch_setting"] = bindings.NewReferenceType(model.IdfwStandaloneHostsSwitchSettingBindingType)
	fieldNameMap["idfw_standalone_hosts_switch_setting"] = "IdfwStandaloneHostsSwitchSetting"
	paramsTypeMap["idfw_standalone_hosts_switch_setting"] = bindings.NewReferenceType(model.IdfwStandaloneHostsSwitchSettingBindingType)
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
		"idfw_standalone_hosts_switch_setting",
		"PUT",
		"/api/v1/idfw/standalone-host-switch-setting",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
