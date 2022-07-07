// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: MasterSwitchSetting
// Used by client-side stubs.

package idfw

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type MasterSwitchSettingClient interface {

	// Fetches IDFW master switch setting to check whether master switch is enabled or disabled
	// @return com.vmware.nsx.model.IdfwMasterSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.IdfwMasterSwitchSetting, error)

	// Update Identity Firewall master switch setting (true=enabled / false=disabled). Identity Firewall master switch setting enables or disables Identity Firewall feature across the system. It affects compute collections, hypervisor and virtual machines. This operation is expensive and also has big impact and implication on system perforamce.
	//
	// @param idfwMasterSwitchSettingParam (required)
	// @return com.vmware.nsx.model.IdfwMasterSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(idfwMasterSwitchSettingParam model.IdfwMasterSwitchSetting) (model.IdfwMasterSwitchSetting, error)
}

type masterSwitchSettingClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewMasterSwitchSettingClient(connector client.Connector) *masterSwitchSettingClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.idfw.master_switch_setting")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := masterSwitchSettingClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *masterSwitchSettingClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *masterSwitchSettingClient) Get() (model.IdfwMasterSwitchSetting, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(masterSwitchSettingGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwMasterSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := masterSwitchSettingGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx.idfw.master_switch_setting", "get", inputDataValue, executionContext)
	var emptyOutput model.IdfwMasterSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), masterSwitchSettingGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwMasterSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *masterSwitchSettingClient) Update(idfwMasterSwitchSettingParam model.IdfwMasterSwitchSetting) (model.IdfwMasterSwitchSetting, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(masterSwitchSettingUpdateInputType(), typeConverter)
	sv.AddStructField("IdfwMasterSwitchSetting", idfwMasterSwitchSettingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwMasterSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := masterSwitchSettingUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx.idfw.master_switch_setting", "update", inputDataValue, executionContext)
	var emptyOutput model.IdfwMasterSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), masterSwitchSettingUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwMasterSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
