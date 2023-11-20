// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SessionTimerProfileBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package locale_services

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func sessionTimerProfileBindingsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["session_timer_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SessionTimerProfileBindingsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.SessionTimerProfileBindingMapBindingType)
}

func sessionTimerProfileBindingsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = vapiBindings_.NewStringType()
	fields["locale_services_id"] = vapiBindings_.NewStringType()
	fields["session_timer_profile_binding_id"] = vapiBindings_.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_services_id"] = "LocaleServicesId"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	paramsTypeMap["tier0_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["locale_services_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["tier0Id"] = vapiBindings_.NewStringType()
	paramsTypeMap["localeServicesId"] = vapiBindings_.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = vapiBindings_.NewStringType()
	pathParams["locale_services_id"] = "localeServicesId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
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
		"/policy/api/v1/global-infra/tier-0s/{tier0Id}/locale-services/{localeServicesId}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
