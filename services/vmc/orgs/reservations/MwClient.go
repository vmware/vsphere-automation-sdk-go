// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Mw
// Used by client-side stubs.

package reservations

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = core.SupportedByRuntimeVersion1

type MwClient interface {

	// get the maintenance window for this SDDC
	//
	// @param orgParam Organization identifier (required)
	// @param reservationParam Reservation Identifier (required)
	// @return com.vmware.vmc.model.MaintenanceWindowGet
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Access not allowed to the operation for the current user
	Get(orgParam string, reservationParam string) (model.MaintenanceWindowGet, error)

	// update the maintenance window for this SDDC
	//
	// @param orgParam Organization identifier (required)
	// @param reservationParam Reservation Identifier (required)
	// @param windowParam Maintenance Window (required)
	// @return com.vmware.vmc.model.MaintenanceWindow
	// @throws Unauthenticated  Unauthorized
	// @throws ConcurrentChange  Conflict with exiting reservation
	// @throws InvalidRequest  The reservation is not in a state that's valid for updates
	// @throws Unauthorized  Access not allowed to the operation for the current user
	Put(orgParam string, reservationParam string, windowParam model.MaintenanceWindow) (model.MaintenanceWindow, error)
}

type mwClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewMwClient(connector client.Connector) *mwClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.reservations.mw")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"put": core.NewMethodIdentifier(interfaceIdentifier, "put"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := mwClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *mwClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *mwClient) Get(orgParam string, reservationParam string) (model.MaintenanceWindowGet, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(mwGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Reservation", reservationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MaintenanceWindowGet
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := mwGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.reservations.mw", "get", inputDataValue, executionContext)
	var emptyOutput model.MaintenanceWindowGet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), mwGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MaintenanceWindowGet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *mwClient) Put(orgParam string, reservationParam string, windowParam model.MaintenanceWindow) (model.MaintenanceWindow, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(mwPutInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Reservation", reservationParam)
	sv.AddStructField("Window", windowParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MaintenanceWindow
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := mwPutRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.reservations.mw", "put", inputDataValue, executionContext)
	var emptyOutput model.MaintenanceWindow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), mwPutOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MaintenanceWindow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
