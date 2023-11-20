// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Bundles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package repository

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

func bundlesCanceluploadInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bundle_id"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BundlesCanceluploadOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func bundlesCanceluploadRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["bundle_id"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["bundle_id"] = "BundleId"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["product"] = vapiBindings_.NewStringType()
	paramsTypeMap["bundle_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["bundleId"] = vapiBindings_.NewStringType()
	pathParams["bundle_id"] = "bundleId"
	queryParams["product"] = "product"
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

func bundlesCreateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["remote_bundle_url"] = vapiBindings_.NewReferenceType(nsxModel.RemoteBundleUrlBindingType)
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["remote_bundle_url"] = "RemoteBundleUrl"
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BundlesCreateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.BundleIdBindingType)
}

func bundlesCreateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["remote_bundle_url"] = vapiBindings_.NewReferenceType(nsxModel.RemoteBundleUrlBindingType)
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["remote_bundle_url"] = "RemoteBundleUrl"
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["product"] = vapiBindings_.NewStringType()
	paramsTypeMap["file_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["remote_bundle_url"] = vapiBindings_.NewReferenceType(nsxModel.RemoteBundleUrlBindingType)
	queryParams["product"] = "product"
	queryParams["file_type"] = "file_type"
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

func bundlesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func BundlesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.BundleIdsBindingType)
}

func bundlesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["product"] = vapiBindings_.NewStringType()
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["product"] = "Product"
	paramsTypeMap["product"] = vapiBindings_.NewStringType()
	paramsTypeMap["file_type"] = vapiBindings_.NewStringType()
	queryParams["product"] = "product"
	queryParams["file_type"] = "file_type"
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
		"/api/v1/repository/bundles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
