// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Settings
// Used by client-side stubs.

package gateway_firewall

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type SettingsClient interface {

	// Get the current gateway firewall configurations.
	// @return com.vmware.nsx_policy.model.GfwFirewallConfiguration
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsx_policyModel.GfwFirewallConfiguration, error)

	// Update gateway firewall related configurations.
	//
	// @param gfwFirewallConfigurationParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(gfwFirewallConfigurationParam nsx_policyModel.GfwFirewallConfiguration) error

	// Update gateway firewall related configurations.
	//
	// @param gfwFirewallConfigurationParam (required)
	// @return com.vmware.nsx_policy.model.GfwFirewallConfiguration
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(gfwFirewallConfigurationParam nsx_policyModel.GfwFirewallConfiguration) (nsx_policyModel.GfwFirewallConfiguration, error)
}

type settingsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewSettingsClient(connector vapiProtocolClient_.Connector) *settingsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.security.gateway_firewall.settings")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := settingsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *settingsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *settingsClient) Get() (nsx_policyModel.GfwFirewallConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := settingsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(settingsGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GfwFirewallConfiguration
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.gateway_firewall.settings", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GfwFirewallConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SettingsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GfwFirewallConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *settingsClient) Patch(gfwFirewallConfigurationParam nsx_policyModel.GfwFirewallConfiguration) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := settingsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(settingsPatchInputType(), typeConverter)
	sv.AddStructField("GfwFirewallConfiguration", gfwFirewallConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.gateway_firewall.settings", "patch", inputDataValue, executionContext)
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

func (sIface *settingsClient) Update(gfwFirewallConfigurationParam nsx_policyModel.GfwFirewallConfiguration) (nsx_policyModel.GfwFirewallConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := settingsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(settingsUpdateInputType(), typeConverter)
	sv.AddStructField("GfwFirewallConfiguration", gfwFirewallConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.GfwFirewallConfiguration
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.security.gateway_firewall.settings", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.GfwFirewallConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), SettingsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.GfwFirewallConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
