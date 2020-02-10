/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Interfaces.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/vcenter/vm/guest"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``IpAddressOrigin`` enumeration class specifies how an IP address was obtained for an interface. See RFC 4293 IpAddressOriginTC. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type InterfacesIpAddressOrigin string

const (
    // Any other type of address configuration other than the below mentioned ones will fall under this category. For e.g., automatic address configuration for the link local address falls under this type. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressOrigin_OTHER InterfacesIpAddressOrigin = "OTHER"
    // The address is configured manually. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressOrigin_MANUAL InterfacesIpAddressOrigin = "MANUAL"
    // The address is configured through dhcp. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressOrigin_DHCP InterfacesIpAddressOrigin = "DHCP"
    // The address is obtained through stateless autoconfiguration (autoconf). See RFC 4862, IPv6 Stateless Address Autoconfiguration. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressOrigin_LINKLAYER InterfacesIpAddressOrigin = "LINKLAYER"
    // The address is chosen by the system at random e.g., an IPv4 address within 169.254/16, or an RFC 3041 privacy address. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressOrigin_RANDOM InterfacesIpAddressOrigin = "RANDOM"
)

func (i InterfacesIpAddressOrigin) InterfacesIpAddressOrigin() bool {
	switch i {
	case InterfacesIpAddressOrigin_OTHER:
		return true
	case InterfacesIpAddressOrigin_MANUAL:
		return true
	case InterfacesIpAddressOrigin_DHCP:
		return true
	case InterfacesIpAddressOrigin_LINKLAYER:
		return true
	case InterfacesIpAddressOrigin_RANDOM:
		return true
	default:
		return false
	}
}


// The ``IpAddressStatus`` enumeration class defines the present status of an address on an interface. See RFC 4293 IpAddressStatusTC. This enumeration was added in vSphere API 7.0.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type InterfacesIpAddressStatus string

const (
    // Indicates that this is a valid address. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_PREFERRED InterfacesIpAddressStatus = "PREFERRED"
    // Indicates that this is a valid but deprecated address that should no longer be used as a source address. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_DEPRECATED InterfacesIpAddressStatus = "DEPRECATED"
    // Indicates that this isn't a valid address. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_INVALID InterfacesIpAddressStatus = "INVALID"
    // Indicates that the address is not accessible because interface is not operational. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_INACCESSIBLE InterfacesIpAddressStatus = "INACCESSIBLE"
    // Indicates that the status cannot be determined. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_UNKNOWN InterfacesIpAddressStatus = "UNKNOWN"
    // Indicates that the uniqueness of the address on the link is presently being verified. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_TENTATIVE InterfacesIpAddressStatus = "TENTATIVE"
    // Indicates the address has been determined to be non-unique on the link, this address will not be reachable. This constant field was added in vSphere API 7.0.0.
	InterfacesIpAddressStatus_DUPLICATE InterfacesIpAddressStatus = "DUPLICATE"
)

func (i InterfacesIpAddressStatus) InterfacesIpAddressStatus() bool {
	switch i {
	case InterfacesIpAddressStatus_PREFERRED:
		return true
	case InterfacesIpAddressStatus_DEPRECATED:
		return true
	case InterfacesIpAddressStatus_INVALID:
		return true
	case InterfacesIpAddressStatus_INACCESSIBLE:
		return true
	case InterfacesIpAddressStatus_UNKNOWN:
		return true
	case InterfacesIpAddressStatus_TENTATIVE:
		return true
	case InterfacesIpAddressStatus_DUPLICATE:
		return true
	default:
		return false
	}
}


// The ``IpAddressInfo`` class describes a specific IP Address. This class was added in vSphere API 7.0.0.
type InterfacesIpAddressInfo struct {
    // IPv4 address is specified using dotted decimal notation. For example, "192.0.2.1". IPv6 addresses are 128-bit addresses specified using eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of the symbol '::' to represent multiple 16-bit groups of contiguous 0's only once in an address as described in RFC 2373. This property was added in vSphere API 7.0.0.
	IpAddress string
    // Denotes the length of a generic Internet network address prefix. Prefix length: the valid range of values is 0-32 for IPv4, and 0-128 for IPv6. A value of n corresponds to an IP address mask that has n contiguous 1-bits from the most significant bit (MSB), with all other bits set to 0. A value of zero is valid only if the calling context defines it. This property was added in vSphere API 7.0.0.
	PrefixLength int64
    // How this address was configured. This property was added in vSphere API 7.0.0.
	Origin *InterfacesIpAddressOrigin
    // The state of this ipAddress. This property was added in vSphere API 7.0.0.
	State InterfacesIpAddressStatus
}

