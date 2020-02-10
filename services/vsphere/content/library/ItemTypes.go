/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Item.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package library

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for item.
const Item_RESOURCE_TYPE = "com.vmware.content.library.Item"


// The ``FindSpec`` class specifies the properties that can be used as a filter to find library items. When multiple properties are specified, all properties of the item must match the specification.
type ItemFindSpec struct {
    // The name of the library item. The name is case-insensitive. See ItemModel#name.
	Name *string
    // The identifier of the library containing the item. See ItemModel#libraryId.
	LibraryId *string
    // The identifier of the library item as reported by the publisher. See ItemModel#sourceId.
	SourceId *string
    // The type of the library item. The type is case-insensitive. See ItemModel#type.
	Type_ *string
    // Whether the item is cached. Possible values are 'true' or 'false'. See ItemModel#cached.
	Cached *bool
}

// The ``DestinationSpec`` class contains information required to publish the library item to a specific subscription. This class was added in vSphere API 6.7.2.
type ItemDestinationSpec struct {
    // Identifier of the subscription associated with the subscribed library. This property was added in vSphere API 6.7.2.
	Subscription string
}



func itemCopyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source_library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["destination_create_spec"] = bindings.NewReferenceType(ItemModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["source_library_item_id"] = "SourceLibraryItemId"
	fieldNameMap["destination_create_spec"] = "DestinationCreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemCopyOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
}

func itemCopyRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source_library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["destination_create_spec"] = bindings.NewReferenceType(ItemModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["source_library_item_id"] = "SourceLibraryItemId"
	fieldNameMap["destination_create_spec"] = "DestinationCreateSpec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"ResourceInaccessible": 500,"NotAllowedInCurrentState": 400})
}

func itemCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(ItemModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
}

func itemCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["create_spec"] = bindings.NewReferenceType(ItemModelBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["create_spec"] = "CreateSpec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"AlreadyExists": 400})
}

func itemDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func itemDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidElementType": 400,"NotFound": 404,"NotAllowedInCurrentState": 400})
}

func itemGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ItemModelBindingType)
}

func itemGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}

func itemListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""), reflect.TypeOf([]string{}))
}

func itemListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_id"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library_id"] = "LibraryId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}

func itemFindInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ItemFindSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemFindOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""), reflect.TypeOf([]string{}))
}

func itemFindRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(ItemFindSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400})
}

func itemUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["update_spec"] = bindings.NewReferenceType(ItemModelBindingType)
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["update_spec"] = "UpdateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func itemUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["update_spec"] = bindings.NewReferenceType(ItemModelBindingType)
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["update_spec"] = "UpdateSpec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"InvalidElementType": 400,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"AlreadyExists": 400})
}

func itemPublishInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["force_sync_content"] = bindings.NewBooleanType()
	fields["subscriptions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ItemDestinationSpecBindingType), reflect.TypeOf([]ItemDestinationSpec{})))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["force_sync_content"] = "ForceSyncContent"
	fieldNameMap["subscriptions"] = "Subscriptions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func itemPublishOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func itemPublishRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["force_sync_content"] = bindings.NewBooleanType()
	fields["subscriptions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ItemDestinationSpecBindingType), reflect.TypeOf([]ItemDestinationSpec{})))
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["force_sync_content"] = "ForceSyncContent"
	fieldNameMap["subscriptions"] = "Subscriptions"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func ItemFindSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["library_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.Library"}, ""))
	fieldNameMap["library_id"] = "LibraryId"
	fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.content.library.Item"}, ""))
	fieldNameMap["source_id"] = "SourceId"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["cached"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cached"] = "Cached"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.find_spec", fields, reflect.TypeOf(ItemFindSpec{}), fieldNameMap, validators)
}

func ItemDestinationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subscription"] = bindings.NewIdType([]string{"com.vmware.content.library.Subscriptions"}, "")
	fieldNameMap["subscription"] = "Subscription"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.destination_spec", fields, reflect.TypeOf(ItemDestinationSpec{}), fieldNameMap, validators)
}

