// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRouters
// Used by client-side stubs.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
)

const _ = core.SupportedByRuntimeVersion1

type ManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient interface {

	// Creates a logical router. The required parameters are router_type (TIER0 or TIER1) and edge_cluster_id (TIER0 only). Optional parameters include internal and external transit network addresses.
	//
	// @param logicalRouterParam (required)
	// @return com.vmware.model.LogicalRouter
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	CreateLogicalRouter(logicalRouterParam model.LogicalRouter) (model.LogicalRouter, error)

	// Deletes the specified logical router. You must delete associated logical router ports before you can delete a logical router. Otherwise use force delete which will delete all related ports and other entities associated with that LR. To force delete logical router pass force=true in query param.
	//
	// @param logicalRouterIdParam (required)
	// @param cascadeDeleteLinkedPortsParam Flag to specify whether to delete related logical switch ports (optional, default to false)
	// @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	DeleteLogicalRouter(logicalRouterIdParam string, cascadeDeleteLinkedPortsParam *bool, forceParam *bool) error

	// Returns routes advertised by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled. It always returns realtime response.
	//
	// @param logicalRouterIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.model.BgpNeighborRouteDetails
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBgpNeighborAdvertisedRoutes(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error)

	// Returns routes advertised by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled in CSV format. It always returns realtime response.
	//
	// @param logicalRouterIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.model.BgpNeighborRouteDetailsInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBgpNeighborAdvertisedRoutesInCsvFormatCsv(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetailsInCsvFormat, error)

	// Returns routes learned by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled. It always returns realtime response.
	//
	// @param logicalRouterIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.model.BgpNeighborRouteDetails
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBgpNeighborRoutes(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error)

	// Returns routes learned by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled in CSV format. It always returns realtime response.
	//
	// @param logicalRouterIdParam (required)
	// @param neighborIdParam (required)
	// @return com.vmware.model.BgpNeighborRouteDetailsInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBgpNeighborRoutesInCsvFormatCsv(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetailsInCsvFormat, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param transportNodeIdParam Transport node id (optional)
	// @return com.vmware.model.BgpNeighborsStatusListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetBgpNeighborsStatus(logicalRouterIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.BgpNeighborsStatusListResult, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param transportNodeIdParam TransportNode Id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param networkPrefixParam IPv4 or IPv6 CIDR Block (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalRouterRouteTable
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterForwardingTable(logicalRouterIdParam string, transportNodeIdParam string, cursorParam *string, includedFieldsParam *string, networkPrefixParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string) (model.LogicalRouterRouteTable, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param transportNodeIdParam TransportNode Id (required)
	// @param networkPrefixParam IPv4 or IPv6 CIDR Block (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalRouterRouteTableInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterForwardingTableInCsvFormatCsv(logicalRouterIdParam string, transportNodeIdParam string, networkPrefixParam *string, sourceParam *string) (model.LogicalRouterRouteTableInCsvFormat, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param transportNodeIdParam TransportNode Id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalRouterRouteTable
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterRouteTable(logicalRouterIdParam string, transportNodeIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string) (model.LogicalRouterRouteTable, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param transportNodeIdParam TransportNode Id (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalRouterRouteTableInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterRouteTableInCsvFormatCsv(logicalRouterIdParam string, transportNodeIdParam string, sourceParam *string) (model.LogicalRouterRouteTableInCsvFormat, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param transportNodeIdParam TransportNode Id (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param networkPrefixParam IPv4 or IPv6 CIDR Block (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param routeSourceParam Route source filter parameter (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param sourceParam Data source type. (optional)
	// @param vrfTableParam VRF filter parameter (optional)
	// @return com.vmware.model.LogicalRouterRouteTable
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterRoutingTable(logicalRouterIdParam string, transportNodeIdParam string, cursorParam *string, includedFieldsParam *string, networkPrefixParam *string, pageSizeParam *int64, routeSourceParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, vrfTableParam *string) (model.LogicalRouterRouteTable, error)

	//
	//
	// @param logicalRouterIdParam (required)
	// @param transportNodeIdParam TransportNode Id (required)
	// @param networkPrefixParam IPv4 or IPv6 CIDR Block (optional)
	// @param routeSourceParam Route source filter parameter (optional)
	// @param sourceParam Data source type. (optional)
	// @param vrfTableParam VRF filter parameter (optional)
	// @return com.vmware.model.LogicalRouterRouteTableInCsvFormat
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterRoutingTableInCsvFormatCsv(logicalRouterIdParam string, transportNodeIdParam string, networkPrefixParam *string, routeSourceParam *string, sourceParam *string, vrfTableParam *string) (model.LogicalRouterRouteTableInCsvFormat, error)

	// Return realized state information of a logical router. Any configuration update that affects the logical router can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of logical router, static routes, etc.
	//
	// @param logicalRouterIdParam (required)
	// @param barrierIdParam (optional)
	// @param requestIdParam Realization request ID (optional)
	// @return com.vmware.model.LogicalRouterState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterState(logicalRouterIdParam string, barrierIdParam *int64, requestIdParam *string) (model.LogicalRouterState, error)

	// Returns status for the Logical Router of the given id.
	//
	// @param logicalRouterIdParam (required)
	// @param sourceParam Data source type. (optional)
	// @return com.vmware.model.LogicalRouterStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalRouterStatus(logicalRouterIdParam string, sourceParam *string) (model.LogicalRouterStatus, error)

	// Return realized state information of a logical service router cluster. Any configuration update that affects the logical service router cluster can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of nat, bgp, bfd, etc. What is a Service Router? When a service cannot be distributed is enabled on a Logical Router, a Service Router (SR) is instantiated. Some examples of services that are not distributed are NAT, DHCP server, Metadata Proxy, Edge Firewall, Load Balancer and so on.
	//
	// @param logicalRouterIdParam (required)
	// @param barrierIdParam (optional)
	// @param requestIdParam Realization request ID (optional)
	// @return com.vmware.model.LogicalServiceRouterClusterState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	GetLogicalServiceRouterClusterState(logicalRouterIdParam string, barrierIdParam *int64, requestIdParam *string) (model.LogicalServiceRouterClusterState, error)

	// Returns information about all logical routers, including the UUID, internal and external transit network addresses, and the router type (TIER0 or TIER1). You can get information for only TIER0 routers or only the TIER1 routers by including the router_type query parameter.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param routerTypeParam Type of Logical Router (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param vrfsOnLogicalRouterIdParam List all VRFs on the specified logical router. (optional)
	// @return com.vmware.model.LogicalRouterListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ListLogicalRouters(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, routerTypeParam *string, sortAscendingParam *bool, sortByParam *string, vrfsOnLogicalRouterIdParam *string) (model.LogicalRouterListResult, error)

	// API to re allocate edge node placement for TIER1 logical router. You can re-allocate service routers of TIER1 in same edge cluster or different edge cluster. You can also place edge nodes manually and provide maximum two indices for HA mode ACTIVE_STANDBY. To re-allocate on new edge cluster you must have existing edge cluster for TIER1 logical router. This will be disruptive operation and all existing statistics of logical router will be remove.
	//
	// @param logicalRouterIdParam (required)
	// @param serviceRouterAllocationConfigParam (required)
	// @return com.vmware.model.LogicalRouter
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReAllocateServiceRoutersReallocate(logicalRouterIdParam string, serviceRouterAllocationConfigParam model.ServiceRouterAllocationConfig) (model.LogicalRouter, error)

	// Reprocess logical router configuration and configuration of related entities like logical router ports, static routing, etc. Any missing Updates are published to controller.
	//
	// @param logicalRouterIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReProcessLogicalRouterReprocess(logicalRouterIdParam string) error

	// Returns information about the specified logical router.
	//
	// @param logicalRouterIdParam (required)
	// @return com.vmware.model.LogicalRouter
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	ReadLogicalRouter(logicalRouterIdParam string) (model.LogicalRouter, error)

	// Modifies the specified logical router. Modifiable attributes include the internal_transit_network, external_transit_networks, and edge_cluster_id (for TIER0 routers).
	//
	// @param logicalRouterIdParam (required)
	// @param logicalRouterParam (required)
	// @return com.vmware.model.LogicalRouter
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	UpdateLogicalRouter(logicalRouterIdParam string, logicalRouterParam model.LogicalRouter) (model.LogicalRouter, error)
}

type managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewManagementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient(connector client.Connector) *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create_logical_router":                                 core.NewMethodIdentifier(interfaceIdentifier, "create_logical_router"),
		"delete_logical_router":                                 core.NewMethodIdentifier(interfaceIdentifier, "delete_logical_router"),
		"get_bgp_neighbor_advertised_routes":                    core.NewMethodIdentifier(interfaceIdentifier, "get_bgp_neighbor_advertised_routes"),
		"get_bgp_neighbor_advertised_routes_in_csv_format_csv":  core.NewMethodIdentifier(interfaceIdentifier, "get_bgp_neighbor_advertised_routes_in_csv_format_csv"),
		"get_bgp_neighbor_routes":                               core.NewMethodIdentifier(interfaceIdentifier, "get_bgp_neighbor_routes"),
		"get_bgp_neighbor_routes_in_csv_format_csv":             core.NewMethodIdentifier(interfaceIdentifier, "get_bgp_neighbor_routes_in_csv_format_csv"),
		"get_bgp_neighbors_status":                              core.NewMethodIdentifier(interfaceIdentifier, "get_bgp_neighbors_status"),
		"get_logical_router_forwarding_table":                   core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_forwarding_table"),
		"get_logical_router_forwarding_table_in_csv_format_csv": core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_forwarding_table_in_csv_format_csv"),
		"get_logical_router_route_table":                        core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_route_table"),
		"get_logical_router_route_table_in_csv_format_csv":      core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_route_table_in_csv_format_csv"),
		"get_logical_router_routing_table":                      core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_routing_table"),
		"get_logical_router_routing_table_in_csv_format_csv":    core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_routing_table_in_csv_format_csv"),
		"get_logical_router_state":                              core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_state"),
		"get_logical_router_status":                             core.NewMethodIdentifier(interfaceIdentifier, "get_logical_router_status"),
		"get_logical_service_router_cluster_state":              core.NewMethodIdentifier(interfaceIdentifier, "get_logical_service_router_cluster_state"),
		"list_logical_routers":                                  core.NewMethodIdentifier(interfaceIdentifier, "list_logical_routers"),
		"re_allocate_service_routers_reallocate":                core.NewMethodIdentifier(interfaceIdentifier, "re_allocate_service_routers_reallocate"),
		"re_process_logical_router_reprocess":                   core.NewMethodIdentifier(interfaceIdentifier, "re_process_logical_router_reprocess"),
		"read_logical_router":                                   core.NewMethodIdentifier(interfaceIdentifier, "read_logical_router"),
		"update_logical_router":                                 core.NewMethodIdentifier(interfaceIdentifier, "update_logical_router"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) CreateLogicalRouter(logicalRouterParam model.LogicalRouter) (model.LogicalRouter, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersCreateLogicalRouterInputType(), typeConverter)
	sv.AddStructField("LogicalRouter", logicalRouterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouter
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersCreateLogicalRouterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "create_logical_router", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouter
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersCreateLogicalRouterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouter), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) DeleteLogicalRouter(logicalRouterIdParam string, cascadeDeleteLinkedPortsParam *bool, forceParam *bool) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersDeleteLogicalRouterInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("CascadeDeleteLinkedPorts", cascadeDeleteLinkedPortsParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersDeleteLogicalRouterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "delete_logical_router", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetBgpNeighborAdvertisedRoutes(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborAdvertisedRoutesInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborRouteDetails
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborAdvertisedRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_bgp_neighbor_advertised_routes", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborRouteDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborAdvertisedRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborRouteDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetBgpNeighborAdvertisedRoutesInCsvFormatCsv(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetailsInCsvFormat, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborAdvertisedRoutesInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborRouteDetailsInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborAdvertisedRoutesInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_bgp_neighbor_advertised_routes_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborRouteDetailsInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborAdvertisedRoutesInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborRouteDetailsInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetBgpNeighborRoutes(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetails, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborRoutesInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborRouteDetails
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborRoutesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_bgp_neighbor_routes", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborRouteDetails
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborRoutesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborRouteDetails), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetBgpNeighborRoutesInCsvFormatCsv(logicalRouterIdParam string, neighborIdParam string) (model.BgpNeighborRouteDetailsInCsvFormat, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborRoutesInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("NeighborId", neighborIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborRouteDetailsInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborRoutesInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_bgp_neighbor_routes_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborRouteDetailsInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborRoutesInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborRouteDetailsInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetBgpNeighborsStatus(logicalRouterIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string, transportNodeIdParam *string) (model.BgpNeighborsStatusListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborsStatusInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.BgpNeighborsStatusListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborsStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_bgp_neighbors_status", inputDataValue, executionContext)
	var emptyOutput model.BgpNeighborsStatusListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetBgpNeighborsStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.BgpNeighborsStatusListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterForwardingTable(logicalRouterIdParam string, transportNodeIdParam string, cursorParam *string, includedFieldsParam *string, networkPrefixParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string) (model.LogicalRouterRouteTable, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterForwardingTableInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NetworkPrefix", networkPrefixParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterRouteTable
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterForwardingTableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_forwarding_table", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterRouteTable
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterForwardingTableOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterRouteTable), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterForwardingTableInCsvFormatCsv(logicalRouterIdParam string, transportNodeIdParam string, networkPrefixParam *string, sourceParam *string) (model.LogicalRouterRouteTableInCsvFormat, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterForwardingTableInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("NetworkPrefix", networkPrefixParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterRouteTableInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterForwardingTableInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_forwarding_table_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterRouteTableInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterForwardingTableInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterRouteTableInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterRouteTable(logicalRouterIdParam string, transportNodeIdParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, sourceParam *string) (model.LogicalRouterRouteTable, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRouteTableInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterRouteTable
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRouteTableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_route_table", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterRouteTable
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRouteTableOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterRouteTable), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterRouteTableInCsvFormatCsv(logicalRouterIdParam string, transportNodeIdParam string, sourceParam *string) (model.LogicalRouterRouteTableInCsvFormat, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRouteTableInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterRouteTableInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRouteTableInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_route_table_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterRouteTableInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRouteTableInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterRouteTableInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterRoutingTable(logicalRouterIdParam string, transportNodeIdParam string, cursorParam *string, includedFieldsParam *string, networkPrefixParam *string, pageSizeParam *int64, routeSourceParam *string, sortAscendingParam *bool, sortByParam *string, sourceParam *string, vrfTableParam *string) (model.LogicalRouterRouteTable, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRoutingTableInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NetworkPrefix", networkPrefixParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RouteSource", routeSourceParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("VrfTable", vrfTableParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterRouteTable
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRoutingTableRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_routing_table", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterRouteTable
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRoutingTableOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterRouteTable), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterRoutingTableInCsvFormatCsv(logicalRouterIdParam string, transportNodeIdParam string, networkPrefixParam *string, routeSourceParam *string, sourceParam *string, vrfTableParam *string) (model.LogicalRouterRouteTableInCsvFormat, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRoutingTableInCsvFormatCsvInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("TransportNodeId", transportNodeIdParam)
	sv.AddStructField("NetworkPrefix", networkPrefixParam)
	sv.AddStructField("RouteSource", routeSourceParam)
	sv.AddStructField("Source", sourceParam)
	sv.AddStructField("VrfTable", vrfTableParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterRouteTableInCsvFormat
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRoutingTableInCsvFormatCsvRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_routing_table_in_csv_format_csv", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterRouteTableInCsvFormat
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterRoutingTableInCsvFormatCsvOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterRouteTableInCsvFormat), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterState(logicalRouterIdParam string, barrierIdParam *int64, requestIdParam *string) (model.LogicalRouterState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterStateInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("BarrierId", barrierIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_state", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalRouterStatus(logicalRouterIdParam string, sourceParam *string) (model.LogicalRouterStatus, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterStatusInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("Source", sourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterStatusRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_router_status", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalRouterStatusOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) GetLogicalServiceRouterClusterState(logicalRouterIdParam string, barrierIdParam *int64, requestIdParam *string) (model.LogicalServiceRouterClusterState, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalServiceRouterClusterStateInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("BarrierId", barrierIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalServiceRouterClusterState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalServiceRouterClusterStateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "get_logical_service_router_cluster_state", inputDataValue, executionContext)
	var emptyOutput model.LogicalServiceRouterClusterState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersGetLogicalServiceRouterClusterStateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalServiceRouterClusterState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) ListLogicalRouters(cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, routerTypeParam *string, sortAscendingParam *bool, sortByParam *string, vrfsOnLogicalRouterIdParam *string) (model.LogicalRouterListResult, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersListLogicalRoutersInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("RouterType", routerTypeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("VrfsOnLogicalRouterId", vrfsOnLogicalRouterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersListLogicalRoutersRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "list_logical_routers", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersListLogicalRoutersOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) ReAllocateServiceRoutersReallocate(logicalRouterIdParam string, serviceRouterAllocationConfigParam model.ServiceRouterAllocationConfig) (model.LogicalRouter, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReAllocateServiceRoutersReallocateInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("ServiceRouterAllocationConfig", serviceRouterAllocationConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouter
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReAllocateServiceRoutersReallocateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "re_allocate_service_routers_reallocate", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouter
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReAllocateServiceRoutersReallocateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouter), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) ReProcessLogicalRouterReprocess(logicalRouterIdParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReProcessLogicalRouterReprocessInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReProcessLogicalRouterReprocessRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "re_process_logical_router_reprocess", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) ReadLogicalRouter(logicalRouterIdParam string) (model.LogicalRouter, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReadLogicalRouterInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouter
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReadLogicalRouterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "read_logical_router", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouter
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersReadLogicalRouterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouter), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersClient) UpdateLogicalRouter(logicalRouterIdParam string, logicalRouterParam model.LogicalRouter) (model.LogicalRouter, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersUpdateLogicalRouterInputType(), typeConverter)
	sv.AddStructField("LogicalRouterId", logicalRouterIdParam)
	sv.AddStructField("LogicalRouter", logicalRouterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouter
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersUpdateLogicalRouterRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.api.management_plane_API_networking_logical_routing_and_services_logical_routers", "update_logical_router", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouter
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), managementPlaneAPINetworkingLogicalRoutingAndServicesLogicalRoutersUpdateLogicalRouterOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouter), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
