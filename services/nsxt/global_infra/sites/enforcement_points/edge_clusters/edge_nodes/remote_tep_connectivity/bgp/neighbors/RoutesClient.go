// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Routes
// Used by client-side stubs.

package neighbors

import (
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RoutesClient interface {

    // Returns routes learned by BGP neighbor from the given edge transport node. It always returns realtime response.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param edgeClusterIdParam (required)
    // @param policyEdgeNodeIdParam (required)
    // @param neighborIdParam (required)
    // @return com.vmware.nsx_policy.model.PolicyEdgeNodeBgpNeighborRouteDetailList
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, edgeClusterIdParam string, policyEdgeNodeIdParam string, neighborIdParam string) (nsx_policyModel.PolicyEdgeNodeBgpNeighborRouteDetailList, error)
}


type routesClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewRoutesClient(connector vapiProtocolClient_.Connector) *routesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.global_infra.sites.enforcement_points.edge_clusters.edge_nodes.remote_tep_connectivity.bgp.neighbors.routes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := routesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *routesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *routesClient) Get(siteIdParam string, enforcementpointIdParam string, edgeClusterIdParam string, policyEdgeNodeIdParam string, neighborIdParam string) (nsx_policyModel.PolicyEdgeNodeBgpNeighborRouteDetailList, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("PolicyEdgeNodeId", policyEdgeNodeIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeNodeBgpNeighborRouteDetailList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.global_infra.sites.enforcement_points.edge_clusters.edge_nodes.remote_tep_connectivity.bgp.neighbors.routes", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeNodeBgpNeighborRouteDetailList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RoutesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeNodeBgpNeighborRouteDetailList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

