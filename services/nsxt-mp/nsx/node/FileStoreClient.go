// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FileStore
// Used by client-side stubs.

package node

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type FileStoreClient interface {

	// Copy a remote file to the file store. If you use scp or sftp, you must provide the remote server's SSH fingerprint. See the *NSX-T Administration Guide* for information and instructions about finding the SSH fingerprint.
	//
	// @param fileNameParam Destination filename (required)
	// @param copyFromRemoteFilePropertiesParam (required)
	// @return com.vmware.nsx.model.FileProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Copyfromremotefile(fileNameParam string, copyFromRemoteFilePropertiesParam model.CopyFromRemoteFileProperties) (model.FileProperties, error)

	// Copy a file in the file store to a remote server. If you use scp or sftp, you must provide the remote server's SSH fingerprint. See the *NSX-T Administration Guide* for information and instructions about finding the SSH fingerprint.
	//
	// @param fileNameParam Destination filename (required)
	// @param copyToRemoteFilePropertiesParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Copytoremotefile(fileNameParam string, copyToRemoteFilePropertiesParam model.CopyToRemoteFileProperties) error

	// Create a directory on the remote remote server. Supports only SFTP. You must provide the remote server's SSH fingerprint. See the *NSX Administration Guide* for information and instructions about finding the SSH fingerprint.
	//
	// @param createRemoteDirectoryPropertiesParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Createremotedirectory(createRemoteDirectoryPropertiesParam model.CreateRemoteDirectoryProperties) error

	// Delete file
	//
	// @param fileNameParam Name of the file to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(fileNameParam string) error

	// Read file properties
	//
	// @param fileNameParam Name of the file to retrieve information about (required)
	// @return com.vmware.nsx.model.FileProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(fileNameParam string) (model.FileProperties, error)

	// List node files
	// @return com.vmware.nsx.model.FilePropertiesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List() (model.FilePropertiesListResult, error)
}

type fileStoreClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewFileStoreClient(connector client.Connector) *fileStoreClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.node.file_store")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"copyfromremotefile":    core.NewMethodIdentifier(interfaceIdentifier, "copyfromremotefile"),
		"copytoremotefile":      core.NewMethodIdentifier(interfaceIdentifier, "copytoremotefile"),
		"createremotedirectory": core.NewMethodIdentifier(interfaceIdentifier, "createremotedirectory"),
		"delete":                core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":                   core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                  core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	fIface := fileStoreClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *fileStoreClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *fileStoreClient) Copyfromremotefile(fileNameParam string, copyFromRemoteFilePropertiesParam model.CopyFromRemoteFileProperties) (model.FileProperties, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(fileStoreCopyfromremotefileInputType(), typeConverter)
	sv.AddStructField("FileName", fileNameParam)
	sv.AddStructField("CopyFromRemoteFileProperties", copyFromRemoteFilePropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FileProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileStoreCopyfromremotefileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.file_store", "copyfromremotefile", inputDataValue, executionContext)
	var emptyOutput model.FileProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), fileStoreCopyfromremotefileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FileProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (fIface *fileStoreClient) Copytoremotefile(fileNameParam string, copyToRemoteFilePropertiesParam model.CopyToRemoteFileProperties) error {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(fileStoreCopytoremotefileInputType(), typeConverter)
	sv.AddStructField("FileName", fileNameParam)
	sv.AddStructField("CopyToRemoteFileProperties", copyToRemoteFilePropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileStoreCopytoremotefileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.file_store", "copytoremotefile", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (fIface *fileStoreClient) Createremotedirectory(createRemoteDirectoryPropertiesParam model.CreateRemoteDirectoryProperties) error {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(fileStoreCreateremotedirectoryInputType(), typeConverter)
	sv.AddStructField("CreateRemoteDirectoryProperties", createRemoteDirectoryPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileStoreCreateremotedirectoryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.file_store", "createremotedirectory", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (fIface *fileStoreClient) Delete(fileNameParam string) error {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(fileStoreDeleteInputType(), typeConverter)
	sv.AddStructField("FileName", fileNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileStoreDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.file_store", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (fIface *fileStoreClient) Get(fileNameParam string) (model.FileProperties, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(fileStoreGetInputType(), typeConverter)
	sv.AddStructField("FileName", fileNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FileProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileStoreGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.file_store", "get", inputDataValue, executionContext)
	var emptyOutput model.FileProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), fileStoreGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FileProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (fIface *fileStoreClient) List() (model.FilePropertiesListResult, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(fileStoreListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FilePropertiesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileStoreListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.file_store", "list", inputDataValue, executionContext)
	var emptyOutput model.FilePropertiesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), fileStoreListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FilePropertiesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
