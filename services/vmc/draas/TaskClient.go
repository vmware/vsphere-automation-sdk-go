// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Task
// Used by client-side stubs.

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

const _ = core.SupportedByRuntimeVersion1

type TaskClient interface {

	// Retrieve details of a task.
	//
	// @param orgParam Organization identifier (required)
	// @param taskParam task identifier (required)
	// @return com.vmware.vmc.draas.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find the task with given identifier
	Get(orgParam string, taskParam string) (model.Task, error)

	// List all tasks with optional filtering.
	//
	// @param orgParam Organization identifier (required)
	// @param filterParam Filter expression Binary Operators: 'eq', 'ne', 'lt', 'gt', 'le', 'ge', 'mul', 'div', 'mod', 'sub', 'add' Unary Operators: 'not', '-' (minus) String Operators: 'startswith', 'endswith', 'length', 'contains', 'tolower', 'toupper', Nested attributes are composed using '.' Dates must be formatted as yyyy-MM-dd or yyyy-MM-ddTHH:mm:ss[.SSS]Z Strings should enclosed in single quotes, escape single quote with two single quotes The special literal 'created' will be mapped to the time the resource was first created. Examples: - $filter=(updated gt 2016-08-09T13:00:00Z) and (org_id eq 278710ff4e-6b6d-4d4e-aefb-ca637f38609e) - $filter=(created eq 2016-08-09) - $filter=(created gt 2016-08-09) and (sddc.status eq 'READY') (optional)
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	List(orgParam string, filterParam *string) ([]model.Task, error)

	// Request that a running task be canceled. This is advisory only, some tasks may not be cancelable, and some tasks might take an arbitrary amount of time to respond to a cancelation request. The task must be monitored to determine subsequent status.
	//
	// @param orgParam Organization identifier (required)
	// @param taskParam task identifier (required)
	// @param actionParam If = 'cancel', task will be canceled (optional)
	// @return com.vmware.vmc.draas.model.Task
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid action or bad argument
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Cannot find the task with given identifier
	Update(orgParam string, taskParam string, actionParam *string) (model.Task, error)
}

type taskClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTaskClient(connector client.Connector) *taskClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.draas.task")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := taskClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *taskClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *taskClient) Get(orgParam string, taskParam string) (model.Task, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(taskGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Task", taskParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := taskGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.task", "get", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), taskGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *taskClient) List(orgParam string, filterParam *string) ([]model.Task, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(taskListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Filter", filterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := taskListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.task", "list", inputDataValue, executionContext)
	var emptyOutput []model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), taskListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *taskClient) Update(orgParam string, taskParam string, actionParam *string) (model.Task, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(taskUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Task", taskParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := taskUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.task", "update", inputDataValue, executionContext)
	var emptyOutput model.Task
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), taskUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
