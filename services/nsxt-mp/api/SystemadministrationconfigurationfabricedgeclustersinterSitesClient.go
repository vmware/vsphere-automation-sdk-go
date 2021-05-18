// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: SystemadministrationconfigurationfabricedgeclustersinterSites
// Used by client-side stubs.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type SystemadministrationconfigurationfabricedgeclustersinterSitesClient interface {

	// Returns the aggregated status for the Edge cluster along with status of all edge nodes in the cluster. It always returns cached response.
	//
	// @param edgeClusterIdParam (required)
	// @return com.vmware.model.EdgeClusterInterSiteStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetEdgeClusterInterSiteStatus(edgeClusterIdParam string) (model.EdgeClusterInterSiteStatus, error)

	// Returns routes advertised by BGP neighbor from the given edge transport node. It always returns realtime response.
	//
	// @param edgeNodeIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.model.BgpNeighborRouteDetails
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetInterSiteEdgeNodeBgpNeighborAdvertisedRoutes(edgeNodeIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error)

	// Returns routes learned by BGP neighbor from the given edge transport node. It always returns realtime response.
	//
	// @param edgeNodeIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.model.BgpNeighborRouteDetails
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetInterSiteEdgeNodeBgpNeighborRoutes(edgeNodeIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error)

	// Returns BGP summary for all configured neighbors in tunnel VRF on the given egde node. It always returns realtime response.
	//
	// @param edgeNodeIdParam (required)
	// @return com.vmware.model.InterSiteBgpSummary
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetInterSiteEdgeNodeBgpSummary(edgeNodeIdParam string) (model.InterSiteBgpSummary, error)

	// Returns RTEP to RTEP tunnel port statistics of the given edge node. It always returns realtime response.
	//
	// @param edgeNodeIdParam (required)
	// @return com.vmware.model.NodeInterSiteStatistics
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetInterSiteEdgeNodeStatistics(edgeNodeIdParam string) (model.NodeInterSiteStatistics, error)

	// Paginated list of BGP Neighbors on edge transport node.
	//
	// @param edgeNodeIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.model.BgpNeighborListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListInterSiteEdgeNodeBgpNeighbors(edgeNodeIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BgpNeighborListResult, error)
}

type systemadministrationconfigurationfabricedgeclustersinterSitesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSystemadministrationconfigurationfabricedgeclustersinterSitesClient(connector client.Connector) *systemadministrationconfigurationfabricedgeclustersinterSitesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get_edge_cluster_inter_site_status":                      core.NewMethodIdentifier(interfaceIdentifier, "get_edge_cluster_inter_site_status"),
		"get_inter_site_edge_node_bgp_neighbor_advertised_routes": core.NewMethodIdentifier(interfaceIdentifier, "get_inter_site_edge_node_bgp_neighbor_advertised_routes"),
		"get_inter_site_edge_node_bgp_neighbor_routes":            core.NewMethodIdentifier(interfaceIdentifier, "get_inter_site_edge_node_bgp_neighbor_routes"),
		"get_inter_site_edge_node_bgp_summary":                    core.NewMethodIdentifier(interfaceIdentifier, "get_inter_site_edge_node_bgp_summary"),
		"get_inter_site_edge_node_statistics":                     core.NewMethodIdentifier(interfaceIdentifier, "get_inter_site_edge_node_statistics"),
		"list_inter_site_edge_node_bgp_neighbors":                 core.NewMethodIdentifier(interfaceIdentifier, "list_inter_site_edge_node_bgp_neighbors"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := systemadministrationconfigurationfabricedgeclustersinterSitesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) GetEdgeClusterInterSiteStatus(edgeClusterIdParam string) (model.EdgeClusterInterSiteStatus, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemadministrationconfigurationfabricedgeclustersinterSitesGetEdgeClusterInterSiteStatusInputType(), typeConverter)
	sv.AddStructField("EdgeClusterId", edgeClusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EdgeClusterInterSiteStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemadministrationconfigurationfabricedgeclustersinterSitesGetEdgeClusterInterSiteStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites", "get_edge_cluster_inter_site_status", inputDataValue, executionContext)
	var emptyOutput model.EdgeClusterInterSiteStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemadministrationconfigurationfabricedgeclustersinterSitesGetEdgeClusterInterSiteStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EdgeClusterInterSiteStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) GetInterSiteEdgeNodeBgpNeighborAdvertisedRoutes(edgeNodeIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborAdvertisedRoutesInputType(), typeConverter)
	sv.AddStructField("EdgeNodeId", edgeNodeIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborRouteDetails
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborAdvertisedRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites", "get_inter_site_edge_node_bgp_neighbor_advertised_routes", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborRouteDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborAdvertisedRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborRouteDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) GetInterSiteEdgeNodeBgpNeighborRoutes(edgeNodeIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborRoutesInputType(), typeConverter)
	sv.AddStructField("EdgeNodeId", edgeNodeIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborRouteDetails
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites", "get_inter_site_edge_node_bgp_neighbor_routes", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborRouteDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpNeighborRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborRouteDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) GetInterSiteEdgeNodeBgpSummary(edgeNodeIdParam string) (model.InterSiteBgpSummary, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpSummaryInputType(), typeConverter)
	sv.AddStructField("EdgeNodeId", edgeNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.InterSiteBgpSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpSummaryRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites", "get_inter_site_edge_node_bgp_summary", inputDataValue, executionContext)
	var emptyOutput model.InterSiteBgpSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeBgpSummaryOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.InterSiteBgpSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) GetInterSiteEdgeNodeStatistics(edgeNodeIdParam string) (model.NodeInterSiteStatistics, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeStatisticsInputType(), typeConverter)
	sv.AddStructField("EdgeNodeId", edgeNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NodeInterSiteStatistics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeStatisticsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites", "get_inter_site_edge_node_statistics", inputDataValue, executionContext)
	var emptyOutput model.NodeInterSiteStatistics
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemadministrationconfigurationfabricedgeclustersinterSitesGetInterSiteEdgeNodeStatisticsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NodeInterSiteStatistics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *systemadministrationconfigurationfabricedgeclustersinterSitesClient) ListInterSiteEdgeNodeBgpNeighbors(edgeNodeIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.BgpNeighborListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(systemadministrationconfigurationfabricedgeclustersinterSitesListInterSiteEdgeNodeBgpNeighborsInputType(), typeConverter)
	sv.AddStructField("EdgeNodeId", edgeNodeIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := systemadministrationconfigurationfabricedgeclustersinterSitesListInterSiteEdgeNodeBgpNeighborsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.api.systemadministrationconfigurationfabricedgeclustersinter_sites", "list_inter_site_edge_node_bgp_neighbors", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), systemadministrationconfigurationfabricedgeclustersinterSitesListInterSiteEdgeNodeBgpNeighborsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
