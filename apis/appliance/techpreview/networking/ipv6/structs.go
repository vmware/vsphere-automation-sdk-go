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



// ``IPv6AddressOrigin`` enumeration class Defines IPv6 address origin values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IPv6AddressOrigin string

const (
    // The IPv6 address is assigned by a DHCP server. See RFC 4293.
     IPv6AddressOrigin_dhcp IPv6AddressOrigin = "dhcp"
    // The IPv6 address is assigned randomly by the system. See RFC 4293.
     IPv6AddressOrigin_random IPv6AddressOrigin = "random"
    // The IPv6 address was manually configured to a specified address, for, example, by user configuration. See RFC 4293.
     IPv6AddressOrigin_manual IPv6AddressOrigin = "manual"
    // The IPv6 address is assigned by a mechanism other than manual, DHCP, SLAAC, or random. See RFC 4293.
     IPv6AddressOrigin_other IPv6AddressOrigin = "other"
    // The IPv6 address is assigned by IPv6 Stateless Address Auto-configuration (SLAAC). See RFC 4293.
     IPv6AddressOrigin_linklayer IPv6AddressOrigin = "linklayer"
)

func (i IPv6AddressOrigin) IPv6AddressOrigin() bool {
    switch i {
        case IPv6AddressOrigin_dhcp:
            return true
        case IPv6AddressOrigin_random:
            return true
        case IPv6AddressOrigin_manual:
            return true
        case IPv6AddressOrigin_other:
            return true
        case IPv6AddressOrigin_linklayer:
            return true
        default:
            return false
    }
}




// ``IPv6AddressStatus`` enumeration class Defines IPv6 address status values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IPv6AddressStatus string

const (
    // This IPv6 address is in the process of being verified as unique. Do not use addresses in this state for general communication. You can use them to determine the uniqueness of the address. See RFC 4293.
     IPv6AddressStatus_tentative IPv6AddressStatus = "tentative"
    // The status of this address cannot be determined. See RFC 4293.
     IPv6AddressStatus_unknown IPv6AddressStatus = "unknown"
    // This IPv6 address is inaccessible because the interface to which this address is assigned is not operational. See RFC 4293.
     IPv6AddressStatus_inaccessible IPv6AddressStatus = "inaccessible"
    // This IPv6 address is not a valid address. It should not appear as the destination or source address of a packet. See RFC 4293.
     IPv6AddressStatus_invalid IPv6AddressStatus = "invalid"
    // This IPv6 address is not unique on the link. Do use this IPv6 address. See RFC 4293.
     IPv6AddressStatus_duplicate IPv6AddressStatus = "duplicate"
    // This is a valid IPv6 address that can appear as the destination or source address of a packet. See RFC 4293.
     IPv6AddressStatus_preferred IPv6AddressStatus = "preferred"
    // This is a valid but deprecated IPv6 address. Do not use this IPv6 address as a source address in new communications, although packets addressed to such an address are processed as expected. See RFC 4293.
     IPv6AddressStatus_deprecated IPv6AddressStatus = "deprecated"
    // This IPv6 address is available for use, subject to restrictions, while its uniqueness on a link is being verified. See RFC 4293.
     IPv6AddressStatus_optimistic IPv6AddressStatus = "optimistic"
)

func (i IPv6AddressStatus) IPv6AddressStatus() bool {
    switch i {
        case IPv6AddressStatus_tentative:
            return true
        case IPv6AddressStatus_unknown:
            return true
        case IPv6AddressStatus_inaccessible:
            return true
        case IPv6AddressStatus_invalid:
            return true
        case IPv6AddressStatus_duplicate:
            return true
        case IPv6AddressStatus_preferred:
            return true
        case IPv6AddressStatus_deprecated:
            return true
        case IPv6AddressStatus_optimistic:
            return true
        default:
            return false
    }
}





// ``IPv6AddressReadOnly`` class Structure that you can use to get information about an IPv6 address along with its origin and status.
 type IPv6AddressReadOnly struct {
    Address string
    Prefix int64
    Origin IPv6AddressOrigin
    Status IPv6AddressStatus
}






// ``IPv6ConfigReadOnly`` class Structure that defines an existing IPv6 configuration on a particular interface. This structure is read only.
 type IPv6ConfigReadOnly struct {
    InterfaceName string
    Dhcp bool
    Autoconf bool
    Addresses []IPv6AddressReadOnly
    DefaultGateway string
    Updateable bool
}






// ``IPv6Address`` class Structure used to name an IPv6 address.
 type IPv6Address struct {
    Address string
    Prefix int64
}






// ``IPv6Config`` class Structure that you can use to configure IPv6 on a particular interface. Because IPv6 permits multiple addresses per interface, addresses can be assigned by DHCP, SLAAC, and can also be statically assigned.
 type IPv6Config struct {
    InterfaceName string
    Dhcp bool
    Autoconf bool
    Addresses []IPv6Address
    DefaultGateway string
}









func setInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewListType(bindings.NewReferenceType(IPv6ConfigBindingType), reflect.TypeOf([]IPv6Config{}))
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
    return bindings.NewListType(bindings.NewReferenceType(IPv6ConfigReadOnlyBindingType), reflect.TypeOf([]IPv6ConfigReadOnly{}))
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
    return bindings.NewListType(bindings.NewReferenceType(IPv6ConfigReadOnlyBindingType), reflect.TypeOf([]IPv6ConfigReadOnly{}))
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



func IPv6AddressReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["origin"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address_origin", reflect.TypeOf(IPv6AddressOrigin(IPv6AddressOrigin_dhcp)))
    fieldNameMap["origin"] = "Origin"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address_status", reflect.TypeOf(IPv6AddressStatus(IPv6AddressStatus_tentative)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address_read_only",fields, reflect.TypeOf(IPv6AddressReadOnly{}), fieldNameMap, validators)
}

func IPv6ConfigReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(IPv6AddressReadOnlyBindingType), reflect.TypeOf([]IPv6AddressReadOnly{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    fields["updateable"] = bindings.NewBooleanType()
    fieldNameMap["updateable"] = "Updateable"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_config_read_only",fields, reflect.TypeOf(IPv6ConfigReadOnly{}), fieldNameMap, validators)
}

func IPv6AddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address",fields, reflect.TypeOf(IPv6Address{}), fieldNameMap, validators)
}

func IPv6ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(IPv6AddressBindingType), reflect.TypeOf([]IPv6Address{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_config",fields, reflect.TypeOf(IPv6Config{}), fieldNameMap, validators)
}


