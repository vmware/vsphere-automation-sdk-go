// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MacAddressExpressions.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package groups

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``action`` of method MacAddressExpressions#create.
const MacAddressExpressions_CREATE_ACTION_ADD = "add"

// Possible value for ``action`` of method MacAddressExpressions#create.
const MacAddressExpressions_CREATE_ACTION_REMOVE = "remove"

func macAddressExpressionsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fields["m_AC_address_list"] = vapiBindings_.NewReferenceType(nsx_policyModel.MACAddressListBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_list"] = "MACAddressList"
	fieldNameMap["action"] = "Action"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MacAddressExpressionsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func macAddressExpressionsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["m_AC_address_list"] = vapiBindings_.NewReferenceType(nsx_policyModel.MACAddressListBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_list"] = "MACAddressList"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["m_AC_address_list"] = vapiBindings_.NewReferenceType(nsx_policyModel.MACAddressListBindingType)
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
		"m_AC_address_list",
		"POST",
		"/policy/api/v1/infra/domains/{domainId}/groups/{groupId}/mac-address-expressions/{expressionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func macAddressExpressionsDeleteInputType() vapiBindings_.StructType {
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

func MacAddressExpressionsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func macAddressExpressionsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/policy/api/v1/infra/domains/{domainId}/groups/{groupId}/mac-address-expressions/{expressionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func macAddressExpressionsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = vapiBindings_.NewStringType()
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["expression_id"] = vapiBindings_.NewStringType()
	fields["m_AC_address_expression"] = vapiBindings_.NewReferenceType(nsx_policyModel.MACAddressExpressionBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_expression"] = "MACAddressExpression"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MacAddressExpressionsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func macAddressExpressionsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["m_AC_address_expression"] = vapiBindings_.NewReferenceType(nsx_policyModel.MACAddressExpressionBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_expression"] = "MACAddressExpression"
	paramsTypeMap["domain_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["m_AC_address_expression"] = vapiBindings_.NewReferenceType(nsx_policyModel.MACAddressExpressionBindingType)
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
		"m_AC_address_expression",
		"PATCH",
		"/policy/api/v1/infra/domains/{domainId}/groups/{groupId}/mac-address-expressions/{expressionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
