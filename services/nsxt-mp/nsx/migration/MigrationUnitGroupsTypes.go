// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MigrationUnitGroups.
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

func migrationUnitGroupsAddmigrationunitsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fields["migration_unit_list"] = bindings.NewReferenceType(model.MigrationUnitListBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["migration_unit_list"] = "MigrationUnitList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsAddmigrationunitsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitListBindingType)
}

func migrationUnitGroupsAddmigrationunitsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fields["migration_unit_list"] = bindings.NewReferenceType(model.MigrationUnitListBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["migration_unit_list"] = "MigrationUnitList"
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["migration_unit_list"] = bindings.NewReferenceType(model.MigrationUnitListBindingType)
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
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
		"action=add_migration_units",
		"migration_unit_list",
		"POST",
		"/api/v1/migration/migration-unit-groups/{groupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitGroupsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["migration_unit_group"] = bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
	fieldNameMap["migration_unit_group"] = "MigrationUnitGroup"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
}

func migrationUnitGroupsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["migration_unit_group"] = bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
	fieldNameMap["migration_unit_group"] = "MigrationUnitGroup"
	paramsTypeMap["migration_unit_group"] = bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
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
		"migration_unit_group",
		"POST",
		"/api/v1/migration/migration-unit-groups",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitGroupsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fieldNameMap["group_id"] = "GroupId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func migrationUnitGroupsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fieldNameMap["group_id"] = "GroupId"
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
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
		"DELETE",
		"/api/v1/migration/migration-unit-groups/{groupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitGroupsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fields["summary"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["summary"] = "Summary"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
}

func migrationUnitGroupsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fields["summary"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["summary"] = "Summary"
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["summary"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	queryParams["summary"] = "summary"
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
		"/api/v1/migration/migration-unit-groups/{groupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitGroupsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["summary"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sync"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["summary"] = "Summary"
	fieldNameMap["sync"] = "Sync"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitGroupListResultBindingType)
}

func migrationUnitGroupsListRestMetadata() protocol.OperationRestMetadata {
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
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["summary"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sync"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["component_type"] = "ComponentType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["summary"] = "Summary"
	fieldNameMap["sync"] = "Sync"
	paramsTypeMap["sync"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["component_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["summary"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["summary"] = "summary"
	queryParams["component_type"] = "component_type"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["sync"] = "sync"
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
		"/api/v1/migration/migration-unit-groups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitGroupsReorderInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fields["reorder_migration_request"] = bindings.NewReferenceType(model.ReorderMigrationRequestBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["reorder_migration_request"] = "ReorderMigrationRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsReorderOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func migrationUnitGroupsReorderRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fields["reorder_migration_request"] = bindings.NewReferenceType(model.ReorderMigrationRequestBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["reorder_migration_request"] = "ReorderMigrationRequest"
	paramsTypeMap["reorder_migration_request"] = bindings.NewReferenceType(model.ReorderMigrationRequestBindingType)
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
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
		"action=reorder",
		"reorder_migration_request",
		"POST",
		"/api/v1/migration/migration-unit-groups/{groupId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func migrationUnitGroupsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fields["migration_unit_group"] = bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["migration_unit_group"] = "MigrationUnitGroup"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitGroupsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
}

func migrationUnitGroupsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fields["migration_unit_group"] = bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["migration_unit_group"] = "MigrationUnitGroup"
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["migration_unit_group"] = bindings.NewReferenceType(model.MigrationUnitGroupBindingType)
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
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
		"migration_unit_group",
		"PUT",
		"/api/v1/migration/migration-unit-groups/{groupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
