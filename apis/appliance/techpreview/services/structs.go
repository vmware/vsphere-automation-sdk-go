/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Services.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package services

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``ServiceOps`` enumeration class Defines service operations.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ServiceOps string

const (
    // The service status can be generated.
     ServiceOps_status ServiceOps = "status"
    // The service can be stopped.
     ServiceOps_stop ServiceOps = "stop"
    // The service can be started or restarted.
     ServiceOps_restart ServiceOps = "restart"
)

func (s ServiceOps) ServiceOps() bool {
    switch s {
        case ServiceOps_status:
            return true
        case ServiceOps_stop:
            return true
        case ServiceOps_restart:
            return true
        default:
            return false
    }
}





// ``ServiceInfo`` class Structure that describes a service and the operations you can perform on a service.
 type ServiceInfo struct {
    Name string
    Description string
    Operations []ServiceOps
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ServiceInfoBindingType), reflect.TypeOf([]ServiceInfo{}))
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
       map[string]int{"Error": 500})
}


func stopInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fields["timeout"] = bindings.NewIntegerType()
    fieldNameMap["name"] = "Name"
    fieldNameMap["timeout"] = "Timeout"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stopOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stopRestMetadata() protocol.OperationRestMetadata {
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


func restartInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fields["timeout"] = bindings.NewIntegerType()
    fieldNameMap["name"] = "Name"
    fieldNameMap["timeout"] = "Timeout"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func restartOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func restartRestMetadata() protocol.OperationRestMetadata {
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



func ServiceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["operations"] = bindings.NewListType(bindings.NewEnumType("com.vmware.appliance.techpreview.services.service_ops", reflect.TypeOf(ServiceOps(ServiceOps_status))), reflect.TypeOf([]ServiceOps{}))
    fieldNameMap["operations"] = "Operations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.services.service_info",fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}


