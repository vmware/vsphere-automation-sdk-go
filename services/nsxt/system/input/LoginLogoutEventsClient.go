// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: LoginLogoutEvents
// Used by client-side stubs.

package input

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type LoginLogoutEventsClient interface {

	// API to receive User Login and Logout events for IDFW
	//
	// @param loginLogoutEventCollectionParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(loginLogoutEventCollectionParam model.LoginLogoutEventCollection) error
}

type loginLogoutEventsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewLoginLogoutEventsClient(connector client.Connector) *loginLogoutEventsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.system.input.login_logout_events")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	lIface := loginLogoutEventsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &lIface
}

func (lIface *loginLogoutEventsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := lIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (lIface *loginLogoutEventsClient) Create(loginLogoutEventCollectionParam model.LoginLogoutEventCollection) error {
	typeConverter := lIface.connector.TypeConverter()
	executionContext := lIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(loginLogoutEventsCreateInputType(), typeConverter)
	sv.AddStructField("LoginLogoutEventCollection", loginLogoutEventCollectionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := loginLogoutEventsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := lIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.system.input.login_logout_events", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
