// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Hidden
// Used by client-side stubs.

package service

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// The ``Hidden`` interface provides methods to retrieve the list of interfaces that are hidden and should not be exposed in various presentation layers of API infrastructure.
type HiddenClient interface {

	// Returns the interface identifiers that should be hidden.
	// @return The list of interface identifiers that should be hidden.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.service``.
	List() ([]string, error)
}

type hiddenClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewHiddenClient(connector client.Connector) *hiddenClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.service.hidden")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	hIface := hiddenClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *hiddenClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *hiddenClient) List() ([]string, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(hiddenListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := hiddenListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.metamodel.service.hidden", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), hiddenListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
