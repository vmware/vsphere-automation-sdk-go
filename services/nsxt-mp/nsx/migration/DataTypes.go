// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Data.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package migration

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
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

// Possible value for ``fileType`` of method Data#get.
const Data_GET_FILE_TYPE_MIGRATION_REPORT = "MIGRATION_REPORT"

func dataDeleteInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DataDeleteOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewVoidType()
}

func dataDeleteRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	paramsTypeMap["file_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["file_type"] = "file_type"
	queryParams["federation_site_id"] = "federation_site_id"
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
		"DELETE",
		"/api/v1/migration/data",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func dataGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DataGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MigrationDataInfoBindingType)
}

func dataGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_type"] = vapiBindings_.NewStringType()
	fields["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["file_type"] = "FileType"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	paramsTypeMap["file_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["federation_site_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["file_type"] = "file_type"
	queryParams["federation_site_id"] = "federation_site_id"
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
		"/api/v1/migration/data",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
