/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ipv4.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ipv4

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``IPv4Mode`` enumeration class Defines different IPv4 modes.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IPv4Mode string

const (
    // IPv4 address is automatically assigned by a DHCP server.
     IPv4Mode_dhcp IPv4Mode = "dhcp"
    // IPv4 address is static.
     IPv4Mode_is_static IPv4Mode = "is_static"
    // The IPv4 protocol is not configured.
     IPv4Mode_unconfigured IPv4Mode = "unconfigured"
)

func (i IPv4Mode) IPv4Mode() bool {
    switch i {
        case IPv4Mode_dhcp:
            return true
        case IPv4Mode_is_static:
            return true
        case IPv4Mode_unconfigured:
            return true
        default:
            return false
    }
}





// ``IPv4Config`` class Structure that defines the IPv4 configuration state of a network interface.
 type IPv4Config struct {
    InterfaceName string
    Mode IPv4Mode
    Address string
    Prefix int64
    DefaultGateway string
}






// ``IPv4ConfigReadOnly`` class Structure that defines the IPv4 configuration state of a network interface.
 type IPv4ConfigReadOnly struct {
    InterfaceName string
    Mode IPv4Mode
    Address string
    Prefix int64
    DefaultGateway string
    Updateable bool
}









func renewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interfaces"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func renewOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func renewRestMetadata() protocol.OperationRestMetadata {
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


func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewListType(bindings.NewReferenceType(IPv4ConfigBindingType), reflect.TypeOf([]IPv4Config{}))
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func setOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func setRestMetadata() protocol.OperationRestMetadata {
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


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(IPv4ConfigReadOnlyBindingType), reflect.TypeOf([]IPv4ConfigReadOnly{}))
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
    fields["interfaces"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(IPv4ConfigReadOnlyBindingType), reflect.TypeOf([]IPv4ConfigReadOnly{}))
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



func IPv4ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_mode", reflect.TypeOf(IPv4Mode(IPv4Mode_dhcp)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_config",fields, reflect.TypeOf(IPv4Config{}), fieldNameMap, validators)
}

func IPv4ConfigReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_mode", reflect.TypeOf(IPv4Mode(IPv4Mode_dhcp)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    fields["updateable"] = bindings.NewBooleanType()
    fieldNameMap["updateable"] = "Updateable"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_config_read_only",fields, reflect.TypeOf(IPv4ConfigReadOnly{}), fieldNameMap, validators)
}


