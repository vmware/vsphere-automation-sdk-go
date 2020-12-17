/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CloudServiceVmcOnAwsTrafficGroup.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package api

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``action`` of method CloudServiceVmcOnAwsTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVmcOnAwsTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_END_RESTORE = "END_RESTORE"
// Possible value for ``action`` of method CloudServiceVmcOnAwsTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVmcOnAwsTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_START_RESTORE = "START_RESTORE"
// Possible value for ``action`` of method CloudServiceVmcOnAwsTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVmcOnAwsTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_SET_RESOURCE_STATUS_SUCCESS = "SET_RESOURCE_STATUS_SUCCESS"
// Possible value for ``action`` of method CloudServiceVmcOnAwsTrafficGroup#triggerTrafficGroupAction.
const CloudServiceVmcOnAwsTrafficGroup_TRIGGER_TRAFFIC_GROUP_ACTION_ACTION_SET_RESOURCE_STATUS_FAILURE = "SET_RESOURCE_STATUS_FAILURE"




func cloudServiceVmcOnAwsTrafficGroupDeleteTrafficGroupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupDeleteTrafficGroupOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVmcOnAwsTrafficGroupDeleteTrafficGroupRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupDeleteTrafficGroupAssociationMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupDeleteTrafficGroupAssociationMapOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVmcOnAwsTrafficGroupDeleteTrafficGroupAssociationMapRestMetadata() protocol.OperationRestMetadata {
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
		"DELETE",
		"/cloud-service/api/v1/infra/traffic-groups/{trafficGroupId}/association-maps/{mapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupGetTrafficGroupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupGetTrafficGroupOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupBindingType)
}

func cloudServiceVmcOnAwsTrafficGroupGetTrafficGroupRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupGetTrafficGroupAssociationMapInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["map_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["map_id"] = "MapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupGetTrafficGroupAssociationMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
}

func cloudServiceVmcOnAwsTrafficGroupGetTrafficGroupAssociationMapRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupListTrafficGroupAssociationMapsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupListTrafficGroupAssociationMapsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupAssociationMapsListResultBindingType)
}

func cloudServiceVmcOnAwsTrafficGroupListTrafficGroupAssociationMapsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupListTrafficGroupsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["verbose"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["verbose"] = "Verbose"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupListTrafficGroupsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupsListResultBindingType)
}

func cloudServiceVmcOnAwsTrafficGroupListTrafficGroupsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["verbose"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["verbose"] = "Verbose"
	paramsTypeMap["verbose"] = bindings.NewOptionalType(bindings.NewBooleanType())
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupTriggerTrafficGroupActionInputType() bindings.StructType {
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

func cloudServiceVmcOnAwsTrafficGroupTriggerTrafficGroupActionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cloudServiceVmcOnAwsTrafficGroupTriggerTrafficGroupActionRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupUpdateTrafficGroupInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["traffic_group_id"] = bindings.NewStringType()
	fields["traffic_group"] = bindings.NewReferenceType(model.TrafficGroupBindingType)
	fieldNameMap["traffic_group_id"] = "TrafficGroupId"
	fieldNameMap["traffic_group"] = "TrafficGroup"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVmcOnAwsTrafficGroupUpdateTrafficGroupOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupBindingType)
}

func cloudServiceVmcOnAwsTrafficGroupUpdateTrafficGroupRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVmcOnAwsTrafficGroupUpdateTrafficGroupAssociationMapInputType() bindings.StructType {
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

func cloudServiceVmcOnAwsTrafficGroupUpdateTrafficGroupAssociationMapOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TrafficGroupAssociationMapBindingType)
}

func cloudServiceVmcOnAwsTrafficGroupUpdateTrafficGroupAssociationMapRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


