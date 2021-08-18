// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TransportNodes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package intelligence

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``transportNodeType`` of method TransportNodes#list.
const TransportNodes_LIST_TRANSPORT_NODE_TYPE_NODE = "HOST_NODE"

func transportNodesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	paramsTypeMap["tn_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["tnId"] = bindings.NewStringType()
	pathParams["tn_id"] = "tnId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/intelligence/transport-nodes/{tnId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
}

func transportNodesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	paramsTypeMap["tn_id"] = bindings.NewStringType()
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["tnId"] = bindings.NewStringType()
	pathParams["tn_id"] = "tnId"
	pathParams["site_id"] = "siteId"
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
		"/policy/api/v1/infra/sites/{siteId}/intelligence/transport-nodes/{tnId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["transport_node_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["standalone"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["transport_node_type"] = "TransportNodeType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enabled"] = "Enabled"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["standalone"] = "Standalone"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeInfoListBindingType)
}

func transportNodesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["transport_node_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["standalone"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["transport_node_type"] = "TransportNodeType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enabled"] = "Enabled"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["standalone"] = "Standalone"
	paramsTypeMap["standalone"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_node_type"] = bindings.NewStringType()
	paramsTypeMap["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["standalone"] = "standalone"
	queryParams["sort_by"] = "sort_by"
	queryParams["transport_node_type"] = "transport_node_type"
	queryParams["enabled"] = "enabled"
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
		"/policy/api/v1/infra/sites/{siteId}/intelligence/transport-nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_info_list"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoListBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["intelligence_transport_node_info_list"] = "IntelligenceTransportNodeInfoList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesPatchOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeInfoListBindingType)
}

func transportNodesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_info_list"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoListBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["intelligence_transport_node_info_list"] = "IntelligenceTransportNodeInfoList"
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["intelligence_transport_node_info_list"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoListBindingType)
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
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
		"intelligence_transport_node_info_list",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/intelligence/transport-nodes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesPatch0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	fieldNameMap["intelligence_transport_node_info"] = "IntelligenceTransportNodeInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesPatch0OutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
}

func transportNodesPatch0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	fieldNameMap["intelligence_transport_node_info"] = "IntelligenceTransportNodeInfo"
	paramsTypeMap["tn_id"] = bindings.NewStringType()
	paramsTypeMap["intelligence_transport_node_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["tnId"] = bindings.NewStringType()
	pathParams["tn_id"] = "tnId"
	pathParams["site_id"] = "siteId"
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
		"intelligence_transport_node_info",
		"PATCH",
		"/policy/api/v1/infra/sites/{siteId}/intelligence/transport-nodes/{tnId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	fieldNameMap["intelligence_transport_node_info"] = "IntelligenceTransportNodeInfo"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
}

func transportNodesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["tn_id"] = bindings.NewStringType()
	fields["intelligence_transport_node_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["tn_id"] = "TnId"
	fieldNameMap["intelligence_transport_node_info"] = "IntelligenceTransportNodeInfo"
	paramsTypeMap["tn_id"] = bindings.NewStringType()
	paramsTypeMap["intelligence_transport_node_info"] = bindings.NewReferenceType(model.IntelligenceTransportNodeInfoBindingType)
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	paramsTypeMap["tnId"] = bindings.NewStringType()
	pathParams["tn_id"] = "tnId"
	pathParams["site_id"] = "siteId"
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
		"intelligence_transport_node_info",
		"PUT",
		"/policy/api/v1/infra/sites/{siteId}/intelligence/transport-nodes/{tnId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
