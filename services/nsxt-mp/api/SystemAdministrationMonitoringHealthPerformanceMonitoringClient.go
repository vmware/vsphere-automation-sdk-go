// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationMonitoringHealthPerformanceMonitoring
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemAdministrationMonitoringHealthPerformanceMonitoringClient interface {

	// This API is executed on a manager node to return current alarms from all NSX nodes. This API is deprecated as part of alarm framework enhancements. Please use below new APIs: GET /alarms GET /alarms/<alarm-id> POST /alarms/<alarm-id>?action=set_status POST /alarms?action=set_status GET /events GET /events/<event-id> PUT /events/<event-id> POST /events/<event-id>?action=set_default
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param fieldsParam Fields to include in query results (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 100)
	// @return com.vmware.model.AlarmListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CollectAlarms(cursorParam *int64, fieldsParam *string, pageSizeParam *int64) (model.AlarmListResult, error)

	// Read global health performance monitoring configuration
	// @return com.vmware.model.GlobalCollectionConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetAggregationServiceGlobalConfig() (model.GlobalCollectionConfiguration, error)

	// Returns the complete set of client type data collection configuration records for the specified feature stack.
	//
	// @param featureStackNameParam (required)
	// @return com.vmware.model.FeatureStackCollectionConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetFeatureStackConfiguration(featureStackNameParam string) (model.FeatureStackCollectionConfiguration, error)

	// List all health performance monitoring feature stacks
	// @return com.vmware.model.FeatureStackCollectionConfigurationList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListFeatureStackConfigurations() (model.FeatureStackCollectionConfigurationList, error)

	// Reset the data collection frequency configuration setting to the default values
	//
	// @param featureStackNameParam (required)
	// @return com.vmware.model.FeatureStackCollectionConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ResetAggregationServiceFeatureStackConfigurationResetCollectionFrequency(featureStackNameParam string) (model.FeatureStackCollectionConfiguration, error)

	// Set the global configuration for aggregation service related data collection
	//
	// @param globalCollectionConfigurationParam (required)
	// @return com.vmware.model.GlobalCollectionConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateAggregationServiceGlobalConfig(globalCollectionConfigurationParam model.GlobalCollectionConfiguration) (model.GlobalCollectionConfiguration, error)

	// Apply the data collection configuration for the specified feature stack.
	//
	// @param featureStackNameParam (required)
	// @param featureStackCollectionConfigurationParam (required)
	// @return com.vmware.model.FeatureStackCollectionConfiguration
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateFeatureStackConfiguration(featureStackNameParam string, featureStackCollectionConfigurationParam model.FeatureStackCollectionConfiguration) (model.FeatureStackCollectionConfiguration, error)
}

