// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Source
// Used by client-side stubs.

package privilege

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

// The “Source“ interface provides methods to manage the sources of privilege metadata information.
//
//	The interface definition language infrastructure provides tools to generate various kinds of metadata in JSON format from the interface definition files and additional properties files. One of the generated files contains privilege information. The generated file can be registered as a source of metadata.
//
//	The privilege file contains all the data present in the interface definition files. Each privilege file contains data about one component element. When a privilege file is added as a source, each source contributes only one component element's metadata.
//
//	Privilege metadata can also be discovered from a remote server that supports the privilege metadata interfaces (see com.vmware.vapi.metadata.privilege). Since multiple components can be registered with a single metadata server, when a remote server is registered as a source, that source can contribute more than one component.
type SourceClient interface {

	// Creates a new metadata source. Once the server validates the registration information of the metadata source, the privilege metadata is retrieved from the source. This populates elements in all the interfaces defined in com.vmware.vapi.metadata.privilege package.
	//
	// @param sourceIdParam metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// @param specParam create specification.
	//
	// @throws AlreadyExists if the metadata source identifier is already registered with the infrastructure.
	// @throws InvalidArgument if the type of the source specified in vapiMetadata_.SourceCreateSpec#type is invalid.
	// @throws InvalidArgument if the file specified in vapiMetadata_.SourceCreateSpec#filepath is not a valid JSON file or if the format of the privilege metadata in the JSON file is invalid.
	// @throws InvalidArgument if the URI specified in vapiMetadata_.SourceCreateSpec#address is unreachable or if there is a transport protocol or message protocol mismatch between the client and the server or if the remote server do not have interfaces present in com.vmware.vapi.metadata.privilege package.
	// @throws NotFound if the file specified in vapiMetadata_.SourceCreateSpec#filepath does not exist.
	Create(sourceIdParam string, specParam SourceCreateSpec) error

	// Deletes an existing privilege metadata source from the infrastructure.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	//
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Delete(sourceIdParam string) error

	// Retrieves information about the metadata source corresponding to ``source_id``.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// @return The SourceInfo instance that corresponds to ``source_id``
	//
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Get(sourceIdParam string) (SourceInfo, error)

	// Returns the identifiers of the metadata sources currently registered with the infrastructure.
	// @return The list of identifiers for metadata sources currently registered.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	List() ([]string, error)

	// Reloads the privilege metadata from all the metadata sources or of a particular metadata source if ``source_id`` is specified.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// If unspecified, all the metadata sources are reloaded.
	//
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Reload(sourceIdParam *string) error

	// Returns the aggregate fingerprint of metadata from all the metadata sources or from a particular metadata source if ``source_id`` is specified.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// If unspecified, the fingerprint of all the metadata sources is returned.
	// @return Aggregate fingerprint of all the metadata sources or of a particular metadata source.
	//
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Fingerprint(sourceIdParam *string) (string, error)
}

type sourceClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSourceClient(connector vapiProtocolClient_.Connector) *sourceClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vapi.metadata.privilege.source")
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "create", inputDataValue, executionContext)
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "delete", inputDataValue, executionContext)
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "get", inputDataValue, executionContext)
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "list", inputDataValue, executionContext)
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "reload", inputDataValue, executionContext)
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

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "fingerprint", inputDataValue, executionContext)
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
