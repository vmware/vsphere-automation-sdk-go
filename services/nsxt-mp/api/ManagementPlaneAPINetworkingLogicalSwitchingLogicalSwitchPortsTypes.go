// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortMacTable.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_MAC_TABLE_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortMacTable.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_MAC_TABLE_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortMacTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_MAC_TABLE_IN_CSV_FORMAT_CSV_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortMacTableInCsvFormatCsv.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_MAC_TABLE_IN_CSV_FORMAT_CSV_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortOperationalStatus.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_OPERATIONAL_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortOperationalStatus.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_OPERATIONAL_STATUS_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatistics.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatistics.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATISTICS_SOURCE_CACHED = "cached"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_VIF = "VIF"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_LOGICALROUTER = "LOGICALROUTER"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_BRIDGEENDPOINT = "BRIDGEENDPOINT"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_DHCP_SERVICE = "DHCP_SERVICE"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_METADATA_PROXY = "METADATA_PROXY"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_L2VPN_SESSION = "L2VPN_SESSION"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_ATTACHMENT_TYPE_NONE = "NONE"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#getLogicalPortStatusSummary.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_GET_LOGICAL_PORT_STATUS_SUMMARY_SOURCE_CACHED = "cached"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_VIF = "VIF"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_LOGICALROUTER = "LOGICALROUTER"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_BRIDGEENDPOINT = "BRIDGEENDPOINT"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_DHCP_SERVICE = "DHCP_SERVICE"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_METADATA_PROXY = "METADATA_PROXY"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_L2VPN_SESSION = "L2VPN_SESSION"

