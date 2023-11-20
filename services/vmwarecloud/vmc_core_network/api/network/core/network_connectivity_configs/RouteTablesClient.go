// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RouteTables
// Used by client-side stubs.

package network_connectivity_configs

import (
	vapiStdErrors_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiCore_ "github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	vapiProtocolClient_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	vmwarecloudVmc_core_networkModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_network/model"
)

const _ = vapiCore_.SupportedByRuntimeVersion2

type RouteTablesClient interface {

	// Get route tables associated with network connectivity config
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @param includeDeletedResourcesParam should the deleted resources be included in the results?
	// @param pageParam Number of requested page (0-based)
	// @param sizeParam Size of each page
	// @param sortParam Sort order of paged requests
	// @param filterParam Filtering criteria of paged requests
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the deployment with given identifier
	GetRouteTablesForNetworkConnectivityConfig(orgIdParam string, idParam string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string, filterParam []string) (vmwarecloudVmc_core_networkModel.RouteTableResponse, error)

	// Get a specific route table associated with core NetworkConnectivityConfig
	//
	// @param orgIdParam organization identifier
	// @param idParam resource identifier
	// @param routeTableIdParam Route Table Id
	// @return OK
	//
	// @throws Unauthenticated Unauthorized
	// @throws Unauthorized Forbidden
	// @throws NotFound Cannot find the deployment with given identifier
	GetAssociatedRouteTableById(orgIdParam string, idParam string, routeTableIdParam string) (vmwarecloudVmc_core_networkModel.RouteTable, error)
}

type routeTablesClient struct {
	connector           vapiProtocolClient_.Connector
	interfaceDefinition vapiCore_.InterfaceDefinition
	errorsBindingMap    map[string]vapiBindings_.BindingType
}

func NewRouteTablesClient(connector vapiProtocolClient_.Connector) *routeTablesClient {
	interfaceIdentifier := vapiCore_.NewInterfaceIdentifier("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.route_tables")
	methodIdentifiers := map[string]vapiCore_.MethodIdentifier{
		"get_route_tables_for_network_connectivity_config": vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_route_tables_for_network_connectivity_config"),
		"get_associated_route_table_by_id":                 vapiCore_.NewMethodIdentifier(interfaceIdentifier, "get_associated_route_table_by_id"),
	}
	interfaceDefinition := vapiCore_.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]vapiBindings_.BindingType)

	rIface := routeTablesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *routeTablesClient) GetErrorBindingType(errorName string) vapiBindings_.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return vapiStdErrors_.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *routeTablesClient) GetRouteTablesForNetworkConnectivityConfig(orgIdParam string, idParam string, includeDeletedResourcesParam *bool, pageParam *int64, sizeParam *int64, sortParam []string, filterParam []string) (vmwarecloudVmc_core_networkModel.RouteTableResponse, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routeTablesGetRouteTablesForNetworkConnectivityConfigRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routeTablesGetRouteTablesForNetworkConnectivityConfigInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("IncludeDeletedResources", includeDeletedResourcesParam)
	sv.AddStructField("Page", pageParam)
	sv.AddStructField("Size", sizeParam)
	sv.AddStructField("Sort", sortParam)
	sv.AddStructField("Filter", filterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.RouteTableResponse
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.route_tables", "get_route_tables_for_network_connectivity_config", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.RouteTableResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RouteTablesGetRouteTablesForNetworkConnectivityConfigOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.RouteTableResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *routeTablesClient) GetAssociatedRouteTableById(orgIdParam string, idParam string, routeTableIdParam string) (vmwarecloudVmc_core_networkModel.RouteTable, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	operationRestMetaData := routeTablesGetAssociatedRouteTableByIdRestMetadata()
	executionContext.SetConnectionMetadata(vapiCore_.RESTMetadataKey, operationRestMetaData)
	executionContext.SetConnectionMetadata(vapiCore_.ResponseTypeKey, vapiCore_.NewResponseType(true, false))

	sv := vapiBindings_.NewStructValueBuilder(routeTablesGetAssociatedRouteTableByIdInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("RouteTableId", routeTableIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput vmwarecloudVmc_core_networkModel.RouteTable
		return emptyOutput, vapiBindings_.VAPIerrorsToError(inputError)
	}

	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.vmwarecloud.vmc_core_network.api.network.core.network_connectivity_configs.route_tables", "get_associated_route_table_by_id", inputDataValue, executionContext)
	var emptyOutput vmwarecloudVmc_core_networkModel.RouteTable
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), RouteTablesGetAssociatedRouteTableByIdOutputType())
		if errorInOutput != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInOutput)
		}
		return output.(vmwarecloudVmc_core_networkModel.RouteTable), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, vapiBindings_.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
