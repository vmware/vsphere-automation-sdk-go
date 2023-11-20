// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: CustomerNetworkPeeringscreatePeering
// Used by client-side stubs.

package deployments

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type CustomerNetworkPeeringscreatePeeringClient interface {

	// add customer network peering
	//
	// @param orgIdParam organization identifier
	// @param deploymentIdParam deployment identifier
	// @param addCustomerNetworkPeeringRequestParam add customer network peering
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	AddCustomerPeeringNetwork(orgIdParam string, deploymentIdParam string, addCustomerNetworkPeeringRequestParam vmwarecloudVmc_core_networkModel.AddCustomerNetworkPeeringRequest) (vmwarecloudVmc_core_networkModel.NetworkActivity, error)
}

type customerNetworkPeeringscreatePeeringClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCustomerNetworkPeeringscreatePeeringClient(connector vapiProtocolClient_.Connector) *customerNetworkPeeringscreatePeeringClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.customer_network_peeringscreate_peering")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"add_customer_peering_network": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "add_customer_peering_network"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := customerNetworkPeeringscreatePeeringClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *customerNetworkPeeringscreatePeeringClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *customerNetworkPeeringscreatePeeringClient) AddCustomerPeeringNetwork(orgIdParam string, deploymentIdParam string, addCustomerNetworkPeeringRequestParam vmwarecloudVmc_core_networkModel.AddCustomerNetworkPeeringRequest) (vmwarecloudVmc_core_networkModel.NetworkActivity, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := customerNetworkPeeringscreatePeeringAddCustomerPeeringNetworkRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(customerNetworkPeeringscreatePeeringAddCustomerPeeringNetworkInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("DeploymentId", deploymentIdParam)
	sv.AddStructField("AddCustomerNetworkPeeringRequest", addCustomerNetworkPeeringRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.NetworkActivity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.deployments.customer_network_peeringscreate_peering", "add_customer_peering_network", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.NetworkActivity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CustomerNetworkPeeringscreatePeeringAddCustomerPeeringNetworkOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.NetworkActivity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
