// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPISecurityIdentityFirewallRealizationData
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

type ManagementPlaneAPISecurityIdentityFirewallRealizationDataClient interface {

	// Get all Identity Firewall NSGroup VM details for a given NSGroup.
	//
	// @param groupIdParam (required)
	// @return com.vmware.model.IdfwNsgroupVmDetailListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetNsgroupVmDetails(groupIdParam string) (model.IdfwNsgroupVmDetailListResult, error)

	// Get IDFW system statistics data.
	// @return com.vmware.model.IdfwSystemStats
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSystemStats() (model.IdfwSystemStats, error)

	// Get IDFW user login events for a given user (all active plus up to 5 most recent archived entries).
	//
	// @param userIdParam (required)
	// @return com.vmware.model.IdfwUserStats
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetUserStats(userIdParam string) (model.IdfwUserStats, error)

	// Get IDFW user login events for a given VM (all active plus up to 5 most recent archived entries).
	//
	// @param vmExtIdParam (required)
	// @return com.vmware.model.IdfwVmStats
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetVmStats(vmExtIdParam string) (model.IdfwVmStats, error)

	// Get user session data.
	// @return com.vmware.model.IdfwUserSessionDataAndMappings
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListUserSessions() (model.IdfwUserSessionDataAndMappings, error)
}

type managementPlaneAPISecurityIdentityFirewallRealizationDataClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPISecurityIdentityFirewallRealizationDataClient(connector client.Connector) *managementPlaneAPISecurityIdentityFirewallRealizationDataClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_security_identity_firewall_realization_data")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_nsgroup_vm_details": core.NewMethodIdentifier(interfaceIdentifier, "get_nsgroup_vm_details"),
		"get_system_stats":       core.NewMethodIdentifier(interfaceIdentifier, "get_system_stats"),
		"get_user_stats":         core.NewMethodIdentifier(interfaceIdentifier, "get_user_stats"),
		"get_vm_stats":           core.NewMethodIdentifier(interfaceIdentifier, "get_vm_stats"),
		"list_user_sessions":     core.NewMethodIdentifier(interfaceIdentifier, "list_user_sessions"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPISecurityIdentityFirewallRealizationDataClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPISecurityIdentityFirewallRealizationDataClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPISecurityIdentityFirewallRealizationDataClient) GetNsgroupVmDetails(groupIdParam string) (model.IdfwNsgroupVmDetailListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallRealizationDataGetNsgroupVmDetailsInputType(), typeConverter)
	sv.AddStructField("GroupId", groupIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwNsgroupVmDetailListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallRealizationDataGetNsgroupVmDetailsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_realization_data", "get_nsgroup_vm_details", inputDataValue, executionContext)
	var emptyOutput model.IdfwNsgroupVmDetailListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallRealizationDataGetNsgroupVmDetailsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwNsgroupVmDetailListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallRealizationDataClient) GetSystemStats() (model.IdfwSystemStats, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallRealizationDataGetSystemStatsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwSystemStats
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallRealizationDataGetSystemStatsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_realization_data", "get_system_stats", inputDataValue, executionContext)
	var emptyOutput model.IdfwSystemStats
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallRealizationDataGetSystemStatsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwSystemStats), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallRealizationDataClient) GetUserStats(userIdParam string) (model.IdfwUserStats, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallRealizationDataGetUserStatsInputType(), typeConverter)
	sv.AddStructField("UserId", userIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwUserStats
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallRealizationDataGetUserStatsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_realization_data", "get_user_stats", inputDataValue, executionContext)
	var emptyOutput model.IdfwUserStats
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallRealizationDataGetUserStatsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwUserStats), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallRealizationDataClient) GetVmStats(vmExtIdParam string) (model.IdfwVmStats, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallRealizationDataGetVmStatsInputType(), typeConverter)
	sv.AddStructField("VmExtId", vmExtIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwVmStats
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallRealizationDataGetVmStatsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_realization_data", "get_vm_stats", inputDataValue, executionContext)
	var emptyOutput model.IdfwVmStats
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallRealizationDataGetVmStatsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwVmStats), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPISecurityIdentityFirewallRealizationDataClient) ListUserSessions() (model.IdfwUserSessionDataAndMappings, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPISecurityIdentityFirewallRealizationDataListUserSessionsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwUserSessionDataAndMappings
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPISecurityIdentityFirewallRealizationDataListUserSessionsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_security_identity_firewall_realization_data", "list_user_sessions", inputDataValue, executionContext)
	var emptyOutput model.IdfwUserSessionDataAndMappings
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPISecurityIdentityFirewallRealizationDataListUserSessionsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwUserSessionDataAndMappings), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
