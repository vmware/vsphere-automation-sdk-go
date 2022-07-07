// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Members.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package mac_sets

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func membersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_set_id"] = bindings.NewStringType()
	fields["m_AC_address_element"] = bindings.NewReferenceType(model.MACAddressElementBindingType)
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["m_AC_address_element"] = "MACAddressElement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MACAddressElementBindingType)
}

func membersCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mac_set_id"] = bindings.NewStringType()
	fields["m_AC_address_element"] = bindings.NewReferenceType(model.MACAddressElementBindingType)
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["m_AC_address_element"] = "MACAddressElement"
	paramsTypeMap["mac_set_id"] = bindings.NewStringType()
	paramsTypeMap["m_AC_address_element"] = bindings.NewReferenceType(model.MACAddressElementBindingType)
	paramsTypeMap["macSetId"] = bindings.NewStringType()
	pathParams["mac_set_id"] = "macSetId"
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
		"m_AC_address_element",
		"POST",
		"/api/v1/mac-sets/{macSetId}/members",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func membersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_set_id"] = bindings.NewStringType()
	fields["mac_address"] = bindings.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["mac_address"] = "MacAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func membersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mac_set_id"] = bindings.NewStringType()
	fields["mac_address"] = bindings.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["mac_address"] = "MacAddress"
	paramsTypeMap["mac_set_id"] = bindings.NewStringType()
	paramsTypeMap["mac_address"] = bindings.NewStringType()
	paramsTypeMap["macSetId"] = bindings.NewStringType()
	paramsTypeMap["macAddress"] = bindings.NewStringType()
	pathParams["mac_address"] = "macAddress"
	pathParams["mac_set_id"] = "macSetId"
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
		"/api/v1/mac-sets/{macSetId}/members/{macAddress}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func membersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_set_id"] = bindings.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MACAddressElementListResultBindingType)
}

func membersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mac_set_id"] = bindings.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	paramsTypeMap["mac_set_id"] = bindings.NewStringType()
	paramsTypeMap["macSetId"] = bindings.NewStringType()
	pathParams["mac_set_id"] = "macSetId"
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
		"/api/v1/mac-sets/{macSetId}/members",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
