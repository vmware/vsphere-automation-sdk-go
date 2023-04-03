// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Operation.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package service

import (
	vapiMetadataPrivilege_ "github.com/vmware/vsphere-automation-sdk-go/lib/vapi/metadata/privilege"
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Resource type for operation.
const Operation_RESOURCE_TYPE = "com.vmware.vapi.operation"

func operationListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func OperationListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf([]string{}))
}

func operationListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	paramsTypeMap["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["serviceId"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
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
		"/vapi/metadata/privilege/service/{serviceId}/operation",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}

func operationGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func OperationGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vapiMetadataPrivilege_.OperationInfoBindingType)
}

func operationGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	fields["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["operation_id"] = "OperationId"
	paramsTypeMap["service_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["operation_id"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	paramsTypeMap["serviceId"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.service"}, "")
	paramsTypeMap["operationId"] = vapiBindings_.NewIdType([]string{"com.vmware.vapi.operation"}, "")
	pathParams["operation_id"] = "operationId"
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
		"/vapi/metadata/privilege/service/{serviceId}/operation/{operationId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.not_found": 404})
}
