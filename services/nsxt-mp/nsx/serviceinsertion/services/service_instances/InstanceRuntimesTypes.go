// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: InstanceRuntimes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service_instances

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``action`` of method InstanceRuntimes#create.
const InstanceRuntimes_CREATE_ACTION_ENABLE_MAINTENANCE_MODE = "enable_maintenance_mode"

// Possible value for ``action`` of method InstanceRuntimes#create.
const InstanceRuntimes_CREATE_ACTION_DISABLE_MAINTENANCE_MODE = "disable_maintenance_mode"

// Possible value for ``action`` of method InstanceRuntimes#create.
const InstanceRuntimes_CREATE_ACTION_IS_HEALTHY = "is_healthy"

// Possible value for ``action`` of method InstanceRuntimes#create.
const InstanceRuntimes_CREATE_ACTION_IS_STOPPED = "is_stopped"

// Possible value for ``action`` of method InstanceRuntimes#create.
const InstanceRuntimes_CREATE_ACTION_IS_NOT_RESPONDING = "is_not_responding"

func instanceRuntimesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["unhealthy_reason"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["unhealthy_reason"] = "UnhealthyReason"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instanceRuntimesCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func instanceRuntimesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["instance_runtime_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["unhealthy_reason"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["unhealthy_reason"] = "UnhealthyReason"
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["unhealthy_reason"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["instance_runtime_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["instanceRuntimeId"] = bindings.NewStringType()
	pathParams["instance_runtime_id"] = "instanceRuntimeId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
	queryParams["action"] = "action"
	queryParams["unhealthy_reason"] = "unhealthy_reason"
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
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes/{instanceRuntimeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceRuntimesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instanceRuntimesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func instanceRuntimesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"action=delete",
		"",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceRuntimesDeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instanceRuntimesDeployOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func instanceRuntimesDeployRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"action=deploy",
		"",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceRuntimesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instanceRuntimesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.InstanceRuntimeListResultBindingType)
}

func instanceRuntimesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceRuntimesUpgradeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instanceRuntimesUpgradeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func instanceRuntimesUpgradeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"action=upgrade",
		"",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
