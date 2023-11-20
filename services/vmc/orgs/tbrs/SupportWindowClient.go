// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SupportWindow
// Used by client-side stubs.

package tbrs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SupportWindowClient interface {

	// Get all available support windows
	//
	// @param orgParam Organization identifier (required)
	// @param minimumSeatsAvailableParam minimum seats available (used as a filter) (optional)
	// @param createdByParam user name which was used to create the support window (used as a filter) (optional)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid request
	// @throws Unauthorized  Forbidden
	// @throws NotFound  No support windows are available
	Get(orgParam string, minimumSeatsAvailableParam *int64, createdByParam *string) ([]vmcModel.SupportWindow, error)

	// Move Sddc to new support window
	//
	// @param orgParam Organization identifier (required)
	// @param idParam Target Support Window ID (required)
	// @param sddcIdParam SDDC to move (required)
	// @return com.vmware.vmc.model.SupportWindowId
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid request
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Feature does not exist
	Put(orgParam string, idParam string, sddcIdParam vmcModel.SddcId) (vmcModel.SupportWindowId, error)
}

type supportWindowClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSupportWindowClient(connector vapiProtocolClient_.Connector) *supportWindowClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.tbrs.support_window")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"put": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "put"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := supportWindowClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *supportWindowClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *supportWindowClient) Get(orgParam string, minimumSeatsAvailableParam *int64, createdByParam *string) ([]vmcModel.SupportWindow, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := supportWindowGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(supportWindowGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("MinimumSeatsAvailable", minimumSeatsAvailableParam)
	sv.AddStructField("CreatedBy", createdByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.SupportWindow
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tbrs.support_window", "get", inputDataValue, executionContext)
	var emptyOutput []vmcModel.SupportWindow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SupportWindowGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.SupportWindow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *supportWindowClient) Put(orgParam string, idParam string, sddcIdParam vmcModel.SddcId) (vmcModel.SupportWindowId, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := supportWindowPutRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(supportWindowPutInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("SddcId", sddcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.SupportWindowId
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tbrs.support_window", "put", inputDataValue, executionContext)
	var emptyOutput vmcModel.SupportWindowId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SupportWindowPutOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.SupportWindowId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
