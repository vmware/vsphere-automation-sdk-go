// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: IpBlockConstraints.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package ip_blocks

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func ipBlockConstraintsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func IpBlockConstraintsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func ipBlockConstraintsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	paramsTypeMap["ip_block_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["ipBlockId"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraintId"] = vapiBindings_.NewStringType()
	pathParams["constraint_id"] = "constraintId"
	pathParams["ip_block_id"] = "ipBlockId"
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
		"DELETE",
		"/policy/api/v1/global-infra/ip-blocks/{ipBlockId}/ip-block-constraints/{constraintId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipBlockConstraintsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func IpBlockConstraintsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
}

func ipBlockConstraintsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	paramsTypeMap["ip_block_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["ipBlockId"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraintId"] = vapiBindings_.NewStringType()
	pathParams["constraint_id"] = "constraintId"
	pathParams["ip_block_id"] = "ipBlockId"
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
		"/policy/api/v1/global-infra/ip-blocks/{ipBlockId}/ip-block-constraints/{constraintId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipBlockConstraintsPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fields["ip_address_block_constraint"] = vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	fieldNameMap["ip_address_block_constraint"] = "IpAddressBlockConstraint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func IpBlockConstraintsPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func ipBlockConstraintsPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fields["ip_address_block_constraint"] = vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	fieldNameMap["ip_address_block_constraint"] = "IpAddressBlockConstraint"
	paramsTypeMap["ip_block_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["ip_address_block_constraint"] = vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
	paramsTypeMap["ipBlockId"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraintId"] = vapiBindings_.NewStringType()
	pathParams["constraint_id"] = "constraintId"
	pathParams["ip_block_id"] = "ipBlockId"
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
		"ip_address_block_constraint",
		"PATCH",
		"/policy/api/v1/global-infra/ip-blocks/{ipBlockId}/ip-block-constraints/{constraintId}",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipBlockConstraintsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fields["ip_address_block_constraint"] = vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	fieldNameMap["ip_address_block_constraint"] = "IpAddressBlockConstraint"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func IpBlockConstraintsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
}

func ipBlockConstraintsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ip_block_id"] = vapiBindings_.NewStringType()
	fields["constraint_id"] = vapiBindings_.NewStringType()
	fields["ip_address_block_constraint"] = vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
	fieldNameMap["ip_block_id"] = "IpBlockId"
	fieldNameMap["constraint_id"] = "ConstraintId"
	fieldNameMap["ip_address_block_constraint"] = "IpAddressBlockConstraint"
	paramsTypeMap["ip_block_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraint_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["ip_address_block_constraint"] = vapiBindings_.NewReferenceType(nsx_policyModel.IpAddressBlockConstraintBindingType)
	paramsTypeMap["ipBlockId"] = vapiBindings_.NewStringType()
	paramsTypeMap["constraintId"] = vapiBindings_.NewStringType()
	pathParams["constraint_id"] = "constraintId"
	pathParams["ip_block_id"] = "ipBlockId"
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
		"ip_address_block_constraint",
		"PUT",
		"/policy/api/v1/global-infra/ip-blocks/{ipBlockId}/ip-block-constraints/{constraintId}",
		"application/json",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
