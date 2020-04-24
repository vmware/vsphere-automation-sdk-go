/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Edges.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networks

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func edgesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_type"] = bindings.NewStringType()
	fields["prev_edge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ld_rname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_type"] = "EdgeType"
	fieldNameMap["prev_edge_id"] = "PrevEdgeId"
	fieldNameMap["start_index"] = "StartIndex"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_order_ascending"] = "SortOrderAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["filter"] = "Filter"
	fieldNameMap["ld_rname"] = "LdRname"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PagedEdgeListBindingType)
}

func edgesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc"] = bindings.NewStringType()
	fields["edge_type"] = bindings.NewStringType()
	fields["prev_edge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ld_rname"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc"] = "Sddc"
	fieldNameMap["edge_type"] = "EdgeType"
	fieldNameMap["prev_edge_id"] = "PrevEdgeId"
	fieldNameMap["start_index"] = "StartIndex"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_order_ascending"] = "SortOrderAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["filter"] = "Filter"
	fieldNameMap["ld_rname"] = "LdRname"
	paramsTypeMap["edge_type"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	paramsTypeMap["filter"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["prev_edge_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ld_rname"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc"] = bindings.NewStringType()
	pathParams["org"] = "org"
	pathParams["sddc"] = "sddc"
	queryParams["filter"] = "filter"
	queryParams["start_index"] = "startIndex"
	queryParams["edge_type"] = "edgeType"
	queryParams["page_size"] = "pageSize"
	queryParams["sort_by"] = "sortBy"
	queryParams["prev_edge_id"] = "prevEdgeId"
	queryParams["sort_order_ascending"] = "sortOrderAscending"
	queryParams["ld_rname"] = "LDRname"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		"/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.not_found": 404})
}


