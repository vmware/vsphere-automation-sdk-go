// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Normalizations.
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

// Possible value for ``preferredNormalizationType`` of method Normalizations#list.
const Normalizations_LIST_PREFERRED_NORMALIZATION_TYPE_NSGROUP = "NSGroup"

// Possible value for ``preferredNormalizationType`` of method Normalizations#list.
const Normalizations_LIST_PREFERRED_NORMALIZATION_TYPE_IPSET = "IPSet"

// Possible value for ``preferredNormalizationType`` of method Normalizations#list.
const Normalizations_LIST_PREFERRED_NORMALIZATION_TYPE_MACSET = "MACSet"

// Possible value for ``preferredNormalizationType`` of method Normalizations#list.
const Normalizations_LIST_PREFERRED_NORMALIZATION_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``preferredNormalizationType`` of method Normalizations#list.
const Normalizations_LIST_PREFERRED_NORMALIZATION_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``preferredNormalizationType`` of method Normalizations#list.
const Normalizations_LIST_PREFERRED_NORMALIZATION_TYPE_DIRECTORYGROUP = "DirectoryGroup"

// Possible value for ``resourceType`` of method Normalizations#list.
const Normalizations_LIST_RESOURCE_TYPE_NSGROUP = "NSGroup"

func normalizationsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["preferred_normalization_type"] = bindings.NewStringType()
	fields["resource_id"] = bindings.NewStringType()
	fields["resource_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["preferred_normalization_type"] = "PreferredNormalizationType"
	fieldNameMap["resource_id"] = "ResourceId"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func normalizationsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NormalizedResourceListResultBindingType)
}

func normalizationsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["preferred_normalization_type"] = bindings.NewStringType()
	fields["resource_id"] = bindings.NewStringType()
	fields["resource_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["preferred_normalization_type"] = "PreferredNormalizationType"
	fieldNameMap["resource_id"] = "ResourceId"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["resource_type"] = bindings.NewStringType()
	paramsTypeMap["resource_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["preferred_normalization_type"] = bindings.NewStringType()
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["resource_type"] = "resource_type"
	queryParams["resource_id"] = "resource_id"
	queryParams["preferred_normalization_type"] = "preferred_normalization_type"
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
		"/api/v1/normalizations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
