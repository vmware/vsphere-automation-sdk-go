// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DeploymentIdmsftLicensingPublish
// Used by client-side stubs.

package deployments

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_aws_cho_providerModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_aws_cho_provider/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DeploymentIdmsftLicensingPublishClient interface {

	// Enables Microsoft Server Licenses provided by VMware
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param publishMsftLicenseIntentParam Payload for Publish Microsoft licenses config on cluster
	// @return Activity accepted for processing
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Not Found
	PublishMicrosoftServerLicenses(orgIdParam string, deploymentIdParam string, publishMsftLicenseIntentParam vmwarecloudVmc_aws_cho_providerModel.PublishMsftLicenseIntent) (vmwarecloudVmc_aws_cho_providerModel.ActivityResponse, error)
}

type deploymentIdmsftLicensingPublishClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDeploymentIdmsftLicensingPublishClient(connector vapiProtocolClient_.Connector) *deploymentIdmsftLicensingPublishClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.vmc_aws.deployments.deployment_idmsft_licensing_publish")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"publish_microsoft_server_licenses": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "publish_microsoft_server_licenses"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := deploymentIdmsftLicensingPublishClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentIdmsftLicensingPublishClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentIdmsftLicensingPublishClient) PublishMicrosoftServerLicenses(orgIdParam string, deploymentIdParam string, publishMsftLicenseIntentParam vmwarecloudVmc_aws_cho_providerModel.PublishMsftLicenseIntent) (vmwarecloudVmc_aws_cho_providerModel.ActivityResponse, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentIdmsftLicensingPublishPublishMicrosoftServerLicensesRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentIdmsftLicensingPublishPublishMicrosoftServerLicensesInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("PublishMsftLicenseIntent", publishMsftLicenseIntentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_aws_cho_providerModel.ActivityResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_aws_cho_provider.api.inventory.vmc_aws.deployments.deployment_idmsft_licensing_publish", "publish_microsoft_server_licenses", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_aws_cho_providerModel.ActivityResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentIdmsftLicensingPublishPublishMicrosoftServerLicensesOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_aws_cho_providerModel.ActivityResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
