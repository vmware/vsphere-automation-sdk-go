// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CommuntyLists.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package bgp

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func communtyListsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communtyListsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.BGPCommunityListBindingType)
}

func communtyListsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = bindings.NewStringType()
	fields["community_list_id"] = bindings.NewStringType()
	fields["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	paramsTypeMap["logical_router_id"] = bindings.NewStringType()
	paramsTypeMap["community_list_id"] = bindings.NewStringType()
	paramsTypeMap["b_GP_community_list"] = bindings.NewReferenceType(model.BGPCommunityListBindingType)
	paramsTypeMap["logicalRouterId"] = bindings.NewStringType()
	paramsTypeMap["communityListId"] = bindings.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	pathParams["community_list_id"] = "communityListId"
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
		"b_GP_community_list",
		"PUT",
		"/api/v1/logical-routers/{logicalRouterId}/routing/bgp/communty-lists/{communityListId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
