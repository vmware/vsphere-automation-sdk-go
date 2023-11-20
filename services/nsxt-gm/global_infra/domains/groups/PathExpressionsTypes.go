// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PathExpressions.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package groups

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

// Possible value for ``action`` of method PathExpressions#create.
const PathExpressions_CREATE_ACTION_ADD = "add"

// Possible value for ``action`` of method PathExpressions#create.
const PathExpressions_CREATE_ACTION_REMOVE = "remove"

func pathExpressionsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fields["group_member_list"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.GroupMemberListBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["group_member_list"] = "GroupMemberList"
	fieldNameMap["action"] = "Action"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PathExpressionsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func pathExpressionsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fields["group_member_list"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.GroupMemberListBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["group_member_list"] = "GroupMemberList"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["group_member_list"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.GroupMemberListBindingType)
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["action"] = vapiBindings_.NewStringType()
	paramsTypeMap["expression_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["expressionId"] = vapiBindings_.NewStringType()
	pathParams["expression_id"] = "expressionId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
	queryParams["action"] = "action"
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
		"group_member_list",
		"POST",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/path-expressions/{expressionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func pathExpressionsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PathExpressionsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func pathExpressionsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["expression_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["expressionId"] = vapiBindings_.NewStringType()
	pathParams["expression_id"] = "expressionId"
	pathParams["group_id"] = "groupId"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/path-expressions/{expressionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func pathExpressionsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fields["path_expression"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PathExpressionBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["path_expression"] = "PathExpression"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func PathExpressionsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func pathExpressionsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fields["path_expression"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PathExpressionBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["path_expression"] = "PathExpression"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["path_expression"] = vapiBindings_.NewReferenceType(nsx_global_policyModel.PathExpressionBindingType)
	paramsTypeMap["expression_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["domainId"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["expressionId"] = vapiBindings_.NewStringType()
	pathParams["expression_id"] = "expressionId"
	pathParams["group_id"] = "groupId"
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
		"path_expression",
		"PATCH",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/path-expressions/{expressionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
