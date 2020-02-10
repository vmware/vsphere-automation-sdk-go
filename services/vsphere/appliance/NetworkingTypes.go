/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Networking.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/appliance/networking"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/appliance/networking/dns"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/appliance/networking/interfaces"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``DNSInfo`` class contains information about the DNS configuration of a virtual appliance. This class was added in vSphere API 6.7.
type NetworkingDNSInfo struct {
    // DNS mode. This property was added in vSphere API 6.7.
	Mode NetworkingDNSInfoDNSMode
    // Hostname. This property was added in vSphere API 6.7.
	Hostname string
    // Servers. This property was added in vSphere API 6.7.
	Servers []string
}

// The ``DNSMode`` enumeration class describes the source of DNS servers. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NetworkingDNSInfoDNSMode string

const (
    // The DNS servers addresses are obtained from a DHCP server. This constant field was added in vSphere API 6.7.
	NetworkingDNSInfoDNSMode_DHCP NetworkingDNSInfoDNSMode = "DHCP"
    // The DNS servers addresses are specified explicitly. This constant field was added in vSphere API 6.7.
	NetworkingDNSInfoDNSMode_STATIC NetworkingDNSInfoDNSMode = "STATIC"
)

func (d NetworkingDNSInfoDNSMode) NetworkingDNSInfoDNSMode() bool {
	switch d {
	case NetworkingDNSInfoDNSMode_DHCP:
		return true
	case NetworkingDNSInfoDNSMode_STATIC:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information about the network configuration of a virtual appliance. This class was added in vSphere API 6.7.
type NetworkingInfo struct {
    // DNS configuration. This property was added in vSphere API 6.7.
	Dns NetworkingDNSInfo
    // Interface configuration as a key-value map where key is a network interface name, for example, "nic0". This property was added in vSphere API 6.7.
	Interfaces map[string]networking.InterfacesInterfaceInfo
}

// The ``UpdateSpec`` class describes whether to enable or disable ipv6 on interfaces. This class was added in vSphere API 6.7.
type NetworkingUpdateSpec struct {
    // IPv6 Enabled or not. This property was added in vSphere API 6.7.
	Ipv6Enabled *bool
}

type NetworkingChangeSpec struct {
    // New hostname to assign to the management network of vCenter appliance. This property was added in vSphere API 6.7.3.
	Hostname string
    // vCenter Server SSO administrator username. This property was added in vSphere API 6.7.3.
	SSOUser string
    // vCenter Server SSO administrator Password. This property was added in vSphere API 6.7.3.
	SSOPassword string
    // DNS Configuration to set for the machine. This property was added in vSphere API 6.7.3.
	Dns *dns.ServersDNSServerConfig
    // IPv4 Configuration to set for the machine. This property was added in vSphere API 6.7.3.
	Ipv4 *interfaces.Ipv4Config
    // IPv6 Configuration to set for the machine. This property was added in vSphere API 6.7.3.
	Ipv6 *interfaces.Ipv6Config
}



func networkingGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NetworkingInfoBindingType)
}

func networkingGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500})
}

func networkingUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(NetworkingUpdateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func networkingUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(NetworkingUpdateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
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
		map[string]int{"Error": 500})
}

func networkingResetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingResetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func networkingResetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		map[string]int{"Error": 500})
}

func networkingChangeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(NetworkingChangeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingChangeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func networkingChangeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(NetworkingChangeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
		map[string]int{"Unsupported": 400,"InvalidArgument": 400,"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}


func NetworkingDNSInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.DNS_info.DNS_mode", reflect.TypeOf(NetworkingDNSInfoDNSMode(NetworkingDNSInfoDNSMode_DHCP)))
	fieldNameMap["mode"] = "Mode"
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["servers"] = "Servers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.DNS_info", fields, reflect.TypeOf(NetworkingDNSInfo{}), fieldNameMap, validators)
}

func NetworkingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dns"] = bindings.NewReferenceType(NetworkingDNSInfoBindingType)
	fieldNameMap["dns"] = "Dns"
	fields["interfaces"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.appliance.networking.interfaces"}, ""), bindings.NewReferenceType(networking.InterfacesInterfaceInfoBindingType),reflect.TypeOf(map[string]networking.InterfacesInterfaceInfo{}))
	fieldNameMap["interfaces"] = "Interfaces"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.info", fields, reflect.TypeOf(NetworkingInfo{}), fieldNameMap, validators)
}

func NetworkingUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipv6_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ipv6_enabled"] = "Ipv6Enabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.update_spec", fields, reflect.TypeOf(NetworkingUpdateSpec{}), fieldNameMap, validators)
}

func NetworkingChangeSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["SSO_user"] = bindings.NewStringType()
	fieldNameMap["SSO_user"] = "SSOUser"
	fields["SSO_password"] = bindings.NewSecretType()
	fieldNameMap["SSO_password"] = "SSOPassword"
	fields["dns"] = bindings.NewOptionalType(bindings.NewReferenceType(dns.ServersDNSServerConfigBindingType))
	fieldNameMap["dns"] = "Dns"
	fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv4ConfigBindingType))
	fieldNameMap["ipv4"] = "Ipv4"
	fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv6ConfigBindingType))
	fieldNameMap["ipv6"] = "Ipv6"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.change_spec", fields, reflect.TypeOf(NetworkingChangeSpec{}), fieldNameMap, validators)
}

