// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts
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

type ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient interface {

	// Creates a new logical switch port. The required parameters are the associated logical_switch_id and admin_state (UP or DOWN). Optional parameters are the attachment and switching_profile_ids. If you don't specify switching_profile_ids, default switching profiles are assigned to the port. If you don't specify an attachment, the switch port remains empty. To configure an attachment, you must specify an id, and optionally you can specify an attachment_type (VIF or LOGICALROUTER). The attachment_type is VIF by default.
	//
	// @param logicalPortParam (required)
	// @return com.vmware.model.LogicalPort
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLogicalPort(logicalPortParam model.LogicalPort) (model.LogicalPort, error)

	// Creates a new, custom qos, port-mirroring, spoof-guard or port-security switching profile. You can override their default switching profile assignments by creating a new switching profile and assigning it to one or more logical switches. You cannot override the default ipfix or ip_discovery switching profiles.
	//
	// @param baseSwitchingProfileParam (required)
	// The parameter must contain all the properties defined in model.BaseSwitchingProfile.
	// @return com.vmware.model.BaseSwitchingProfile
	// The return value will contain all the properties defined in model.BaseSwitchingProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateSwitchingProfile(baseSwitchingProfileParam *data.StructValue) (*data.StructValue, error)

	// Deletes the specified logical switch port. By default, if logical port has attachments, or it is added to any NSGroup, the deletion will be failed. Option detach could be used for deleting logical port forcibly.
	//
	// @param lportIdParam (required)
	// @param detachParam force delete even if attached or referenced by a group (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLogicalPort(lportIdParam string, detachParam *bool) error

	// Deletes the specified switching profile.
	//
	// @param switchingProfileIdParam (required)
	// @param unbindParam force unbinding of logical switches and ports from a switching profile (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteSwitchingProfile(switchingProfileIdParam string, unbindParam *bool) error

	// Returns information about a specified logical port.
	//
	// @param lportIdParam (required)
	// @return com.vmware.model.LogicalPort
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPort(lportIdParam string) (model.LogicalPort, error)

	// Returns MAC table of a specified logical port. If the target transport node id is not provided, the NSX manager will ask the controller for the transport node where the logical port is located. The query parameter \"source=cached\" is not supported. MAC table retrieval is not supported on logical ports that are attached to a logical router.
	//
	// @param lportIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam TransportNode Id (optional)
	// @return com.vmware.model.LogicalPortMacAddressListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPortMacTable(lportIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.LogicalPortMacAddressListResult, error)

	// Returns MAC table in CSV format of a specified logical port. If the target transport node id is not provided, the NSX manager will ask the controller for the transport node where the logical port is located. The query parameter \"source=cached\" is not supported. MAC table retrieval is not supported on logical ports that are attached to a logical router.
	//
	// @param lportIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam TransportNode Id (optional)
	// @return com.vmware.model.LogicalPortMacAddressCsvListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPortMacTableInCsvFormatCsv(lportIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.LogicalPortMacAddressCsvListResult, error)

	// Returns operational status of a specified logical port.
	//
	// @param lportIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalPortOperationalStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPortOperationalStatus(lportIdParam string, sourceParam *string) (model.LogicalPortOperationalStatus, error)

	// Returns transport node id for a specified logical port. Also returns information about all address bindings of the specified logical port. This includes address bindings discovered via various snooping methods like ARP snooping, DHCP snooping etc. and addressing bindings that are realized based on user configuration.
	//
	// @param lportIdParam (required)
	// @return com.vmware.model.LogicalPortState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPortState(lportIdParam string) (model.LogicalPortState, error)

	// Returns statistics of a specified logical port. If the logical port is attached to a logical router port, query parameter \"source=realtime\" is not supported.
	//
	// @param lportIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalPortStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPortStatistics(lportIdParam string, sourceParam *string) (model.LogicalPortStatistics, error)

	// Returns operational status of all logical ports. The query parameter \"source=realtime\" is not supported. Pagination is not supported for this API. The query parameters \"cursor\", \"sort_ascending\", \"sort_by\", \"page_size\" and \"included_fields\" will be ignored.
	//
	// @param attachmentIdParam Logical Port attachment Id (optional)
	// @param attachmentTypeParam Type of attachment for logical port; for query only. (optional)
	// @param bridgeClusterIdParam Bridge Cluster identifier (optional)
	// @param containerPortsOnlyParam Only container VIF logical ports will be returned if true (optional, default to false)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param diagnosticParam Flag to enable showing of transit logical port. (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param logicalSwitchIdParam Logical Switch identifier (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param parentVifIdParam ID of the VIF of type PARENT (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param switchingProfileIdParam Network Profile identifier (optional)
	// @param transportNodeIdParam Transport node identifier (optional)
	// @param transportZoneIdParam Transport zone identifier (optional)
	// @return com.vmware.model.LogicalPortStatusSummary
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalPortStatusSummary(attachmentIdParam *string, attachmentTypeParam *string, bridgeClusterIdParam *string, containerPortsOnlyParam *bool, cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, parentVifIdParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchingProfileIdParam *string, transportNodeIdParam *string, transportZoneIdParam *string) (model.LogicalPortStatusSummary, error)

	// Returns information about a specified switching profile.
	//
	// @param switchingProfileIdParam (required)
	// @return com.vmware.model.BaseSwitchingProfile
	// The return value will contain all the properties defined in model.BaseSwitchingProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSwitchingProfile(switchingProfileIdParam string) (*data.StructValue, error)

	// Get Counts of Ports and Switches Using This Switching Profile
	//
	// @param switchingProfileIdParam (required)
	// @return com.vmware.model.SwitchingProfileStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetSwitchingProfileStatus(switchingProfileIdParam string) (model.SwitchingProfileStatus, error)

	// Returns information about all configured logical switch ports. Logical switch ports connect to VM virtual network interface cards (NICs). Each logical port is associated with one logical switch.
	//
	// @param attachmentIdParam Logical Port attachment Id (optional)
	// @param attachmentTypeParam Type of attachment for logical port; for query only. (optional)
	// @param bridgeClusterIdParam Bridge Cluster identifier (optional)
	// @param containerPortsOnlyParam Only container VIF logical ports will be returned if true (optional, default to false)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param diagnosticParam Flag to enable showing of transit logical port. (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param logicalSwitchIdParam Logical Switch identifier (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param parentVifIdParam ID of the VIF of type PARENT (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param switchingProfileIdParam Network Profile identifier (optional)
	// @param transportNodeIdParam Transport node identifier (optional)
	// @param transportZoneIdParam Transport zone identifier (optional)
	// @return com.vmware.model.LogicalPortListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLogicalPorts(attachmentIdParam *string, attachmentTypeParam *string, bridgeClusterIdParam *string, containerPortsOnlyParam *bool, cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, parentVifIdParam *string, sortAscendingParam *bool, sortByParam *string, switchingProfileIdParam *string, transportNodeIdParam *string, transportZoneIdParam *string) (model.LogicalPortListResult, error)

	// Returns information about the system-default and user-configured switching profiles. Each switching profile has a unique ID, a display name, and various other read-only and configurable properties. The default switching profiles are assigned automatically to each switch.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeSystemOwnedParam Whether the list result contains system resources (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param switchingProfileTypeParam comma-separated list of switching profile types, e.g. ?switching_profile_type=QosSwitchingProfile,IpDiscoverySwitchingProfile (optional)
	// @return com.vmware.model.SwitchingProfilesListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListSwitchingProfiles(cursorParam *string, includeSystemOwnedParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, switchingProfileTypeParam *string) (model.SwitchingProfilesListResult, error)

	// Modifies an existing logical switch port. Parameters that can be modified include attachment_type (LOGICALROUTER, VIF), admin_state (UP or DOWN), attachment id and switching_profile_ids. You cannot modify the logical_switch_id. In other words, you cannot move an existing port from one switch to another switch.
	//
	// @param lportIdParam (required)
	// @param logicalPortParam (required)
	// @return com.vmware.model.LogicalPort
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLogicalPort(lportIdParam string, logicalPortParam model.LogicalPort) (model.LogicalPort, error)

	// Updates the user-configurable parameters of a switching profile. Only the qos, port-mirroring, spoof-guard and port-security switching profiles can be modified. You cannot modify the ipfix or ip-discovery switching profiles.
	//
	// @param switchingProfileIdParam (required)
	// @param baseSwitchingProfileParam (required)
	// The parameter must contain all the properties defined in model.BaseSwitchingProfile.
	// @return com.vmware.model.BaseSwitchingProfile
	// The return value will contain all the properties defined in model.BaseSwitchingProfile.
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateSwitchingProfile(switchingProfileIdParam string, baseSwitchingProfileParam *data.StructValue) (*data.StructValue, error)
}

type managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient(connector client.Connector) *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_logical_port":                          core.NewMethodIdentifier(interfaceIdentifier, "create_logical_port"),
		"create_switching_profile":                     core.NewMethodIdentifier(interfaceIdentifier, "create_switching_profile"),
		"delete_logical_port":                          core.NewMethodIdentifier(interfaceIdentifier, "delete_logical_port"),
		"delete_switching_profile":                     core.NewMethodIdentifier(interfaceIdentifier, "delete_switching_profile"),
		"get_logical_port":                             core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port"),
		"get_logical_port_mac_table":                   core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port_mac_table"),
		"get_logical_port_mac_table_in_csv_format_csv": core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port_mac_table_in_csv_format_csv"),
		"get_logical_port_operational_status":          core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port_operational_status"),
		"get_logical_port_state":                       core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port_state"),
		"get_logical_port_statistics":                  core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port_statistics"),
		"get_logical_port_status_summary":              core.NewMethodIdentifier(interfaceIdentifier, "get_logical_port_status_summary"),
		"get_switching_profile":                        core.NewMethodIdentifier(interfaceIdentifier, "get_switching_profile"),
		"get_switching_profile_status":                 core.NewMethodIdentifier(interfaceIdentifier, "get_switching_profile_status"),
		"list_logical_ports":                           core.NewMethodIdentifier(interfaceIdentifier, "list_logical_ports"),
		"list_switching_profiles":                      core.NewMethodIdentifier(interfaceIdentifier, "list_switching_profiles"),
		"update_logical_port":                          core.NewMethodIdentifier(interfaceIdentifier, "update_logical_port"),
		"update_switching_profile":                     core.NewMethodIdentifier(interfaceIdentifier, "update_switching_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) CreateLogicalPort(logicalPortParam model.LogicalPort) (model.LogicalPort, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateLogicalPortInputType(), typeConverter)
	sv.AddStructField("LogicalPort", logicalPortParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPort
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateLogicalPortRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "create_logical_port", inputDataValue, executionContext)
	var emptyOutput model.LogicalPort
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateLogicalPortOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPort), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) CreateSwitchingProfile(baseSwitchingProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateSwitchingProfileInputType(), typeConverter)
	sv.AddStructField("BaseSwitchingProfile", baseSwitchingProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateSwitchingProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "create_switching_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateSwitchingProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) DeleteLogicalPort(lportIdParam string, detachParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteLogicalPortInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("Detach", detachParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteLogicalPortRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "delete_logical_port", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) DeleteSwitchingProfile(switchingProfileIdParam string, unbindParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteSwitchingProfileInputType(), typeConverter)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("Unbind", unbindParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteSwitchingProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "delete_switching_profile", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPort(lportIdParam string) (model.LogicalPort, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPort
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port", inputDataValue, executionContext)
	var emptyOutput model.LogicalPort
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPort), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPortMacTable(lportIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.LogicalPortMacAddressListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortMacAddressListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port_mac_table", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortMacAddressListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortMacAddressListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPortMacTableInCsvFormatCsv(lportIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.LogicalPortMacAddressCsvListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortMacAddressCsvListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port_mac_table_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortMacAddressCsvListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortMacAddressCsvListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPortOperationalStatus(lportIdParam string, sourceParam *string) (model.LogicalPortOperationalStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOperationalStatusInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortOperationalStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOperationalStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port_operational_status", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortOperationalStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOperationalStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortOperationalStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPortState(lportIdParam string) (model.LogicalPortState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStateInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port_state", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPortStatistics(lportIdParam string, sourceParam *string) (model.LogicalPortStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatisticsInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port_statistics", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetLogicalPortStatusSummary(attachmentIdParam *string, attachmentTypeParam *string, bridgeClusterIdParam *string, containerPortsOnlyParam *bool, cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, parentVifIdParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchingProfileIdParam *string, transportNodeIdParam *string, transportZoneIdParam *string) (model.LogicalPortStatusSummary, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatusSummaryInputType(), typeConverter)
	sv.AddStructField("AttachmentId", attachmentIdParam)
	sv.AddStructField("AttachmentType", attachmentTypeParam)
	sv.AddStructField("BridgeClusterId", bridgeClusterIdParam)
	sv.AddStructField("ContainerPortsOnly", containerPortsOnlyParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Diagnostic", diagnosticParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ParentVifId", parentVifIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("TransportZoneId", transportZoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortStatusSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatusSummaryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_logical_port_status_summary", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortStatusSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatusSummaryOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortStatusSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetSwitchingProfile(switchingProfileIdParam string) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileInputType(), typeConverter)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_switching_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) GetSwitchingProfileStatus(switchingProfileIdParam string) (model.SwitchingProfileStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileStatusInputType(), typeConverter)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SwitchingProfileStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "get_switching_profile_status", inputDataValue, executionContext)
	var emptyOutput model.SwitchingProfileStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SwitchingProfileStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) ListLogicalPorts(attachmentIdParam *string, attachmentTypeParam *string, bridgeClusterIdParam *string, containerPortsOnlyParam *bool, cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, parentVifIdParam *string, sortAscendingParam *bool, sortByParam *string, switchingProfileIdParam *string, transportNodeIdParam *string, transportZoneIdParam *string) (model.LogicalPortListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListLogicalPortsInputType(), typeConverter)
	sv.AddStructField("AttachmentId", attachmentIdParam)
	sv.AddStructField("AttachmentType", attachmentTypeParam)
	sv.AddStructField("BridgeClusterId", bridgeClusterIdParam)
	sv.AddStructField("ContainerPortsOnly", containerPortsOnlyParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Diagnostic", diagnosticParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("ParentVifId", parentVifIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("TransportZoneId", transportZoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPortListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListLogicalPortsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "list_logical_ports", inputDataValue, executionContext)
	var emptyOutput model.LogicalPortListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListLogicalPortsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPortListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) ListSwitchingProfiles(cursorParam *string, includeSystemOwnedParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, switchingProfileTypeParam *string) (model.SwitchingProfilesListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListSwitchingProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeSystemOwned", includeSystemOwnedParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("SwitchingProfileType", switchingProfileTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SwitchingProfilesListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListSwitchingProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "list_switching_profiles", inputDataValue, executionContext)
	var emptyOutput model.SwitchingProfilesListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListSwitchingProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SwitchingProfilesListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) UpdateLogicalPort(lportIdParam string, logicalPortParam model.LogicalPort) (model.LogicalPort, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateLogicalPortInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("LogicalPort", logicalPortParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalPort
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateLogicalPortRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "update_logical_port", inputDataValue, executionContext)
	var emptyOutput model.LogicalPort
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateLogicalPortOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalPort), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsClient) UpdateSwitchingProfile(switchingProfileIdParam string, baseSwitchingProfileParam *data.StructValue) (*data.StructValue, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateSwitchingProfileInputType(), typeConverter)
	sv.AddStructField("SwitchingProfileId", switchingProfileIdParam)
	sv.AddStructField("BaseSwitchingProfile", baseSwitchingProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateSwitchingProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_switching_logical_switch_ports", "update_switching_profile", inputDataValue, executionContext)
	var emptyOutput *data.StructValue
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateSwitchingProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
