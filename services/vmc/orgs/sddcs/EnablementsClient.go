// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Enablements
// Used by client-side stubs.

package sddcs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type EnablementsClient interface {

	// Enable/disable the add-on or set the the default enablement status for future add-ons in the SDDC.
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	// @param enablementParam Add-on name, for example \"hcx, \"draas\", \"skynet\". If it is \"default\", the operation is for default settings of future add-ons. (required)
	// @param actionParam Enable or disable the add-on or change default settings for future add-ons in the SDDC. (required)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request / Validation error.
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Add-on not found
	EnableDisableAddon(orgParam string, sddcParam string, enablementParam string, actionParam string) error

	// Return a list of enablement status for current add-ons and the default settings for future add-ons in the SDDC.
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam Sddc identifier (required)
	//
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request / Validation error.
	// @throws Unauthorized  Forbidden
	List(orgParam string, sddcParam string) ([]vmcModel.EnablementInfo, error)
}

type enablementsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewEnablementsClient(connector vapiProtocolClient_.Connector) *enablementsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmc.orgs.sddcs.enablements")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"enable_disable_addon": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "enable_disable_addon"),
		"list":                 vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	eIface := enablementsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *enablementsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *enablementsClient) EnableDisableAddon(orgParam string, sddcParam string, enablementParam string, actionParam string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := enablementsEnableDisableAddonRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(enablementsEnableDisableAddonInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Enablement", enablementParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.enablements", "enable_disable_addon", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *enablementsClient) List(orgParam string, sddcParam string) ([]vmcModel.EnablementInfo, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	operationRestMetaData := enablementsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(enablementsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []vmcModel.EnablementInfo
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.vmc.orgs.sddcs.enablements", "list", inputDataValue, executionContext)
	var emptyOutput []vmcModel.EnablementInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), EnablementsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.([]vmcModel.EnablementInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
