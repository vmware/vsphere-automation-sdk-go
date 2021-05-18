// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationMonitoringLogsCoreDumps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

func systemAdministrationMonitoringLogsCoreDumpsCopyCoreDumpToRemoteFileCopyToRemoteFileInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["file_name"] = bindings.NewStringType()
	fields["copy_to_remote_file_properties"] = bindings.NewReferenceType(model.CopyToRemoteFilePropertiesBindingType)
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["copy_to_remote_file_properties"] = "CopyToRemoteFileProperties"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringLogsCoreDumpsCopyCoreDumpToRemoteFileCopyToRemoteFileOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationMonitoringLogsCoreDumpsCopyCoreDumpToRemoteFileCopyToRemoteFileRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["file_name"] = bindings.NewStringType()
	fields["copy_to_remote_file_properties"] = bindings.NewReferenceType(model.CopyToRemoteFilePropertiesBindingType)
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["copy_to_remote_file_properties"] = "CopyToRemoteFileProperties"
	paramsTypeMap["copy_to_remote_file_properties"] = bindings.NewReferenceType(model.CopyToRemoteFilePropertiesBindingType)
	paramsTypeMap["file_name"] = bindings.NewStringType()
	paramsTypeMap["fileName"] = bindings.NewStringType()
	pathParams["file_name"] = "fileName"
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
		"action=copy_to_remote_file",
		"copy_to_remote_file_properties",
		"POST",
		"/api/v1/node/core-dumps/{fileName}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.timed_out": 500, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationMonitoringLogsCoreDumpsListCoreDumpsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationMonitoringLogsCoreDumpsListCoreDumpsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.FilePropertiesListResultBindingType)
}

func systemAdministrationMonitoringLogsCoreDumpsListCoreDumpsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/api/v1/node/core-dumps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
