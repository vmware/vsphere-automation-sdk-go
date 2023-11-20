// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package logical_ports

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Status#get.
const Status_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Status#get.
const Status_GET_SOURCE_CACHED = "cached"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_VIF = "VIF"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_LOGICALROUTER = "LOGICALROUTER"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_BRIDGEENDPOINT = "BRIDGEENDPOINT"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_DHCP_SERVICE = "DHCP_SERVICE"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_METADATA_PROXY = "METADATA_PROXY"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_L2VPN_SESSION = "L2VPN_SESSION"

// Possible value for ``attachmentType`` of method Status#getall.
const Status_GETALL_ATTACHMENT_TYPE_NONE = "NONE"

// Possible value for ``source`` of method Status#getall.
const Status_GETALL_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Status#getall.
const Status_GETALL_SOURCE_CACHED = "cached"

func statusGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatusGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LogicalPortOperationalStatusBindingType)
}

func statusGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lport_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["lport_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["lportId"] = vapiBindings_.NewStringType()
	pathParams["lport_id"] = "lportId"
	queryParams["source"] = "source"
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
		"/api/v1/logical-ports/{lportId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func statusGetallInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["attachment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["bridge_cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["container_ports_only"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["diagnostic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["logical_switch_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["parent_vif_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["switching_profile_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
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
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatusGetallOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LogicalPortStatusSummaryBindingType)
}

func statusGetallRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["attachment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["bridge_cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["container_ports_only"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["diagnostic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["logical_switch_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["parent_vif_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["switching_profile_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
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
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["switching_profile_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["bridge_cluster_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["attachment_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["attachment_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_zone_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["container_ports_only"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["diagnostic"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["parent_vif_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["logical_switch_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
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
		"/api/v1/logical-ports/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
