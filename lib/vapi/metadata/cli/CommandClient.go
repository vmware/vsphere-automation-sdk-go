// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Command
// Used by client-side stubs.

package cli

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// The ``Command`` interface provides methods to get information about command line interface (CLI) commands.
type CommandClient interface {

	// Returns the identifiers of all commands, or commands in a specific namespace.
	//
	// @param pathParam The dot-separated path of the namespace for which command identifiers should be returned.
	// If null identifiers of all commands registered with the infrastructure will be returned.
	// @return Identifiers of the requested commands.
	// @throws NotFound if a namespace corresponding to ``path`` doesn't exist.
	List(pathParam *string) ([]CommandIdentity, error)

	// Retrieves information about a command including information about how to execute that command.
	//
	// @param identityParam Identifier of the command for which to retreive information.
	// @return Information about the command including information about how to execute that command.
	// @throws NotFound if a command corresponding to ``identity`` doesn't exist.
	Get(identityParam CommandIdentity) (CommandInfo, error)

	// Returns the aggregate fingerprint of all the command metadata from all the metadata sources.
	//
	//  The fingerprint provides clients an efficient way to check if the metadata for commands has been modified on the server.
	// @return Fingerprint of all the command metadata present on the server.
	Fingerprint() (string, error)
}

type commandClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCommandClient(connector client.Connector) *commandClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.cli.command")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list":        core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":         core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"fingerprint": core.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := commandClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *commandClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *commandClient) List(pathParam *string) ([]CommandIdentity, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(commandListInputType(), typeConverter)
	sv.AddStructField("Path", pathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []CommandIdentity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := commandListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.cli.command", "list", inputDataValue, executionContext)
	var emptyOutput []CommandIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), commandListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]CommandIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *commandClient) Get(identityParam CommandIdentity) (CommandInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(commandGetInputType(), typeConverter)
	sv.AddStructField("Identity", identityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput CommandInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := commandGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.cli.command", "get", inputDataValue, executionContext)
	var emptyOutput CommandInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), commandGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(CommandInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *commandClient) Fingerprint() (string, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(commandFingerprintInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := commandFingerprintRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.cli.command", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), commandFingerprintOutputType())
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
