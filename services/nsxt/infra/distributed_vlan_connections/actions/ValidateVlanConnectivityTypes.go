// Copyright (c) 2019-2026 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ValidateVlanConnectivity.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package actions

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func validateVlanConnectivityCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["distributed_vlan_connection_id"] = vapiBindings_.NewStringType()
	fields["distributed_vlan_connection_validation_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVlanConnectionValidationConfigBindingType)
	fieldNameMap["distributed_vlan_connection_id"] = "DistributedVlanConnectionId"
	fieldNameMap["distributed_vlan_connection_validation_config"] = "DistributedVlanConnectionValidationConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ValidateVlanConnectivityCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func validateVlanConnectivityCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["distributed_vlan_connection_id"] = vapiBindings_.NewStringType()
	fields["distributed_vlan_connection_validation_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVlanConnectionValidationConfigBindingType)
	fieldNameMap["distributed_vlan_connection_id"] = "DistributedVlanConnectionId"
	fieldNameMap["distributed_vlan_connection_validation_config"] = "DistributedVlanConnectionValidationConfig"
	paramsTypeMap["distributed_vlan_connection_validation_config"] = vapiBindings_.NewReferenceType(nsx_policyModel.DistributedVlanConnectionValidationConfigBindingType)
	paramsTypeMap["distributed_vlan_connection_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["distributedVlanConnectionId"] = vapiBindings_.NewStringType()
	pathParams["distributed_vlan_connection_id"] = "distributedVlanConnectionId"
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
		"distributed_vlan_connection_validation_config",
		"POST",
		"/policy/api/v1/infra/distributed-vlan-connections/{distributedVlanConnectionId}/actions/validate-vlan-connectivity",
		"application/json",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
