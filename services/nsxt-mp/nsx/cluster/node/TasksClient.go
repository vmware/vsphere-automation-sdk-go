// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Tasks
// Used by client-side stubs.

package node

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TasksClient interface {

	// Cancel specified task
	//
	// @param taskIdParam ID of task to delete (required)
	// @param clusterNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Cancel(taskIdParam string, clusterNodeIdParam string) error

	// Delete task
	//
	// @param taskIdParam ID of task to delete (required)
	// @param clusterNodeIdParam (required)
	//
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(taskIdParam string, clusterNodeIdParam string) error

	// Read task properties
	//
	// @param taskIdParam ID of task to read (required)
	// @param clusterNodeIdParam (required)
	// @param suppressRedirectParam Do not return a redirect HTTP status. (optional, default to false)
	// @return com.vmware.nsx.model.ApplianceManagementTaskProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(taskIdParam string, clusterNodeIdParam string, suppressRedirectParam *bool) (nsxModel.ApplianceManagementTaskProperties, error)

	// List appliance management tasks
	//
	// @param clusterNodeIdParam (required)
	// @param fieldsParam Comma-separated field names to include in query result (optional)
	// @param requestMethodParam Comma-separated request methods to include in query result (optional)
	// @param requestPathParam Comma-separated request paths to include in query result (optional)
	// @param requestUriParam Comma-separated request URIs to include in query result (optional)
	// @param statusParam Comma-separated status values to include in query result (optional)
	// @param userParam Comma-separated user names to include in query result (optional)
	// @return com.vmware.nsx.model.ApplianceManagementTaskListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws UnableToAllocateResource  Insufficient Storage
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(clusterNodeIdParam string, fieldsParam *string, requestMethodParam *string, requestPathParam *string, requestUriParam *string, statusParam *string, userParam *string) (nsxModel.ApplianceManagementTaskListResult, error)
}

type tasksClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTasksClient(connector vapiProtocolClient_.Connector) *tasksClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.cluster.node.tasks")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"cancel": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "cancel"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := tasksClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *tasksClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *tasksClient) Cancel(taskIdParam string, clusterNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksCancelRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksCancelInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.tasks", "cancel", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tasksClient) Delete(taskIdParam string, clusterNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksDeleteInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.tasks", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *tasksClient) Get(taskIdParam string, clusterNodeIdParam string, suppressRedirectParam *bool) (nsxModel.ApplianceManagementTaskProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksGetInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	sv.AddStructField("SuppressRedirect", suppressRedirectParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ApplianceManagementTaskProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.tasks", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ApplianceManagementTaskProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ApplianceManagementTaskProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tasksClient) List(clusterNodeIdParam string, fieldsParam *string, requestMethodParam *string, requestPathParam *string, requestUriParam *string, statusParam *string, userParam *string) (nsxModel.ApplianceManagementTaskListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksListInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	sv.AddStructField("Fields", fieldsParam)
	sv.AddStructField("RequestMethod", requestMethodParam)
	sv.AddStructField("RequestPath", requestPathParam)
	sv.AddStructField("RequestUri", requestUriParam)
	sv.AddStructField("Status", statusParam)
	sv.AddStructField("User", userParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ApplianceManagementTaskListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.tasks", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ApplianceManagementTaskListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ApplianceManagementTaskListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
