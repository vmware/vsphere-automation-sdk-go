// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricProfilesHostswitchProfiles
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

type SystemAdministrationConfigurationFabricProfilesHostswitchProfilesClient interface {

	// Creates a hostswitch profile. The resource_type is required. For uplink profiles, the teaming and policy parameters are required. By default, the mtu is 1600 and the transport_vlan is 0. The supported MTU range is 1280 through (uplink_mtu_threshold). (uplink_mtu_threshold) is 9000 by default. Range can be extended by modifying (uplink_mtu_threshold) in SwitchingGlobalConfig to the required upper threshold.
	//
	// @param baseHostSwitchProfileParam (required)
	// The parameter must contain all the properties defined in model.BaseHostSwitchProfile.
	// @return com.vmware.model.BaseHostSwitchProfile
	// The return value will contain all the properties defined in model.BaseHostSwitchProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateHostSwitchProfile(baseHostSwitchProfileParam *data.StructValue) (*data.StructValue, error)

	// Deletes a specified hostswitch profile.
	//
	// @param hostSwitchProfileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteHostSwitchProfile(hostSwitchProfileIdParam string) error

	// Returns information about a specified hostswitch profile.
	//
	// @param hostSwitchProfileIdParam (required)
	// @return com.vmware.model.BaseHostSwitchProfile
	// The return value will contain all the properties defined in model.BaseHostSwitchProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetHostSwitchProfile(hostSwitchProfileIdParam string) (*data.StructValue, error)

	// Returns information about the configured hostswitch profiles. Hostswitch profiles define networking policies for hostswitches (sometimes referred to as bridges in OVS). Currently, only uplink teaming is supported. Uplink teaming allows NSX to load balance traffic across different physical NICs (PNICs) on the hypervisor hosts. Multiple teaming policies are supported, including LACP active, LACP passive, load balancing based on source ID, and failover order.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param deploymentTypeParam Supported edge deployment type. (optional)
	// @param hostswitchProfileTypeParam Supported HostSwitch profiles. (optional)
	// @param includeSystemOwnedParam Whether the list result contains system resources (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nodeTypeParam Fabric node type for which uplink profiles are to be listed (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param uplinkTeamingPolicyNameParam The host switch profile's uplink teaming policy name (optional)
	// @return com.vmware.model.HostSwitchProfilesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListHostSwitchProfiles(cursorParam *string, deploymentTypeParam *string, hostswitchProfileTypeParam *string, includeSystemOwnedParam *bool, includedFieldsParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, uplinkTeamingPolicyNameParam *string) (model.HostSwitchProfilesListResult, error)

	// Modifies a specified hostswitch profile. The body of the PUT request must include the resource_type. For uplink profiles, the put request must also include teaming parameters. Modifiable attributes include display_name, mtu, and transport_vlan. For uplink teaming policies, uplink_name and policy are also modifiable.
	//
	// @param hostSwitchProfileIdParam (required)
	// @param baseHostSwitchProfileParam (required)
	// The parameter must contain all the properties defined in model.BaseHostSwitchProfile.
	// @return com.vmware.model.BaseHostSwitchProfile
	// The return value will contain all the properties defined in model.BaseHostSwitchProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateHostSwitchProfile(hostSwitchProfileIdParam string, baseHostSwitchProfileParam *data.StructValue) (*data.StructValue, error)
}

type systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricProfilesHostswitchProfilesClient(connector client.Connector) *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_profiles_hostswitch_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_host_switch_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_host_switch_profile"),
		"delete_host_switch_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_host_switch_profile"),
		"get_host_switch_profile":    core.NewMethodIdentifier(interfaceIdentifier, "get_host_switch_profile"),
		"list_host_switch_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_host_switch_profiles"),
		"update_host_switch_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_host_switch_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient) CreateHostSwitchProfile(baseHostSwitchProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesHostswitchProfilesCreateHostSwitchProfileInputType(), typeConverter)
	sv.AddStructField("BaseHostSwitchProfile", baseHostSwitchProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesHostswitchProfilesCreateHostSwitchProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_hostswitch_profiles", "create_host_switch_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesHostswitchProfilesCreateHostSwitchProfileOutputType())
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

func (sIface *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient) DeleteHostSwitchProfile(hostSwitchProfileIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesHostswitchProfilesDeleteHostSwitchProfileInputType(), typeConverter)
	sv.AddStructField("HostSwitchProfileId", hostSwitchProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesHostswitchProfilesDeleteHostSwitchProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_hostswitch_profiles", "delete_host_switch_profile", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient) GetHostSwitchProfile(hostSwitchProfileIdParam string) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesHostswitchProfilesGetHostSwitchProfileInputType(), typeConverter)
	sv.AddStructField("HostSwitchProfileId", hostSwitchProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesHostswitchProfilesGetHostSwitchProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_hostswitch_profiles", "get_host_switch_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesHostswitchProfilesGetHostSwitchProfileOutputType())
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

func (sIface *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient) ListHostSwitchProfiles(cursorParam *string, deploymentTypeParam *string, hostswitchProfileTypeParam *string, includeSystemOwnedParam *bool, includedFieldsParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, uplinkTeamingPolicyNameParam *string) (model.HostSwitchProfilesListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesHostswitchProfilesListHostSwitchProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("DeploymentType", deploymentTypeParam)
	sv.AddStructField("HostswitchProfileType", hostswitchProfileTypeParam)
	sv.AddStructField("IncludeSystemOwned", includeSystemOwnedParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeType", nodeTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("UplinkTeamingPolicyName", uplinkTeamingPolicyNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HostSwitchProfilesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesHostswitchProfilesListHostSwitchProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_hostswitch_profiles", "list_host_switch_profiles", inputDataValue, executionContext)
	var emptyOutput model.HostSwitchProfilesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesHostswitchProfilesListHostSwitchProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HostSwitchProfilesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricProfilesHostswitchProfilesClient) UpdateHostSwitchProfile(hostSwitchProfileIdParam string, baseHostSwitchProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricProfilesHostswitchProfilesUpdateHostSwitchProfileInputType(), typeConverter)
	sv.AddStructField("HostSwitchProfileId", hostSwitchProfileIdParam)
	sv.AddStructField("BaseHostSwitchProfile", baseHostSwitchProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricProfilesHostswitchProfilesUpdateHostSwitchProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_profiles_hostswitch_profiles", "update_host_switch_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricProfilesHostswitchProfilesUpdateHostSwitchProfileOutputType())
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
