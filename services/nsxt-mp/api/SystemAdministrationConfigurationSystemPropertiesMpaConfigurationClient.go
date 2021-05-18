// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationSystemPropertiesMpaConfiguration
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

type SystemAdministrationConfigurationSystemPropertiesMpaConfigurationClient interface {

	// Delete the MPA configuration for this node.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteMPAConfiguration() error

	// Retrieve the MPA configuration for this node to identify the Manager nodes with which this node is communicating.
	// @return com.vmware.model.MPAConfigProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadMPAConfiguration() (model.MPAConfigProperties, error)

	// Update the MPA configuration for this node.
	//
	// @param mPAConfigPropertiesParam (required)
	// @return com.vmware.model.MPAConfigProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateMPAConfiguration(mPAConfigPropertiesParam model.MPAConfigProperties) (model.MPAConfigProperties, error)
}

type systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationSystemPropertiesMpaConfigurationClient(connector client.Connector) *systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_system_properties_mpa_configuration")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete_MPA_configuration": core.NewMethodIdentifier(interfaceIdentifier, "delete_MPA_configuration"),
		"read_MPA_configuration":   core.NewMethodIdentifier(interfaceIdentifier, "read_MPA_configuration"),
		"update_MPA_configuration": core.NewMethodIdentifier(interfaceIdentifier, "update_MPA_configuration"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient) DeleteMPAConfiguration() error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesMpaConfigurationDeleteMPAConfigurationInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesMpaConfigurationDeleteMPAConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_mpa_configuration", "delete_MPA_configuration", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient) ReadMPAConfiguration() (model.MPAConfigProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesMpaConfigurationReadMPAConfigurationInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MPAConfigProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesMpaConfigurationReadMPAConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_mpa_configuration", "read_MPA_configuration", inputDataValue, executionContext)
	var emptyOutput model.MPAConfigProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationSystemPropertiesMpaConfigurationReadMPAConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MPAConfigProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationSystemPropertiesMpaConfigurationClient) UpdateMPAConfiguration(mPAConfigPropertiesParam model.MPAConfigProperties) (model.MPAConfigProperties, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesMpaConfigurationUpdateMPAConfigurationInputType(), typeConverter)
	sv.AddStructField("MPAConfigProperties", mPAConfigPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MPAConfigProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesMpaConfigurationUpdateMPAConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_mpa_configuration", "update_MPA_configuration", inputDataValue, executionContext)
	var emptyOutput model.MPAConfigProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationSystemPropertiesMpaConfigurationUpdateMPAConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MPAConfigProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
