// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: InstanceEndpoints.
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

func instanceEndpointsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_endpoint"] = vapiBindings_.NewReferenceType(nsxModel.InstanceEndpointBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint"] = "InstanceEndpoint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceEndpointsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.InstanceEndpointBindingType)
}

func instanceEndpointsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["instance_endpoint"] = vapiBindings_.NewReferenceType(nsxModel.InstanceEndpointBindingType)
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint"] = "InstanceEndpoint"
	paramsTypeMap["instance_endpoint"] = vapiBindings_.NewReferenceType(nsxModel.InstanceEndpointBindingType)
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
		"instance_endpoint",
		"POST",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceEndpointsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_endpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceEndpointsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func instanceEndpointsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["instance_endpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["instance_endpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["instanceEndpointId"] = vapiBindings_.NewStringType()
	pathParams["instance_endpoint_id"] = "instanceEndpointId"
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
		"DELETE",
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints/{instanceEndpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceEndpointsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fields["instance_endpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceEndpointsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.InstanceEndpointBindingType)
}

func instanceEndpointsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["instance_endpoint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["instance_endpoint_id"] = "InstanceEndpointId"
	paramsTypeMap["service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_instance_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["instance_endpoint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceInstanceId"] = vapiBindings_.NewStringType()
	paramsTypeMap["instanceEndpointId"] = vapiBindings_.NewStringType()
	pathParams["instance_endpoint_id"] = "instanceEndpointId"
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints/{instanceEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func instanceEndpointsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_id"] = vapiBindings_.NewStringType()
	fields["service_instance_id"] = vapiBindings_.NewStringType()
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InstanceEndpointsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.InstanceEndpointListResultBindingType)
}

func instanceEndpointsListRestMetadata() vapiProtocol_.OperationRestMetadata {
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
		"/api/v1/serviceinsertion/services/nsxt-mp/{serviceId}/service-instances/{serviceInstanceId}/instance-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
