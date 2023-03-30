// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Model
// Used by client-side stubs.

package resource

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Model`` interface provides methods to retrieve information about models.
//
//  A structure is used as a model if it is used for persisting data about an entity. Some of the fields in the model structure are also used for creating indexes for querying.
//
//  One or more services can operate on the same resource type. One or more services can provide the model structure for an entity of this resource type. Using ``Model`` interface you can retrieve the list of all the structure elements that are model structures for a given resource type.
type ModelClient interface {

	// Returns the set of identifiers for the structure elements that are models for the resource type corresponding to ``resource_id``.
	//
	//  The Structure interface provides methods to retrieve more details about the structure elements corresponding to the identifiers returned by this method.
	//
	// @param resourceIdParam Identifier of the resource type.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.resource``.
	// @return The set of identifiers for the models that are associated with the resource type in ``resource_id``.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.structure``.
	//
	// @throws NotFound if the resource type associated with ``resource_id`` does not exist.
	List(resourceIdParam string) (map[string]bool, error)
}

type modelClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewModelClient(connector vapiProtocolClient_.Connector) *modelClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.resource.model")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	mIface := modelClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *modelClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *modelClient) List(resourceIdParam string) (map[string]bool, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := modelListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(modelListInputType(), typeConverter)
	sv.AddStructField("ResourceId", resourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput map[string]bool
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.resource.model", "list", inputDataValue, executionContext)
	var emptyOutput map[string]bool
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ModelListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string]bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
