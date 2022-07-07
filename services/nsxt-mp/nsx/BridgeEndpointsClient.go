// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: BridgeEndpoints
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

type BridgeEndpointsClient interface {

	// Creates a Bridge Endpoint. It describes the physical attributes of the bridge like vlan. A logical port can be attached to a vif providing bridging functionality from the logical overlay network to the physical vlan network
	//
	// @param bridgeEndpointParam (required)
	// @return com.vmware.nsx.model.BridgeEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error)

	// Deletes the specified Bridge Endpoint.
	//
	// @param bridgeendpointIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(bridgeendpointIdParam string) error

	// Returns information about a specified bridge endpoint.
	//
	// @param bridgeendpointIdParam (required)
	// @return com.vmware.nsx.model.BridgeEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(bridgeendpointIdParam string) (model.BridgeEndpoint, error)

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
	// @return com.vmware.nsx.model.BridgeEndpointListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(bridgeClusterIdParam *string, bridgeEndpointProfileIdParam *string, cursorParam *string, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vlanTransportZoneIdParam *string) (model.BridgeEndpointListResult, error)

	// Modifies a existing bridge endpoint.
	//
	// @param bridgeendpointIdParam (required)
	// @param bridgeEndpointParam (required)
	// @return com.vmware.nsx.model.BridgeEndpoint
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(bridgeendpointIdParam string, bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error)
}

type bridgeEndpointsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewBridgeEndpointsClient(connector client.Connector) *bridgeEndpointsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.bridge_endpoints")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	bIface := bridgeEndpointsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *bridgeEndpointsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *bridgeEndpointsClient) Create(bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointsCreateInputType(), typeConverter)
	sv.AddStructField("BridgeEndpoint", bridgeEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoints", "create", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bridgeEndpointsClient) Delete(bridgeendpointIdParam string) error {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointsDeleteInputType(), typeConverter)
	sv.AddStructField("BridgeendpointId", bridgeendpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoints", "delete", inputDataValue, executionContext)
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

func (bIface *bridgeEndpointsClient) Get(bridgeendpointIdParam string) (model.BridgeEndpoint, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointsGetInputType(), typeConverter)
	sv.AddStructField("BridgeendpointId", bridgeendpointIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoints", "get", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bridgeEndpointsClient) List(bridgeClusterIdParam *string, bridgeEndpointProfileIdParam *string, cursorParam *string, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, vlanTransportZoneIdParam *string) (model.BridgeEndpointListResult, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointsListInputType(), typeConverter)
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
	operationRestMetaData := bridgeEndpointsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoints", "list", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpointListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpointListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bridgeEndpointsClient) Update(bridgeendpointIdParam string, bridgeEndpointParam model.BridgeEndpoint) (model.BridgeEndpoint, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(bridgeEndpointsUpdateInputType(), typeConverter)
	sv.AddStructField("BridgeendpointId", bridgeendpointIdParam)
	sv.AddStructField("BridgeEndpoint", bridgeEndpointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BridgeEndpoint
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := bridgeEndpointsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx.bridge_endpoints", "update", inputDataValue, executionContext)
	var emptyOutput model.BridgeEndpoint
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), bridgeEndpointsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BridgeEndpoint), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
