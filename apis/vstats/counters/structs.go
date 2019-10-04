/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Counters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package counters

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for counters.
const Counters_RESOURCE_TYPE = "com.vmware.vstats.model.Counter"



// The ``Info`` class contains the counter information. Counter is derived from metric. It applies the metric to a particular class of a resource.
 type Info struct {
    Cid string
    Metric string
    ResourceAddressSchema string
}






// The ``FilterSpec`` class contains fields that can be used to filter list of counters.
 type FilterSpec struct {
    Cid *string
    Types []string
    Metric *string
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
    return bindings.NewListType(bindings.NewReferenceType(InfoBindingType), reflect.TypeOf([]Info{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.cid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, ""))
    paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
    paramsTypeMap["filter.metric"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Metric"}, ""))
    queryParams["filter.types"] = "types"
    queryParams["filter.metric"] = "metric"
    queryParams["filter.cid"] = "cid"
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
      "/stats/counters",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
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
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    pathParams["cid"] = "cid"
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
      "/stats/counters/{cid}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    fields["metric"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Metric"}, "")
    fieldNameMap["metric"] = "Metric"
    fields["resource_address_schema"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcAddrSchema"}, "")
    fieldNameMap["resource_address_schema"] = "ResourceAddressSchema"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counters.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, ""))
    fieldNameMap["cid"] = "Cid"
    fields["types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["types"] = "Types"
    fields["metric"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Metric"}, ""))
    fieldNameMap["metric"] = "Metric"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counters.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