// The ``IpConfigInfo`` class describes the protocol version independent address reporting data object for network interfaces. This class was added in vSphere API 7.0.0.
type InterfacesIpConfigInfo struct {
    // IP addresses configured on the interface. This property was added in vSphere API 7.0.0.
	IpAddresses []InterfacesIpAddressInfo
    // Client side DHCP for an interface. This property was added in vSphere API 7.0.0.
	Dhcp *guest.DhcpConfigInfo
}

// The ``Info`` class describes a virtual network adapter configured in the guest operating system. This class was added in vSphere API 7.0.0.
type InterfacesInfo struct {
    // Client DNS values. Data assigned by DNS. This property was added in vSphere API 7.0.0.
	DnsValues *guest.DnsAssignedValues
    // MAC address of the adapter. This property was added in vSphere API 7.0.0.
	MacAddress *string
    // DNS configuration of the adapter. See guest.NetworkingInfo#dns for system wide settings. This property was added in vSphere API 7.0.0.
	Dns *guest.DnsConfigInfo
    // IP configuration settings of the adapter. This property was added in vSphere API 7.0.0.
	Ip *InterfacesIpConfigInfo
    // The IP addresses of any WINS name servers for the adapter. This property was added in vSphere API 7.0.0.
	WinsServers []string
    // Link to the corresponding virtual device. This property was added in vSphere API 7.0.0.
	Nic *string
}



func interfacesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(InterfacesInfoBindingType), reflect.TypeOf([]InterfacesInfo{}))
}

func interfacesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}


func InterfacesIpAddressInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_address"] = bindings.NewStringType()
	fieldNameMap["ip_address"] = "IpAddress"
	fields["prefix_length"] = bindings.NewIntegerType()
	fieldNameMap["prefix_length"] = "PrefixLength"
	fields["origin"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_address_origin", reflect.TypeOf(InterfacesIpAddressOrigin(InterfacesIpAddressOrigin_OTHER))))
	fieldNameMap["origin"] = "Origin"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_address_status", reflect.TypeOf(InterfacesIpAddressStatus(InterfacesIpAddressStatus_PREFERRED)))
	fieldNameMap["state"] = "State"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_address_info", fields, reflect.TypeOf(InterfacesIpAddressInfo{}), fieldNameMap, validators)
}

func InterfacesIpConfigInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_addresses"] = bindings.NewListType(bindings.NewReferenceType(InterfacesIpAddressInfoBindingType), reflect.TypeOf([]InterfacesIpAddressInfo{}))
	fieldNameMap["ip_addresses"] = "IpAddresses"
	fields["dhcp"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.DhcpConfigInfoBindingType))
	fieldNameMap["dhcp"] = "Dhcp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.interfaces.ip_config_info", fields, reflect.TypeOf(InterfacesIpConfigInfo{}), fieldNameMap, validators)
}

func InterfacesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dns_values"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.DnsAssignedValuesBindingType))
	fieldNameMap["dns_values"] = "DnsValues"
	fields["mac_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["mac_address"] = "MacAddress"
	fields["dns"] = bindings.NewOptionalType(bindings.NewReferenceType(guest.DnsConfigInfoBindingType))
	fieldNameMap["dns"] = "Dns"
	fields["ip"] = bindings.NewOptionalType(bindings.NewReferenceType(InterfacesIpConfigInfoBindingType))
	fieldNameMap["ip"] = "Ip"
	fields["wins_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["wins_servers"] = "WinsServers"
	fields["nic"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, ""))
	fieldNameMap["nic"] = "Nic"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.interfaces.info", fields, reflect.TypeOf(InterfacesInfo{}), fieldNameMap, validators)
}

