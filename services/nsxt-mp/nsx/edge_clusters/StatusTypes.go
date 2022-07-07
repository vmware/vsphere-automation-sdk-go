// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package edge_clusters

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Status#get.
const Status_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Status#get.
const Status_GET_SOURCE_CACHED = "cached"

func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterStatusBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
	queryParams["source"] = "source"
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
		"/api/v1/edge-clusters/{edgeClusterId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
