// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationSystemPropertiesRealizationStateBarrier
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

type SystemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient interface {

	// Returns the current global realization barrier number for NSX. This method has been deprecated. To track realization state, use X-NSX-REQUESTID request header instead.
	// @return com.vmware.model.CurrentRealizationStateBarrier
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetCurrentBarrier() (model.CurrentRealizationStateBarrier, error)

	// Returns the current barrier configuration
	// @return com.vmware.model.RealizationStateBarrierConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetRealizationStateBarrierConfig() (model.RealizationStateBarrierConfig, error)

	// Increment the current barrier number by 1 for NSX. This method has been deprecated. To track realization state, use X-NSX-REQUESTID request header instead.
	// @return com.vmware.model.CurrentRealizationStateBarrier
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	IncrementRealizationStateBarrierIncrement() (model.CurrentRealizationStateBarrier, error)

	// Updates the barrier configuration having interval set in milliseconds The new interval that automatically increments the global realization number
	//
	// @param realizationStateBarrierConfigParam (required)
	// @return com.vmware.model.RealizationStateBarrierConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateRealizationStateBarrierConfig(realizationStateBarrierConfigParam model.RealizationStateBarrierConfig) (model.RealizationStateBarrierConfig, error)
}

type systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient(connector client.Connector) *systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_system_properties_realization_state_barrier")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_current_barrier":                           core.NewMethodIdentifier(interfaceIdentifier, "get_current_barrier"),
		"get_realization_state_barrier_config":          core.NewMethodIdentifier(interfaceIdentifier, "get_realization_state_barrier_config"),
		"increment_realization_state_barrier_increment": core.NewMethodIdentifier(interfaceIdentifier, "increment_realization_state_barrier_increment"),
		"update_realization_state_barrier_config":       core.NewMethodIdentifier(interfaceIdentifier, "update_realization_state_barrier_config"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient) GetCurrentBarrier() (model.CurrentRealizationStateBarrier, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierGetCurrentBarrierInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CurrentRealizationStateBarrier
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierGetCurrentBarrierRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_realization_state_barrier", "get_current_barrier", inputDataValue, executionContext)
	var emptyOutput model.CurrentRealizationStateBarrier
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierGetCurrentBarrierOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CurrentRealizationStateBarrier), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient) GetRealizationStateBarrierConfig() (model.RealizationStateBarrierConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierGetRealizationStateBarrierConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RealizationStateBarrierConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierGetRealizationStateBarrierConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_realization_state_barrier", "get_realization_state_barrier_config", inputDataValue, executionContext)
	var emptyOutput model.RealizationStateBarrierConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierGetRealizationStateBarrierConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RealizationStateBarrierConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient) IncrementRealizationStateBarrierIncrement() (model.CurrentRealizationStateBarrier, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierIncrementRealizationStateBarrierIncrementInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CurrentRealizationStateBarrier
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierIncrementRealizationStateBarrierIncrementRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_realization_state_barrier", "increment_realization_state_barrier_increment", inputDataValue, executionContext)
	var emptyOutput model.CurrentRealizationStateBarrier
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierIncrementRealizationStateBarrierIncrementOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CurrentRealizationStateBarrier), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierClient) UpdateRealizationStateBarrierConfig(realizationStateBarrierConfigParam model.RealizationStateBarrierConfig) (model.RealizationStateBarrierConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierUpdateRealizationStateBarrierConfigInputType(), typeConverter)
	sv.AddStructField("RealizationStateBarrierConfig", realizationStateBarrierConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RealizationStateBarrierConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierUpdateRealizationStateBarrierConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_system_properties_realization_state_barrier", "update_realization_state_barrier_config", inputDataValue, executionContext)
	var emptyOutput model.RealizationStateBarrierConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationSystemPropertiesRealizationStateBarrierUpdateRealizationStateBarrierConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RealizationStateBarrierConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
