// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemAdministrationConfigurationFabricTransportZones
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

type SystemAdministrationConfigurationFabricTransportZonesClient interface {

	// Creates a new transport zone. The required parameters are host_switch_name and transport_type (OVERLAY or VLAN). The optional parameters are description and display_name.
	//
	// @param transportZoneParam (required)
	// @return com.vmware.model.TransportZone
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateTransportZone(transportZoneParam model.TransportZone) (model.TransportZone, error)

	// Deletes an existing transport zone.
	//
	// @param zoneIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteTransportZone(zoneIdParam string) error

	// Get high-level summary of a transport zone. The service layer does not support source = realtime or cached.
	// @return com.vmware.model.HeatMapTransportNodesAggregateStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetAllTransportZoneStatus() (model.HeatMapTransportNodesAggregateStatus, error)

	// Get high-level summary of a transport zone
	//
	// @param zoneIdParam ID of transport zone (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.HeatMapTransportZoneStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetHeatmapTransportZoneStatus(zoneIdParam string, sourceParam *string) (model.HeatMapTransportZoneStatus, error)

	// Returns information about a single transport zone.
	//
	// @param zoneIdParam (required)
	// @return com.vmware.model.TransportZone
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTransportZone(zoneIdParam string) (model.TransportZone, error)

	// Returns information about a specified transport zone, including the number of logical switches in the transport zone, number of logical spitch ports assigned to the transport zone, and number of transport nodes in the transport zone.
	//
	// @param zoneIdParam (required)
	// @return com.vmware.model.TransportZoneStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetTransportZoneStatus(zoneIdParam string) (model.TransportZoneStatus, error)

	// Returns information about configured transport zones. NSX requires at least one transport zone. NSX uses transport zones to provide connectivity based on the topology of the underlying network, trust zones, or organizational separations. For example, you might have hypervisors that use one network for management traffic and a different network for VM traffic. This architecture would require two transport zones. The combination of transport zones plus transport connectors enables NSX to form tunnels between hypervisors. Transport zones define which interfaces on the hypervisors can communicate with which other interfaces on other hypervisors to establish overlay tunnels or provide connectivity to a VLAN. A logical switch can be in one (and only one) transport zone. This means that all of a switch's interfaces must be in the same transport zone. However, each hypervisor virtual switch (OVS or VDS) has multiple interfaces (connectors), and each connector can be attached to a different logical switch. For example, on a single hypervisor with two connectors, connector A can be attached to logical switch 1 in transport zone A, while connector B is attached to logical switch 2 in transport zone B. In this way, a single hypervisor can participate in multiple transport zones. The API for creating a transport zone requires that a single host switch be specified for each transport zone, and multiple transport zones can share the same host switch.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeSystemOwnedParam Filter to indicate whether to include system owned Transport Zones. (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param isDefaultParam Filter to choose if default transport zones will be returned (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param transportTypeParam Filter to choose the type of transport zones to return (optional)
	// @param uplinkTeamingPolicyNameParam The transport zone's uplink teaming policy name (optional)
	// @return com.vmware.model.TransportZoneListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListTransportZones(cursorParam *string, includeSystemOwnedParam *bool, includedFieldsParam *string, isDefaultParam *bool, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportTypeParam *string, uplinkTeamingPolicyNameParam *string) (model.TransportZoneListResult, error)

	// Updates an existing transport zone. Modifiable parameters are is_default, description, and display_name. The request must include the existing host_switch_name.
	//
	// @param zoneIdParam (required)
	// @param transportZoneParam (required)
	// @return com.vmware.model.TransportZone
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateTransportZone(zoneIdParam string, transportZoneParam model.TransportZone) (model.TransportZone, error)
}

