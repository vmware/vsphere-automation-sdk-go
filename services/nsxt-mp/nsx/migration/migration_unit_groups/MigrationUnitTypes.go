// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MigrationUnit.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration_unit_groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func migrationUnitReorderInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fields["migration_unit_id"] = bindings.NewStringType()
	fields["reorder_migration_request"] = bindings.NewReferenceType(model.ReorderMigrationRequestBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["migration_unit_id"] = "MigrationUnitId"
	fieldNameMap["reorder_migration_request"] = "ReorderMigrationRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func migrationUnitReorderOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func migrationUnitReorderRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fields["migration_unit_id"] = bindings.NewStringType()
	fields["reorder_migration_request"] = bindings.NewReferenceType(model.ReorderMigrationRequestBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["migration_unit_id"] = "MigrationUnitId"
	fieldNameMap["reorder_migration_request"] = "ReorderMigrationRequest"
	paramsTypeMap["reorder_migration_request"] = bindings.NewReferenceType(model.ReorderMigrationRequestBindingType)
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["migration_unit_id"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["migrationUnitId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
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
		"action=reorder",
		"reorder_migration_request",
		"POST",
		"/api/v1/migration/migration-unit-groups/{groupId}/migration-unit/{migrationUnitId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
