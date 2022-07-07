// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MigrationUnits.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func migrationUnitsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["migration_unit_id"] = bindings.NewStringType()
	fieldNameMap["migration_unit_id"] = "MigrationUnitId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitBindingType)
}

func migrationUnitsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["migration_unit_id"] = bindings.NewStringType()
	fieldNameMap["migration_unit_id"] = "MigrationUnitId"
	paramsTypeMap["migration_unit_id"] = bindings.NewStringType()
	paramsTypeMap["migrationUnitId"] = bindings.NewStringType()
	pathParams["migration_unit_id"] = "migrationUnitId"
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
		"/api/v1/migration/migration-units/{migrationUnitId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["current_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["has_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["migration_unit_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["current_version"] = "CurrentVersion"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["has_warnings"] = "HasWarnings"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["metadata"] = "Metadata"
	fieldNameMap["migration_unit_type"] = "MigrationUnitType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitListResultBindingType)
}

func migrationUnitsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["current_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["has_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["migration_unit_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["current_version"] = "CurrentVersion"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["has_warnings"] = "HasWarnings"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["metadata"] = "Metadata"
	fieldNameMap["migration_unit_type"] = "MigrationUnitType"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["has_warnings"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["current_version"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["migration_unit_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["group_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["component_type"] = "component_type"
	queryParams["migration_unit_type"] = "migration_unit_type"
	queryParams["metadata"] = "metadata"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["group_id"] = "group_id"
	queryParams["included_fields"] = "included_fields"
	queryParams["current_version"] = "current_version"
	queryParams["has_warnings"] = "has_warnings"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/migration/migration-units",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
