// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Topology.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nvds_urt

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func topologyApplyInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upgrade_topology"] = vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
	fields["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["use_recommended_topology_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["upgrade_topology"] = "UpgradeTopology"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["use_recommended_topology_config"] = "UseRecommendedTopologyConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TopologyApplyOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
}

func topologyApplyRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["upgrade_topology"] = vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
	fields["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["use_recommended_topology_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["upgrade_topology"] = "UpgradeTopology"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["use_recommended_topology_config"] = "UseRecommendedTopologyConfig"
	paramsTypeMap["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["use_recommended_topology_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["upgrade_topology"] = vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
	queryParams["cluster_id"] = "cluster_id"
	queryParams["use_recommended_topology_config"] = "use_recommended_topology_config"
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

func topologyGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["precheck_id"] = vapiBindings_.NewStringType()
	fields["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["compute_manager_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["show_vds_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["precheck_id"] = "PrecheckId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_manager_id"] = "ComputeManagerId"
	fieldNameMap["show_vds_config"] = "ShowVdsConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TopologyGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
}

func topologyGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["precheck_id"] = vapiBindings_.NewStringType()
	fields["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["compute_manager_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["show_vds_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["precheck_id"] = "PrecheckId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_manager_id"] = "ComputeManagerId"
	fieldNameMap["show_vds_config"] = "ShowVdsConfig"
	paramsTypeMap["compute_manager_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["show_vds_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["precheck_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["precheckId"] = vapiBindings_.NewStringType()
	pathParams["precheck_id"] = "precheckId"
	queryParams["compute_manager_id"] = "compute_manager_id"
	queryParams["cluster_id"] = "cluster_id"
	queryParams["show_vds_config"] = "show_vds_config"
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
		"/api/v1/nvds-urt/topology/{precheckId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func topologyRevertInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["upgrade_topology"] = vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
	fieldNameMap["upgrade_topology"] = "UpgradeTopology"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TopologyRevertOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
}

func topologyRevertRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["upgrade_topology"] = vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
	fieldNameMap["upgrade_topology"] = "UpgradeTopology"
	paramsTypeMap["upgrade_topology"] = vapiBindings_.NewReferenceType(nsxModel.UpgradeTopologyBindingType)
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
		"action=revert",
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
