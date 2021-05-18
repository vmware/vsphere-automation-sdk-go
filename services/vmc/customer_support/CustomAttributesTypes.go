// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CustomAttributes.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package customer_support

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func customAttributesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["inputs_sddc"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["inputs_sddc"] = "InputsSddc"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func customAttributesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SddcCustomAttributesBindingType)
}

func customAttributesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["user_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["inputs_sddc"] = bindings.NewStringType()
	fieldNameMap["user_id"] = "UserId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["inputs_sddc"] = "InputsSddc"
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["inputs_sddc"] = bindings.NewStringType()
	paramsTypeMap["user_id"] = bindings.NewStringType()
	queryParams["inputs_sddc"] = "inputs-sddc"
	queryParams["user_id"] = "userId"
	queryParams["org_id"] = "orgId"
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
		"/vmc/api/customer-support/custom-attributes",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
