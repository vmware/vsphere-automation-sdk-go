// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SupportWindow
// Used by client-side stubs.

package tbrs

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type SupportWindowClient interface {

	// Get all available support windows
	//
	// @param orgParam Organization identifier (required)
	// @param minimumSeatsAvailableParam minimum seats available (used as a filter) (optional)
	// @param createdByParam user name which was used to create the support window (used as a filter) (optional)
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid request
	// @throws Unauthorized  Forbidden
	// @throws NotFound  No support windows are available
	Get(orgParam string, minimumSeatsAvailableParam *int64, createdByParam *string) ([]model.SupportWindow, error)

	// Move Sddc to new support window
	//
	// @param orgParam Organization identifier (required)
	// @param idParam Target Support Window ID (required)
	// @param sddcIdParam SDDC to move (required)
	// @return com.vmware.vmc.model.SupportWindowId
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Invalid request
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Feature does not exist
	Put(orgParam string, idParam string, sddcIdParam model.SddcId) (model.SupportWindowId, error)
}

type supportWindowClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSupportWindowClient(connector client.Connector) *supportWindowClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.tbrs.support_window")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"put": core.NewMethodIdentifier(interfaceIdentifier, "put"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := supportWindowClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *supportWindowClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *supportWindowClient) Get(orgParam string, minimumSeatsAvailableParam *int64, createdByParam *string) ([]model.SupportWindow, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(supportWindowGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("MinimumSeatsAvailable", minimumSeatsAvailableParam)
	sv.AddStructField("CreatedBy", createdByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.SupportWindow
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := supportWindowGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tbrs.support_window", "get", inputDataValue, executionContext)
	var emptyOutput []model.SupportWindow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), supportWindowGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.SupportWindow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *supportWindowClient) Put(orgParam string, idParam string, sddcIdParam model.SddcId) (model.SupportWindowId, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(supportWindowPutInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("SddcId", sddcIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SupportWindowId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := supportWindowPutRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tbrs.support_window", "put", inputDataValue, executionContext)
	var emptyOutput model.SupportWindowId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), supportWindowPutOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SupportWindowId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
