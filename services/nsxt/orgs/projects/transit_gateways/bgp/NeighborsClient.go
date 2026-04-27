// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Neighbors
// Used by client-side stubs.

package bgp

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type NeighborsClient interface {

	// Delete a BGP neighbor configuration from the specified transit gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param neighborIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string) error

	// Read a BGP neighbor configuration under the specified transit gateway by ID.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.nsx_policy.model.BgpNeighborConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string) (nsx_policyModel.BgpNeighborConfig, error)

	// List all BGP neighbor configurations under the specified transit gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includeMarkForDeleteObjectsParam If true, resources that are marked for deletion will be included in the results. By default, these resources are not included. (optional, default to false)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.TgwBgpNeighborConfigListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, transitGatewayIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.TgwBgpNeighborConfigListResult, error)

	// Partially update a BGP neighbor configuration under the specified transit gateway. Only the provided fields are updated.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param neighborIdParam (required)
	// @param bgpNeighborConfigParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string, bgpNeighborConfigParam nsx_policyModel.BgpNeighborConfig) error

	// Create or replace a BGP neighbor configuration under the specified transit gateway. Uses the same BgpNeighborConfig schema as tier-0 gateways. For TGW, the source_attachment field is used instead of source_addresses to specify the centralized network attachment as the BGP peering source.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param neighborIdParam (required)
	// @param bgpNeighborConfigParam (required)
	// @return com.vmware.nsx_policy.model.BgpNeighborConfig
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string, bgpNeighborConfigParam nsx_policyModel.BgpNeighborConfig) (nsx_policyModel.BgpNeighborConfig, error)
}

type neighborsClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewNeighborsClient(connector vapiProtocolClient_.Connector) *neighborsClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp.neighbors")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	nIface := neighborsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *neighborsClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *neighborsClient) Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string) error {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := neighborsDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(neighborsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp.neighbors", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *neighborsClient) Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string) (nsx_policyModel.BgpNeighborConfig, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := neighborsGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(neighborsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.BgpNeighborConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp.neighbors", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.BgpNeighborConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), NeighborsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.BgpNeighborConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) List(orgIdParam string, projectIdParam string, transitGatewayIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.TgwBgpNeighborConfigListResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := neighborsListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(neighborsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludeMarkForDeleteObjects", includeMarkForDeleteObjectsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TgwBgpNeighborConfigListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp.neighbors", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TgwBgpNeighborConfigListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), NeighborsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TgwBgpNeighborConfigListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *neighborsClient) Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string, bgpNeighborConfigParam nsx_policyModel.BgpNeighborConfig) error {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := neighborsPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(neighborsPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	sv.AddStructField("BgpNeighborConfig", bgpNeighborConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp.neighbors", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (nIface *neighborsClient) Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, neighborIdParam string, bgpNeighborConfigParam nsx_policyModel.BgpNeighborConfig) (nsx_policyModel.BgpNeighborConfig, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	operationRestMetaData := neighborsUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(neighborsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	sv.AddStructField("BgpNeighborConfig", bgpNeighborConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.BgpNeighborConfig
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bgp.neighbors", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.BgpNeighborConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), NeighborsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.BgpNeighborConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
