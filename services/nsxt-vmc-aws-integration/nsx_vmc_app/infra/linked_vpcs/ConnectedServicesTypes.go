// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ConnectedServices.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package linked_vpcs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_vmc_appModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-vmc-aws-integration/nsx_vmc_app/model"
	"reflect"
)

func connectedServicesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectedServicesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.ConnectedServiceListResultBindingType)
}

func connectedServicesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	paramsTypeMap["linked_vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["linkedVpcId"] = vapiBindings_.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
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
		"/cloud-service/api/v1/infra/linked-vpcs/{linkedVpcId}/connected-services",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func connectedServicesUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fields["service_name"] = vapiBindings_.NewStringType()
	fields["connected_service_status"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ConnectedServiceStatusBindingType)
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fieldNameMap["service_name"] = "ServiceName"
	fieldNameMap["connected_service_status"] = "ConnectedServiceStatus"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectedServicesUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_vmc_appModel.ConnectedServiceStatusBindingType)
}

func connectedServicesUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["linked_vpc_id"] = vapiBindings_.NewStringType()
	fields["service_name"] = vapiBindings_.NewStringType()
	fields["connected_service_status"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ConnectedServiceStatusBindingType)
	fieldNameMap["linked_vpc_id"] = "LinkedVpcId"
	fieldNameMap["service_name"] = "ServiceName"
	fieldNameMap["connected_service_status"] = "ConnectedServiceStatus"
	paramsTypeMap["connected_service_status"] = vapiBindings_.NewReferenceType(nsx_vmc_appModel.ConnectedServiceStatusBindingType)
	paramsTypeMap["linked_vpc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["service_name"] = vapiBindings_.NewStringType()
	paramsTypeMap["linkedVpcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["serviceName"] = vapiBindings_.NewStringType()
	pathParams["linked_vpc_id"] = "linkedVpcId"
	pathParams["service_name"] = "serviceName"
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
