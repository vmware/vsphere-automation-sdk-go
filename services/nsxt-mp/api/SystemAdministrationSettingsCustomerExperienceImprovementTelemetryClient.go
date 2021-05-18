// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationSettingsCustomerExperienceImprovementTelemetry
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

type SystemAdministrationSettingsCustomerExperienceImprovementTelemetryClient interface {

	// Returns telemetry agreement information.
	// @return com.vmware.model.TelemetryAgreement
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTelemetryAgreement() (model.TelemetryAgreement, error)

	// Returns the telemetry configuration.
	// @return com.vmware.model.TelemetryConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTelemetryConfig() (model.TelemetryConfig, error)

	// Set telemetry agreement information.
	//
	// @param telemetryAgreementParam (required)
	// @return com.vmware.model.TelemetryAgreement
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTelemetryAgreement(telemetryAgreementParam model.TelemetryAgreement) (model.TelemetryAgreement, error)

	// Updates or creates the telemetry configuration, and returns the new configuration.
	//
	// @param telemetryConfigParam (required)
	// @return com.vmware.model.TelemetryConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTelemetryConfig(telemetryConfigParam model.TelemetryConfig) (model.TelemetryConfig, error)
}

type systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationSettingsCustomerExperienceImprovementTelemetryClient(connector client.Connector) *systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_settings_customer_experience_improvement_telemetry")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_telemetry_agreement":    core.NewMethodIdentifier(interfaceIdentifier, "get_telemetry_agreement"),
		"get_telemetry_config":       core.NewMethodIdentifier(interfaceIdentifier, "get_telemetry_config"),
		"update_telemetry_agreement": core.NewMethodIdentifier(interfaceIdentifier, "update_telemetry_agreement"),
		"update_telemetry_config":    core.NewMethodIdentifier(interfaceIdentifier, "update_telemetry_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient) GetTelemetryAgreement() (model.TelemetryAgreement, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCustomerExperienceImprovementTelemetryGetTelemetryAgreementInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TelemetryAgreement
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCustomerExperienceImprovementTelemetryGetTelemetryAgreementRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_customer_experience_improvement_telemetry", "get_telemetry_agreement", inputDataValue, executionContext)
	var emptyOutput model.TelemetryAgreement
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCustomerExperienceImprovementTelemetryGetTelemetryAgreementOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TelemetryAgreement), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient) GetTelemetryConfig() (model.TelemetryConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCustomerExperienceImprovementTelemetryGetTelemetryConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TelemetryConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCustomerExperienceImprovementTelemetryGetTelemetryConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_customer_experience_improvement_telemetry", "get_telemetry_config", inputDataValue, executionContext)
	var emptyOutput model.TelemetryConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCustomerExperienceImprovementTelemetryGetTelemetryConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TelemetryConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient) UpdateTelemetryAgreement(telemetryAgreementParam model.TelemetryAgreement) (model.TelemetryAgreement, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCustomerExperienceImprovementTelemetryUpdateTelemetryAgreementInputType(), typeConverter)
	sv.AddStructField("TelemetryAgreement", telemetryAgreementParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TelemetryAgreement
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCustomerExperienceImprovementTelemetryUpdateTelemetryAgreementRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_customer_experience_improvement_telemetry", "update_telemetry_agreement", inputDataValue, executionContext)
	var emptyOutput model.TelemetryAgreement
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCustomerExperienceImprovementTelemetryUpdateTelemetryAgreementOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TelemetryAgreement), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationSettingsCustomerExperienceImprovementTelemetryClient) UpdateTelemetryConfig(telemetryConfigParam model.TelemetryConfig) (model.TelemetryConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationSettingsCustomerExperienceImprovementTelemetryUpdateTelemetryConfigInputType(), typeConverter)
	sv.AddStructField("TelemetryConfig", telemetryConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TelemetryConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationSettingsCustomerExperienceImprovementTelemetryUpdateTelemetryConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_settings_customer_experience_improvement_telemetry", "update_telemetry_config", inputDataValue, executionContext)
	var emptyOutput model.TelemetryConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationSettingsCustomerExperienceImprovementTelemetryUpdateTelemetryConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TelemetryConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
