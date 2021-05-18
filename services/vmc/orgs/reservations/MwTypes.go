// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Mw.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package reservations

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func mwGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["reservation"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func mwGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MaintenanceWindowGetBindingType)
}

func mwGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["reservation"] = bindings.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["reservation"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["reservation"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["reservation"] = "reservation"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/vmc/api/orgs/{org}/reservations/{reservation}/mw",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func mwPutInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["reservation"] = bindings.NewStringType()
	fields["window"] = bindings.NewReferenceType(model.MaintenanceWindowBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	fieldNameMap["window"] = "Window"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func mwPutOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MaintenanceWindowBindingType)
}

func mwPutRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["reservation"] = bindings.NewStringType()
	fields["window"] = bindings.NewReferenceType(model.MaintenanceWindowBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["reservation"] = "Reservation"
	fieldNameMap["window"] = "Window"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["reservation"] = bindings.NewStringType()
	paramsTypeMap["window"] = bindings.NewReferenceType(model.MaintenanceWindowBindingType)
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["reservation"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["reservation"] = "reservation"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
