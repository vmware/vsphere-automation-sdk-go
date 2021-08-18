// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: HeartbeatConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package antrea

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func heartbeatConfigGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func heartbeatConfigGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
}

func heartbeatConfigGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/policy/api/v1/infra/antrea/heartbeat-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func heartbeatConfigPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["antrea_heartbeat_config"] = bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
	fieldNameMap["antrea_heartbeat_config"] = "AntreaHeartbeatConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func heartbeatConfigPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func heartbeatConfigPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["antrea_heartbeat_config"] = bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
	fieldNameMap["antrea_heartbeat_config"] = "AntreaHeartbeatConfig"
	paramsTypeMap["antrea_heartbeat_config"] = bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
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
		"antrea_heartbeat_config",
		"PATCH",
		"/policy/api/v1/infra/antrea/heartbeat-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func heartbeatConfigUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["antrea_heartbeat_config"] = bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
	fieldNameMap["antrea_heartbeat_config"] = "AntreaHeartbeatConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func heartbeatConfigUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
}

func heartbeatConfigUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["antrea_heartbeat_config"] = bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
	fieldNameMap["antrea_heartbeat_config"] = "AntreaHeartbeatConfig"
	paramsTypeMap["antrea_heartbeat_config"] = bindings.NewReferenceType(model.AntreaHeartbeatConfigBindingType)
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
		"antrea_heartbeat_config",
		"PUT",
		"/policy/api/v1/infra/antrea/heartbeat-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
