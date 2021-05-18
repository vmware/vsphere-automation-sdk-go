// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationNSXManagersManagerConfiguration
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationNSXManagersManagerConfigurationClient interface {

	// Returns the NSX Management nodes global configuration.
	// @return com.vmware.model.ManagementConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadManagementConfig() (model.ManagementConfig, error)

	// Modifies the NSX Management nodes global configuration.
	//
	// @param managementConfigParam (required)
	// @return com.vmware.model.ManagementConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateManagementConfig(managementConfigParam model.ManagementConfig) (model.ManagementConfig, error)
}

type systemAdministrationConfigurationNSXManagersManagerConfigurationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationNSXManagersManagerConfigurationClient(connector client.Connector) *systemAdministrationConfigurationNSXManagersManagerConfigurationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_NSX_managers_manager_configuration")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"read_management_config":   core.NewMethodIdentifier(interfaceIdentifier, "read_management_config"),
		"update_management_config": core.NewMethodIdentifier(interfaceIdentifier, "update_management_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationNSXManagersManagerConfigurationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationNSXManagersManagerConfigurationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationNSXManagersManagerConfigurationClient) ReadManagementConfig() (model.ManagementConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersManagerConfigurationReadManagementConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ManagementConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersManagerConfigurationReadManagementConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_manager_configuration", "read_management_config", inputDataValue, executionContext)
	var emptyOutput model.ManagementConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersManagerConfigurationReadManagementConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ManagementConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationNSXManagersManagerConfigurationClient) UpdateManagementConfig(managementConfigParam model.ManagementConfig) (model.ManagementConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationNSXManagersManagerConfigurationUpdateManagementConfigInputType(), typeConverter)
	sv.AddStructField("ManagementConfig", managementConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ManagementConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationNSXManagersManagerConfigurationUpdateManagementConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_NSX_managers_manager_configuration", "update_management_config", inputDataValue, executionContext)
	var emptyOutput model.ManagementConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationNSXManagersManagerConfigurationUpdateManagementConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ManagementConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
