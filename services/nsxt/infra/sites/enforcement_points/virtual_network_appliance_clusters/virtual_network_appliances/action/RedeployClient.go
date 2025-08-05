// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Redeploy
// Used by client-side stubs.

package action

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RedeployClient interface {

	// Redeploys a virtual network appliance(VNA) at NSX Manager that replaces the virtual network appliance(VNA) with identifier <id>. If NSX Manager can access the specified virtual network appliance(VNA), then the appliance is put into maintenance mode and then the associated VM is deleted. This is a means to reset all configuration on the virtual network appliance(VNA). The communication channel between NSX Manager and service is established after this operation.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param virtualNetworkApplianceClusterIdParam (required)
	// @param virtualNetworkApplianceIdParam (required)
	// @param baseNetworkApplianceParam (required)
	// The parameter must contain all the properties defined in nsx_policyModel.BaseNetworkAppliance.
	// @return com.vmware.nsx_policy.model.BaseNetworkAppliance
	// The return value will contain all the properties defined in nsx_policyModel.BaseNetworkAppliance.
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, baseNetworkApplianceParam *vapiData_.StructValue) (*vapiData_.StructValue, error)
}

type redeployClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRedeployClient(connector vapiProtocolClient_.Connector) *redeployClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances.action.redeploy")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := redeployClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *redeployClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *redeployClient) Create(siteIdParam string, enforcementpointIdParam string, virtualNetworkApplianceClusterIdParam string, virtualNetworkApplianceIdParam string, baseNetworkApplianceParam *vapiData_.StructValue) (*vapiData_.StructValue, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := redeployCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(redeployCreateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("VirtualNetworkApplianceClusterId", virtualNetworkApplianceClusterIdParam)
	sv.AddStructField("VirtualNetworkApplianceId", virtualNetworkApplianceIdParam)
	sv.AddStructField("BaseNetworkAppliance", baseNetworkApplianceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *vapiData_.StructValue
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.virtual_network_appliance_clusters.virtual_network_appliances.action.redeploy", "create", inputDataValue, executionContext)
	var emptyOutput *vapiData_.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RedeployCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(*vapiData_.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
