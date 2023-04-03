// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Operation
// Used by client-side stubs.

package service

import (
	vapiMetadataPrivilege_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/privilege"
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Operation`` interface provides methods to retrieve privilege information of an operation element.
//
//  An operation element is said to contain privilege information if there are any privileges assigned to the operation element or if one of the parameter elements contained in it has privileges assigned in privilege definition file.
type OperationClient interface {

	// Returns the identifiers for the operation elements contained in the service element corresponding to ``service_id`` that have privilege information.
	//
	// @param serviceIdParam Identifier of the service element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
	// @return List of identifiers for the operation elements contained in the service element that have privilege information.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.operation``.
	//
	// @throws NotFound if the service element associated with ``service_id`` does not have any operation elements that have privilege information.
	List(serviceIdParam string) ([]string, error)

	// Retrieves the privilege information about an operation element corresponding to ``operation_id`` contained in the service element corresponding to ``service_id``.
	//
	// @param serviceIdParam Identifier of the service element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.service``.
	// @param operationIdParam Identifier of the operation element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.operation``.
	// @return The vapiMetadataPrivilege_.OperationInfo instance that corresponds to ``operation_id``.
	//
	// @throws NotFound if the service element associated with ``service_id`` does not exist.
	// @throws NotFound if the operation element associated with ``operation_id`` does not exist.
	// @throws NotFound if the operation element associated with ``operation_id`` does not have any privilege information.
	Get(serviceIdParam string, operationIdParam string) (vapiMetadataPrivilege_.OperationInfo, error)
}

type operationClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewOperationClient(connector vapiProtocolClient_.Connector) *operationClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.privilege.service.operation")
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

func (oIface *operationClient) List(serviceIdParam string) ([]string, error) {
	typeConverter := oIface.connector.TypeConverter()
	executionContext := oIface.connector.NewExecutionContext()
	operationRestMetaData := operationListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(operationListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.service.operation", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *operationClient) Get(serviceIdParam string, operationIdParam string) (vapiMetadataPrivilege_.OperationInfo, error) {
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
		var emptyOutput vapiMetadataPrivilege_.OperationInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := oIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.service.operation", "get", inputDataValue, executionContext)
	var emptyOutput vapiMetadataPrivilege_.OperationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), OperationGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vapiMetadataPrivilege_.OperationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
