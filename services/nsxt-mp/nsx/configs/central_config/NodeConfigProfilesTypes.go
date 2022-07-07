// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: NodeConfigProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package central_config

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func nodeConfigProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile_id"] = bindings.NewStringType()
	fields["show_sensitive_data"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["profile_id"] = "ProfileId"
	fieldNameMap["show_sensitive_data"] = "ShowSensitiveData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodeConfigProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
}

func nodeConfigProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["profile_id"] = bindings.NewStringType()
	fields["show_sensitive_data"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["profile_id"] = "ProfileId"
	fieldNameMap["show_sensitive_data"] = "ShowSensitiveData"
	paramsTypeMap["profile_id"] = bindings.NewStringType()
	paramsTypeMap["show_sensitive_data"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["profileId"] = bindings.NewStringType()
	pathParams["profile_id"] = "profileId"
	queryParams["show_sensitive_data"] = "show_sensitive_data"
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
		"/api/v1/configs/central-config/node-config-profiles/{profileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeConfigProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodeConfigProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CentralNodeConfigProfileListResultBindingType)
}

func nodeConfigProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/configs/central-config/node-config-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func nodeConfigProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_config_profile_id"] = bindings.NewStringType()
	fields["central_node_config_profile"] = bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
	fieldNameMap["node_config_profile_id"] = "NodeConfigProfileId"
	fieldNameMap["central_node_config_profile"] = "CentralNodeConfigProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodeConfigProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
}

func nodeConfigProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_config_profile_id"] = bindings.NewStringType()
	fields["central_node_config_profile"] = bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
	fieldNameMap["node_config_profile_id"] = "NodeConfigProfileId"
	fieldNameMap["central_node_config_profile"] = "CentralNodeConfigProfile"
	paramsTypeMap["central_node_config_profile"] = bindings.NewReferenceType(model.CentralNodeConfigProfileBindingType)
	paramsTypeMap["node_config_profile_id"] = bindings.NewStringType()
	paramsTypeMap["nodeConfigProfileId"] = bindings.NewStringType()
	pathParams["node_config_profile_id"] = "nodeConfigProfileId"
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
		"central_node_config_profile",
		"PUT",
		"/api/v1/configs/central-config/node-config-profiles/{nodeConfigProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
