// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: LinkedVpcs.
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

// Possible value for ``action`` of method LinkedVpcs#create.
const LinkedVpcs_CREATE_ACTION_ENABLE_MANAGED_PREFIX_LIST_MODE = "enable_managed_prefix_list_mode"

// Possible value for ``action`` of method LinkedVpcs#create.
const LinkedVpcs_CREATE_ACTION_DISABLE_MANAGED_PREFIX_LIST_MODE = "disable_managed_prefix_list_mode"

// Possible value for ``action`` of method LinkedVpcs#create.
const LinkedVpcs_CREATE_ACTION_ENABLE_IPV6_MODE = "enable_ipv6_mode"

// Possible value for ``action`` of method LinkedVpcs#create.
const LinkedVpcs_CREATE_ACTION_DISABLE_IPV6_MODE = "disable_ipv6_mode"

func linkedVpcsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewStringType()
	fields["update_default_route_table"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["update_default_route_table"] = "UpdateDefaultRouteTable"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LinkedVpcsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func linkedVpcsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewStringType()
	fields["update_default_route_table"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["update_default_route_table"] = "UpdateDefaultRouteTable"
	paramsTypeMap["linked_vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["update_default_route_table"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["action"] = vapiBindings_.NewStringType()
	paramsTypeMap["linkedVpcId"] = vapiBindings_.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
	queryParams["update_default_route_table"] = "update_default_route_table"
	queryParams["action"] = "action"
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
		"POST",
		"/cloud-service/api/v1/infra/linked-vpcs/{linkedVpcId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func linkedVpcsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LinkedVpcsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.LinkedVpcInfoBindingType)
}

func linkedVpcsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	paramsTypeMap["linked_vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["linkedVpcId"] = vapiBindings_.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
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
		"/cloud-service/api/v1/infra/linked-vpcs/{linkedVpcId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func linkedVpcsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func LinkedVpcsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.LinkedVpcsListResultBindingType)
}

func linkedVpcsListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/linked-vpcs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
