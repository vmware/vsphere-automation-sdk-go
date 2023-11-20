// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CompatibleSubnets.
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

func compatibleSubnetsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["linked_account_id"] = vapiBindings_.NewStringType()
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sddc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["force_refresh"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sddc_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["num_of_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_id"] = "LinkedAccountId"
	fieldNameMap["region"] = "Region"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["force_refresh"] = "ForceRefresh"
	fieldNameMap["instance_type"] = "InstanceType"
	fieldNameMap["sddc_type"] = "SddcType"
	fieldNameMap["num_of_hosts"] = "NumOfHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CompatibleSubnetsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.AwsCompatibleSubnetsBindingType)
}

func compatibleSubnetsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["linked_account_id"] = vapiBindings_.NewStringType()
	fields["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sddc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["force_refresh"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sddc_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["num_of_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_id"] = "LinkedAccountId"
	fieldNameMap["region"] = "Region"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["force_refresh"] = "ForceRefresh"
	fieldNameMap["instance_type"] = "InstanceType"
	fieldNameMap["sddc_type"] = "SddcType"
	fieldNameMap["num_of_hosts"] = "NumOfHosts"
	paramsTypeMap["num_of_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["linked_account_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["region"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["force_refresh"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["instance_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	queryParams["linked_account_id"] = "linkedAccountId"
	queryParams["num_of_hosts"] = "numOfHosts"
	queryParams["sddc"] = "sddc"
	queryParams["instance_type"] = "instanceType"
	queryParams["region"] = "region"
	queryParams["sddc_type"] = "sddcType"
	queryParams["force_refresh"] = "forceRefresh"
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
		"/vmc/api/orgs/{org}/account-link/compatible-subnets",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func compatibleSubnetsPostInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CompatibleSubnetsPostOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.AwsSubnetBindingType)
}

func compatibleSubnetsPostRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
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
		"",
		"POST",
		"/vmc/api/orgs/{org}/account-link/compatible-subnets",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}
