/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CloudConnectionState.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package edge_nodes

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func cloudConnectionStateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_node_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudConnectionStateGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CloudConnectionStateBindingType)
}

func cloudConnectionStateGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["enforcement_point_id"] = bindings.NewStringType()
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_node_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_node_id"] = "EdgeNodeId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["enforcement_point_id"] = bindings.NewStringType()
	paramsTypeMap["edge_node_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["enforcementPointId"] = bindings.NewStringType()
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	paramsTypeMap["edgeNodeId"] = bindings.NewStringType()
	pathParams["edge_node_id"] = "edgeNodeId"
	pathParams["site_id"] = "siteId"
	pathParams["edge_cluster_id"] = "edgeClusterId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/global-infra/sites/{siteId}/enforcement-points/{enforcementPointId}/edge-clusters/{edgeClusterId}/edge-nodes/{edgeNodeId}/cloud-connection-state",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


