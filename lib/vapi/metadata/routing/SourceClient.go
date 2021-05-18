// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Source
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

// Operations to manage the metadata sources for routing information
type SourceClient interface {

	// Create a new metadata source.
	//
	// @param sourceIdParam  metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	// @param specParam  create specification.
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
	// @throws NotFound  If the metadata source identifier is not found.
	Delete(sourceIdParam string) error

	// Get the details about a metadata source.
	//
	// @param sourceIdParam  Metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	// @return Metadata source info.
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
	// @throws NotFound  If the metadata source identifier is not found.
	Reload(sourceIdParam *string) error

	// Returns the fingerprint of all the sources or of a particular source.
	//
	// @param sourceIdParam  Metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.source``.
	//  If unspecified, fingerprint of all the sources is returned
	// @return fingerprint of all the sources or of a particular source.
	// @throws NotFound  If the metadata source identifier is not found.
	Fingerprint(sourceIdParam *string) (string, error)
}

type sourceClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSourceClient(connector client.Connector) *sourceClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.routing.source")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":      core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":      core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":         core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":        core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reload":      core.NewMethodIdentifier(interfaceIdentifier, "reload"),
		"fingerprint": core.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := sourceClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sourceClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sourceClient) Create(sourceIdParam string, specParam SourceCreateSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceCreateInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Delete(sourceIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceDeleteInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Get(sourceIdParam string) (SourceInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceGetInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput SourceInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "get", inputDataValue, executionContext)
	var emptyOutput SourceInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sourceGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(SourceInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sourceClient) List() ([]string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sourceListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sourceClient) Reload(sourceIdParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceReloadInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceReloadRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "reload", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Fingerprint(sourceIdParam *string) (string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceFingerprintInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceFingerprintRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.routing.source", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sourceFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
