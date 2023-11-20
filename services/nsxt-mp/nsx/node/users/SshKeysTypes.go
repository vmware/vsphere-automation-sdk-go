// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SshKeys.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package users

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func sshKeysAddsshkeyInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["ssh_key_properties"] = vapiBindings_.NewReferenceType(nsxModel.SshKeyPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_properties"] = "SshKeyProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SshKeysAddsshkeyOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sshKeysAddsshkeyRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["ssh_key_properties"] = vapiBindings_.NewReferenceType(nsxModel.SshKeyPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_properties"] = "SshKeyProperties"
	paramsTypeMap["ssh_key_properties"] = vapiBindings_.NewReferenceType(nsxModel.SshKeyPropertiesBindingType)
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["userid"] = "userid"
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

func sshKeysListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SshKeysListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.SshKeyPropertiesListResultBindingType)
}

func sshKeysListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fieldNameMap["userid"] = "Userid"
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["userid"] = "userid"
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
		"/api/v1/node/users/{userid}/ssh-keys",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func sshKeysRemovesshkeyInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = vapiBindings_.NewStringType()
	fields["ssh_key_base_properties"] = vapiBindings_.NewReferenceType(nsxModel.SshKeyBasePropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_base_properties"] = "SshKeyBaseProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SshKeysRemovesshkeyOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func sshKeysRemovesshkeyRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["userid"] = vapiBindings_.NewStringType()
	fields["ssh_key_base_properties"] = vapiBindings_.NewReferenceType(nsxModel.SshKeyBasePropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_base_properties"] = "SshKeyBaseProperties"
	paramsTypeMap["ssh_key_base_properties"] = vapiBindings_.NewReferenceType(nsxModel.SshKeyBasePropertiesBindingType)
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	paramsTypeMap["userid"] = vapiBindings_.NewStringType()
	pathParams["userid"] = "userid"
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
