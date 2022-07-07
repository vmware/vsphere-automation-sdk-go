// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Data.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``fileType`` of method Data#delete.
const Data_DELETE_FILE_TYPE_VRA_INPUT = "VRA_INPUT"

// Possible value for ``fileType`` of method Data#delete.
const Data_DELETE_FILE_TYPE_EDGE_CUTOVER_MAPPING = "EDGE_CUTOVER_MAPPING"

// Possible value for ``fileType`` of method Data#delete.
const Data_DELETE_FILE_TYPE_BYOT_L3_MAPPING = "BYOT_L3_MAPPING"

// Possible value for ``fileType`` of method Data#delete.
const Data_DELETE_FILE_TYPE_AVI_LB_MAPPING = "AVI_LB_MAPPING"

// Possible value for ``fileType`` of method Data#get.
const Data_GET_FILE_TYPE_VRA_INPUT = "VRA_INPUT"

// Possible value for ``fileType`` of method Data#get.
const Data_GET_FILE_TYPE_VRA_OUTPUT = "VRA_OUTPUT"

// Possible value for ``fileType`` of method Data#get.
const Data_GET_FILE_TYPE_EDGE_CUTOVER_MAPPING = "EDGE_CUTOVER_MAPPING"

// Possible value for ``fileType`` of method Data#get.
const Data_GET_FILE_TYPE_BYOT_L3_MAPPING = "BYOT_L3_MAPPING"

// Possible value for ``fileType`` of method Data#get.
const Data_GET_FILE_TYPE_AVI_LB_MAPPING = "AVI_LB_MAPPING"

func dataDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = bindings.NewStringType()
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dataDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func dataDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = bindings.NewStringType()
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	paramsTypeMap["file_type"] = bindings.NewStringType()
	paramsTypeMap["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["file_type"] = "file_type"
	queryParams["federation_site_id"] = "federation_site_id"
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
		"DELETE",
		"/api/v1/migration/data",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dataGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = bindings.NewStringType()
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func dataGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.MigrationDataInfoBindingType)
}

func dataGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = bindings.NewStringType()
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	paramsTypeMap["file_type"] = bindings.NewStringType()
	paramsTypeMap["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["file_type"] = "file_type"
	queryParams["federation_site_id"] = "federation_site_id"
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
		"/api/v1/migration/data",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
