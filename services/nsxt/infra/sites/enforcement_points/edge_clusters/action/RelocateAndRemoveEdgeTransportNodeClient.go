// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RelocateAndRemoveEdgeTransportNode
// Used by client-side stubs.

package action

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RelocateAndRemoveEdgeTransportNodeClient interface {

	// Relocate auto allocated service contexts from policy edge node at given id. For API to perform relocate and remove action the edge node at given id must only have auto allocated service contexts. If any manually allocated service context is present on the edge cluster member, then the task will not be performed. Also, it is recommended to move edge node for which relocate and remove action is being performed into maintenance mode, before executing the API. If edge is not not moved into maintenance mode, then API will move edge node into maintenance mode before performing the actual relocate and remove task.To maintain high availability, Edge cluster should have at least two healthy edge nodes for relocate and removal. Once relocate action is performed successfully, the policy edge node will be removed from the edge cluster.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param edgeClusterIdParam (required)
	// @param policyEdgeNodeRelocateAndRemoveMemberParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(siteIdParam string, enforcementpointIdParam string, edgeClusterIdParam string, policyEdgeNodeRelocateAndRemoveMemberParam nsx_policyModel.PolicyEdgeNodeRelocateAndRemoveMember) error
}

type relocateAndRemoveEdgeTransportNodeClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRelocateAndRemoveEdgeTransportNodeClient(connector vapiProtocolClient_.Connector) *relocateAndRemoveEdgeTransportNodeClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_clusters.action.relocate_and_remove_edge_transport_node")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := relocateAndRemoveEdgeTransportNodeClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *relocateAndRemoveEdgeTransportNodeClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *relocateAndRemoveEdgeTransportNodeClient) Create(siteIdParam string, enforcementpointIdParam string, edgeClusterIdParam string, policyEdgeNodeRelocateAndRemoveMemberParam nsx_policyModel.PolicyEdgeNodeRelocateAndRemoveMember) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := relocateAndRemoveEdgeTransportNodeCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(relocateAndRemoveEdgeTransportNodeCreateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("PolicyEdgeNodeRelocateAndRemoveMember", policyEdgeNodeRelocateAndRemoveMemberParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_clusters.action.relocate_and_remove_edge_transport_node", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
