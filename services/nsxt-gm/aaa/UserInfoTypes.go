// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: UserInfo.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package aaa

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsx_global_policyModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"reflect"
)

func userInfoGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provide_flat_listing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provide_flat_listing"] = "ProvideFlatListing"
	fieldNameMap["root_path"] = "RootPath"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func UserInfoGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsx_global_policyModel.UserInfoBindingType)
}

func userInfoGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["provide_flat_listing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provide_flat_listing"] = "ProvideFlatListing"
	fieldNameMap["root_path"] = "RootPath"
	paramsTypeMap["root_path"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["provide_flat_listing"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	queryParams["root_path"] = "root_path"
	queryParams["provide_flat_listing"] = "provide_flat_listing"
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
		"/global-manager/api/v1/aaa/user-info",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
