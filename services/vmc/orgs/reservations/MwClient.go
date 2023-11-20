// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Mw
// Used by client-side stubs.

package reservations

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type MwClient interface {

	// get the maintenance window for this SDDC
	//
	// Deprecated: This API element is deprecated.
	//
	// @param orgParam Organization identifier (required)
	// @param reservationParam Reservation Identifier (required)
	// @return com.vmware.vmc.model.MaintenanceWindowGet
	//
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Access not allowed to the operation for the current user
	Get(orgParam string, reservationParam string) (vmcModel.MaintenanceWindowGet, error)

	// update the maintenance window for this SDDC
	//
	// Deprecated: This API element is deprecated.
	//
	// @param orgParam Organization identifier (required)
	// @param reservationParam Reservation Identifier (required)
	// @param windowParam Maintenance Window (required)
	// @return com.vmware.vmc.model.MaintenanceWindow
	//
	// @throws Unauthenticated  Unauthorized
	// @throws ConcurrentChange  Conflict with exiting reservation
	// @throws InvalidRequest  The reservation is not in a state that's valid for updates
	// @throws Unauthorized  Access not allowed to the operation for the current user
	Put(orgParam string, reservationParam string, windowParam vmcModel.MaintenanceWindow) (vmcModel.MaintenanceWindow, error)
}

type mwClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewMwClient(connector vapiProtocolClient_.Connector) *mwClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.reservations.mw")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"put": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "put"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	mIface := mwClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *mwClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *mwClient) Get(orgParam string, reservationParam string) (vmcModel.MaintenanceWindowGet, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := mwGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(mwGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Reservation", reservationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.MaintenanceWindowGet
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.reservations.mw", "get", inputDataValue, executionContext)
	var emptyOutput vmcModel.MaintenanceWindowGet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), MwGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.MaintenanceWindowGet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *mwClient) Put(orgParam string, reservationParam string, windowParam vmcModel.MaintenanceWindow) (vmcModel.MaintenanceWindow, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	operationRestMetaData := mwPutRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(mwPutInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Reservation", reservationParam)
	sv.AddStructField("Window", windowParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmcModel.MaintenanceWindow
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.reservations.mw", "put", inputDataValue, executionContext)
	var emptyOutput vmcModel.MaintenanceWindow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), MwPutOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmcModel.MaintenanceWindow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
