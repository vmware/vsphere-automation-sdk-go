// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationGlobalConfigurations
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationConfigurationGlobalConfigurationsClient interface {

	// Returns global configurations that belong to the config type
	//
	// @param configTypeParam (required)
	// @return com.vmware.model.GlobalConfigs
	// The return value will contain all the properties defined in model.GlobalConfigs.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetGlobalConfigs(configTypeParam string) (*data.StructValue, error)

	// Returns list of all Central Node Config profiles.
	// @return com.vmware.model.CentralNodeConfigProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListCentralNodeConfigProfiles() (model.CentralNodeConfigProfileListResult, error)

	// Returns global configurations of a NSX domain grouped by the config types. These global configurations are valid across NSX domain for their respective types unless they are overridden by a more granular configurations.
	// @return com.vmware.model.GlobalConfigsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListGlobalConfigs() (model.GlobalConfigsListResult, error)

	// Returns properties in specified Central Node Config profile. Sensitive data (like SNMP v2c community strings) are included only if query parameter \"show_sensitive_data\" is true.
	//
	// @param profileIdParam Central Node Config profile id (required)
	// @param showSensitiveDataParam Show sensitive data in Central Node Config profile (optional, default to false)
	// @return com.vmware.model.CentralNodeConfigProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadCentralNodeConfigProfile(profileIdParam string, showSensitiveDataParam *bool) (model.CentralNodeConfigProfile, error)

	// It is similar to update global configurations but this request would trigger update even if the configs are unmodified. However, the realization of the new configurations is config-type specific. Refer to config-type specific documentation for details about the configuration push state.
	//
	// @param configTypeParam (required)
	// @param globalConfigsParam (required)
	// The parameter must contain all the properties defined in model.GlobalConfigs.
	// @return com.vmware.model.GlobalConfigs
	// The return value will contain all the properties defined in model.GlobalConfigs.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResyncGlobalConfigsResyncConfig(configTypeParam string, globalConfigsParam *data.StructValue) (*data.StructValue, error)

	// Updates properties in the specified Central Node Config profile.
	//
	// @param nodeConfigProfileIdParam (required)
	// @param centralNodeConfigProfileParam (required)
	// @return com.vmware.model.CentralNodeConfigProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateCentralNodeConfigProfile(nodeConfigProfileIdParam string, centralNodeConfigProfileParam model.CentralNodeConfigProfile) (model.CentralNodeConfigProfile, error)

	//
	//
	// @param configTypeParam (required)
	// @param globalConfigsParam (required)
	// The parameter must contain all the properties defined in model.GlobalConfigs.
	// @return com.vmware.model.GlobalConfigs
	// The return value will contain all the properties defined in model.GlobalConfigs.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateGlobalConfigs(configTypeParam string, globalConfigsParam *data.StructValue) (*data.StructValue, error)
}

type systemAdministrationConfigurationGlobalConfigurationsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationGlobalConfigurationsClient(connector client.Connector) *systemAdministrationConfigurationGlobalConfigurationsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_global_configurations")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_global_configs":                  core.NewMethodIdentifier(interfaceIdentifier, "get_global_configs"),
		"list_central_node_config_profiles":   core.NewMethodIdentifier(interfaceIdentifier, "list_central_node_config_profiles"),
		"list_global_configs":                 core.NewMethodIdentifier(interfaceIdentifier, "list_global_configs"),
		"read_central_node_config_profile":    core.NewMethodIdentifier(interfaceIdentifier, "read_central_node_config_profile"),
		"resync_global_configs_resync_config": core.NewMethodIdentifier(interfaceIdentifier, "resync_global_configs_resync_config"),
		"update_central_node_config_profile":  core.NewMethodIdentifier(interfaceIdentifier, "update_central_node_config_profile"),
		"update_global_configs":               core.NewMethodIdentifier(interfaceIdentifier, "update_global_configs"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationGlobalConfigurationsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) GetGlobalConfigs(configTypeParam string) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsGetGlobalConfigsInputType(), typeConverter)
	sv.AddStructField("ConfigType", configTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsGetGlobalConfigsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "get_global_configs", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsGetGlobalConfigsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) ListCentralNodeConfigProfiles() (model.CentralNodeConfigProfileListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsListCentralNodeConfigProfilesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CentralNodeConfigProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsListCentralNodeConfigProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "list_central_node_config_profiles", inputDataValue, executionContext)
	var emptyOutput model.CentralNodeConfigProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsListCentralNodeConfigProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CentralNodeConfigProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) ListGlobalConfigs() (model.GlobalConfigsListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsListGlobalConfigsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GlobalConfigsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsListGlobalConfigsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "list_global_configs", inputDataValue, executionContext)
	var emptyOutput model.GlobalConfigsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsListGlobalConfigsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GlobalConfigsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) ReadCentralNodeConfigProfile(profileIdParam string, showSensitiveDataParam *bool) (model.CentralNodeConfigProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsReadCentralNodeConfigProfileInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("ShowSensitiveData", showSensitiveDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CentralNodeConfigProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsReadCentralNodeConfigProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "read_central_node_config_profile", inputDataValue, executionContext)
	var emptyOutput model.CentralNodeConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsReadCentralNodeConfigProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CentralNodeConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) ResyncGlobalConfigsResyncConfig(configTypeParam string, globalConfigsParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsResyncGlobalConfigsResyncConfigInputType(), typeConverter)
	sv.AddStructField("ConfigType", configTypeParam)
	sv.AddStructField("GlobalConfigs", globalConfigsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsResyncGlobalConfigsResyncConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "resync_global_configs_resync_config", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsResyncGlobalConfigsResyncConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) UpdateCentralNodeConfigProfile(nodeConfigProfileIdParam string, centralNodeConfigProfileParam model.CentralNodeConfigProfile) (model.CentralNodeConfigProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsUpdateCentralNodeConfigProfileInputType(), typeConverter)
	sv.AddStructField("NodeConfigProfileId", nodeConfigProfileIdParam)
	sv.AddStructField("CentralNodeConfigProfile", centralNodeConfigProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CentralNodeConfigProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsUpdateCentralNodeConfigProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "update_central_node_config_profile", inputDataValue, executionContext)
	var emptyOutput model.CentralNodeConfigProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsUpdateCentralNodeConfigProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CentralNodeConfigProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationGlobalConfigurationsClient) UpdateGlobalConfigs(configTypeParam string, globalConfigsParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationGlobalConfigurationsUpdateGlobalConfigsInputType(), typeConverter)
	sv.AddStructField("ConfigType", configTypeParam)
	sv.AddStructField("GlobalConfigs", globalConfigsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationGlobalConfigurationsUpdateGlobalConfigsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_global_configurations", "update_global_configs", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationGlobalConfigurationsUpdateGlobalConfigsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
