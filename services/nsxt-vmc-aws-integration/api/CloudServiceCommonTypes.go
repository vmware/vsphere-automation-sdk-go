// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CloudServiceCommon.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"reflect"
)

// Possible value for ``action`` of method CloudServiceCommon#upgradeVMCAppData.
const CloudServiceCommon_UPGRADE_VMCAPP_DATA_ACTION_UPGRADE = "UPGRADE"

func cloudServiceCommonGetExternalConnectivityConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonGetExternalConnectivityConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExternalConnectivityConfigBindingType)
}

func cloudServiceCommonGetExternalConnectivityConfigRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/external/config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonGetMgmtVmInfoInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm_id"] = bindings.NewStringType()
	fieldNameMap["vm_id"] = "VmId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonGetMgmtVmInfoOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MgmtVmInfoBindingType)
}

func cloudServiceCommonGetMgmtVmInfoRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["vm_id"] = bindings.NewStringType()
	fieldNameMap["vm_id"] = "VmId"
	paramsTypeMap["vm_id"] = bindings.NewStringType()
	paramsTypeMap["vmId"] = bindings.NewStringType()
	pathParams["vm_id"] = "vmId"
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
		"/cloud-service/api/v1/infra/mgmt-vms/{vmId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonGetRealizedEntitiesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["intent_path"] = bindings.NewStringType()
	fields["site_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["site_path"] = "SitePath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonGetRealizedEntitiesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VmcRealizedEntitiesBindingType)
}

func cloudServiceCommonGetRealizedEntitiesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["intent_path"] = bindings.NewStringType()
	fields["site_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["site_path"] = "SitePath"
	paramsTypeMap["site_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["intent_path"] = bindings.NewStringType()
	queryParams["site_path"] = "site_path"
	queryParams["intent_path"] = "intent_path"
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
		"/cloud-service/api/v1/infra/realized-state/realized-entities",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonGetRealizedStateStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["intent_path"] = bindings.NewStringType()
	fields["site_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["site_path"] = "SitePath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonGetRealizedStateStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VmcConsolidatedRealizedStatusBindingType)
}

func cloudServiceCommonGetRealizedStateStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["intent_path"] = bindings.NewStringType()
	fields["site_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["site_path"] = "SitePath"
	paramsTypeMap["site_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["intent_path"] = bindings.NewStringType()
	queryParams["site_path"] = "site_path"
	queryParams["intent_path"] = "intent_path"
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
		"/cloud-service/api/v1/infra/realized-state/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonGetSddcUserConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonGetSddcUserConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SddcUserConfigurationBindingType)
}

func cloudServiceCommonGetSddcUserConfigRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/sddc-user-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonGetVmcFeatureFlagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["internal_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_name"] = "InternalName"
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonGetVmcFeatureFlagsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VmcFeatureFlagsBindingType)
}

func cloudServiceCommonGetVmcFeatureFlagsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["internal_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["internal_name"] = "InternalName"
	fieldNameMap["name"] = "Name"
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["internal_name"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["internal_name"] = "internal_name"
	queryParams["name"] = "name"
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
		"/cloud-service/api/v1/infra/feature-flags",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonListMgmtVmsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonListMgmtVmsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MgmtVmsListResultBindingType)
}

func cloudServiceCommonListMgmtVmsRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/mgmt-vms",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonUpdateIntranetUplinkMtuInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["external_connectivity_config"] = bindings.NewReferenceType(model.ExternalConnectivityConfigBindingType)
	fieldNameMap["external_connectivity_config"] = "ExternalConnectivityConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonUpdateIntranetUplinkMtuOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExternalConnectivityConfigBindingType)
}

func cloudServiceCommonUpdateIntranetUplinkMtuRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["external_connectivity_config"] = bindings.NewReferenceType(model.ExternalConnectivityConfigBindingType)
	fieldNameMap["external_connectivity_config"] = "ExternalConnectivityConfig"
	paramsTypeMap["external_connectivity_config"] = bindings.NewReferenceType(model.ExternalConnectivityConfigBindingType)
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
		"external_connectivity_config",
		"PUT",
		"/cloud-service/api/v1/infra/external/config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceCommonUpgradeVMCAppDataInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upgrade_version_info"] = bindings.NewReferenceType(model.UpgradeVersionInfoBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["upgrade_version_info"] = "UpgradeVersionInfo"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceCommonUpgradeVMCAppDataOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceCommonUpgradeVMCAppDataRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["upgrade_version_info"] = bindings.NewReferenceType(model.UpgradeVersionInfoBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["upgrade_version_info"] = "UpgradeVersionInfo"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["upgrade_version_info"] = bindings.NewReferenceType(model.UpgradeVersionInfoBindingType)
	queryParams["action"] = "action"
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
		"upgrade_version_info",
		"POST",
		"/cloud-service/api/v1/infra/sddc-network-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
