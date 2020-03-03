/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: MacAddressExpressions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package groups

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``action`` of method MacAddressExpressions#create.
const MacAddressExpressions_CREATE_ACTION_ADD = "add"
// Possible value for ``action`` of method MacAddressExpressions#create.
const MacAddressExpressions_CREATE_ACTION_REMOVE = "remove"




func macAddressExpressionsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["expression_id"] = bindings.NewStringType()
	fields["m_AC_address_list"] = bindings.NewReferenceType(model.MACAddressListBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_list"] = "MACAddressList"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func macAddressExpressionsCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func macAddressExpressionsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["expression_id"] = bindings.NewStringType()
	fields["m_AC_address_list"] = bindings.NewReferenceType(model.MACAddressListBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_list"] = "MACAddressList"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["m_AC_address_list"] = bindings.NewReferenceType(model.MACAddressListBindingType)
	paramsTypeMap["expression_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["expressionId"] = bindings.NewStringType()
	pathParams["expression_id"] = "expressionId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
	queryParams["action"] = "action"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"m_AC_address_list",
		"POST",
		"/policy/api/v1/global-infra/domains/{domainId}/groups/{groupId}/mac-address-expressions/{expressionId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func macAddressExpressionsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["expression_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func macAddressExpressionsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func macAddressExpressionsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["expression_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["expression_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["expressionId"] = bindings.NewStringType()
	pathParams["expression_id"] = "expressionId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/policy/api/v1/global-infra/domains/{domainId}/groups/{groupId}/mac-address-expressions/{expressionId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func macAddressExpressionsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["expression_id"] = bindings.NewStringType()
	fields["m_AC_address_expression"] = bindings.NewReferenceType(model.MACAddressExpressionBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_expression"] = "MACAddressExpression"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func macAddressExpressionsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func macAddressExpressionsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["expression_id"] = bindings.NewStringType()
	fields["m_AC_address_expression"] = bindings.NewReferenceType(model.MACAddressExpressionBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["expression_id"] = "ExpressionId"
	fieldNameMap["m_AC_address_expression"] = "MACAddressExpression"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["expression_id"] = bindings.NewStringType()
	paramsTypeMap["m_AC_address_expression"] = bindings.NewReferenceType(model.MACAddressExpressionBindingType)
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["expressionId"] = bindings.NewStringType()
	pathParams["expression_id"] = "expressionId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"m_AC_address_expression",
		"PATCH",
		"/policy/api/v1/global-infra/domains/{domainId}/groups/{groupId}/mac-address-expressions/{expressionId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


