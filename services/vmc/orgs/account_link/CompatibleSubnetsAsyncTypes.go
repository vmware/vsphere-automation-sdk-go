/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CompatibleSubnetsAsync.
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





func compatibleSubnetsAsyncGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["linked_account_id"] = bindings.NewStringType()
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sddc"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["num_of_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_id"] = "LinkedAccountId"
	fieldNameMap["region"] = "Region"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["instance_type"] = "InstanceType"
	fieldNameMap["sddc_type"] = "SddcType"
	fieldNameMap["num_of_hosts"] = "NumOfHosts"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func compatibleSubnetsAsyncGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TaskBindingType)
}

func compatibleSubnetsAsyncGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["linked_account_id"] = bindings.NewStringType()
	fields["region"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sddc"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["num_of_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["linked_account_id"] = "LinkedAccountId"
	fieldNameMap["region"] = "Region"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["instance_type"] = "InstanceType"
	fieldNameMap["sddc_type"] = "SddcType"
	fieldNameMap["num_of_hosts"] = "NumOfHosts"
	paramsTypeMap["sddc"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["instance_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["region"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["linked_account_id"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["num_of_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["sddc_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	pathParams["org"] = "org"
	queryParams["linked_account_id"] = "linkedAccountId"
	queryParams["num_of_hosts"] = "numOfHosts"
	queryParams["sddc"] = "sddc"
	queryParams["instance_type"] = "instanceType"
	queryParams["region"] = "region"
	queryParams["sddc_type"] = "sddcType"
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
		"/vmc/api/orgs/{org}/account-link/compatible-subnets-async",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403})
}

func compatibleSubnetsAsyncPostInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["aws_subnet"] = bindings.NewReferenceType(model.AwsSubnetBindingType)
	fields["org"] = bindings.NewStringType()
	fieldNameMap["aws_subnet"] = "AwsSubnet"
	fieldNameMap["org"] = "Org"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func compatibleSubnetsAsyncPostOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TaskBindingType)
}

func compatibleSubnetsAsyncPostRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["aws_subnet"] = bindings.NewReferenceType(model.AwsSubnetBindingType)
	fields["org"] = bindings.NewStringType()
	fieldNameMap["aws_subnet"] = "AwsSubnet"
	fieldNameMap["org"] = "Org"
	paramsTypeMap["aws_subnet"] = bindings.NewReferenceType(model.AwsSubnetBindingType)
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	pathParams["org"] = "org"
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
		"aws_subnet",
		"POST",
		"/vmc/api/orgs/{org}/account-link/compatible-subnets-async",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401,"com.vmware.vapi.std.errors.unauthorized": 403})
}


