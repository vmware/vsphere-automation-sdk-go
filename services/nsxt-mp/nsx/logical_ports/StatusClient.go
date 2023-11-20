// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Status
// Used by client-side stubs.

package logical_ports

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type StatusClient interface {

	// Returns operational status of a specified logical port.
	//
	//  This api is deprecated from 3.2.2. Please use policy api -
	//  /infra/segments/<segment-id>/ports/<segment-port-id>/status
	//
	// Deprecated: This API element is deprecated.
	//
	// @param lportIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.nsx.model.LogicalPortOperationalStatus
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(lportIdParam string, sourceParam *string) (nsxModel.LogicalPortOperationalStatus, error)

	// Returns operational status of all logical ports. The query parameter \"source=realtime\" is not supported. Pagination is not supported for this API. The query parameters \"cursor\", \"sort_ascending\", \"sort_by\", \"page_size\" and \"included_fields\" will be ignored.
	//
	//  This api is deprecated from 3.2.2. Please use policy api -
	//  /search/query?query=resource_type:SegmentPort&included_fields=id&included_fields=admin_state
	//
	// Deprecated: This API element is deprecated.
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
	// @return com.vmware.nsx.model.LogicalPortStatusSummary
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Getall(attachmentIdParam *string, attachmentTypeParam *string, bridgeClusterIdParam *string, containerPortsOnlyParam *bool, cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, parentVifIdParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchingProfileIdParam *string, transportNodeIdParam *string, transportZoneIdParam *string) (nsxModel.LogicalPortStatusSummary, error)
}

type statusClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewStatusClient(connector vapiProtocolClient_.Connector) *statusClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx.logical_ports.status")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"getall": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "getall"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	sIface := statusClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *statusClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *statusClient) Get(lportIdParam string, sourceParam *string) (nsxModel.LogicalPortOperationalStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statusGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statusGetInputType(), typeConverter)
	sv.AddStructField("LportId", lportIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsxModel.LogicalPortOperationalStatus
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_ports.status", "get", inputDataValue, executionContext)
	var emptyOutput nsxModel.LogicalPortOperationalStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatusGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LogicalPortOperationalStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *statusClient) Getall(attachmentIdParam *string, attachmentTypeParam *string, bridgeClusterIdParam *string, containerPortsOnlyParam *bool, cursorParam *string, diagnosticParam *bool, includedFieldsParam *string, logicalSwitchIdParam *string, pageSizeParam *int64, parentVifIdParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, switchingProfileIdParam *string, transportNodeIdParam *string, transportZoneIdParam *string) (nsxModel.LogicalPortStatusSummary, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	operationRestMetaData := statusGetallRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(statusGetallInputType(), typeConverter)
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
		var emptyOutput nsxModel.LogicalPortStatusSummary
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_ports.status", "getall", inputDataValue, executionContext)
	var emptyOutput nsxModel.LogicalPortStatusSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), StatusGetallOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsxModel.LogicalPortStatusSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
