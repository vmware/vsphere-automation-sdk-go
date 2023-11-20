// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DetailedStatus.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package lb_services

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``source`` of method DetailedStatus#get.
const DetailedStatus_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method DetailedStatus#get.
const DetailedStatus_GET_SOURCE_CACHED = "cached"

func detailedStatusGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lb_service_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_instance_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lb_service_id"] = "LbServiceId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["include_instance_details"] = "IncludeInstanceDetails"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_ids"] = "TransportNodeIds"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DetailedStatusGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.AggregateLBServiceStatusBindingType)
}

func detailedStatusGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lb_service_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["include_instance_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["lb_service_id"] = "LbServiceId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["include_instance_details"] = "IncludeInstanceDetails"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_ids"] = "TransportNodeIds"
	paramsTypeMap["include_instance_details"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["transport_node_ids"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["lb_service_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["lbServiceId"] = vapiBindings_.NewStringType()
	pathParams["lb_service_id"] = "lbServiceId"
	queryParams["include_instance_details"] = "include_instance_details"
	queryParams["transport_node_ids"] = "transport_node_ids"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/lb-services/{lbServiceId}/detailed-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
