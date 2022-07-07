// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: DiscoveredNodes
// Used by client-side stubs.

package fabric

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type DiscoveredNodesClient interface {

	// NSX components are installaed on host and transport node is created with given configurations.
	//
	// @param nodeExtIdParam (required)
	// @param transportNodeParam (required)
	// @return com.vmware.nsx.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createtransportnode(nodeExtIdParam string, transportNodeParam model.TransportNode) (model.TransportNode, error)

	// Returns information about a specific discovered node.
	//
	// @param nodeExtIdParam (required)
	// @return com.vmware.nsx.model.DiscoveredNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(nodeExtIdParam string) (model.DiscoveredNode, error)

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
	// @return com.vmware.nsx.model.DiscoveredNodeListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cmLocalIdParam *string, cursorParam *string, displayNameParam *string, externalIdParam *string, hasParentParam *string, includedFieldsParam *string, ipAddressParam *string, nodeIdParam *string, nodeTypeParam *string, originIdParam *string, pageSizeParam *int64, parentComputeCollectionParam *string, sortAscendingParam *bool, sortByParam *string) (model.DiscoveredNodeListResult, error)

	// When transport node profile (TNP) is applied to a cluster, if any validation fails (e.g. VMs running on host) then transport node (TN) is not created. In that case after the required action is taken (e.g. VMs powered off), you can call this API to try to create TN for that discovered node. Do not call this API if Transport Node already exists for the discovered node. In that case use API on transport node. /transport-nodes/<transport-node-id>?action=restore_cluster_config
	//
	// @param nodeExtIdParam (required)
	// @return com.vmware.nsx.model.TransportNode
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Reapplyclusterconfig(nodeExtIdParam string) (model.TransportNode, error)
}

type discoveredNodesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewDiscoveredNodesClient(connector client.Connector) *discoveredNodesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.fabric.discovered_nodes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"createtransportnode":  core.NewMethodIdentifier(interfaceIdentifier, "createtransportnode"),
		"get":                  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                 core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reapplyclusterconfig": core.NewMethodIdentifier(interfaceIdentifier, "reapplyclusterconfig"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	dIface := discoveredNodesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *discoveredNodesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *discoveredNodesClient) Createtransportnode(nodeExtIdParam string, transportNodeParam model.TransportNode) (model.TransportNode, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(discoveredNodesCreatetransportnodeInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	sv.AddStructField("TransportNode", transportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := discoveredNodesCreatetransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.discovered_nodes", "createtransportnode", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), discoveredNodesCreatetransportnodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *discoveredNodesClient) Get(nodeExtIdParam string) (model.DiscoveredNode, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(discoveredNodesGetInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.DiscoveredNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := discoveredNodesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.discovered_nodes", "get", inputDataValue, executionContext)
	var emptyOutput model.DiscoveredNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), discoveredNodesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DiscoveredNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *discoveredNodesClient) List(cmLocalIdParam *string, cursorParam *string, displayNameParam *string, externalIdParam *string, hasParentParam *string, includedFieldsParam *string, ipAddressParam *string, nodeIdParam *string, nodeTypeParam *string, originIdParam *string, pageSizeParam *int64, parentComputeCollectionParam *string, sortAscendingParam *bool, sortByParam *string) (model.DiscoveredNodeListResult, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(discoveredNodesListInputType(), typeConverter)
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
	operationRestMetaData := discoveredNodesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.discovered_nodes", "list", inputDataValue, executionContext)
	var emptyOutput model.DiscoveredNodeListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), discoveredNodesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.DiscoveredNodeListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *discoveredNodesClient) Reapplyclusterconfig(nodeExtIdParam string) (model.TransportNode, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(discoveredNodesReapplyclusterconfigInputType(), typeConverter)
	sv.AddStructField("NodeExtId", nodeExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNode
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := discoveredNodesReapplyclusterconfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	dIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx.fabric.discovered_nodes", "reapplyclusterconfig", inputDataValue, executionContext)
	var emptyOutput model.TransportNode
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), discoveredNodesReapplyclusterconfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNode), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
