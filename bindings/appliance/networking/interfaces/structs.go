/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Interfaces.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package interfaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance/networking/interfaces/ipv4"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/appliance/networking/interfaces/ipv6"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``InterfaceStatus`` enumeration class Defines interface status
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type InterfaceStatus string

const (
    // The interface is down.
     InterfaceStatus_down InterfaceStatus = "down"
    // The interface is up.
     InterfaceStatus_up InterfaceStatus = "up"
)

func (i InterfaceStatus) InterfaceStatus() bool {
    switch i {
        case InterfaceStatus_down:
            return true
        case InterfaceStatus_up:
            return true
        default:
            return false
    }
}





// ``InterfaceInfo`` class Structure that defines properties and status of a network interface.
 type InterfaceInfo struct {
    Name string
    Status InterfaceStatus
    Mac string
    Ipv4 *ipv4.Info
    Ipv6 *ipv6.Info
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InterfaceInfoBindingType), reflect.TypeOf([]InterfaceInfo{}))
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InterfaceInfoBindingType)
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
       map[string]int{"NotFound": 404,"Error": 500})
}



func InterfaceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.interface_status", reflect.TypeOf(InterfaceStatus(InterfaceStatus_down)))
    fieldNameMap["status"] = "Status"
    fields["mac"] = bindings.NewStringType()
    fieldNameMap["mac"] = "Mac"
    fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(ipv4.InfoBindingType))
    fieldNameMap["ipv4"] = "Ipv4"
    fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(ipv6.InfoBindingType))
    fieldNameMap["ipv6"] = "Ipv6"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.interface_info",fields, reflect.TypeOf(InterfaceInfo{}), fieldNameMap, validators)
}


