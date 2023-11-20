// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Reservation.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package tbrs

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func reservationPostInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vmcModel.SddcStateRequestBindingType))
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_state"] = "SddcState"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ReservationPostOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmcModel.ReservationWindowBindingType), reflect.TypeOf([]vmcModel.ReservationWindow{})), reflect.TypeOf(map[string][]vmcModel.ReservationWindow{}))
}

func reservationPostRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vmcModel.SddcStateRequestBindingType))
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_state"] = "SddcState"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(vmcModel.SddcStateRequestBindingType))
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
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
		"sddc_state",
		"POST",
		"/vmc/api/orgs/{org}/tbrs/reservation",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
