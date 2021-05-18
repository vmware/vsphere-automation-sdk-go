// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricProfilesTransportNodeProfiles
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

type SystemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient interface {

	// Transport node profile captures the configuration needed to create a transport node. A transport node profile can be attached to compute collections for automatic TN creation of member hosts.
	//
	// @param transportNodeProfileParam (required)
	// @return com.vmware.model.TransportNodeProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateTransportNodeProfile(transportNodeProfileParam model.TransportNodeProfile) (model.TransportNodeProfile, error)

	// Deletes the specified transport node profile. A transport node profile can be deleted only when it is not attached to any compute collection.
	//
	// @param transportNodeProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteTransportNodeProfile(transportNodeProfileIdParam string) error

	// Returns information about a specified transport node profile.
	//
	// @param transportNodeProfileIdParam (required)
	// @return com.vmware.model.TransportNodeProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTransportNodeProfile(transportNodeProfileIdParam string) (model.TransportNodeProfile, error)

	// Returns information about all transport node profiles.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.TransportNodeProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTransportNodeProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.TransportNodeProfileListResult, error)

	// When configurations of a transport node profile(TNP) is updated, all the transport nodes in all the compute collections to which this TNP is attached are updated to reflect the updated configuration.
	//
	// @param transportNodeProfileIdParam (required)
	// @param transportNodeProfileParam (required)
	// @param esxMgmtIfMigrationDestParam The network ids to which the ESX vmk interfaces will be migrated (optional)
	// @param ifIdParam The ESX vmk interfaces to migrate (optional)
	// @param pingIpParam IP Addresses to ping right after ESX vmk interfaces were migrated. (optional)
	// @param skipValidationParam Whether to skip front-end validation for vmk/vnic/pnic migration (optional, default to false)
	// @param vnicParam The ESX vmk interfaces and/or VM NIC to migrate (optional)
	// @param vnicMigrationDestParam The migration destinations of ESX vmk interfaces and/or VM NIC (optional)
	// @return com.vmware.model.TransportNodeProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTransportNodeProfile(transportNodeProfileIdParam string, transportNodeProfileParam model.TransportNodeProfile, esxMgmtIfMigrationDestParam *string, ifIdParam *string, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (model.TransportNodeProfile, error)
}

type systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient(connector client.Connector) *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_profiles_transport_node_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_transport_node_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_transport_node_profile"),
		"delete_transport_node_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_transport_node_profile"),
		"get_transport_node_profile":    core.NewMethodIdentifier(interfaceIdentifier, "get_transport_node_profile"),
		"list_transport_node_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_transport_node_profiles"),
		"update_transport_node_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_transport_node_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient) CreateTransportNodeProfile(transportNodeProfileParam model.TransportNodeProfile) (model.TransportNodeProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesTransportNodeProfilesCreateTransportNodeProfileInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfile", transportNodeProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesTransportNodeProfilesCreateTransportNodeProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_transport_node_profiles", "create_transport_node_profile", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesTransportNodeProfilesCreateTransportNodeProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient) DeleteTransportNodeProfile(transportNodeProfileIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesTransportNodeProfilesDeleteTransportNodeProfileInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfileId", transportNodeProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesTransportNodeProfilesDeleteTransportNodeProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_transport_node_profiles", "delete_transport_node_profile", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient) GetTransportNodeProfile(transportNodeProfileIdParam string) (model.TransportNodeProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesTransportNodeProfilesGetTransportNodeProfileInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfileId", transportNodeProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesTransportNodeProfilesGetTransportNodeProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_transport_node_profiles", "get_transport_node_profile", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesTransportNodeProfilesGetTransportNodeProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient) ListTransportNodeProfiles(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.TransportNodeProfileListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesTransportNodeProfilesListTransportNodeProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesTransportNodeProfilesListTransportNodeProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_transport_node_profiles", "list_transport_node_profiles", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesTransportNodeProfilesListTransportNodeProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricProfilesTransportNodeProfilesClient) UpdateTransportNodeProfile(transportNodeProfileIdParam string, transportNodeProfileParam model.TransportNodeProfile, esxMgmtIfMigrationDestParam *string, ifIdParam *string, pingIpParam *string, skipValidationParam *bool, vnicParam *string, vnicMigrationDestParam *string) (model.TransportNodeProfile, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesTransportNodeProfilesUpdateTransportNodeProfileInputType(), typeConverter)
	sv.AddStructField("TransportNodeProfileId", transportNodeProfileIdParam)
	sv.AddStructField("TransportNodeProfile", transportNodeProfileParam)
	sv.AddStructField("EsxMgmtIfMigrationDest", esxMgmtIfMigrationDestParam)
	sv.AddStructField("IfId", ifIdParam)
	sv.AddStructField("PingIp", pingIpParam)
	sv.AddStructField("SkipValidation", skipValidationParam)
	sv.AddStructField("Vnic", vnicParam)
	sv.AddStructField("VnicMigrationDest", vnicMigrationDestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesTransportNodeProfilesUpdateTransportNodeProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_transport_node_profiles", "update_transport_node_profile", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesTransportNodeProfilesUpdateTransportNodeProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
