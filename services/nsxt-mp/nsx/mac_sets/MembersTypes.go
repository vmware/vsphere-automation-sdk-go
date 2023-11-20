// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Members.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package mac_sets

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func membersCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_set_id"] = vapiBindings_.NewStringType()
	fields["m_AC_address_element"] = vapiBindings_.NewReferenceType(nsxModel.MACAddressElementBindingType)
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["m_AC_address_element"] = "MACAddressElement"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MembersCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MACAddressElementBindingType)
}

func membersCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mac_set_id"] = vapiBindings_.NewStringType()
	fields["m_AC_address_element"] = vapiBindings_.NewReferenceType(nsxModel.MACAddressElementBindingType)
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["m_AC_address_element"] = "MACAddressElement"
	paramsTypeMap["mac_set_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["m_AC_address_element"] = vapiBindings_.NewReferenceType(nsxModel.MACAddressElementBindingType)
	paramsTypeMap["macSetId"] = vapiBindings_.NewStringType()
	pathParams["mac_set_id"] = "macSetId"
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

func membersDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_set_id"] = vapiBindings_.NewStringType()
	fields["mac_address"] = vapiBindings_.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["mac_address"] = "MacAddress"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MembersDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func membersDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mac_set_id"] = vapiBindings_.NewStringType()
	fields["mac_address"] = vapiBindings_.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	fieldNameMap["mac_address"] = "MacAddress"
	paramsTypeMap["mac_set_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["mac_address"] = vapiBindings_.NewStringType()
	paramsTypeMap["macSetId"] = vapiBindings_.NewStringType()
	paramsTypeMap["macAddress"] = vapiBindings_.NewStringType()
	pathParams["mac_address"] = "macAddress"
	pathParams["mac_set_id"] = "macSetId"
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
		"DELETE",
		"/api/v1/mac-sets/{macSetId}/members/{macAddress}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func membersListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mac_set_id"] = vapiBindings_.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MembersListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MACAddressElementListResultBindingType)
}

func membersListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mac_set_id"] = vapiBindings_.NewStringType()
	fieldNameMap["mac_set_id"] = "MacSetId"
	paramsTypeMap["mac_set_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["macSetId"] = vapiBindings_.NewStringType()
	pathParams["mac_set_id"] = "macSetId"
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
		"/api/v1/mac-sets/{macSetId}/members",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
