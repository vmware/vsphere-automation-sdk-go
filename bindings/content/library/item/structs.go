/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Item.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/content/library"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for item.
const Item_RESOURCE_TYPE = "com.vmware.content.library.Item"



// The ``FindSpec`` class specifies the properties that can be used as a filter to find library items. When multiple properties are specified, all properties of the item must match the specification.
 type FindSpec struct {
    Name *string
    LibraryId *string
    SourceId *string
    Type_ *string
    Cached *bool
}






// The ``DestinationSpec`` class contains information required to publish the library item to a specific subscription.
 type DestinationSpec struct {
    Subscription string
}









func copyInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["source_library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["destination_create_spec"] = bindings.NewReferenceType(library.ItemModelBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["source_library_item_id"] = "SourceLibraryItemId"
    fieldNameMap["destination_create_spec"] = "DestinationCreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func copyOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
}

func copyRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"ResourceInaccessible": 500,"NotAllowedInCurrentState": 400})
}


func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["create_spec"] = bindings.NewReferenceType(library.ItemModelBindingType)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"AlreadyExists": 400})
}


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["library_item_id"] = "LibraryItemId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidElementType": 400,"NotFound": 404,"NotAllowedInCurrentState": 400})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["library_item_id"] = "LibraryItemId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(library.ItemModelBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404})
}


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_id"] = bindings.NewIdType([]string {"com.vmware.content.Library"}, "")
    fieldNameMap["library_id"] = "LibraryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""), reflect.TypeOf([]string{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404})
}


func findInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(FindSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func findOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""), reflect.TypeOf([]string{}))
}

func findRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400})
}


func updateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["update_spec"] = bindings.NewReferenceType(library.ItemModelBindingType)
    fieldNameMap["library_item_id"] = "LibraryItemId"
    fieldNameMap["update_spec"] = "UpdateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidElementType": 400,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"AlreadyExists": 400})
}


func publishInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item_id"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["force_sync_content"] = bindings.NewBooleanType()
    fields["subscriptions"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DestinationSpecBindingType), reflect.TypeOf([]DestinationSpec{})))
    fieldNameMap["library_item_id"] = "LibraryItemId"
    fieldNameMap["force_sync_content"] = "ForceSyncContent"
    fieldNameMap["subscriptions"] = "Subscriptions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func publishOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func publishRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"InvalidArgument": 400,"InvalidElementType": 400,"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func FindSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["library_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["library_id"] = "LibraryId"
    fields["source_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["source_id"] = "SourceId"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["cached"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cached"] = "Cached"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.find_spec",fields, reflect.TypeOf(FindSpec{}), fieldNameMap, validators)
}

func DestinationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["subscription"] = bindings.NewIdType([]string {"com.vmware.content.library.Subscriptions"}, "")
    fieldNameMap["subscription"] = "Subscription"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.destination_spec",fields, reflect.TypeOf(DestinationSpec{}), fieldNameMap, validators)
}


