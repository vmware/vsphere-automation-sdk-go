// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AdvertisedRoutes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package neighbors

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func advertisedRoutesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func AdvertisedRoutesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.BgpNeighborRouteDetailsBindingType)
}

func advertisedRoutesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = vapiBindings_.NewStringType()
	fields["neighbor_id"] = vapiBindings_.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	paramsTypeMap["neighbor_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edge_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["edgeNodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["neighborId"] = vapiBindings_.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["edge_node_id"] = "edgeNodeId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/bgp/neighbors/{neighborId}/advertised-routes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
