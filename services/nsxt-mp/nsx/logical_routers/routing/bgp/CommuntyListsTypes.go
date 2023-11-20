// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CommuntyLists.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package bgp

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func communtyListsUpdateInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_id"] = vapiBindings_.NewStringType()
	fields["community_list_id"] = vapiBindings_.NewStringType()
	fields["b_GP_community_list"] = vapiBindings_.NewReferenceType(nsxModel.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func CommuntyListsUpdateOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.BGPCommunityListBindingType)
}

func communtyListsUpdateRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_id"] = vapiBindings_.NewStringType()
	fields["community_list_id"] = vapiBindings_.NewStringType()
	fields["b_GP_community_list"] = vapiBindings_.NewReferenceType(nsxModel.BGPCommunityListBindingType)
	fieldNameMap["logical_router_id"] = "LogicalRouterId"
	fieldNameMap["community_list_id"] = "CommunityListId"
	fieldNameMap["b_GP_community_list"] = "BGPCommunityList"
	paramsTypeMap["logical_router_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["community_list_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["b_GP_community_list"] = vapiBindings_.NewReferenceType(nsxModel.BGPCommunityListBindingType)
	paramsTypeMap["logicalRouterId"] = vapiBindings_.NewStringType()
	paramsTypeMap["communityListId"] = vapiBindings_.NewStringType()
	pathParams["logical_router_id"] = "logicalRouterId"
	pathParams["community_list_id"] = "communityListId"
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
