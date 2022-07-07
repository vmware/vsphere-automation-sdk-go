// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ContainerClusters.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package fabric

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_PAS = "PAS"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_PKS = "PKS"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_KUBERNETES = "Kubernetes"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_OPENSHIFT = "Openshift"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_WCP = "WCP"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_WCP_GUEST = "WCP_Guest"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_AKS = "AKS"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_EKS = "EKS"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_TKGM = "TKGm"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_TKGI = "TKGi"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_GKE = "GKE"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_GARDENER = "Gardener"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_RANCHER = "Rancher"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_TAS = "TAS"

// Possible value for ``clusterType`` of method ContainerClusters#list.
const ContainerClusters_LIST_CLUSTER_TYPE_OTHER = "Other"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_VSPHERE = "vSphere"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_AWS = "AWS"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_AZURE = "Azure"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_GOOGLE = "Google"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_VMC = "VMC"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_KVM = "KVM"

// Possible value for ``infraType`` of method ContainerClusters#list.
const ContainerClusters_LIST_INFRA_TYPE_BAREMETAL = "Baremetal"

func containerClustersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["container_cluster_id"] = bindings.NewStringType()
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClustersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ContainerClusterBindingType)
}

func containerClustersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["container_cluster_id"] = bindings.NewStringType()
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	paramsTypeMap["container_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["containerClusterId"] = bindings.NewStringType()
	pathParams["container_cluster_id"] = "containerClusterId"
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
		"/api/v1/fabric/container-clusters/{containerClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClustersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["infra_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["scope_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_type"] = "ClusterType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["infra_type"] = "InfraType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["scope_id"] = "ScopeId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClustersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ContainerClusterListResultBindingType)
}

func containerClustersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["infra_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["scope_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_type"] = "ClusterType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["infra_type"] = "InfraType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["scope_id"] = "ScopeId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cluster_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["infra_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["scope_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["infra_type"] = "infra_type"
	queryParams["scope_id"] = "scope_id"
	queryParams["cluster_type"] = "cluster_type"
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
		"/api/v1/fabric/container-clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
