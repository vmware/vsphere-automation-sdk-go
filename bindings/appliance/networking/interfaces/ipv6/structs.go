/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ipv6.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ipv6

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Origin`` enumeration class defines IPv6 address origin values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Origin string

const (
    // The IPv6 address is assigned by a DHCP server. See RFC 4293.
     Origin_DHCP Origin = "DHCP"
    // The IPv6 address is assigned randomly by the system. See RFC 4293.
     Origin_RANDOM Origin = "RANDOM"
    // The IPv6 address was manually configured to a specified address, for example, by user configuration. See RFC 4293.
     Origin_MANUAL Origin = "MANUAL"
    // The IPv6 address is assigned by IPv6 Stateless Address Auto-configuration (SLAAC). See RFC 4293.
     Origin_LINKLAYER Origin = "LINKLAYER"
    // The IPv6 address is assigned by a mechanism other than manual, DHCP, SLAAC, or random. See RFC 4293.
     Origin_OTHER Origin = "OTHER"
)

func (o Origin) Origin() bool {
    switch o {
        case Origin_DHCP:
            return true
        case Origin_RANDOM:
            return true
        case Origin_MANUAL:
            return true
        case Origin_LINKLAYER:
            return true
        case Origin_OTHER:
            return true
        default:
            return false
    }
}




// The ``Status`` enumeration class defines IPv6 address status values. See RFC 4293.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Status string

const (
    // The IPv6 address is in the process of being verified as unique. An address in this state cannot be used for general communication. It can be used to determine the uniqueness of the address.
     Status_TENTATIVE Status = "TENTATIVE"
    // The status of this address cannot be determined.
     Status_UNKNOWN Status = "UNKNOWN"
    // The IPv6 address is inaccessible because the interface to which this address is assigned is not operational.
     Status_INACCESSIBLE Status = "INACCESSIBLE"
    // The IPv6 address is not a valid address. It should not appear as the destination or source address of a packet.
     Status_INVALID Status = "INVALID"
    // The IPv6 address is not unique on the link and cannot be used.
     Status_DUPLICATE Status = "DUPLICATE"
    // This is a valid IPv6 address that can appear as the destination or source address of a packet.
     Status_PREFERRED Status = "PREFERRED"
    // The is a valid but deprecated IPv6 address. This address cannot be used as a source address in new communications, although packets addressed to such an address are processed as expected.
     Status_DEPRECATED Status = "DEPRECATED"
    // The IPv6 address is available for use, subject to restrictions, while its uniqueness on a link is being verified.
     Status_OPTIMISTIC Status = "OPTIMISTIC"
)

func (s Status) Status() bool {
    switch s {
        case Status_TENTATIVE:
            return true
        case Status_UNKNOWN:
            return true
        case Status_INACCESSIBLE:
            return true
        case Status_INVALID:
            return true
        case Status_DUPLICATE:
            return true
        case Status_PREFERRED:
            return true
        case Status_DEPRECATED:
            return true
        case Status_OPTIMISTIC:
            return true
        default:
            return false
    }
}





// The ``Address`` class provides the structure used to name an IPv6 address.
 type Address struct {
    Address string
    Prefix int64
}






// The ``AddressInfo`` class provides the structure that you can use to get information about an IPv6 address along with its origin and status.
 type AddressInfo struct {
    Origin Origin
    Status Status
    Address string
    Prefix int64
}






// The ``Config`` class provides the structure that you can use to configure IPv6 on a particular interface. Addresses can be assigned by DHCP, SLAAC or STATIC, as IPv6 permits multiple addresses per interface.
 type Config struct {
    Dhcp bool
    Autoconf bool
    Addresses []Address
    DefaultGateway string
}






// The ``Info`` class provides the structure that defines an existing IPv6 configuration on a particular interface. This structure is read only.
 type Info struct {
    Dhcp bool
    Autoconf bool
    Addresses []AddressInfo
    DefaultGateway string
    Configurable bool
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
       map[string]int{"ResourceBusy": 500,"NotFound": 404,"Error": 500})
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



func AddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.address",fields, reflect.TypeOf(Address{}), fieldNameMap, validators)
}

func AddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["origin"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv6.origin", reflect.TypeOf(Origin(Origin_DHCP)))
    fieldNameMap["origin"] = "Origin"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv6.status", reflect.TypeOf(Status(Status_TENTATIVE)))
    fieldNameMap["status"] = "Status"
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.address_info",fields, reflect.TypeOf(AddressInfo{}), fieldNameMap, validators)
}

func ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(AddressBindingType), reflect.TypeOf([]Address{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.config",fields, reflect.TypeOf(Config{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(AddressInfoBindingType), reflect.TypeOf([]AddressInfo{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    fields["configurable"] = bindings.NewBooleanType()
    fieldNameMap["configurable"] = "Configurable"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


