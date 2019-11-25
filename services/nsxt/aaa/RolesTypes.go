/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Roles.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package aaa

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func rolesCloneInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role_id"] = bindings.NewStringType()
	fields["new_role"] = bindings.NewReferenceType(model.NewRoleBindingType)
	fieldNameMap["role_id"] = "RoleId"
	fieldNameMap["new_role"] = "NewRole"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rolesCloneOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NewRoleBindingType)
}

func rolesCloneRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["role_id"] = bindings.NewStringType()
	fields["new_role"] = bindings.NewReferenceType(model.NewRoleBindingType)
	fieldNameMap["role_id"] = "RoleId"
	fieldNameMap["new_role"] = "NewRole"
	paramsTypeMap["role_id"] = bindings.NewStringType()
	paramsTypeMap["new_role"] = bindings.NewReferenceType(model.NewRoleBindingType)
	paramsTypeMap["roleId"] = bindings.NewStringType()
	pathParams["role_id"] = "roleId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=clone",
		"new_role",
		"POST",
		"/policy/api/v1/aaa/roles/{roleId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func rolesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role_id"] = bindings.NewStringType()
	fieldNameMap["role_id"] = "RoleId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rolesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func rolesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["role_id"] = bindings.NewStringType()
	fieldNameMap["role_id"] = "RoleId"
	paramsTypeMap["role_id"] = bindings.NewStringType()
	paramsTypeMap["roleId"] = bindings.NewStringType()
	pathParams["role_id"] = "roleId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/policy/api/v1/aaa/roles/{roleId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func rolesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role"] = bindings.NewStringType()
	fieldNameMap["role"] = "Role"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rolesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleWithFeaturesBindingType)
}

func rolesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["role"] = bindings.NewStringType()
	fieldNameMap["role"] = "Role"
	paramsTypeMap["role"] = bindings.NewStringType()
	paramsTypeMap["role"] = bindings.NewStringType()
	pathParams["role"] = "role"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/aaa/roles/{role}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func rolesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rolesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleListResultBindingType)
}

func rolesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/aaa/roles",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func rolesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role_id"] = bindings.NewStringType()
	fields["role_with_features"] = bindings.NewReferenceType(model.RoleWithFeaturesBindingType)
	fieldNameMap["role_id"] = "RoleId"
	fieldNameMap["role_with_features"] = "RoleWithFeatures"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rolesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleWithFeaturesBindingType)
}

func rolesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["role_id"] = bindings.NewStringType()
	fields["role_with_features"] = bindings.NewReferenceType(model.RoleWithFeaturesBindingType)
	fieldNameMap["role_id"] = "RoleId"
	fieldNameMap["role_with_features"] = "RoleWithFeatures"
	paramsTypeMap["role_id"] = bindings.NewStringType()
	paramsTypeMap["role_with_features"] = bindings.NewReferenceType(model.RoleWithFeaturesBindingType)
	paramsTypeMap["roleId"] = bindings.NewStringType()
	pathParams["role_id"] = "roleId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"role_with_features",
		"PUT",
		"/policy/api/v1/aaa/roles/{roleId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func rolesValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["feature_permission_array"] = bindings.NewReferenceType(model.FeaturePermissionArrayBindingType)
	fieldNameMap["feature_permission_array"] = "FeaturePermissionArray"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rolesValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RecommendedFeaturePermissionListResultBindingType)
}

func rolesValidateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["feature_permission_array"] = bindings.NewReferenceType(model.FeaturePermissionArrayBindingType)
	fieldNameMap["feature_permission_array"] = "FeaturePermissionArray"
	paramsTypeMap["feature_permission_array"] = bindings.NewReferenceType(model.FeaturePermissionArrayBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"action=validate",
		"feature_permission_array",
		"POST",
		"/policy/api/v1/aaa/roles",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


