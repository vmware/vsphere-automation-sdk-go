// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Notifications.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package notification_watchers

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func notificationsAddurifiltersInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fields["notification"] = vapiBindings_.NewReferenceType(nsxModel.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NotificationsAddurifiltersOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
}

func notificationsAddurifiltersRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fields["notification"] = vapiBindings_.NewReferenceType(nsxModel.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	paramsTypeMap["watcher_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["notification"] = vapiBindings_.NewReferenceType(nsxModel.NotificationBindingType)
	paramsTypeMap["watcherId"] = vapiBindings_.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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

func notificationsDeleteurifiltersInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fields["notification"] = vapiBindings_.NewReferenceType(nsxModel.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NotificationsDeleteurifiltersOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
}

func notificationsDeleteurifiltersRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fields["notification"] = vapiBindings_.NewReferenceType(nsxModel.NotificationBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notification"] = "Notification"
	paramsTypeMap["watcher_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["notification"] = vapiBindings_.NewReferenceType(nsxModel.NotificationBindingType)
	paramsTypeMap["watcherId"] = vapiBindings_.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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

func notificationsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fieldNameMap["watcher_id"] = "WatcherId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NotificationsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
}

func notificationsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fieldNameMap["watcher_id"] = "WatcherId"
	paramsTypeMap["watcher_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["watcherId"] = vapiBindings_.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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
		"/api/v1/notification-watchers/{watcherId}/notifications",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func notificationsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fields["notifications_list"] = vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notifications_list"] = "NotificationsList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func NotificationsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
}

func notificationsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["watcher_id"] = vapiBindings_.NewStringType()
	fields["notifications_list"] = vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
	fieldNameMap["watcher_id"] = "WatcherId"
	fieldNameMap["notifications_list"] = "NotificationsList"
	paramsTypeMap["watcher_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["notifications_list"] = vapiBindings_.NewReferenceType(nsxModel.NotificationsListBindingType)
	paramsTypeMap["watcherId"] = vapiBindings_.NewStringType()
	pathParams["watcher_id"] = "watcherId"
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
