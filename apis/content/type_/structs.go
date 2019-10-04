/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Type.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package type_

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// The ``Info`` class describes support for a specific type of data in an library.ItemModel. The ``Info`` can be queried through the Type interface. Type support describes plugins in the Content Library which can provide metadata on library items and help manage the transfer process by adding dependent files when a current file is added.
 type Info struct {
    Description string
    Name string
    Type_ string
    Vendor string
    Version string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["vendor"] = bindings.NewStringType()
    fieldNameMap["vendor"] = "Vendor"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.content.type.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


