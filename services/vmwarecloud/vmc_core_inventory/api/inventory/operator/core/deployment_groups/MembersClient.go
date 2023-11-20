// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Members
// Used by client-side stubs.

package deployment_groups

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_inventoryModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type MembersClient interface {

	// Get the effective members of the deployment group
	//
	// @param deploymentGroupIdParam deployment group identifier
	// @param orgIdParam organization identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	GetCoreDeploymentGroupMembersAsOperator(deploymentGroupIdParam string, orgIdParam *string) ([]vmwarecloudVmc_core_inventoryModel.GroupMember, error)
}

type membersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewMembersClient(connector vapiProtocolClient_.Connector) *membersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployment_groups.members")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_core_deployment_group_members_as_operator": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_core_deployment_group_members_as_operator"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	mIface := membersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *membersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *membersClient) GetCoreDeploymentGroupMembersAsOperator(deploymentGroupIdParam string, orgIdParam *string) ([]vmwarecloudVmc_core_inventoryModel.GroupMember, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := membersGetCoreDeploymentGroupMembersAsOperatorRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(membersGetCoreDeploymentGroupMembersAsOperatorInputType(), typeConverter)
	sv.AddStructField("DeploymentGroupId", deploymentGroupIdParam)
	sv.AddStructField("OrgId", orgIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmwarecloudVmc_core_inventoryModel.GroupMember
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_inventory.api.inventory.operator.core.deployment_groups.members", "get_core_deployment_group_members_as_operator", inputDataValue, executionContext)
	var emptyOutput []vmwarecloudVmc_core_inventoryModel.GroupMember
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), MembersGetCoreDeploymentGroupMembersAsOperatorOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmwarecloudVmc_core_inventoryModel.GroupMember), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