// Possible value for ``attachmentType`` of method ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts#listLogicalPorts.
const ManagementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPorts_LIST_LOGICAL_PORTS_ATTACHMENT_TYPE_NONE = "NONE"

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateLogicalPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_port"] = bindings.NewReferenceType(model.LogicalPortBindingType)
	fieldNameMap["logical_port"] = "LogicalPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateLogicalPortOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateLogicalPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_port"] = bindings.NewReferenceType(model.LogicalPortBindingType)
	fieldNameMap["logical_port"] = "LogicalPort"
	paramsTypeMap["logical_port"] = bindings.NewReferenceType(model.LogicalPortBindingType)
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
		"logical_port",
		"POST",
		"/api/v1/logical-ports",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateSwitchingProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["base_switching_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
	fieldNameMap["base_switching_profile"] = "BaseSwitchingProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateSwitchingProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsCreateSwitchingProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["base_switching_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
	fieldNameMap["base_switching_profile"] = "BaseSwitchingProfile"
	paramsTypeMap["base_switching_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
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
		"base_switching_profile",
		"POST",
		"/api/v1/switching-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteLogicalPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["detach"] = "Detach"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteLogicalPortOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteLogicalPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["detach"] = "Detach"
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["detach"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["detach"] = "detach"
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
		"/api/v1/logical-ports/{lportId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteSwitchingProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["switching_profile_id"] = bindings.NewStringType()
	fields["unbind"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["unbind"] = "Unbind"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteSwitchingProfileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsDeleteSwitchingProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["switching_profile_id"] = bindings.NewStringType()
	fields["unbind"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["unbind"] = "Unbind"
	paramsTypeMap["switching_profile_id"] = bindings.NewStringType()
	paramsTypeMap["unbind"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["switchingProfileId"] = bindings.NewStringType()
	pathParams["switching_profile_id"] = "switchingProfileId"
	queryParams["unbind"] = "unbind"
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
		"/api/v1/switching-profiles/{switchingProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fieldNameMap["lport_id"] = "LportId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fieldNameMap["lport_id"] = "LportId"
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
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
		"/api/v1/logical-ports/{lportId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortMacAddressListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"/api/v1/logical-ports/{lportId}/mac-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInCsvFormatCsvInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInCsvFormatCsvOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortMacAddressCsvListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortMacTableInCsvFormatCsvRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"format=csv",
		"",
		"GET",
		"/api/v1/logical-ports/{lportId}/mac-table",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOperationalStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOperationalStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortOperationalStatusBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortOperationalStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["source"] = "source"
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
		"/api/v1/logical-ports/{lportId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fieldNameMap["lport_id"] = "LportId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortStateBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fieldNameMap["lport_id"] = "LportId"
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
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
		"/api/v1/logical-ports/{lportId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortStatisticsBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["source"] = "source"
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
		"/api/v1/logical-ports/{lportId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatusSummaryInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["attachment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["attachment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["container_ports_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["parent_vif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fieldNameMap["attachment_type"] = "AttachmentType"
	fieldNameMap["bridge_cluster_id"] = "BridgeClusterId"
	fieldNameMap["container_ports_only"] = "ContainerPortsOnly"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["parent_vif_id"] = "ParentVifId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatusSummaryOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortStatusSummaryBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetLogicalPortStatusSummaryRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["attachment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["attachment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["container_ports_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["parent_vif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fieldNameMap["attachment_type"] = "AttachmentType"
	fieldNameMap["bridge_cluster_id"] = "BridgeClusterId"
	fieldNameMap["container_ports_only"] = "ContainerPortsOnly"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["parent_vif_id"] = "ParentVifId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	paramsTypeMap["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["parent_vif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["container_ports_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["attachment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["attachment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["switching_profile_id"] = "switching_profile_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
	queryParams["bridge_cluster_id"] = "bridge_cluster_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["attachment_type"] = "attachment_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["attachment_id"] = "attachment_id"
	queryParams["transport_zone_id"] = "transport_zone_id"
	queryParams["container_ports_only"] = "container_ports_only"
	queryParams["diagnostic"] = "diagnostic"
	queryParams["parent_vif_id"] = "parent_vif_id"
	queryParams["logical_switch_id"] = "logical_switch_id"
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
		"/api/v1/logical-ports/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["switching_profile_id"] = bindings.NewStringType()
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["switching_profile_id"] = bindings.NewStringType()
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	paramsTypeMap["switching_profile_id"] = bindings.NewStringType()
	paramsTypeMap["switchingProfileId"] = bindings.NewStringType()
	pathParams["switching_profile_id"] = "switchingProfileId"
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
		"/api/v1/switching-profiles/{switchingProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["switching_profile_id"] = bindings.NewStringType()
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SwitchingProfileStatusBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsGetSwitchingProfileStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["switching_profile_id"] = bindings.NewStringType()
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	paramsTypeMap["switching_profile_id"] = bindings.NewStringType()
	paramsTypeMap["switchingProfileId"] = bindings.NewStringType()
	pathParams["switching_profile_id"] = "switchingProfileId"
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
		"/api/v1/switching-profiles/{switchingProfileId}/summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListLogicalPortsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["attachment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["attachment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["container_ports_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["parent_vif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fieldNameMap["attachment_type"] = "AttachmentType"
	fieldNameMap["bridge_cluster_id"] = "BridgeClusterId"
	fieldNameMap["container_ports_only"] = "ContainerPortsOnly"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["parent_vif_id"] = "ParentVifId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListLogicalPortsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListLogicalPortsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["attachment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["attachment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["container_ports_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["parent_vif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["attachment_id"] = "AttachmentId"
	fieldNameMap["attachment_type"] = "AttachmentType"
	fieldNameMap["bridge_cluster_id"] = "BridgeClusterId"
	fieldNameMap["container_ports_only"] = "ContainerPortsOnly"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["diagnostic"] = "Diagnostic"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["parent_vif_id"] = "ParentVifId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	fieldNameMap["transport_zone_id"] = "TransportZoneId"
	paramsTypeMap["bridge_cluster_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["switching_profile_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["parent_vif_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["container_ports_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["diagnostic"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["attachment_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["attachment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["transport_zone_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["switching_profile_id"] = "switching_profile_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["transport_node_id"] = "transport_node_id"
	queryParams["bridge_cluster_id"] = "bridge_cluster_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["attachment_type"] = "attachment_type"
	queryParams["included_fields"] = "included_fields"
	queryParams["attachment_id"] = "attachment_id"
	queryParams["transport_zone_id"] = "transport_zone_id"
	queryParams["container_ports_only"] = "container_ports_only"
	queryParams["diagnostic"] = "diagnostic"
	queryParams["parent_vif_id"] = "parent_vif_id"
	queryParams["logical_switch_id"] = "logical_switch_id"
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
		"/api/v1/logical-ports",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListSwitchingProfilesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_system_owned"] = "IncludeSystemOwned"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switching_profile_type"] = "SwitchingProfileType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListSwitchingProfilesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SwitchingProfilesListResultBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsListSwitchingProfilesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["switching_profile_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_system_owned"] = "IncludeSystemOwned"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["switching_profile_type"] = "SwitchingProfileType"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["include_system_owned"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["switching_profile_type"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["cursor"] = "cursor"
	queryParams["include_system_owned"] = "include_system_owned"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["switching_profile_type"] = "switching_profile_type"
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
		"/api/v1/switching-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateLogicalPortInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["logical_port"] = bindings.NewReferenceType(model.LogicalPortBindingType)
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["logical_port"] = "LogicalPort"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateLogicalPortOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortBindingType)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateLogicalPortRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = bindings.NewStringType()
	fields["logical_port"] = bindings.NewReferenceType(model.LogicalPortBindingType)
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["logical_port"] = "LogicalPort"
	paramsTypeMap["lport_id"] = bindings.NewStringType()
	paramsTypeMap["logical_port"] = bindings.NewReferenceType(model.LogicalPortBindingType)
	paramsTypeMap["lportId"] = bindings.NewStringType()
	pathParams["lport_id"] = "lportId"
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
		"logical_port",
		"PUT",
		"/api/v1/logical-ports/{lportId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateSwitchingProfileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["switching_profile_id"] = bindings.NewStringType()
	fields["base_switching_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["base_switching_profile"] = "BaseSwitchingProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateSwitchingProfileOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
}

func managementPlaneAPINetworkingLogicalSwitchingLogicalSwitchPortsUpdateSwitchingProfileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["switching_profile_id"] = bindings.NewStringType()
	fields["base_switching_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
	fieldNameMap["switching_profile_id"] = "SwitchingProfileId"
	fieldNameMap["base_switching_profile"] = "BaseSwitchingProfile"
	paramsTypeMap["switching_profile_id"] = bindings.NewStringType()
	paramsTypeMap["base_switching_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.BaseSwitchingProfileBindingType)}, bindings.REST)
	paramsTypeMap["switchingProfileId"] = bindings.NewStringType()
	pathParams["switching_profile_id"] = "switchingProfileId"
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
		"base_switching_profile",
		"PUT",
		"/api/v1/switching-profiles/{switchingProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
