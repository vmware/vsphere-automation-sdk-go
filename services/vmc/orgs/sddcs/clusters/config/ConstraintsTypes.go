// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Constraints.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package config

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"reflect"
)

func constraintsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["cluster"] = vapiBindings_.NewStringType()
	fields["expected_num_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["expected_num_hosts"] = "ExpectedNumHosts"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func ConstraintsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcModel.VsanClusterReconfigConstraintsBindingType)
}

func constraintsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["cluster"] = vapiBindings_.NewStringType()
	fields["expected_num_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["expected_num_hosts"] = "ExpectedNumHosts"
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["expected_num_hosts"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	pathParams["cluster"] = "cluster"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["expected_num_hosts"] = "expected_num_hosts"
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/config/constraints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
