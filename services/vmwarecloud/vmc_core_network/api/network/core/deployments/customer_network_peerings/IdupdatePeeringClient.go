// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: IdupdatePeering
// Used by client-side stubs.

package customer_network_peerings

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type IdupdatePeeringClient interface {

	// update customer network peering
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param idParam resource identifier
	// @param updateCustomerNetworkPeeringRequestParam update customer network peering
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	UpdateCustomerNetworkPeering(orgIdParam string, deploymentIdParam string, idParam string, updateCustomerNetworkPeeringRequestParam vmwarecloudVmc_core_networkModel.UpdateCustomerNetworkPeeringRequest) (vmwarecloudVmc_core_networkModel.NetworkActivity, error)
}

type idupdatePeeringClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewIdupdatePeeringClient(connector vapiProtocolClient_.Connector) *idupdatePeeringClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.customer_network_peerings.idupdate_peering")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"update_customer_network_peering": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update_customer_network_peering"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	iIface := idupdatePeeringClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *idupdatePeeringClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *idupdatePeeringClient) UpdateCustomerNetworkPeering(orgIdParam string, deploymentIdParam string, idParam string, updateCustomerNetworkPeeringRequestParam vmwarecloudVmc_core_networkModel.UpdateCustomerNetworkPeeringRequest) (vmwarecloudVmc_core_networkModel.NetworkActivity, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	operationRestMetaData := idupdatePeeringUpdateCustomerNetworkPeeringRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(idupdatePeeringUpdateCustomerNetworkPeeringInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("UpdateCustomerNetworkPeeringRequest", updateCustomerNetworkPeeringRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.NetworkActivity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.customer_network_peerings.idupdate_peering", "update_customer_network_peering", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.NetworkActivity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), IdupdatePeeringUpdateCustomerNetworkPeeringOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.NetworkActivity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
