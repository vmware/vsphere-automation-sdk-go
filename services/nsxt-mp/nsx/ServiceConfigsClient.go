// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ServiceConfigs
// Used by client-side stubs.

package nsx

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ServiceConfigsClient interface {

	// Creates a new service config that can group profiles and configs
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param serviceConfigParam (required)
	// @return com.vmware.nsx.model.ServiceConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(serviceConfigParam nsxModel.ServiceConfig) (nsxModel.ServiceConfig, error)

	// Deletes the specified service config
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param configSetIdParam service Ccnfig Id (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(configSetIdParam string) error

	// Returns information about the specified Service Config.
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param configSetIdParam Service Config Id (required)
	// @return com.vmware.nsx.model.ServiceConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(configSetIdParam string) (nsxModel.ServiceConfig, error)

	// List of all service configs.
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param profileTypeParam It fetches ServiceConfig for the given profile_type. Only one type of supported profile type can be mentioned in a single API call. API will return all ServiceConfig if this field is not passed. (optional)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.ServiceConfigListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypeParam *string, sortAscendingParam *bool, sortByParam *string) (nsxModel.ServiceConfigListResult, error)

	// Updates the specified ServiceConfig.
	//  Please use APIs in the Policy section corresponding to the profile type - IPFIX_SWITCH_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Profiles IPFIX_COLLECTOR_PROFILE - Policy > Monitoring > IPFIX > Switch IPFIX Collectors DFW_CPU_MEM_THRESHOLDS_PROFILE - Policy > Security > Security Profiles > Cpu Memory Thresholds Profiles DFW_SESSIONTIMER_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles GI_SERVICE_PROFILE - Policy > Security > Service Insertion > Service References DFW_FLOOD_PROTECTION_PROFILE - Policy > Security > Security Profiles > Flood Protection Profiles GENERAL_SECURITY_SETTINGS_PROFILE - Policy > Security > Security Profiles > General Security Profiles DFW_DNS_PROFILE - Policy > Security > Security Profiles > DNS Security Profiles PACE_HOST_CONFIG_PROFILE SHA_PROFILE - Policy > Monitoring > System Health Agent LATENCY_STAT_PROFILE - Policy > Monitoring > Latency IP_DISCOVERY_SWITCHING_UPM_PROFILE - Policy > Networking > Connectivity > Segments > Segment Profiles > IP Discovery Profiles SYSTEM_HEALTH_PLUGIN_PROFILE - Policy > Monitoring > System Health Agent COMMON_AGENT_HOST_CONFIG_PROFILE ODS_PROFILE - Policy > Monitoring > Online Diagnostic System ODS_DYNAMIC_INSTANCE - Policy > Monitoring > Online Diagnostic System
	//
	// Deprecated: This API element is deprecated.
	//
	// @param configSetIdParam service config Id (required)
	// @param serviceConfigParam (required)
	// @return com.vmware.nsx.model.ServiceConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(configSetIdParam string, serviceConfigParam nsxModel.ServiceConfig) (nsxModel.ServiceConfig, error)
}

type serviceConfigsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewServiceConfigsClient(connector vapiProtocolClient_.Connector) *serviceConfigsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.service_configs")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"create": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := serviceConfigsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *serviceConfigsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *serviceConfigsClient) Create(serviceConfigParam nsxModel.ServiceConfig) (nsxModel.ServiceConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceConfigsCreateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceConfigsCreateInputType(), typeConverter)
	sv.AddStructField("ServiceConfig", serviceConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs", "create", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceConfigsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceConfigsClient) Delete(configSetIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceConfigsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceConfigsDeleteInputType(), typeConverter)
	sv.AddStructField("ConfigSetId", configSetIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *serviceConfigsClient) Get(configSetIdParam string) (nsxModel.ServiceConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceConfigsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceConfigsGetInputType(), typeConverter)
	sv.AddStructField("ConfigSetId", configSetIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceConfigsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceConfigsClient) List(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, profileTypeParam *string, sortAscendingParam *bool, sortByParam *string) (nsxModel.ServiceConfigListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceConfigsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceConfigsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ProfileType", profileTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceConfigListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceConfigsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *serviceConfigsClient) Update(configSetIdParam string, serviceConfigParam nsxModel.ServiceConfig) (nsxModel.ServiceConfig, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := serviceConfigsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(serviceConfigsUpdateInputType(), typeConverter)
	sv.AddStructField("ConfigSetId", configSetIdParam)
	sv.AddStructField("ServiceConfig", serviceConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.ServiceConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.service_configs", "update", inputDataValue, executionContext)
	var emptyOutput nsxModel.ServiceConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ServiceConfigsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.ServiceConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
