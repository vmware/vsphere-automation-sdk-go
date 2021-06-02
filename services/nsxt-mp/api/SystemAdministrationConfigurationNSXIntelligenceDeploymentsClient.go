// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationNSXIntelligenceDeployments
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationNSXIntelligenceDeploymentsClient interface {

	// Deploys a Intelligence cluster node VM as specified by the deployment config.
	//
	// @param addIntelligenceClusterNodeVMInfoParam (required)
	// @return com.vmware.model.IntelligenceClusterNodeVMDeploymentRequestList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	AddPaceClusterNodeVM(addIntelligenceClusterNodeVMInfoParam model.AddIntelligenceClusterNodeVMInfo) (model.IntelligenceClusterNodeVMDeploymentRequestList, error)

	// Attempts to unregister and undeploy a specified auto-deployed cluster node VM. If it is a member of a cluster, then the VM will be automatically detached from the cluster before being unregistered and undeployed. Alternatively, if the original deployment attempt failed or the VM is not found, cleans up the deployment information associated with the deployment attempt. Note: If a VM has been successfully auto-deployed, then the associated deployment information will not be deleted unless and until the VM is successfully deleted.
	//
	// @param nodeIdParam (required)
	// @param forceDeleteParam Delete by force (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteAutoDeployedPaceClusterNodeVMDelete(nodeIdParam string, forceDeleteParam *bool) error

	// Returns request information for every attempted deployment of a cluster node VM.
	// @return com.vmware.model.IntelligenceClusterNodeVMDeploymentRequestList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListPaceClusterNodeVMDeploymentRequests() (model.IntelligenceClusterNodeVMDeploymentRequestList, error)

	// Returns deployment request information for a specific attempted deployment of a cluster node VM.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.IntelligenceClusterNodeVMDeploymentRequest
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadPaceClusterNodeVMDeploymentRequest(nodeIdParam string) (model.IntelligenceClusterNodeVMDeploymentRequest, error)

	// Returns the current deployment or undeployment status for a VM along with any other relevant current information, such as error messages.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.model.IntelligenceClusterNodeVMDeploymentStatusReport
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadPaceClusterNodeVMDeploymentStatus(nodeIdParam string) (model.IntelligenceClusterNodeVMDeploymentStatusReport, error)
}

type systemAdministrationConfigurationNSXIntelligenceDeploymentsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationNSXIntelligenceDeploymentsClient(connector client.Connector) *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_NSX_intelligence_deployments")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"add_pace_cluster_node_VM":                         core.NewMethodIdentifier(interfaceIdentifier, "add_pace_cluster_node_VM"),
		"delete_auto_deployed_pace_cluster_node_VM_delete": core.NewMethodIdentifier(interfaceIdentifier, "delete_auto_deployed_pace_cluster_node_VM_delete"),
		"list_pace_cluster_node_VM_deployment_requests":    core.NewMethodIdentifier(interfaceIdentifier, "list_pace_cluster_node_VM_deployment_requests"),
		"read_pace_cluster_node_VM_deployment_request":     core.NewMethodIdentifier(interfaceIdentifier, "read_pace_cluster_node_VM_deployment_request"),
		"read_pace_cluster_node_VM_deployment_status":      core.NewMethodIdentifier(interfaceIdentifier, "read_pace_cluster_node_VM_deployment_status"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationNSXIntelligenceDeploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient) AddPaceClusterNodeVM(addIntelligenceClusterNodeVMInfoParam model.AddIntelligenceClusterNodeVMInfo) (model.IntelligenceClusterNodeVMDeploymentRequestList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceDeploymentsAddPaceClusterNodeVMInputType(), typeConverter)
	sv.AddStructField("AddIntelligenceClusterNodeVMInfo", addIntelligenceClusterNodeVMInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceDeploymentsAddPaceClusterNodeVMRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_deployments", "add_pace_cluster_node_VM", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceDeploymentsAddPaceClusterNodeVMOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient) DeleteAutoDeployedPaceClusterNodeVMDelete(nodeIdParam string, forceDeleteParam *bool) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceDeploymentsDeleteAutoDeployedPaceClusterNodeVMDeleteInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("ForceDelete", forceDeleteParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceDeploymentsDeleteAutoDeployedPaceClusterNodeVMDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_deployments", "delete_auto_deployed_pace_cluster_node_VM_delete", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient) ListPaceClusterNodeVMDeploymentRequests() (model.IntelligenceClusterNodeVMDeploymentRequestList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceDeploymentsListPaceClusterNodeVMDeploymentRequestsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceDeploymentsListPaceClusterNodeVMDeploymentRequestsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_deployments", "list_pace_cluster_node_VM_deployment_requests", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceDeploymentsListPaceClusterNodeVMDeploymentRequestsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient) ReadPaceClusterNodeVMDeploymentRequest(nodeIdParam string) (model.IntelligenceClusterNodeVMDeploymentRequest, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceDeploymentsReadPaceClusterNodeVMDeploymentRequestInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequest
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceDeploymentsReadPaceClusterNodeVMDeploymentRequestRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_deployments", "read_pace_cluster_node_VM_deployment_request", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceDeploymentsReadPaceClusterNodeVMDeploymentRequestOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXIntelligenceDeploymentsClient) ReadPaceClusterNodeVMDeploymentStatus(nodeIdParam string) (model.IntelligenceClusterNodeVMDeploymentStatusReport, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXIntelligenceDeploymentsReadPaceClusterNodeVMDeploymentStatusInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IntelligenceClusterNodeVMDeploymentStatusReport
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXIntelligenceDeploymentsReadPaceClusterNodeVMDeploymentStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_intelligence_deployments", "read_pace_cluster_node_VM_deployment_status", inputDataValue, executionContext)
	var emptyOutput model.IntelligenceClusterNodeVMDeploymentStatusReport
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXIntelligenceDeploymentsReadPaceClusterNodeVMDeploymentStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IntelligenceClusterNodeVMDeploymentStatusReport), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
