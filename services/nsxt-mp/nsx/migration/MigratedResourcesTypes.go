// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: MigratedResources.
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

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_IPBLOCK = "IPBLOCK"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_IPPOOL = "IPPOOL"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_IPBLOCK_SUBNET = "IPBLOCK_SUBNET"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_IPPOOL_ALLOCATION = "IPPOOL_ALLOCATION"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_TIER0 = "TIER0"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_TIER1 = "TIER1"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_TIER0_LOGICAL_ROUTER_PORT = "TIER0_LOGICAL_ROUTER_PORT"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_TIER1_LOGICAL_ROUTER_PORT = "TIER1_LOGICAL_ROUTER_PORT"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_TIER0_LOGICAL_ROUTER_CONFIG = "TIER0_LOGICAL_ROUTER_CONFIG"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_SPOOFGUARD_PROFILES = "SPOOFGUARD_PROFILES"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LOGICAL_SWITCH = "LOGICAL_SWITCH"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LOGICAL_PORT = "LOGICAL_PORT"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_NAT = "NAT"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_IP_SET = "IP_SET"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_NS_GROUP = "NS_GROUP"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_CERTIFICATE = "CERTIFICATE"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_CRL = "CRL"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_MONITOR = "LB_MONITOR"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_POOL = "LB_POOL"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_PERSISTENCE_PROFILE = "LB_PERSISTENCE_PROFILE"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_APPLICATION_PROFILE = "LB_APPLICATION_PROFILE"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_SERVICE = "LB_SERVICE"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_CLIENT_SSL_PROFILE = "LB_CLIENT_SSL_PROFILE"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_LB_VIRTUAL_SERVER = "LB_VIRTUAL_SERVER"

// Possible value for ``resourceType`` of method MigratedResources#list.
const MigratedResources_LIST_RESOURCE_TYPE_DFW_SECTION = "DFW_SECTION"

func migratedResourcesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_type"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_id"] = "ResourceId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func MigratedResourcesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.MigratedObjectListResultBindingType)
}

func migratedResourcesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["resource_type"] = vapiBindings_.NewStringType()
	fields["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fields["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	fields["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["resource_id"] = "ResourceId"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["cursor"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_ascending"] = vapiBindings_.NewOptionalType(vapiBindings_.NewBooleanType())
	paramsTypeMap["included_fields"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["resource_type"] = vapiBindings_.NewStringType()
	paramsTypeMap["resource_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["sort_by"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["page_size"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["resource_type"] = "resource_type"
	queryParams["resource_id"] = "resource_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
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
		"/api/v1/migration/migrated-resources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
