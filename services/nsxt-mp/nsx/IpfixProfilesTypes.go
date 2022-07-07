// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: IpfixProfiles.
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

// Possible value for ``appliedToEntityType`` of method IpfixProfiles#list.
const IpfixProfiles_LIST_APPLIED_TO_ENTITY_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``appliedToEntityType`` of method IpfixProfiles#list.
const IpfixProfiles_LIST_APPLIED_TO_ENTITY_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``appliedToEntityType`` of method IpfixProfiles#list.
const IpfixProfiles_LIST_APPLIED_TO_ENTITY_TYPE_NSGROUP = "NSGroup"

func ipfixProfilesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixProfilesCreateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
}

func ipfixProfilesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	paramsTypeMap["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
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
		"ipfix_upm_profile",
		"POST",
		"/api/v1/ipfix-profiles",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipfixProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipfixProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	paramsTypeMap["ipfix_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixProfileId"] = bindings.NewStringType()
	pathParams["ipfix_profile_id"] = "ipfixProfileId"
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
		"/api/v1/ipfix-profiles/{ipfixProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipfixProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixProfilesGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
}

func ipfixProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	paramsTypeMap["ipfix_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixProfileId"] = bindings.NewStringType()
	pathParams["ipfix_profile_id"] = "ipfixProfileId"
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
		"/api/v1/ipfix-profiles/{ipfixProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipfixProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["applied_to_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["applied_to_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_to_entity_id"] = "AppliedToEntityId"
	fieldNameMap["applied_to_entity_type"] = "AppliedToEntityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["profile_types"] = "ProfileTypes"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IpfixUpmProfileListResultBindingType)
}

func ipfixProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["applied_to_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["applied_to_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["applied_to_entity_id"] = "AppliedToEntityId"
	fieldNameMap["applied_to_entity_type"] = "AppliedToEntityType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["profile_types"] = "ProfileTypes"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["profile_types"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_to_entity_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["applied_to_entity_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["applied_to_entity_id"] = "applied_to_entity_id"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["profile_types"] = "profile_types"
	queryParams["sort_by"] = "sort_by"
	queryParams["applied_to_entity_type"] = "applied_to_entity_type"
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
		"/api/v1/ipfix-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func ipfixProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
}

func ipfixProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["ipfix_profile_id"] = bindings.NewStringType()
	fields["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	fieldNameMap["ipfix_profile_id"] = "IpfixProfileId"
	fieldNameMap["ipfix_upm_profile"] = "IpfixUpmProfile"
	paramsTypeMap["ipfix_upm_profile"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IpfixUpmProfileBindingType)}, bindings.REST)
	paramsTypeMap["ipfix_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixProfileId"] = bindings.NewStringType()
	pathParams["ipfix_profile_id"] = "ipfixProfileId"
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
		"ipfix_upm_profile",
		"PUT",
		"/api/v1/ipfix-profiles/{ipfixProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
