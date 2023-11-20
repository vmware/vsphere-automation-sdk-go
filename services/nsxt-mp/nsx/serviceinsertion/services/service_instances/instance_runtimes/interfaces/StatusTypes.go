// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package interfaces

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

func statusGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_runtime_id"] = vapiBindings_.NewStringType()
	fields["interface_index"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["interface_index"] = "InterfaceIndex"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatusGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.RuntimeInterfaceOperationalStatusBindingType)
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
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_runtime_id"] = vapiBindings_.NewStringType()
	fields["interface_index"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_runtime_id"] = "InstanceRuntimeId"
	fieldNameMap["interface_index"] = "InterfaceIndex"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["interface_index"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["instance_runtime_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["instanceRuntimeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["interfaceIndex"] = vapiBindings_.NewStringType()
	pathParams["instance_runtime_id"] = "instanceRuntimeId"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["interface_index"] = "interfaceIndex"
	pathParams["service_id"] = "serviceId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-runtimes/{instanceRuntimeId}/interfaces/{interfaceIndex}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
