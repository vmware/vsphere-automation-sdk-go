// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Ipv6StaticBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package servers

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func ipv6StaticBindingsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["dhcp_v6_static_binding"] = vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func Ipv6StaticBindingsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
}

func ipv6StaticBindingsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["dhcp_v6_static_binding"] = vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	paramsTypeMap["dhcp_v6_static_binding"] = vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	pathParams["server_id"] = "serverId"
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
		"dhcp_v6_static_binding",
		"POST",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func Ipv6StaticBindingsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func ipv6StaticBindingsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["binding_id"] = "bindingId"
	pathParams["server_id"] = "serverId"
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
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func Ipv6StaticBindingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
}

func ipv6StaticBindingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["binding_id"] = "bindingId"
	pathParams["server_id"] = "serverId"
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
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func Ipv6StaticBindingsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingListResultBindingType)
}

func ipv6StaticBindingsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	pathParams["server_id"] = "serverId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipv6StaticBindingsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["dhcp_v6_static_binding"] = vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func Ipv6StaticBindingsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
}

func ipv6StaticBindingsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["server_id"] = vapiBindings_.NewStringType()
	fields["binding_id"] = vapiBindings_.NewStringType()
	fields["dhcp_v6_static_binding"] = vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
	fieldNameMap["server_id"] = "ServerId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["dhcp_v6_static_binding"] = "DhcpV6StaticBinding"
	paramsTypeMap["dhcp_v6_static_binding"] = vapiBindings_.NewReferenceType(nsxModel.DhcpV6StaticBindingBindingType)
	paramsTypeMap["server_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serverId"] = vapiBindings_.NewStringType()
	paramsTypeMap["bindingId"] = vapiBindings_.NewStringType()
	pathParams["binding_id"] = "bindingId"
	pathParams["server_id"] = "serverId"
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
		"dhcp_v6_static_binding",
		"PUT",
		"/api/v1/dhcp/servers/{serverId}/ipv6-static-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
