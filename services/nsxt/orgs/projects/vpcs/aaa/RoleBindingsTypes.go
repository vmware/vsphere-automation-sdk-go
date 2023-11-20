// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RoleBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package aaa

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``identitySourceType`` of method RoleBindings#delete.
const RoleBindings_DELETE_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#delete.
const RoleBindings_DELETE_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#delete.
const RoleBindings_DELETE_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

// Possible value for ``identitySourceType`` of method RoleBindings#get.
const RoleBindings_GET_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#get.
const RoleBindings_GET_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#get.
const RoleBindings_GET_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

// Possible value for ``identitySourceType`` of method RoleBindings#list.
const RoleBindings_LIST_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#list.
const RoleBindings_LIST_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#list.
const RoleBindings_LIST_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

func roleBindingsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RoleBindingsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func roleBindingsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["vpc_id"] = "vpcId"
	pathParams["binding_id"] = "bindingId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
	queryParams["path"] = "path"
	queryParams["root_path"] = "root_path"
	queryParams["role"] = "role"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["name"] = "name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["identity_source_type"] = "identity_source_type"
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
		"DELETE",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RoleBindingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RoleBindingBindingType)
}

func roleBindingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["vpc_id"] = "vpcId"
	pathParams["binding_id"] = "bindingId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
	queryParams["path"] = "path"
	queryParams["root_path"] = "root_path"
	queryParams["role"] = "role"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["name"] = "name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["identity_source_type"] = "identity_source_type"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RoleBindingsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RoleBindingListResultBindingType)
}

func roleBindingsListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_source_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["role"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["identity_source_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	pathParams["vpc_id"] = "vpcId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
	queryParams["path"] = "path"
	queryParams["root_path"] = "root_path"
	queryParams["role"] = "role"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["name"] = "name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["identity_source_type"] = "identity_source_type"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["project_id"] = vapiBindings_.NewStringType()
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["role_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.RoleBindingBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["role_binding"] = "RoleBinding"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RoleBindingsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.RoleBindingBindingType)
}

func roleBindingsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["vpc_id"] = vapiBindings_.NewStringType()
	fields["role_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.RoleBindingBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["role_binding"] = "RoleBinding"
	paramsTypeMap["role_binding"] = vapiBindings_.NewReferenceType(nsx_policyModel.RoleBindingBindingType)
	paramsTypeMap["project_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["projectId"] = vapiBindings_.NewStringType()
	paramsTypeMap["vpcId"] = vapiBindings_.NewStringType()
	pathParams["vpc_id"] = "vpcId"
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
		"role_binding",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
