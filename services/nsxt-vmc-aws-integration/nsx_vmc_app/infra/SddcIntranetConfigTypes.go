// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SddcIntranetConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package infra

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
	"reflect"
)

func sddcIntranetConfigDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["intranet_config_id"] = vapiBindings_.NewStringType()
	fieldNameMap["intranet_config_id"] = "IntranetConfigId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SddcIntranetConfigDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sddcIntranetConfigDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["intranet_config_id"] = vapiBindings_.NewStringType()
	fieldNameMap["intranet_config_id"] = "IntranetConfigId"
	paramsTypeMap["intranet_config_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["intranetConfigId"] = vapiBindings_.NewStringType()
	pathParams["intranet_config_id"] = "intranetConfigId"
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
		"/cloud-service/api/v1/infra/sddc-intranet-config/{intranetConfigId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sddcIntranetConfigGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["intranet_configuration_id"] = vapiBindings_.NewStringType()
	fieldNameMap["intranet_configuration_id"] = "IntranetConfigurationId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SddcIntranetConfigGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.IntranetConnectionConfigurationBindingType)
}

func sddcIntranetConfigGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["intranet_configuration_id"] = vapiBindings_.NewStringType()
	fieldNameMap["intranet_configuration_id"] = "IntranetConfigurationId"
	paramsTypeMap["intranet_configuration_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["intranetConfigurationId"] = vapiBindings_.NewStringType()
	pathParams["intranet_configuration_id"] = "intranetConfigurationId"
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
		"/cloud-service/api/v1/infra/sddc-intranet-config/{intranetConfigurationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sddcIntranetConfigListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SddcIntranetConfigListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.IntranetConnectionConfigurationListResultBindingType)
}

func sddcIntranetConfigListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/cloud-service/api/v1/infra/sddc-intranet-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sddcIntranetConfigUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["intranet_configuration_id"] = vapiBindings_.NewStringType()
	fields["intranet_connection_configuration"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.IntranetConnectionConfigurationBindingType)
	fieldNameMap["intranet_configuration_id"] = "IntranetConfigurationId"
	fieldNameMap["intranet_connection_configuration"] = "IntranetConnectionConfiguration"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SddcIntranetConfigUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.IntranetConnectionConfigurationBindingType)
}

func sddcIntranetConfigUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["intranet_configuration_id"] = vapiBindings_.NewStringType()
	fields["intranet_connection_configuration"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.IntranetConnectionConfigurationBindingType)
	fieldNameMap["intranet_configuration_id"] = "IntranetConfigurationId"
	fieldNameMap["intranet_connection_configuration"] = "IntranetConnectionConfiguration"
	paramsTypeMap["intranet_configuration_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["intranet_connection_configuration"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.IntranetConnectionConfigurationBindingType)
	paramsTypeMap["intranetConfigurationId"] = vapiBindings_.NewStringType()
	pathParams["intranet_configuration_id"] = "intranetConfigurationId"
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
		"intranet_connection_configuration",
		"PUT",
		"/cloud-service/api/v1/infra/sddc-intranet-config/{intranetConfigurationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
