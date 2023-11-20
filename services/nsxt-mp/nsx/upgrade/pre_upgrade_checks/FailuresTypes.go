// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Failures.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package pre_upgrade_checks

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``originType`` of method Failures#list.
const Failures_LIST_ORIGIN_TYPE_COMPONENT = "COMPONENT"

// Possible value for ``originType`` of method Failures#list.
const Failures_LIST_ORIGIN_TYPE_UPGRADE_UNIT = "UPGRADE_UNIT"

// Possible value for ``type`` of method Failures#list.
const Failures_LIST_TYPE_FAILURE = "FAILURE"

// Possible value for ``type`` of method Failures#list.
const Failures_LIST_TYPE_WARNING = "WARNING"

func failuresListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_text"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["needs_ack"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["origin_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["unit_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["unit_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["filter_text"] = "FilterText"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_name"] = "GroupName"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["needs_ack"] = "NeedsAck"
	fieldNameMap["origin_type"] = "OriginType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	fieldNameMap["unit_id"] = "UnitId"
	fieldNameMap["unit_name"] = "UnitName"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FailuresListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.UpgradeCheckFailureListResultBindingType)
}

func failuresListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["filter_text"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["needs_ack"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["origin_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["unit_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["unit_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["filter_text"] = "FilterText"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["group_name"] = "GroupName"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["needs_ack"] = "NeedsAck"
	fieldNameMap["origin_type"] = "OriginType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	fieldNameMap["unit_id"] = "UnitId"
	fieldNameMap["unit_name"] = "UnitName"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["filter_text"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["group_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["origin_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["unit_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["component_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["group_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["unit_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["needs_ack"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["filter_text"] = "filter_text"
	queryParams["group_name"] = "group_name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["origin_type"] = "origin_type"
	queryParams["unit_name"] = "unit_name"
	queryParams["component_type"] = "component_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["group_id"] = "group_id"
	queryParams["included_fields"] = "included_fields"
	queryParams["unit_id"] = "unit_id"
	queryParams["needs_ack"] = "needs_ack"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/upgrade/pre-upgrade-checks/failures",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
