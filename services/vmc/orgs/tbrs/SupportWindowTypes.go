// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SupportWindow.
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

func supportWindowGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["minimum_seats_available"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["created_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["minimum_seats_available"] = "MinimumSeatsAvailable"
	fieldNameMap["created_by"] = "CreatedBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SupportWindowGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmcModel.SupportWindowBindingType), reflect.TypeOf([]vmcModel.SupportWindow{}))
}

func supportWindowGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["minimum_seats_available"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["created_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["minimum_seats_available"] = "MinimumSeatsAvailable"
	fieldNameMap["created_by"] = "CreatedBy"
	paramsTypeMap["minimum_seats_available"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["created_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	queryParams["created_by"] = "createdBy"
	queryParams["minimum_seats_available"] = "minimumSeatsAvailable"
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
		"/vmc/api/orgs/{org}/tbrs/support-window",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}

func supportWindowPutInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["id"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewReferenceType(vmcModel.SddcIdBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["id"] = "Id"
	fieldNameMap["sddc_id"] = "SddcId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func SupportWindowPutOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.SupportWindowIdBindingType)
}

func supportWindowPutRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["id"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewReferenceType(vmcModel.SddcIdBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["id"] = "Id"
	fieldNameMap["sddc_id"] = "SddcId"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_id"] = vapiBindings_.NewReferenceType(vmcModel.SddcIdBindingType)
	paramsTypeMap["id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["id"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["id"] = "id"
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
		"sddc_id",
		"PUT",
		"/vmc/api/orgs/{org}/tbrs/support-window/{id}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}
