// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SystemAdministrationConfigurationFabricComputeCollections.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``action`` of method SystemAdministrationConfigurationFabricComputeCollections#performActionOnComputeCollection.
const SystemAdministrationConfigurationFabricComputeCollections_PERFORM_ACTION_ON_COMPUTE_COLLECTION_ACTION_NSX = "remove_nsx"

func systemAdministrationConfigurationFabricComputeCollectionsCreateComputeCollectionFabricTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compute_collection_fabric_template"] = bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
	fieldNameMap["compute_collection_fabric_template"] = "ComputeCollectionFabricTemplate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsCreateComputeCollectionFabricTemplateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsCreateComputeCollectionFabricTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["compute_collection_fabric_template"] = bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
	fieldNameMap["compute_collection_fabric_template"] = "ComputeCollectionFabricTemplate"
	paramsTypeMap["compute_collection_fabric_template"] = bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
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
		"compute_collection_fabric_template",
		"POST",
		"/api/v1/fabric/compute-collection-fabric-templates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsDeleteComputeCollectionFabricTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fabric_template_id"] = bindings.NewStringType()
	fieldNameMap["fabric_template_id"] = "FabricTemplateId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsDeleteComputeCollectionFabricTemplateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricComputeCollectionsDeleteComputeCollectionFabricTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["fabric_template_id"] = bindings.NewStringType()
	fieldNameMap["fabric_template_id"] = "FabricTemplateId"
	paramsTypeMap["fabric_template_id"] = bindings.NewStringType()
	paramsTypeMap["fabricTemplateId"] = bindings.NewStringType()
	pathParams["fabric_template_id"] = "fabricTemplateId"
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
		"/api/v1/fabric/compute-collection-fabric-templates/{fabricTemplateId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsGetComputeCollectionFabricTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fabric_template_id"] = bindings.NewStringType()
	fieldNameMap["fabric_template_id"] = "FabricTemplateId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsGetComputeCollectionFabricTemplateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsGetComputeCollectionFabricTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["fabric_template_id"] = bindings.NewStringType()
	fieldNameMap["fabric_template_id"] = "FabricTemplateId"
	paramsTypeMap["fabric_template_id"] = bindings.NewStringType()
	paramsTypeMap["fabricTemplateId"] = bindings.NewStringType()
	pathParams["fabric_template_id"] = "fabricTemplateId"
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
		"/api/v1/fabric/compute-collection-fabric-templates/{fabricTemplateId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsGetHostNodeStatusOnComputeCollectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsGetHostNodeStatusOnComputeCollectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.HostNodeStatusListResultBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsGetHostNodeStatusOnComputeCollectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
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
		"/api/v1/fabric/compute-collections/{ccExtId}/member-status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionFabricTemplatesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compute_collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["compute_collection_id"] = "ComputeCollectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionFabricTemplatesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionFabricTemplateListResultBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionFabricTemplatesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["compute_collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["compute_collection_id"] = "ComputeCollectionId"
	paramsTypeMap["compute_collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["compute_collection_id"] = "compute_collection_id"
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
		"/api/v1/fabric/compute-collection-fabric-templates",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionPhysicalNetworkInterfacesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionPhysicalNetworkInterfacesOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionNetworkInterfacesListResultBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionPhysicalNetworkInterfacesRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
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
		"/api/v1/fabric/compute-collections/{ccExtId}/network/physical-interfaces",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cm_local_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["discovered_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["origin_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["origin_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["owner_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cm_local_id"] = "CmLocalId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["discovered_node_id"] = "DiscoveredNodeId"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["origin_id"] = "OriginId"
	fieldNameMap["origin_type"] = "OriginType"
	fieldNameMap["owner_id"] = "OwnerId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionListResultBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsListComputeCollectionsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cm_local_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["discovered_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["origin_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["origin_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["owner_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cm_local_id"] = "CmLocalId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["discovered_node_id"] = "DiscoveredNodeId"
	fieldNameMap["display_name"] = "DisplayName"
	fieldNameMap["external_id"] = "ExternalId"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["origin_id"] = "OriginId"
	fieldNameMap["origin_type"] = "OriginType"
	fieldNameMap["owner_id"] = "OwnerId"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["origin_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["origin_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["display_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["owner_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cm_local_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["external_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["discovered_node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["node_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["cm_local_id"] = "cm_local_id"
	queryParams["owner_id"] = "owner_id"
	queryParams["external_id"] = "external_id"
	queryParams["origin_id"] = "origin_id"
	queryParams["sort_by"] = "sort_by"
	queryParams["display_name"] = "display_name"
	queryParams["origin_type"] = "origin_type"
	queryParams["discovered_node_id"] = "discovered_node_id"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["node_id"] = "node_id"
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
		"/api/v1/fabric/compute-collections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsPerformActionOnComputeCollectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cc_ext_id"] = "CcExtId"
	fieldNameMap["action"] = "Action"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsPerformActionOnComputeCollectionOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func systemAdministrationConfigurationFabricComputeCollectionsPerformActionOnComputeCollectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fields["action"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cc_ext_id"] = "CcExtId"
	fieldNameMap["action"] = "Action"
	paramsTypeMap["action"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
	queryParams["action"] = "action"
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
		"POST",
		"/api/v1/fabric/compute-collections/{ccExtId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsReadComputeCollectionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsReadComputeCollectionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsReadComputeCollectionRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cc_ext_id"] = bindings.NewStringType()
	fieldNameMap["cc_ext_id"] = "CcExtId"
	paramsTypeMap["cc_ext_id"] = bindings.NewStringType()
	paramsTypeMap["ccExtId"] = bindings.NewStringType()
	pathParams["cc_ext_id"] = "ccExtId"
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
		"/api/v1/fabric/compute-collections/{ccExtId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func systemAdministrationConfigurationFabricComputeCollectionsUpdateComputeCollectionFabricTemplateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["fabric_template_id"] = bindings.NewStringType()
	fields["compute_collection_fabric_template"] = bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
	fieldNameMap["fabric_template_id"] = "FabricTemplateId"
	fieldNameMap["compute_collection_fabric_template"] = "ComputeCollectionFabricTemplate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemAdministrationConfigurationFabricComputeCollectionsUpdateComputeCollectionFabricTemplateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
}

func systemAdministrationConfigurationFabricComputeCollectionsUpdateComputeCollectionFabricTemplateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["fabric_template_id"] = bindings.NewStringType()
	fields["compute_collection_fabric_template"] = bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
	fieldNameMap["fabric_template_id"] = "FabricTemplateId"
	fieldNameMap["compute_collection_fabric_template"] = "ComputeCollectionFabricTemplate"
	paramsTypeMap["compute_collection_fabric_template"] = bindings.NewReferenceType(model.ComputeCollectionFabricTemplateBindingType)
	paramsTypeMap["fabric_template_id"] = bindings.NewStringType()
	paramsTypeMap["fabricTemplateId"] = bindings.NewStringType()
	pathParams["fabric_template_id"] = "fabricTemplateId"
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
		"compute_collection_fabric_template",
		"PUT",
		"/api/v1/fabric/compute-collection-fabric-templates/{fabricTemplateId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
