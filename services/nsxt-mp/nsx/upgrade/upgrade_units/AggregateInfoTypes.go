// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: AggregateInfo.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package upgrade_units

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``selectionStatus`` of method AggregateInfo#list.
const AggregateInfo_LIST_SELECTION_STATUS_SELECTED = "SELECTED"

// Possible value for ``selectionStatus`` of method AggregateInfo#list.
const AggregateInfo_LIST_SELECTION_STATUS_DESELECTED = "DESELECTED"

// Possible value for ``selectionStatus`` of method AggregateInfo#list.
const AggregateInfo_LIST_SELECTION_STATUS_ALL = "ALL"

func aggregateInfoListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["has_errors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["selection_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["upgrade_unit_display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["has_errors"] = "HasErrors"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["metadata"] = "Metadata"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["selection_status"] = "SelectionStatus"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["upgrade_unit_display_name"] = "UpgradeUnitDisplayName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func aggregateInfoListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UpgradeUnitAggregateInfoListResultBindingType)
}

func aggregateInfoListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["has_errors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["selection_status"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["upgrade_unit_display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["has_errors"] = "HasErrors"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["metadata"] = "Metadata"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["selection_status"] = "SelectionStatus"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["upgrade_unit_display_name"] = "UpgradeUnitDisplayName"
	paramsTypeMap["selection_status"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["has_errors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["upgrade_unit_display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["has_errors"] = "has_errors"
	queryParams["component_type"] = "component_type"
	queryParams["metadata"] = "metadata"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["group_id"] = "group_id"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["selection_status"] = "selection_status"
	queryParams["upgrade_unit_display_name"] = "upgrade_unit_display_name"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/upgrade/upgrade-units/aggregate-info",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
