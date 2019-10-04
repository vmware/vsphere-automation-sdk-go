/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for data providers.
const Providers_RESOURCE_TYPE = "com.vmware.vstats.model.Provider"



// ``Summary`` class describes a statistical data provider.
 type Summary struct {
    Id string
    IdValue string
    Type_ string
    Scheme *string
    Metadata *string
    TrackingSn *int64
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
      "GET",
      "/stats/providers",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Provider"}, "")
    fieldNameMap["id"] = "Id"
    fields["id_value"] = bindings.NewStringType()
    fieldNameMap["id_value"] = "IdValue"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["scheme"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["scheme"] = "Scheme"
    fields["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["metadata"] = "Metadata"
    fields["tracking_sn"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["tracking_sn"] = "TrackingSn"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.providers.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


