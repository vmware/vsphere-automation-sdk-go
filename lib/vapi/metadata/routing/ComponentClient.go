// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Component
// Used by client-side stubs.

package routing

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// Operations to retrieve information about the routing information in a vAPI component. A Component is said to contain routing information if any of its packages, services or methods contain routing information
type ComponentClient interface {

	// List all the vAPI components that contain operations which have routing information.
	// @return list of fully qualified component names
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.component``.
	List() ([]string, error)

	// Get the routing information for a vAPI component
	//
	// @param componentIdParam  fully qualified component name
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
	// @return routing information for the vAPI component
	// @throws NotFound  If the component name does not exist
	Get(componentIdParam string) (ComponentData, error)

	// Checksum of all routing metadata for a vAPI component on the server
	//
	// @param componentIdParam  fully qualified component name
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.component``.
	// @return checksum of routing metadata for a vAPI component
	// @throws NotFound  If the component name does not exist
	Fingerprint(componentIdParam string) (string, error)
}

type componentClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewComponentClient(connector client.Connector) *componentClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.routing.component")
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
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.component", "list", inputDataValue, executionContext)
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
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.component", "get", inputDataValue, executionContext)
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
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.component", "fingerprint", inputDataValue, executionContext)
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
