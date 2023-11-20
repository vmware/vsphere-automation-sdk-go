// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: StandaloneHostSwitchSetting
// Used by client-side stubs.

package idfw

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StandaloneHostSwitchSettingClient interface {

	// Fetches IDFW standalone hosts switch setting to check whether standalone hosts is enabled or disabled
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/settings/firewall/idfw/standalone-host-switch-setting
	//
	// Deprecated: This API element is deprecated.
	// @return com.vmware.nsx.model.IdfwStandaloneHostsSwitchSetting
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (nsxModel.IdfwStandaloneHostsSwitchSetting, error)

	// Update Identity Firewall standalone hosts switch setting (true=enabled / false=disabled).
	//
	//  Use the following Policy API -
	//  PUT/PATCH /policy/api/v1/infra/settings/firewall/idfw/standalone-host-switch-setting
	//
	// Deprecated: This API element is deprecated.
	//
	// @param idfwStandaloneHostsSwitchSettingParam (required)
	// @return com.vmware.nsx.model.IdfwStandaloneHostsSwitchSetting
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(idfwStandaloneHostsSwitchSettingParam nsxModel.IdfwStandaloneHostsSwitchSetting) (nsxModel.IdfwStandaloneHostsSwitchSetting, error)
}

type standaloneHostSwitchSettingClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStandaloneHostSwitchSettingClient(connector vapiProtocolClient_.Connector) *standaloneHostSwitchSettingClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.idfw.standalone_host_switch_setting")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := standaloneHostSwitchSettingClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *standaloneHostSwitchSettingClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *standaloneHostSwitchSettingClient) Get() (nsxModel.IdfwStandaloneHostsSwitchSetting, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := standaloneHostSwitchSettingGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(standaloneHostSwitchSettingGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IdfwStandaloneHostsSwitchSetting
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.idfw.standalone_host_switch_setting", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.IdfwStandaloneHostsSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StandaloneHostSwitchSettingGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IdfwStandaloneHostsSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *standaloneHostSwitchSettingClient) Update(idfwStandaloneHostsSwitchSettingParam nsxModel.IdfwStandaloneHostsSwitchSetting) (nsxModel.IdfwStandaloneHostsSwitchSetting, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := standaloneHostSwitchSettingUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(standaloneHostSwitchSettingUpdateInputType(), typeConverter)
	sv.AddStructField("IdfwStandaloneHostsSwitchSetting", idfwStandaloneHostsSwitchSettingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.IdfwStandaloneHostsSwitchSetting
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.idfw.standalone_host_switch_setting", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.IdfwStandaloneHostsSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StandaloneHostSwitchSettingUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.IdfwStandaloneHostsSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
