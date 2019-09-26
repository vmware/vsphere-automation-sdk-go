/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ResourceAddresses.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package resourceAddresses

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vstats"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for ``ResourceAddresses``
const ResourceAddresses_RESOURCE_TYPE = "com.vmware.vstats.model.RsrcAddr"



// The ``Info`` class contains global address of a specific Resource.
 type Info struct {
    Id string
    ResourceIds []vstats.RsrcId
}






// The ListResult class contains properties used to return the resource addresses.
 type ListResult struct {
    RsrcAddrs []Info
    Next *string
}






// ``FilterSpec`` class describes filter criteria for resource addresses.
 type FilterSpec struct {
    Types []string
    Resources []string
    Page *string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ListResultBindingType)
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
    paramsTypeMap["filter.page"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["filter.resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    queryParams["filter.types"] = "types"
    queryParams["filter.page"] = "page"
    queryParams["filter.resources"] = "rsrcs"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/stats/rsrc-addrs",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcAddr"}, "")
    fieldNameMap["id"] = "Id"
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
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcAddr"}, "")
    paramsTypeMap["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcAddr"}, "")
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/stats/rsrc-addrs/{id}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcAddr"}, "")
    fieldNameMap["id"] = "Id"
    fields["resource_ids"] = bindings.NewListType(bindings.NewReferenceType(vstats.RsrcIdBindingType), reflect.TypeOf([]vstats.RsrcId{}))
    fieldNameMap["resource_ids"] = "ResourceIds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.resource_addresses.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func ListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["rsrc_addrs"] = bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
    fieldNameMap["rsrc_addrs"] = "RsrcAddrs"
    fields["next"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["next"] = "Next"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.resource_addresses.list_result",fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["types"] = "Types"
    fields["resources"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["resources"] = "Resources"
    fields["page"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["page"] = "Page"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.resource_addresses.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


