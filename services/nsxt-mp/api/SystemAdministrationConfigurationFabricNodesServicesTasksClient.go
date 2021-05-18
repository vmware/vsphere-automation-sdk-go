// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricNodesServicesTasks
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

type SystemAdministrationConfigurationFabricNodesServicesTasksClient interface {

	// Cancel specified task
	//
	// @param taskIdParam ID of task to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CancelApplianceManagementTaskCancel(taskIdParam string) error

	// Delete task
	//
	// @param taskIdParam ID of task to delete (required)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteApplianceManagementTask(taskIdParam string) error

	// List appliance management tasks
	//
	// @param fieldsParam Fields to include in query results (optional)
	// @param requestMethodParam Request method(s) to include in query result (optional)
	// @param requestPathParam Request URI path(s) to include in query result (optional)
	// @param requestUriParam Request URI(s) to include in query result (optional)
	// @param statusParam Status(es) to include in query result (optional)
	// @param userParam Names of users to include in query result (optional)
	// @return com.vmware.model.ApplianceManagementTaskListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListApplianceManagementTasks(fieldsParam *string, requestMethodParam *string, requestPathParam *string, requestUriParam *string, statusParam *string, userParam *string) (model.ApplianceManagementTaskListResult, error)

	// Read task properties
	//
	// @param taskIdParam ID of task to read (required)
	// @param suppressRedirectParam Suppress redirect status if applicable (optional, default to false)
	// @return com.vmware.model.ApplianceManagementTaskProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadApplianceManagementTaskProperties(taskIdParam string, suppressRedirectParam *bool) (model.ApplianceManagementTaskProperties, error)

	// Read asynchronous task response
	//
	// @param taskIdParam ID of task to read (required)
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error, Bad Gateway
	// @throws NotFound  Not Found, Gone
	ReadAsyncApplianceManagementTaskResponse(taskIdParam string) error
}

type systemAdministrationConfigurationFabricNodesServicesTasksClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricNodesServicesTasksClient(connector client.Connector) *systemAdministrationConfigurationFabricNodesServicesTasksClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_nodes_services_tasks")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"cancel_appliance_management_task_cancel":       core.NewMethodIdentifier(interfaceIdentifier, "cancel_appliance_management_task_cancel"),
		"delete_appliance_management_task":              core.NewMethodIdentifier(interfaceIdentifier, "delete_appliance_management_task"),
		"list_appliance_management_tasks":               core.NewMethodIdentifier(interfaceIdentifier, "list_appliance_management_tasks"),
		"read_appliance_management_task_properties":     core.NewMethodIdentifier(interfaceIdentifier, "read_appliance_management_task_properties"),
		"read_async_appliance_management_task_response": core.NewMethodIdentifier(interfaceIdentifier, "read_async_appliance_management_task_response"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricNodesServicesTasksClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricNodesServicesTasksClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricNodesServicesTasksClient) CancelApplianceManagementTaskCancel(taskIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesServicesTasksCancelApplianceManagementTaskCancelInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesServicesTasksCancelApplianceManagementTaskCancelRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_services_tasks", "cancel_appliance_management_task_cancel", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricNodesServicesTasksClient) DeleteApplianceManagementTask(taskIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesServicesTasksDeleteApplianceManagementTaskInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesServicesTasksDeleteApplianceManagementTaskRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_services_tasks", "delete_appliance_management_task", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricNodesServicesTasksClient) ListApplianceManagementTasks(fieldsParam *string, requestMethodParam *string, requestPathParam *string, requestUriParam *string, statusParam *string, userParam *string) (model.ApplianceManagementTaskListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesServicesTasksListApplianceManagementTasksInputType(), typeConverter)
	sv.AddStructField("Fields", fieldsParam)
	sv.AddStructField("RequestMethod", requestMethodParam)
	sv.AddStructField("RequestPath", requestPathParam)
	sv.AddStructField("RequestUri", requestUriParam)
	sv.AddStructField("Status", statusParam)
	sv.AddStructField("User", userParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ApplianceManagementTaskListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesServicesTasksListApplianceManagementTasksRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_services_tasks", "list_appliance_management_tasks", inputDataValue, executionContext)
	var emptyOutput model.ApplianceManagementTaskListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesServicesTasksListApplianceManagementTasksOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ApplianceManagementTaskListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesServicesTasksClient) ReadApplianceManagementTaskProperties(taskIdParam string, suppressRedirectParam *bool) (model.ApplianceManagementTaskProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesServicesTasksReadApplianceManagementTaskPropertiesInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	sv.AddStructField("SuppressRedirect", suppressRedirectParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ApplianceManagementTaskProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesServicesTasksReadApplianceManagementTaskPropertiesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_services_tasks", "read_appliance_management_task_properties", inputDataValue, executionContext)
	var emptyOutput model.ApplianceManagementTaskProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricNodesServicesTasksReadApplianceManagementTaskPropertiesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ApplianceManagementTaskProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricNodesServicesTasksClient) ReadAsyncApplianceManagementTaskResponse(taskIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricNodesServicesTasksReadAsyncApplianceManagementTaskResponseInputType(), typeConverter)
	sv.AddStructField("TaskId", taskIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricNodesServicesTasksReadAsyncApplianceManagementTaskResponseRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_nodes_services_tasks", "read_async_appliance_management_task_response", inputDataValue, executionContext)
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
