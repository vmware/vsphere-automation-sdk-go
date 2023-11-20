// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ReservedNetworks.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package sddc_network_config

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
	"reflect"
)

func reservedNetworksDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reserved_network_id"] = vapiBindings_.NewStringType()
	fieldNameMap["reserved_network_id"] = "ReservedNetworkId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReservedNetworksDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func reservedNetworksDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["reserved_network_id"] = vapiBindings_.NewStringType()
	fieldNameMap["reserved_network_id"] = "ReservedNetworkId"
	paramsTypeMap["reserved_network_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservedNetworkId"] = vapiBindings_.NewStringType()
	pathParams["reserved_network_id"] = "reservedNetworkId"
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
		"/cloud-service/api/v1/infra/sddc-network-config/reserved-networks/{reservedNetworkId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func reservedNetworksGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reserved_network_id"] = vapiBindings_.NewStringType()
	fieldNameMap["reserved_network_id"] = "ReservedNetworkId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReservedNetworksGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
}

func reservedNetworksGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["reserved_network_id"] = vapiBindings_.NewStringType()
	fieldNameMap["reserved_network_id"] = "ReservedNetworkId"
	paramsTypeMap["reserved_network_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservedNetworkId"] = vapiBindings_.NewStringType()
	pathParams["reserved_network_id"] = "reservedNetworkId"
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
		"/cloud-service/api/v1/infra/sddc-network-config/reserved-networks/{reservedNetworkId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func reservedNetworksListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReservedNetworksListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlocksListResultBindingType)
}

func reservedNetworksListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/sddc-network-config/reserved-networks",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func reservedNetworksUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reserved_network_id"] = vapiBindings_.NewStringType()
	fields["reserved_CIDR_block"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
	fieldNameMap["reserved_network_id"] = "ReservedNetworkId"
	fieldNameMap["reserved_CIDR_block"] = "ReservedCIDRBlock"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReservedNetworksUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
}

func reservedNetworksUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["reserved_network_id"] = vapiBindings_.NewStringType()
	fields["reserved_CIDR_block"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
	fieldNameMap["reserved_network_id"] = "ReservedNetworkId"
	fieldNameMap["reserved_CIDR_block"] = "ReservedCIDRBlock"
	paramsTypeMap["reserved_CIDR_block"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
	paramsTypeMap["reserved_network_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservedNetworkId"] = vapiBindings_.NewStringType()
	pathParams["reserved_network_id"] = "reservedNetworkId"
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
		"reserved_CIDR_block",
		"PUT",
		"/cloud-service/api/v1/infra/sddc-network-config/reserved-networks/{reservedNetworkId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func reservedNetworksValidateonlyInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["reserved_CIDR_block"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
	fieldNameMap["reserved_CIDR_block"] = "ReservedCIDRBlock"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReservedNetworksValidateonlyOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func reservedNetworksValidateonlyRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["reserved_CIDR_block"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
	fieldNameMap["reserved_CIDR_block"] = "ReservedCIDRBlock"
	paramsTypeMap["reserved_CIDR_block"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ReservedCIDRBlockBindingType)
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
		"action=VALIDATE_ONLY",
		"reserved_CIDR_block",
		"POST",
		"/cloud-service/api/v1/infra/sddc-network-config/reserved-networks",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
