// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Cluster.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package idfw

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func clusterDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
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
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["cluster_id"] = "clusterId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/settings/firewall/idfw/cluster/{clusterId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
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
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["cluster_id"] = "clusterId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/settings/firewall/idfw/cluster/{clusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationListResultBindingType)
}

func clusterListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/settings/firewall/idfw/cluster",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clusterPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["compute_cluster_idfw_configuration"] = bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_cluster_idfw_configuration"] = "ComputeClusterIdfwConfiguration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clusterPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["compute_cluster_idfw_configuration"] = bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_cluster_idfw_configuration"] = "ComputeClusterIdfwConfiguration"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["compute_cluster_idfw_configuration"] = bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["cluster_id"] = "clusterId"
	pathParams["org_id"] = "orgId"
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
		"compute_cluster_idfw_configuration",
		"PATCH",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/settings/firewall/idfw/cluster/{clusterId}",
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
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["compute_cluster_idfw_configuration"] = bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_cluster_idfw_configuration"] = "ComputeClusterIdfwConfiguration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
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
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["cluster_id"] = bindings.NewStringType()
	fields["compute_cluster_idfw_configuration"] = bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["compute_cluster_idfw_configuration"] = "ComputeClusterIdfwConfiguration"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["compute_cluster_idfw_configuration"] = bindings.NewReferenceType(model.ComputeClusterIdfwConfigurationBindingType)
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["cluster_id"] = "clusterId"
	pathParams["org_id"] = "orgId"
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
		"compute_cluster_idfw_configuration",
		"PUT",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/settings/firewall/idfw/cluster/{clusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
