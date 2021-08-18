// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SessionTimerProfileBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package tier_1s

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func sessionTimerProfileBindingsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sessionTimerProfileBindingsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["org_id"] = "orgId"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sessionTimerProfileBindingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
}

func sessionTimerProfileBindingsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["org_id"] = "orgId"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
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
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sessionTimerProfileBindingsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sessionTimerProfileBindingsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["org_id"] = "orgId"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
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
		"session_timer_profile_binding_map",
		"PATCH",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sessionTimerProfileBindingsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionTimerProfileBindingsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
}

func sessionTimerProfileBindingsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["deployment_group_id"] = bindings.NewStringType()
	fields["vhc_id"] = bindings.NewStringType()
	fields["tier1_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_id"] = bindings.NewStringType()
	fields["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	fieldNameMap["vhc_id"] = "VhcId"
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["session_timer_profile_binding_id"] = "SessionTimerProfileBindingId"
	fieldNameMap["session_timer_profile_binding_map"] = "SessionTimerProfileBindingMap"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["vhc_id"] = bindings.NewStringType()
	paramsTypeMap["session_timer_profile_binding_map"] = bindings.NewReferenceType(model.SessionTimerProfileBindingMapBindingType)
	paramsTypeMap["deployment_group_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["deploymentGroupId"] = bindings.NewStringType()
	paramsTypeMap["vhcId"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["sessionTimerProfileBindingId"] = bindings.NewStringType()
	pathParams["vhc_id"] = "vhcId"
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["org_id"] = "orgId"
	pathParams["session_timer_profile_binding_id"] = "sessionTimerProfileBindingId"
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
		"session_timer_profile_binding_map",
		"PUT",
		"/policy/api/v1/org-root/orgs/{orgId}/deployment-groups/{deploymentGroupId}/vhcs/{vhcId}/infra/tier-1s/{tier1Id}/session-timer-profile-bindings/{sessionTimerProfileBindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
