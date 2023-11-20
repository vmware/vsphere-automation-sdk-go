// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Deployments
// Used by client-side stubs.

package controller_nodes

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type DeploymentsClient interface {

	// Deploys a Advanced Load Balancer controller node VM as specified by the deployment config.
	//
	// @param addALBControllerNodeVMInfoParam (required)
	// @return com.vmware.nsx_policy.model.ALBControllerNodeVMDeploymentRequestList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(addALBControllerNodeVMInfoParam nsx_policyModel.AddALBControllerNodeVMInfo) (nsx_policyModel.ALBControllerNodeVMDeploymentRequestList, error)

	// Attempts to unregister and undeploy a specified auto-deployed cluster node VM. If it is a member of a cluster, then the VM will be automatically detached from the cluster before being unregistered and undeployed. Alternatively, if the original deployment attempt failed or the VM is not found, cleans up the deployment information associated with the deployment attempt. Note: If a VM has been successfully auto-deployed, then the associated deployment information will not be deleted unless and until the VM is successfully deleted.
	//
	// @param nodeIdParam (required)
	// @param forceDeleteParam Delete by force (optional)
	// @param inaccessibleParam Delete when controller is inaccessible (optional)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(nodeIdParam string, forceDeleteParam *bool, inaccessibleParam *string) error

	// Returns deployment request information for a specific attempted deployment of a cluster node VM.
	//
	// @param nodeIdParam (required)
	// @return com.vmware.nsx_policy.model.ALBControllerNodeVMDeploymentRequest
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(nodeIdParam string) (nsx_policyModel.ALBControllerNodeVMDeploymentRequest, error)

	// Returns request information for every attempted deployment of a cluster node VM.
	//
	// @param clusteringIdParam Unique Id for NSX Advanced Load Balancer Controller Cluster used in VCF managed NSX. (optional)
	// @param stateParam the current state of the Advanced Load Balancer controller VM (optional)
	// @return com.vmware.nsx_policy.model.ALBControllerNodeVMDeploymentRequestList
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(clusteringIdParam *string, stateParam *string) (nsx_policyModel.ALBControllerNodeVMDeploymentRequestList, error)

	// Update Advanced Load Balancer Controller node VM details
	//
	// @param nodeIdParam (required)
	// @param aLBControllerNodeVMDeploymentRequestParam (required)
	// @param runningConfigParam Update Advanced Load Balancer Controller runtime config as well (optional)
	// @return com.vmware.nsx_policy.model.ALBControllerNodeVMDeploymentRequest
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(nodeIdParam string, aLBControllerNodeVMDeploymentRequestParam nsx_policyModel.ALBControllerNodeVMDeploymentRequest, runningConfigParam *bool) (nsx_policyModel.ALBControllerNodeVMDeploymentRequest, error)
}

type deploymentsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewDeploymentsClient(connector vapiProtocolClient_.Connector) *deploymentsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.alb.controller_nodes.deployments")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	dIface := deploymentsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &dIface
}

func (dIface *deploymentsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := dIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (dIface *deploymentsClient) Create(addALBControllerNodeVMInfoParam nsx_policyModel.AddALBControllerNodeVMInfo) (nsx_policyModel.ALBControllerNodeVMDeploymentRequestList, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsCreateInputType(), typeConverter)
	sv.AddStructField("AddALBControllerNodeVMInfo", addALBControllerNodeVMInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequestList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.deployments", "create", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ALBControllerNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentsClient) Delete(nodeIdParam string, forceDeleteParam *bool, inaccessibleParam *string) error {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsDeleteInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("ForceDelete", forceDeleteParam)
	sv.AddStructField("Inaccessible", inaccessibleParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.deployments", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (dIface *deploymentsClient) Get(nodeIdParam string) (nsx_policyModel.ALBControllerNodeVMDeploymentRequest, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsGetInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequest
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.deployments", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ALBControllerNodeVMDeploymentRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentsClient) List(clusteringIdParam *string, stateParam *string) (nsx_policyModel.ALBControllerNodeVMDeploymentRequestList, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsListInputType(), typeConverter)
	sv.AddStructField("ClusteringId", clusteringIdParam)
	sv.AddStructField("State", stateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequestList
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.deployments", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequestList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ALBControllerNodeVMDeploymentRequestList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (dIface *deploymentsClient) Update(nodeIdParam string, aLBControllerNodeVMDeploymentRequestParam nsx_policyModel.ALBControllerNodeVMDeploymentRequest, runningConfigParam *bool) (nsx_policyModel.ALBControllerNodeVMDeploymentRequest, error) {
	typeConverter := dIface.connector.TypeConverter()
	executionContext := dIface.connector.NewExecutionContext()
	operationRestMetaData := deploymentsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(deploymentsUpdateInputType(), typeConverter)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("ALBControllerNodeVMDeploymentRequest", aLBControllerNodeVMDeploymentRequestParam)
	sv.AddStructField("RunningConfig", runningConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequest
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := dIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.alb.controller_nodes.deployments", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.ALBControllerNodeVMDeploymentRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), DeploymentsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.ALBControllerNodeVMDeploymentRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), dIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
