// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Excludelist.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func excludelistAddmemberInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_reference"] = vapiBindings_.NewReferenceType(nsxModel.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExcludelistAddmemberOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ResourceReferenceBindingType)
}

func excludelistAddmemberRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_reference"] = vapiBindings_.NewReferenceType(nsxModel.ResourceReferenceBindingType)
	fieldNameMap["resource_reference"] = "ResourceReference"
	paramsTypeMap["resource_reference"] = vapiBindings_.NewReferenceType(nsxModel.ResourceReferenceBindingType)
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

func excludelistCheckifexistsInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = vapiBindings_.NewStringType()
	fields["deep_check"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["object_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExcludelistCheckifexistsOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ResourceReferenceBindingType)
}

func excludelistCheckifexistsRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = vapiBindings_.NewStringType()
	fields["deep_check"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["object_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	paramsTypeMap["deep_check"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["object_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["object_id"] = vapiBindings_.NewStringType()
	queryParams["deep_check"] = "deep_check"
	queryParams["object_type"] = "object_type"
	queryParams["object_id"] = "object_id"
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

func excludelistGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExcludelistGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ExcludeListBindingType)
}

func excludelistGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/firewall/excludelist",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func excludelistRemovememberInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = vapiBindings_.NewStringType()
	fields["deep_check"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["object_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExcludelistRemovememberOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ResourceReferenceBindingType)
}

func excludelistRemovememberRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["object_id"] = vapiBindings_.NewStringType()
	fields["deep_check"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["object_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["object_id"] = "ObjectId"
	fieldNameMap["deep_check"] = "DeepCheck"
	fieldNameMap["object_type"] = "ObjectType"
	paramsTypeMap["deep_check"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["object_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["object_id"] = vapiBindings_.NewStringType()
	queryParams["deep_check"] = "deep_check"
	queryParams["object_type"] = "object_type"
	queryParams["object_id"] = "object_id"
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

func excludelistUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["exclude_list"] = vapiBindings_.NewReferenceType(nsxModel.ExcludeListBindingType)
	fieldNameMap["exclude_list"] = "ExcludeList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ExcludelistUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.ExcludeListBindingType)
}

func excludelistUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["exclude_list"] = vapiBindings_.NewReferenceType(nsxModel.ExcludeListBindingType)
	fieldNameMap["exclude_list"] = "ExcludeList"
	paramsTypeMap["exclude_list"] = vapiBindings_.NewReferenceType(nsxModel.ExcludeListBindingType)
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
