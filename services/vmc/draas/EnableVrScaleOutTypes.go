// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EnableVrScaleOut.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package draas

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmcDraasModel "github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
	"reflect"
)

// Possible value for ``dryRun`` of method EnableVrScaleOut#post.
const EnableVrScaleOut_POST_DRY_RUN_TRUE = "true"

// Possible value for ``dryRun`` of method EnableVrScaleOut#post.
const EnableVrScaleOut_POST_DRY_RUN_FALSE = "false"

// Possible value for ``dryRun`` of method EnableVrScaleOut#post.
const EnableVrScaleOut_POST_DRY_RUN_SKIP = "skip"

func enableVrScaleOutPostInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["scale_out"] = vapiBindings_.NewBooleanType()
	fields["dry_run"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["is_retry"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["scale_out"] = "ScaleOut"
	fieldNameMap["dry_run"] = "DryRun"
	fieldNameMap["is_retry"] = "IsRetry"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnableVrScaleOutPostOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcDraasModel.TaskBindingType)
}

func enableVrScaleOutPostRestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["scale_out"] = vapiBindings_.NewBooleanType()
	fields["dry_run"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["is_retry"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["scale_out"] = "ScaleOut"
	fieldNameMap["dry_run"] = "DryRun"
	fieldNameMap["is_retry"] = "IsRetry"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["is_retry"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["dry_run"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["scale_out"] = vapiBindings_.NewBooleanType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["is_retry"] = "is_retry"
	queryParams["dry_run"] = "dry_run"
	queryParams["scale_out"] = "scale_out"
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
		"/vmc/draas/api/orgs/{org}/sddcs/{sddc}/site-recovery/vr-scale-out",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func enableVrScaleOutPost0InputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = vapiBindings_.NewStringType()
	fields["sddc"] = vapiBindings_.NewStringType()
	fields["scale_out"] = vapiBindings_.NewBooleanType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["scale_out"] = "ScaleOut"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func EnableVrScaleOutPost0OutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmcDraasModel.TaskBindingType)
}

func enableVrScaleOutPost0RestMetadata() vapiProtocol_.OperationRestMetadata {
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
	fields["scale_out"] = vapiBindings_.NewBooleanType()
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["scale_out"] = "ScaleOut"
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	paramsTypeMap["scale_out"] = vapiBindings_.NewBooleanType()
	paramsTypeMap["org"] = vapiBindings_.NewStringType()
	paramsTypeMap["sddc"] = vapiBindings_.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["scale_out"] = "scale_out"
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
		"PUT",
		"/vmc/draas/api/orgs/{org}/sddcs/{sddc}/site-recovery/vr-scale-out-config",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403})
}
