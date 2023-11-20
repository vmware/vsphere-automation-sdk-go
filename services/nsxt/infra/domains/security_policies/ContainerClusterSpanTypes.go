// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ContainerClusterSpan.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package security_policies

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func containerClusterSpanDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["antrea_cluster1"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContainerClusterSpanDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func containerClusterSpanDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["antrea_cluster1"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["antrea_cluster1"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["antreaCluster1"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["antrea_cluster1"] = "antreaCluster1"
	pathParams["domain_id"] = "domainId"
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
		"DELETE",
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span/{antreaCluster1}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["antrea_cluster1"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContainerClusterSpanGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
}

func containerClusterSpanGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["antrea_cluster1"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["antrea_cluster1"] = "AntreaCluster1"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["antrea_cluster1"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["antreaCluster1"] = vapiBindings_.NewStringType()
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["antrea_cluster1"] = "antreaCluster1"
	pathParams["domain_id"] = "domainId"
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span/{antreaCluster1}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContainerClusterSpanListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterListResultBindingType)
}

func containerClusterSpanListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
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
		"/policy/api/v1/infra/domains/{domainId}/security-policies/{securityPolicyId}/container-cluster-span",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func containerClusterSpanPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["container_cluster_id"] = vapiBindings_.NewStringType()
	fields["security_policy_container_cluster"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContainerClusterSpanPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func containerClusterSpanPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["container_cluster_id"] = vapiBindings_.NewStringType()
	fields["security_policy_container_cluster"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["container_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy_container_cluster"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["containerClusterId"] = vapiBindings_.NewStringType()
	pathParams["container_cluster_id"] = "containerClusterId"
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["domain_id"] = "domainId"
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

func containerClusterSpanUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["container_cluster_id"] = vapiBindings_.NewStringType()
	fields["security_policy_container_cluster"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ContainerClusterSpanUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
}

func containerClusterSpanUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["security_policy_id"] = vapiBindings_.NewStringType()
	fields["container_cluster_id"] = vapiBindings_.NewStringType()
	fields["security_policy_container_cluster"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["security_policy_id"] = "SecurityPolicyId"
	fieldNameMap["container_cluster_id"] = "ContainerClusterId"
	fieldNameMap["security_policy_container_cluster"] = "SecurityPolicyContainerCluster"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["container_cluster_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["security_policy_container_cluster"] = vapiBindings_.NewReferenceType(nsx_policyModel.SecurityPolicyContainerClusterBindingType)
	paramsTypeMap["security_policy_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["securityPolicyId"] = vapiBindings_.NewStringType()
	paramsTypeMap["containerClusterId"] = vapiBindings_.NewStringType()
	pathParams["container_cluster_id"] = "containerClusterId"
	pathParams["security_policy_id"] = "securityPolicyId"
	pathParams["domain_id"] = "domainId"
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
