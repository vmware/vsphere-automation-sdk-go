// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CloudServiceVMCOnAWSLinkedVPC.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/model"
	"reflect"
)

func cloudServiceVMCOnAWSLinkedVPCGetLinkedVpcInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = bindings.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSLinkedVPCGetLinkedVpcOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LinkedVpcInfoBindingType)
}

func cloudServiceVMCOnAWSLinkedVPCGetLinkedVpcRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = bindings.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	paramsTypeMap["linked_vpc_id"] = bindings.NewStringType()
	paramsTypeMap["linkedVpcId"] = bindings.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
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
		"/cloud-service/api/v1/infra/linked-vpcs/{linkedVpcId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSLinkedVPCListConnectedServicesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = bindings.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSLinkedVPCListConnectedServicesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConnectedServiceListResultBindingType)
}

func cloudServiceVMCOnAWSLinkedVPCListConnectedServicesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = bindings.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	paramsTypeMap["linked_vpc_id"] = bindings.NewStringType()
	paramsTypeMap["linkedVpcId"] = bindings.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
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
		"/cloud-service/api/v1/infra/linked-vpcs/{linkedVpcId}/connected-services",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSLinkedVPCListLinkedVpcsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSLinkedVPCListLinkedVpcsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LinkedVpcsListResultBindingType)
}

func cloudServiceVMCOnAWSLinkedVPCListLinkedVpcsRestMetadata() protocol.OperationRestMetadata {
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
		"/cloud-service/api/v1/infra/linked-vpcs",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cloudServiceVMCOnAWSLinkedVPCUpdateConnectedServiceInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = bindings.NewStringType()
	fields["service_name"] = bindings.NewStringType()
	fields["connected_service_status"] = bindings.NewReferenceType(model.ConnectedServiceStatusBindingType)
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fieldNameMap["service_name"] = "ServiceName"
	fieldNameMap["connected_service_status"] = "ConnectedServiceStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cloudServiceVMCOnAWSLinkedVPCUpdateConnectedServiceOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConnectedServiceStatusBindingType)
}

func cloudServiceVMCOnAWSLinkedVPCUpdateConnectedServiceRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = bindings.NewStringType()
	fields["service_name"] = bindings.NewStringType()
	fields["connected_service_status"] = bindings.NewReferenceType(model.ConnectedServiceStatusBindingType)
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fieldNameMap["service_name"] = "ServiceName"
	fieldNameMap["connected_service_status"] = "ConnectedServiceStatus"
	paramsTypeMap["linked_vpc_id"] = bindings.NewStringType()
	paramsTypeMap["connected_service_status"] = bindings.NewReferenceType(model.ConnectedServiceStatusBindingType)
	paramsTypeMap["service_name"] = bindings.NewStringType()
	paramsTypeMap["linkedVpcId"] = bindings.NewStringType()
	paramsTypeMap["serviceName"] = bindings.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
	pathParams["service_name"] = "serviceName"
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
		"connected_service_status",
		"PUT",
		"/cloud-service/api/v1/infra/linked-vpcs/{linkedVpcId}/connected-services/{serviceName}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
