// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ContainerClusterSpan.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package security_policies

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func containerClusterSpanDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["antrea_cluster1"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClusterSpanDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func containerClusterSpanDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["antrea_cluster1"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["security_policy_id"] = bindings.NewStringType()
	paramsTypeMap["antrea_cluster1"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["securityPolicyId"] = bindings.NewStringType()
	paramsTypeMap["antreaCluster1"] = bindings.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["antrea_cluster1"] = "antreaCluster1"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span/{antreaCluster1}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["antrea_cluster1"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClusterSpanGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
}

func containerClusterSpanGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["antrea_cluster1"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["security_policy_id"] = bindings.NewStringType()
	paramsTypeMap["antrea_cluster1"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["securityPolicyId"] = bindings.NewStringType()
	paramsTypeMap["antreaCluster1"] = bindings.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["antrea_cluster1"] = "antreaCluster1"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span/{antreaCluster1}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClusterSpanListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SecurityPolicyContainerClusterListResultBindingType)
}

func containerClusterSpanListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["security_policy_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["securityPolicyId"] = bindings.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["domain_id"] = "domainId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["container_cluster_id"] = bindings.NewStringType()
	fields["security_policy_container_cluster"] = bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClusterSpanPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func containerClusterSpanPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["container_cluster_id"] = bindings.NewStringType()
	fields["security_policy_container_cluster"] = bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["container_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["security_policy_container_cluster"] = bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
	paramsTypeMap["security_policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["securityPolicyId"] = bindings.NewStringType()
	paramsTypeMap["containerClusterId"] = bindings.NewStringType()
	pathParams["container_cluster_id"] = "containerClusterId"
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["domain_id"] = "domainId"
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
		"security_policy_container_cluster",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span/{containerClusterId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["container_cluster_id"] = bindings.NewStringType()
	fields["security_policy_container_cluster"] = bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func containerClusterSpanUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
}

func containerClusterSpanUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["security_policy_id"] = bindings.NewStringType()
	fields["container_cluster_id"] = bindings.NewStringType()
	fields["security_policy_container_cluster"] = bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["container_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["security_policy_container_cluster"] = bindings.NewReferenceType(model.SecurityPolicyContainerClusterBindingType)
	paramsTypeMap["security_policy_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["securityPolicyId"] = bindings.NewStringType()
	paramsTypeMap["containerClusterId"] = bindings.NewStringType()
	pathParams["container_cluster_id"] = "containerClusterId"
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["domain_id"] = "domainId"
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
		"security_policy_container_cluster",
		"PUT",
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span/{containerClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
