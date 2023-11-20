// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Traits.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package network_connectivity_configs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
	"reflect"
)

const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_HOST_VCENTER_TRAIT = "HostVcenterTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_HOST_VCENTER_EXTENDED_TRAIT = "HostVcenterExtendedTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_HOST_VMC_AWS_TRAIT = "HostVmcAwsTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_CLUSTER_VCENTER_TRAIT = "ClusterVcenterTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_CLUSTER_VCENTER_EXTENDED_TRAIT = "ClusterVcenterExtendedTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_CLUSTER_VMC_AWS_TRAIT = "ClusterVmcAwsTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_DEPLOYMENT_VCENTER_TRAIT = "DeploymentVcenterTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_DEPLOYMENT_VCENTER_EXTENDED_TRAIT = "DeploymentVcenterExtendedTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_DEPLOYMENT_VMC_AWS_TRAIT = "DeploymentVmcAwsTrait"
const Traits_GET_NETWORK_CONNECTIVITY_CONFIG_TRAIT_TRAIT_NAME_DEPLOYMENT_VMC_AWS_CREDENTIALS_TRAIT = "DeploymentVmcAwsCredentialsTrait"

func traitsGetNetworkConnectivityConfigTraitInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["id"] = vapiBindings_.NewStringType()
	fields["trait_name"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["trait_name"] = "TraitName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TraitsGetNetworkConnectivityConfigTraitOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_networkModel.TraitBindingType)
}

func traitsGetNetworkConnectivityConfigTraitRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["trait_name"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["trait_name"] = "TraitName"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["id"] = vapiBindings_.NewStringType()
	paramsTypeMap["trait_name"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["id"] = vapiBindings_.NewStringType()
	paramsTypeMap["traitName"] = vapiBindings_.NewStringType()
	pathParams["trait_name"] = "traitName"
	pathParams["id"] = "id"
	pathParams["org_id"] = "orgId"
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
		"/api/network/{orgId}/core/network-connectivity-configs/{id}/traits/{traitName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}
