// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitches
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

type ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient interface {

	// Creates a new logical switch. The request must include the transport_zone_id, display_name, and admin_state (UP or DOWN). The replication_mode (MTEP or SOURCE) is required for overlay logical switches, but not for VLAN-based logical switches. A vlan needs to be provided for VLAN-based logical switches
	//
	// @param logicalSwitchParam (required)
	// @return com.vmware.model.LogicalSwitch
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLogicalSwitch(logicalSwitchParam model.LogicalSwitch) (model.LogicalSwitch, error)

	// Removes a logical switch from the associated overlay or VLAN transport zone. By default, a logical switch cannot be deleted if there are logical ports on the switch, or it is added to a NSGroup. Cascade option can be used to delete all ports and the logical switch. Detach option can be used to delete the logical switch forcibly.
	//
	// @param lswitchIdParam (required)
	// @param cascadeParam Delete a Logical Switch and all the logical ports in it, if none of the logical ports have any attachment. (optional, default to false)
	// @param detachParam Force delete a logical switch (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLogicalSwitch(lswitchIdParam string, cascadeParam *bool, detachParam *bool) error

	// Returns information about the specified logical switch Id.
	//
	// @param lswitchIdParam (required)
	// @return com.vmware.model.LogicalSwitch
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitch(lswitchIdParam string) (model.LogicalSwitch, error)

	// Returns MAC table of a specified logical switch from the given transport node if a transport node id is given in the query parameter from the Central Controller Plane. The query parameter \"source=cached\" is not supported.
	//
	// @param lswitchIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam TransportNode Id (optional)
	// @return com.vmware.model.MacAddressListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchMacTable(lswitchIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.MacAddressListResult, error)

	// Returns MAC table of a specified logical switch in CSV format from the given transport node if a transport node id is given in the query parameter from the Central Controller Plane. The query parameter \"source=cached\" is not supported.
	//
	// @param lswitchIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam TransportNode Id (optional)
	// @return com.vmware.model.MacAddressCsvListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchMacTableInCsvFormatCsv(lswitchIdParam string, sourceParam *string, transportNodeIdParam *string) (model.MacAddressCsvListResult, error)

	// Returns current state of the logical switch configuration and details of only out-of-sync transport nodes.
	//
	// @param lswitchIdParam (required)
	// @return com.vmware.model.LogicalSwitchState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchState(lswitchIdParam string) (model.LogicalSwitchState, error)

	// Returns statistics of a specified logical switch. The query parameter \"source=realtime\" is not supported.
	//
	// @param lswitchIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalSwitchStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchStatistics(lswitchIdParam string, sourceParam *string) (model.LogicalSwitchStatistics, error)

	// Returns the number of ports assigned to a logical switch.
	//
	// @param lswitchIdParam (required)
	// @return com.vmware.model.LogicalSwitchStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchStatus(lswitchIdParam string) (model.LogicalSwitchStatus, error)

	// Returns Operational status of all logical switches. The query parameter \"source=realtime\" is not supported.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param diagnosticParam Flag to enable showing of transit logical switch. (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param switchingProfileIdParam Switching Profile identifier (optional)
	// @param transportTypeParam Mode of transport supported in the transport zone for this logical switch (optional)
	// @param transportZoneIdParam Transport zone identifier (optional)
	// @param uplinkTeamingPolicyNameParam The logical switch's uplink teaming policy name (optional)
	// @param vlanParam Virtual Local Area Network Identifier (optional)
	// @param vniParam VNI of the OVERLAY LogicalSwitch(es) to return. (optional)
	// @return com.vmware.model.LogicalSwitchStatusSummary
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchStatusSummary(cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchingProfileIdParam *string, transportTypeParam *string, transportZoneIdParam *string, uplinkTeamingPolicyNameParam *string, vlanParam *int64, vniParam *int64) (model.LogicalSwitchStatusSummary, error)

	// Returns the virtual tunnel endpoint table of a specified logical switch from the given transport node if a transport node id is given in the query parameter, from the Central Controller Plane. The query parameter \"source=cached\" is not supported.
	//
	// @param lswitchIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam TransportNode Id (optional)
	// @return com.vmware.model.VtepListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchVtepTable(lswitchIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.VtepListResult, error)

	// Returns virtual tunnel endpoint table of a specified logical switch in CSV format from the given transport node if a transport node id is given in the query parameter from the Central Controller Plane. The query parameter \"source=cached\" is not supported.
	//
	// @param lswitchIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam TransportNode Id (optional)
	// @return com.vmware.model.VtepCsvListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalSwitchVtepTableInCsvFormatCsv(lswitchIdParam string, sourceParam *string, transportNodeIdParam *string) (model.VtepCsvListResult, error)

	// Returns information about all configured logical switches.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param diagnosticParam Flag to enable showing of transit logical switch. (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param switchingProfileIdParam Switching Profile identifier (optional)
	// @param transportTypeParam Mode of transport supported in the transport zone for this logical switch (optional)
	// @param transportZoneIdParam Transport zone identifier (optional)
	// @param uplinkTeamingPolicyNameParam The logical switch's uplink teaming policy name (optional)
	// @param vlanParam Virtual Local Area Network Identifier (optional)
	// @param vniParam VNI of the OVERLAY LogicalSwitch(es) to return. (optional)
	// @return com.vmware.model.LogicalSwitchListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLogicalSwitches(cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, switchingProfileIdParam *string, transportTypeParam *string, transportZoneIdParam *string, uplinkTeamingPolicyNameParam *string, vlanParam *int64, vniParam *int64) (model.LogicalSwitchListResult, error)

	// Returns a list of logical switches states that have realized state as provided as query parameter.
	//
	// @param statusParam Realized state of logical switches (optional)
	// @return com.vmware.model.LogicalSwitchStateListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLogicalSwitchesByState(statusParam *string) (model.LogicalSwitchStateListResult, error)

	// Modifies attributes of an existing logical switch. Modifiable attributes include admin_state, replication_mode, switching_profile_ids and VLAN spec. You cannot modify the original transport_zone_id.
	//
	// @param lswitchIdParam (required)
	// @param logicalSwitchParam (required)
	// @return com.vmware.model.LogicalSwitch
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLogicalSwitch(lswitchIdParam string, logicalSwitchParam model.LogicalSwitch) (model.LogicalSwitch, error)
}

type managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient(connector client.Connector) *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_logical_switch":                           core.NewMethodIdentifier(interfaceIdentifier, "create_logical_switch"),
		"delete_logical_switch":                           core.NewMethodIdentifier(interfaceIdentifier, "delete_logical_switch"),
		"get_logical_switch":                              core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch"),
		"get_logical_switch_mac_table":                    core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_mac_table"),
		"get_logical_switch_mac_table_in_csv_format_csv":  core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_mac_table_in_csv_format_csv"),
		"get_logical_switch_state":                        core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_state"),
		"get_logical_switch_statistics":                   core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_statistics"),
		"get_logical_switch_status":                       core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_status"),
		"get_logical_switch_status_summary":               core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_status_summary"),
		"get_logical_switch_vtep_table":                   core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_vtep_table"),
		"get_logical_switch_vtep_table_in_csv_format_csv": core.NewMethodIdentifier(interfaceIdentifier, "get_logical_switch_vtep_table_in_csv_format_csv"),
		"list_logical_switches":                           core.NewMethodIdentifier(interfaceIdentifier, "list_logical_switches"),
		"list_logical_switches_by_state":                  core.NewMethodIdentifier(interfaceIdentifier, "list_logical_switches_by_state"),
		"update_logical_switch":                           core.NewMethodIdentifier(interfaceIdentifier, "update_logical_switch"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) CreateLogicalSwitch(logicalSwitchParam model.LogicalSwitch) (model.LogicalSwitch, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesCreateLogicalSwitchInputType(), typeConverter)
	sv.AddStructField("LogicalSwitch", logicalSwitchParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitch
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesCreateLogicalSwitchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "create_logical_switch", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitch
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesCreateLogicalSwitchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitch), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) DeleteLogicalSwitch(lswitchIdParam string, cascadeParam *bool, detachParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesDeleteLogicalSwitchInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("Cascade", cascadeParam)
	sv.AddStructField("Detach", detachParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesDeleteLogicalSwitchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "delete_logical_switch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitch(lswitchIdParam string) (model.LogicalSwitch, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitch
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitch
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitch), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchMacTable(lswitchIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.MacAddressListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MacAddressListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_mac_table", inputDataValue, executionContext)
	var emptyOutput model.MacAddressListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MacAddressListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchMacTableInCsvFormatCsv(lswitchIdParam string, sourceParam *string, transportNodeIdParam *string) (model.MacAddressCsvListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MacAddressCsvListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_mac_table_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.MacAddressCsvListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchMacTableInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MacAddressCsvListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchState(lswitchIdParam string) (model.LogicalSwitchState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStateInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_state", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchStatistics(lswitchIdParam string, sourceParam *string) (model.LogicalSwitchStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatisticsInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_statistics", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchStatus(lswitchIdParam string) (model.LogicalSwitchStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_status", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchStatusSummary(cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchingProfileIdParam *string, transportTypeParam *string, transportZoneIdParam *string, uplinkTeamingPolicyNameParam *string, vlanParam *int64, vniParam *int64) (model.LogicalSwitchStatusSummary, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusSummaryInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Diagnostic", diagnosticParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("TransportType", transportTypeParam)
	sv.AddStructField("TransportZoneId", transportZoneIdParam)
	sv.AddStructField("UplinkTeamingPolicyName", uplinkTeamingPolicyNameParam)
	sv.AddStructField("Vlan", vlanParam)
	sv.AddStructField("Vni", vniParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchStatusSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusSummaryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_status_summary", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchStatusSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchStatusSummaryOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchStatusSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchVtepTable(lswitchIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.VtepListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VtepListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_vtep_table", inputDataValue, executionContext)
	var emptyOutput model.VtepListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VtepListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) GetLogicalSwitchVtepTableInCsvFormatCsv(lswitchIdParam string, sourceParam *string, transportNodeIdParam *string) (model.VtepCsvListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.VtepCsvListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "get_logical_switch_vtep_table_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.VtepCsvListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesGetLogicalSwitchVtepTableInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.VtepCsvListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) ListLogicalSwitches(cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, switchingProfileIdParam *string, transportTypeParam *string, transportZoneIdParam *string, uplinkTeamingPolicyNameParam *string, vlanParam *int64, vniParam *int64) (model.LogicalSwitchListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Diagnostic", diagnosticParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("TransportType", transportTypeParam)
	sv.AddStructField("TransportZoneId", transportZoneIdParam)
	sv.AddStructField("UplinkTeamingPolicyName", uplinkTeamingPolicyNameParam)
	sv.AddStructField("Vlan", vlanParam)
	sv.AddStructField("Vni", vniParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "list_logical_switches", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) ListLogicalSwitchesByState(statusParam *string) (model.LogicalSwitchStateListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesByStateInputType(), typeConverter)
	sv.AddStructField("Status", statusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitchStateListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesByStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "list_logical_switches_by_state", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitchStateListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesListLogicalSwitchesByStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitchStateListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesClient) UpdateLogicalSwitch(lswitchIdParam string, logicalSwitchParam model.LogicalSwitch) (model.LogicalSwitch, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesUpdateLogicalSwitchInputType(), typeConverter)
	sv.AddStructField("LswitchId", lswitchIdParam)
	sv.AddStructField("LogicalSwitch", logicalSwitchParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalSwitch
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesUpdateLogicalSwitchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switches", "update_logical_switch", inputDataValue, executionContext)
	var emptyOutput model.LogicalSwitch
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchesUpdateLogicalSwitchOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalSwitch), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
