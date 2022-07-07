// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Bundles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package repository

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``product`` of method Bundles#cancelupload.
const Bundles_CANCELUPLOAD_PRODUCT_SAMPLE = "SAMPLE"

// Possible value for ``product`` of method Bundles#cancelupload.
const Bundles_CANCELUPLOAD_PRODUCT_ALB_CONTROLLER = "ALB_CONTROLLER"

// Possible value for ``product`` of method Bundles#cancelupload.
const Bundles_CANCELUPLOAD_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``fileType`` of method Bundles#create.
const Bundles_CREATE_FILE_TYPE_OVA = "OVA"

// Possible value for ``product`` of method Bundles#create.
const Bundles_CREATE_PRODUCT_SAMPLE = "SAMPLE"

// Possible value for ``product`` of method Bundles#create.
const Bundles_CREATE_PRODUCT_ALB_CONTROLLER = "ALB_CONTROLLER"

// Possible value for ``product`` of method Bundles#create.
const Bundles_CREATE_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

// Possible value for ``fileType`` of method Bundles#get.
const Bundles_GET_FILE_TYPE_OVA = "OVA"

// Possible value for ``product`` of method Bundles#get.
const Bundles_GET_PRODUCT_SAMPLE = "SAMPLE"

// Possible value for ``product`` of method Bundles#get.
const Bundles_GET_PRODUCT_ALB_CONTROLLER = "ALB_CONTROLLER"

// Possible value for ``product`` of method Bundles#get.
const Bundles_GET_PRODUCT_INTELLIGENCE = "INTELLIGENCE"

func bundlesCanceluploadInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bundlesCanceluploadOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func bundlesCanceluploadRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["bundle_id"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	paramsTypeMap["bundleId"] = bindings.NewStringType()
	pathParams["bundle_id"] = "bundleId"
	queryParams["product"] = "product"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=cancel_upload",
		"",
		"POST",
		"/api/v1/repository/bundles/{bundleId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bundlesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["remote_bundle_url"] = bindings.NewReferenceType(model.RemoteBundleUrlBindingType)
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["remote_bundle_url"] = "RemoteBundleUrl"
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bundlesCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BundleIdBindingType)
}

func bundlesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["remote_bundle_url"] = bindings.NewReferenceType(model.RemoteBundleUrlBindingType)
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["remote_bundle_url"] = "RemoteBundleUrl"
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["remote_bundle_url"] = bindings.NewReferenceType(model.RemoteBundleUrlBindingType)
	paramsTypeMap["file_type"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	queryParams["product"] = "product"
	queryParams["file_type"] = "file_type"
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
		"remote_bundle_url",
		"POST",
		"/api/v1/repository/bundles",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func bundlesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func bundlesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BundleIdsBindingType)
}

func bundlesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewStringType()
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["file_type"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewStringType()
	queryParams["product"] = "product"
	queryParams["file_type"] = "file_type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
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
		"/api/v1/repository/bundles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
