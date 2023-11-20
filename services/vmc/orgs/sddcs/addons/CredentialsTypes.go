// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Credentials.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package addons

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

// Possible value for ``addonType`` of method Credentials#create.
const Credentials_CREATE_ADDON_TYPE_HCX = "HCX"

// Possible value for ``addonType`` of method Credentials#get.
const Credentials_GET_ADDON_TYPE_HCX = "HCX"

// Possible value for ``addonType`` of method Credentials#list.
const Credentials_LIST_ADDON_TYPE_HCX = "HCX"

// Possible value for ``addonType`` of method Credentials#update.
const Credentials_UPDATE_ADDON_TYPE_HCX = "HCX"

func credentialsCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fields["credentials"] = vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["credentials"] = "Credentials"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CredentialsCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType)
}

func credentialsCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fields["credentials"] = vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["credentials"] = "Credentials"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["credentials"] = vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType)
	paramsTypeMap["addon_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["addonType"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
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
		"credentials",
		"POST",
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func credentialsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["name"] = "Name"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CredentialsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType)
}

func credentialsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["name"] = "Name"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["addon_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["addonType"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
	pathParams["name"] = "name"
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
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthorized": 403})
}

func credentialsListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CredentialsListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewListType(vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType), reflect.TypeOf([]vmcModel.NewCredentials{}))
}

func credentialsListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["addon_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["addonType"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
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
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthorized": 403})
}

func credentialsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fields["name"] = vapiBindings_.NewStringType()
	fields["credentials"] = vapiBindings_.NewReferenceType(vmcModel.UpdateCredentialsBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["name"] = "Name"
	fieldNameMap["credentials"] = "Credentials"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CredentialsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.NewCredentialsBindingType)
}

func credentialsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc_id"] = vapiBindings_.NewStringType()
	fields["addon_type"] = vapiBindings_.NewStringType()
	fields["name"] = vapiBindings_.NewStringType()
	fields["credentials"] = vapiBindings_.NewReferenceType(vmcModel.UpdateCredentialsBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_id"] = "SddcId"
	fieldNameMap["addon_type"] = "AddonType"
	fieldNameMap["name"] = "Name"
	fieldNameMap["credentials"] = "Credentials"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["credentials"] = vapiBindings_.NewReferenceType(vmcModel.UpdateCredentialsBindingType)
	paramsTypeMap["addon_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddcId"] = vapiBindings_.NewStringType()
	paramsTypeMap["addonType"] = vapiBindings_.NewStringType()
	paramsTypeMap["name"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc_id"] = "sddcId"
	pathParams["addon_type"] = "addonType"
	pathParams["name"] = "name"
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
		"credentials",
		"PUT",
		"/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
