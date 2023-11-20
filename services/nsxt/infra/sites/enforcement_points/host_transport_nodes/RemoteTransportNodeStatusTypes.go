// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RemoteTransportNodeStatus.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package host_transport_nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_0 = "0"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_NO_DIAGNOSTIC = "NO_DIAGNOSTIC"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_1 = "1"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_CONTROL_DETECTION_TIME_EXPIRED = "CONTROL_DETECTION_TIME_EXPIRED"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_2 = "2"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_ECHO_FUNCTION_FAILED = "ECHO_FUNCTION_FAILED"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_3 = "3"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_NEIGHBOR_SIGNALED_SESSION_DOWN = "NEIGHBOR_SIGNALED_SESSION_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_4 = "4"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_FORWARDING_PLANE_RESET = "FORWARDING_PLANE_RESET"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_5 = "5"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_PATH_DOWN = "PATH_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_6 = "6"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_CONCATENATED_PATH_DOWN = "CONCATENATED_PATH_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_7 = "7"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_ADMINISTRATIVELY_DOWN = "ADMINISTRATIVELY_DOWN"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_8 = "8"

// Possible value for ``bfdDiagnosticCode`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_BFD_DIAGNOSTIC_CODE_REVERSE_CONCATENATED_PATH_DOWN = "REVERSE_CONCATENATED_PATH_DOWN"

// Possible value for ``source`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_SOURCE_CACHED = "cached"

// Possible value for ``tunnelStatus`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_TUNNEL_STATUS_UP = "UP"

// Possible value for ``tunnelStatus`` of method RemoteTransportNodeStatus#list.
const RemoteTransportNodeStatus_LIST_TUNNEL_STATUS_DOWN = "DOWN"

func remoteTransportNodeStatusListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["bfd_diagnostic_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["tunnel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["bfd_diagnostic_code"] = "BfdDiagnosticCode"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RemoteTransportNodeStatusListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeStatusListResultBindingType)
}

func remoteTransportNodeStatusListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["bfd_diagnostic_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["tunnel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["bfd_diagnostic_code"] = "BfdDiagnosticCode"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["source"] = "Source"
	fieldNameMap["tunnel_status"] = "TunnelStatus"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["tunnel_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bfd_diagnostic_code"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	paramsTypeMap["nodeId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["node_id"] = "nodeId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["tunnel_status"] = "tunnel_status"
	queryParams["bfd_diagnostic_code"] = "bfd_diagnostic_code"
	queryParams["sort_by"] = "sort_by"
	queryParams["source"] = "source"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/host-transport-nodes/{nodeId}/remote-transport-node-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
