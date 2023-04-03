// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Operation
// Used by client-side stubs.

package introspection

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The Operation service provides operations to retrieve information about the operations present in a vAPI service.
type OperationClient interface {

	// Returns the set of operation identifiers for a given vAPI service.
	//
	// @param serviceIdParam service identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
	// @return set of operation identifiers for a given vAPI service.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
	//
	// @throws NotFound If the service identifier does not exist.
	List(serviceIdParam string) (map[string]bool, error)

	// Returns the OperationInfo for a given vAPI operation.
	//
	// @param serviceIdParam service identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
	// @param operationIdParam operation identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
	// @return OperationInfo for a given vAPI operation.
	//
	// @throws NotFound If the operation identifier does not exist.
	// @throws NotFound If the service identifier does not exist.
	Get(serviceIdParam string, operationIdParam string) (OperationInfo, error)
}

type operationClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOperationClient(connector vapiProtocolClient_.Connector) *operationClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.std.introspection.operation")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	oIface := operationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &oIface
}

func (oIface *operationClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := oIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (oIface *operationClient) List(serviceIdParam string) (map[string]bool, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := operationListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(operationListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput map[string]bool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vapi.std.introspection.operation", "list", inputDataValue, executionContext)
	var emptyOutput map[string]bool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string]bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *operationClient) Get(serviceIdParam string, operationIdParam string) (OperationInfo, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := operationGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(operationGetInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("OperationId", operationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput OperationInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vapi.std.introspection.operation", "get", inputDataValue, executionContext)
	var emptyOutput OperationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(OperationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
