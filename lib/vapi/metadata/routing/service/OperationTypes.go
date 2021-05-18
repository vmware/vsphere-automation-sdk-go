// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Operation.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/metadata/routing"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Resource type for vAPI operation.
const Operation_RESOURCE_TYPE = "com.vmware.vapi.operation"

func operationListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf([]string{}))
}

func operationListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["serviceId"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
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
		"/vapi/metadata/routing/service/{serviceId}/operation",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func operationGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(routing.OperationInfoBindingType)
}

func operationGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	paramsTypeMap["service_id"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["operation_id"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	paramsTypeMap["serviceId"] = bindings.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["operationId"] = bindings.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	pathParams["operation_id"] = "operationId"
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
		"/vapi/metadata/routing/service/{serviceId}/operation/{operationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}
