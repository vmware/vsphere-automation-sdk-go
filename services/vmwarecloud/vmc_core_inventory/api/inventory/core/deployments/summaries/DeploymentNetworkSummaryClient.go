// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DeploymentNetworkSummary
// Used by client-side stubs.

package summaries

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_inventoryModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DeploymentNetworkSummaryClient interface {

	// Get the DeploymentNetworkSummary specific core deployment summary
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	GetCoreDeploymentNetworkSummary(orgIdParam string, deploymentIdParam string) (vmwarecloudVmc_core_inventoryModel.DeploymentNetworkSummary, error)
}

type deploymentNetworkSummaryClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDeploymentNetworkSummaryClient(connector vapiProtocolClient_.Connector) *deploymentNetworkSummaryClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.core.deployments.summaries.deployment_network_summary")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_core_deployment_network_summary": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_deployment_network_summary"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := deploymentNetworkSummaryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentNetworkSummaryClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentNetworkSummaryClient) GetCoreDeploymentNetworkSummary(orgIdParam string, deploymentIdParam string) (vmwarecloudVmc_core_inventoryModel.DeploymentNetworkSummary, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentNetworkSummaryGetCoreDeploymentNetworkSummaryRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentNetworkSummaryGetCoreDeploymentNetworkSummaryInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentNetworkSummary
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.core.deployments.summaries.deployment_network_summary", "get_core_deployment_network_summary", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentNetworkSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentNetworkSummaryGetCoreDeploymentNetworkSummaryOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.DeploymentNetworkSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
