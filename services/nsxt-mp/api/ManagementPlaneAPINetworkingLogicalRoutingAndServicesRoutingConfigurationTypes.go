// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfiguration.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``action`` of method ManagementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfiguration#unSetPasswordOnBgpNeighbor.
const ManagementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfiguration_UN_SET_PASSWORD_ON_BGP_NEIGHBOR_ACTION_PASSWORD = "clear_password"

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddBgpNeighborInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bgp_neighbor"] = bindings.NewReferenceType(model.BgpNeighborBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bgp_neighbor"] = "BgpNeighbor"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddBgpNeighborOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddBgpNeighborRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bgp_neighbor"] = bindings.NewReferenceType(model.BgpNeighborBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bgp_neighbor"] = "BgpNeighbor"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["bgp_neighbor"] = bindings.NewReferenceType(model.BgpNeighborBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"bgp_neighbor",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddIPPrefixListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["ip_prefix_list"] = bindings.NewReferenceType(model.IPPrefixListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["ip_prefix_list"] = "IpPrefixList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddIPPrefixListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPPrefixListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddIPPrefixListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["ip_prefix_list"] = bindings.NewReferenceType(model.IPPrefixListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["ip_prefix_list"] = "IpPrefixList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["ip_prefix_list"] = bindings.NewReferenceType(model.IPPrefixListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"ip_prefix_list",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/ip-prefix-lists",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddRouteMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["route_map"] = bindings.NewReferenceType(model.RouteMapBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["route_map"] = "RouteMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddRouteMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RouteMapBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddRouteMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["route_map"] = bindings.NewReferenceType(model.RouteMapBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["route_map"] = "RouteMap"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["route_map"] = bindings.NewReferenceType(model.RouteMapBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"route_map",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/route-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddStaticRouteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["static_route"] = bindings.NewReferenceType(model.StaticRouteBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["static_route"] = "StaticRoute"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddStaticRouteOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticRouteBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationAddStaticRouteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["static_route"] = bindings.NewReferenceType(model.StaticRouteBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["static_route"] = "StaticRoute"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["static_route"] = bindings.NewReferenceType(model.StaticRouteBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"static_route",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateBGPCommunityListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateBGPCommunityListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPCommunityListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateBGPCommunityListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"b_GP_community_list",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/community-lists",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateDADProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["d_AD_profile"] = bindings.NewReferenceType(model.DADProfileBindingType)
	fieldNameMap["d_AD_profile"] = "DADProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateDADProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DADProfileBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateDADProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["d_AD_profile"] = bindings.NewReferenceType(model.DADProfileBindingType)
	fieldNameMap["d_AD_profile"] = "DADProfile"
	paramsTypeMap["d_AD_profile"] = bindings.NewReferenceType(model.DADProfileBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"d_AD_profile",
		"POST",
		"/api/v1/ipv6/dad-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateNDRAProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["n_DRA_profile"] = bindings.NewReferenceType(model.NDRAProfileBindingType)
	fieldNameMap["n_DRA_profile"] = "NDRAProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateNDRAProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NDRAProfileBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationCreateNDRAProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["n_DRA_profile"] = bindings.NewReferenceType(model.NDRAProfileBindingType)
	fieldNameMap["n_DRA_profile"] = "NDRAProfile"
	paramsTypeMap["n_DRA_profile"] = bindings.NewReferenceType(model.NDRAProfileBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"n_DRA_profile",
		"POST",
		"/api/v1/ipv6/nd-ra-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteBGPCommunityListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteBGPCommunityListOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteBGPCommunityListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["community_list_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["communityListId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	pathParams["community_list_id"] = "communityListId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/community-lists/{communityListId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteBgpNeighborInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteBgpNeighborOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteBgpNeighborRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors/{id}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteDADProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dad_profile_id"] = bindings.NewStringType()
	fieldNameMap["dad_profile_id"] = "DadProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteDADProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteDADProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["dad_profile_id"] = bindings.NewStringType()
	fieldNameMap["dad_profile_id"] = "DadProfileId"
	paramsTypeMap["dad_profile_id"] = bindings.NewStringType()
	paramsTypeMap["dadProfileId"] = bindings.NewStringType()
	pathParams["dad_profile_id"] = "dadProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/ipv6/dad-profiles/{dadProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteIPPrefixListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteIPPrefixListOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteIPPrefixListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/logical-routers/{logicalRouterId}/routing/ip-prefix-lists/{id}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteNDRAProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nd_ra_profile_id"] = bindings.NewStringType()
	fieldNameMap["nd_ra_profile_id"] = "NdRaProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteNDRAProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteNDRAProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["nd_ra_profile_id"] = bindings.NewStringType()
	fieldNameMap["nd_ra_profile_id"] = "NdRaProfileId"
	paramsTypeMap["nd_ra_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ndRaProfileId"] = bindings.NewStringType()
	pathParams["nd_ra_profile_id"] = "ndRaProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/ipv6/nd-ra-profiles/{ndRaProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteRouteMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteRouteMapOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteRouteMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/logical-routers/{logicalRouterId}/routing/route-maps/{id}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteStaticRouteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteStaticRouteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationDeleteStaticRouteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"DELETE",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/{id}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListBGPCommunityListsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListBGPCommunityListsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPCommunityListListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListBGPCommunityListsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/community-lists",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListBgpNeighborsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListBgpNeighborsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListBgpNeighborsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListDADProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListDADProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DADProfileListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListDADProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/ipv6/dad-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListIPPrefixListsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListIPPrefixListsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPPrefixListListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListIPPrefixListsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/ip-prefix-lists",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListNDRAProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListNDRAProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NDRAProfileListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListNDRAProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/ipv6/nd-ra-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListRouteMapsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListRouteMapsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RouteMapListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListRouteMapsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/route-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListStaticRoutesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListStaticRoutesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticRouteListResultBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationListStaticRoutesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadAdvertiseRuleListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadAdvertiseRuleListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AdvertiseRuleListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadAdvertiseRuleListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/advertisement/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadAdvertisementConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadAdvertisementConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AdvertisementConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadAdvertisementConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/advertisement",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBGPCommunityListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBGPCommunityListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPCommunityListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBGPCommunityListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["community_list_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["communityListId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	pathParams["community_list_id"] = "communityListId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/community-lists/{communityListId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpNeighborInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpNeighborOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpNeighborRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpNeighborWithPasswordShowSensitiveDataInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpNeighborWithPasswordShowSensitiveDataOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadBgpNeighborWithPasswordShowSensitiveDataRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=show-sensitive-data",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadDADProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dad_profile_id"] = bindings.NewStringType()
	fieldNameMap["dad_profile_id"] = "DadProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadDADProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DADProfileBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadDADProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["dad_profile_id"] = bindings.NewStringType()
	fieldNameMap["dad_profile_id"] = "DadProfileId"
	paramsTypeMap["dad_profile_id"] = bindings.NewStringType()
	paramsTypeMap["dadProfileId"] = bindings.NewStringType()
	pathParams["dad_profile_id"] = "dadProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/ipv6/dad-profiles/{dadProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadDebugInfoTextInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadDebugInfoTextOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadDebugInfoTextRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"format=text",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/debug-info",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadIPPrefixListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadIPPrefixListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPPrefixListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadIPPrefixListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/ip-prefix-lists/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadNDRAProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nd_ra_profile_id"] = bindings.NewStringType()
	fieldNameMap["nd_ra_profile_id"] = "NdRaProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadNDRAProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NDRAProfileBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadNDRAProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["nd_ra_profile_id"] = bindings.NewStringType()
	fieldNameMap["nd_ra_profile_id"] = "NdRaProfileId"
	paramsTypeMap["nd_ra_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ndRaProfileId"] = bindings.NewStringType()
	pathParams["nd_ra_profile_id"] = "ndRaProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/ipv6/nd-ra-profiles/{ndRaProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRedistributionConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRedistributionConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RedistributionConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRedistributionConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/redistribution",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRedistributionRuleListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRedistributionRuleListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RedistributionRuleListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRedistributionRuleListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/redistribution/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRouteMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRouteMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RouteMapBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRouteMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/route-maps/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRoutingConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRoutingConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoutingConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadRoutingConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadStaticRouteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadStaticRouteOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticRouteBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationReadStaticRouteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUnSetPasswordOnBgpNeighborInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUnSetPasswordOnBgpNeighborOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUnSetPasswordOnBgpNeighborRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	queryParams["action"] = "action"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"POST",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateAdvertiseRuleListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["advertise_rule_list"] = bindings.NewReferenceType(model.AdvertiseRuleListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["advertise_rule_list"] = "AdvertiseRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateAdvertiseRuleListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AdvertiseRuleListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateAdvertiseRuleListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["advertise_rule_list"] = bindings.NewReferenceType(model.AdvertiseRuleListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["advertise_rule_list"] = "AdvertiseRuleList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["advertise_rule_list"] = bindings.NewReferenceType(model.AdvertiseRuleListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"advertise_rule_list",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/advertisement/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateAdvertisementConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["advertisement_config"] = bindings.NewReferenceType(model.AdvertisementConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["advertisement_config"] = "AdvertisementConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateAdvertisementConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AdvertisementConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateAdvertisementConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["advertisement_config"] = bindings.NewReferenceType(model.AdvertisementConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["advertisement_config"] = "AdvertisementConfig"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["advertisement_config"] = bindings.NewReferenceType(model.AdvertisementConfigBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"advertisement_config",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/advertisement",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBGPCommunityListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBGPCommunityListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPCommunityListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBGPCommunityListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["community_list_id"] = bindings.NewStringType()
	paramsTypeMap["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["communityListId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	pathParams["community_list_id"] = "communityListId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"b_GP_community_list",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/community-lists/{communityListId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBGPCommunityListOldInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBGPCommunityListOldOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPCommunityListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBGPCommunityListOldRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["community_list_id"] = bindings.NewStringType()
	paramsTypeMap["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["communityListId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	pathParams["community_list_id"] = "communityListId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"b_GP_community_list",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/communty-lists/{communityListId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBgpConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bgp_config"] = bindings.NewReferenceType(model.BgpConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bgp_config"] = "BgpConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBgpConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBgpConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["bgp_config"] = bindings.NewReferenceType(model.BgpConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["bgp_config"] = "BgpConfig"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["bgp_config"] = bindings.NewReferenceType(model.BgpConfigBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"bgp_config",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBgpNeighborInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["bgp_neighbor"] = bindings.NewReferenceType(model.BgpNeighborBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["bgp_neighbor"] = "BgpNeighbor"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBgpNeighborOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateBgpNeighborRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["bgp_neighbor"] = bindings.NewReferenceType(model.BgpNeighborBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["bgp_neighbor"] = "BgpNeighbor"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["bgp_neighbor"] = bindings.NewReferenceType(model.BgpNeighborBindingType)
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"bgp_neighbor",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/neighbors/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateDADProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dad_profile_id"] = bindings.NewStringType()
	fields["d_AD_profile"] = bindings.NewReferenceType(model.DADProfileBindingType)
	fieldNameMap["dad_profile_id"] = "DadProfileId"
	fieldNameMap["d_AD_profile"] = "DADProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateDADProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DADProfileBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateDADProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["dad_profile_id"] = bindings.NewStringType()
	fields["d_AD_profile"] = bindings.NewReferenceType(model.DADProfileBindingType)
	fieldNameMap["dad_profile_id"] = "DadProfileId"
	fieldNameMap["d_AD_profile"] = "DADProfile"
	paramsTypeMap["d_AD_profile"] = bindings.NewReferenceType(model.DADProfileBindingType)
	paramsTypeMap["dad_profile_id"] = bindings.NewStringType()
	paramsTypeMap["dadProfileId"] = bindings.NewStringType()
	pathParams["dad_profile_id"] = "dadProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"d_AD_profile",
		"PUT",
		"/api/v1/ipv6/dad-profiles/{dadProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateIPPrefixListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["ip_prefix_list"] = bindings.NewReferenceType(model.IPPrefixListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["ip_prefix_list"] = "IpPrefixList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateIPPrefixListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPPrefixListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateIPPrefixListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["ip_prefix_list"] = bindings.NewReferenceType(model.IPPrefixListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["ip_prefix_list"] = "IpPrefixList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["ip_prefix_list"] = bindings.NewReferenceType(model.IPPrefixListBindingType)
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"ip_prefix_list",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/ip-prefix-lists/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateNDRAProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nd_ra_profile_id"] = bindings.NewStringType()
	fields["n_DRA_profile"] = bindings.NewReferenceType(model.NDRAProfileBindingType)
	fieldNameMap["nd_ra_profile_id"] = "NdRaProfileId"
	fieldNameMap["n_DRA_profile"] = "NDRAProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateNDRAProfileOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NDRAProfileBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateNDRAProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["nd_ra_profile_id"] = bindings.NewStringType()
	fields["n_DRA_profile"] = bindings.NewReferenceType(model.NDRAProfileBindingType)
	fieldNameMap["nd_ra_profile_id"] = "NdRaProfileId"
	fieldNameMap["n_DRA_profile"] = "NDRAProfile"
	paramsTypeMap["nd_ra_profile_id"] = bindings.NewStringType()
	paramsTypeMap["n_DRA_profile"] = bindings.NewReferenceType(model.NDRAProfileBindingType)
	paramsTypeMap["ndRaProfileId"] = bindings.NewStringType()
	pathParams["nd_ra_profile_id"] = "ndRaProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"n_DRA_profile",
		"PUT",
		"/api/v1/ipv6/nd-ra-profiles/{ndRaProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRedistributionConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["redistribution_config"] = bindings.NewReferenceType(model.RedistributionConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["redistribution_config"] = "RedistributionConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRedistributionConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RedistributionConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRedistributionConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["redistribution_config"] = bindings.NewReferenceType(model.RedistributionConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["redistribution_config"] = "RedistributionConfig"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["redistribution_config"] = bindings.NewReferenceType(model.RedistributionConfigBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"redistribution_config",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/redistribution",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRedistributionRuleListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["redistribution_rule_list"] = bindings.NewReferenceType(model.RedistributionRuleListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["redistribution_rule_list"] = "RedistributionRuleList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRedistributionRuleListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RedistributionRuleListBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRedistributionRuleListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["redistribution_rule_list"] = bindings.NewReferenceType(model.RedistributionRuleListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["redistribution_rule_list"] = "RedistributionRuleList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["redistribution_rule_list"] = bindings.NewReferenceType(model.RedistributionRuleListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"redistribution_rule_list",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/redistribution/rules",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRouteMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["route_map"] = bindings.NewReferenceType(model.RouteMapBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["route_map"] = "RouteMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRouteMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RouteMapBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRouteMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["route_map"] = bindings.NewReferenceType(model.RouteMapBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["route_map"] = "RouteMap"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["route_map"] = bindings.NewReferenceType(model.RouteMapBindingType)
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"route_map",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/route-maps/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRoutingConfigInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["routing_config"] = bindings.NewReferenceType(model.RoutingConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["routing_config"] = "RoutingConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRoutingConfigOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoutingConfigBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateRoutingConfigRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["routing_config"] = bindings.NewReferenceType(model.RoutingConfigBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["routing_config"] = "RoutingConfig"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["routing_config"] = bindings.NewReferenceType(model.RoutingConfigBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"routing_config",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateStaticRouteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["static_route"] = bindings.NewReferenceType(model.StaticRouteBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["static_route"] = "StaticRoute"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateStaticRouteOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.StaticRouteBindingType)
}

func managementPlaneAPINetworkingLogicalRoutingAndServicesRoutingConfigurationUpdateStaticRouteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["id"] = bindings.NewStringType()
	fields["static_route"] = bindings.NewReferenceType(model.StaticRouteBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["id"] = "Id"
	fieldNameMap["static_route"] = "StaticRoute"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["static_route"] = bindings.NewReferenceType(model.StaticRouteBindingType)
	paramsTypeMap["id"] = bindings.NewStringType()
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["id"] = bindings.NewStringType()
	pathParams["id"] = "id"
	pathParams["logical_router_id"] = "logicalRouterId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"static_route",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/static-routes/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
