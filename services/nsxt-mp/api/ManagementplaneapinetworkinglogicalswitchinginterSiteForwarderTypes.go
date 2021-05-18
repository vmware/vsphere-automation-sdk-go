// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementplaneapinetworkinglogicalswitchinginterSiteForwarder.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``source`` of method ManagementplaneapinetworkinglogicalswitchinginterSiteForwarder#getL2ForwarderStatus.
const ManagementplaneapinetworkinglogicalswitchinginterSiteForwarder_GET_L2FORWARDER_STATUS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementplaneapinetworkinglogicalswitchinginterSiteForwarder#getL2ForwarderStatus.
const ManagementplaneapinetworkinglogicalswitchinginterSiteForwarder_GET_L2FORWARDER_STATUS_SOURCE_CACHED = "cached"

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderRemoteMacsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_switch_id"] = bindings.NewStringType()
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderRemoteMacsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L2ForwarderRemoteMacsBindingType)
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderRemoteMacsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_switch_id"] = bindings.NewStringType()
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	paramsTypeMap["logical_switch_id"] = bindings.NewStringType()
	paramsTypeMap["logicalSwitchId"] = bindings.NewStringType()
	pathParams["logical_switch_id"] = "logicalSwitchId"
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
		"/api/v1/logical-switches/{logicalSwitchId}/inter-site-forwarder/site-span-info",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_switch_id"] = bindings.NewStringType()
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L2ForwarderStatisticsBindingType)
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_switch_id"] = bindings.NewStringType()
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	paramsTypeMap["logical_switch_id"] = bindings.NewStringType()
	paramsTypeMap["logicalSwitchId"] = bindings.NewStringType()
	pathParams["logical_switch_id"] = "logicalSwitchId"
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
		"/api/v1/logical-switches/{logicalSwitchId}/inter-site-forwarder/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_switch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L2ForwarderStatusBindingType)
}

func managementplaneapinetworkinglogicalswitchinginterSiteForwarderGetL2ForwarderStatusRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_switch_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["logical_switch_id"] = "LogicalSwitchId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["logical_switch_id"] = bindings.NewStringType()
	paramsTypeMap["logicalSwitchId"] = bindings.NewStringType()
	pathParams["logical_switch_id"] = "logicalSwitchId"
	queryParams["source"] = "source"
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
		"/api/v1/logical-switches/{logicalSwitchId}/inter-site-forwarder/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
