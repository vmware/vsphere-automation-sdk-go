// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Source
// Used by client-side stubs.

package routing

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// Operations to manage the metadata sources for routing information
type SourceClient interface {

	// Create a new metadata source.
	//
	// @param sourceIdParam  metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	// @param specParam  create specification.
	//
	// @throws AlreadyExists  If the metadata source identifier is already present.
	// @throws InvalidArgument  If type of the source specified in \\\\@{link CreateSpec#type} is invalid.
	// @throws InvalidArgument  If the file specified in \\\\@{link CreateSpec#filepath} is not a valid json file.
	// @throws InvalidArgument  If the URI specified in \\\\@{link CreateSpec#address} is unreachable or not a vAPI compatible server.
	// @throws NotFound  If the file specified in \\\\@{link CreateSpec#filepath} does not exist.
	Create(sourceIdParam string, specParam SourceCreateSpec) error

	// Delete a metadata source.
	//
	// @param sourceIdParam  Metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	//
	// @throws NotFound  If the metadata source identifier is not found.
	Delete(sourceIdParam string) error

	// Get the details about a metadata source.
	//
	// @param sourceIdParam  Metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	// @return Metadata source info.
	//
	// @throws NotFound  If the metadata source identifier is not found.
	Get(sourceIdParam string) (SourceInfo, error)

	// List all the metadata sources.
	// @return List of all metadata sources.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.metadata.source``.
	List() ([]string, error)

	// Reload metadata from all the sources or of a particular source.
	//
	// @param sourceIdParam  Metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	//  If unspecified, all the sources are reloaded
	//
	// @throws NotFound  If the metadata source identifier is not found.
	Reload(sourceIdParam *string) error

	// Returns the fingerprint of all the sources or of a particular source.
	//
	// @param sourceIdParam  Metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	//  If unspecified, fingerprint of all the sources is returned
	// @return fingerprint of all the sources or of a particular source.
	//
	// @throws NotFound  If the metadata source identifier is not found.
	Fingerprint(sourceIdParam *string) (string, error)
}

type sourceClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSourceClient(connector vapiProtocolClient_.Connector) *sourceClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.routing.source")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":         vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":        vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reload":      vapiCore_.NewMethodIdentifier(interfaceIdentifier, "reload"),
		"fingerprint": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := sourceClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sourceClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sourceClient) Create(sourceIdParam string, specParam SourceCreateSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sourceCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sourceCreateInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Delete(sourceIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sourceDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sourceDeleteInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Get(sourceIdParam string) (SourceInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sourceGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sourceGetInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput SourceInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "get", inputDataValue, executionContext)
	var emptyOutput SourceInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SourceGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(SourceInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sourceClient) List() ([]string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sourceListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sourceListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SourceListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sourceClient) Reload(sourceIdParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sourceReloadRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sourceReloadInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "reload", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Fingerprint(sourceIdParam *string) (string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := sourceFingerprintRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(sourceFingerprintInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SourceFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
