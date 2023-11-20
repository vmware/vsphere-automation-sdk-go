// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Iddelete
// Used by client-side stubs.

package vpc_peerings

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type IddeleteClient interface {

	// Delete VPC Peering
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param idParam resource identifier
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	DeleteVpcPeering(orgIdParam string, deploymentIdParam string, idParam string) (vmwarecloudVmc_core_networkModel.Activity, error)
}

type iddeleteClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewIddeleteClient(connector vapiProtocolClient_.Connector) *iddeleteClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings.iddelete")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete_vpc_peering": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete_vpc_peering"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := iddeleteClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *iddeleteClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *iddeleteClient) DeleteVpcPeering(orgIdParam string, deploymentIdParam string, idParam string) (vmwarecloudVmc_core_networkModel.Activity, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := iddeleteDeleteVpcPeeringRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(iddeleteDeleteVpcPeeringInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.Activity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.vpc_peerings.iddelete", "delete_vpc_peering", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.Activity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IddeleteDeleteVpcPeeringOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.Activity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
