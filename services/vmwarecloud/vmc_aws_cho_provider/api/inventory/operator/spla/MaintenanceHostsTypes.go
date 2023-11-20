// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MaintenanceHosts.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package spla

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_aws_cho_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_cho_provider/model"
	"reflect"
)

func maintenanceHostsGetSplaMaintenanceHostsInfoInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_region"] = vapiBindings_.NewStringType()
	fieldNameMap["aws_region"] = "AwsRegion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MaintenanceHostsGetSplaMaintenanceHostsInfoOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfoBindingType)
}

func maintenanceHostsGetSplaMaintenanceHostsInfoRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["aws_region"] = vapiBindings_.NewStringType()
	fieldNameMap["aws_region"] = "AwsRegion"
	paramsTypeMap["aws_region"] = vapiBindings_.NewStringType()
	queryParams["aws_region"] = "aws_region"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/api/inventory/operator/spla/maintenance-hosts",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func maintenanceHostsAddSplaMaintenanceHostsInfoInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["add_spla_maintenance_hosts_info_request_body"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.MaintenanceHostBindingType), reflect.TypeOf([]vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost{}))
	fields["aws_region"] = vapiBindings_.NewStringType()
	fieldNameMap["add_spla_maintenance_hosts_info_request_body"] = "AddSplaMaintenanceHostsInfoRequestBody"
	fieldNameMap["aws_region"] = "AwsRegion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MaintenanceHostsAddSplaMaintenanceHostsInfoOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.SplaMaintenanceHostsInfoBindingType)
}

func maintenanceHostsAddSplaMaintenanceHostsInfoRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["add_spla_maintenance_hosts_info_request_body"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.MaintenanceHostBindingType), reflect.TypeOf([]vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost{}))
	fields["aws_region"] = vapiBindings_.NewStringType()
	fieldNameMap["add_spla_maintenance_hosts_info_request_body"] = "AddSplaMaintenanceHostsInfoRequestBody"
	fieldNameMap["aws_region"] = "AwsRegion"
	paramsTypeMap["aws_region"] = vapiBindings_.NewStringType()
	paramsTypeMap["add_spla_maintenance_hosts_info_request_body"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.MaintenanceHostBindingType), reflect.TypeOf([]vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost{}))
	queryParams["aws_region"] = "aws_region"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"add_spla_maintenance_hosts_info_request_body",
		"POST",
		"/api/inventory/operator/spla/maintenance-hosts",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func maintenanceHostsRemoveSplaMaintenanceHostsInfoInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["remove_spla_maintenance_hosts_info_request_body"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.MaintenanceHostBindingType), reflect.TypeOf([]vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost{}))
	fields["aws_region"] = vapiBindings_.NewStringType()
	fieldNameMap["remove_spla_maintenance_hosts_info_request_body"] = "RemoveSplaMaintenanceHostsInfoRequestBody"
	fieldNameMap["aws_region"] = "AwsRegion"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MaintenanceHostsRemoveSplaMaintenanceHostsInfoOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func maintenanceHostsRemoveSplaMaintenanceHostsInfoRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["remove_spla_maintenance_hosts_info_request_body"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.MaintenanceHostBindingType), reflect.TypeOf([]vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost{}))
	fields["aws_region"] = vapiBindings_.NewStringType()
	fieldNameMap["remove_spla_maintenance_hosts_info_request_body"] = "RemoveSplaMaintenanceHostsInfoRequestBody"
	fieldNameMap["aws_region"] = "AwsRegion"
	paramsTypeMap["remove_spla_maintenance_hosts_info_request_body"] = vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_aws_cho_providerModel.MaintenanceHostBindingType), reflect.TypeOf([]vmwarecloudVmc_aws_cho_providerModel.MaintenanceHost{}))
	paramsTypeMap["aws_region"] = vapiBindings_.NewStringType()
	queryParams["aws_region"] = "aws_region"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"remove_spla_maintenance_hosts_info_request_body",
		"DELETE",
		"/api/inventory/operator/spla/maintenance-hosts",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}
