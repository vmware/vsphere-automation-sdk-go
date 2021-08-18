// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cluster_control_planes

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["enforcementpoint_id"] = bindings.NewStringType()
	fields["cluster_control_plane_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["cluster_control_plane_id"] = "ClusterControlPlaneId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AntreaClusterInfoBindingType)
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
	fields["site_id"] = bindings.NewStringType()
	fields["enforcementpoint_id"] = bindings.NewStringType()
	fields["cluster_control_plane_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcementpoint_id"] = "EnforcementpointId"
	fieldNameMap["cluster_control_plane_id"] = "ClusterControlPlaneId"
	paramsTypeMap["enforcementpoint_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["cluster_control_plane_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["enforcementpointId"] = bindings.NewStringType()
	paramsTypeMap["clusterControlPlaneId"] = bindings.NewStringType()
	pathParams["enforcementpoint_id"] = "enforcementpointId"
	pathParams["site_id"] = "siteId"
	pathParams["cluster_control_plane_id"] = "clusterControlPlaneId"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementpointId}/cluster-control-planes/{clusterControlPlaneId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
