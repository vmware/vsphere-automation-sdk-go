// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DeploymentIdaddHost
// Used by client-side stubs.

package deployments

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_aws_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_provider/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DeploymentIdaddHostClient interface {

	// Adds a host to the cluster associated with a deployment
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param addHostIntentParam Payload for the Add Host request
	// @return Activity accepted for processing
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	PostDeploymentAddHost(orgIdParam string, deploymentIdParam string, addHostIntentParam vmwarecloudVmc_aws_providerModel.AddHostIntent) (vmwarecloudVmc_aws_providerModel.ActivityResponse, error)
}

type deploymentIdaddHostClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDeploymentIdaddHostClient(connector vapiProtocolClient_.Connector) *deploymentIdaddHostClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_aws_provider.api.inventory.vmc_aws.deployments.deployment_idadd_host")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"post_deployment_add_host": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "post_deployment_add_host"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := deploymentIdaddHostClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentIdaddHostClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentIdaddHostClient) PostDeploymentAddHost(orgIdParam string, deploymentIdParam string, addHostIntentParam vmwarecloudVmc_aws_providerModel.AddHostIntent) (vmwarecloudVmc_aws_providerModel.ActivityResponse, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentIdaddHostPostDeploymentAddHostRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentIdaddHostPostDeploymentAddHostInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("AddHostIntent", addHostIntentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_providerModel.ActivityResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_provider.api.inventory.vmc_aws.deployments.deployment_idadd_host", "post_deployment_add_host", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_providerModel.ActivityResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentIdaddHostPostDeploymentAddHostOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_aws_providerModel.ActivityResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
