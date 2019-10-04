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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/vcenter/vm/guest"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``IpAddressOrigin`` enumeration class specifies how an IP address was obtained for an interface. See RFC 4293 IpAddressOriginTC.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IpAddressOrigin string

const (
    // Any other type of address configuration other than the below mentioned ones will fall under this category. For e.g., automatic address configuration for the link local address falls under this type.
     IpAddressOrigin_OTHER IpAddressOrigin = "OTHER"
    // The address is configured manually.
     IpAddressOrigin_MANUAL IpAddressOrigin = "MANUAL"
    // The address is configured through dhcp.
     IpAddressOrigin_DHCP IpAddressOrigin = "DHCP"
    // The address is obtained through stateless autoconfiguration (autoconf). See RFC 4862, IPv6 Stateless Address Autoconfiguration.
     IpAddressOrigin_LINKLAYER IpAddressOrigin = "LINKLAYER"
    // The address is chosen by the system at random e.g., an IPv4 address within 169.254/16, or an RFC 3041 privacy address.
     IpAddressOrigin_RANDOM IpAddressOrigin = "RANDOM"
)

func (i IpAddressOrigin) IpAddressOrigin() bool {
    switch i {
        case IpAddressOrigin_OTHER:
            return true
        case IpAddressOrigin_MANUAL:
            return true
        case IpAddressOrigin_DHCP:
            return true
        case IpAddressOrigin_LINKLAYER:
            return true
        case IpAddressOrigin_RANDOM:
            return true
        default:
            return false
    }
}




// The ``IpAddressStatus`` enumeration class defines the present status of an address on an interface. See RFC 4293 IpAddressStatusTC.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type IpAddressStatus string

const (
    // Indicates that this is a valid address.
     IpAddressStatus_PREFERRED IpAddressStatus = "PREFERRED"
    // Indicates that this is a valid but deprecated address that should no longer be used as a source address.
     IpAddressStatus_DEPRECATED IpAddressStatus = "DEPRECATED"
    // Indicates that this isn't a valid address.
     IpAddressStatus_INVALID IpAddressStatus = "INVALID"
    // Indicates that the address is not accessible because interface is not operational.
     IpAddressStatus_INACCESSIBLE IpAddressStatus = "INACCESSIBLE"
    // Indicates that the status cannot be determined.
     IpAddressStatus_UNKNOWN IpAddressStatus = "UNKNOWN"
    // Indicates that the uniqueness of the address on the link is presently being verified.
     IpAddressStatus_TENTATIVE IpAddressStatus = "TENTATIVE"
    // Indicates the address has been determined to be non-unique on the link, this address will not be reachable.
     IpAddressStatus_DUPLICATE IpAddressStatus = "DUPLICATE"
)

func (i IpAddressStatus) IpAddressStatus() bool {
    switch i {
        case IpAddressStatus_PREFERRED:
            return true
        case IpAddressStatus_DEPRECATED:
            return true
        case IpAddressStatus_INVALID:
            return true
        case IpAddressStatus_INACCESSIBLE:
            return true
        case IpAddressStatus_UNKNOWN:
            return true
        case IpAddressStatus_TENTATIVE:
            return true
        case IpAddressStatus_DUPLICATE:
            return true
        default:
            return false
    }
}





// The ``IpAddressInfo`` class describes a specific IP Address.
 type IpAddressInfo struct {
    IpAddress string
    PrefixLength int64
    Origin *IpAddressOrigin
    State IpAddressStatus
}






// The ``IpConfigInfo`` class describes the protocol version independent address reporting data object for network interfaces.
 type IpConfigInfo struct {
    IpAddresses []IpAddressInfo
    Dhcp *guest.DhcpConfigInfo
}






// The ``Info`` class describes a virtual network adapter configured in the guest operating system.
 type Info struct {
    DnsValues *guest.DnsAssignedValues
    MacAddress *string
    Dns *guest.DnsConfigInfo
    Ip *IpConfigInfo
    WinsServers []string
    Nic *string
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
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
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/vm/{vm}/guest/networking/interfaces",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}



func IpAddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_address"] = bindings.NewStringType()
    fieldNameMap["ip_address"] = "IpAddress"
    fields["prefix_length"] = bindings.NewIntegerType()
    fieldNameMap["prefix_length"] = "PrefixLength"
    fields["origin"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_address_origin", reflect.TypeOf(IpAddressOrigin(IpAddressOrigin_OTHER))))
    fieldNameMap["origin"] = "Origin"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_address_status", reflect.TypeOf(IpAddressStatus(IpAddressStatus_PREFERRED)))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_address_info",fields, reflect.TypeOf(IpAddressInfo{}), fieldNameMap, validators)
}

func IpConfigInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_addresses"] = bindings.NewListType(bindings.NewReferenceType(IpAddressInfoBindingType), reflect.TypeOf([]IpAddressInfo{}))
    fieldNameMap["ip_addresses"] = "IpAddresses"
    fields["dhcp"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.DhcpConfigInfoBindingType))
    fieldNameMap["dhcp"] = "Dhcp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_config_info",fields, reflect.TypeOf(IpConfigInfo{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dns_values"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.DnsAssignedValuesBindingType))
    fieldNameMap["dns_values"] = "DnsValues"
    fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["mac_address"] = "MacAddress"
    fields["dns"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.DnsConfigInfoBindingType))
    fieldNameMap["dns"] = "Dns"
    fields["ip"] = bindings.NewOptionalType(bindings.NewReferenceType(IpConfigInfoBindingType))
    fieldNameMap["ip"] = "Ip"
    fields["wins_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["wins_servers"] = "WinsServers"
    fields["nic"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, ""))
    fieldNameMap["nic"] = "Nic"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.interfaces.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


