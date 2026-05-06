// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Telemetry
// Used by client-side stubs.

package services

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TelemetryClient interface {

	// Read Telemetry service properties
	//  This API is deprecated. Use:
	// GET https://<nsx-mgr>/api/v1/node/services/nsxt-mp/monitoring
	//
	// Deprecated: This API element is deprecated.
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.NodePhonehomeCoordinatorServiceProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(clusterNodeIdParam string) (nsxModel.NodePhonehomeCoordinatorServiceProperties, error)

	// Reset the logging levels to default values
	//  This API is deprecated. Use:
	// POST https://<nsx-mgr>/api/v1/node/services/nsxt-mp/monitoring?action=reset-monitoring-logging-levels
	//
	// Deprecated: This API element is deprecated.
	//
	// @param clusterNodeIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resettelemetrylogginglevels(clusterNodeIdParam string) error

	// Restart, start or stop Telemetry service
	//  This API is deprecated. Use:
	// POST https://<nsx-mgr>/api/v1/node/services/nsxt-mp/monitoring?action=restart|start|stop
	//
	// Deprecated: This API element is deprecated.
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.NodeServiceStatusProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Restart(clusterNodeIdParam string) (nsxModel.NodeServiceStatusProperties, error)

	// Restart, start or stop Telemetry service
	//  This API is deprecated. Use:
	// POST https://<nsx-mgr>/api/v1/node/services/nsxt-mp/monitoring?action=restart|start|stop
	//
	// Deprecated: This API element is deprecated.
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.NodeServiceStatusProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Start(clusterNodeIdParam string) (nsxModel.NodeServiceStatusProperties, error)

	// Restart, start or stop Telemetry service
	//  This API is deprecated. Use:
	// POST https://<nsx-mgr>/api/v1/node/services/nsxt-mp/monitoring?action=restart|start|stop
	//
	// Deprecated: This API element is deprecated.
	//
	// @param clusterNodeIdParam (required)
	// @return com.vmware.nsx.model.NodeServiceStatusProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Stop(clusterNodeIdParam string) (nsxModel.NodeServiceStatusProperties, error)

	// Update Telemetry service properties
	//  This API is deprecated. Use:
	// PUT https://<nsx-mgr>/api/v1/node/services/nsxt-mp/monitoring
	//
	// Deprecated: This API element is deprecated.
	//
	// @param clusterNodeIdParam (required)
	// @param nodePhonehomeCoordinatorServicePropertiesParam (required)
	// @return com.vmware.nsx.model.NodePhonehomeCoordinatorServiceProperties
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(clusterNodeIdParam string, nodePhonehomeCoordinatorServicePropertiesParam nsxModel.NodePhonehomeCoordinatorServiceProperties) (nsxModel.NodePhonehomeCoordinatorServiceProperties, error)
}

type telemetryClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTelemetryClient(connector vapiProtocolClient_.Connector) *telemetryClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.cluster.node.services.telemetry")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":                         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"resettelemetrylogginglevels": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "resettelemetrylogginglevels"),
		"restart":                     vapiCore_.NewMethodIdentifier(interfaceIdentifier, "restart"),
		"start":                       vapiCore_.NewMethodIdentifier(interfaceIdentifier, "start"),
		"stop":                        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "stop"),
		"update":                      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := telemetryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *telemetryClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *telemetryClient) Get(clusterNodeIdParam string) (nsxModel.NodePhonehomeCoordinatorServiceProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := telemetryGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(telemetryGetInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.NodePhonehomeCoordinatorServiceProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.telemetry", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.NodePhonehomeCoordinatorServiceProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TelemetryGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.NodePhonehomeCoordinatorServiceProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *telemetryClient) Resettelemetrylogginglevels(clusterNodeIdParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := telemetryResettelemetrylogginglevelsRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(telemetryResettelemetrylogginglevelsInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.telemetry", "resettelemetrylogginglevels", inputDataValue, executionContext)
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

func (tIface *telemetryClient) Restart(clusterNodeIdParam string) (nsxModel.NodeServiceStatusProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := telemetryRestartRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(telemetryRestartInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.NodeServiceStatusProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.telemetry", "restart", inputDataValue, executionContext)
	var emptyOutput nsxModel.NodeServiceStatusProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TelemetryRestartOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.NodeServiceStatusProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *telemetryClient) Start(clusterNodeIdParam string) (nsxModel.NodeServiceStatusProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := telemetryStartRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(telemetryStartInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.NodeServiceStatusProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.telemetry", "start", inputDataValue, executionContext)
	var emptyOutput nsxModel.NodeServiceStatusProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TelemetryStartOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.NodeServiceStatusProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *telemetryClient) Stop(clusterNodeIdParam string) (nsxModel.NodeServiceStatusProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := telemetryStopRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(telemetryStopInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.NodeServiceStatusProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.telemetry", "stop", inputDataValue, executionContext)
	var emptyOutput nsxModel.NodeServiceStatusProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TelemetryStopOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.NodeServiceStatusProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *telemetryClient) Update(clusterNodeIdParam string, nodePhonehomeCoordinatorServicePropertiesParam nsxModel.NodePhonehomeCoordinatorServiceProperties) (nsxModel.NodePhonehomeCoordinatorServiceProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := telemetryUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(telemetryUpdateInputType(), typeConverter)
	sv.AddStructField("ClusterNodeId", clusterNodeIdParam)
	sv.AddStructField("NodePhonehomeCoordinatorServiceProperties", nodePhonehomeCoordinatorServicePropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.NodePhonehomeCoordinatorServiceProperties
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.cluster.node.services.telemetry", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.NodePhonehomeCoordinatorServiceProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TelemetryUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.NodePhonehomeCoordinatorServiceProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
