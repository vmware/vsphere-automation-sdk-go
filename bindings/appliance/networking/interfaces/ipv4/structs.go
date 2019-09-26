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



// The ``Mode`` enumeration class defines different IPv4 address assignment modes.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Mode string

const (
    // The IPv4 address is automatically assigned by a DHCP server.
     Mode_DHCP Mode = "DHCP"
    // The IPv4 address is static.
     Mode_STATIC Mode = "STATIC"
    // The IPv4 protocol is not configured.
     Mode_UNCONFIGURED Mode = "UNCONFIGURED"
)

func (m Mode) Mode() bool {
    switch m {
        case Mode_DHCP:
            return true
        case Mode_STATIC:
            return true
        case Mode_UNCONFIGURED:
            return true
        default:
            return false
    }
}





// The ``Config`` class provides defines the IPv4 configuration of a network interface.
 type Config struct {
    Mode Mode
    Address *string
    Prefix *int64
    DefaultGateway *string
}






// The ``Info`` class defines current IPv4 configuration state of a network interface.
 type Info struct {
    Configurable bool
    Mode Mode
    Address *string
    Prefix *int64
    DefaultGateway *string
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fields["config"] = bindings.NewReferenceType(ConfigBindingType)
    fieldNameMap["interface_name"] = "InterfaceName"
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Error": 500})
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
    return bindings.NewReferenceType(InfoBindingType)
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



func ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv4.mode", reflect.TypeOf(Mode(Mode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("prefix", true),
            },
            "DHCP": []bindings.FieldData {},
            "UNCONFIGURED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv4.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configurable"] = bindings.NewBooleanType()
    fieldNameMap["configurable"] = "Configurable"
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv4.mode", reflect.TypeOf(Mode(Mode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("default_gateway", true),
            },
            "DHCP": []bindings.FieldData {
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("default_gateway", true),
            },
            "UNCONFIGURED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv4.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


