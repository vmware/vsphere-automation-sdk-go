/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: EffectivePermissions.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package aaa

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func effectivePermissionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["feature_name"] = bindings.NewStringType()
	fields["object_path"] = bindings.NewStringType()
	fieldNameMap["feature_name"] = "FeatureName"
	fieldNameMap["object_path"] = "ObjectPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func effectivePermissionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PathPermissionGroupBindingType)
}

func effectivePermissionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["feature_name"] = bindings.NewStringType()
	fields["object_path"] = bindings.NewStringType()
	fieldNameMap["feature_name"] = "FeatureName"
	fieldNameMap["object_path"] = "ObjectPath"
	paramsTypeMap["object_path"] = bindings.NewStringType()
	paramsTypeMap["feature_name"] = bindings.NewStringType()
	queryParams["feature_name"] = "feature_name"
	queryParams["object_path"] = "object_path"
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
		"/policy/api/v1/aaa/effective-permissions",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


