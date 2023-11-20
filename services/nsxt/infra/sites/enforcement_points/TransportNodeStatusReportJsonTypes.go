// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TransportNodeStatusReportJson.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package enforcement_points

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``source`` of method TransportNodeStatusReportJson#getall.
const TransportNodeStatusReportJson_GETALL_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method TransportNodeStatusReportJson#getall.
const TransportNodeStatusReportJson_GETALL_SOURCE_CACHED = "cached"

// Possible value for ``status`` of method TransportNodeStatusReportJson#getall.
const TransportNodeStatusReportJson_GETALL_STATUS_UP = "UP"

// Possible value for ``status`` of method TransportNodeStatusReportJson#getall.
const TransportNodeStatusReportJson_GETALL_STATUS_DOWN = "DOWN"

// Possible value for ``status`` of method TransportNodeStatusReportJson#getall.
const TransportNodeStatusReportJson_GETALL_STATUS_DEGRADED = "DEGRADED"

func transportNodeStatusReportJsonGetallInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func TransportNodeStatusReportJsonGetallOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.TransportNodeStatusReportListResultBindingType)
}

func transportNodeStatusReportJsonGetallRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = vapiBindings_.NewStringType()
	fields["enforcement_point_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["enforcement_point_id"] = "EnforcementPointId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["status"] = "Status"
	paramsTypeMap["site_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcement_point_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["siteId"] = vapiBindings_.NewStringType()
	paramsTypeMap["enforcementPointId"] = vapiBindings_.NewStringType()
	pathParams["site_id"] = "siteId"
	pathParams["enforcement_point_id"] = "enforcementPointId"
	queryParams["source"] = "source"
	queryParams["status"] = "status"
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
		"/policy/api/v1/infra/sites/{siteId}/enforcement-points/{enforcementPointId}/transport-node-status-report-json",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