type systemAdministrationConfigurationFabricTransportZonesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemAdministrationConfigurationFabricTransportZonesClient(connector client.Connector) *systemAdministrationConfigurationFabricTransportZonesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.system_administration_configuration_fabric_transport_zones")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_transport_zone":             core.NewMethodIdentifier(interfaceIdentifier, "create_transport_zone"),
		"delete_transport_zone":             core.NewMethodIdentifier(interfaceIdentifier, "delete_transport_zone"),
		"get_all_transport_zone_status":     core.NewMethodIdentifier(interfaceIdentifier, "get_all_transport_zone_status"),
		"get_heatmap_transport_zone_status": core.NewMethodIdentifier(interfaceIdentifier, "get_heatmap_transport_zone_status"),
		"get_transport_zone":                core.NewMethodIdentifier(interfaceIdentifier, "get_transport_zone"),
		"get_transport_zone_status":         core.NewMethodIdentifier(interfaceIdentifier, "get_transport_zone_status"),
		"list_transport_zones":              core.NewMethodIdentifier(interfaceIdentifier, "list_transport_zones"),
		"update_transport_zone":             core.NewMethodIdentifier(interfaceIdentifier, "update_transport_zone"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemAdministrationConfigurationFabricTransportZonesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) CreateTransportZone(transportZoneParam model.TransportZone) (model.TransportZone, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesCreateTransportZoneInputType(), typeConverter)
	sv.AddStructField("TransportZone", transportZoneParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportZone
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesCreateTransportZoneRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "create_transport_zone", inputDataValue, executionContext)
	var emptyOutput model.TransportZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesCreateTransportZoneOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) DeleteTransportZone(zoneIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesDeleteTransportZoneInputType(), typeConverter)
	sv.AddStructField("ZoneId", zoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesDeleteTransportZoneRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "delete_transport_zone", inputDataValue, executionContext)
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

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) GetAllTransportZoneStatus() (model.HeatMapTransportNodesAggregateStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesGetAllTransportZoneStatusInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HeatMapTransportNodesAggregateStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesGetAllTransportZoneStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "get_all_transport_zone_status", inputDataValue, executionContext)
	var emptyOutput model.HeatMapTransportNodesAggregateStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesGetAllTransportZoneStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HeatMapTransportNodesAggregateStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) GetHeatmapTransportZoneStatus(zoneIdParam string, sourceParam *string) (model.HeatMapTransportZoneStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesGetHeatmapTransportZoneStatusInputType(), typeConverter)
	sv.AddStructField("ZoneId", zoneIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.HeatMapTransportZoneStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesGetHeatmapTransportZoneStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "get_heatmap_transport_zone_status", inputDataValue, executionContext)
	var emptyOutput model.HeatMapTransportZoneStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesGetHeatmapTransportZoneStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.HeatMapTransportZoneStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) GetTransportZone(zoneIdParam string) (model.TransportZone, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesGetTransportZoneInputType(), typeConverter)
	sv.AddStructField("ZoneId", zoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportZone
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesGetTransportZoneRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "get_transport_zone", inputDataValue, executionContext)
	var emptyOutput model.TransportZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesGetTransportZoneOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) GetTransportZoneStatus(zoneIdParam string) (model.TransportZoneStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesGetTransportZoneStatusInputType(), typeConverter)
	sv.AddStructField("ZoneId", zoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportZoneStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesGetTransportZoneStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "get_transport_zone_status", inputDataValue, executionContext)
	var emptyOutput model.TransportZoneStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesGetTransportZoneStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportZoneStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) ListTransportZones(cursorParam *string, includeSystemOwnedParam *bool, includedFieldsParam *string, isDefaultParam *bool, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, transportTypeParam *string, uplinkTeamingPolicyNameParam *string) (model.TransportZoneListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesListTransportZonesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeSystemOwned", includeSystemOwnedParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("IsDefault", isDefaultParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("TransportType", transportTypeParam)
	sv.AddStructField("UplinkTeamingPolicyName", uplinkTeamingPolicyNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportZoneListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesListTransportZonesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "list_transport_zones", inputDataValue, executionContext)
	var emptyOutput model.TransportZoneListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesListTransportZonesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportZoneListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemAdministrationConfigurationFabricTransportZonesClient) UpdateTransportZone(zoneIdParam string, transportZoneParam model.TransportZone) (model.TransportZone, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemAdministrationConfigurationFabricTransportZonesUpdateTransportZoneInputType(), typeConverter)
	sv.AddStructField("ZoneId", zoneIdParam)
	sv.AddStructField("TransportZone", transportZoneParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportZone
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemAdministrationConfigurationFabricTransportZonesUpdateTransportZoneRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.system_administration_configuration_fabric_transport_zones", "update_transport_zone", inputDataValue, executionContext)
	var emptyOutput model.TransportZone
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemAdministrationConfigurationFabricTransportZonesUpdateTransportZoneOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportZone), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
