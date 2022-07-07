// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Associations.
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

// Possible value for ``associatedResourceType`` of method Associations#list.
const Associations_LIST_ASSOCIATED_RESOURCE_TYPE_NSGROUP = "NSGroup"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_NSGROUP = "NSGroup"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_IPSET = "IPSet"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_MACSET = "MACSet"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_LOGICALSWITCH = "LogicalSwitch"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_LOGICALPORT = "LogicalPort"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_VIRTUALMACHINE = "VirtualMachine"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_DIRECTORYGROUP = "DirectoryGroup"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_VIRTUALNETWORKINTERFACE = "VirtualNetworkInterface"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_TRANSPORTNODE = "TransportNode"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_IPADDRESS = "IPAddress"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_CLOUDNATIVESERVICEINSTANCE = "CloudNativeServiceInstance"

// Possible value for ``resourceType`` of method Associations#list.
const Associations_LIST_RESOURCE_TYPE_PHYSICALSERVER = "PhysicalServer"

func associationsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["associated_resource_type"] = bindings.NewStringType()
	fields["resource_id"] = bindings.NewStringType()
	fields["resource_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["fetch_ancestors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["associated_resource_type"] = "AssociatedResourceType"
	fieldNameMap["resource_id"] = "ResourceId"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["fetch_ancestors"] = "FetchAncestors"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associationsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AssociationListResultBindingType)
}

func associationsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["associated_resource_type"] = bindings.NewStringType()
	fields["resource_id"] = bindings.NewStringType()
	fields["resource_type"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["fetch_ancestors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["associated_resource_type"] = "AssociatedResourceType"
	fieldNameMap["resource_id"] = "ResourceId"
	fieldNameMap["resource_type"] = "ResourceType"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["fetch_ancestors"] = "FetchAncestors"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["resource_type"] = bindings.NewStringType()
	paramsTypeMap["resource_id"] = bindings.NewStringType()
	paramsTypeMap["fetch_ancestors"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["associated_resource_type"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["resource_type"] = "resource_type"
	queryParams["resource_id"] = "resource_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["associated_resource_type"] = "associated_resource_type"
	queryParams["fetch_ancestors"] = "fetch_ancestors"
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
		"/api/v1/associations",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
