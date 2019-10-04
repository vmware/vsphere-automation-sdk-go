/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TagAssociation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagAssociation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``BatchResult`` class describes the result of performing the same method on several tags or objects in a single invocation.
 type BatchResult struct {
    Success bool
    ErrorMessages []std.LocalizableMessage
}






// The ``TagToObjects`` class describes a tag and its related objects. Use the TagAssociation#listAttachedObjectsOnTags method to retrieve a array with each element containing a tag and the objects to which it is attached.
 type TagToObjects struct {
    TagId string
    ObjectIds []std.DynamicID
}






// The ``ObjectToTags`` class describes an object and its related tags. Use the TagAssociation#listAttachedTagsOnObjects method to retrieve a array with each element containing an object and the tags attached to it.
 type ObjectToTags struct {
    ObjectId std.DynamicID
    TagIds []string
}









func attachInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["tag_id"] = "TagId"
    fieldNameMap["object_id"] = "ObjectId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func attachOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func attachRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthorized": 403,"Unauthenticated": 401})
}


func attachMultipleTagsToObjectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["object_id"] = "ObjectId"
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func attachMultipleTagsToObjectOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BatchResultBindingType)
}

func attachMultipleTagsToObjectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}


func attachTagToMultipleObjectsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
    fieldNameMap["tag_id"] = "TagId"
    fieldNameMap["object_ids"] = "ObjectIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func attachTagToMultipleObjectsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BatchResultBindingType)
}

func attachTagToMultipleObjectsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
}


func detachInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["tag_id"] = "TagId"
    fieldNameMap["object_id"] = "ObjectId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detachOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func detachRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
}


func detachMultipleTagsFromObjectInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["object_id"] = "ObjectId"
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detachMultipleTagsFromObjectOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BatchResultBindingType)
}

func detachMultipleTagsFromObjectRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}


func detachTagFromMultipleObjectsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
    fieldNameMap["tag_id"] = "TagId"
    fieldNameMap["object_ids"] = "ObjectIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func detachTagFromMultipleObjectsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BatchResultBindingType)
}

func detachTagFromMultipleObjectsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
}


func listAttachedObjectsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fieldNameMap["tag_id"] = "TagId"
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
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


func listAttachedTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["object_id"] = "ObjectId"
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
       map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
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


func listAttachableTagsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["object_id"] = "ObjectId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listAttachableTagsOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func listAttachableTagsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"Unauthenticated": 401})
}



func BatchResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["success"] = bindings.NewBooleanType()
    fieldNameMap["success"] = "Success"
    fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["error_messages"] = "ErrorMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.tag_association.batch_result",fields, reflect.TypeOf(BatchResult{}), fieldNameMap, validators)
}

func TagToObjectsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, "")
    fieldNameMap["tag_id"] = "TagId"
    fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
    fieldNameMap["object_ids"] = "ObjectIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.tag_association.tag_to_objects",fields, reflect.TypeOf(TagToObjects{}), fieldNameMap, validators)
}

func ObjectToTagsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
    fieldNameMap["object_id"] = "ObjectId"
    fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["tag_ids"] = "TagIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.tag_association.object_to_tags",fields, reflect.TypeOf(ObjectToTags{}), fieldNameMap, validators)
}


