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
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/appliance/networking/interfaces"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// ``InterfaceStatus`` enumeration class Defines interface status
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type InterfacesInterfaceStatus string

const (
    // The interface is down.
	InterfacesInterfaceStatus_down InterfacesInterfaceStatus = "down"
    // The interface is up.
	InterfacesInterfaceStatus_up InterfacesInterfaceStatus = "up"
)

func (i InterfacesInterfaceStatus) InterfacesInterfaceStatus() bool {
	switch i {
	case InterfacesInterfaceStatus_down:
		return true
	case InterfacesInterfaceStatus_up:
		return true
	default:
		return false
	}
}


// ``InterfaceInfo`` class Structure that defines properties and status of a network interface.
type InterfacesInterfaceInfo struct {
    // Interface name, for example, "nic0", "nic1".
	Name string
    // Interface status.
	Status InterfacesInterfaceStatus
    // MAC address. For example 00:0C:29:94:BB:5A.
	Mac string
    // IPv4 Address information. This property was added in vSphere API 6.7.
	Ipv4 *interfaces.Ipv4Info
    // IPv6 Address information. This property was added in vSphere API 6.7.
	Ipv6 *interfaces.Ipv6Info
}



func interfacesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(InterfacesInterfaceInfoBindingType), reflect.TypeOf([]InterfacesInterfaceInfo{}))
}

func interfacesListRestMetadata() protocol.OperationRestMetadata {
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

func interfacesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["interface_name"] = bindings.NewIdType([]string{"com.vmware.appliance.networking.interfaces"}, "")
	fieldNameMap["interface_name"] = "InterfaceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(InterfacesInterfaceInfoBindingType)
}

func interfacesGetRestMetadata() protocol.OperationRestMetadata {
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


func InterfacesInterfaceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.interface_status", reflect.TypeOf(InterfacesInterfaceStatus(InterfacesInterfaceStatus_down)))
	fieldNameMap["status"] = "Status"
	fields["mac"] = bindings.NewStringType()
	fieldNameMap["mac"] = "Mac"
	fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv4InfoBindingType))
	fieldNameMap["ipv4"] = "Ipv4"
	fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv6InfoBindingType))
	fieldNameMap["ipv6"] = "Ipv6"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.interfaces.interface_info", fields, reflect.TypeOf(InterfacesInterfaceInfo{}), fieldNameMap, validators)
}

