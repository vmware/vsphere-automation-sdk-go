/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.vm.guest.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The {\\\\@name DnsAssignedValues) class describes values assigned by a Domain Name Server (DNS). This class was added in vSphere API 7.0.0.
type DnsAssignedValues struct {
    // The host name portion of DNS name. For example, "esx01" part of esx01.example.com. This property was added in vSphere API 7.0.0.
	HostName string
    // The domain name portion of the DNS name. "example.com" part of esx01.example.com. This property was added in vSphere API 7.0.0.
	DomainName string
}

// The ``DnsConfigInfo`` class describes the configuration of RFC 1034 DNS settings. This class was added in vSphere API 7.0.0.
type DnsConfigInfo struct {
    // The IP addresses of the DNS servers in order of use. IPv4 addresses are specified using dotted decimal notation. For example, "192.0.2.1". IPv6 addresses are 128-bit addresses represented as eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of the symbol '::' to represent multiple 16-bit groups of contiguous 0's only once in an address as described in RFC 2373. This property was added in vSphere API 7.0.0.
	IpAddresses []string
    // The domain in which to search for hosts, placed in order of preference. These are the domain name portion of the DNS names. This property was added in vSphere API 7.0.0.
	SearchDomains []string
}

// The ``DhcpConfigInfo`` class specifies when Dynamic Host Configuration Protocol is enabled. This class was added in vSphere API 7.0.0.
type DhcpConfigInfo struct {
    // True if IPv4 DHCP is enabled, false otherwise. This property was added in vSphere API 7.0.0.
	Ipv4Enabled bool
    // True if IPv6 DHCP is enabled, false otherwise. This property was added in vSphere API 7.0.0.
	Ipv6Enabled bool
}




func DnsAssignedValuesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host_name"] = bindings.NewStringType()
	fieldNameMap["host_name"] = "HostName"
	fields["domain_name"] = bindings.NewStringType()
	fieldNameMap["domain_name"] = "DomainName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.dns_assigned_values", fields, reflect.TypeOf(DnsAssignedValues{}), fieldNameMap, validators)
}

func DnsConfigInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ip_addresses"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["ip_addresses"] = "IpAddresses"
	fields["search_domains"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["search_domains"] = "SearchDomains"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.dns_config_info", fields, reflect.TypeOf(DnsConfigInfo{}), fieldNameMap, validators)
}

func DhcpConfigInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv4_enabled"] = bindings.NewBooleanType()
	fieldNameMap["ipv4_enabled"] = "Ipv4Enabled"
	fields["ipv6_enabled"] = bindings.NewBooleanType()
	fieldNameMap["ipv6_enabled"] = "Ipv6Enabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.dhcp_config_info", fields, reflect.TypeOf(DhcpConfigInfo{}), fieldNameMap, validators)
}


