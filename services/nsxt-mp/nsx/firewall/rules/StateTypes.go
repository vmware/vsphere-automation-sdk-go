// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: State.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package rules

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func stateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rule_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stateGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RuleStateBindingType)
}

func stateGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["rule_id"] = bindings.NewStringType()
	fields["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["rule_id"] = "RuleId"
	fieldNameMap["barrier_id"] = "BarrierId"
	fieldNameMap["request_id"] = "RequestId"
	paramsTypeMap["rule_id"] = bindings.NewStringType()
	paramsTypeMap["barrier_id"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ruleId"] = bindings.NewStringType()
	pathParams["rule_id"] = "ruleId"
	queryParams["barrier_id"] = "barrier_id"
	queryParams["request_id"] = "request_id"
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
		"/api/v1/firewall/rules/{ruleId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
