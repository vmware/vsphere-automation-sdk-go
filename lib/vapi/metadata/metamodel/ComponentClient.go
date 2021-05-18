// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Component
// Used by client-side stubs.

package metamodel

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// The ``Component`` interface providers methods to retrieve metamodel information of a component element.
//
//  A component defines a set of functionality that is deployed together and versioned together. For example, all the interfaces that belong to VMware Content Library are part of a single component. A component element describes a component. A component element contains one or more package elements.
//
//  The methods for package elements are provided by interface Package.
type ComponentClient interface {

	// Returns the identifiers for the component elements that are registered with the infrastructure.
	// @return The list of identifiers for the component elements that are registered with the infrastructure.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.component``.
	List() ([]string, error)

	// Retrieves metamodel information about the component element corresponding to ``component_id``.
	//
	//  The ComponentData contains the metamodel information about the component and it's fingerprint. It contains information about all the package elements that are contained in this component element.
	//
	// @param componentIdParam Identifier of the component element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
	// @return The ComponentData instance that corresponds to ``component_id``.
	// @throws NotFound if the component element associated with ``component_id`` is not registered with the infrastructure.
	Get(componentIdParam string) (ComponentData, error)

	// Retrieves the fingerprint computed from the metamodel metadata of the component element corresponding to ``component_id``.
	//
	//  The fingerprint provides clients an efficient way to check if the metadata for a particular component element has been modified on the server. The client can do this by comparing the result of this operation with the fingerprint returned in the result of Component#get.
	//
	// @param componentIdParam Identifier of the component element.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
	// @return The fingerprint computed from the metamodel metadata of the component element.
	// @throws NotFound if the component element associated with ``component_id`` is not registered with the infrastructure.
	Fingerprint(componentIdParam string) (string, error)
}

type componentClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewComponentClient(connector client.Connector) *componentClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.component")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list":        core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":         core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"fingerprint": core.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := componentClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *componentClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *componentClient) List() ([]string, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(componentListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := componentListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.component", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), componentListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *componentClient) Get(componentIdParam string) (ComponentData, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(componentGetInputType(), typeConverter)
	sv.AddStructField("ComponentId", componentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput ComponentData
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := componentGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.component", "get", inputDataValue, executionContext)
	var emptyOutput ComponentData
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), componentGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(ComponentData), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *componentClient) Fingerprint(componentIdParam string) (string, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(componentFingerprintInputType(), typeConverter)
	sv.AddStructField("ComponentId", componentIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := componentFingerprintRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.component", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), componentFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
