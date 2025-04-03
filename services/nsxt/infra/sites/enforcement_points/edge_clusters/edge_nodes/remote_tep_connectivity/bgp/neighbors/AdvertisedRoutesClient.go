// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AdvertisedRoutes
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

type AdvertisedRoutesClient interface {

    // Returns routes advertised by BGP neighbor from the given edge transport node. It always returns realtime response.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param edgeClusterIdParam (required)
    // @param policyEdgeNodeIdParam (required)
    // @param neighborIdParam (required)
    // @return com.vmware.nsx_policy.model.PolicyEdgeNodeBgpNeighborAdvertisedRoutes
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, edgeClusterIdParam string, policyEdgeNodeIdParam string, neighborIdParam string) (nsx_policyModel.PolicyEdgeNodeBgpNeighborAdvertisedRoutes, error)
}


type advertisedRoutesClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewAdvertisedRoutesClient(connector vapiProtocolClient_.Connector) *advertisedRoutesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_clusters.edge_nodes.remote_tep_connectivity.bgp.neighbors.advertised_routes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := advertisedRoutesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *advertisedRoutesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *advertisedRoutesClient) Get(siteIdParam string, enforcementpointIdParam string, edgeClusterIdParam string, policyEdgeNodeIdParam string, neighborIdParam string) (nsx_policyModel.PolicyEdgeNodeBgpNeighborAdvertisedRoutes, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := advertisedRoutesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(advertisedRoutesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("PolicyEdgeNodeId", policyEdgeNodeIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeNodeBgpNeighborAdvertisedRoutes
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_clusters.edge_nodes.remote_tep_connectivity.bgp.neighbors.advertised_routes", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeNodeBgpNeighborAdvertisedRoutes
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AdvertisedRoutesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeNodeBgpNeighborAdvertisedRoutes), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

