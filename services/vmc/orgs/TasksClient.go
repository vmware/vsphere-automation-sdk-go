// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Tasks
// Used by client-side stubs.

package orgs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TasksClient interface {

	// Retrieve details of a task.
	//
	// @param orgParam Organization identifier (required)
	// @param taskParam Task identifier (required)
	// @return com.vmware.vmc.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find the task with given identifier
	Get(orgParam string, taskParam string) (vmcModel.Task, error)

	// List all tasks with optional filtering.
	//
	// @param orgParam Organization identifier (required)
	// @param filterParam Filter expression Binary Operators: 'eq', 'ne', 'lt', 'gt', 'le', 'ge', 'mul', 'div', 'mod', 'sub', 'add' Unary Operators: 'not', '-' (minus) String Operators: 'startswith', 'endswith', 'length', 'contains', 'tolower', 'toupper', Nested attributes are composed using '.' Dates must be formatted as yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss[.SSS]Z Strings should enclosed in single quotes, escape single quote with two single quotes The special literal 'created' will be mapped to the time the resource was first created. Examples: - $filter=(updated gt 2016-08-09T13:00:00Z) and (org_id eq 278710ff4e-6b6d-4d4e-aefb-ca637f38609e) - $filter=(created eq 2016-08-09) - $filter=(created gt 2016-08-09) and (sddc.status eq 'READY') (optional)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	List(orgParam string, filterParam *string) ([]vmcModel.Task, error)

	// Request that a running task be canceled. This is advisory only, some tasks may not be cancelable, and some tasks might take an arbitrary amount of time to respond to a cancelation request. The task must be monitored to determine subsequent status.
	//
	// @param orgParam Organization identifier (required)
	// @param taskParam Task identifier (required)
	// @param actionParam If = 'cancel', task will be canceled (optional)
	// @return com.vmware.vmc.model.Task
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find the task with given identifier
	Update(orgParam string, taskParam string, actionParam *string) (vmcModel.Task, error)
}

type tasksClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTasksClient(connector vapiProtocolClient_.Connector) *tasksClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.tasks")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
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

func (tIface *tasksClient) Get(orgParam string, taskParam string) (vmcModel.Task, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Task", taskParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tasks", "get", inputDataValue, executionContext)
	var emptyOutput vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tasksClient) List(orgParam string, filterParam *string) ([]vmcModel.Task, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Filter", filterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tasks", "list", inputDataValue, executionContext)
	var emptyOutput []vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tasksClient) Update(orgParam string, taskParam string, actionParam *string) (vmcModel.Task, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := tasksUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(tasksUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Task", taskParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.Task
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tasks", "update", inputDataValue, executionContext)
	var emptyOutput vmcModel.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TasksUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
