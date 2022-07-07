// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Notifications.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package notification_watchers

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func notificationsAddurifiltersInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = bindings.NewStringType()
	fields["notification"] = bindings.NewReferenceType(model.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func notificationsAddurifiltersOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NotificationsListBindingType)
}

func notificationsAddurifiltersRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = bindings.NewStringType()
	fields["notification"] = bindings.NewReferenceType(model.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	paramsTypeMap["watcher_id"] = bindings.NewStringType()
	paramsTypeMap["notification"] = bindings.NewReferenceType(model.NotificationBindingType)
	paramsTypeMap["watcherId"] = bindings.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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
		"action=add_uri_filters",
		"notification",
		"POST",
		"/api/v1/notification-watchers/{watcherId}/notifications",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func notificationsDeleteurifiltersInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = bindings.NewStringType()
	fields["notification"] = bindings.NewReferenceType(model.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func notificationsDeleteurifiltersOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NotificationsListBindingType)
}

func notificationsDeleteurifiltersRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = bindings.NewStringType()
	fields["notification"] = bindings.NewReferenceType(model.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	paramsTypeMap["watcher_id"] = bindings.NewStringType()
	paramsTypeMap["notification"] = bindings.NewReferenceType(model.NotificationBindingType)
	paramsTypeMap["watcherId"] = bindings.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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
		"action=delete_uri_filters",
		"notification",
		"POST",
		"/api/v1/notification-watchers/{watcherId}/notifications",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func notificationsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = bindings.NewStringType()
	fieldNameMap["watcher_id"] = "WatcherId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func notificationsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NotificationsListBindingType)
}

func notificationsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = bindings.NewStringType()
	fieldNameMap["watcher_id"] = "WatcherId"
	paramsTypeMap["watcher_id"] = bindings.NewStringType()
	paramsTypeMap["watcherId"] = bindings.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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
		"/api/v1/notification-watchers/{watcherId}/notifications",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func notificationsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = bindings.NewStringType()
	fields["notifications_list"] = bindings.NewReferenceType(model.NotificationsListBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notifications_list"] = "NotificationsList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func notificationsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NotificationsListBindingType)
}

func notificationsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = bindings.NewStringType()
	fields["notifications_list"] = bindings.NewReferenceType(model.NotificationsListBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notifications_list"] = "NotificationsList"
	paramsTypeMap["watcher_id"] = bindings.NewStringType()
	paramsTypeMap["notifications_list"] = bindings.NewReferenceType(model.NotificationsListBindingType)
	paramsTypeMap["watcherId"] = bindings.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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
		"notifications_list",
		"PUT",
		"/api/v1/notification-watchers/{watcherId}/notifications",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
