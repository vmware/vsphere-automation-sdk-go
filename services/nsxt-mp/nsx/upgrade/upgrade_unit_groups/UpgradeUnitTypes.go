// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: UpgradeUnit.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package upgrade_unit_groups

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func upgradeUnitReorderInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["upgrade_unit_id"] = vapiBindings_.NewStringType()
	fields["reorder_request"] = vapiBindings_.NewReferenceType(nsxModel.ReorderRequestBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["upgrade_unit_id"] = "UpgradeUnitId"
	fieldNameMap["reorder_request"] = "ReorderRequest"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UpgradeUnitReorderOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func upgradeUnitReorderRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = vapiBindings_.NewStringType()
	fields["upgrade_unit_id"] = vapiBindings_.NewStringType()
	fields["reorder_request"] = vapiBindings_.NewReferenceType(nsxModel.ReorderRequestBindingType)
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["upgrade_unit_id"] = "UpgradeUnitId"
	fieldNameMap["reorder_request"] = "ReorderRequest"
	paramsTypeMap["reorder_request"] = vapiBindings_.NewReferenceType(nsxModel.ReorderRequestBindingType)
	paramsTypeMap["group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["upgrade_unit_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["groupId"] = vapiBindings_.NewStringType()
	paramsTypeMap["upgradeUnitId"] = vapiBindings_.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["upgrade_unit_id"] = "upgradeUnitId"
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
		"action=reorder",
		"reorder_request",
		"POST",
		"/api/v1/upgrade/upgrade-unit-groups/{groupId}/upgrade-unit/{upgradeUnitId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
