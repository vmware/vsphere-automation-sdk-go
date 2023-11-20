// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: StatusSummary.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package upgrade

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``selectionStatus`` of method StatusSummary#get.
const StatusSummary_GET_SELECTION_STATUS_SELECTED = "SELECTED"

// Possible value for ``selectionStatus`` of method StatusSummary#get.
const StatusSummary_GET_SELECTION_STATUS_DESELECTED = "DESELECTED"

// Possible value for ``selectionStatus`` of method StatusSummary#get.
const StatusSummary_GET_SELECTION_STATUS_ALL = "ALL"

func statusSummaryGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["selection_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["show_history"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["selection_status"] = "SelectionStatus"
	fieldNameMap["show_history"] = "ShowHistory"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatusSummaryGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.UpgradeStatusBindingType)
}

func statusSummaryGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["selection_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["show_history"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["selection_status"] = "SelectionStatus"
	fieldNameMap["show_history"] = "ShowHistory"
	paramsTypeMap["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["selection_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["show_history"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["component_type"] = "component_type"
	queryParams["selection_status"] = "selection_status"
	queryParams["show_history"] = "show_history"
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
		"/api/v1/upgrade/status-summary",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
