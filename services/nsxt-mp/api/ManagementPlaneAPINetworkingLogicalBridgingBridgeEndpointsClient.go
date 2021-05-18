// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpoints
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

type ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient interface {

	// Creates a Bridge Endpoint. It describes the physical attributes of the bridge like vlan. A logical port can be attached to a vif providing bridging functionality from the logical overlay network to the physical vlan network
	//
	// @param bridgeEndpointParam (required)
	// @return com.vmware.model.BridgeEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateBridgeEndpoint(bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error)

	// Deletes the specified Bridge Endpoint.
	//
	// @param bridgeendpointIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteBridgeEndpoint(bridgeendpointIdParam string) error

	// Returns information about a specified bridge endpoint.
	//
	// @param bridgeendpointIdParam (required)
	// @return com.vmware.model.BridgeEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBridgeEndpoint(bridgeendpointIdParam string) (model.BridgeEndpoint, error)

	// Get the statistics for the Bridge Endpoint of the given Endpoint id (endpoint-id)
	//
	// @param endpointIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.BridgeEndpointStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBridgeEndpointStatistics(endpointIdParam string, sourceParam *string) (model.BridgeEndpointStatistics, error)

	// Get the status for the Bridge Endpoint of the given Endpoint id
	//
	// @param endpointIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.BridgeEndpointStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBridgeEndpointStatus(endpointIdParam string, sourceParam *string) (model.BridgeEndpointStatus, error)

	// Returns information about all configured bridge endoints
	//
	// @param bridgeClusterIdParam Bridge Cluster Identifier (optional)
	// @param bridgeEndpointProfileIdParam Bridge endpoint profile used by the edge cluster (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param logicalSwitchIdParam Logical Switch Identifier (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param vlanTransportZoneIdParam VLAN transport zone id used by the edge cluster (optional)
	// @return com.vmware.model.BridgeEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListBridgeEndpoints(bridgeClusterIdParam *string, bridgeEndpointProfileIdParam *string, cursorParam *string, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vlanTransportZoneIdParam *string) (model.BridgeEndpointListResult, error)

	// Modifies a existing bridge endpoint.
	//
	// @param bridgeendpointIdParam (required)
	// @param bridgeEndpointParam (required)
	// @return com.vmware.model.BridgeEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateBridgeEndpoint(bridgeendpointIdParam string, bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error)
}

type managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient(connector client.Connector) *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_bridge_endpoint":         core.NewMethodIdentifier(interfaceIdentifier, "create_bridge_endpoint"),
		"delete_bridge_endpoint":         core.NewMethodIdentifier(interfaceIdentifier, "delete_bridge_endpoint"),
		"get_bridge_endpoint":            core.NewMethodIdentifier(interfaceIdentifier, "get_bridge_endpoint"),
		"get_bridge_endpoint_statistics": core.NewMethodIdentifier(interfaceIdentifier, "get_bridge_endpoint_statistics"),
		"get_bridge_endpoint_status":     core.NewMethodIdentifier(interfaceIdentifier, "get_bridge_endpoint_status"),
		"list_bridge_endpoints":          core.NewMethodIdentifier(interfaceIdentifier, "list_bridge_endpoints"),
		"update_bridge_endpoint":         core.NewMethodIdentifier(interfaceIdentifier, "update_bridge_endpoint"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) CreateBridgeEndpoint(bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsCreateBridgeEndpointInputType(), typeConverter)
	sv.AddStructField("BridgeEndpoint", bridgeEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsCreateBridgeEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "create_bridge_endpoint", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsCreateBridgeEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) DeleteBridgeEndpoint(bridgeendpointIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsDeleteBridgeEndpointInputType(), typeConverter)
	sv.AddStructField("BridgeendpointId", bridgeendpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsDeleteBridgeEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "delete_bridge_endpoint", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) GetBridgeEndpoint(bridgeendpointIdParam string) (model.BridgeEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointInputType(), typeConverter)
	sv.AddStructField("BridgeendpointId", bridgeendpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "get_bridge_endpoint", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) GetBridgeEndpointStatistics(endpointIdParam string, sourceParam *string) (model.BridgeEndpointStatistics, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatisticsInputType(), typeConverter)
	sv.AddStructField("EndpointId", endpointIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "get_bridge_endpoint_statistics", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) GetBridgeEndpointStatus(endpointIdParam string, sourceParam *string) (model.BridgeEndpointStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatusInputType(), typeConverter)
	sv.AddStructField("EndpointId", endpointIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "get_bridge_endpoint_status", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsGetBridgeEndpointStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) ListBridgeEndpoints(bridgeClusterIdParam *string, bridgeEndpointProfileIdParam *string, cursorParam *string, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vlanTransportZoneIdParam *string) (model.BridgeEndpointListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsListBridgeEndpointsInputType(), typeConverter)
	sv.AddStructField("BridgeClusterId", bridgeClusterIdParam)
	sv.AddStructField("BridgeEndpointProfileId", bridgeEndpointProfileIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("LogicalSwitchId", logicalSwitchIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("VlanTransportZoneId", vlanTransportZoneIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsListBridgeEndpointsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "list_bridge_endpoints", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsListBridgeEndpointsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsClient) UpdateBridgeEndpoint(bridgeendpointIdParam string, bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsUpdateBridgeEndpointInputType(), typeConverter)
	sv.AddStructField("BridgeendpointId", bridgeendpointIdParam)
	sv.AddStructField("BridgeEndpoint", bridgeEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsUpdateBridgeEndpointRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoints", "update_bridge_endpoint", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointsUpdateBridgeEndpointOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
