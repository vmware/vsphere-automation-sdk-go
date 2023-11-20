// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Deployments
// Used by client-side stubs.

package core

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_inventoryModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DeploymentsClient interface {

	// Get a collection of deployment resources scoped by the operator
	//
	// @param orgIdParam organization identifier
	// @param summaryParam List of deployment summaries to be returned
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @param pageParam Zero-based page index (0..N)
	// @param sizeParam The size of the page to be returned
	// @param sortParam Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetCoreDeploymentsAsOperator(orgIdParam *string, summaryParam []string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.DeploymentResponse, error)

	// Get details for a specific core deployment
	//
	// @param deploymentIdParam deployment identifier
	// @param orgIdParam organization identifier
	// @param summaryParam List of deployment summaries to be returned
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	GetCoreDeploymentByIdAsOperator(deploymentIdParam string, orgIdParam *string, summaryParam []string, includeDeletedResourcesParam *bool) (vmwarecloudVmc_core_inventoryModel.Deployment, error)
}

type deploymentsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDeploymentsClient(connector vapiProtocolClient_.Connector) *deploymentsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployments")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_core_deployments_as_operator":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_deployments_as_operator"),
		"get_core_deployment_by_id_as_operator": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_deployment_by_id_as_operator"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := deploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentsClient) GetCoreDeploymentsAsOperator(orgIdParam *string, summaryParam []string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.DeploymentResponse, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsGetCoreDeploymentsAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsGetCoreDeploymentsAsOperatorInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Summary", summaryParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	sv.AddStructField("Sort", sortParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployments", "get_core_deployments_as_operator", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentsGetCoreDeploymentsAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.DeploymentResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentsClient) GetCoreDeploymentByIdAsOperator(deploymentIdParam string, orgIdParam *string, summaryParam []string, includeDeletedResourcesParam *bool) (vmwarecloudVmc_core_inventoryModel.Deployment, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsGetCoreDeploymentByIdAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsGetCoreDeploymentByIdAsOperatorInputType(), typeConverter)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Summary", summaryParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.Deployment
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployments", "get_core_deployment_by_id_as_operator", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.Deployment
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentsGetCoreDeploymentByIdAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.Deployment), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
