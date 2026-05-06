// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TransportNodeStatus
// Used by client-side stubs.

package transport_zones

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type TransportNodeStatusClient interface {

	// To improve performance, we have removed 'memory_used', 'memory_total', and 'cpu_usage' from the resource_usage field, and 'last_status_update_time' from the health_metrics field in the agent_status returned by the API. To obtain these metrics, use the GET /api/v1/systemhealth-extended/host-transport-nodes endpoint.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeDfwHeapStatsParam If true, DFW heap stats information will be returned in API (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param nodeTypeParam Transport node type. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param statusParam Rolled-up status of pNIC, management connection, control connection, tunnel status and agent status. UP means all of these are up; DOWN represents the state when pNIC or agent status is down. DEGRADED status here represents the state for a node when its pNIC bond status is DEGRADED, or, its Control connection status is either DEGRADED or DOWN. UNKNOWN is the case when both control connection, tunnel and agent status are unknown. If none of these conditions are true, the node status is considered DOWN. (optional)
	// @return com.vmware.nsx.model.TransportNodeStatusListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Getall(cursorParam *string, includeDfwHeapStatsParam *bool, includedFieldsParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, statusParam *string) (nsxModel.TransportNodeStatusListResult, error)

	// Read status of transport nodes in a transport zone
	//
	// @param zoneIdParam ID of transport zone (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeDfwHeapStatsParam If true, DFW heap stats information will be returned in API (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param nodeTypeParam Transport node type. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param statusParam Rolled-up status of pNIC, management connection, control connection, tunnel status and agent status. UP means all of these are up; DOWN represents the state when pNIC or agent status is down. DEGRADED status here represents the state for a node when its pNIC bond status is DEGRADED, or, its Control connection status is either DEGRADED or DOWN. UNKNOWN is the case when both control connection, tunnel and agent status are unknown. If none of these conditions are true, the node status is considered DOWN. (optional)
	// @return com.vmware.nsx.model.TransportNodeStatusListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(zoneIdParam string, cursorParam *string, includeDfwHeapStatsParam *bool, includedFieldsParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, statusParam *string) (nsxModel.TransportNodeStatusListResult, error)
}

type transportNodeStatusClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewTransportNodeStatusClient(connector vapiProtocolClient_.Connector) *transportNodeStatusClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.transport_zones.transport_node_status")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"getall": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "getall"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	tIface := transportNodeStatusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *transportNodeStatusClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *transportNodeStatusClient) Getall(cursorParam *string, includeDfwHeapStatsParam *bool, includedFieldsParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, statusParam *string) (nsxModel.TransportNodeStatusListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeStatusGetallRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeStatusGetallInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeDfwHeapStats", includeDfwHeapStatsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeType", nodeTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("Status", statusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeStatusListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_zones.transport_node_status", "getall", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeStatusGetallOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *transportNodeStatusClient) List(zoneIdParam string, cursorParam *string, includeDfwHeapStatsParam *bool, includedFieldsParam *string, nodeTypeParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, statusParam *string) (nsxModel.TransportNodeStatusListResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	operationRestMetaData := transportNodeStatusListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(transportNodeStatusListInputType(), typeConverter)
	sv.AddStructField("ZoneId", zoneIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeDfwHeapStats", includeDfwHeapStatsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeType", nodeTypeParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("Status", statusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.TransportNodeStatusListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.transport_zones.transport_node_status", "list", inputDataValue, executionContext)
	var emptyOutput nsxModel.TransportNodeStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), TransportNodeStatusListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.TransportNodeStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
