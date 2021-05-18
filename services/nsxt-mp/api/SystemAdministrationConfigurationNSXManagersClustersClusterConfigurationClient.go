// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationNSXManagersClustersClusterConfiguration
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

type SystemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient interface {

	// Add a new controller to the NSX cluster. Deprecated. Use POST /cluster?action=join_cluster to join a node to cluster. The controller comes with the new node.
	//
	// @param addClusterNodeSpecParam (required)
	// @param actionParam (required)
	// @return com.vmware.model.ClusterNodeConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddClusterNode(addClusterNodeSpecParam model.AddClusterNodeSpec, actionParam string) (model.ClusterNodeConfig, error)

	//
	//
	// @param nodeIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteClusterNodeConfig(nodeIdParam string) error

	// Detach a node from the Cluster
	//
	// @param nodeIdParam UUID of the node (required)
	// @param forceParam (optional)
	// @param gracefulShutdownParam (optional, default to false)
	// @param ignoreRepositoryIpCheckParam (optional, default to false)
	// @return com.vmware.model.ClusterConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DetachClusterNodeRemoveNode(nodeIdParam string, forceParam *string, gracefulShutdownParam *string, ignoreRepositoryIpCheckParam *string) (model.ClusterConfiguration, error)

	// Returns list of backup frames and some metadata to be used by UI.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param uiTabTypeParam (optional, default to LOCAL_MANAGER_TAB)
	// @return com.vmware.model.BackupUiFramesInfoList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBackupUiFramesInfo(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, uiTabTypeParam *string) (model.BackupUiFramesInfoList, error)

	// Returns information about the specified NSX cluster node.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.ClusterNodeInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetClusterNodeConfig(nodeIdParam string) (model.ClusterNodeInfo, error)

	// Invoke DELETE request on target cluster node
	//
	// @param targetNodeIdParam Target node UUID or keyword self (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokeDeleteClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Invoke GET request on target cluster node
	//
	// @param targetNodeIdParam Target node UUID or keyword self (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokeGetClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Invoke POST request on target cluster node
	//
	// @param targetNodeIdParam Target node UUID or keyword self (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokePostClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Invoke PUT request on target cluster node
	//
	// @param targetNodeIdParam Target node UUID or keyword self (required)
	// @param targetUriParam URI of API to invoke on target node (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	InvokePutClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error

	// Join this node to a NSX Cluster
	//
	// @param joinClusterParametersParam (required)
	// @return com.vmware.model.ClusterConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	JoinClusterJoinCluster(joinClusterParametersParam model.JoinClusterParameters) (model.ClusterConfiguration, error)

	// Returns information about all NSX cluster nodes.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.ClusterNodeConfigListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListClusterNodeConfigs(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ClusterNodeConfigListResult, error)

	// Returns information about the NSX cluster configuration. An NSX cluster has two functions or purposes, commonly referred to as \"roles.\" These two roles are control and management. Each NSX installation has a single cluster. Separate NSX clusters do not share data. In other words, a given data-plane node is attached to only one cluster, not to multiple clusters.
	// @return com.vmware.model.ClusterConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadClusterConfig() (model.ClusterConfig, error)

	//
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.ClusterNodeConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadClusterNodeConfig(nodeIdParam string) (model.ClusterNodeConfig, error)
}

type systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient(connector client.Connector) *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_cluster_node":                  core.NewMethodIdentifier(interfaceIdentifier, "add_cluster_node"),
		"delete_cluster_node_config":        core.NewMethodIdentifier(interfaceIdentifier, "delete_cluster_node_config"),
		"detach_cluster_node_remove_node":   core.NewMethodIdentifier(interfaceIdentifier, "detach_cluster_node_remove_node"),
		"get_backup_ui_frames_info":         core.NewMethodIdentifier(interfaceIdentifier, "get_backup_ui_frames_info"),
		"get_cluster_node_config":           core.NewMethodIdentifier(interfaceIdentifier, "get_cluster_node_config"),
		"invoke_delete_cluster_central_API": core.NewMethodIdentifier(interfaceIdentifier, "invoke_delete_cluster_central_API"),
		"invoke_get_cluster_central_API":    core.NewMethodIdentifier(interfaceIdentifier, "invoke_get_cluster_central_API"),
		"invoke_post_cluster_central_API":   core.NewMethodIdentifier(interfaceIdentifier, "invoke_post_cluster_central_API"),
		"invoke_put_cluster_central_API":    core.NewMethodIdentifier(interfaceIdentifier, "invoke_put_cluster_central_API"),
		"join_cluster_join_cluster":         core.NewMethodIdentifier(interfaceIdentifier, "join_cluster_join_cluster"),
		"list_cluster_node_configs":         core.NewMethodIdentifier(interfaceIdentifier, "list_cluster_node_configs"),
		"read_cluster_config":               core.NewMethodIdentifier(interfaceIdentifier, "read_cluster_config"),
		"read_cluster_node_config":          core.NewMethodIdentifier(interfaceIdentifier, "read_cluster_node_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) AddClusterNode(addClusterNodeSpecParam model.AddClusterNodeSpec, actionParam string) (model.ClusterNodeConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationAddClusterNodeInputType(), typeConverter)
	sv.AddStructField("AddClusterNodeSpec", addClusterNodeSpecParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationAddClusterNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "add_cluster_node", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationAddClusterNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) DeleteClusterNodeConfig(nodeIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationDeleteClusterNodeConfigInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationDeleteClusterNodeConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "delete_cluster_node_config", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) DetachClusterNodeRemoveNode(nodeIdParam string, forceParam *string, gracefulShutdownParam *string, ignoreRepositoryIpCheckParam *string) (model.ClusterConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationDetachClusterNodeRemoveNodeInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("Force", forceParam)
	sv.AddStructField("GracefulShutdown", gracefulShutdownParam)
	sv.AddStructField("IgnoreRepositoryIpCheck", ignoreRepositoryIpCheckParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationDetachClusterNodeRemoveNodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "detach_cluster_node_remove_node", inputDataValue, executionContext)
	var emptyOutput model.ClusterConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationDetachClusterNodeRemoveNodeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) GetBackupUiFramesInfo(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, uiTabTypeParam *string) (model.BackupUiFramesInfoList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationGetBackupUiFramesInfoInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("UiTabType", uiTabTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BackupUiFramesInfoList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationGetBackupUiFramesInfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "get_backup_ui_frames_info", inputDataValue, executionContext)
	var emptyOutput model.BackupUiFramesInfoList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationGetBackupUiFramesInfoOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BackupUiFramesInfoList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) GetClusterNodeConfig(nodeIdParam string) (model.ClusterNodeInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationGetClusterNodeConfigInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationGetClusterNodeConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "get_cluster_node_config", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationGetClusterNodeConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) InvokeDeleteClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokeDeleteClusterCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokeDeleteClusterCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "invoke_delete_cluster_central_API", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) InvokeGetClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokeGetClusterCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokeGetClusterCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "invoke_get_cluster_central_API", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) InvokePostClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokePostClusterCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokePostClusterCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "invoke_post_cluster_central_API", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) InvokePutClusterCentralAPI(targetNodeIdParam string, targetUriParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokePutClusterCentralAPIInputType(), typeConverter)
	sv.AddStructField("TargetNodeId", targetNodeIdParam)
	sv.AddStructField("TargetUri", targetUriParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationInvokePutClusterCentralAPIRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "invoke_put_cluster_central_API", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) JoinClusterJoinCluster(joinClusterParametersParam model.JoinClusterParameters) (model.ClusterConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationJoinClusterJoinClusterInputType(), typeConverter)
	sv.AddStructField("JoinClusterParameters", joinClusterParametersParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationJoinClusterJoinClusterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "join_cluster_join_cluster", inputDataValue, executionContext)
	var emptyOutput model.ClusterConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationJoinClusterJoinClusterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) ListClusterNodeConfigs(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ClusterNodeConfigListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationListClusterNodeConfigsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeConfigListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationListClusterNodeConfigsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "list_cluster_node_configs", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationListClusterNodeConfigsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) ReadClusterConfig() (model.ClusterConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationReadClusterConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationReadClusterConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "read_cluster_config", inputDataValue, executionContext)
	var emptyOutput model.ClusterConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationReadClusterConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterConfigurationClient) ReadClusterNodeConfig(nodeIdParam string) (model.ClusterNodeConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterConfigurationReadClusterNodeConfigInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterConfigurationReadClusterNodeConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_configuration", "read_cluster_node_config", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterConfigurationReadClusterNodeConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
