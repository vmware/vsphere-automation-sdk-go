// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: DeploymentGroups.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package core

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	vmwarecloudVmc_core_inventoryModel "github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
	"reflect"
)

func deploymentGroupsGetCoreDeploymentGroupsInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["include_deleted_resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["include_deleted_resources"] = "IncludeDeletedResources"
	fieldNameMap["page"] = "Page"
	fieldNameMap["size"] = "Size"
	fieldNameMap["sort"] = "Sort"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentGroupsGetCoreDeploymentGroupsOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_inventoryModel.DeploymentGroupResponseBindingType)
}

func deploymentGroupsGetCoreDeploymentGroupsRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["include_deleted_resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["page"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["sort"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["include_deleted_resources"] = "IncludeDeletedResources"
	fieldNameMap["page"] = "Page"
	fieldNameMap["size"] = "Size"
	fieldNameMap["sort"] = "Sort"
	paramsTypeMap["size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["include_deleted_resources"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["page"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	paramsTypeMap["sort"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	pathParams["org_id"] = "orgId"
	queryParams["size"] = "size"
	queryParams["include_deleted_resources"] = "include_deleted_resources"
	queryParams["page"] = "page"
	queryParams["sort"] = "sort"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/api/inventory/{orgId}/core/deployment-groups",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403})
}

func deploymentGroupsGetCoreDeploymentGroupByIdInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["deployment_group_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func DeploymentGroupsGetCoreDeploymentGroupByIdOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(vmwarecloudVmc_core_inventoryModel.DeploymentGroupBindingType)
}

func deploymentGroupsGetCoreDeploymentGroupByIdRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = vapiBindings_.NewStringType()
	fields["deployment_group_id"] = vapiBindings_.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["deployment_group_id"] = "DeploymentGroupId"
	paramsTypeMap["org_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["deployment_group_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["orgId"] = vapiBindings_.NewStringType()
	paramsTypeMap["deploymentGroupId"] = vapiBindings_.NewStringType()
	pathParams["deployment_group_id"] = "deploymentGroupId"
	pathParams["org_id"] = "orgId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"] = make(map[string]string)
	errorHeaders["com.vmware.vapi.std.errors.unauthenticated"]["challenge"] = "WWW-Authenticate"
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
		"/api/inventory/{orgId}/core/deployment-groups/{deploymentGroupId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.unauthenticated": 401, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.not_found": 404})
}
