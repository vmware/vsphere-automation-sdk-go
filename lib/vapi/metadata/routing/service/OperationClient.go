// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Operation
// Used by client-side stubs.

package service

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// Operations to retrieve information about routing in a vAPI operation
type OperationClient interface {

	// Get the IDs of all the vAPI operations in the given service that contain routing information
	//
	// @param serviceIdParam Identifier of the service
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
	// @return list of operation identifiers
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
	// @throws NotFound If the service identifier does not exist.
	List(serviceIdParam string) ([]string, error)

	// Get information about a vAPI operation that contains routing information
	//
	// @param serviceIdParam Identifier of the service
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
	// @param operationIdParam Identifier of the operation
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
	// @return Operation info for the vAPI operation that contains routing information.
	// @throws NotFound If the service identifier does not exist or if the specified operation identifier does not exist in the service.
	Get(serviceIdParam string, operationIdParam string) (routing.OperationInfo, error)
}

type operationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewOperationClient(connector client.Connector) *operationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.routing.service.operation")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	oIface := operationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *operationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *operationClient) List(serviceIdParam string) ([]string, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(operationListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := operationListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.service.operation", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), operationListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *operationClient) Get(serviceIdParam string, operationIdParam string) (routing.OperationInfo, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(operationGetInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("OperationId", operationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput routing.OperationInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := operationGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.service.operation", "get", inputDataValue, executionContext)
	var emptyOutput routing.OperationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), operationGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(routing.OperationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
