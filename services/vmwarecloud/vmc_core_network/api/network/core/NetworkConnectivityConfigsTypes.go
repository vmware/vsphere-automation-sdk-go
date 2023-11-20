// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: NetworkConnectivityConfigs.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package core

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
	"reflect"
)

const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIGS_TRAIT_AWS_REALIZED_SDDC_CONNECTIVITY_TRAIT = "AwsRealizedSddcConnectivityTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIGS_TRAIT_AWS_DIRECT_CONNECT_GATEWAY_ASSOCIATIONS_TRAIT = "AwsDirectConnectGatewayAssociationsTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIGS_TRAIT_AWS_VPC_ATTACHMENTS_TRAIT = "AwsVpcAttachmentsTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIGS_TRAIT_AWS_NETWORK_CONNECTIVITY_TRAIT = "AwsNetworkConnectivityTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIGS_TRAIT_AWS_CUSTOMER_TRANSIT_GATEWAY_ASSOCIATIONS_TRAIT = "AwsCustomerTransitGatewayAssociationsTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIG_BY_ID_TRAIT_AWS_REALIZED_SDDC_CONNECTIVITY_TRAIT = "AwsRealizedSddcConnectivityTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIG_BY_ID_TRAIT_AWS_DIRECT_CONNECT_GATEWAY_ASSOCIATIONS_TRAIT = "AwsDirectConnectGatewayAssociationsTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIG_BY_ID_TRAIT_AWS_VPC_ATTACHMENTS_TRAIT = "AwsVpcAttachmentsTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIG_BY_ID_TRAIT_AWS_NETWORK_CONNECTIVITY_TRAIT = "AwsNetworkConnectivityTrait"
const NetworkConnectivityConfigs_GET_NETWORK_CONNECTIVITY_CONFIG_BY_ID_TRAIT_AWS_CUSTOMER_TRANSIT_GATEWAY_ASSOCIATIONS_TRAIT = "AwsCustomerTransitGatewayAssociationsTrait"

func networkConnectivityConfigsGetNetworkConnectivityConfigsInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fields["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fields["include_deleted_resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["trait"] = "Trait"
	fieldNameMap["include_deleted_resources"] = "IncludeDeletedResources"
	fieldNameMap["page"] = "Page"
	fieldNameMap["size"] = "Size"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NetworkConnectivityConfigsGetNetworkConnectivityConfigsOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.NetworkConnectivityConfigBindingType), reflect.TypeOf([]vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig{}))
}

func networkConnectivityConfigsGetNetworkConnectivityConfigsRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fields["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fields["include_deleted_resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["trait"] = "Trait"
	fieldNameMap["include_deleted_resources"] = "IncludeDeletedResources"
	fieldNameMap["page"] = "Page"
	fieldNameMap["size"] = "Size"
	paramsTypeMap["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_deleted_resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	paramsTypeMap["page"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	pathParams["org_id"] = "orgId"
	queryParams["size"] = "size"
	queryParams["group_id"] = "group_id"
	queryParams["include_deleted_resources"] = "include_deleted_resources"
	queryParams["trait"] = "trait"
	queryParams["page"] = "page"
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
		"/api/network/{orgId}/core/network-connectivity-configs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func networkConnectivityConfigsGetNetworkConnectivityConfigByIdInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["id"] = vapiBindings_.NewStringType()
	fields["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["trait"] = "Trait"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NetworkConnectivityConfigsGetNetworkConnectivityConfigByIdOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.NetworkConnectivityConfigBindingType)
}

func networkConnectivityConfigsGetNetworkConnectivityConfigByIdRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["id"] = vapiBindings_.NewStringType()
	fields["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["trait"] = "Trait"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["trait"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	paramsTypeMap["id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["id"] = vapiBindings_.NewStringType()
	pathParams["id"] = "id"
	pathParams["org_id"] = "orgId"
	queryParams["trait"] = "trait"
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
		"/api/network/{orgId}/core/network-connectivity-configs/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}
