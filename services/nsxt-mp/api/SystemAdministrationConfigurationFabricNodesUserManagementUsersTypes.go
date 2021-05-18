// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricNodesUserManagementUsers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func systemAdministrationConfigurationFabricNodesUserManagementUsersActivateNodeUserActivateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersActivateNodeUserActivateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersActivateNodeUserActivateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=activate",
		"node_user_password_property",
		"POST",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersAddNodeUserSshKeyAddSshKeyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["ssh_key_properties"] = bindings.NewReferenceType(model.SshKeyPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_properties"] = "SshKeyProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersAddNodeUserSshKeyAddSshKeyOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersAddNodeUserSshKeyAddSshKeyRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["ssh_key_properties"] = bindings.NewReferenceType(model.SshKeyPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_properties"] = "SshKeyProperties"
	paramsTypeMap["ssh_key_properties"] = bindings.NewReferenceType(model.SshKeyPropertiesBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=add_ssh_key",
		"ssh_key_properties",
		"POST",
		"/api/v1/node/users/{userid}/ssh-keys",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersDeactivateNodeUserDeactivateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersDeactivateNodeUserDeactivateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersDeactivateNodeUserDeactivateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=deactivate",
		"",
		"POST",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersDeleteNodeUserSshKeyRemoveSshKeyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["ssh_key_base_properties"] = bindings.NewReferenceType(model.SshKeyBasePropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_base_properties"] = "SshKeyBaseProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersDeleteNodeUserSshKeyRemoveSshKeyOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersDeleteNodeUserSshKeyRemoveSshKeyRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["ssh_key_base_properties"] = bindings.NewReferenceType(model.SshKeyBasePropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_base_properties"] = "SshKeyBaseProperties"
	paramsTypeMap["ssh_key_base_properties"] = bindings.NewReferenceType(model.SshKeyBasePropertiesBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=remove_ssh_key",
		"ssh_key_base_properties",
		"POST",
		"/api/v1/node/users/{userid}/ssh-keys",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUserSshKeysInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUserSshKeysOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SshKeyPropertiesListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUserSshKeysRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"/api/v1/node/users/{userid}/ssh-keys",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUsersInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUsersOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesListResultBindingType)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersListNodeUsersRestMetadata() protocol.OperationRestMetadata {
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
		"/api/v1/node/users",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersReadNodeUserInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersReadNodeUserOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersReadNodeUserRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserOwnPasswordResetOwnPasswordInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserOwnPasswordResetOwnPasswordOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserOwnPasswordResetOwnPasswordRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
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
		"action=reset_own_password",
		"node_user_password_property",
		"POST",
		"/api/v1/node/users",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserPasswordResetPasswordInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserPasswordResetPasswordOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersResetNodeUserPasswordResetPasswordRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_password_property"] = "NodeUserPasswordProperty"
	paramsTypeMap["node_user_password_property"] = bindings.NewReferenceType(model.NodeUserPasswordPropertyBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"action=reset_password",
		"node_user_password_property",
		"POST",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersUpdateNodeUserInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["node_user_properties"] = bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersUpdateNodeUserOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
}

func systemAdministrationConfigurationFabricNodesUserManagementUsersUpdateNodeUserRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = bindings.NewStringType()
	fields["node_user_properties"] = bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["node_user_properties"] = "NodeUserProperties"
	paramsTypeMap["node_user_properties"] = bindings.NewReferenceType(model.NodeUserPropertiesBindingType)
	paramsTypeMap["userid"] = bindings.NewStringType()
	paramsTypeMap["userid"] = bindings.NewStringType()
	pathParams["userid"] = "userid"
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
		"node_user_properties",
		"PUT",
		"/api/v1/node/users/{userid}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
