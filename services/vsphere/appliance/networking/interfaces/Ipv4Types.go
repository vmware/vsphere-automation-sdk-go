/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Ipv4.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package interfaces

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``Mode`` enumeration class defines different IPv4 address assignment modes. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type Ipv4Mode string

const (
    // The IPv4 address is automatically assigned by a DHCP server. This constant field was added in vSphere API 6.7.
	Ipv4Mode_DHCP Ipv4Mode = "DHCP"
    // The IPv4 address is static. This constant field was added in vSphere API 6.7.
	Ipv4Mode_STATIC Ipv4Mode = "STATIC"
    // The IPv4 protocol is not configured. This constant field was added in vSphere API 6.7.
	Ipv4Mode_UNCONFIGURED Ipv4Mode = "UNCONFIGURED"
)

func (m Ipv4Mode) Ipv4Mode() bool {
	switch m {
	case Ipv4Mode_DHCP:
		return true
	case Ipv4Mode_STATIC:
		return true
	case Ipv4Mode_UNCONFIGURED:
		return true
	default:
		return false
	}
}


// The ``Config`` class provides defines the IPv4 configuration of a network interface. This class was added in vSphere API 6.7.
type Ipv4Config struct {
    // The Address assignment mode. This property was added in vSphere API 6.7.
	Mode Ipv4Mode
    // The IPv4 address, for example, "10.20.80.191". This property was added in vSphere API 6.7.
	Address *string
    // The IPv4 CIDR prefix, for example, 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. This property was added in vSphere API 6.7.
	Prefix *int64
    // The IPv4 address of the default gateway. This configures the global default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces. This property was added in vSphere API 6.7.
	DefaultGateway *string
}

// The ``Info`` class defines current IPv4 configuration state of a network interface. This class was added in vSphere API 6.7.
type Ipv4Info struct {
    // The specified network interface is configurable or not. This property was added in vSphere API 6.7.
	Configurable bool
    // The Address assignment mode. This property was added in vSphere API 6.7.
	Mode Ipv4Mode
    // The IPv4 address, for example, "10.20.80.191". This property was added in vSphere API 6.7.
	Address *string
    // The IPv4 CIDR prefix, for example, 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. This property was added in vSphere API 6.7.
	Prefix *int64
    // The IPv4 address of the default gateway. This configures the global default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces. This property was added in vSphere API 6.7.
	DefaultGateway *string
}



func ipv4SetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["interface_name"] = bindings.NewIdType([]string{"com.vmware.appliance.networking.interfaces"}, "")
	fields["config"] = bindings.NewReferenceType(Ipv4ConfigBindingType)
	fieldNameMap["interface_name"] = "InterfaceName"
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4SetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipv4SetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["interface_name"] = bindings.NewIdType([]string{"com.vmware.appliance.networking.interfaces"}, "")
	fields["config"] = bindings.NewReferenceType(Ipv4ConfigBindingType)
	fieldNameMap["interface_name"] = "InterfaceName"
	fieldNameMap["config"] = "Config"
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Error": 500})
}

func ipv4GetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["interface_name"] = bindings.NewIdType([]string{"com.vmware.appliance.networking.interfaces"}, "")
	fieldNameMap["interface_name"] = "InterfaceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4GetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(Ipv4InfoBindingType)
}

func ipv4GetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["interface_name"] = bindings.NewIdType([]string{"com.vmware.appliance.networking.interfaces"}, "")
	fieldNameMap["interface_name"] = "InterfaceName"
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
		map[string]int{"NotFound": 404,"Error": 500})
}


func Ipv4ConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv4.mode", reflect.TypeOf(Ipv4Mode(Ipv4Mode_DHCP)))
	fieldNameMap["mode"] = "Mode"
	fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["address"] = "Address"
	fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["prefix"] = "Prefix"
	fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_gateway"] = "DefaultGateway"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("mode",
		map[string][]bindings.FieldData{
			"STATIC": []bindings.FieldData{
				bindings.NewFieldData("address", true),
				bindings.NewFieldData("prefix", true),
			},
			"DHCP": []bindings.FieldData{},
			"UNCONFIGURED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv4.config", fields, reflect.TypeOf(Ipv4Config{}), fieldNameMap, validators)
}

func Ipv4InfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["configurable"] = bindings.NewBooleanType()
	fieldNameMap["configurable"] = "Configurable"
	fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv4.mode", reflect.TypeOf(Ipv4Mode(Ipv4Mode_DHCP)))
	fieldNameMap["mode"] = "Mode"
	fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["address"] = "Address"
	fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["prefix"] = "Prefix"
	fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_gateway"] = "DefaultGateway"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("mode",
		map[string][]bindings.FieldData{
			"STATIC": []bindings.FieldData{
				bindings.NewFieldData("address", true),
				bindings.NewFieldData("prefix", true),
				bindings.NewFieldData("default_gateway", true),
			},
			"DHCP": []bindings.FieldData{
				bindings.NewFieldData("address", true),
				bindings.NewFieldData("prefix", true),
				bindings.NewFieldData("default_gateway", true),
			},
			"UNCONFIGURED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv4.info", fields, reflect.TypeOf(Ipv4Info{}), fieldNameMap, validators)
}

