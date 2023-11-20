// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: UpdateNetworkingState
// Used by client-side stubs.

package nsx

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type UpdateNetworkingStateClient interface {

	// API to update SDDC networking-state in core-network for deployment Id
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param deploymentNetworkingStateInfoParam networking state info request payload
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the deployment with given identifier
	UpdateSddcNetworkingState(orgIdParam string, deploymentIdParam string, deploymentNetworkingStateInfoParam vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo) (vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo, error)
}

type updateNetworkingStateClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewUpdateNetworkingStateClient(connector vapiProtocolClient_.Connector) *updateNetworkingStateClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.nsx.update_networking_state")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"update_sddc_networking_state": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update_sddc_networking_state"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	uIface := updateNetworkingStateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *updateNetworkingStateClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *updateNetworkingStateClient) UpdateSddcNetworkingState(orgIdParam string, deploymentIdParam string, deploymentNetworkingStateInfoParam vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo) (vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	operationRestMetaData := updateNetworkingStateUpdateSddcNetworkingStateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(updateNetworkingStateUpdateSddcNetworkingStateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("DeploymentNetworkingStateInfo", deploymentNetworkingStateInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.nsx.update_networking_state", "update_sddc_networking_state", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), UpdateNetworkingStateUpdateSddcNetworkingStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.DeploymentNetworkingStateInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
