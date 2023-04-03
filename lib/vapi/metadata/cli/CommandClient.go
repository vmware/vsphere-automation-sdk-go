// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Command
// Used by client-side stubs.

package cli

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The ``Command`` interface provides methods to get information about command line interface (CLI) commands.
type CommandClient interface {

	// Returns the identifiers of all commands, or commands in a specific namespace.
	//
	// @param pathParam The dot-separated path of the namespace for which command identifiers should be returned.
	// If null identifiers of all commands registered with the infrastructure will be returned.
	// @return Identifiers of the requested commands.
	//
	// @throws NotFound if a namespace corresponding to ``path`` doesn't exist.
	List(pathParam *string) ([]CommandIdentity, error)

	// Retrieves information about a command including information about how to execute that command.
	//
	// @param identityParam Identifier of the command for which to retreive information.
	// @return Information about the command including information about how to execute that command.
	//
	// @throws NotFound if a command corresponding to ``identity`` doesn't exist.
	Get(identityParam CommandIdentity) (CommandInfo, error)

	// Returns the aggregate fingerprint of all the command metadata from all the metadata sources.
	//
	//  The fingerprint provides clients an efficient way to check if the metadata for commands has been modified on the server.
	// @return Fingerprint of all the command metadata present on the server.
	Fingerprint() (string, error)
}

type commandClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewCommandClient(connector vapiProtocolClient_.Connector) *commandClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.cli.command")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list":        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"get":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"fingerprint": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	cIface := commandClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *commandClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *commandClient) List(pathParam *string) ([]CommandIdentity, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := commandListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(commandListInputType(), typeConverter)
	sv.AddStructField("Path", pathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []CommandIdentity
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.cli.command", "list", inputDataValue, executionContext)
	var emptyOutput []CommandIdentity
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommandListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]CommandIdentity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *commandClient) Get(identityParam CommandIdentity) (CommandInfo, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := commandGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(commandGetInputType(), typeConverter)
	sv.AddStructField("Identity", identityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput CommandInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.cli.command", "get", inputDataValue, executionContext)
	var emptyOutput CommandInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommandGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(CommandInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *commandClient) Fingerprint() (string, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	operationRestMetaData := commandFingerprintRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(commandFingerprintInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.cli.command", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), CommandFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
