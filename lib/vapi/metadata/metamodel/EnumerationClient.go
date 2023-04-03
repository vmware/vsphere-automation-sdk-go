// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Enumeration
// Used by client-side stubs.

package metamodel

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Enumeration`` interface provides methods to retrieve metamodel information about an enumeration element in the interface definition language.
//
//  The ``Enumeration`` has a list of enumeration value elements.
type EnumerationClient interface {

	// Returns the identifiers for the enumeration elements that are contained in all the package elements, service elements and structure elements.
	// @return The list of identifiers for the enumeration elements.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.enumeration``.
	List() ([]string, error)

	// Retrieves information about the enumeration element corresponding to ``enumeration_id``.
	//
	//  The EnumerationInfo contains the metamodel information about the enumeration value element contained in the enumeration element.
	//
	// @param enumerationIdParam Identifier of the enumeration element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.enumeration``.
	// @return The EnumerationInfo instance that corresponds to ``enumeration_id``
	//
	// @throws NotFound if the enumeration element associated with ``enumeration_id`` is not contained in any of the package elements, service elements and structure elements.
	Get(enumerationIdParam string) (EnumerationInfo, error)
}

type enumerationClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEnumerationClient(connector vapiProtocolClient_.Connector) *enumerationClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.enumeration")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := enumerationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *enumerationClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *enumerationClient) List() ([]string, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := enumerationListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(enumerationListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.enumeration", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EnumerationListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *enumerationClient) Get(enumerationIdParam string) (EnumerationInfo, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := enumerationGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(enumerationGetInputType(), typeConverter)
	sv.AddStructField("EnumerationId", enumerationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput EnumerationInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.enumeration", "get", inputDataValue, executionContext)
	var emptyOutput EnumerationInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EnumerationGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(EnumerationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
