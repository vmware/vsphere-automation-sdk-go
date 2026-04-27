// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: BfdPeers
// Used by client-side stubs.

package transit_gateways

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type BfdPeersClient interface {

	// Delete a BFD peer configuration from the specified transit gateway.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param bfdPeerIdParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string) error

	// Read a BFD peer configuration for a transit gateway by ID.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param bfdPeerIdParam (required)
	// @return com.vmware.nsx_policy.model.TransitGatewayBfdPeer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string) (nsx_policyModel.TransitGatewayBfdPeer, error)

	// List all BFD peer configurations under the specified transit gateway.
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
	// @return com.vmware.nsx_policy.model.TransitGatewayBfdPeerListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, transitGatewayIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.TransitGatewayBfdPeerListResult, error)

	// Partially update a BFD peer configuration for a transit gateway. Only the provided fields are updated.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param bfdPeerIdParam (required)
	// @param transitGatewayBfdPeerParam (required)
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string, transitGatewayBfdPeerParam nsx_policyModel.TransitGatewayBfdPeer) error

	// Create or replace a BFD peer configuration for fast liveness detection of static route next-hop peers on a transit gateway centralized network attachment.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param transitGatewayIdParam (required)
	// @param bfdPeerIdParam (required)
	// @param transitGatewayBfdPeerParam (required)
	// @return com.vmware.nsx_policy.model.TransitGatewayBfdPeer
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string, transitGatewayBfdPeerParam nsx_policyModel.TransitGatewayBfdPeer) (nsx_policyModel.TransitGatewayBfdPeer, error)
}

type bfdPeersClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewBfdPeersClient(connector vapiProtocolClient_.Connector) *bfdPeersClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.transit_gateways.bfd_peers")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"delete": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  vapiCore_.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	bIface := bfdPeersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &bIface
}

func (bIface *bfdPeersClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := bIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (bIface *bfdPeersClient) Delete(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string) error {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bfdPeersDeleteRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bfdPeersDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("BfdPeerId", bfdPeerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bfd_peers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (bIface *bfdPeersClient) Get(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string) (nsx_policyModel.TransitGatewayBfdPeer, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bfdPeersGetRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bfdPeersGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("BfdPeerId", bfdPeerIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TransitGatewayBfdPeer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bfd_peers", "get", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TransitGatewayBfdPeer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BfdPeersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TransitGatewayBfdPeer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bfdPeersClient) List(orgIdParam string, projectIdParam string, transitGatewayIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.TransitGatewayBfdPeerListResult, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bfdPeersListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bfdPeersListInputType(), typeConverter)
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
		var emptyOutput nsx_policyModel.TransitGatewayBfdPeerListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bfd_peers", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TransitGatewayBfdPeerListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BfdPeersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TransitGatewayBfdPeerListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *bfdPeersClient) Patch(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string, transitGatewayBfdPeerParam nsx_policyModel.TransitGatewayBfdPeer) error {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bfdPeersPatchRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bfdPeersPatchInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("BfdPeerId", bfdPeerIdParam)
	sv.AddStructField("TransitGatewayBfdPeer", transitGatewayBfdPeerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bfd_peers", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (bIface *bfdPeersClient) Update(orgIdParam string, projectIdParam string, transitGatewayIdParam string, bfdPeerIdParam string, transitGatewayBfdPeerParam nsx_policyModel.TransitGatewayBfdPeer) (nsx_policyModel.TransitGatewayBfdPeer, error) {
	typeConverter := bIface.connector.TypeConverter()
	executionContext := bIface.connector.NewExecutionContext()
	operationRestMetaData := bfdPeersUpdateRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(bfdPeersUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("TransitGatewayId", transitGatewayIdParam)
	sv.AddStructField("BfdPeerId", bfdPeerIdParam)
	sv.AddStructField("TransitGatewayBfdPeer", transitGatewayBfdPeerParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.TransitGatewayBfdPeer
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := bIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.transit_gateways.bfd_peers", "update", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.TransitGatewayBfdPeer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), BfdPeersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.TransitGatewayBfdPeer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
