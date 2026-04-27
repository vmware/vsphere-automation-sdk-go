// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CentralizedNetworkAttachments.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package projects

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func centralizedNetworkAttachmentsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CentralizedNetworkAttachmentsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func centralizedNetworkAttachmentsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cna_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["cnaId"] = vapiBindings_.NewStringType()
	pathParams["cna_id"] = "cnaId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/centralized-network-attachments/{cnaId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func centralizedNetworkAttachmentsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CentralizedNetworkAttachmentsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
}

func centralizedNetworkAttachmentsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["cna_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["cnaId"] = vapiBindings_.NewStringType()
	pathParams["cna_id"] = "cnaId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/centralized-network-attachments/{cnaId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func centralizedNetworkAttachmentsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CentralizedNetworkAttachmentsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentListResultBindingType)
}

func centralizedNetworkAttachmentsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["include_mark_for_delete_objects"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/centralized-network-attachments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func centralizedNetworkAttachmentsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fields["centralized_network_attachment"] = vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	fieldNameMap["centralized_network_attachment"] = "CentralizedNetworkAttachment"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CentralizedNetworkAttachmentsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func centralizedNetworkAttachmentsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fields["centralized_network_attachment"] = vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	fieldNameMap["centralized_network_attachment"] = "CentralizedNetworkAttachment"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["centralized_network_attachment"] = vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
	paramsTypeMap["cna_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["cnaId"] = vapiBindings_.NewStringType()
	pathParams["cna_id"] = "cnaId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"centralized_network_attachment",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/centralized-network-attachments/{cnaId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func centralizedNetworkAttachmentsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fields["centralized_network_attachment"] = vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	fieldNameMap["centralized_network_attachment"] = "CentralizedNetworkAttachment"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CentralizedNetworkAttachmentsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
}

func centralizedNetworkAttachmentsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["cna_id"] = vapiBindings_.NewStringType()
	fields["centralized_network_attachment"] = vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cna_id"] = "CnaId"
	fieldNameMap["centralized_network_attachment"] = "CentralizedNetworkAttachment"
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["centralized_network_attachment"] = vapiBindings_.NewReferenceType(nsx_policyModel.CentralizedNetworkAttachmentBindingType)
	paramsTypeMap["cna_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["cnaId"] = vapiBindings_.NewStringType()
	pathParams["cna_id"] = "cnaId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"centralized_network_attachment",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/centralized-network-attachments/{cnaId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
