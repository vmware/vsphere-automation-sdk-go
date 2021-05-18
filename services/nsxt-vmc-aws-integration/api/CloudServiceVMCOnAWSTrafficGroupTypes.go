// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CloudServiceVMCOnAWSTrafficGroup.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"reflect"
)

// Possible value for ``state`` of method CloudServiceVMCOnAWSTrafficGroup#listTrafficGroups.
const CloudServiceVMCOnAWSTrafficGroup_LIST_TRAFFIC_GROUPS_STATE_FAILURE = "FAILURE"

// Possible value for ``state`` of method CloudServiceVMCOnAWSTrafficGroup#listTrafficGroups.
const CloudServiceVMCOnAWSTrafficGroup_LIST_TRAFFIC_GROUPS_STATE_UNAVAILABLE = "UNAVAILABLE"

// Possible value for ``state`` of method CloudServiceVMCOnAWSTrafficGroup#listTrafficGroups.
const CloudServiceVMCOnAWSTrafficGroup_LIST_TRAFFIC_GROUPS_STATE_SUCCESS = "SUCCESS"

// Possible value for ``state`` of method CloudServiceVMCOnAWSTrafficGroup#listTrafficGroups.
const CloudServiceVMCOnAWSTrafficGroup_LIST_TRAFFIC_GROUPS_STATE_PENDING = "PENDING"

// Possible value for ``state`` of method CloudServiceVMCOnAWSTrafficGroup#listTrafficGroups.
const CloudServiceVMCOnAWSTrafficGroup_LIST_TRAFFIC_GROUPS_STATE_IN_PROGRESS = "IN_PROGRESS"

// Possible value for ``action`` of method CloudServiceVMCOnAWSTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVMCOnAWSTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_END_RESTORE = "END_RESTORE"

// Possible value for ``action`` of method CloudServiceVMCOnAWSTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVMCOnAWSTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_START_RESTORE = "START_RESTORE"

// Possible value for ``action`` of method CloudServiceVMCOnAWSTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVMCOnAWSTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_SET_RESOURCE_STATUS_SUCCESS = "SET_RESOURCE_STATUS_SUCCESS"

// Possible value for ``action`` of method CloudServiceVMCOnAWSTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVMCOnAWSTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_SET_RESOURCE_STATUS_FAILURE = "SET_RESOURCE_STATUS_FAILURE"

func cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupAssociationMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupAssociationMapOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSTrafficGroupDeleteTrafficGroupAssociationMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["map_id"] = bindings.NewStringType()
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	paramsTypeMap["mapId"] = bindings.NewStringType()
	pathParams["map_id"] = "mapId"
	pathParams["traffic_group_id"] = "trafficGroupId"
	queryParams["force"] = "force"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupBindingType)
}

func cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupAssociationMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupAssociationMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
}

func cloudServiceVMCOnAWSTrafficGroupGetTrafficGroupAssociationMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	paramsTypeMap["map_id"] = bindings.NewStringType()
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	paramsTypeMap["mapId"] = bindings.NewStringType()
	pathParams["map_id"] = "mapId"
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupListTrafficGroupAssociationMapsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupListTrafficGroupAssociationMapsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupAssociationMapsListResultBindingType)
}

func cloudServiceVMCOnAWSTrafficGroupListTrafficGroupAssociationMapsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupListTrafficGroupsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["verbose"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["verbose"] = "Verbose"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupListTrafficGroupsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupsListResultBindingType)
}

func cloudServiceVMCOnAWSTrafficGroupListTrafficGroupsRestMetadata() protocol.OperationRestMetadata {
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
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["verbose"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["verbose"] = "Verbose"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["verbose"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["state"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["state"] = "state"
	queryParams["page_size"] = "page_size"
	queryParams["verbose"] = "verbose"
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
		"/cloud-service/api/v1/infra/traffic-groups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupTriggerTrafficGroupActionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["action_message"] = bindings.NewReferenceType(model.ActionMessageBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["action_message"] = "ActionMessage"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupTriggerTrafficGroupActionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVMCOnAWSTrafficGroupTriggerTrafficGroupActionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["action_message"] = bindings.NewReferenceType(model.ActionMessageBindingType)
	fields["action"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["action_message"] = "ActionMessage"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["action_message"] = bindings.NewReferenceType(model.ActionMessageBindingType)
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
	queryParams["action"] = "action"
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
		"action_message",
		"POST",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["traffic_group"] = bindings.NewReferenceType(model.TrafficGroupBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["traffic_group"] = "TrafficGroup"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupBindingType)
}

func cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["traffic_group"] = bindings.NewReferenceType(model.TrafficGroupBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["traffic_group"] = "TrafficGroup"
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["traffic_group"] = bindings.NewReferenceType(model.TrafficGroupBindingType)
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"traffic_group",
		"PUT",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupAssociationMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fields["traffic_group_association_map"] = bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["traffic_group_association_map"] = "TrafficGroupAssociationMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupAssociationMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
}

func cloudServiceVMCOnAWSTrafficGroupUpdateTrafficGroupAssociationMapRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fields["traffic_group_association_map"] = bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	fieldNameMap["traffic_group_association_map"] = "TrafficGroupAssociationMap"
	paramsTypeMap["traffic_group_association_map"] = bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
	paramsTypeMap["map_id"] = bindings.NewStringType()
	paramsTypeMap["traffic_group_id"] = bindings.NewStringType()
	paramsTypeMap["trafficGroupId"] = bindings.NewStringType()
	paramsTypeMap["mapId"] = bindings.NewStringType()
	pathParams["map_id"] = "mapId"
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"traffic_group_association_map",
		"PUT",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
