// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Reservation
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

type ReservationClient interface {

	// Retreive all reservations for all SDDCs in this org
	//
	// @param orgParam Organization identifier (required)
	// @param sddcStateParam SDDCs and/or states to query (optional)
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Call
	// @throws Unauthorized  Forbidden
	Post(orgParam string, sddcStateParam *model.SddcStateRequest) (map[string][]model.ReservationWindow, error)
}

type reservationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewReservationClient(connector client.Connector) *reservationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.orgs.tbrs.reservation")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"post": core.NewMethodIdentifier(interfaceIdentifier, "post"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := reservationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *reservationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *reservationClient) Post(orgParam string, sddcStateParam *model.SddcStateRequest) (map[string][]model.ReservationWindow, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(reservationPostInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcState", sddcStateParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput map[string][]model.ReservationWindow
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := reservationPostRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.tbrs.reservation", "post", inputDataValue, executionContext)
	var emptyOutput map[string][]model.ReservationWindow
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), reservationPostOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string][]model.ReservationWindow), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
