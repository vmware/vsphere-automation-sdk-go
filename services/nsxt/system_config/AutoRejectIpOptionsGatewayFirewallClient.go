// Copyright (c) 2019-2025 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AutoRejectIpOptionsGatewayFirewall
// Used by client-side stubs.

package system_config

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type AutoRejectIpOptionsGatewayFirewallClient interface {

	// Read Configuration values for auto-reject-ip-options-gateway-firewall
	// @return com.vmware.nsx_policy.model.AutoRejectIPOptionsGatewayFirewallConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig, error)

	// Update auto reject ip options gateway firewall configuration value
	//
	// @param autoRejectIPOptionsGatewayFirewallConfigParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(autoRejectIPOptionsGatewayFirewallConfigParam nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig) error
}

type autoRejectIpOptionsGatewayFirewallClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewAutoRejectIpOptionsGatewayFirewallClient(connector vapiProtocolClient_.Connector) *autoRejectIpOptionsGatewayFirewallClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.system_config.auto_reject_ip_options_gateway_firewall")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	aIface := autoRejectIpOptionsGatewayFirewallClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *autoRejectIpOptionsGatewayFirewallClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *autoRejectIpOptionsGatewayFirewallClient) Get() (nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := autoRejectIpOptionsGatewayFirewallGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(autoRejectIpOptionsGatewayFirewallGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.system_config.auto_reject_ip_options_gateway_firewall", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), AutoRejectIpOptionsGatewayFirewallGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *autoRejectIpOptionsGatewayFirewallClient) Patch(autoRejectIPOptionsGatewayFirewallConfigParam nsx_policyModel.AutoRejectIPOptionsGatewayFirewallConfig) error {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	operationRestMetaData := autoRejectIpOptionsGatewayFirewallPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(autoRejectIpOptionsGatewayFirewallPatchInputType(), typeConverter)
	sv.AddStructField("AutoRejectIPOptionsGatewayFirewallConfig", autoRejectIPOptionsGatewayFirewallConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.system_config.auto_reject_ip_options_gateway_firewall", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
