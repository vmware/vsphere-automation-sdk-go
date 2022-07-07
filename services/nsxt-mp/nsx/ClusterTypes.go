// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Cluster.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``frameType`` of method Cluster#backuptoremote.
const Cluster_BACKUPTOREMOTE_FRAME_TYPE_GLOBAL_MANAGER = "GLOBAL_MANAGER"

// Possible value for ``frameType`` of method Cluster#backuptoremote.
const Cluster_BACKUPTOREMOTE_FRAME_TYPE_LOCAL_MANAGER = "LOCAL_MANAGER"

// Possible value for ``frameType`` of method Cluster#backuptoremote.
const Cluster_BACKUPTOREMOTE_FRAME_TYPE_LOCAL_LOCAL_MANAGER = "LOCAL_LOCAL_MANAGER"

// Possible value for ``frameType`` of method Cluster#backuptoremote.
const Cluster_BACKUPTOREMOTE_FRAME_TYPE_NSX_INTELLIGENCE = "NSX_INTELLIGENCE"

// Possible value for ``force`` of method Cluster#removenode.
const Cluster_REMOVENODE_FORCE_TRUE = "true"

// Possible value for ``force`` of method Cluster#removenode.
const Cluster_REMOVENODE_FORCE_FALSE = "false"

// Possible value for ``gracefulShutdown`` of method Cluster#removenode.
const Cluster_REMOVENODE_GRACEFUL_SHUTDOWN_TRUE = "true"

// Possible value for ``gracefulShutdown`` of method Cluster#removenode.
const Cluster_REMOVENODE_GRACEFUL_SHUTDOWN_FALSE = "false"

// Possible value for ``ignoreRepositoryIpCheck`` of method Cluster#removenode.
const Cluster_REMOVENODE_IGNORE_REPOSITORY_IP_CHECK_TRUE = "true"

// Possible value for ``ignoreRepositoryIpCheck`` of method Cluster#removenode.
const Cluster_REMOVENODE_IGNORE_REPOSITORY_IP_CHECK_FALSE = "false"

func clusterBackuptoremoteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["site_id"] = "SiteId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterBackuptoremoteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterBackuptoremoteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["frame_type"] = "FrameType"
	fieldNameMap["site_id"] = "SiteId"
	paramsTypeMap["frame_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["site_id"] = "site_id"
	queryParams["frame_type"] = "frame_type"
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
		"action=backup_to_remote",
		"",
		"POST",
		"/api/v1/cluster",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"/api/v1/cluster/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"/api/v1/cluster/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterConfigBindingType)
}

func clusterGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/cluster",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterGet0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterGet0OutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterNodeInfoBindingType)
}

func clusterGet0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fieldNameMap["node_id"] = "NodeId"
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
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
		"/api/v1/cluster/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterGet1InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterGet1OutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterGet1RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"/api/v1/cluster/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterJoinclusterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["join_cluster_parameters"] = bindings.NewReferenceType(model.JoinClusterParametersBindingType)
	fieldNameMap["join_cluster_parameters"] = "JoinClusterParameters"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterJoinclusterOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterConfigurationBindingType)
}

func clusterJoinclusterRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["join_cluster_parameters"] = bindings.NewReferenceType(model.JoinClusterParametersBindingType)
	fieldNameMap["join_cluster_parameters"] = "JoinClusterParameters"
	paramsTypeMap["join_cluster_parameters"] = bindings.NewReferenceType(model.JoinClusterParametersBindingType)
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
		"action=join_cluster",
		"join_cluster_parameters",
		"POST",
		"/api/v1/cluster",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterRemovenodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["graceful_shutdown"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ignore_repository_ip_check"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["force"] = "Force"
	fieldNameMap["graceful_shutdown"] = "GracefulShutdown"
	fieldNameMap["ignore_repository_ip_check"] = "IgnoreRepositoryIpCheck"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterRemovenodeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterConfigurationBindingType)
}

func clusterRemovenodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["graceful_shutdown"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ignore_repository_ip_check"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["force"] = "Force"
	fieldNameMap["graceful_shutdown"] = "GracefulShutdown"
	fieldNameMap["ignore_repository_ip_check"] = "IgnoreRepositoryIpCheck"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ignore_repository_ip_check"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["graceful_shutdown"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_id"] = bindings.NewStringType()
	paramsTypeMap["nodeId"] = bindings.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["graceful_shutdown"] = "graceful-shutdown"
	queryParams["ignore_repository_ip_check"] = "ignore-repository-ip-check"
	queryParams["force"] = "force"
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
		"action=remove_node",
		"",
		"POST",
		"/api/v1/cluster/{nodeId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterSummarizeinventorytoremoteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterSummarizeinventorytoremoteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterSummarizeinventorytoremoteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"action=summarize_inventory_to_remote",
		"",
		"POST",
		"/api/v1/cluster",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["target_node_id"] = bindings.NewStringType()
	fields["target_uri"] = bindings.NewStringType()
	fieldNameMap["target_node_id"] = "TargetNodeId"
	fieldNameMap["target_uri"] = "TargetUri"
	paramsTypeMap["target_uri"] = bindings.NewStringType()
	paramsTypeMap["target_node_id"] = bindings.NewStringType()
	paramsTypeMap["targetNodeId"] = bindings.NewStringType()
	paramsTypeMap["targetUri"] = bindings.NewStringType()
	pathParams["target_uri"] = "targetUri"
	pathParams["target_node_id"] = "targetNodeId"
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
		"PUT",
		"/api/v1/cluster/{targetNodeId}/{targetUri}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
