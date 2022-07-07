// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AdvertisedRoutes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package neighbors

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func advertisedRoutesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_node_id"] = bindings.NewStringType()
	fields["neighbor_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func advertisedRoutesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BgpNeighborRouteDetailsBindingType)
}

func advertisedRoutesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_node_id"] = bindings.NewStringType()
	fields["neighbor_id"] = bindings.NewStringType()
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	fieldNameMap["neighbor_id"] = "NeighborId"
	paramsTypeMap["neighbor_id"] = bindings.NewStringType()
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	paramsTypeMap["neighborId"] = bindings.NewStringType()
	pathParams["neighbor_id"] = "neighborId"
	pathParams["edge_node_id"] = "edgeNodeId"
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
		"/api/v1/transport-nodes/{edgeNodeId}/inter-site/bgp/neighbors/{neighborId}/advertised-routes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
