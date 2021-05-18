// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Onboarding.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package sites

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func onboardingCheckconflictInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["config_onboarding_conflict_request"] = bindings.NewReferenceType(model.ConfigOnboardingConflictRequestBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["config_onboarding_conflict_request"] = "ConfigOnboardingConflictRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func onboardingCheckconflictOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConfigOnboardingConflictStatusBindingType)
}

func onboardingCheckconflictRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["config_onboarding_conflict_request"] = bindings.NewReferenceType(model.ConfigOnboardingConflictRequestBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["config_onboarding_conflict_request"] = "ConfigOnboardingConflictRequest"
	paramsTypeMap["config_onboarding_conflict_request"] = bindings.NewReferenceType(model.ConfigOnboardingConflictRequestBindingType)
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
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
		"action=check_conflict",
		"config_onboarding_conflict_request",
		"POST",
		"/global-manager/api/v1/global-infra/sites/{siteId}/onboarding",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func onboardingStartonboardingInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["config_onboarding_request"] = bindings.NewReferenceType(model.ConfigOnboardingRequestBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["config_onboarding_request"] = "ConfigOnboardingRequest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func onboardingStartonboardingOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ConfigOnboardingStatusBindingType)
}

func onboardingStartonboardingRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["config_onboarding_request"] = bindings.NewReferenceType(model.ConfigOnboardingRequestBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["config_onboarding_request"] = "ConfigOnboardingRequest"
	paramsTypeMap["config_onboarding_request"] = bindings.NewReferenceType(model.ConfigOnboardingRequestBindingType)
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
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
		"action=start_onboarding",
		"config_onboarding_request",
		"POST",
		"/global-manager/api/v1/global-infra/sites/{siteId}/onboarding",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
