// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Redeploy
// Used by client-side stubs.

package action

import (
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RedeployClient interface {

    // Redeploys a service node at NSX Manager that replaces the service node with identifier <node-id>. If NSX Manager can access the specified service node, then the node is put into maintenance mode and then the associated VM is deleted. This is a means to reset all configuration on the service node. The communication channel between NSX Manager and service is established after this operation.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param serviceClusterIdParam (required)
    // @param serviceNodeIdParam (required)
    // @param policyServiceNodeParam (required)
    // @return com.vmware.nsx_policy.model.PolicyServiceNode
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Create(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, policyServiceNodeParam nsx_policyModel.PolicyServiceNode) (nsx_policyModel.PolicyServiceNode, error)
}


type redeployClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewRedeployClient(connector vapiProtocolClient_.Connector) *redeployClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes.action.redeploy")
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

func (rIface *redeployClient) Create(siteIdParam string, enforcementpointIdParam string, serviceClusterIdParam string, serviceNodeIdParam string, policyServiceNodeParam nsx_policyModel.PolicyServiceNode) (nsx_policyModel.PolicyServiceNode, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := redeployCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(redeployCreateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("ServiceClusterId", serviceClusterIdParam)
	sv.AddStructField("ServiceNodeId", serviceNodeIdParam)
	sv.AddStructField("PolicyServiceNode", policyServiceNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyServiceNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.service_clusters.service_nodes.action.redeploy", "create", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyServiceNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RedeployCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyServiceNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

