// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Alarms.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``newStatus`` of method Alarms#setstatus.
const Alarms_SETSTATUS_NEW_STATUS_OPEN = "OPEN"

// Possible value for ``newStatus`` of method Alarms#setstatus.
const Alarms_SETSTATUS_NEW_STATUS_ACKNOWLEDGED = "ACKNOWLEDGED"

// Possible value for ``newStatus`` of method Alarms#setstatus.
const Alarms_SETSTATUS_NEW_STATUS_SUPPRESSED = "SUPPRESSED"

// Possible value for ``newStatus`` of method Alarms#setstatus.
const Alarms_SETSTATUS_NEW_STATUS_RESOLVED = "RESOLVED"

// Possible value for ``newStatus`` of method Alarms#setstatus0.
const Alarms_SETSTATUS_0_NEW_STATUS_OPEN = "OPEN"

// Possible value for ``newStatus`` of method Alarms#setstatus0.
const Alarms_SETSTATUS_0_NEW_STATUS_ACKNOWLEDGED = "ACKNOWLEDGED"

// Possible value for ``newStatus`` of method Alarms#setstatus0.
const Alarms_SETSTATUS_0_NEW_STATUS_SUPPRESSED = "SUPPRESSED"

// Possible value for ``newStatus`` of method Alarms#setstatus0.
const Alarms_SETSTATUS_0_NEW_STATUS_RESOLVED = "RESOLVED"

func alarmsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alarm_id"] = bindings.NewStringType()
	fieldNameMap["alarm_id"] = "AlarmId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func alarmsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AlarmBindingType)
}

func alarmsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alarm_id"] = bindings.NewStringType()
	fieldNameMap["alarm_id"] = "AlarmId"
	paramsTypeMap["alarm_id"] = bindings.NewStringType()
	paramsTypeMap["alarmId"] = bindings.NewStringType()
	pathParams["alarm_id"] = "alarmId"
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
		"/api/v1/alarms/{alarmId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func alarmsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["after"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["before"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["event_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["feature_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["severity"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["after"] = "After"
	fieldNameMap["before"] = "Before"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["event_type"] = "EventType"
	fieldNameMap["feature_name"] = "FeatureName"
	fieldNameMap["id"] = "Id"
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_resource_type"] = "NodeResourceType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["severity"] = "Severity"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func alarmsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AlarmsListResultBindingType)
}

func alarmsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["after"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["before"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["event_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["feature_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["severity"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["after"] = "After"
	fieldNameMap["before"] = "Before"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["event_type"] = "EventType"
	fieldNameMap["feature_name"] = "FeatureName"
	fieldNameMap["id"] = "Id"
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_resource_type"] = "NodeResourceType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["severity"] = "Severity"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["after"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["event_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["feature_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["before"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["severity"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["severity"] = "severity"
	queryParams["feature_name"] = "feature_name"
	queryParams["before"] = "before"
	queryParams["intent_path"] = "intent_path"
	queryParams["sort_by"] = "sort_by"
	queryParams["event_type"] = "event_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["after"] = "after"
	queryParams["id"] = "id"
	queryParams["node_resource_type"] = "node_resource_type"
	queryParams["node_id"] = "node_id"
	queryParams["page_size"] = "page_size"
	queryParams["status"] = "status"
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
		"/api/v1/alarms",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func alarmsSetstatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["alarm_id"] = bindings.NewStringType()
	fields["new_status"] = bindings.NewStringType()
	fields["suppress_duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["alarm_id"] = "AlarmId"
	fieldNameMap["new_status"] = "NewStatus"
	fieldNameMap["suppress_duration"] = "SuppressDuration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func alarmsSetstatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AlarmBindingType)
}

func alarmsSetstatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["alarm_id"] = bindings.NewStringType()
	fields["new_status"] = bindings.NewStringType()
	fields["suppress_duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["alarm_id"] = "AlarmId"
	fieldNameMap["new_status"] = "NewStatus"
	fieldNameMap["suppress_duration"] = "SuppressDuration"
	paramsTypeMap["suppress_duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["alarm_id"] = bindings.NewStringType()
	paramsTypeMap["new_status"] = bindings.NewStringType()
	paramsTypeMap["alarmId"] = bindings.NewStringType()
	pathParams["alarm_id"] = "alarmId"
	queryParams["new_status"] = "new_status"
	queryParams["suppress_duration"] = "suppress_duration"
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
		"action=set_status",
		"",
		"POST",
		"/api/v1/alarms/{alarmId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func alarmsSetstatus0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["new_status"] = bindings.NewStringType()
	fields["after"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["before"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["event_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["feature_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["severity"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["suppress_duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["new_status"] = "NewStatus"
	fieldNameMap["after"] = "After"
	fieldNameMap["before"] = "Before"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["event_type"] = "EventType"
	fieldNameMap["feature_name"] = "FeatureName"
	fieldNameMap["id"] = "Id"
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_resource_type"] = "NodeResourceType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["severity"] = "Severity"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["status"] = "Status"
	fieldNameMap["suppress_duration"] = "SuppressDuration"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func alarmsSetstatus0OutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func alarmsSetstatus0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["new_status"] = bindings.NewStringType()
	fields["after"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["before"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["event_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["feature_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["severity"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["suppress_duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["new_status"] = "NewStatus"
	fieldNameMap["after"] = "After"
	fieldNameMap["before"] = "Before"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["event_type"] = "EventType"
	fieldNameMap["feature_name"] = "FeatureName"
	fieldNameMap["id"] = "Id"
	fieldNameMap["intent_path"] = "IntentPath"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["node_resource_type"] = "NodeResourceType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["severity"] = "Severity"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["status"] = "Status"
	fieldNameMap["suppress_duration"] = "SuppressDuration"
	paramsTypeMap["id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["node_resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["after"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["event_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["feature_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["new_status"] = bindings.NewStringType()
	paramsTypeMap["intent_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["suppress_duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["before"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["severity"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["severity"] = "severity"
	queryParams["feature_name"] = "feature_name"
	queryParams["before"] = "before"
	queryParams["new_status"] = "new_status"
	queryParams["intent_path"] = "intent_path"
	queryParams["sort_by"] = "sort_by"
	queryParams["suppress_duration"] = "suppress_duration"
	queryParams["event_type"] = "event_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["after"] = "after"
	queryParams["id"] = "id"
	queryParams["node_resource_type"] = "node_resource_type"
	queryParams["node_id"] = "node_id"
	queryParams["page_size"] = "page_size"
	queryParams["status"] = "status"
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
		"action=set_status",
		"",
		"POST",
		"/api/v1/alarms",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
