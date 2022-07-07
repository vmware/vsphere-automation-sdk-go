// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package logical_ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lport_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["lport_id"] = "LportId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortOperationalStatusBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
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

func statusGetallInputType() bindings.StructType {
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

func statusGetallOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LogicalPortStatusSummaryBindingType)
}

func statusGetallRestMetadata() protocol.OperationRestMetadata {
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
