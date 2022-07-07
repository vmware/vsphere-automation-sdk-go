// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: VirtualServers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package loadbalancer

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func virtualServersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lb_virtual_server"] = bindings.NewReferenceType(model.LbVirtualServerBindingType)
	fieldNameMap["lb_virtual_server"] = "LbVirtualServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbVirtualServerBindingType)
}

func virtualServersCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lb_virtual_server"] = bindings.NewReferenceType(model.LbVirtualServerBindingType)
	fieldNameMap["lb_virtual_server"] = "LbVirtualServer"
	paramsTypeMap["lb_virtual_server"] = bindings.NewReferenceType(model.LbVirtualServerBindingType)
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
		"lb_virtual_server",
		"POST",
		"/api/v1/loadbalancer/virtual-servers",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualServersCreatewithrulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lb_virtual_server_with_rule"] = bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
	fieldNameMap["lb_virtual_server_with_rule"] = "LbVirtualServerWithRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersCreatewithrulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
}

func virtualServersCreatewithrulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lb_virtual_server_with_rule"] = bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
	fieldNameMap["lb_virtual_server_with_rule"] = "LbVirtualServerWithRule"
	paramsTypeMap["lb_virtual_server_with_rule"] = bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
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
		"action=create_with_rules",
		"lb_virtual_server_with_rule",
		"POST",
		"/api/v1/loadbalancer/virtual-servers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualServersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_server_id"] = bindings.NewStringType()
	fields["delete_associated_rules"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["delete_associated_rules"] = "DeleteAssociatedRules"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func virtualServersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_server_id"] = bindings.NewStringType()
	fields["delete_associated_rules"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["delete_associated_rules"] = "DeleteAssociatedRules"
	paramsTypeMap["delete_associated_rules"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["virtual_server_id"] = bindings.NewStringType()
	paramsTypeMap["virtualServerId"] = bindings.NewStringType()
	pathParams["virtual_server_id"] = "virtualServerId"
	queryParams["delete_associated_rules"] = "delete_associated_rules"
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
		"/api/v1/loadbalancer/virtual-servers/{virtualServerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualServersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_server_id"] = bindings.NewStringType()
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbVirtualServerBindingType)
}

func virtualServersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_server_id"] = bindings.NewStringType()
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	paramsTypeMap["virtual_server_id"] = bindings.NewStringType()
	paramsTypeMap["virtualServerId"] = bindings.NewStringType()
	pathParams["virtual_server_id"] = "virtualServerId"
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
		"/api/v1/loadbalancer/virtual-servers/{virtualServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualServersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbVirtualServerListResultBindingType)
}

func virtualServersListRestMetadata() protocol.OperationRestMetadata {
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
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
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
		"/api/v1/loadbalancer/virtual-servers",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualServersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_server_id"] = bindings.NewStringType()
	fields["lb_virtual_server"] = bindings.NewReferenceType(model.LbVirtualServerBindingType)
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["lb_virtual_server"] = "LbVirtualServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbVirtualServerBindingType)
}

func virtualServersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_server_id"] = bindings.NewStringType()
	fields["lb_virtual_server"] = bindings.NewReferenceType(model.LbVirtualServerBindingType)
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["lb_virtual_server"] = "LbVirtualServer"
	paramsTypeMap["lb_virtual_server"] = bindings.NewReferenceType(model.LbVirtualServerBindingType)
	paramsTypeMap["virtual_server_id"] = bindings.NewStringType()
	paramsTypeMap["virtualServerId"] = bindings.NewStringType()
	pathParams["virtual_server_id"] = "virtualServerId"
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
		"lb_virtual_server",
		"PUT",
		"/api/v1/loadbalancer/virtual-servers/{virtualServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func virtualServersUpdatewithrulesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["virtual_server_id"] = bindings.NewStringType()
	fields["lb_virtual_server_with_rule"] = bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["lb_virtual_server_with_rule"] = "LbVirtualServerWithRule"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func virtualServersUpdatewithrulesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
}

func virtualServersUpdatewithrulesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["virtual_server_id"] = bindings.NewStringType()
	fields["lb_virtual_server_with_rule"] = bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
	fieldNameMap["virtual_server_id"] = "VirtualServerId"
	fieldNameMap["lb_virtual_server_with_rule"] = "LbVirtualServerWithRule"
	paramsTypeMap["lb_virtual_server_with_rule"] = bindings.NewReferenceType(model.LbVirtualServerWithRuleBindingType)
	paramsTypeMap["virtual_server_id"] = bindings.NewStringType()
	paramsTypeMap["virtualServerId"] = bindings.NewStringType()
	pathParams["virtual_server_id"] = "virtualServerId"
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
		"action=update_with_rules",
		"lb_virtual_server_with_rule",
		"PUT",
		"/api/v1/loadbalancer/virtual-servers/{virtualServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
