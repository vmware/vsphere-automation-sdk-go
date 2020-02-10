/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Datacenter.
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

// The resource type for the vCenter Datacenter
const Datacenter_RESOURCE_TYPE = "Datacenter"


// The ``CreateSpec`` class defines the information used to create a datacenter.
type DatacenterCreateSpec struct {
    // The name of the datacenter to be created.
	Name string
    // Datacenter folder in which the new datacenter should be created.
	Folder *string
}

// The ``FilterSpec`` class contains properties used to filter the results when listing datacenters (see Datacenter#list). If multiple properties are specified, only datacenters matching all of the properties match the filter.
type DatacenterFilterSpec struct {
    // Identifiers of datacenters that can match the filter.
	Datacenters map[string]bool
    // Names that datacenters must have to match the filter (see DatacenterInfo#name).
	Names map[string]bool
    // Folders that must contain the datacenters for the datacenter to match the filter.
	Folders map[string]bool
}

// The ``Summary`` class contains commonly used information about a datacenter in vCenter Server.
type DatacenterSummary struct {
    // Identifier of the datacenter.
	Datacenter string
    // Name of the datacenter.
	Name string
}

// The ``Info`` class contains information about a datacenter in vCenter Server.
type DatacenterInfo struct {
    // The name of the datacenter.
	Name string
    // The root datastore folder associated with the datacenter.
	DatastoreFolder string
    // The root host and cluster folder associated with the datacenter.
	HostFolder string
    // The root network folder associated with the datacenter.
	NetworkFolder string
    // The root virtual machine folder associated with the datacenter.
	VmFolder string
}



func datacenterCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(DatacenterCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datacenterCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"Datacenter"}, "")
}

func datacenterCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(DatacenterCreateSpecBindingType)
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
		map[string]int{"Error": 500,"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func datacenterDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datacenter"] = bindings.NewIdType([]string{"Datacenter"}, "")
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["datacenter"] = "Datacenter"
	fieldNameMap["force"] = "Force"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datacenterDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func datacenterDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["datacenter"] = bindings.NewIdType([]string{"Datacenter"}, "")
	fields["force"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["datacenter"] = "Datacenter"
	fieldNameMap["force"] = "Force"
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInUse": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func datacenterListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DatacenterFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datacenterListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(DatacenterSummaryBindingType), reflect.TypeOf([]DatacenterSummary{}))
}

func datacenterListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DatacenterFilterSpecBindingType))
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
		map[string]int{"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func datacenterGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datacenter"] = bindings.NewIdType([]string{"Datacenter"}, "")
	fieldNameMap["datacenter"] = "Datacenter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datacenterGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(DatacenterInfoBindingType)
}

func datacenterGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["datacenter"] = bindings.NewIdType([]string{"Datacenter"}, "")
	fieldNameMap["datacenter"] = "Datacenter"
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
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func DatacenterCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.datacenter.create_spec", fields, reflect.TypeOf(DatacenterCreateSpec{}), fieldNameMap, validators)
}

func DatacenterFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["datacenters"] = "Datacenters"
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["folders"] = "Folders"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.datacenter.filter_spec", fields, reflect.TypeOf(DatacenterFilterSpec{}), fieldNameMap, validators)
}

func DatacenterSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datacenter"] = bindings.NewIdType([]string{"Datacenter"}, "")
	fieldNameMap["datacenter"] = "Datacenter"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.datacenter.summary", fields, reflect.TypeOf(DatacenterSummary{}), fieldNameMap, validators)
}

func DatacenterInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["datastore_folder"] = bindings.NewIdType([]string{"Folder"}, "")
	fieldNameMap["datastore_folder"] = "DatastoreFolder"
	fields["host_folder"] = bindings.NewIdType([]string{"Folder"}, "")
	fieldNameMap["host_folder"] = "HostFolder"
	fields["network_folder"] = bindings.NewIdType([]string{"Folder"}, "")
	fieldNameMap["network_folder"] = "NetworkFolder"
	fields["vm_folder"] = bindings.NewIdType([]string{"Folder"}, "")
	fieldNameMap["vm_folder"] = "VmFolder"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.datacenter.info", fields, reflect.TypeOf(DatacenterInfo{}), fieldNameMap, validators)
}

