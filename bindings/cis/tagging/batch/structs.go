/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Batch.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package batch

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/cis/tagging"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``TagToObjects`` class describes a tag and its related objects. Use the Batch#listAttachedObjectsOnTags method or the Batch#listAllAttachedObjectsOnTags method to retrieve a array with each element containing a tag and objects its attached to.
 type TagToObjects struct {
    TagId string
    ObjectIds []std.DynamicID
}






// The ``ObjectToTags`` class describes an object and its related tags. Use the Batch#listAttachedTagsOnObjects method to retrieve a array with each element containing an object and tags attached to it.
 type ObjectToTags struct {
    ObjectId std.DynamicID
    TagIds []string
}









func getCategoriesInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["category_ids"] = "CategoryIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getCategoriesOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(tagging.CategoryModelBindingType), reflect.TypeOf([]tagging.CategoryModel{}))
}

func getCategoriesRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func getAllCategoriesInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getAllCategoriesOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(tagging.CategoryModelBindingType), reflect.TypeOf([]tagging.CategoryModel{}))
}

func getAllCategoriesRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func getTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getTagsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(tagging.TagModelBindingType), reflect.TypeOf([]tagging.TagModel{}))
}

func getTagsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func getAllTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getAllTagsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(tagging.TagModelBindingType), reflect.TypeOf([]tagging.TagModel{}))
}

func getAllTagsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func listTagsForCategoriesInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["category_ids"] = "CategoryIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listTagsForCategoriesOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func listTagsForCategoriesRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func findTagsByNameInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_name"] = bindings.NewStringType()
    fieldNameMap["tag_name"] = "TagName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func findTagsByNameOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func findTagsByNameRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func listAttachedObjectsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listAttachedObjectsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
}

func listAttachedObjectsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401})
}


func listAttachedObjectsOnTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listAttachedObjectsOnTagsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(TagToObjectsBindingType), reflect.TypeOf([]TagToObjects{}))
}

func listAttachedObjectsOnTagsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401})
}


func listAllAttachedObjectsOnTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listAllAttachedObjectsOnTagsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(TagToObjectsBindingType), reflect.TypeOf([]TagToObjects{}))
}

func listAllAttachedObjectsOnTagsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401})
}


func listAttachedTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
    fieldNameMap["object_ids"] = "ObjectIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listAttachedTagsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func listAttachedTagsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401})
}


func listAttachedTagsOnObjectsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
    fieldNameMap["object_ids"] = "ObjectIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listAttachedTagsOnObjectsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ObjectToTagsBindingType), reflect.TypeOf([]ObjectToTags{}))
}

func listAttachedTagsOnObjectsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401})
}



func TagToObjectsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fieldNameMap["tag_id"] = "TagId"
    fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
    fieldNameMap["object_ids"] = "ObjectIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.batch.tag_to_objects",fields, reflect.TypeOf(TagToObjects{}), fieldNameMap, validators)
}

func ObjectToTagsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["object_id"] = "ObjectId"
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.batch.object_to_tags",fields, reflect.TypeOf(ObjectToTags{}), fieldNameMap, validators)
}


