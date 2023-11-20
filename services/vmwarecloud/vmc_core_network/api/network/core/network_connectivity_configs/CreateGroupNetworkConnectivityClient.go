// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CreateGroupNetworkConnectivity
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

type CreateGroupNetworkConnectivityClient interface {

	// An aggregate api to create Deployment Group, NetworkConnectivityConfig and apply the NetworkConnectivityConfig
	//
	// @param orgIdParam organization identifier
	// @param createGroupNetworkConnectivityRequestParam Deployment Ids for the members to be associated with a NetworkConnectivityConfig
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws ConcurrentChange Conflict
	CreateGroupNetworkConnectivity(orgIdParam string, createGroupNetworkConnectivityRequestParam vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityRequest) (vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityResponse, error)
}

type createGroupNetworkConnectivityClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCreateGroupNetworkConnectivityClient(connector vapiProtocolClient_.Connector) *createGroupNetworkConnectivityClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.create_group_network_connectivity")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create_group_network_connectivity": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create_group_network_connectivity"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := createGroupNetworkConnectivityClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *createGroupNetworkConnectivityClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *createGroupNetworkConnectivityClient) CreateGroupNetworkConnectivity(orgIdParam string, createGroupNetworkConnectivityRequestParam vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityRequest) (vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityResponse, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := createGroupNetworkConnectivityCreateGroupNetworkConnectivityRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(createGroupNetworkConnectivityCreateGroupNetworkConnectivityInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("CreateGroupNetworkConnectivityRequest", createGroupNetworkConnectivityRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.create_group_network_connectivity", "create_group_network_connectivity", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CreateGroupNetworkConnectivityCreateGroupNetworkConnectivityOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.CreateGroupNetworkConnectivityResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
