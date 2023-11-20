// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: FloodProtectionProfileBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package tier_1s

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func floodProtectionProfileBindingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func FloodProtectionProfileBindingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.FloodProtectionProfileBindingMapBindingType)
}

func floodProtectionProfileBindingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = vapiBindings_.NewStringType()
	fields["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["flood_protection_profile_binding_id"] = "FloodProtectionProfileBindingId"
	paramsTypeMap["tier1_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["flood_protection_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier1Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["floodProtectionProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["flood_protection_profile_binding_id"] = "floodProtectionProfileBindingId"
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
		"/policy/api/v1/global-infra/tier-1s/{tier1Id}/flood-protection-profile-bindings/{floodProtectionProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
