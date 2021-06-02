// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkTransportTransportNodes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func managementPlaneAPINetworkTransportTransportNodesCreateNvdsUpgradePrecheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesCreateNvdsUpgradePrecheckOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradePrecheckIdBindingType)
}

func managementPlaneAPINetworkTransportTransportNodesCreateNvdsUpgradePrecheckRestMetadata() protocol.OperationRestMetadata {
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
		"POST",
		"/api/v1/nvds-urt/precheck",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesCreateNvdsUpgradePrecheckByClustersInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["precheck_parameters"] = bindings.NewReferenceType(model.PrecheckParametersBindingType)
	fieldNameMap["precheck_parameters"] = "PrecheckParameters"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesCreateNvdsUpgradePrecheckByClustersOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradePrecheckIdBindingType)
}

func managementPlaneAPINetworkTransportTransportNodesCreateNvdsUpgradePrecheckByClustersRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["precheck_parameters"] = bindings.NewReferenceType(model.PrecheckParametersBindingType)
	fieldNameMap["precheck_parameters"] = "PrecheckParameters"
	paramsTypeMap["precheck_parameters"] = bindings.NewReferenceType(model.PrecheckParametersBindingType)
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
		"precheck_parameters",
		"POST",
		"/api/v1/nvds-urt/precheck-by-clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesGetNvdsUpgradePrecheckIdInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesGetNvdsUpgradePrecheckIdOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradePrecheckIdBindingType)
}

func managementPlaneAPINetworkTransportTransportNodesGetNvdsUpgradePrecheckIdRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/nvds-urt/precheck",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesGetNvdsUpgradeReadinessCheckSummaryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["precheck_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["precheck_id"] = "PrecheckId"
	fieldNameMap["cluster_id"] = "ClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesGetNvdsUpgradeReadinessCheckSummaryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradeStatusSummaryBindingType)
}

func managementPlaneAPINetworkTransportTransportNodesGetNvdsUpgradeReadinessCheckSummaryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["precheck_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["precheck_id"] = "PrecheckId"
	fieldNameMap["cluster_id"] = "ClusterId"
	paramsTypeMap["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["precheck_id"] = bindings.NewStringType()
	paramsTypeMap["precheckId"] = bindings.NewStringType()
	pathParams["precheck_id"] = "precheckId"
	queryParams["cluster_id"] = "cluster_id"
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
		"/api/v1/nvds-urt/status-summary/{precheckId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesGetRecommendedVdsTopologyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["precheck_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["compute_manager_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["show_vds_config"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["precheck_id"] = "PrecheckId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_manager_id"] = "ComputeManagerId"
	fieldNameMap["show_vds_config"] = "ShowVdsConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesGetRecommendedVdsTopologyOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeTopologyBindingType)
}

func managementPlaneAPINetworkTransportTransportNodesGetRecommendedVdsTopologyRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["precheck_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["compute_manager_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["show_vds_config"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["precheck_id"] = "PrecheckId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_manager_id"] = "ComputeManagerId"
	fieldNameMap["show_vds_config"] = "ShowVdsConfig"
	paramsTypeMap["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["precheck_id"] = bindings.NewStringType()
	paramsTypeMap["show_vds_config"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["compute_manager_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["precheckId"] = bindings.NewStringType()
	pathParams["precheck_id"] = "precheckId"
	queryParams["compute_manager_id"] = "compute_manager_id"
	queryParams["cluster_id"] = "cluster_id"
	queryParams["show_vds_config"] = "show_vds_config"
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
		"/api/v1/nvds-urt/topology/{precheckId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesIgnoreMigrateStatusIgnoreMigrateStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesIgnoreMigrateStatusIgnoreMigrateStatusOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkTransportTransportNodesIgnoreMigrateStatusIgnoreMigrateStatusRestMetadata() protocol.OperationRestMetadata {
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
		"action=ignore_migrate_status",
		"",
		"POST",
		"/api/v1/nvds-urt",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesMigrateTransportNodeFromNvdsToVdsMigrateToVdsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesMigrateTransportNodeFromNvdsToVdsMigrateToVdsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkTransportTransportNodesMigrateTransportNodeFromNvdsToVdsMigrateToVdsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeId"] = bindings.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
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
		"action=migrate_to_vds",
		"",
		"POST",
		"/api/v1/transport-nodes/{transportNodeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesNvdsUpgradeCleanupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesNvdsUpgradeCleanupOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkTransportTransportNodesNvdsUpgradeCleanupRestMetadata() protocol.OperationRestMetadata {
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
		"action=cleanup",
		"",
		"POST",
		"/api/v1/nvds-urt",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkTransportTransportNodesSetTargetVdsTopologyApplyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upgrade_topology"] = bindings.NewReferenceType(model.UpgradeTopologyBindingType)
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["use_recommended_topology_config"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["upgrade_topology"] = "UpgradeTopology"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["use_recommended_topology_config"] = "UseRecommendedTopologyConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkTransportTransportNodesSetTargetVdsTopologyApplyOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeTopologyBindingType)
}

func managementPlaneAPINetworkTransportTransportNodesSetTargetVdsTopologyApplyRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["upgrade_topology"] = bindings.NewReferenceType(model.UpgradeTopologyBindingType)
	fields["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["use_recommended_topology_config"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["upgrade_topology"] = "UpgradeTopology"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["use_recommended_topology_config"] = "UseRecommendedTopologyConfig"
	paramsTypeMap["cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["use_recommended_topology_config"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["upgrade_topology"] = bindings.NewReferenceType(model.UpgradeTopologyBindingType)
	queryParams["cluster_id"] = "cluster_id"
	queryParams["use_recommended_topology_config"] = "use_recommended_topology_config"
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
		"action=apply",
		"upgrade_topology",
		"POST",
		"/api/v1/nvds-urt/topology",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
