// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Status
// Used by client-side stubs.

package nsx_advanced_security_addon

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StatusClient interface {

	// Respond with the current status of the NSX addon for given resource
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @return OK
	//
	// @throws InvalidRequest Bad Request
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the deployment with given identifier
	GetNsxAdvancedSecurityAddonStatus(orgIdParam string, deploymentIdParam string) (vmwarecloudVmc_core_networkModel.NsxAddonResponse, error)
}

type statusClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStatusClient(connector vapiProtocolClient_.Connector) *statusClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.nsx_advanced_security_addon.status")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_nsx_advanced_security_addon_status": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_nsx_advanced_security_addon_status"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := statusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statusClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statusClient) GetNsxAdvancedSecurityAddonStatus(orgIdParam string, deploymentIdParam string) (vmwarecloudVmc_core_networkModel.NsxAddonResponse, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statusGetNsxAdvancedSecurityAddonStatusRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statusGetNsxAdvancedSecurityAddonStatusInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.NsxAddonResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.nsx_advanced_security_addon.status", "get_nsx_advanced_security_addon_status", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.NsxAddonResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatusGetNsxAdvancedSecurityAddonStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.NsxAddonResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
