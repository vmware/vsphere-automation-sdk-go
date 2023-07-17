// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RemediationPolicy.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package clusters

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/autoscaler/model"
	"reflect"
)

func remediationPolicyGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["cluster"] = vapiBindings_.NewStringType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RemediationPolicyGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(model.RemediationClusterInfoBindingType)
}

func remediationPolicyGetRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	pathParams["cluster"] = "cluster"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
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
		"/vmc/autoscaler/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/remediation-policy",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func remediationPolicyPatchInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["cluster"] = vapiBindings_.NewStringType()
	fields["remediation_policy_patch"] = vapiBindings_.NewReferenceType(model.RemediationPolicyPatchBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["remediation_policy_patch"] = "RemediationPolicyPatch"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func RemediationPolicyPatchOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(model.TaskBindingType)
}

func remediationPolicyPatchRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["remediation_policy_patch"] = vapiBindings_.NewReferenceType(model.RemediationPolicyPatchBindingType)
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["remediation_policy_patch"] = "RemediationPolicyPatch"
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["remediation_policy_patch"] = vapiBindings_.NewReferenceType(model.RemediationPolicyPatchBindingType)
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["cluster"] = vapiBindings_.NewStringType()
	pathParams["cluster"] = "cluster"
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
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
		"remediation_policy_patch",
		"PATCH",
		"/vmc/autoscaler/api/orgs/{org}/sddcs/{sddc}/clusters/{cluster}/remediation-policy",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
