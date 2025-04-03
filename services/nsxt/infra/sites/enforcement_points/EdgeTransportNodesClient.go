// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EdgeTransportNodes
// Used by client-side stubs.

package enforcement_points

import (
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EdgeTransportNodesClient interface {

    // Delete Edge Transport Node.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param edgeTransportNodeIdParam (required)
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string) error

    // Read an Edge Transport Node under an Enforcement Point
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param edgeTransportNodeIdParam (required)
    // @return com.vmware.nsx_policy.model.PolicyEdgeTransportNode
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string) (nsx_policyModel.PolicyEdgeTransportNode, error)

    // List Edge Transport Nodes under an Enforcement Point
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param inMaintenanceModeParam Maintenance mode flag (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param managementIpParam Edge transport node management IP address (optional)
    // @param nodeTypeParam Supported edge transport node type. (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @param transportZonePathParam Transport zone path (optional)
    // @return com.vmware.nsx_policy.model.PolicyEdgeTransportNodeListResult
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(siteIdParam string, enforcementpointIdParam string, cursorParam *string, inMaintenanceModeParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, managementIpParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.PolicyEdgeTransportNodeListResult, error)

    // If the passed Edge Transport Node does not already exist, create a new Edge Transport Node. If it already exists, patch it. Transport nodes are hypervisor NSX Edges that will participate in an NSX-T overlay.this means that it will have Tier0/TIer1 uplinks and downlinks. This API creates/updates the edge node (router) in the transport network. Additional documentation on creating a transport node can be found in the NSX-T Installation Guide. In order for the transport node to forward packets, the switch_spec property must be specified. When creating a edge transport node, you need to specify if the edge TN switches are already manually preconfigured on the node, or if NSX should create and manage the edge TN switches. You specify this choice by the type of host switches you pass in the switch_spec property of the Edge Transport Node request payload. For a NSX edge node, NSX Manager always configures the edge TN switch. To allow NSX to manage the Edge TN switch configuration on NSX Edge nodes, pass an array of switches objects in the switch_spec property, and NSX will automatically create edge TN switches with the properties you provide. In the current NSX-T release, up to 16 host switches can be automatically managed. See the switch_spec schema definition for documentation on the properties that must be provided. If the edge node (router) is already added in system then it can be converted to transport node by providing node_id in request. If edge transport node (router) is not already present in system then new edge transport node can be created using this API.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param edgeTransportNodeIdParam (required)
    // @param policyEdgeTransportNodeParam (required)
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string, policyEdgeTransportNodeParam nsx_policyModel.PolicyEdgeTransportNode) error

    // If the passed Edge Transport Node does not already exist, create a new Edge Transport Node. If it already exists, update it.
    //
    // @param siteIdParam (required)
    // @param enforcementpointIdParam (required)
    // @param edgeTransportNodeIdParam (required)
    // @param policyEdgeTransportNodeParam (required)
    // @return com.vmware.nsx_policy.model.PolicyEdgeTransportNode
    //
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string, policyEdgeTransportNodeParam nsx_policyModel.PolicyEdgeTransportNode) (nsx_policyModel.PolicyEdgeTransportNode, error)
}


type edgeTransportNodesClient struct {
	connector           	   vapiProtocolClient_.Connector
	interfaceDefinition 	   vapiCore_.InterfaceDefinition
	errorsBindingMap           map[string]vapiBindings_.BindingType
}

func NewEdgeTransportNodesClient(connector vapiProtocolClient_.Connector) *edgeTransportNodesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_transport_nodes")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := edgeTransportNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *edgeTransportNodesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *edgeTransportNodesClient) Delete(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeTransportNodesDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeTransportNodesDeleteInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeTransportNodeId", edgeTransportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_transport_nodes", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *edgeTransportNodesClient) Get(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string) (nsx_policyModel.PolicyEdgeTransportNode, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeTransportNodesGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeTransportNodesGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeTransportNodeId", edgeTransportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeTransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_transport_nodes", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EdgeTransportNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeTransportNodesClient) List(siteIdParam string, enforcementpointIdParam string, cursorParam *string, inMaintenanceModeParam *bool, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, managementIpParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportZonePathParam *string) (nsx_policyModel.PolicyEdgeTransportNodeListResult, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeTransportNodesListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeTransportNodesListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("InMaintenanceMode", inMaintenanceModeParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("ManagementIp", managementIpParam)
	sv.AddStructField("NodeType", nodeTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("TransportZonePath", transportZonePathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeTransportNodeListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_transport_nodes", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeTransportNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EdgeTransportNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeTransportNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeTransportNodesClient) Patch(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string, policyEdgeTransportNodeParam nsx_policyModel.PolicyEdgeTransportNode) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeTransportNodesPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeTransportNodesPatchInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeTransportNodeId", edgeTransportNodeIdParam)
	sv.AddStructField("PolicyEdgeTransportNode", policyEdgeTransportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_transport_nodes", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *edgeTransportNodesClient) Update(siteIdParam string, enforcementpointIdParam string, edgeTransportNodeIdParam string, policyEdgeTransportNodeParam nsx_policyModel.PolicyEdgeTransportNode) (nsx_policyModel.PolicyEdgeTransportNode, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := edgeTransportNodesUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(edgeTransportNodesUpdateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("EdgeTransportNodeId", edgeTransportNodeIdParam)
	sv.AddStructField("PolicyEdgeTransportNode", policyEdgeTransportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.PolicyEdgeTransportNode
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.edge_transport_nodes", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.PolicyEdgeTransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EdgeTransportNodesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.PolicyEdgeTransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

