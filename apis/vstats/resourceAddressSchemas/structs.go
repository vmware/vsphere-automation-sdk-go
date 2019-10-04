/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ResourceAddressSchemas.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package resourceAddressSchemas

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Resource type for resource addressing schemas.
const ResourceAddressSchemas_RESOURCE_TYPE = "com.vmware.vstats.model.RsrcAddrSchema"


// Declares which predicates are supported by the ``ResourceIdDefinition``.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type QueryCapabilities string

const (
    // Equal.
     QueryCapabilities_EQUAL QueryCapabilities = "EQUAL"
    // Supports both Equality and Aggregation.
     QueryCapabilities_EQUAL_ALL QueryCapabilities = "EQUAL_ALL"
)

func (q QueryCapabilities) QueryCapabilities() bool {
    switch q {
        case QueryCapabilities_EQUAL:
            return true
        case QueryCapabilities_EQUAL_ALL:
            return true
        default:
            return false
    }
}





// The ``ResourceIdDefinition`` class describes a single identifier of the Resource Addressing Schema.
 type ResourceIdDefinition struct {
    Key string
    Type_ string
    QueryOptions QueryCapabilities
}






// The ``Info`` class defines addressing schema for a counter. This is set of named placeholders for different resource types. For example a network link between VMs will take two arguments "source" and "destination" both of type VM. For each argument query capability is defined.
 type Info struct {
    Id string
    Schema []ResourceIdDefinition
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
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
    paramsTypeMap["id"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
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
      "/stats/rsrc-addr-schemas/{id}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func ResourceIdDefinitionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key"] = bindings.NewStringType()
    fieldNameMap["key"] = "Key"
    fields["type"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcType"}, "")
    fieldNameMap["type"] = "Type_"
    fields["query_options"] = bindings.NewEnumType("com.vmware.vstats.resource_address_schemas.query_capabilities", reflect.TypeOf(QueryCapabilities(QueryCapabilities_EQUAL)))
    fieldNameMap["query_options"] = "QueryOptions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.resource_address_schemas.resource_id_definition",fields, reflect.TypeOf(ResourceIdDefinition{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vstats.model.RsrcAddrSchema"}, "")
    fieldNameMap["id"] = "Id"
    fields["schema"] = bindings.NewListType(bindings.NewReferenceType(ResourceIdDefinitionBindingType), reflect.TypeOf([]ResourceIdDefinition{}))
    fieldNameMap["schema"] = "Schema"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.resource_address_schemas.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


