// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationNSXManagersClustersClusterNodeDeployments
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

type SystemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient interface {

	// Deploys a cluster node VM as specified by the deployment config. Once the VM is deployed and powered on, it will automatically join the existing cluster.
	//
	// @param addClusterNodeVMInfoParam (required)
	// @return com.vmware.model.ClusterNodeVMDeploymentRequestList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddClusterNodeVM(addClusterNodeVMInfoParam model.AddClusterNodeVMInfo) (model.ClusterNodeVMDeploymentRequestList, error)

	// Attempts to unregister and undeploy a specified auto-deployed cluster node VM. If it is a member of a cluster, then the VM will be automatically detached from the cluster before being unregistered and undeployed. Alternatively, if the original deployment attempt failed or the VM is not found, cleans up the deployment information associated with the deployment attempt. Note: If a VM has been successfully auto-deployed, then the associated deployment information will not be deleted unless and until the VM is successfully deleted.
	//
	// @param nodeIdParam (required)
	// @param forceDeleteParam Delete by force (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteAutoDeployedClusterNodeVMDelete(nodeIdParam string, forceDeleteParam *bool) error

	//
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.RepoSyncStatusReport
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetRepoSyncStatus(nodeIdParam string) (model.RepoSyncStatusReport, error)

	// Returns request information for every attempted deployment of a cluster node VM.
	// @return com.vmware.model.ClusterNodeVMDeploymentRequestList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListClusterNodeVMDeploymentRequests() (model.ClusterNodeVMDeploymentRequestList, error)

	// Attempts to synchronize the repository partition on nsx manager. Repository partition contains packages required for the install and upgrade of nsx components.Normally there is no need to call this API explicitely by the user.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	PerformRepoSyncRepoSync() error

	// Returns deployment request information for a specific attempted deployment of a cluster node VM.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.ClusterNodeVMDeploymentRequest
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadClusterNodeVMDeploymentRequest(nodeIdParam string) (model.ClusterNodeVMDeploymentRequest, error)

	// Returns the current deployment or undeployment status for a VM along with any other relevant current information, such as error messages.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.ClusterNodeVMDeploymentStatusReport
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadClusterNodeVMDeploymentStatus(nodeIdParam string) (model.ClusterNodeVMDeploymentStatusReport, error)
}

type systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient(connector client.Connector) *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_cluster_node_VM":                         core.NewMethodIdentifier(interfaceIdentifier, "add_cluster_node_VM"),
		"delete_auto_deployed_cluster_node_VM_delete": core.NewMethodIdentifier(interfaceIdentifier, "delete_auto_deployed_cluster_node_VM_delete"),
		"get_repo_sync_status":                        core.NewMethodIdentifier(interfaceIdentifier, "get_repo_sync_status"),
		"list_cluster_node_VM_deployment_requests":    core.NewMethodIdentifier(interfaceIdentifier, "list_cluster_node_VM_deployment_requests"),
		"perform_repo_sync_repo_sync":                 core.NewMethodIdentifier(interfaceIdentifier, "perform_repo_sync_repo_sync"),
		"read_cluster_node_VM_deployment_request":     core.NewMethodIdentifier(interfaceIdentifier, "read_cluster_node_VM_deployment_request"),
		"read_cluster_node_VM_deployment_status":      core.NewMethodIdentifier(interfaceIdentifier, "read_cluster_node_VM_deployment_status"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) AddClusterNodeVM(addClusterNodeVMInfoParam model.AddClusterNodeVMInfo) (model.ClusterNodeVMDeploymentRequestList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsAddClusterNodeVMInputType(), typeConverter)
	sv.AddStructField("AddClusterNodeVMInfo", addClusterNodeVMInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeVMDeploymentRequestList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsAddClusterNodeVMRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "add_cluster_node_VM", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsAddClusterNodeVMOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) DeleteAutoDeployedClusterNodeVMDelete(nodeIdParam string, forceDeleteParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsDeleteAutoDeployedClusterNodeVMDeleteInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("ForceDelete", forceDeleteParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsDeleteAutoDeployedClusterNodeVMDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "delete_auto_deployed_cluster_node_VM_delete", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) GetRepoSyncStatus(nodeIdParam string) (model.RepoSyncStatusReport, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsGetRepoSyncStatusInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RepoSyncStatusReport
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsGetRepoSyncStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "get_repo_sync_status", inputDataValue, executionContext)
	var emptyOutput model.RepoSyncStatusReport
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsGetRepoSyncStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RepoSyncStatusReport), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) ListClusterNodeVMDeploymentRequests() (model.ClusterNodeVMDeploymentRequestList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsListClusterNodeVMDeploymentRequestsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeVMDeploymentRequestList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsListClusterNodeVMDeploymentRequestsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "list_cluster_node_VM_deployment_requests", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsListClusterNodeVMDeploymentRequestsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) PerformRepoSyncRepoSync() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsPerformRepoSyncRepoSyncInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsPerformRepoSyncRepoSyncRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "perform_repo_sync_repo_sync", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) ReadClusterNodeVMDeploymentRequest(nodeIdParam string) (model.ClusterNodeVMDeploymentRequest, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsReadClusterNodeVMDeploymentRequestInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeVMDeploymentRequest
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsReadClusterNodeVMDeploymentRequestRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "read_cluster_node_VM_deployment_request", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeVMDeploymentRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsReadClusterNodeVMDeploymentRequestOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeVMDeploymentRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsClient) ReadClusterNodeVMDeploymentStatus(nodeIdParam string) (model.ClusterNodeVMDeploymentStatusReport, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsReadClusterNodeVMDeploymentStatusInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ClusterNodeVMDeploymentStatusReport
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsReadClusterNodeVMDeploymentStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_clusters_cluster_node_deployments", "read_cluster_node_VM_deployment_status", inputDataValue, executionContext)
	var emptyOutput model.ClusterNodeVMDeploymentStatusReport
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersClustersClusterNodeDeploymentsReadClusterNodeVMDeploymentStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ClusterNodeVMDeploymentStatusReport), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
