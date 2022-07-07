// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RoleBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
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

// Possible value for ``identitySourceType`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#deletestalebindings.
const RoleBindings_DELETESTALEBINDINGS_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

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

func roleBindingsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role_binding"] = bindings.NewReferenceType(model.RoleBindingBindingType)
	fieldNameMap["role_binding"] = "RoleBinding"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleBindingBindingType)
}

func roleBindingsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["role_binding"] = bindings.NewReferenceType(model.RoleBindingBindingType)
	fieldNameMap["role_binding"] = "RoleBinding"
	paramsTypeMap["role_binding"] = bindings.NewReferenceType(model.RoleBindingBindingType)
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
		"role_binding",
		"POST",
		"/global-manager/api/v1/aaa/role-bindings",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func roleBindingsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["binding_id"] = "bindingId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
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
		"/global-manager/api/v1/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsDeletestalebindingsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsDeletestalebindingsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func roleBindingsDeletestalebindingsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
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
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=delete_stale_bindings",
		"",
		"POST",
		"/global-manager/api/v1/aaa/role-bindings",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleBindingBindingType)
}

func roleBindingsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["binding_id"] = "bindingId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
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
		"/global-manager/api/v1/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleBindingListResultBindingType)
}

func roleBindingsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["role"] = "Role"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
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
		"/global-manager/api/v1/aaa/role-bindings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["binding_id"] = bindings.NewStringType()
	fields["role_binding"] = bindings.NewReferenceType(model.RoleBindingBindingType)
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["role_binding"] = "RoleBinding"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleBindingBindingType)
}

func roleBindingsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["binding_id"] = bindings.NewStringType()
	fields["role_binding"] = bindings.NewReferenceType(model.RoleBindingBindingType)
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["role_binding"] = "RoleBinding"
	paramsTypeMap["role_binding"] = bindings.NewReferenceType(model.RoleBindingBindingType)
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["binding_id"] = "bindingId"
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
		"role_binding",
		"PUT",
		"/global-manager/api/v1/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
