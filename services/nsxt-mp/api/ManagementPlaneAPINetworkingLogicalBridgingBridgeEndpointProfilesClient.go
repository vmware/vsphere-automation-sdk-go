// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfiles
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

type ManagementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient interface {

	// Creates a Bridge Endpoint Profile. Profile contains edge cluster id, indexes of the member nodes, fialover mode and high availability mode for a Bridge EndPoint
	//
	// @param bridgeEndpointProfileParam (required)
	// @return com.vmware.model.BridgeEndpointProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateBridgeEndpointProfile(bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error)

	// Deletes the specified Bridge Endpoint Profile.
	//
	// @param bridgeendpointprofileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteBridgeEndpointProfile(bridgeendpointprofileIdParam string) error

	// Returns information about a specified bridge endpoint profile.
	//
	// @param bridgeendpointprofileIdParam (required)
	// @return com.vmware.model.BridgeEndpointProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBridgeEndpointProfile(bridgeendpointprofileIdParam string) (model.BridgeEndpointProfile, error)

	// Returns information about all configured bridge endoint profiles
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param edgeClusterIdParam Edge Cluster Identifier (optional)
	// @param failoverModeParam (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.BridgeEndpointProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListBridgeEndpointProfiles(cursorParam *string, edgeClusterIdParam *string, failoverModeParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BridgeEndpointProfileListResult, error)

	// Modifies a existing bridge endpoint profile.
	//
	// @param bridgeendpointprofileIdParam (required)
	// @param bridgeEndpointProfileParam (required)
	// @return com.vmware.model.BridgeEndpointProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateBridgeEndpointProfile(bridgeendpointprofileIdParam string, bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error)
}

type managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient(connector client.Connector) *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoint_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_bridge_endpoint_profile": core.NewMethodIdentifier(interfaceIdentifier, "create_bridge_endpoint_profile"),
		"delete_bridge_endpoint_profile": core.NewMethodIdentifier(interfaceIdentifier, "delete_bridge_endpoint_profile"),
		"get_bridge_endpoint_profile":    core.NewMethodIdentifier(interfaceIdentifier, "get_bridge_endpoint_profile"),
		"list_bridge_endpoint_profiles":  core.NewMethodIdentifier(interfaceIdentifier, "list_bridge_endpoint_profiles"),
		"update_bridge_endpoint_profile": core.NewMethodIdentifier(interfaceIdentifier, "update_bridge_endpoint_profile"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient) CreateBridgeEndpointProfile(bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesCreateBridgeEndpointProfileInputType(), typeConverter)
	sv.AddStructField("BridgeEndpointProfile", bridgeEndpointProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesCreateBridgeEndpointProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoint_profiles", "create_bridge_endpoint_profile", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesCreateBridgeEndpointProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient) DeleteBridgeEndpointProfile(bridgeendpointprofileIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesDeleteBridgeEndpointProfileInputType(), typeConverter)
	sv.AddStructField("BridgeendpointprofileId", bridgeendpointprofileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesDeleteBridgeEndpointProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoint_profiles", "delete_bridge_endpoint_profile", inputDataValue, executionContext)
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

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient) GetBridgeEndpointProfile(bridgeendpointprofileIdParam string) (model.BridgeEndpointProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesGetBridgeEndpointProfileInputType(), typeConverter)
	sv.AddStructField("BridgeendpointprofileId", bridgeendpointprofileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesGetBridgeEndpointProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoint_profiles", "get_bridge_endpoint_profile", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesGetBridgeEndpointProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient) ListBridgeEndpointProfiles(cursorParam *string, edgeClusterIdParam *string, failoverModeParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BridgeEndpointProfileListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesListBridgeEndpointProfilesInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	sv.AddStructField("FailoverMode", failoverModeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfileListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesListBridgeEndpointProfilesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoint_profiles", "list_bridge_endpoint_profiles", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesListBridgeEndpointProfilesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesClient) UpdateBridgeEndpointProfile(bridgeendpointprofileIdParam string, bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesUpdateBridgeEndpointProfileInputType(), typeConverter)
	sv.AddStructField("BridgeendpointprofileId", bridgeendpointprofileIdParam)
	sv.AddStructField("BridgeEndpointProfile", bridgeEndpointProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesUpdateBridgeEndpointProfileRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_bridging_bridge_endpoint_profiles", "update_bridge_endpoint_profile", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalBridgingBridgeEndpointProfilesUpdateBridgeEndpointProfileOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
