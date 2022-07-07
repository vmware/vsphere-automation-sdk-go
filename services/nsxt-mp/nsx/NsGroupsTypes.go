// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: NsGroups.
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

// Possible value for ``action`` of method NsGroups#addorremoveexpression.
const NsGroups_ADDORREMOVEEXPRESSION_ACTION_ADD_MEMBERS = "ADD_MEMBERS"

// Possible value for ``action`` of method NsGroups#addorremoveexpression.
const NsGroups_ADDORREMOVEEXPRESSION_ACTION_REMOVE_MEMBERS = "REMOVE_MEMBERS"

func nsGroupsAddorremoveexpressionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ns_group_id"] = bindings.NewStringType()
	fields["ns_group_expression_list"] = bindings.NewReferenceType(model.NSGroupExpressionListBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["ns_group_expression_list"] = "NsGroupExpressionList"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nsGroupsAddorremoveexpressionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NSGroupBindingType)
}

func nsGroupsAddorremoveexpressionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ns_group_id"] = bindings.NewStringType()
	fields["ns_group_expression_list"] = bindings.NewReferenceType(model.NSGroupExpressionListBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["ns_group_expression_list"] = "NsGroupExpressionList"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["ns_group_id"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["ns_group_expression_list"] = bindings.NewReferenceType(model.NSGroupExpressionListBindingType)
	paramsTypeMap["nsGroupId"] = bindings.NewStringType()
	pathParams["ns_group_id"] = "nsGroupId"
	queryParams["action"] = "action"
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
		"ns_group_expression_list",
		"POST",
		"/api/v1/ns-groups/{nsGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nsGroupsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ns_group"] = bindings.NewReferenceType(model.NSGroupBindingType)
	fieldNameMap["ns_group"] = "NsGroup"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nsGroupsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NSGroupBindingType)
}

func nsGroupsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ns_group"] = bindings.NewReferenceType(model.NSGroupBindingType)
	fieldNameMap["ns_group"] = "NsGroup"
	paramsTypeMap["ns_group"] = bindings.NewReferenceType(model.NSGroupBindingType)
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
		"ns_group",
		"POST",
		"/api/v1/ns-groups",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nsGroupsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ns_group_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nsGroupsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func nsGroupsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ns_group_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["ns_group_id"] = bindings.NewStringType()
	paramsTypeMap["nsGroupId"] = bindings.NewStringType()
	pathParams["ns_group_id"] = "nsGroupId"
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
		"",
		"",
		"DELETE",
		"/api/v1/ns-groups/{nsGroupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nsGroupsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ns_group_id"] = bindings.NewStringType()
	fields["populate_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["populate_references"] = "PopulateReferences"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nsGroupsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NSGroupBindingType)
}

func nsGroupsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ns_group_id"] = bindings.NewStringType()
	fields["populate_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["populate_references"] = "PopulateReferences"
	paramsTypeMap["ns_group_id"] = bindings.NewStringType()
	paramsTypeMap["populate_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["nsGroupId"] = bindings.NewStringType()
	pathParams["ns_group_id"] = "nsGroupId"
	queryParams["populate_references"] = "populate_references"
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
		"/api/v1/ns-groups/{nsGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nsGroupsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["member_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["populate_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["member_types"] = "MemberTypes"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["populate_references"] = "PopulateReferences"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nsGroupsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NSGroupListResultBindingType)
}

func nsGroupsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["member_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["populate_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["member_types"] = "MemberTypes"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["populate_references"] = "PopulateReferences"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["member_types"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["populate_references"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["populate_references"] = "populate_references"
	queryParams["member_types"] = "member_types"
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
		"/api/v1/ns-groups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nsGroupsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ns_group_id"] = bindings.NewStringType()
	fields["ns_group"] = bindings.NewReferenceType(model.NSGroupBindingType)
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["ns_group"] = "NsGroup"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nsGroupsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NSGroupBindingType)
}

func nsGroupsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ns_group_id"] = bindings.NewStringType()
	fields["ns_group"] = bindings.NewReferenceType(model.NSGroupBindingType)
	fieldNameMap["ns_group_id"] = "NsGroupId"
	fieldNameMap["ns_group"] = "NsGroup"
	paramsTypeMap["ns_group"] = bindings.NewReferenceType(model.NSGroupBindingType)
	paramsTypeMap["ns_group_id"] = bindings.NewStringType()
	paramsTypeMap["nsGroupId"] = bindings.NewStringType()
	pathParams["ns_group_id"] = "nsGroupId"
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
		"ns_group",
		"PUT",
		"/api/v1/ns-groups/{nsGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
