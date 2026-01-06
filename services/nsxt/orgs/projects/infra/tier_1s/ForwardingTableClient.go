// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ForwardingTable
// Used by client-side stubs.

package tier_1s

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type ForwardingTableClient interface {

	// Get forwarding table from tier-1
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param tier1IdParam (required)
	// @param componentTypeParam Component type define to take the route from CCP. (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param edgeIdParam UUID of edge node. Edge should be member of enforcement point. (optional)
	// @param edgePathParam Policy path of edge node. Edge should be member of enforcement point. (optional)
	// @param enforcementPointPathParam String Path of the enforcement point. When not specified, routes from all enforcement-points are returned. This property is required for retrieving routes in CSV format. (optional)
	// @param includedFieldsParam Note - this parameter currently only works when used with the search APIs /policy/api/v1/search/query and /policy/api/v1/search/dsl. It is ignored for other list APIs. (optional)
	// @param networkPrefixParam IPAddress or CIDR network address to filter entries in the table. (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param routeSourceParam Filter routes based on the source from which route is learned. (optional)
	// @param sortAscendingParam If true, results are sorted in ascending order (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.RoutingTableListResult
	//
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws TimedOut  Gateway Timeout
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, tier1IdParam string, componentTypeParam *string, cursorParam *string, edgeIdParam *string, edgePathParam *string, enforcementPointPathParam *string, includedFieldsParam *string, networkPrefixParam *string, pageSizeParam *int64, routeSourceParam *string, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.RoutingTableListResult, error)
}

type forwardingTableClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewForwardingTableClient(connector vapiProtocolClient_.Connector) *forwardingTableClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.infra.tier_1s.forwarding_table")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"list": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	fIface := forwardingTableClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *forwardingTableClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *forwardingTableClient) List(orgIdParam string, projectIdParam string, tier1IdParam string, componentTypeParam *string, cursorParam *string, edgeIdParam *string, edgePathParam *string, enforcementPointPathParam *string, includedFieldsParam *string, networkPrefixParam *string, pageSizeParam *int64, routeSourceParam *string, sortAscendingParam *bool, sortByParam *string) (nsx_policyModel.RoutingTableListResult, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	operationRestMetaData := forwardingTableListRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(forwardingTableListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("Tier1Id", tier1IdParam)
	sv.AddStructField("ComponentType", componentTypeParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("EdgeId", edgeIdParam)
	sv.AddStructField("EdgePath", edgePathParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NetworkPrefix", networkPrefixParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RouteSource", routeSourceParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput nsx_policyModel.RoutingTableListResult
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.tier_1s.forwarding_table", "list", inputDataValue, executionContext)
	var emptyOutput nsx_policyModel.RoutingTableListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ForwardingTableListOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(nsx_policyModel.RoutingTableListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
