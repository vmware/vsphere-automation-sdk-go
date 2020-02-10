/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Network.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// The resource type for the vCenter network
const Network_RESOURCE_TYPE = "Network"


// The ``Type`` enumeration class defines the type of a vCenter Server network. The type of a network can be used to determine what features it supports and which APIs can be used to find more information about the network or change its configuration.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NetworkType string

const (
    // XXX: ESX based (created and managed on ESX)
	NetworkType_STANDARD_PORTGROUP NetworkType = "STANDARD_PORTGROUP"
    // XXX: vCenter based (create and managed through vCenter)
	NetworkType_DISTRIBUTED_PORTGROUP NetworkType = "DISTRIBUTED_PORTGROUP"
    // A network for whose configuration is managed outside of vSphere. The identifer and name of the network is made available through vSphere so that host and virtual machine virtual ethernet devices can connect to them.
	NetworkType_OPAQUE_NETWORK NetworkType = "OPAQUE_NETWORK"
)

func (t NetworkType) NetworkType() bool {
	switch t {
	case NetworkType_STANDARD_PORTGROUP:
		return true
	case NetworkType_DISTRIBUTED_PORTGROUP:
		return true
	case NetworkType_OPAQUE_NETWORK:
		return true
	default:
		return false
	}
}


// The ``FilterSpec`` class contains properties used to filter the results when listing networks (see Network#list). If multiple properties are specified, only networks matching all of the properties match the filter.
type NetworkFilterSpec struct {
    // Identifiers of networks that can match the filter.
	Networks map[string]bool
    // Names that networks must have to match the filter (see NetworkSummary#name).
	Names map[string]bool
    // Types that networks must have to match the filter (see NetworkSummary#type).
	Types map[NetworkType]bool
    // Folders that must contain the network for the network to match the filter.
	Folders map[string]bool
    // Datacenters that must contain the network for the network to match the filter.
	Datacenters map[string]bool
}

// The ``Summary`` class contains commonly used information about a network.
type NetworkSummary struct {
    // Identifier of the network.
	Network string
    // Name of the network.
	Name string
    // Type (STANDARD_PORTGROUP, DISTRIBUTED_PORTGROUP, OPAQUE_NETWORK) of the vCenter Server network.
	Type_ NetworkType
}



func networkListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(NetworkSummaryBindingType), reflect.TypeOf([]NetworkSummary{}))
}

func networkListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(NetworkFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
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
		map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func NetworkFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["networks"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Network"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["networks"] = "Networks"
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.network.type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP))), reflect.TypeOf(map[NetworkType]bool{})))
	fieldNameMap["types"] = "Types"
	fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["folders"] = "Folders"
	fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["datacenters"] = "Datacenters"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.network.filter_spec", fields, reflect.TypeOf(NetworkFilterSpec{}), fieldNameMap, validators)
}

func NetworkSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network"] = bindings.NewIdType([]string{"Network"}, "")
	fieldNameMap["network"] = "Network"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.network.type", reflect.TypeOf(NetworkType(NetworkType_STANDARD_PORTGROUP)))
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.network.summary", fields, reflect.TypeOf(NetworkSummary{}), fieldNameMap, validators)
}

