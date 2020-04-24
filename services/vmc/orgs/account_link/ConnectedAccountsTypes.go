/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ConnectedAccounts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package account_link

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func connectedAccountsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["linked_account_path_id"] = bindings.NewStringType()
	fields["force_even_when_sddc_present"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_path_id"] = "LinkedAccountPathId"
	fieldNameMap["force_even_when_sddc_present"] = "ForceEvenWhenSddcPresent"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func connectedAccountsDeleteOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AwsCustomerConnectedAccountBindingType)
}

func connectedAccountsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["linked_account_path_id"] = bindings.NewStringType()
	fields["force_even_when_sddc_present"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_path_id"] = "LinkedAccountPathId"
	fieldNameMap["force_even_when_sddc_present"] = "ForceEvenWhenSddcPresent"
	paramsTypeMap["force_even_when_sddc_present"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["linked_account_path_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["linkedAccountPathId"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["linked_account_path_id"] = "linkedAccountPathId"
	queryParams["force_even_when_sddc_present"] = "forceEvenWhenSddcPresent"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		"/vmc/api/orgs/{org}/account-link/connected-accounts/{linkedAccountPathId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403})
}

func connectedAccountsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func connectedAccountsGetOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(model.AwsCustomerConnectedAccountBindingType), reflect.TypeOf([]model.AwsCustomerConnectedAccount{}))
}

func connectedAccountsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["provider"] = "Provider"
	paramsTypeMap["provider"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	pathParams["org"] = "org"
	queryParams["provider"] = "provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		"/vmc/api/orgs/{org}/account-link/connected-accounts",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403})
}


