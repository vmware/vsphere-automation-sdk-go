// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Vmgroup.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func vmgroupPostmigrateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["post_vm_group_migration_spec"] = vapiBindings_.NewReferenceType(nsxModel.PostVmGroupMigrationSpecBindingType)
	fieldNameMap["post_vm_group_migration_spec"] = "PostVmGroupMigrationSpec"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VmgroupPostmigrateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func vmgroupPostmigrateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["post_vm_group_migration_spec"] = vapiBindings_.NewReferenceType(nsxModel.PostVmGroupMigrationSpecBindingType)
	fieldNameMap["post_vm_group_migration_spec"] = "PostVmGroupMigrationSpec"
	paramsTypeMap["post_vm_group_migration_spec"] = vapiBindings_.NewReferenceType(nsxModel.PostVmGroupMigrationSpecBindingType)
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
		"action=post_migrate",
		"post_vm_group_migration_spec",
		"POST",
		"/api/v1/migration/vmgroup",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func vmgroupPremigrateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pre_vm_group_migration_spec"] = vapiBindings_.NewReferenceType(nsxModel.PreVmGroupMigrationSpecBindingType)
	fieldNameMap["pre_vm_group_migration_spec"] = "PreVmGroupMigrationSpec"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func VmgroupPremigrateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func vmgroupPremigrateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["pre_vm_group_migration_spec"] = vapiBindings_.NewReferenceType(nsxModel.PreVmGroupMigrationSpecBindingType)
	fieldNameMap["pre_vm_group_migration_spec"] = "PreVmGroupMigrationSpec"
	paramsTypeMap["pre_vm_group_migration_spec"] = vapiBindings_.NewReferenceType(nsxModel.PreVmGroupMigrationSpecBindingType)
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
		"action=pre_migrate",
		"pre_vm_group_migration_spec",
		"POST",
		"/api/v1/migration/vmgroup",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