type systemAdministrationMonitoringHealthPerformanceMonitoringClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationMonitoringHealthPerformanceMonitoringClient(connector client.Connector) *systemAdministrationMonitoringHealthPerformanceMonitoringClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_monitoring_health_performance_monitoring")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"collect_alarms":                        core.NewMethodIdentifier(interfaceIdentifier, "collect_alarms"),
		"get_aggregation_service_global_config": core.NewMethodIdentifier(interfaceIdentifier, "get_aggregation_service_global_config"),
		"get_feature_stack_configuration":       core.NewMethodIdentifier(interfaceIdentifier, "get_feature_stack_configuration"),
		"list_feature_stack_configurations":     core.NewMethodIdentifier(interfaceIdentifier, "list_feature_stack_configurations"),
		"reset_aggregation_service_feature_stack_configuration_reset_collection_frequency": core.NewMethodIdentifier(interfaceIdentifier, "reset_aggregation_service_feature_stack_configuration_reset_collection_frequency"),
		"update_aggregation_service_global_config":                                         core.NewMethodIdentifier(interfaceIdentifier, "update_aggregation_service_global_config"),
		"update_feature_stack_configuration":                                               core.NewMethodIdentifier(interfaceIdentifier, "update_feature_stack_configuration"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationMonitoringHealthPerformanceMonitoringClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) CollectAlarms(cursorParam *int64, fieldsParam *string, pageSizeParam *int64) (model.AlarmListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringCollectAlarmsInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Fields", fieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AlarmListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringCollectAlarmsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "collect_alarms", inputDataValue, executionContext)
	var emptyOutput model.AlarmListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringCollectAlarmsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AlarmListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) GetAggregationServiceGlobalConfig() (model.GlobalCollectionConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringGetAggregationServiceGlobalConfigInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GlobalCollectionConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringGetAggregationServiceGlobalConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "get_aggregation_service_global_config", inputDataValue, executionContext)
	var emptyOutput model.GlobalCollectionConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringGetAggregationServiceGlobalConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GlobalCollectionConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) GetFeatureStackConfiguration(featureStackNameParam string) (model.FeatureStackCollectionConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringGetFeatureStackConfigurationInputType(), typeConverter)
	sv.AddStructField("FeatureStackName", featureStackNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FeatureStackCollectionConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringGetFeatureStackConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "get_feature_stack_configuration", inputDataValue, executionContext)
	var emptyOutput model.FeatureStackCollectionConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringGetFeatureStackConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FeatureStackCollectionConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) ListFeatureStackConfigurations() (model.FeatureStackCollectionConfigurationList, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringListFeatureStackConfigurationsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FeatureStackCollectionConfigurationList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringListFeatureStackConfigurationsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "list_feature_stack_configurations", inputDataValue, executionContext)
	var emptyOutput model.FeatureStackCollectionConfigurationList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringListFeatureStackConfigurationsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FeatureStackCollectionConfigurationList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) ResetAggregationServiceFeatureStackConfigurationResetCollectionFrequency(featureStackNameParam string) (model.FeatureStackCollectionConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringResetAggregationServiceFeatureStackConfigurationResetCollectionFrequencyInputType(), typeConverter)
	sv.AddStructField("FeatureStackName", featureStackNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FeatureStackCollectionConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringResetAggregationServiceFeatureStackConfigurationResetCollectionFrequencyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "reset_aggregation_service_feature_stack_configuration_reset_collection_frequency", inputDataValue, executionContext)
	var emptyOutput model.FeatureStackCollectionConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringResetAggregationServiceFeatureStackConfigurationResetCollectionFrequencyOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FeatureStackCollectionConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) UpdateAggregationServiceGlobalConfig(globalCollectionConfigurationParam model.GlobalCollectionConfiguration) (model.GlobalCollectionConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringUpdateAggregationServiceGlobalConfigInputType(), typeConverter)
	sv.AddStructField("GlobalCollectionConfiguration", globalCollectionConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.GlobalCollectionConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringUpdateAggregationServiceGlobalConfigRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "update_aggregation_service_global_config", inputDataValue, executionContext)
	var emptyOutput model.GlobalCollectionConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringUpdateAggregationServiceGlobalConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.GlobalCollectionConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringHealthPerformanceMonitoringClient) UpdateFeatureStackConfiguration(featureStackNameParam string, featureStackCollectionConfigurationParam model.FeatureStackCollectionConfiguration) (model.FeatureStackCollectionConfiguration, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringHealthPerformanceMonitoringUpdateFeatureStackConfigurationInputType(), typeConverter)
	sv.AddStructField("FeatureStackName", featureStackNameParam)
	sv.AddStructField("FeatureStackCollectionConfiguration", featureStackCollectionConfigurationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.FeatureStackCollectionConfiguration
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringHealthPerformanceMonitoringUpdateFeatureStackConfigurationRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_health_performance_monitoring", "update_feature_stack_configuration", inputDataValue, executionContext)
	var emptyOutput model.FeatureStackCollectionConfiguration
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringHealthPerformanceMonitoringUpdateFeatureStackConfigurationOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.FeatureStackCollectionConfiguration), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
