// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: BridgeEndpointProfiles
// Used by client-side stubs.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type BridgeEndpointProfilesClient interface {

	// Creates a Bridge Endpoint Profile. Profile contains edge cluster id, indexes of the member nodes, fialover mode and high availability mode for a Bridge EndPoint
	//
	// @param bridgeEndpointProfileParam (required)
	// @return com.vmware.nsx.model.BridgeEndpointProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error)

	// Deletes the specified Bridge Endpoint Profile.
	//
	// @param bridgeendpointprofileIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(bridgeendpointprofileIdParam string) error

	// Returns information about a specified bridge endpoint profile.
	//
	// @param bridgeendpointprofileIdParam (required)
	// @return com.vmware.nsx.model.BridgeEndpointProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(bridgeendpointprofileIdParam string) (model.BridgeEndpointProfile, error)

	// Returns information about all configured bridge endoint profiles
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param edgeClusterIdParam Edge Cluster Identifier (optional)
	// @param failoverModeParam (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.BridgeEndpointProfileListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, edgeClusterIdParam *string, failoverModeParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BridgeEndpointProfileListResult, error)

	// Modifies a existing bridge endpoint profile.
	//
	// @param bridgeendpointprofileIdParam (required)
	// @param bridgeEndpointProfileParam (required)
	// @return com.vmware.nsx.model.BridgeEndpointProfile
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(bridgeendpointprofileIdParam string, bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error)
}

type bridgeEndpointProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewBridgeEndpointProfilesClient(connector client.Connector) *bridgeEndpointProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.bridge_endpoint_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	bIface := bridgeEndpointProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *bridgeEndpointProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *bridgeEndpointProfilesClient) Create(bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointProfilesCreateInputType(), typeConverter)
	sv.AddStructField("BridgeEndpointProfile", bridgeEndpointProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointProfilesCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoint_profiles", "create", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointProfilesCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bridgeEndpointProfilesClient) Delete(bridgeendpointprofileIdParam string) error {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointProfilesDeleteInputType(), typeConverter)
	sv.AddStructField("BridgeendpointprofileId", bridgeendpointprofileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointProfilesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoint_profiles", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (bIface *bridgeEndpointProfilesClient) Get(bridgeendpointprofileIdParam string) (model.BridgeEndpointProfile, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointProfilesGetInputType(), typeConverter)
	sv.AddStructField("BridgeendpointprofileId", bridgeendpointprofileIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointProfilesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoint_profiles", "get", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bridgeEndpointProfilesClient) List(cursorParam *string, edgeClusterIdParam *string, failoverModeParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BridgeEndpointProfileListResult, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointProfilesListInputType(), typeConverter)
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
	operationRestMetaData := bridgeEndpointProfilesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoint_profiles", "list", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfileListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointProfilesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfileListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bridgeEndpointProfilesClient) Update(bridgeendpointprofileIdParam string, bridgeEndpointProfileParam model.BridgeEndpointProfile) (model.BridgeEndpointProfile, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointProfilesUpdateInputType(), typeConverter)
	sv.AddStructField("BridgeendpointprofileId", bridgeendpointprofileIdParam)
	sv.AddStructField("BridgeEndpointProfile", bridgeEndpointProfileParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpointProfile
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointProfilesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoint_profiles", "update", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointProfile
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointProfilesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointProfile), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
