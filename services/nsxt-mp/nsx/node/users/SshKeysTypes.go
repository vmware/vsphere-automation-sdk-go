// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SshKeys.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package users

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func sshKeysAddsshkeyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["ssh_key_properties"] = bindings.NewReferenceType(model.SshKeyPropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_properties"] = "SshKeyProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sshKeysAddsshkeyOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sshKeysAddsshkeyRestMetadata() protocol.OperationRestMetadata {
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

func sshKeysListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fieldNameMap["userid"] = "Userid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sshKeysListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SshKeyPropertiesListResultBindingType)
}

func sshKeysListRestMetadata() protocol.OperationRestMetadata {
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

func sshKeysRemovesshkeyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["userid"] = bindings.NewStringType()
	fields["ssh_key_base_properties"] = bindings.NewReferenceType(model.SshKeyBasePropertiesBindingType)
	fieldNameMap["userid"] = "Userid"
	fieldNameMap["ssh_key_base_properties"] = "SshKeyBaseProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sshKeysRemovesshkeyOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func sshKeysRemovesshkeyRestMetadata() protocol.OperationRestMetadata {
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
