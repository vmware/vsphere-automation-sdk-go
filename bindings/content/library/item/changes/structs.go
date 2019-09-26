/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Changes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package changes

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "time"
)

// Resource type for library item versions.
const Changes_RESOURCE_TYPE = "com.vmware.content.library.item.Version"



// The ``Summary`` class contains commonly used information about a library item change.
 type Summary struct {
    Version string
    Time time.Time
    User *string
    ShortMessage *string
}






// The ``Info`` class contains information about a library item change.
 type Info struct {
    Time time.Time
    User *string
    Message *string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fieldNameMap["library_item"] = "LibraryItem"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["library_item"] = bindings.NewIdType([]string {"com.vmware.content.library.Item"}, "")
    fields["version"] = bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
    fieldNameMap["library_item"] = "LibraryItem"
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.content.library.item.Version"}, "")
    fieldNameMap["version"] = "Version"
    fields["time"] = bindings.NewDateTimeType()
    fieldNameMap["time"] = "Time"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    fields["short_message"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["short_message"] = "ShortMessage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.changes.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["time"] = bindings.NewDateTimeType()
    fieldNameMap["time"] = "Time"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    fields["message"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.library.item.changes.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


