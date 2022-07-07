// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: GroupedFeedbackRequests.
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

// Possible value for ``networkLayer`` of method GroupedFeedbackRequests#list.
const GroupedFeedbackRequests_LIST_NETWORK_LAYER_L2 = "L2"

// Possible value for ``networkLayer`` of method GroupedFeedbackRequests#list.
const GroupedFeedbackRequests_LIST_NETWORK_LAYER_L3_L7 = "L3_L7"

// Possible value for ``state`` of method GroupedFeedbackRequests#list.
const GroupedFeedbackRequests_LIST_STATE_ALL = "ALL"

// Possible value for ``state`` of method GroupedFeedbackRequests#list.
const GroupedFeedbackRequests_LIST_STATE_RESOLVED = "RESOLVED"

// Possible value for ``state`` of method GroupedFeedbackRequests#list.
const GroupedFeedbackRequests_LIST_STATE_UNRESOLVED = "UNRESOLVED"

func groupedFeedbackRequestsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hash"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sub_category"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["category"] = "Category"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	fieldNameMap["hash"] = "Hash"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_layer"] = "NetworkLayer"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["sub_category"] = "SubCategory"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func groupedFeedbackRequestsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.GroupedMigrationFeedbackRequestListResultBindingType)
}

func groupedFeedbackRequestsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["category"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["hash"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sub_category"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["category"] = "Category"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	fieldNameMap["hash"] = "Hash"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["network_layer"] = "NetworkLayer"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["state"] = "State"
	fieldNameMap["sub_category"] = "SubCategory"
	paramsTypeMap["category"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["hash"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["network_layer"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["sub_category"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["state"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["network_layer"] = "network_layer"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sub_category"] = "sub_category"
	queryParams["federation_site_id"] = "federation_site_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["state"] = "state"
	queryParams["category"] = "category"
	queryParams["hash"] = "hash"
	queryParams["page_size"] = "page_size"
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
		"",
		"GET",
		"/api/v1/migration/grouped-feedback-requests",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
