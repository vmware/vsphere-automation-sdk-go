// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricNodesDiscoveredNodes
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationFabricNodesDiscoveredNodesClient interface {

	// NSX components are installaed on host and transport node is created with given configurations.
	//
	// @param nodeExtIdParam (required)
	// @param transportNodeParam (required)
	// @return com.vmware.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateTransportNodeForDiscoveredNodeCreateTransportNode(nodeExtIdParam string, transportNodeParam model.TransportNode) (model.TransportNode, error)

	//
	//
	// @param nodeExtIdParam (required)
	// @return com.vmware.model.Node
	// The return value will contain all the properties defined in model.Node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	HostPrepDiscoveredNodeHostprep(nodeExtIdParam string) (*data.StructValue, error)

	// Returns information about all discovered nodes.
	//
	// @param cmLocalIdParam Local Id of the discovered node in the Compute Manager (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param displayNameParam Display name of discovered node (optional)
	// @param externalIdParam External id of the discovered node, ex. a mo-ref from VC (optional)
	// @param hasParentParam Discovered node has a parent compute collection or is a standalone host (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param ipAddressParam IP address of the discovered node (optional)
	// @param nodeIdParam Id of the fabric node created from the discovered node (optional)
	// @param nodeTypeParam Discovered Node type like HostNode (optional)
	// @param originIdParam Id of the compute manager from where this node was discovered (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param parentComputeCollectionParam External id of the compute collection to which this node belongs (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.DiscoveredNodeListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListDiscoveredNodes(cmLocalIdParam *string, cursorParam *string, displayNameParam *string, externalIdParam *string, hasParentParam *string, includedFieldsParam *string, ipAddressParam *string, nodeIdParam *string, nodeTypeParam *string, originIdParam *string, pageSizeParam *int64, parentComputeCollectionParam *string, sortAscendingParam *bool, sortByParam *string) (model.DiscoveredNodeListResult, error)

	// Returns information about a specific discovered node.
	//
	// @param nodeExtIdParam (required)
	// @return com.vmware.model.DiscoveredNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadDiscoveredNode(nodeExtIdParam string) (model.DiscoveredNode, error)

	//
	//
	// @param nodeExtIdParam (required)
	// @return com.vmware.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReapplyTNProfileOnDiscoveredNodeReapplyClusterConfig(nodeExtIdParam string) (model.TransportNode, error)
}

type systemAdministrationConfigurationFabricNodesDiscoveredNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricNodesDiscoveredNodesClient(connector client.Connector) *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_nodes_discovered_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_transport_node_for_discovered_node_create_transport_node": core.NewMethodIdentifier(interfaceIdentifier, "create_transport_node_for_discovered_node_create_transport_node"),
		"host_prep_discovered_node_hostprep":                              core.NewMethodIdentifier(interfaceIdentifier, "host_prep_discovered_node_hostprep"),
		"list_discovered_nodes":                                           core.NewMethodIdentifier(interfaceIdentifier, "list_discovered_nodes"),
		"read_discovered_node":                                            core.NewMethodIdentifier(interfaceIdentifier, "read_discovered_node"),
		"reapply_TN_profile_on_discovered_node_reapply_cluster_config":    core.NewMethodIdentifier(interfaceIdentifier, "reapply_TN_profile_on_discovered_node_reapply_cluster_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricNodesDiscoveredNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient) CreateTransportNodeForDiscoveredNodeCreateTransportNode(nodeExtIdParam string, transportNodeParam model.TransportNode) (model.TransportNode, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDiscoveredNodesCreateTransportNodeForDiscoveredNodeCreateTransportNodeInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	sv.AddStructField("TransportNode", transportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDiscoveredNodesCreateTransportNodeForDiscoveredNodeCreateTransportNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_discovered_nodes", "create_transport_node_for_discovered_node_create_transport_node", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDiscoveredNodesCreateTransportNodeForDiscoveredNodeCreateTransportNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient) HostPrepDiscoveredNodeHostprep(nodeExtIdParam string) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDiscoveredNodesHostPrepDiscoveredNodeHostprepInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDiscoveredNodesHostPrepDiscoveredNodeHostprepRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_discovered_nodes", "host_prep_discovered_node_hostprep", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDiscoveredNodesHostPrepDiscoveredNodeHostprepOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient) ListDiscoveredNodes(cmLocalIdParam *string, cursorParam *string, displayNameParam *string, externalIdParam *string, hasParentParam *string, includedFieldsParam *string, ipAddressParam *string, nodeIdParam *string, nodeTypeParam *string, originIdParam *string, pageSizeParam *int64, parentComputeCollectionParam *string, sortAscendingParam *bool, sortByParam *string) (model.DiscoveredNodeListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDiscoveredNodesListDiscoveredNodesInputType(), typeConverter)
	sv.AddStructField("CmLocalId", cmLocalIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DisplayName", displayNameParam)
	sv.AddStructField("ExternalId", externalIdParam)
	sv.AddStructField("HasParent", hasParentParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IpAddress", ipAddressParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("NodeType", nodeTypeParam)
	sv.AddStructField("OriginId", originIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ParentComputeCollection", parentComputeCollectionParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DiscoveredNodeListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDiscoveredNodesListDiscoveredNodesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_discovered_nodes", "list_discovered_nodes", inputDataValue, executionContext)
	var emptyOutput model.DiscoveredNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDiscoveredNodesListDiscoveredNodesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DiscoveredNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient) ReadDiscoveredNode(nodeExtIdParam string) (model.DiscoveredNode, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDiscoveredNodesReadDiscoveredNodeInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DiscoveredNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDiscoveredNodesReadDiscoveredNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_discovered_nodes", "read_discovered_node", inputDataValue, executionContext)
	var emptyOutput model.DiscoveredNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDiscoveredNodesReadDiscoveredNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DiscoveredNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesDiscoveredNodesClient) ReapplyTNProfileOnDiscoveredNodeReapplyClusterConfig(nodeExtIdParam string) (model.TransportNode, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesDiscoveredNodesReapplyTNProfileOnDiscoveredNodeReapplyClusterConfigInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesDiscoveredNodesReapplyTNProfileOnDiscoveredNodeReapplyClusterConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_discovered_nodes", "reapply_TN_profile_on_discovered_node_reapply_cluster_config", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesDiscoveredNodesReapplyTNProfileOnDiscoveredNodeReapplyClusterConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
