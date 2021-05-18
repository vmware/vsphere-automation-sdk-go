// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationMonitoringSystemHealthProfiles
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

type SystemAdministrationMonitoringSystemHealthProfilesClient interface {

	// Create a system health profile.
	//
	// @param systemHealthAgentProfileParam (required)
	// @return com.vmware.model.SystemHealthAgentProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateSystemHealthAgentProfile(systemHealthAgentProfileParam model.SystemHealthAgentProfile) (model.SystemHealthAgentProfile, error)

	// Delete an existing system health profile by ID.
	//
	// @param profileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteSystemHealthAgentProfile(profileIdParam string) error

	// List all system health profiles.
	// @return com.vmware.model.SystemHealthAgentProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListSystemHealthAgentProfiles() (model.SystemHealthAgentProfileListResult, error)

	// Show the details of a system health profile.
	//
	// @param profileIdParam (required)
	// @return com.vmware.model.SystemHealthAgentProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ShowSystemHealthAgentProfile(profileIdParam string) (model.SystemHealthAgentProfile, error)

	// Update a system health profile definition.
	//
	// @param profileIdParam (required)
	// @param systemHealthAgentProfileParam (required)
	// @return com.vmware.model.SystemHealthAgentProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateSystemHealthAgentProfile(profileIdParam string, systemHealthAgentProfileParam model.SystemHealthAgentProfile) (model.SystemHealthAgentProfile, error)
}

type systemAdministrationMonitoringSystemHealthProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationMonitoringSystemHealthProfilesClient(connector client.Connector) *systemAdministrationMonitoringSystemHealthProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_monitoring_system_health_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_system_health_agent_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_system_health_agent_profile"),
		"delete_system_health_agent_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_system_health_agent_profile"),
		"list_system_health_agent_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_system_health_agent_profiles"),
		"show_system_health_agent_profile":   core.NewMethodIdentifier(interfaceIdentifier, "show_system_health_agent_profile"),
		"update_system_health_agent_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_system_health_agent_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationMonitoringSystemHealthProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationMonitoringSystemHealthProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationMonitoringSystemHealthProfilesClient) CreateSystemHealthAgentProfile(systemHealthAgentProfileParam model.SystemHealthAgentProfile) (model.SystemHealthAgentProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringSystemHealthProfilesCreateSystemHealthAgentProfileInputType(), typeConverter)
	sv.AddStructField("SystemHealthAgentProfile", systemHealthAgentProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SystemHealthAgentProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringSystemHealthProfilesCreateSystemHealthAgentProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_system_health_profiles", "create_system_health_agent_profile", inputDataValue, executionContext)
	var emptyOutput model.SystemHealthAgentProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringSystemHealthProfilesCreateSystemHealthAgentProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SystemHealthAgentProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringSystemHealthProfilesClient) DeleteSystemHealthAgentProfile(profileIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringSystemHealthProfilesDeleteSystemHealthAgentProfileInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringSystemHealthProfilesDeleteSystemHealthAgentProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_system_health_profiles", "delete_system_health_agent_profile", inputDataValue, executionContext)
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

func (sIface *systemAdministrationMonitoringSystemHealthProfilesClient) ListSystemHealthAgentProfiles() (model.SystemHealthAgentProfileListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringSystemHealthProfilesListSystemHealthAgentProfilesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SystemHealthAgentProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringSystemHealthProfilesListSystemHealthAgentProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_system_health_profiles", "list_system_health_agent_profiles", inputDataValue, executionContext)
	var emptyOutput model.SystemHealthAgentProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringSystemHealthProfilesListSystemHealthAgentProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SystemHealthAgentProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringSystemHealthProfilesClient) ShowSystemHealthAgentProfile(profileIdParam string) (model.SystemHealthAgentProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringSystemHealthProfilesShowSystemHealthAgentProfileInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SystemHealthAgentProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringSystemHealthProfilesShowSystemHealthAgentProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_system_health_profiles", "show_system_health_agent_profile", inputDataValue, executionContext)
	var emptyOutput model.SystemHealthAgentProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringSystemHealthProfilesShowSystemHealthAgentProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SystemHealthAgentProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationMonitoringSystemHealthProfilesClient) UpdateSystemHealthAgentProfile(profileIdParam string, systemHealthAgentProfileParam model.SystemHealthAgentProfile) (model.SystemHealthAgentProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationMonitoringSystemHealthProfilesUpdateSystemHealthAgentProfileInputType(), typeConverter)
	sv.AddStructField("ProfileId", profileIdParam)
	sv.AddStructField("SystemHealthAgentProfile", systemHealthAgentProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SystemHealthAgentProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationMonitoringSystemHealthProfilesUpdateSystemHealthAgentProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_monitoring_system_health_profiles", "update_system_health_agent_profile", inputDataValue, executionContext)
	var emptyOutput model.SystemHealthAgentProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationMonitoringSystemHealthProfilesUpdateSystemHealthAgentProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SystemHealthAgentProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
