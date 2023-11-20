// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ConnectedAccounts.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package account_link

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func connectedAccountsDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["linked_account_path_id"] = vapiBindings_.NewStringType()
	fields["force_even_when_sddc_present"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_path_id"] = "LinkedAccountPathId"
	fieldNameMap["force_even_when_sddc_present"] = "ForceEvenWhenSddcPresent"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectedAccountsDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.AwsCustomerConnectedAccountBindingType)
}

func connectedAccountsDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["linked_account_path_id"] = vapiBindings_.NewStringType()
	fields["force_even_when_sddc_present"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_path_id"] = "LinkedAccountPathId"
	fieldNameMap["force_even_when_sddc_present"] = "ForceEvenWhenSddcPresent"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["force_even_when_sddc_present"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["linked_account_path_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["linkedAccountPathId"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["linked_account_path_id"] = "linkedAccountPathId"
	queryParams["force_even_when_sddc_present"] = "forceEvenWhenSddcPresent"
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
		"DELETE",
		"/vmc/api/orgs/{org}/account-link/connected-accounts/{linkedAccountPathId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func connectedAccountsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["provider"] = "Provider"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConnectedAccountsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmcModel.AwsCustomerConnectedAccountBindingType), reflect.TypeOf([]vmcModel.AwsCustomerConnectedAccount{}))
}

func connectedAccountsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["provider"] = "Provider"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	queryParams["provider"] = "provider"
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
		"/vmc/api/orgs/{org}/account-link/connected-accounts",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}
