// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: InstanceRuntimes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service_instances

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func instanceRuntimesCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_runtime_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["unhealthy_reason"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["unhealthy_reason"] = "UnhealthyReason"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceRuntimesCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func instanceRuntimesCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_runtime_id"] = vapiBindings_.NewStringType()
	fields["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["unhealthy_reason"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["unhealthy_reason"] = "UnhealthyReason"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["action"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["instance_runtime_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["unhealthy_reason"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["instanceRuntimeId"] = vapiBindings_.NewStringType()
	pathParams["instance_runtime_id"] = "instanceRuntimeId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
	queryParams["action"] = "action"
	queryParams["unhealthy_reason"] = "unhealthy_reason"
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
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes/{instanceRuntimeId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceRuntimesDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceRuntimesDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func instanceRuntimesDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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

func instanceRuntimesDeployInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceRuntimesDeployOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func instanceRuntimesDeployRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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

func instanceRuntimesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceRuntimesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.InstanceRuntimeListResultBindingType)
}

func instanceRuntimesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceRuntimesUpgradeInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceRuntimesUpgradeOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func instanceRuntimesUpgradeRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["service_id"] = "serviceId"
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
