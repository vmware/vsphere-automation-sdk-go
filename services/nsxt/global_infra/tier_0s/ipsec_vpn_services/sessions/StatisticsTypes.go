// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Statistics.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package sessions

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

// Possible value for ``action`` of method Statistics#create.
const Statistics_CREATE_ACTION_RESET = "reset"

func statisticsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["service_id"] = bindings.NewStringType()
	fields["session_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["session_id"] = "SessionId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statisticsCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func statisticsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["service_id"] = bindings.NewStringType()
	fields["session_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["service_id"] = "ServiceId"
	fieldNameMap["session_id"] = "SessionId"
	fieldNameMap["action"] = "Action"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["service_id"] = bindings.NewStringType()
	paramsTypeMap["session_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["serviceId"] = bindings.NewStringType()
	paramsTypeMap["sessionId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["session_id"] = "sessionId"
	pathParams["service_id"] = "serviceId"
	queryParams["action"] = "action"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"POST",
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/ipsec-vpn-services/{serviceId}/sessions/{sessionId}/statistics",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
