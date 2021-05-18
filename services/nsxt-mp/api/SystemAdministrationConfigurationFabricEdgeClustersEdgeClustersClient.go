// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricEdgeClustersEdgeClusters
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient interface {

	// Creates a new edge cluster. It only supports homogeneous members. The TransportNodes backed by EdgeNode are only allowed in cluster members. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types.
	//
	// @param edgeClusterParam (required)
	// @return com.vmware.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateEdgeCluster(edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error)

	// Deletes the specified edge cluster.
	//
	// @param edgeClusterIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteEdgeCluster(edgeClusterIdParam string) error

	// Returns the allocation details of cluster and its members. Lists the edge node members, active and standby services of each node, utilization details of configured sub-pools. These allocation details can be monitored by customers to trigger migration of certain service contexts to different edge nodes, to balance the utilization of edge node resources.
	//
	// @param edgeClusterIdParam (required)
	// @return com.vmware.model.EdgeClusterAllocationStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetEdgeClusterAllocationStatus(edgeClusterIdParam string) (model.EdgeClusterAllocationStatus, error)

	// Return realized state information of a edge cluster. Any configuration update that affects the edge cluster can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of edge cluster.
	//
	// @param edgeClusterIdParam (required)
	// @param barrierIdParam (optional)
	// @param requestIdParam Realization request ID (optional)
	// @return com.vmware.model.EdgeClusterState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetEdgeClusterState(edgeClusterIdParam string, barrierIdParam *int64, requestIdParam *string) (model.EdgeClusterState, error)

	// Returns the aggregated status for the Edge cluster along with status of all edge nodes in the cluster. Query parameter \"source=realtime\" is the only supported source.
	//
	// @param edgeClusterIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.EdgeClusterStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetEdgeClusterStatus(edgeClusterIdParam string, sourceParam *string) (model.EdgeClusterStatus, error)

	// Returns information about the configured edge clusters, which enable you to group together transport nodes of the type EdgeNode and apply fabric profiles to all members of the edge cluster. Each edge node can participate in only one edge cluster.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.EdgeClusterListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListEdgeClusters(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EdgeClusterListResult, error)

	// Returns information about the specified edge cluster.
	//
	// @param edgeClusterIdParam (required)
	// @return com.vmware.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadEdgeCluster(edgeClusterIdParam string) (model.EdgeCluster, error)

	// Replace the transport node in the specified member of the edge-cluster. This is a disruptive action. This will move all the LogicalRouterPorts(uplink and routerLink) host on the old transport_node to the new transport_node. The transportNode cannot be present in another member of any edgeClusters.
	//
	// @param edgeClusterIdParam (required)
	// @param edgeClusterMemberTransportNodeParam (required)
	// @return com.vmware.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReplaceEdgeClusterMemberTransportNodeReplaceTransportNode(edgeClusterIdParam string, edgeClusterMemberTransportNodeParam model.EdgeClusterMemberTransportNode) (model.EdgeCluster, error)

	// Modifies the specified edge cluster. Modifiable parameters include the description, display_name, transport-node-id. If the optional fabric_profile_binding is included, resource_type and profile_id are required. User should do a GET on the edge-cluster and obtain the payload and retain the member_index of the existing members as returning in the GET output. For new member additions, the member_index cannot be defined by the user, user can read the system allocated index to the new member in the output of this API call or by doing a GET call. User cannot use this PUT api to replace the transport_node of an existing member because this is a disruption action, we have exposed a explicit API for doing so, refer to \"ReplaceEdgeClusterMemberTransportNode\" EdgeCluster only supports homogeneous members. The TransportNodes backed by EdgeNode are only allowed in cluster members. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types.
	//
	// @param edgeClusterIdParam (required)
	// @param edgeClusterParam (required)
	// @return com.vmware.model.EdgeCluster
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateEdgeCluster(edgeClusterIdParam string, edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error)
}

type systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient(connector client.Connector) *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_edge_cluster":                core.NewMethodIdentifier(interfaceIdentifier, "create_edge_cluster"),
		"delete_edge_cluster":                core.NewMethodIdentifier(interfaceIdentifier, "delete_edge_cluster"),
		"get_edge_cluster_allocation_status": core.NewMethodIdentifier(interfaceIdentifier, "get_edge_cluster_allocation_status"),
		"get_edge_cluster_state":             core.NewMethodIdentifier(interfaceIdentifier, "get_edge_cluster_state"),
		"get_edge_cluster_status":            core.NewMethodIdentifier(interfaceIdentifier, "get_edge_cluster_status"),
		"list_edge_clusters":                 core.NewMethodIdentifier(interfaceIdentifier, "list_edge_clusters"),
		"read_edge_cluster":                  core.NewMethodIdentifier(interfaceIdentifier, "read_edge_cluster"),
		"replace_edge_cluster_member_transport_node_replace_transport_node": core.NewMethodIdentifier(interfaceIdentifier, "replace_edge_cluster_member_transport_node_replace_transport_node"),
		"update_edge_cluster": core.NewMethodIdentifier(interfaceIdentifier, "update_edge_cluster"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) CreateEdgeCluster(edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersCreateEdgeClusterInputType(), typeConverter)
	sv.AddStructField("EdgeCluster", edgeClusterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersCreateEdgeClusterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "create_edge_cluster", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersCreateEdgeClusterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) DeleteEdgeCluster(edgeClusterIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersDeleteEdgeClusterInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersDeleteEdgeClusterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "delete_edge_cluster", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) GetEdgeClusterAllocationStatus(edgeClusterIdParam string) (model.EdgeClusterAllocationStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterAllocationStatusInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeClusterAllocationStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterAllocationStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "get_edge_cluster_allocation_status", inputDataValue, executionContext)
	var emptyOutput model.EdgeClusterAllocationStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterAllocationStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeClusterAllocationStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) GetEdgeClusterState(edgeClusterIdParam string, barrierIdParam *int64, requestIdParam *string) (model.EdgeClusterState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStateInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("BarrierId", barrierIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeClusterState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "get_edge_cluster_state", inputDataValue, executionContext)
	var emptyOutput model.EdgeClusterState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeClusterState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) GetEdgeClusterStatus(edgeClusterIdParam string, sourceParam *string) (model.EdgeClusterStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStatusInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeClusterStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "get_edge_cluster_status", inputDataValue, executionContext)
	var emptyOutput model.EdgeClusterStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersGetEdgeClusterStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeClusterStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) ListEdgeClusters(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.EdgeClusterListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersListEdgeClustersInputType(), typeConverter)
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
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersListEdgeClustersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "list_edge_clusters", inputDataValue, executionContext)
	var emptyOutput model.EdgeClusterListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersListEdgeClustersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeClusterListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) ReadEdgeCluster(edgeClusterIdParam string) (model.EdgeCluster, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReadEdgeClusterInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReadEdgeClusterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "read_edge_cluster", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReadEdgeClusterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) ReplaceEdgeClusterMemberTransportNodeReplaceTransportNode(edgeClusterIdParam string, edgeClusterMemberTransportNodeParam model.EdgeClusterMemberTransportNode) (model.EdgeCluster, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReplaceEdgeClusterMemberTransportNodeReplaceTransportNodeInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("EdgeClusterMemberTransportNode", edgeClusterMemberTransportNodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReplaceEdgeClusterMemberTransportNodeReplaceTransportNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "replace_edge_cluster_member_transport_node_replace_transport_node", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersReplaceEdgeClusterMemberTransportNodeReplaceTransportNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricEdgeClustersEdgeClustersClient) UpdateEdgeCluster(edgeClusterIdParam string, edgeClusterParam model.EdgeCluster) (model.EdgeCluster, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricEdgeClustersEdgeClustersUpdateEdgeClusterInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("EdgeCluster", edgeClusterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeCluster
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricEdgeClustersEdgeClustersUpdateEdgeClusterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_edge_clusters_edge_clusters", "update_edge_cluster", inputDataValue, executionContext)
	var emptyOutput model.EdgeCluster
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricEdgeClustersEdgeClustersUpdateEdgeClusterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeCluster), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
