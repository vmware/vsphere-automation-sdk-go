// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SharedPrefixList
// Used by client-side stubs.

package network_connectivity_configs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SharedPrefixListClient interface {

	// gets all shared prefix lists of given network connectivity config
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetAllSharedPrefixListsForNetworkConnectivityConfig(orgIdParam string, idParam string) ([]vmwarecloudVmc_core_networkModel.AwsSharedPrefixList, error)

	// gets shared prefix list of given id and network connectivity config
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @param sharedPrefixListIdParam Shared Prefix List Id
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the shared prefix list with given identifier
	GetSharedPrefixList(orgIdParam string, idParam string, sharedPrefixListIdParam string) (vmwarecloudVmc_core_networkModel.AwsSharedPrefixList, error)
}

type sharedPrefixListClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSharedPrefixListClient(connector vapiProtocolClient_.Connector) *sharedPrefixListClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.shared_prefix_list")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_all_shared_prefix_lists_for_network_connectivity_config": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_all_shared_prefix_lists_for_network_connectivity_config"),
		"get_shared_prefix_list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_shared_prefix_list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := sharedPrefixListClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sharedPrefixListClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sharedPrefixListClient) GetAllSharedPrefixListsForNetworkConnectivityConfig(orgIdParam string, idParam string) ([]vmwarecloudVmc_core_networkModel.AwsSharedPrefixList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sharedPrefixListGetAllSharedPrefixListsForNetworkConnectivityConfigRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sharedPrefixListGetAllSharedPrefixListsForNetworkConnectivityConfigInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmwarecloudVmc_core_networkModel.AwsSharedPrefixList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.shared_prefix_list", "get_all_shared_prefix_lists_for_network_connectivity_config", inputDataValue, executionContext)
	var emptyOutput []vmwarecloudVmc_core_networkModel.AwsSharedPrefixList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SharedPrefixListGetAllSharedPrefixListsForNetworkConnectivityConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmwarecloudVmc_core_networkModel.AwsSharedPrefixList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sharedPrefixListClient) GetSharedPrefixList(orgIdParam string, idParam string, sharedPrefixListIdParam string) (vmwarecloudVmc_core_networkModel.AwsSharedPrefixList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sharedPrefixListGetSharedPrefixListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sharedPrefixListGetSharedPrefixListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("SharedPrefixListId", sharedPrefixListIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.AwsSharedPrefixList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.shared_prefix_list", "get_shared_prefix_list", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.AwsSharedPrefixList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SharedPrefixListGetSharedPrefixListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.AwsSharedPrefixList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
