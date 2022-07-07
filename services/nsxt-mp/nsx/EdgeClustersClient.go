// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EdgeClusters
// Used by client-side stubs.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type EdgeClustersClient interface {

	// Creates a new edge cluster. It only supports homogeneous members. The TransportNodes backed by EdgeNode are only allowed in cluster members. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types.
	//
	// @param edgeClusterParam (required)
	// @return com.vmware.nsx.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error)

	// Deletes the specified edge cluster.
	//
	// @param edgeClusterIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(edgeClusterIdParam string) error

	// Returns information about the specified edge cluster.
	//
	// @param edgeClusterIdParam (required)
	// @return com.vmware.nsx.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(edgeClusterIdParam string) (model.EdgeCluster, error)

	// Returns information about the configured edge clusters, which enable you to group together transport nodes of the type EdgeNode and apply fabric profiles to all members of the edge cluster. Each edge node can participate in only one edge cluster.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.EdgeClusterListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EdgeClusterListResult, error)

	// Replace the transport node in the specified member of the edge-cluster. This is a disruptive action. This will move all the LogicalRouterPorts(uplink and routerLink) host on the old transport_node to the new transport_node. The transportNode cannot be present in another member of any edgeClusters.
	//
	// @param edgeClusterIdParam (required)
	// @param edgeClusterMemberTransportNodeParam (required)
	// @return com.vmware.nsx.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Replacetransportnode(edgeClusterIdParam string, edgeClusterMemberTransportNodeParam model.EdgeClusterMemberTransportNode) (model.EdgeCluster, error)

	// Modifies the specified edge cluster. Modifiable parameters include the description, display_name, transport-node-id. If the optional fabric_profile_binding is included, resource_type and profile_id are required. User should do a GET on the edge-cluster and obtain the payload and retain the member_index of the existing members as returning in the GET output. For new member additions, the member_index cannot be defined by the user, user can read the system allocated index to the new member in the output of this API call or by doing a GET call. User cannot use this PUT api to replace the transport_node of an existing member because this is a disruption action, we have exposed a explicit API for doing so, refer to \"ReplaceEdgeClusterMemberTransportNode\" EdgeCluster only supports homogeneous members. The TransportNodes backed by EdgeNode are only allowed in cluster members. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types.
	//
	// @param edgeClusterIdParam (required)
	// @param edgeClusterParam (required)
	// @return com.vmware.nsx.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(edgeClusterIdParam string, edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error)
}

type edgeClustersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewEdgeClustersClient(connector client.Connector) *edgeClustersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.edge_clusters")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":               core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":               core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":                  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                 core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"replacetransportnode": core.NewMethodIdentifier(interfaceIdentifier, "replacetransportnode"),
		"update":               core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	eIface := edgeClustersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *edgeClustersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *edgeClustersClient) Create(edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(edgeClustersCreateInputType(), typeConverter)
	sv.AddStructField("EdgeCluster", edgeClusterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgeClustersCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.edge_clusters", "create", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), edgeClustersCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeClustersClient) Delete(edgeClusterIdParam string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(edgeClustersDeleteInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgeClustersDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.edge_clusters", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *edgeClustersClient) Get(edgeClusterIdParam string) (model.EdgeCluster, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(edgeClustersGetInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgeClustersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.edge_clusters", "get", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), edgeClustersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeClustersClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EdgeClusterListResult, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(edgeClustersListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeClusterListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgeClustersListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.edge_clusters", "list", inputDataValue, executionContext)
	var emptyOutput model.EdgeClusterListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), edgeClustersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeClusterListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeClustersClient) Replacetransportnode(edgeClusterIdParam string, edgeClusterMemberTransportNodeParam model.EdgeClusterMemberTransportNode) (model.EdgeCluster, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(edgeClustersReplacetransportnodeInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("EdgeClusterMemberTransportNode", edgeClusterMemberTransportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgeClustersReplacetransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.edge_clusters", "replacetransportnode", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), edgeClustersReplacetransportnodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *edgeClustersClient) Update(edgeClusterIdParam string, edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(edgeClustersUpdateInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("EdgeCluster", edgeClusterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgeClustersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx.edge_clusters", "update", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), edgeClustersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
