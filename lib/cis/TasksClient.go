// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Tasks
// Used by client-side stubs.

package cis

import (
	cisTask_ "github.com/vmware/vsphere-automation-sdk-go/lib/cis/task"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Tasks`` interface provides methods for managing the task related to a long running operation.
type TasksClient interface {

	// Returns information about a task.
	//
	// @param taskParam Task identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
	// @param specParam Specification on what to get for a task.
	// If null, the behavior is equivalent to a TasksGetSpec with all properties null which means only the data described in cisTask_.Info will be returned and the result of the operation will be return.
	// @return Information about the specified task.
	//
	// @throws Error if the system reports an error while responding to the request.
	// @throws NotFound if the task is not found.
	// @throws ResourceInaccessible if the task's state cannot be accessed.
	// @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
	// @throws Unauthenticated if the user can not be authenticated.
	// @throws Unauthorized if the user doesn't have the required privileges.
	Get(taskParam string, specParam *TasksGetSpec) (cisTask_.Info, error)

	// Returns information about at most 1000 visible (subject to permission checks) tasks matching the TasksFilterSpec. All tasks must be in the same provider.
	//
	// @param filterSpecParam Specification of matching tasks.
	// This is currently required. In the future, if it is null, the behavior is equivalent to a TasksFilterSpec with all properties null which means all tasks match the filter.
	// @param resultSpecParam Specification of what to return for a task.
	// If null, the behavior is equivalent to a TasksGetSpec with all properties null which means only the data describe in cisTask_.Info will be returned and the result of the operation will be return.
	// @return Map of task identifier to information about the task.
	// The key in the return value map will be an identifier for the resource type: ``com.vmware.cis.task``.
	//
	// @throws InvalidArgument if any of the specified parameters are invalid.
	// @throws ResourceInaccessible if a task's state cannot be accessed or over 1000 tasks matching the TasksFilterSpec.
	// @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
	// @throws Unauthenticated if the user can not be authenticated.
	// @throws Unauthorized if the user doesn't have the required privileges.
	List(filterSpecParam *TasksFilterSpec, resultSpecParam *TasksGetSpec) (map[string]cisTask_.Info, error)

	// Cancel a running operation associated with the task. This is the best effort attempt. Operation may not be cancelled anymore once it reaches certain stage.
	//
	// @param taskParam Task identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.cis.task``.
	//
	// @throws Error if the system reports an error while responding to the request.
	// @throws NotAllowedInCurrentState if the task is already canceled or completed.
	// @throws NotFound if the task is not found.
	// @throws ResourceInaccessible if the task's state cannot be accessed.
	// @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
	// @throws Unauthenticated if the user can not be authenticated.
	// @throws Unauthorized if the user doesn't have the required privileges.
	// @throws Unsupported if the task is not cancelable.
	Cancel(taskParam string) error
}

type tasksClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTasksClient(connector vapiProtocolClient_.Connector) *tasksClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.cis.tasks")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"cancel": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "cancel"),
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

func (tIface *tasksClient) Get(taskParam string, specParam *TasksGetSpec) (cisTask_.Info, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksGetInputType(), typeConverter)
	sv.AddStructField("Task", taskParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput cisTask_.Info
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.cis.tasks", "get", inputDataValue, executionContext)
	var emptyOutput cisTask_.Info
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(cisTask_.Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tasksClient) List(filterSpecParam *TasksFilterSpec, resultSpecParam *TasksGetSpec) (map[string]cisTask_.Info, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksListInputType(), typeConverter)
	sv.AddStructField("FilterSpec", filterSpecParam)
	sv.AddStructField("ResultSpec", resultSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput map[string]cisTask_.Info
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.cis.tasks", "list", inputDataValue, executionContext)
	var emptyOutput map[string]cisTask_.Info
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string]cisTask_.Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tasksClient) Cancel(taskParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksCancelRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksCancelInputType(), typeConverter)
	sv.AddStructField("Task", taskParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.cis.tasks", "cancel", inputDataValue, executionContext)
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
