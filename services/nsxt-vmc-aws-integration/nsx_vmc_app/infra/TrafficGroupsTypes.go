// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TrafficGroups.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package infra

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
	"reflect"
)

// Possible value for ``action`` of method TrafficGroups#create.
const TrafficGroups_CREATE_ACTION_END_RESTORE = "END_RESTORE"

// Possible value for ``action`` of method TrafficGroups#create.
const TrafficGroups_CREATE_ACTION_START_RESTORE = "START_RESTORE"

// Possible value for ``action`` of method TrafficGroups#create.
const TrafficGroups_CREATE_ACTION_SET_RESOURCE_STATUS_SUCCESS = "SET_RESOURCE_STATUS_SUCCESS"

// Possible value for ``action`` of method TrafficGroups#create.
const TrafficGroups_CREATE_ACTION_SET_RESOURCE_STATUS_FAILURE = "SET_RESOURCE_STATUS_FAILURE"

// Possible value for ``state`` of method TrafficGroups#list.
const TrafficGroups_LIST_STATE_FAILURE = "FAILURE"

// Possible value for ``state`` of method TrafficGroups#list.
const TrafficGroups_LIST_STATE_UNAVAILABLE = "UNAVAILABLE"

// Possible value for ``state`` of method TrafficGroups#list.
const TrafficGroups_LIST_STATE_SUCCESS = "SUCCESS"

// Possible value for ``state`` of method TrafficGroups#list.
const TrafficGroups_LIST_STATE_PENDING = "PENDING"

// Possible value for ``state`` of method TrafficGroups#list.
const TrafficGroups_LIST_STATE_IN_PROGRESS = "IN_PROGRESS"

func trafficGroupsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["action_message"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ActionMessageBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["action_message"] = "ActionMessage"
	fieldNameMap["action"] = "Action"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TrafficGroupsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func trafficGroupsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["action_message"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ActionMessageBindingType)
	fields["action"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["action_message"] = "ActionMessage"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["action_message"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ActionMessageBindingType)
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["action"] = vapiBindings_.NewStringType()
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
	queryParams["action"] = "action"
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

func trafficGroupsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["force"] = "Force"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TrafficGroupsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func trafficGroupsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["force"] = "Force"
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["force"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
	queryParams["force"] = "force"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func trafficGroupsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TrafficGroupsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupBindingType)
}

func trafficGroupsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func trafficGroupsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["verbose"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["verbose"] = "Verbose"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TrafficGroupsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupsListResultBindingType)
}

func trafficGroupsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["verbose"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["verbose"] = "Verbose"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["verbose"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["state"] = "state"
	queryParams["page_size"] = "page_size"
	queryParams["verbose"] = "verbose"
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
		"/cloud-service/api/v1/infra/traffic-groups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func trafficGroupsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["traffic_group"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["traffic_group"] = "TrafficGroup"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TrafficGroupsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupBindingType)
}

func trafficGroupsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["traffic_group_id"] = vapiBindings_.NewStringType()
	fields["traffic_group"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["traffic_group"] = "TrafficGroup"
	paramsTypeMap["traffic_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["traffic_group"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.TrafficGroupBindingType)
	paramsTypeMap["trafficGroupId"] = vapiBindings_.NewStringType()
	pathParams["traffic_group_id"] = "trafficGroupId"
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
