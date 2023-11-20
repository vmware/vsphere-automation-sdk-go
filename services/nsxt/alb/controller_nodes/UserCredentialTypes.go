// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: UserCredential.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package controller_nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func userCredentialCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["a_LB_controller_user_credential"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialBindingType)
	fieldNameMap["a_LB_controller_user_credential"] = "ALBControllerUserCredential"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UserCredentialCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialResponseBindingType)
}

func userCredentialCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["a_LB_controller_user_credential"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialBindingType)
	fieldNameMap["a_LB_controller_user_credential"] = "ALBControllerUserCredential"
	paramsTypeMap["a_LB_controller_user_credential"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialBindingType)
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
		"a_LB_controller_user_credential",
		"POST",
		"/policy/api/v1/alb/controller-nodes/user-credential",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func userCredentialUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["a_LB_controller_user_credential"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialBindingType)
	fields["running_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["a_LB_controller_user_credential"] = "ALBControllerUserCredential"
	fieldNameMap["running_config"] = "RunningConfig"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UserCredentialUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialResponseBindingType)
}

func userCredentialUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["a_LB_controller_user_credential"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialBindingType)
	fields["running_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["a_LB_controller_user_credential"] = "ALBControllerUserCredential"
	fieldNameMap["running_config"] = "RunningConfig"
	paramsTypeMap["a_LB_controller_user_credential"] = vapiBindings_.NewReferenceType(nsx_policyModel.ALBControllerUserCredentialBindingType)
	paramsTypeMap["running_config"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["running_config"] = "running_config"
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
		"a_LB_controller_user_credential",
		"PUT",
		"/policy/api/v1/alb/controller-nodes/user-credential",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
