// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: TransportNodeCollections.
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

func transportNodeCollectionsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_collection"] = bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
	fields["apply_profile"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_collection"] = "TransportNodeCollection"
	fieldNameMap["apply_profile"] = "ApplyProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeCollectionsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
}

func transportNodeCollectionsCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_collection"] = bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
	fields["apply_profile"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["transport_node_collection"] = "TransportNodeCollection"
	fieldNameMap["apply_profile"] = "ApplyProfile"
	paramsTypeMap["apply_profile"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["transport_node_collection"] = bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
	queryParams["apply_profile"] = "apply_profile"
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
		"transport_node_collection",
		"POST",
		"/api/v1/transport-node-collections",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeCollectionsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeCollectionsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodeCollectionsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	paramsTypeMap["transport_node_collection_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeCollectionId"] = bindings.NewStringType()
	pathParams["transport_node_collection_id"] = "transportNodeCollectionId"
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
		"/api/v1/transport-node-collections/{transportNodeCollectionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeCollectionsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeCollectionsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
}

func transportNodeCollectionsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	paramsTypeMap["transport_node_collection_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeCollectionId"] = bindings.NewStringType()
	pathParams["transport_node_collection_id"] = "transportNodeCollectionId"
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
		"/api/v1/transport-node-collections/{transportNodeCollectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeCollectionsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_moid"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["compute_collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vc_instance_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_moid"] = "ClusterMoid"
	fieldNameMap["compute_collection_id"] = "ComputeCollectionId"
	fieldNameMap["vc_instance_uuid"] = "VcInstanceUuid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeCollectionsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeCollectionListResultBindingType)
}

func transportNodeCollectionsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_moid"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["compute_collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["vc_instance_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cluster_moid"] = "ClusterMoid"
	fieldNameMap["compute_collection_id"] = "ComputeCollectionId"
	fieldNameMap["vc_instance_uuid"] = "VcInstanceUuid"
	paramsTypeMap["compute_collection_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vc_instance_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cluster_moid"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["cluster_moid"] = "cluster_moid"
	queryParams["compute_collection_id"] = "compute_collection_id"
	queryParams["vc_instance_uuid"] = "vc_instance_uuid"
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
		"/api/v1/transport-node-collections",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeCollectionsRetryprofilerealizationInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeCollectionsRetryprofilerealizationOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func transportNodeCollectionsRetryprofilerealizationRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	paramsTypeMap["transport_node_collection_id"] = bindings.NewStringType()
	paramsTypeMap["transportNodeCollectionId"] = bindings.NewStringType()
	pathParams["transport_node_collection_id"] = "transportNodeCollectionId"
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
		"action=retry_profile_realization",
		"",
		"POST",
		"/api/v1/transport-node-collections/{transportNodeCollectionId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func transportNodeCollectionsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fields["transport_node_collection"] = bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	fieldNameMap["transport_node_collection"] = "TransportNodeCollection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func transportNodeCollectionsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
}

func transportNodeCollectionsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_collection_id"] = bindings.NewStringType()
	fields["transport_node_collection"] = bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
	fieldNameMap["transport_node_collection_id"] = "TransportNodeCollectionId"
	fieldNameMap["transport_node_collection"] = "TransportNodeCollection"
	paramsTypeMap["transport_node_collection_id"] = bindings.NewStringType()
	paramsTypeMap["transport_node_collection"] = bindings.NewReferenceType(model.TransportNodeCollectionBindingType)
	paramsTypeMap["transportNodeCollectionId"] = bindings.NewStringType()
	pathParams["transport_node_collection_id"] = "transportNodeCollectionId"
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
		"transport_node_collection",
		"PUT",
		"/api/v1/transport-node-collections/{transportNodeCollectionId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
