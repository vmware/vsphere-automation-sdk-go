// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TransportNodeProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func transportNodeProfilesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_profile"] = bindings.NewReferenceType(model.TransportNodeProfileBindingType)
	fieldNameMap["transport_node_profile"] = "TransportNodeProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeProfilesCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeProfileBindingType)
}

func transportNodeProfilesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_profile"] = bindings.NewReferenceType(model.TransportNodeProfileBindingType)
	fieldNameMap["transport_node_profile"] = "TransportNodeProfile"
	paramsTypeMap["transport_node_profile"] = bindings.NewReferenceType(model.TransportNodeProfileBindingType)
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
		"transport_node_profile",
		"POST",
		"/api/v1/transport-node-profiles",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_profile_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_profile_id"] = "TransportNodeProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodeProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_profile_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_profile_id"] = "TransportNodeProfileId"
	paramsTypeMap["transport_node_profile_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeProfileId"] = bindings.NewStringType()
	pathParams["transport_node_profile_id"] = "transportNodeProfileId"
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
		"/api/v1/transport-node-profiles/{transportNodeProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_profile_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_profile_id"] = "TransportNodeProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeProfileBindingType)
}

func transportNodeProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_profile_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_profile_id"] = "TransportNodeProfileId"
	paramsTypeMap["transport_node_profile_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeProfileId"] = bindings.NewStringType()
	pathParams["transport_node_profile_id"] = "transportNodeProfileId"
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
		"/api/v1/transport-node-profiles/{transportNodeProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeProfileListResultBindingType)
}

func transportNodeProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/transport-node-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_profile_id"] = bindings.NewStringType()
	fields["transport_node_profile"] = bindings.NewReferenceType(model.TransportNodeProfileBindingType)
	fields["esx_mgmt_if_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["if_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ping_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["skip_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vnic_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_profile_id"] = "TransportNodeProfileId"
	fieldNameMap["transport_node_profile"] = "TransportNodeProfile"
	fieldNameMap["esx_mgmt_if_migration_dest"] = "EsxMgmtIfMigrationDest"
	fieldNameMap["if_id"] = "IfId"
	fieldNameMap["ping_ip"] = "PingIp"
	fieldNameMap["skip_validation"] = "SkipValidation"
	fieldNameMap["vnic"] = "Vnic"
	fieldNameMap["vnic_migration_dest"] = "VnicMigrationDest"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeProfileBindingType)
}

func transportNodeProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_profile_id"] = bindings.NewStringType()
	fields["transport_node_profile"] = bindings.NewReferenceType(model.TransportNodeProfileBindingType)
	fields["esx_mgmt_if_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["if_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ping_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["skip_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vnic_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["transport_node_profile_id"] = "TransportNodeProfileId"
	fieldNameMap["transport_node_profile"] = "TransportNodeProfile"
	fieldNameMap["esx_mgmt_if_migration_dest"] = "EsxMgmtIfMigrationDest"
	fieldNameMap["if_id"] = "IfId"
	fieldNameMap["ping_ip"] = "PingIp"
	fieldNameMap["skip_validation"] = "SkipValidation"
	fieldNameMap["vnic"] = "Vnic"
	fieldNameMap["vnic_migration_dest"] = "VnicMigrationDest"
	paramsTypeMap["transport_node_profile"] = bindings.NewReferenceType(model.TransportNodeProfileBindingType)
	paramsTypeMap["vnic"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["skip_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["if_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["esx_mgmt_if_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ping_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vnic_migration_dest"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["transport_node_profile_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeProfileId"] = bindings.NewStringType()
	pathParams["transport_node_profile_id"] = "transportNodeProfileId"
	queryParams["ping_ip"] = "ping_ip"
	queryParams["vnic"] = "vnic"
	queryParams["skip_validation"] = "skip_validation"
	queryParams["esx_mgmt_if_migration_dest"] = "esx_mgmt_if_migration_dest"
	queryParams["if_id"] = "if_id"
	queryParams["vnic_migration_dest"] = "vnic_migration_dest"
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
		"transport_node_profile",
		"PUT",
		"/api/v1/transport-node-profiles/{transportNodeProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
