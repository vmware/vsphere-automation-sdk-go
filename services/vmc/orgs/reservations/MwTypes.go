// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Mw.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package reservations

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func mwGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["reservation"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MwGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.MaintenanceWindowGetBindingType)
}

func mwGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["reservation"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservation"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservation"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["reservation"] = "reservation"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/vmc/api/orgs/{org}/reservations/{reservation}/mw",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func mwPutInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["reservation"] = vapiBindings_.NewStringType()
	fields["window"] = vapiBindings_.NewReferenceType(vmcModel.MaintenanceWindowBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	fieldNameMap["window"] = "Window"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MwPutOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.MaintenanceWindowBindingType)
}

func mwPutRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["reservation"] = vapiBindings_.NewStringType()
	fields["window"] = vapiBindings_.NewReferenceType(vmcModel.MaintenanceWindowBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	fieldNameMap["window"] = "Window"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservation"] = vapiBindings_.NewStringType()
	paramsTypeMap["window"] = vapiBindings_.NewReferenceType(vmcModel.MaintenanceWindowBindingType)
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["reservation"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["reservation"] = "reservation"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"window",
		"PUT",
		"/vmc/api/orgs/{org}/reservations/{reservation}/mw",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
