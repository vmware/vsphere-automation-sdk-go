// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: NetworkConnectivityConfigs
// Used by client-side stubs.

package core

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type NetworkConnectivityConfigsClient interface {

	// Get a collection of NetworkConnectivityConfig resources scoped by organization
	//
	// @param orgIdParam organization identifier
	// @param groupIdParam List of deployment group identifier
	// @param traitParam List of resource traits to be returned
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @param pageParam Number of requested page (0-based)
	// @param sizeParam Size of each page
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetNetworkConnectivityConfigs(orgIdParam string, groupIdParam []string, traitParam []string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64) ([]vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig, error)

	// Get details for a specific core NetworkConnectivityConfig
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @param traitParam List of resource traits to be returned
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the deployment with given identifier
	GetNetworkConnectivityConfigById(orgIdParam string, idParam string, traitParam []string) (vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig, error)
}

type networkConnectivityConfigsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewNetworkConnectivityConfigsClient(connector vapiProtocolClient_.Connector) *networkConnectivityConfigsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_network_connectivity_configs":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_network_connectivity_configs"),
		"get_network_connectivity_config_by_id": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_network_connectivity_config_by_id"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	nIface := networkConnectivityConfigsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *networkConnectivityConfigsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *networkConnectivityConfigsClient) GetNetworkConnectivityConfigs(orgIdParam string, groupIdParam []string, traitParam []string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64) ([]vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := networkConnectivityConfigsGetNetworkConnectivityConfigsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(networkConnectivityConfigsGetNetworkConnectivityConfigsInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("Trait", traitParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs", "get_network_connectivity_configs", inputDataValue, executionContext)
	var emptyOutput []vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), NetworkConnectivityConfigsGetNetworkConnectivityConfigsOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *networkConnectivityConfigsClient) GetNetworkConnectivityConfigById(orgIdParam string, idParam string, traitParam []string) (vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := networkConnectivityConfigsGetNetworkConnectivityConfigByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(networkConnectivityConfigsGetNetworkConnectivityConfigByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Trait", traitParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs", "get_network_connectivity_config_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), NetworkConnectivityConfigsGetNetworkConnectivityConfigByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.NetworkConnectivityConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
