// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DeploymentGroups
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

type DeploymentGroupsClient interface {

	// Get a collection of deployment groups
	//
	// @param orgIdParam organization identifier
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @param pageParam Zero-based page index (0..N)
	// @param sizeParam The size of the page to be returned
	// @param sortParam Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	GetCoreDeploymentGroupsAsOperator(orgIdParam *string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.DeploymentGroupResponse, error)

	// Get details for a specific core deployment group
	//
	// @param deploymentGroupIdParam deployment group identifier
	// @param orgIdParam organization identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	GetCoreDeploymentGroupByIdAsOperator(deploymentGroupIdParam string, orgIdParam *string) (vmwarecloudVmc_core_inventoryModel.DeploymentGroup, error)
}

type deploymentGroupsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDeploymentGroupsClient(connector vapiProtocolClient_.Connector) *deploymentGroupsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployment_groups")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_core_deployment_groups_as_operator":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_deployment_groups_as_operator"),
		"get_core_deployment_group_by_id_as_operator": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_deployment_group_by_id_as_operator"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := deploymentGroupsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentGroupsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentGroupsClient) GetCoreDeploymentGroupsAsOperator(orgIdParam *string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string) (vmwarecloudVmc_core_inventoryModel.DeploymentGroupResponse, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentGroupsGetCoreDeploymentGroupsAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentGroupsGetCoreDeploymentGroupsAsOperatorInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	sv.AddStructField("Sort", sortParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentGroupResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployment_groups", "get_core_deployment_groups_as_operator", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentGroupResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentGroupsGetCoreDeploymentGroupsAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.DeploymentGroupResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentGroupsClient) GetCoreDeploymentGroupByIdAsOperator(deploymentGroupIdParam string, orgIdParam *string) (vmwarecloudVmc_core_inventoryModel.DeploymentGroup, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentGroupsGetCoreDeploymentGroupByIdAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentGroupsGetCoreDeploymentGroupByIdAsOperatorInputType(), typeConverter)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("OrgId", orgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentGroup
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployment_groups", "get_core_deployment_group_by_id_as_operator", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_inventoryModel.DeploymentGroup
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentGroupsGetCoreDeploymentGroupByIdAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_inventoryModel.DeploymentGroup), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
