// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Excludelist.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``objectType`` of method Excludelist#checkifexists.
const Excludelist_CHECKIFEXISTS_OBJECT_TYPE_NSGROUP = "NSGroup"

// Possible value for ``objectType`` of method Excludelist#checkifexists.
const Excludelist_CHECKIFEXISTS_OBJECT_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``objectType`` of method Excludelist#checkifexists.
const Excludelist_CHECKIFEXISTS_OBJECT_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``objectType`` of method Excludelist#removemember.
const Excludelist_REMOVEMEMBER_OBJECT_TYPE_NSGROUP = "NSGroup"

// Possible value for ``objectType`` of method Excludelist#removemember.
const Excludelist_REMOVEMEMBER_OBJECT_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``objectType`` of method Excludelist#removemember.
const Excludelist_REMOVEMEMBER_OBJECT_TYPE_LOGICALSWITCH = "LogicalSwitch"

func excludelistAddmemberInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func excludelistAddmemberOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func excludelistAddmemberRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	paramsTypeMap["resource_reference"] = bindings.NewReferenceType(model.ResourceReferenceBindingType)
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
		"action=add_member",
		"resource_reference",
		"POST",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func excludelistCheckifexistsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func excludelistCheckifexistsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func excludelistCheckifexistsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	paramsTypeMap["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["object_id"] = bindings.NewStringType()
	paramsTypeMap["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["deep_check"] = "deep_check"
	queryParams["object_type"] = "object_type"
	queryParams["object_id"] = "object_id"
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
		"action=check_if_exists",
		"",
		"POST",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func excludelistGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func excludelistGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExcludeListBindingType)
}

func excludelistGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func excludelistRemovememberInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func excludelistRemovememberOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ResourceReferenceBindingType)
}

func excludelistRemovememberRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = bindings.NewStringType()
	fields["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	paramsTypeMap["object_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["object_id"] = bindings.NewStringType()
	paramsTypeMap["deep_check"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["deep_check"] = "deep_check"
	queryParams["object_type"] = "object_type"
	queryParams["object_id"] = "object_id"
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
		"action=remove_member",
		"",
		"POST",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func excludelistUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["exclude_list"] = bindings.NewReferenceType(model.ExcludeListBindingType)
	fieldNameMap["exclude_list"] = "ExcludeList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func excludelistUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ExcludeListBindingType)
}

func excludelistUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["exclude_list"] = bindings.NewReferenceType(model.ExcludeListBindingType)
	fieldNameMap["exclude_list"] = "ExcludeList"
	paramsTypeMap["exclude_list"] = bindings.NewReferenceType(model.ExcludeListBindingType)
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
		"exclude_list",
		"PUT",
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
