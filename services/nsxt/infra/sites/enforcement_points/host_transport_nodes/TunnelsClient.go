// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Tunnels
// Used by client-side stubs.

package host_transport_nodes

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type TunnelsClient interface {

	// Tunnel properties
	//
	// @param siteIdParam site ID (required)
	// @param enforcementPointIdParam enforcement point ID (required)
	// @param nodeIdParam ID of transport node (required)
	// @param tunnelNameParam Tunnel name (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx_policy.model.TunnelProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(siteIdParam string, enforcementPointIdParam string, nodeIdParam string, tunnelNameParam string, sourceParam *string) (model.TunnelProperties, error)

	// List of tunnels
	//
	// @param siteIdParam site ID (required)
	// @param enforcementPointIdParam enforcement point ID (required)
	// @param nodeIdParam ID of transport node (required)
	// @param bfdDiagnosticCodeParam BFD diagnostic code of Tunnel as defined in RFC 5880 (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param remoteNodeIdParam (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param statusParam Tunnel status (optional)
	// @return com.vmware.nsx_policy.model.TunnelList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(siteIdParam string, enforcementPointIdParam string, nodeIdParam string, bfdDiagnosticCodeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, remoteNodeIdParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, statusParam *string) (model.TunnelList, error)
}

type tunnelsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTunnelsClient(connector client.Connector) *tunnelsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.tunnels")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":  core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := tunnelsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *tunnelsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *tunnelsClient) Get(siteIdParam string, enforcementPointIdParam string, nodeIdParam string, tunnelNameParam string, sourceParam *string) (model.TunnelProperties, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tunnelsGetInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementPointId", enforcementPointIdParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("TunnelName", tunnelNameParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TunnelProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tunnelsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.tunnels", "get", inputDataValue, executionContext)
	var emptyOutput model.TunnelProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tunnelsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TunnelProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *tunnelsClient) List(siteIdParam string, enforcementPointIdParam string, nodeIdParam string, bfdDiagnosticCodeParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, remoteNodeIdParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, statusParam *string) (model.TunnelList, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(tunnelsListInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementPointId", enforcementPointIdParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("BfdDiagnosticCode", bfdDiagnosticCodeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RemoteNodeId", remoteNodeIdParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("Status", statusParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TunnelList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tunnelsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.tunnels", "list", inputDataValue, executionContext)
	var emptyOutput model.TunnelList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tunnelsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TunnelList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
