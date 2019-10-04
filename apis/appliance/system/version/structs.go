/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Version.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package version

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)




// ``VersionStruct`` class Structure representing appliance version information.
 type VersionStruct struct {
    Version string
    Product string
    Build string
    Type_ string
    Summary string
    Releasedate string
    InstallTime string
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(VersionStructBindingType)
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
       map[string]int{"Error": 500})
}



func VersionStructBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["product"] = bindings.NewStringType()
    fieldNameMap["product"] = "Product"
    fields["build"] = bindings.NewStringType()
    fieldNameMap["build"] = "Build"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["summary"] = bindings.NewStringType()
    fieldNameMap["summary"] = "Summary"
    fields["releasedate"] = bindings.NewStringType()
    fieldNameMap["releasedate"] = "Releasedate"
    fields["install_time"] = bindings.NewStringType()
    fieldNameMap["install_time"] = "InstallTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.system.version.version_struct",fields, reflect.TypeOf(VersionStruct{}), fieldNameMap, validators)
}


